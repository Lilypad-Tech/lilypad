package web3

import (
	"context"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

type EventChannels struct {
	Token       *TokenEventChannels
	Payment     *PaymentEventChannels
	Storage     *StorageEventChannels
	JobCreator  *JobCreatorEventChannels
	Mediation   *MediationEventChannels
	Pow         *PowEventChannels
	collections []EventChannelCollection
}

func NewEventChannels() *EventChannels {
	tokenChannels := NewTokenEventChannels()
	paymentChannels := NewPaymentEventChannels()
	storageChannels := NewStorageEventChannels()
	jobCreatorChannels := NewJobCreatorEventChannels()
	mediationChannels := NewMediationEventChannels()
	powChannels := NewPowEventChannels()
	collections := []EventChannelCollection{
		tokenChannels,
		paymentChannels,
		storageChannels,
		jobCreatorChannels,
		mediationChannels,
		powChannels,
	}
	return &EventChannels{
		Token:       tokenChannels,
		Payment:     paymentChannels,
		Storage:     storageChannels,
		JobCreator:  jobCreatorChannels,
		Mediation:   mediationChannels,
		Pow:         powChannels,
		collections: collections,
	}
}

func (eventChannels *EventChannels) Start(
	ctx context.Context,
	cm *system.CleanupManager,
	sdk *Web3SDK,
) error {
	for _, collection := range eventChannels.collections {
		c := collection
		go func() {
			for {
				err := c.Start(ctx, cm, sdk)
				if err != nil {
					log.Error().Msgf("error starting listeners: %s reconnect in 2 seconds", err.Error())
				}

				time.Sleep(time.Second * 2)
			}
		}()
	}
	return nil
}
