package jobcreator

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

type JobCreatorController struct {
	web3SDK    *web3.Web3SDK
	web3Events *web3.EventChannels
}

func NewJobCreatorController(
	web3SDK *web3.Web3SDK,
) (*JobCreatorController, error) {
	controller := &JobCreatorController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
	}
	return controller, nil
}

func (controller *JobCreatorController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *JobCreatorController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event *token.TokenTransfer) {
		log.Info().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *JobCreatorController) Start(ctx context.Context, cm *system.CleanupManager) error {
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
