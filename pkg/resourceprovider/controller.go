package resourceprovider

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

type ResourceProviderController struct {
	solverClient *solver.SolverClient
	web3SDK      *web3.ContractSDK
	web3Events   *web3.EventChannels
}

func NewResourceProviderController(
	web3SDK *web3.ContractSDK,
) (*ResourceProviderController, error) {
	controller := &ResourceProviderController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
	}
	return controller, nil
}

func (controller *ResourceProviderController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *ResourceProviderController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event *token.TokenTransfer) {
		log.Info().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *ResourceProviderController) Start(ctx context.Context, cm *system.CleanupManager) error {
	// get the local subscriptions setup
	err := controller.subscribeToWeb3()
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
