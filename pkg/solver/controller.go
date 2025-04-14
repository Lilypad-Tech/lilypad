package solver

import (
	"context"
	"fmt"
	"time"

	"github.com/lilypad-tech/lilypad/v2/pkg/data"
	"github.com/lilypad-tech/lilypad/v2/pkg/metricsDashboard"
	"github.com/lilypad-tech/lilypad/v2/pkg/solver/matcher"
	"github.com/lilypad-tech/lilypad/v2/pkg/solver/stats"
	"github.com/lilypad-tech/lilypad/v2/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/v2/pkg/system"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/mediation"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/storage"
	"github.com/rs/zerolog"
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
	stats           stats.Stats
	loop            *system.ControlLoop
	solverEventSubs []func(SolverEvent)
	options         SolverOptions
	log             *zerolog.Logger
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
	stats stats.Stats,
	options SolverOptions,
	tracer trace.Tracer,
	meter metric.Meter,
) (*SolverController, error) {
	controller := &SolverController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
		store:      store,
		stats:      stats,
		options:    options,
		log:        system.GetLogger(system.SolverService),
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
	controller.log.Debug().Msgf("controller.web3Events.Start")
	err = controller.web3Events.Start(ctx, cm, controller.web3SDK)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	// make sure we are registered as a solver
	// so that users can lookup our URL
	controller.log.Debug().Msgf("controller.registerAsSolver")
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
	controller.log.Debug().Msgf("controller.loop.Start")
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
			controller.log.Error().Err(err).Msg("error updating deal state")
			return
		}
		controller.log.Info().Str("StorageDealStateChange", data.GetAgreementStateString(ev.State)).Msg("")
		system.DumpObjectDebug(ev)
		// update the store with the state change
		controller.loop.Trigger()
	})

	// update the mediator
	controller.web3Events.Mediation.SubscribeMediationRequested(func(ev mediation.MediationMediationRequested) {
		controller.log.Info().Msg("MediationMediationRequested")
		system.DumpObjectDebug(ev)
		_, err := controller.updateDealMediator(ev.DealId, ev.Mediator.String())
		if err != nil {
			controller.log.Error().Err(err).Msg("error updating deal state")
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

	controller.log.Debug().Msgf("GetUser with selfAddress: %s", selfAddress.String())
	selfUser, err := controller.web3SDK.GetUser(selfAddress)
	if err != nil {
		return err
	}

	// TODO: check the other props and call update if they have changed
	controller.log.Debug().Msgf("selfUser.Url: %s", selfUser.Url)
	controller.log.Debug().Msgf("controller.options.Server.URL: %s", controller.options.Server.URL)
	if selfUser.Url != controller.options.Server.URL {
		controller.log.Info().
			Str("address", selfAddress.String()).
			Str("oldURL", selfUser.Url).
			Str("newURL", controller.options.Server.URL).
			Msg("solver will be updated because URL has changed")

		err = controller.web3SDK.UpdateUser(
			"",
			controller.options.Server.URL,
			[]uint8{solverType},
		)
		if err != nil {
			return err
		}
	} else {
		controller.log.Info().
			Str("address", selfAddress.String()).
			Str("url", controller.options.Server.URL).
			Msg("solver url already correct")
	}

	existingSolvers, err := controller.web3SDK.GetSolverAddresses()
	if err != nil {
		return err
	}
	foundSolver := false
	for _, existingSolver := range existingSolvers {
		if existingSolver.String() == selfAddress.String() {
			controller.log.Info().Str("address", selfAddress.String()).Msg("solver exists")
			foundSolver = true
			break
		}
	}
	if !foundSolver {
		controller.log.Info().Msg("solver registering")
		// add the solver to the storage contract
		err = controller.web3SDK.AddUserToList(
			solverType,
		)
		if err != nil {
			return err
		}
		controller.log.Info().Str("address", selfAddress.String()).Msg("solver registered")
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

	// Remove expired job offers
	err := controller.cancelExpiredJobs(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "remove expired offers failed")
		span.RecordError(err)
		return err
	}

	// Match job offers with resource offers to make deals
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

	// Add deals to the store, update offer states, and notify network
	span.AddEvent("add_deals.start")
	for _, deal := range deals {
		_, err := controller.addDeal(ctx, deal)
		if err != nil {
			return err
		}
	}
	span.AddEvent("add_deals.done")

	// Report deal state metrics
	span.AddEvent("report_deal_metrics.start")
	storedDeals, err := controller.store.GetDealsAll()
	if err != nil {
		span.SetStatus(codes.Error, "get all deals failed")
		span.RecordError(err)
		return err
	}
	jobOffers, err := controller.store.GetJobOffers(store.GetJobOffersQuery{Cancelled: system.BoolPointer(true)})
	if err != nil {
		span.SetStatus(codes.Error, "get cancelled job offers failed")
		span.RecordError(err)
		return err
	}
	err = reportJobMetrics(ctx, controller.meter, storedDeals, jobOffers)
	if err != nil {
		span.SetStatus(codes.Error, "report deal metrics failed")
		span.RecordError(err)
	}
	span.AddEvent("report_deal_metrics.done")

	return nil
}

func (controller *SolverController) cancelExpiredJobs(ctx context.Context) error {
	ctx, span := controller.tracer.Start(ctx, "cancel_expired_jobs")
	defer span.End()

	// Get active job offers
	span.AddEvent("db.get_job_offers.start")
	jobOffers, err := controller.store.GetJobOffers(store.GetJobOffersQuery{
		Active: true,
	})
	if err != nil {
		controller.log.Error().Func(system.AddTraceContext(span)).Err(err).Msg("get job offers failed")
		span.SetStatus(codes.Error, "get job offers failed")
		span.RecordError(err)
		return err
	}
	span.AddEvent("db.get_job_offers.done")

	// Check active job offers, and cancel expired offers
	// and associated resource offers and deals
	span.AddEvent("expire_jobs.start")
	expiredOffers := []string{}
	expiredDeals := []string{}
	for _, jobOffer := range jobOffers {
		now := time.Now().UnixMilli()
		if now-int64(jobOffer.JobOffer.CreatedAt) > int64(controller.options.JobTimeoutSeconds*1000) {
			if jobOffer.DealID == "" {
				// Cancel expired job offers
				_, err := controller.updateJobOfferState(jobOffer.ID, jobOffer.DealID, data.GetAgreementStateIndex("JobTimedOut"))
				if err != nil {
					controller.log.Error().Func(system.AddTraceContext(span)).Err(err).Msg("update expired job offer state failed")
					span.SetStatus(codes.Error, "update expired job offer state failed")
					span.RecordError(err)
				}
			} else {
				// Cancel expired job offers, resource offers, and deals
				deal, err := controller.updateDealState(jobOffer.DealID, data.GetAgreementStateIndex("JobTimedOut"))
				if err != nil {
					controller.log.Error().Func(system.AddTraceContext(span)).Err(err).Msg("update expired deal state failed")
					span.SetStatus(codes.Error, "update expired deal state failed")
					span.RecordError(err)
				}
				expiredDeals = append(expiredDeals, jobOffer.DealID)
				controller.stats.PostJobRun(deal)
			}
			expiredOffers = append(expiredOffers, jobOffer.ID)
		}
	}
	span.SetAttributes(attribute.KeyValue{
		Key:   "expired_job_offers",
		Value: attribute.StringSliceValue(expiredOffers),
	})
	span.SetAttributes(attribute.KeyValue{
		Key:   "expired_deals",
		Value: attribute.StringSliceValue(expiredDeals),
	})
	span.AddEvent("expire_jobs.end")

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

	controller.log.Info().Str("cid", jobOffer.ID).
		Str("address", jobOffer.JobCreator).
		Any("jobOffer", jobOffer).
		Msg("add job offer")

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

	// Check if we have an unmatched resource offer for the resource provider
	existingOffers, err := controller.store.GetResourceOffers(store.GetResourceOffersQuery{
		ResourceProvider: resourceOffer.ResourceProvider,
		NotMatched:       true,
	})
	if len(existingOffers) > 0 {
		controller.log.Warn().
			Str("address", resourceOffer.ResourceProvider).
			Msg("resource provider posted a resource offer when an unmatched offer already exists")
		// TODO(bgins) Return error to resource provider
		// The resource provider currently crashes when an error is returned. Once we update the
		// resource provider to selectively handle errors, we should return one here.
		return nil, nil
	}

	// Check the resource provider's ETH balance
	balance, err := controller.web3SDK.GetBalance(resourceOffer.ResourceProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ETH balance for resource provider: %v", err)
	}
	// Convert InstructionPrice from ETH to Wei
	requiredBalanceWei := web3.EtherToWei(REQUIRED_BALANCE_IN_WEI) // 0.0006 based on the required balance for a job

	// If the balance is less than the required balance, don't add the resource offer
	if balance.Cmp(requiredBalanceWei) < 0 {
		controller.log.Error().Err(err).
			Str("addresss", resourceOffer.ResourceProvider).
			Str("balance", balance.String()).
			Str("requiredBalance", requiredBalanceWei.String()).
			Msg("resource provider does not have enough ETH to post offer")
		return nil, nil
	}

	// required LP balance
	requiredBalanceLp := web3.EtherToWei(float64(resourceOffer.DefaultPricing.InstructionPrice)) // based on the required LP balance for a job
	balanceLp, err := controller.web3SDK.GetLPBalance(resourceOffer.ResourceProvider)
	if err != nil {
		controller.log.Error().Err(err).Msg("failed to retrieve LP balance for resource provider")
		return nil, nil
	}
	if balanceLp.Cmp(requiredBalanceLp) < 0 {
		controller.log.Error().Err(err).
			Str("addresss", resourceOffer.ResourceProvider).
			Str("balance", balanceLp.String()).
			Str("requiredBalance", requiredBalanceLp.String()).
			Msg("resource provider does not have enough LP to post offer")
		return nil, nil
	}

	controller.log.Info().Str("cid", resourceOffer.ID).
		Str("address", resourceOffer.ResourceProvider).
		Any("resourceOffer", resourceOffer).
		Msg("add resource offer")

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

// Remove resource offers in an unmatched DealNegotiating[0] state
func (controller *SolverController) removeUnmatchedResourceOffers(resourceProviderID string) error {
	controller.log.Info().Str("address", resourceProviderID).Msg("remove unmatched resource offers")
	resourceOffers, err := controller.store.GetResourceOffers(store.GetResourceOffersQuery{
		ResourceProvider: resourceProviderID,
	})
	if err != nil {
		return err
	}

	for _, offer := range resourceOffers {
		if offer.State == 0 {
			err = controller.store.RemoveResourceOffer(offer.ID)
			if err != nil {
				controller.log.Error().Err(err).
					Str("address", resourceProviderID).
					Str("cid", offer.ID).
					Msg("remove resource offer failed")
			}
		}
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

	controller.log.Info().Func(system.AddTraceContext(span)).
		Str("cid", deal.ID).
		Str("resourceProvider", deal.Members.ResourceProvider).
		Str("jobCreator", deal.Members.JobCreator).
		Any("deal", deal).
		Msg("add deal")

	//creates deal container and sets state to agreed
	dealContainer := data.GetDealContainer(deal)
	dealContainer.State = data.GetAgreementStateIndex("DealAgreed")

	span.AddEvent("store.add_deal.start")
	ret, err := controller.store.AddDeal(dealContainer)
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
	controller.log.Info().Str("cid", id).Str("state", data.GetAgreementStateString(state)).Msg("update job offer")

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
	controller.log.Info().Str("cid", id).Str("state", data.GetAgreementStateString(state)).Msg("update resource offer")

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
	controller.log.Info().Str("cid", id).Str("state", data.GetAgreementStateString(state)).Msg("update deal")

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
	controller.log.Info().Str("cid", id).Str("address", mediator).Msg("update deal mediator")
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
	controller.log.Info().Any("payload", payload).Msg("update resource provider txs")
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
	controller.log.Info().Any("payload", payload).Msg("update job creator txs")
	dealContainer, err := controller.store.UpdateDealTransactionsJobCreator(id, payload)
	if err != nil {
		return nil, err
	}
	if payload.AcceptResult != "" {
		return controller.updateDealState(id, data.GetAgreementStateIndex("ResultsAccepted"))
	}
	controller.writeEvent(SolverEvent{
		EventType: JobCreatorTransactionsUpdated,
		Deal:      dealContainer,
	})
	return dealContainer, nil
}

func (controller *SolverController) updateDealTransactionsMediator(id string, payload data.DealTransactionsMediator) (*data.DealContainer, error) {
	controller.log.Info().Any("payload", payload).Msg("update mediator txs")
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
