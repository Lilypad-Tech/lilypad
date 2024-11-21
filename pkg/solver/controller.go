package solver

import (
	"context"
	"fmt"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/metricsDashboard"
	"github.com/lilypad-tech/lilypad/pkg/solver/matcher"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/mediation"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/storage"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// add an enum for various types of event
type SolverEventType string

const (
	JobOfferAdded                       SolverEventType = "JobOfferAdded"
	ResourceOfferAdded                  SolverEventType = "ResourceOfferAdded"
	ResourceOfferRemoved                SolverEventType = "ResourceOfferRemoved"
	DealAdded                           SolverEventType = "DealAdded"
	JobOfferStateUpdated                SolverEventType = "JobOfferStateUpdated"
	ResourceOfferStateUpdated           SolverEventType = "ResourceOfferStateUpdated"
	DealStateUpdated                    SolverEventType = "DealStateUpdated"
	DealMediatorUpdated                 SolverEventType = "DealMediatorUpdated"
	ResourceProviderTransactionsUpdated SolverEventType = "ResourceProviderTransactionsUpdated"
	JobCreatorTransactionsUpdated       SolverEventType = "JobCreatorTransactionsUpdated"
	MediatorTransactionsUpdated         SolverEventType = "MediatorTransactionsUpdated"
)

type SolverEvent struct {
	EventType     SolverEventType              `json:"event_type"`
	JobOffer      *data.JobOfferContainer      `json:"job_offer"`
	ResourceOffer *data.ResourceOfferContainer `json:"resource_offer"`
	Deal          *data.DealContainer          `json:"deal"`
}

type SolverController struct {
	web3SDK         *web3.Web3SDK
	web3Events      *web3.EventChannels
	store           store.SolverStore
	loop            *system.ControlLoop
	solverEventSubs []func(SolverEvent)
	options         SolverOptions
	log             *system.ServiceLogger
	tracer          trace.Tracer
	meter           metric.Meter
}

// the background "even if we have not heard of an event" loop
// i.e. things will not wait 10 seconds - the control loop
// reacts to events in the system - this 10 second background
// loop is just for in case we miss any events
const CONTROL_LOOP_INTERVAL = 10 * time.Second
const REQUIRED_BALANCE_IN_WEI = 0.0006

func NewSolverController(
	web3SDK *web3.Web3SDK,
	store store.SolverStore,
	options SolverOptions,
	tracer trace.Tracer,
	meter metric.Meter,
) (*SolverController, error) {
	controller := &SolverController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
		store:      store,
		options:    options,
		log:        system.NewServiceLogger(system.SolverService),
		tracer:     tracer,
		meter:      meter,
	}
	return controller, nil
}

func (controller *SolverController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error, 1)
	// get the local subscriptions setup
	err := controller.subscribeToWeb3()
	if err != nil {
		errorChan <- err
		return errorChan
	}

	// activate the web3 event listeners
	log.Debug().Msgf("controller.web3Events.Start")
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	// make sure we are registered as a solver
	// so that users can lookup our URL
	log.Debug().Msgf("controller.registerAsSolver")
	err = controller.registerAsSolver()
	if err != nil {
		errorChan <- err
		return errorChan
	}

	controller.loop = system.NewControlLoop(
		system.SolverService,
		ctx,
		CONTROL_LOOP_INTERVAL,
		func() error {
			err := controller.solve(ctx)
			if err != nil {
				errorChan <- err
			}
			return err
		},
	)
	log.Debug().Msgf("controller.loop.Start")
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

 Events

 *
 *
 *
*/

// * get the deal id
// * see if we have the deal locally
// * update the deal state locally
func (controller *SolverController) subscribeToWeb3() error {

	// change the deal state
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		_, err := controller.updateDealState(ev.DealId, ev.State)
		if err != nil {
			controller.log.Error("error updating deal state", err)
			return
		}
		controller.log.Info("StorageDealStateChange", data.GetAgreementStateString(ev.State))
		system.DumpObjectDebug(ev)
		// update the store with the state change
		controller.loop.Trigger()
	})

	// update the mediator
	controller.web3Events.Mediation.SubscribeMediationRequested(func(ev mediation.MediationMediationRequested) {
		controller.log.Info("MediationMediationRequested", "")
		system.DumpObjectDebug(ev)
		_, err := controller.updateDealMediator(ev.DealId, ev.Mediator.String())
		if err != nil {
			controller.log.Error("error updating deal state", err)
			return
		}

		// update the store with the state change
		controller.loop.Trigger()
	})

	return nil
}

// return a new event channel that will hear about events
// coming out of this controller
func (controller *SolverController) subscribeEvents(handler func(SolverEvent)) {
	controller.solverEventSubs = append(controller.solverEventSubs, handler)
}

func (controller *SolverController) reactToEvent(ev SolverEvent) {
	// both of these should trigger a solve
	if ev.EventType == ResourceOfferAdded || ev.EventType == JobOfferAdded {
		controller.loop.Trigger()
	}
}

// write the given event to all generated event channels
func (controller *SolverController) writeEvent(ev SolverEvent) {
	controller.reactToEvent(ev)
	for _, handler := range controller.solverEventSubs {
		handler(ev)
	}
}

/*
 *
 *
 *

 Register

 *
 *
 *
*/

func (controller *SolverController) registerAsSolver() error {
	selfAddress := controller.web3SDK.GetAddress()
	solverType, err := data.GetServiceType("Solver")
	if err != nil {
		return err
	}

	log.Debug().Msgf("GetUser with selfAddress: %s", selfAddress.String())
	selfUser, err := controller.web3SDK.GetUser(selfAddress)
	if err != nil {
		return err
	}

	// TODO: check the other props and call update if they have changed
	log.Debug().Msgf("selfUser.Url: %s", selfUser.Url)
	log.Debug().Msgf("controller.options.Server.URL: %s", controller.options.Server.URL)
	if selfUser.Url != controller.options.Server.URL {
		controller.log.Info("url change", fmt.Sprintf("solver will be updated because URL has changed: %s %s != %s", selfAddress.String(), selfUser.Url, controller.options.Server.URL))
		err = controller.web3SDK.UpdateUser(
			"",
			controller.options.Server.URL,
			[]uint8{solverType},
		)
		if err != nil {
			return err
		}
	} else {
		controller.log.Info("url same", fmt.Sprintf("solver url already correct: %s %s", selfAddress.String(), controller.options.Server.URL))
	}

	existingSolvers, err := controller.web3SDK.GetSolverAddresses()
	if err != nil {
		return err
	}
	foundSolver := false
	for _, existingSolver := range existingSolvers {
		if existingSolver.String() == selfAddress.String() {
			controller.log.Info("solver exists", selfAddress.String())
			foundSolver = true
			break
		}
	}
	if !foundSolver {
		controller.log.Info("solver registering", "")
		// add the solver to the storage contract
		err = controller.web3SDK.AddUserToList(
			solverType,
		)
		if err != nil {
			return err
		}
		controller.log.Info("solver registered", selfAddress.String())
	}
	return nil
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

func (controller *SolverController) solve(ctx context.Context) error {
	ctx, span := controller.tracer.Start(ctx, "solve")
	defer span.End()

	// find out which deals we can make from matching the offers
	deals, err := matcher.GetMatchingDeals(ctx, controller.store, controller.updateJobOfferState, controller.tracer, controller.meter)
	if err != nil {
		span.SetStatus(codes.Error, "get matching deals failed")
		span.RecordError(err)
		return err
	}
	span.SetAttributes(attribute.KeyValue{
		Key:   "deal_ids",
		Value: attribute.StringSliceValue(data.GetDealIDs(deals)),
	})

	// loop over each of the deals add add them to the store and emit events
	span.AddEvent("add_deals.start")
	for _, deal := range deals {
		_, err := controller.addDeal(ctx, deal)
		if err != nil {
			return err
		}
	}
	span.AddEvent("add_deals.done")

	span.AddEvent("report_deal_metrics.start")
	storedDeals, err := controller.store.GetDealsAll()
	if err != nil {
		span.SetStatus(codes.Error, "get all deals failed")
		span.RecordError(err)
		return err
	}
	err = reportDealMetrics(ctx, controller.meter, storedDeals)
	if err != nil {
		span.SetStatus(codes.Error, "report deal metrics failed")
		span.RecordError(err)
	}
	span.AddEvent("report_deal_metrics.done")

	return nil
}

/*
*
*
*

# Add Handlers

*
*
*
*/
func (controller *SolverController) addJobOffer(jobOffer data.JobOffer) (*data.JobOfferContainer, error) {
	id, err := data.GetJobOfferID(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOffer.ID = id

	controller.log.Info("add job offer", jobOffer)

	ret, err := controller.store.AddJobOffer(data.GetJobOfferContainer(jobOffer))
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: JobOfferAdded,
		JobOffer:  ret,
	})
	return ret, nil
}

func (controller *SolverController) addResourceOffer(resourceOffer data.ResourceOffer) (*data.ResourceOfferContainer, error) {
	id, err := data.GetResourceOfferID(resourceOffer)
	if err != nil {
		return nil, err
	}
	resourceOffer.ID = id

	// Check the resource provider's ETH balance
	balance, err := controller.web3SDK.GetBalance(resourceOffer.ResourceProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ETH balance for resource provider: %v", err)
	}
	// Convert InstructionPrice from ETH to Wei
	requiredBalanceWei := web3.EtherToWei(REQUIRED_BALANCE_IN_WEI) // 0.0006 based on the required balance for a job

	// If the balance is less than the required balance, don't add the resource offer
	if balance.Cmp(requiredBalanceWei) < 0 {
		err := fmt.Errorf("address %s doesn't have enough ETH balance. The required balance is %s but current balance is %s", resourceOffer.ResourceProvider, requiredBalanceWei, balance)
		controller.log.Error("ETH balance check failed", err)
		return nil, nil
	}

	// required LP balance
	requiredBalanceLp := web3.EtherToWei(float64(resourceOffer.DefaultPricing.InstructionPrice)) // based on the required LP balance for a job
	balanceLp, err := controller.web3SDK.GetLPBalance(resourceOffer.ResourceProvider)
	if err != nil {
		err := fmt.Errorf("failed to retrieve LP balance for resource provider: %v", err)
		controller.log.Error("LP Balance error", err)
		return nil, nil
	}
	if balanceLp.Cmp(requiredBalanceLp) < 0 {
		err := fmt.Errorf("address %s doesn't have enough LP balance. The required balance is %s but current balance is %s", resourceOffer.ResourceProvider, requiredBalanceLp, balanceLp)
		controller.log.Error("LP balance check failed", err)
		return nil, nil
	}

	controller.log.Info("add resource offer", resourceOffer)

	metricsDashboard.TrackNodeInfo(resourceOffer)

	ret, err := controller.store.AddResourceOffer(data.GetResourceOfferContainer(resourceOffer))
	if err != nil {
		return nil, err
	}

	controller.writeEvent(SolverEvent{
		EventType:     ResourceOfferAdded,
		ResourceOffer: ret,
	})
	return ret, nil
}

func (controller *SolverController) removeResourceOfferByResourceProvider(ID string) error {
	controller.log.Info("remove resource offer", ID)
	resourceOffers, err := controller.store.GetResourceOffers(store.GetResourceOffersQuery{
		ResourceProvider: ID,
	})
	if err != nil {
		return err
	}

	if len(resourceOffers) == 0 {
		return nil
	}

	err = controller.store.RemoveResourceOffer(resourceOffers[0].ID)
	if err != nil {
		return err
	}

	controller.writeEvent(SolverEvent{
		EventType:     ResourceOfferRemoved,
		ResourceOffer: nil,
	})
	return nil
}

func (controller *SolverController) addDeal(ctx context.Context, deal data.Deal) (*data.DealContainer, error) {
	ctx, span := controller.tracer.Start(ctx, "add_deal")
	defer span.End()

	span.AddEvent("data.get_deal_id.start")
	id, err := data.GetDealID(deal)
	if err != nil {
		span.SetStatus(codes.Error, "get deal ID failed")
		span.RecordError(err)
		return nil, err
	}
	deal.ID = id
	span.SetAttributes(attribute.String("deal.id", deal.ID))
	span.AddEvent("data.get_deal_id.done")

	controller.log.Info("add deal", deal)

	span.AddEvent("store.add_deal.start")
	ret, err := controller.store.AddDeal(data.GetDealContainer(deal))
	if err != nil {
		span.SetStatus(codes.Error, "add deal to store failed")
		span.RecordError(err)
		return nil, err
	}
	span.AddEvent("store.add_deal.done")

	span.AddEvent("write_event.start")
	controller.writeEvent(SolverEvent{
		EventType: DealAdded,
		Deal:      ret,
	})
	span.AddEvent("write_event.done")

	span.AddEvent("update_job_offer_state.start")
	_, err = controller.updateJobOfferState(ret.JobOffer, ret.ID, ret.State)
	if err != nil {
		span.SetStatus(codes.Error, "updated job offer state failed")
		span.RecordError(err)
		return nil, err
	}
	span.AddEvent("update_job_offer_state.done")

	span.AddEvent("update_resource_offer_state.start")
	_, err = controller.updateResourceOfferState(ret.ResourceOffer, ret.ID, ret.State)
	if err != nil {
		span.SetStatus(codes.Error, "updated resource offer state failed")
		span.RecordError(err)
		return nil, err
	}
	span.AddEvent("update_resource_offer_state.done")

	return ret, nil
}

/*
*
*
*

# Update Handlers

*
*
*
*/
func (controller *SolverController) updateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	controller.log.Info("update job offer", fmt.Sprintf("%s %s", id, data.GetAgreementStateString(state)))

	ret, err := controller.store.UpdateJobOfferState(id, dealID, state)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: JobOfferStateUpdated,
		JobOffer:  ret,
	})
	return ret, nil
}

func (controller *SolverController) updateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	controller.log.Info("update resource offer", fmt.Sprintf("%s %s", id, data.GetAgreementStateString(state)))

	ret, err := controller.store.UpdateResourceOfferState(id, dealID, state)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType:     ResourceOfferStateUpdated,
		ResourceOffer: ret,
	})
	return ret, nil
}

// this will also update the job and resource offer states
func (controller *SolverController) updateDealState(id string, state uint8) (*data.DealContainer, error) {
	controller.log.Info("update deal", fmt.Sprintf("%s %s", id, data.GetAgreementStateString(state)))

	dealContainer, err := controller.store.UpdateDealState(id, state)
	if err != nil {
		return nil, err
	}

	controller.writeEvent(SolverEvent{
		EventType: DealStateUpdated,
		Deal:      dealContainer,
	})
	_, err = controller.updateJobOfferState(dealContainer.JobOffer, dealContainer.ID, dealContainer.State)
	if err != nil {
		return nil, err
	}
	_, err = controller.updateResourceOfferState(dealContainer.ResourceOffer, dealContainer.ID, dealContainer.State)
	if err != nil {
		return nil, err
	}
	return dealContainer, nil
}

// this will also update the job and resource offer states
func (controller *SolverController) updateDealMediator(id string, mediator string) (*data.DealContainer, error) {
	controller.log.Info("update mediator", fmt.Sprintf("%s %s", id, mediator))
	dealContainer, err := controller.store.UpdateDealMediator(id, mediator)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: DealMediatorUpdated,
		Deal:      dealContainer,
	})
	return dealContainer, nil
}

/*
*
*
*

# Update TX Handlers

*
*
*
*/
func (controller *SolverController) updateDealTransactionsResourceProvider(id string, payload data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	controller.log.Info("update resource provider txs", payload)
	dealContainer, err := controller.store.UpdateDealTransactionsResourceProvider(id, payload)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: ResourceProviderTransactionsUpdated,
		Deal:      dealContainer,
	})
	return dealContainer, nil
}

func (controller *SolverController) updateDealTransactionsJobCreator(id string, payload data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	controller.log.Info("update job creator txs", payload)
	dealContainer, err := controller.store.UpdateDealTransactionsJobCreator(id, payload)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: JobCreatorTransactionsUpdated,
		Deal:      dealContainer,
	})
	return dealContainer, nil
}

func (controller *SolverController) updateDealTransactionsMediator(id string, payload data.DealTransactionsMediator) (*data.DealContainer, error) {
	controller.log.Info("update mediator txs", payload)
	dealContainer, err := controller.store.UpdateDealTransactionsMediator(id, payload)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: MediatorTransactionsUpdated,
		Deal:      dealContainer,
	})
	return dealContainer, nil
}

/*
*
*
*

# Run onchain job

*
*
*
*/

// func (controller *SolverController) runJob(ev jobcreatorweb3.JobcreatorJobAdded) (*data.DealContainer, error) {
// 	options := optionsfactory.NewJobCreatorOptions()
// 	fmt.Printf("options --------------------------------------\n")
// 	spew.Dump(options)
// 	fmt.Printf("ev --------------------------------------\n")
// 	spew.Dump(ev)
// 	return nil, nil
// }
