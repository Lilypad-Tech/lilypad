package mediator

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

type MediatorController struct {
	web3SDK    *web3.Web3SDK
	web3Events *web3.EventChannels
}

func NewMediatorController(
	web3SDK *web3.Web3SDK,
) (*MediatorController, error) {
	controller := &MediatorController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
	}
	return controller, nil
}

func (controller *MediatorController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *MediatorController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *MediatorController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error)
	// get the local subscriptions setup
	err := controller.subscribeToWeb3()
	if err != nil {
		errorChan <- err
		return errorChan
	}
	// activate the web3 event listeners
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := controller.solve()
				if err != nil {
					log.Error().Err(err).Msgf("error solving")
					errorChan <- err
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return errorChan
}
