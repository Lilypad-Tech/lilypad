package jobcreator

import (
	"context"
	"fmt"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
)

type JobCreatorController struct {
	solverClient *solver.SolverClient
	options      JobCreatorOptions
	web3SDK      *web3.Web3SDK
	web3Events   *web3.EventChannels
	loop         *system.ControlLoop
	log          *system.ServiceLogger
}

// the background "even if we have not heard of an event" loop
// i.e. things will not wait 10 seconds - the control loop
// reacts to events in the system - this 10 second background
// loop is just for in case we miss any events
const CONTROL_LOOP_INTERVAL = 10 * time.Second

func NewJobCreatorController(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
) (*JobCreatorController, error) {
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Offer.Services.Solver)
	if err != nil {
		return nil, err
	}

	solverClient, err := solver.NewSolverClient(http.ClientOptions{
		URL:        solverUrl,
		PrivateKey: options.Web3.PrivateKey,
	})
	if err != nil {
		return nil, err
	}

	controller := &JobCreatorController{
		solverClient: solverClient,
		options:      options,
		web3SDK:      web3SDK,
		web3Events:   web3.NewEventChannels(),
		log:          system.NewServiceLogger(system.JobCreatorService),
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
	controller.log.Info("add job offer", offer)
	return controller.solverClient.AddJobOffer(offer)
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
		// we need to agree to the deal now we've heard about it
		if ev.EventType == solver.DealAdded {
			if ev.Deal == nil {
				controller.log.Error("solver event", fmt.Errorf("RP received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.JobCreator != controller.web3SDK.GetAddress().String() {
				return
			}

			solver.ServiceLogSolverEvent(system.JobCreatorService, ev)

			// trigger the solver
			controller.loop.Trigger()
		}
	})
	return nil
}

func (controller *JobCreatorController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		deal, err := controller.solverClient.GetDeal(ev.DealId.String())
		if err != nil {
			controller.log.Error("error getting deal", err)
			return
		}
		if deal.ResourceProvider != controller.web3SDK.GetAddress().String() {
			return
		}
		controller.log.Info("StorageDealStateChange", ev)
		controller.loop.Trigger()
	})
	return nil
}

func (controller *JobCreatorController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error)
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
	err = controller.solverClient.Start(ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
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
	// load the deals that are in DealNegotiating
	// and do not have a TransactionsResourceProvider.Agree tx
	matchedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			JobCreator: controller.web3SDK.GetAddress().String(),
			State:      "DealNegotiating",
		},
		// this is where the solver has found us a match and we need to agree to it
		func(dealContainer data.DealContainer) bool {
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
		controller.log.Info("agree", dealContainer)
		tx, err := controller.web3SDK.Agree(dealContainer.Deal)
		if err != nil {
			// TODO: we need a way of deciding based on certain classes of error what happens
			// some will be retryable - otherwise will be fatal
			// we need a way to exit a job loop as a baseline
			controller.log.Error("error calling agree tx for deal", err)
			continue
		}
		controller.log.Info("agree tx", tx)

		// we have agreed to the deal so we need to update the tx in the solver
		_, err = controller.solverClient.UpdateTransactionsJobCreator(dealContainer.ID, data.DealTransactionsJobCreator{
			Agree: tx,
		})
		if err != nil {
			// TODO: we need a way of deciding based on certain classes of error what happens
			// some will be retryable - otherwise will be fatal
			// we need a way to exit a job loop as a baseline
			controller.log.Error("error calling agree tx for deal", err)
			continue
		}
		controller.log.Info("updated deal with agree tx", tx)
	}

	return err
}
