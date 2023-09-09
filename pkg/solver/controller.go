package solver

import (
	"context"
	"math/big"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

// add an enum for various types of event
type SolverEventType int

const (
	JobOfferAdded SolverEventType = iota
	ResourceOfferAdded
	MatchFound
)

type SolverEvent struct {
	EventType     SolverEventType     `json:"event_type"`
	JobOffer      *data.JobOffer      `json:"job_offer"`
	ResourceOffer *data.ResourceOffer `json:"resource_offer"`
}

type SolverEventChannel chan SolverEvent

type SolverController struct {
	web3SDK             *web3.Web3SDK
	web3Events          *web3.EventChannels
	store               store.SolverStore
	solverEventChannels []SolverEventChannel
}

func NewSolverController(
	web3SDK *web3.Web3SDK,
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

	// THIS IS JUST FOR TESTING
	log.Info().Msgf("sending tx")
	tx, err := controller.web3SDK.Contracts.Token.Transfer(
		controller.web3SDK.TransactOpts,
		common.HexToAddress("0x2546BcD3c84621e976D8185a91A922aE77ECEc30"),
		big.NewInt(1),
	)
	if err != nil {
		log.Info().Msgf("error sending tx: %s\n", err.Error())

	} else {
		log.Info().Msgf("tx sent: %s\n", tx.Hash())
	}
	return nil
}

func (controller *SolverController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("solver Transfer. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

// return a new event channel that will hear about events
// coming out of this controller
func (controller *SolverController) getEventChannel() SolverEventChannel {
	eventChannel := make(SolverEventChannel)
	controller.solverEventChannels = append(controller.solverEventChannels, eventChannel)
	return eventChannel
}

// write the given event to all generated event channels
func (controller *SolverController) writeEvent(ev SolverEvent) {
	for _, eventChannel := range controller.solverEventChannels {
		eventChannel <- ev
	}
}

func (controller *SolverController) Start(ctx context.Context, cm *system.CleanupManager) error {
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

func (controller *SolverController) addJobOffer(jobOffer data.JobOffer) (*data.JobOffer, error) {
	jobOffer.ID = ""
	id, err := data.CalculateCID(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOffer.ID = id
	ret, err := controller.store.AddJobOffer(jobOffer)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: JobOfferAdded,
		JobOffer:  ret,
	})
	return ret, nil
}

func (controller *SolverController) addResourceOffer(resourceOffer data.ResourceOffer) (*data.ResourceOffer, error) {
	resourceOffer.ID = ""
	id, err := data.CalculateCID(resourceOffer)
	if err != nil {
		return nil, err
	}
	resourceOffer.ID = id
	ret, err := controller.store.AddResourceOffer(resourceOffer)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType:     ResourceOfferAdded,
		ResourceOffer: ret,
	})
	return ret, nil
}
