package solver

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

type SolverController struct {
	web3SDK    *web3.ContractSDK
	web3Events *web3.EventChannels
	store      store.SolverStore
}

func NewSolverController(
	web3SDK *web3.ContractSDK,
	web3Events *web3.EventChannels,
	store store.SolverStore,
) (*SolverController, error) {
	controller := &SolverController{
		web3SDK:    web3SDK,
		web3Events: web3Events,
		store:      store,
	}
	return controller, nil
}

func (controller *SolverController) Solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *SolverController) Subscribe() error {
	controller.web3Events.Token.SubscribeTransfer(func(event *token.TokenTransfer) {
		log.Debug().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *SolverController) Start(ctx context.Context, cm *system.CleanupManager) error {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := controller.Solve()
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
