package solver

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
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
	store store.SolverStore,
) (*SolverController, error) {
	controller := &SolverController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
		store:      store,
	}
	return controller, nil
}

func (controller *SolverController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *SolverController) subscribe() error {
	controller.web3Events.Token.SubscribeTransfer(func(event *token.TokenTransfer) {
		log.Info().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *SolverController) Start(ctx context.Context, cm *system.CleanupManager) error {
	// get the local subscriptions setup
	err := controller.subscribe()
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

func (controller *SolverController) addJobOffer(jobOffer data.JobOffer) (*data.JobOffer, error) {
	jobOffer.ID = ""
	id, err := data.CalculateCID(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOffer.ID = id
	return controller.store.AddJobOffer(jobOffer)
}

func (controller *SolverController) addResourceOffer(resourceOffer data.ResourceOffer) (*data.ResourceOffer, error) {
	resourceOffer.ID = ""
	id, err := data.CalculateCID(resourceOffer)
	if err != nil {
		return nil, err
	}
	resourceOffer.ID = id
	return controller.store.AddResourceOffer(resourceOffer)
}
