package jobcreator

import (
	"context"
	"fmt"
	"time"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/http"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/metricsDashboard"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/storage"
	"go.opentelemetry.io/otel/trace"
)

// JobOfferSubscriber is a function that gets called when a job offer is updated
type JobOfferSubscriber func(offer data.JobOfferContainer)

// JobOfferSubscription represents a subscription to job offer updates
type JobOfferSubscription struct {
	// The callback function to call when a job offer is updated
	Callback JobOfferSubscriber
}

// JobCreatorController manages a single job identified by jobID.
// It handles all aspects of job processing including subscriptions,
// deal management, and result handling for this specific job.
type JobCreatorController struct {
	solverClient          *solver.SolverClient
	options               JobCreatorOptions
	web3SDK               *web3.Web3SDK
	web3Events            *web3.EventChannels
	loop                  *system.ControlLoop
	log                   *system.ServiceLogger
	jobOfferSubscriptions map[string]JobOfferSubscription
	tracer                trace.Tracer
	jobID                 string // The single job this controller instance manages
}

// the background "even if we have not heard of an event" loop
// i.e. things will not wait 10 seconds - the control loop
// reacts to events in the system - this 10 second background
// loop is just for in case we miss any events
const CONTROL_LOOP_INTERVAL = 10 * time.Second

func NewJobCreatorController(
	jobID string,
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
	tracer trace.Tracer,
) (*JobCreatorController, error) {
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Offer.Services.Solver)
	if err != nil {
		return nil, err
	}

	solverClient, err := solver.NewSolverClient(
		http.ClientOptions{
			URL:           solverUrl,
			PrivateKey:    options.Web3.PrivateKey,
			Type:          "JobCreator",
			PublicAddress: web3SDK.GetAddress().String(),
		})
	if err != nil {
		return nil, err
	}

	metricsDashboard.Init(options.Offer.Services.APIHost)

	controller := &JobCreatorController{
		solverClient:          solverClient,
		options:               options,
		web3SDK:               web3SDK,
		web3Events:            web3.NewEventChannels(),
		log:                   system.NewServiceLogger(system.JobCreatorService),
		jobOfferSubscriptions: make(map[string]JobOfferSubscription),
		tracer:                tracer,
		jobID:                 jobID,
	}
	return controller, nil
}

/*
 *
 *
 *

 Public API

 *
 *
 *
*/

func (controller *JobCreatorController) AddJobOffer(offer data.JobOffer) (data.JobOfferContainer, error) {
	controller.log.Debug("add job offer", offer)
	container, err := controller.solverClient.AddJobOffer(offer)
	if err != nil {
		return container, err
	}

	// Set the controller's jobID to track this specific job
	controller.jobID = container.JobOffer.ID
	controller.log.Debug("Set controller jobID", map[string]interface{}{
		"jobID": controller.jobID,
	})

	return container, nil
}

// SubscribeToJobOfferUpdates subscribes to job offer updates for this controller's job
// Returns a function that can be called to unsubscribe
func (controller *JobCreatorController) SubscribeToJobOfferUpdates(sub JobOfferSubscriber) func() {
	controller.jobOfferSubscriptions[controller.jobID] = JobOfferSubscription{
		Callback: sub,
	}

	return func() {
		delete(controller.jobOfferSubscriptions, controller.jobID)
	}
}

/*
*
*
*

	Setup

*
*
*
*/
func (controller *JobCreatorController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(ev solver.SolverEvent) {
		// First check if this event is for our specific job
		if controller.jobID != "" {
			if ev.Deal != nil && ev.Deal.JobOffer != controller.jobID {
				// Skip events for other jobs
				return
			}
			if ev.JobOffer != nil && ev.JobOffer.JobOffer.ID != controller.jobID {
				// Skip events for other jobs
				return
			}
		}

		if ev.EventType == "DealStateUpdated" {
			metricsDashboard.TrackDeal(metricsDashboard.DealPayload{
				ID:               ev.Deal.ID,
				JobID:            ev.Deal.JobOffer,
				JobCreator:       ev.Deal.JobCreator,
				ResourceProvider: ev.Deal.ResourceProvider,
			})
		}

		switch ev.EventType {
		case solver.DealAdded:
			if ev.Deal == nil {
				controller.log.Error("solver event", fmt.Errorf("JC received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.JobCreator != controller.web3SDK.GetAddress().String() {
				return
			}

			controller.log.Debug("Received deal event", map[string]interface{}{
				"dealID": ev.Deal.ID,
				"jobID":  controller.jobID,
			})

			solver.ServiceLogSolverEvent(system.JobCreatorService, ev)
			controller.loop.Trigger()

		case solver.JobOfferStateUpdated:
			if ev.JobOffer == nil {
				controller.log.Error("solver event", fmt.Errorf("RP received nil job offer"))
				return
			}
			metricsDashboard.TrackJobOfferUpdate(*ev.JobOffer)

			// Make a copy of the subscriptions to avoid modification during iteration
			for jobID, sub := range controller.jobOfferSubscriptions {
				// If JobID is empty (global subscription) or matches the job offer ID, send the update
				if jobID == "" || jobID == ev.JobOffer.JobOffer.ID {
					sub.Callback(*ev.JobOffer)
				}
			}
		}
	})
	return nil
}

func (controller *JobCreatorController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		deal, err := controller.solverClient.GetDeal(ev.DealId)
		if err != nil {
			controller.log.Error("error getting deal", err)
			return
		}
		if deal.JobCreator != controller.web3SDK.GetAddress().String() {
			return
		}
		controller.log.Debug("StorageDealStateChange", data.GetAgreementStateString(ev.State))
		system.DumpObjectDebug(ev)
		controller.loop.Trigger()
	})
	return nil
}

func (controller *JobCreatorController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error, 1)
	err := controller.subscribeToSolver()
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.subscribeToWeb3()
	if err != nil {
		errorChan <- err
		return errorChan
	}

	// this connects the websocket client
	err = controller.solverClient.Start(ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.web3Events.Start(ctx, cm, controller.web3SDK)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	controller.loop = system.NewControlLoop(
		system.JobCreatorService,
		ctx,
		CONTROL_LOOP_INTERVAL,
		func() error {
			err := controller.solve()
			if err != nil {
				errorChan <- err
			}
			return err
		},
	)

	err = controller.loop.Start(true)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	return errorChan
}

/*
 *
 *
 *

 Solve

 *
 *
 *
*/

func (controller *JobCreatorController) solve() error {
	controller.log.Debug("solving", "")
	err := controller.agreeToDeals()
	if err != nil {
		return err
	}
	err = controller.checkResults()
	if err != nil {
		return err
	}
	return nil
}

/*
 *
 *
 *

 Agree to deals

 *
 *
 *
*/

// list the deals we have been assigned to that we have not yet posted and agree tx to the contract for
func (controller *JobCreatorController) agreeToDeals() error {
	// If no jobID is set, we shouldn't process any deals
	if controller.jobID == "" {
		return nil
	}

	matchedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			JobCreator: controller.web3SDK.GetAddress().String(),
			State:      "DealNegotiating",
		},
		func(dealContainer data.DealContainer) bool {
			// Only agree to deals for our specific job
			if dealContainer.Deal.JobOffer.ID != controller.jobID {
				controller.log.Debug("Skipping deal - wrong jobID", map[string]interface{}{
					"dealJobID": dealContainer.Deal.JobOffer.ID,
					"ourJobID":  controller.jobID,
				})
				return false
			}

			return dealContainer.Transactions.JobCreator.Agree == ""
		},
	)
	if err != nil {
		return err
	}
	if len(matchedDeals) <= 0 {
		return nil
	}

	// map over the deals and agree to them
	for _, dealContainer := range matchedDeals {
		controller.log.Debug("agree", dealContainer)
		txHash, err := controller.web3SDK.Agree(dealContainer.Deal)
		if err != nil {
			// TODO: error handling - is it terminal or retryable?
			controller.log.Error("error calling agree tx for deal", err)
			continue
		}
		controller.log.Debug("agree tx", txHash)

		// we have agreed to the deal so we need to update the tx in the solver
		_, err = controller.solverClient.UpdateTransactionsJobCreator(dealContainer.ID, data.DealTransactionsJobCreator{
			Agree: txHash,
		})
		if err != nil {
			// TODO: error handling - is it terminal or retryable?
			controller.log.Error("error adding agree tx hash for deal", err)
			continue
		}
		controller.log.Debug("updated deal with agree tx", txHash)
	}

	return nil
}

// list the deals that have results posted but we have not yet checked
// we do this synchronously to prevent us racing with large result sets
// also we are the client so have a lower chance of there being a chunky backlog
func (controller *JobCreatorController) checkResults() error {
	controller.log.Debug("Checking results for jobID", controller.jobID)

	// If no jobID is set, we shouldn't process any results
	if controller.jobID == "" {
		return nil
	}

	completedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			JobCreator: controller.web3SDK.GetAddress().String(),
			State:      "ResultsSubmitted",
		},
		func(dealContainer data.DealContainer) bool {
			// First check if this deal belongs to our job
			if dealContainer.Deal.JobOffer.ID != controller.jobID {
				controller.log.Debug("Skipping deal - wrong jobID", map[string]interface{}{
					"dealJobID": dealContainer.Deal.JobOffer.ID,
					"ourJobID":  controller.jobID,
				})
				return false
			}

			// Then check if we haven't processed it yet
			match := dealContainer.Transactions.JobCreator.AcceptResult == "" &&
				dealContainer.Transactions.JobCreator.CheckResult == ""

			controller.log.Debug("Filtering deal", map[string]interface{}{
				"dealID":          dealContainer.ID,
				"dealJobOfferID":  dealContainer.Deal.JobOffer.ID,
				"controllerJobID": controller.jobID,
				"match":           match,
			})

			return match
		},
	)
	if err != nil {
		return err
	}
	if len(completedDeals) <= 0 {
		return nil
	}

	// Process each deal
	for _, dealContainer := range completedDeals {

		result, err := controller.solverClient.GetResult(dealContainer.ID)
		if err != nil {
			controller.log.Debug("failed to get result metadata", map[string]interface{}{
				"dealID": dealContainer.ID,
			})
			controller.log.Error("failed to get result", err)
			continue // Continue to next deal instead of returning
		}

		if result.Error != "" {
			err := fmt.Errorf("result contains error: %s", result.Error)
			controller.log.Debug("result error metadata", map[string]interface{}{
				"dealID": dealContainer.ID,
			})
			controller.log.Error("result contains error", err)
			continue // Continue to next deal
		}

		if result.DataID == "" {
			err := fmt.Errorf("result missing DataID for deal %s", dealContainer.ID)
			controller.log.Debug("missing DataID metadata", map[string]interface{}{
				"dealID": dealContainer.ID,
			})
			controller.log.Error("result missing DataID", err)
			continue // Continue to next deal
		}

		err = controller.downloadResult(dealContainer)
		if err != nil {
			controller.log.Error("failed to download results", err)
			continue // Continue to next deal
		}

		err = controller.acceptResult(dealContainer)
		if err != nil {
			controller.log.Error("failed to accept results", err)
			continue // Continue to next deal
		}
	}

	return nil
}

func (controller *JobCreatorController) downloadResult(dealContainer data.DealContainer) error {
	err := controller.solverClient.DownloadResultFiles(dealContainer.ID, solver.GetDownloadsFilePath(dealContainer.ID))
	if err != nil {
		return fmt.Errorf("error downloading results for deal: %s", err.Error())
	}

	controller.log.Debug("Downloaded results for job", solver.GetDownloadsFilePath(dealContainer.ID))

	// TODO: activate the mediation check here
	err = controller.acceptResult(dealContainer)
	if err != nil {
		controller.log.Error("failed to accept results", err)
		return err
	}

	// work out if we should check or accept the results
	// if controller.options.Mediation.CheckResultsPercentage >= rand.Intn(100) {
	// 	err = controller.checkResult(dealContainer)

	// 	if err != nil {
	// 		// TODO: error handling - is it terminal or retryable?
	// 		controller.log.Error("error checking deal results", err)
	// 		return nil
	// 	}

	// 	controller.log.Debug("Checked results for job", dealContainer.ID)
	// } else {
	// 	err = controller.acceptResult(dealContainer)

	// 	if err != nil {
	// 		// TODO: error handling - is it terminal or retryable?
	// 		controller.log.Error("error accepting deal results", err)
	// 		return nil
	// 	}

	// 	controller.log.Debug("Accepted results for job", dealContainer.ID)
	// }
	return nil
}

func (controller *JobCreatorController) acceptResult(deal data.DealContainer) error {
	controller.log.Debug("Accepting results for job", deal.ID)
	// !TODO MAINNET: get txHash for accepting results on chain or get this flow
	txHash := "0x"

	// we have agreed to the deal so we need to update the tx in the solver
	_, err := controller.solverClient.UpdateTransactionsJobCreator(deal.ID, data.DealTransactionsJobCreator{
		AcceptResult: txHash,
	})
	if err != nil {
		return fmt.Errorf("error adding AcceptResult tx hash for deal: %s", err.Error())
	}

	return nil
}
