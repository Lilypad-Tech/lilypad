package directory

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/directory/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

// add an enum for various types of event
type DirectoryEventType int

const (
	DealAdded DirectoryEventType = iota
)

type DirectoryEvent struct {
	EventType DirectoryEventType `json:"event_type"`
	Deal      *data.Deal         `json:"deal"`
}

type DirectoryEventChannel chan DirectoryEvent

type DirectoryController struct {
	web3SDK                *web3.Web3SDK
	web3Events             *web3.EventChannels
	store                  store.DirectoryStore
	directoryEventChannels []DirectoryEventChannel
}

func NewDirectoryController(
	web3SDK *web3.Web3SDK,
	store store.DirectoryStore,
) (*DirectoryController, error) {
	controller := &DirectoryController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
		store:      store,
	}
	return controller, nil
}

func (controller *DirectoryController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
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

	go func() {
		for {
			err := controller.solve()
			if err != nil {
				log.Error().Err(err).Msgf("error solving")
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
	return nil
}

func (controller *DirectoryController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *DirectoryController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *DirectoryController) getEventChannel() DirectoryEventChannel {
	eventChannel := make(DirectoryEventChannel)
	controller.directoryEventChannels = append(controller.directoryEventChannels, eventChannel)
	return eventChannel
}

func (controller *DirectoryController) writeEvent(ev DirectoryEvent) {
	for _, eventChannel := range controller.directoryEventChannels {
		eventChannel <- ev
	}
}

func (controller *DirectoryController) addDeal(deal data.Deal) (*data.Deal, error) {
	deal.ID = ""
	id, err := data.CalculateCID(deal)
	if err != nil {
		return nil, err
	}
	deal.ID = id
	ret, err := controller.store.AddDeal(deal)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(DirectoryEvent{
		EventType: DealAdded,
		Deal:      ret,
	})
	return ret, nil
}
