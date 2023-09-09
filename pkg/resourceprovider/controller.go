package resourceprovider

import (
	"context"
	"fmt"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver"
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
	controller := &ResourceProviderController{
		options:    options,
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
	}
	return controller, nil
}

func (controller *ResourceProviderController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *ResourceProviderController) subscribeToSolver() error {
	if controller.solverClient == nil {
		return fmt.Errorf("controller.solverClient has not been setup")
	}

	return nil
}

func (controller *ResourceProviderController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *ResourceProviderController) Start(ctx context.Context, cm *system.CleanupManager) error {
	// we know the address of the solver but what is it's url?
	solverUrl, err := controller.web3SDK.GetSolverUrl(controller.options.Web3.SolverAddress)
	if err != nil {
		return err
	}
	// setup the solver client
	controller.solverClient, err = solver.NewSolverClient(http.ClientOptions{
		URL:        solverUrl,
		PrivateKey: controller.options.Web3.PrivateKey,
	})
	if err != nil {
		return err
	}

	err = controller.subscribeToSolver()
	if err != nil {
		return err
	}

	err = controller.subscribeToWeb3()
	if err != nil {
		return err
	}

	// activate the web3 event listeners
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := controller.solve()
				if err != nil {
					log.Error().Err(err).Msgf("error solving")
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
