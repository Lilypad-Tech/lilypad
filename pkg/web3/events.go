package web3

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

type EventChannels struct {
	Token       *TokenEventChannels
	Payment     *PaymentEventChannels
	Storage     *StorageEventChannels
	Controller  *ControllerEventChannels
	collections []EventChannelCollection
}

func NewEventChannels() *EventChannels {
	tokenChannels := NewTokenEventChannels()
	paymentChannels := NewPaymentEventChannels()
	storageChannels := NewStorageEventChannels()
	controllerChannels := NewControllerEventChannels()
	collections := []EventChannelCollection{
		tokenChannels,
		paymentChannels,
		storageChannels,
		controllerChannels,
	}
	return &EventChannels{
		Token:       tokenChannels,
		Payment:     paymentChannels,
		Storage:     storageChannels,
		Controller:  controllerChannels,
		collections: collections,
	}
}

func (eventChannels *EventChannels) Start(
	sdk *ContractSDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	for _, collection := range eventChannels.collections {
		c := collection
		go func() {
			err := c.Start(sdk, ctx, cm)
			if err != nil {
				log.Error().Msgf("error starting listeners: %s", err.Error())
			}
		}()
	}
	return nil
}
