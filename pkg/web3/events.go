package web3

import (
	"context"

	"github.com/rs/zerolog/log"
)

type EventChannels struct {
	Token       *TokenEventChannels
	Payment     *PaymentEventChannels
	Storage     *StorageEventChannels
	Controller  *ControllerEventChannels
	collections []EventChannelCollection
}

func NewEventChannels() (*EventChannels, error) {
	tokenChannels, err := NewTokenEventChannels()
	if err != nil {
		return nil, err
	}
	paymentChannels, err := NewPaymentEventChannels()
	if err != nil {
		return nil, err
	}
	storageChannels, err := NewStorageEventChannels()
	if err != nil {
		return nil, err
	}
	controllerChannels, err := NewControllerEventChannels()
	if err != nil {
		return nil, err
	}
	collections := []EventChannelCollection{
		tokenChannels,
		// paymentChannels,
		// storageChannels,
		// controllerChannels,
	}
	return &EventChannels{
		Token:       tokenChannels,
		Payment:     paymentChannels,
		Storage:     storageChannels,
		Controller:  controllerChannels,
		collections: collections,
	}, nil
}

func (eventChannels *EventChannels) Start(ctx context.Context, sdk *ContractSDK) error {
	for _, collection := range eventChannels.collections {
		c := collection
		go func() {
			err := c.Start(ctx, sdk)
			if err != nil {
				log.Error().Msgf("error starting listeners: %s", err.Error())
			}
		}()
	}
	return nil
}
