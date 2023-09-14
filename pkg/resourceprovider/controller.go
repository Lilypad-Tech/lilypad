package resourceprovider

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

type ResourceProviderController struct {
	solverClient *solver.SolverClient
	options      ResourceProviderOptions
	web3SDK      *web3.Web3SDK
	web3Events   *web3.EventChannels
	loop         *system.ControlLoop
}

func NewResourceProviderController(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
) (*ResourceProviderController, error) {
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Web3.SolverAddress)
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

	controller := &ResourceProviderController{
		solverClient: solverClient,
		options:      options,
		web3SDK:      web3SDK,
		web3Events:   web3.NewEventChannels(),
	}
	return controller, nil
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
func (controller *ResourceProviderController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(ev solver.SolverEvent) {
		// we need to agree to the deal now we've heard about it
		if ev.EventType == solver.DealAdded {
			if ev.Deal == nil {
				system.Error(system.ResourceProviderService, "solver event", fmt.Errorf("RP received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.ResourceProvider != controller.web3SDK.GetAddress().String() {
				return
			}

			solver.ServiceLogSolverEvent(system.ResourceProviderService, ev)

			// trigger the solver
			controller.loop.Trigger()
		}
	})
	return nil
}

func (controller *ResourceProviderController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		deal, err := controller.solverClient.GetDeal(ev.DealId.String())
		if err != nil {
			system.Error(system.ResourceProviderService, "error getting deal", err)
			return
		}
		if deal.ResourceProvider != controller.web3SDK.GetAddress().String() {
			return
		}
		system.Info(system.ResourceProviderService, "StorageDealStateChange", ev)
		controller.loop.Trigger()
	})
	return nil
}

func (controller *ResourceProviderController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
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
		system.ResourceProviderService,
		ctx,
		time.Second*10,
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

func (controller *ResourceProviderController) solve() error {
	system.Debug(system.ResourceProviderService, "solving", "")
	err := controller.agreeToDeals()
	if err != nil {
		return err
	}
	err = controller.ensureResourceOffers()
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
func (controller *ResourceProviderController) agreeToDeals() error {
	// load the deals that are in DealNegotiating
	// and do not have a TransactionsResourceProvider.Agree tx
	matchedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			ResourceProvider: controller.web3SDK.GetAddress().String(),
			State:            "DealNegotiating",
		},
		// this is where the solver has found us a match and we need to agree to it
		func(dealContainer data.DealContainer) bool {
			return dealContainer.TransactionsResourceProvider.Agree == ""
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
		system.Info(system.ResourceProviderService, "agree", dealContainer)
		tx, err := controller.web3SDK.Agree(dealContainer.Deal)
		if err != nil {
			// TODO: we need a way of deciding based on certain classes of error what happens
			// some will be retryable - otherwise will be fatal
			// we need a way to exit a job loop as a baseline
			system.Error(system.ResourceProviderService, "error calling agree tx for deal", err)
			continue
		}
		system.Info(system.ResourceProviderService, "agree tx", tx)

		// we have agreed to the deal so we need to update the tx in the solver
		_, err = controller.solverClient.UpdateTransactionsResourceProvider(dealContainer.ID, data.DealTransactionsResourceProvider{
			Agree: tx,
		})
		if err != nil {
			// TODO: we need a way of deciding based on certain classes of error what happens
			// some will be retryable - otherwise will be fatal
			// we need a way to exit a job loop as a baseline
			system.Error(system.ResourceProviderService, "error calling agree tx for deal", err)
			continue
		}
		system.Info(system.ResourceProviderService, "updated deal with agree tx", tx)
	}

	return err
}

/*
 *
 *
 *

 Ensure resource offers

 *
 *
 *
*/

/*
Ensure resource offers are posted to the solve
*/

func (controller *ResourceProviderController) getResourceOffer(index int, spec data.MachineSpec) data.ResourceOffer {
	return data.ResourceOffer{
		// assign CreatedAt to the current millisecond timestamp
		CreatedAt:        int(time.Now().UnixNano() / int64(time.Millisecond)),
		ResourceProvider: controller.web3SDK.GetAddress().String(),
		Index:            index,
		Spec:             spec,
		Modules:          controller.options.Offers.Modules,
		Mode:             controller.options.Offers.Mode,
		DefaultPricing:   controller.options.Offers.DefaultPricing,
		DefaultTimeouts:  controller.options.Offers.DefaultTimeouts,
		ModulePricing:    map[string]data.DealPricing{},
		ModuleTimeouts:   map[string]data.DealTimeouts{},
		TrustedParties:   controller.options.Offers.TrustedParties,
	}
}

func (controller *ResourceProviderController) ensureResourceOffers() error {
	// load the resource offers that are currently active and so should not be replaced
	activeResourceOffers, err := controller.solverClient.GetResourceOffers(store.GetResourceOffersQuery{
		ResourceProvider: controller.web3SDK.GetAddress().String(),
		Active:           true,
	})
	if err != nil {
		return err
	}

	// create a map of the ids of resource offers we have
	// this will allow us to check if we need to create a new one
	// or update an existing one - we use the "index" because
	// the id's are changing because of the timestamps
	existingResourceOffersMap := map[int]data.ResourceOfferContainer{}
	for _, existingResourceOffer := range activeResourceOffers {
		existingResourceOffersMap[existingResourceOffer.ResourceOffer.Index] = existingResourceOffer
	}

	addResourceOffers := []data.ResourceOffer{}

	// map over the specs we have in the config
	for index, spec := range controller.options.Offers.Specs {

		// check if the resource offer already exists
		// if it does then we need to update it
		// if it doesn't then we need to add it
		_, ok := existingResourceOffersMap[index]
		if !ok {
			addResourceOffers = append(addResourceOffers, controller.getResourceOffer(index, spec))
		}
	}

	// add the resource offers we need to add
	for _, resourceOffer := range addResourceOffers {
		system.Info(system.ResourceProviderService, "add resource offer", resourceOffer)
		_, err := controller.solverClient.AddResourceOffer(resourceOffer)
		if err != nil {
			return err
		}
	}

	return err
}
