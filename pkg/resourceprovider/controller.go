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
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

type ResourceProviderController struct {
	solverClient *solver.SolverClient
	options      ResourceProviderOptions
	web3SDK      *web3.Web3SDK
	web3Events   *web3.EventChannels
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

	go func() {
		for {
			err := controller.solve()
			if err != nil {
				log.Error().Msgf("error solving: %s", err.Error())
				errorChan <- err
				return
			}
			select {
			case <-time.After(1 * time.Second):
			case <-ctx.Done():
				return
			}
		}
	}()
	return errorChan
}

/*
Solve
*/
func (controller *ResourceProviderController) solve() error {
	err := controller.ensureResourceOffers()
	if err != nil {
		return err
	}
	return nil
}

/*
Subscribe
*/
func (controller *ResourceProviderController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(event solver.SolverEvent) {
		log.Info().Msgf("New solver event %+v", event)
	})
	return nil
}

func (controller *ResourceProviderController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("New SubscribeTransfer. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

/*
Ensure resource offers are posted to the solve
*/

// convert the config string values into big ints
func (controller *ResourceProviderController) getPricing() data.Pricing {
	config := controller.options.Offers.DefaultPricing
	return data.Pricing{
		InstructionPrice:          web3.ConvertStringToInt64(config.InstructionPrice),
		Timeout:                   web3.ConvertStringToInt64(config.Timeout),
		TimeoutCollateral:         web3.ConvertStringToInt64(config.TimeoutCollateral),
		PaymentCollateral:         web3.ConvertStringToInt64(config.PaymentCollateral),
		ResultsCollateralMultiple: web3.ConvertStringToInt64(config.ResultsCollateralMultiple),
		MediationFee:              web3.ConvertStringToInt64(config.MediationFee),
	}
}

func (controller *ResourceProviderController) ensureResourceOffers() error {
	existingResourceOffers, err := controller.solverClient.GetResourceOffers(store.GetResourceOffersQuery{
		ResourceProvider: controller.web3SDK.GetAddress().String(),
	})
	if err != nil {
		return err
	}

	// create a map of the ids of resource offers we have
	// this will allow us to check if we need to create a new one
	// or update an existing one
	existingResourceOffersMap := map[string]data.ResourceOffer{}
	for _, existingResourceOffer := range existingResourceOffers {
		existingResourceOffersMap[existingResourceOffer.ID] = existingResourceOffer
	}

	addResourceOffers := []data.ResourceOffer{}

	// map over the specs we have in the config
	for index, spec := range controller.options.Offers.Specs {

		resourceOffer := data.ResourceOffer{
			ResourceProvider: controller.web3SDK.GetAddress().String(),
			Index:            index,
			Spec:             spec,
			Modules:          controller.options.Offers.Modules,
			DefaultPricing:   controller.getPricing(),
			ModulePricing:    map[string]data.Pricing{},
		}

		resourceOfferID, err := data.GetResourceOfferID(resourceOffer)
		if err != nil {
			return err
		}

		// check if the resource offer already exists
		// if it does then we need to update it
		// if it doesn't then we need to add it
		_, ok := existingResourceOffersMap[resourceOfferID]
		if !ok {
			addResourceOffers = append(addResourceOffers, resourceOffer)
		}
	}

	// add the resource offers we need to add
	for _, resourceOffer := range addResourceOffers {
		log.Info().
			Str("add resource offer", fmt.Sprintf("%+v", resourceOffer)).
			Msgf("")
		_, err := controller.solverClient.AddResourceOffer(resourceOffer)
		if err != nil {
			return err
		}
	}

	return err
}
