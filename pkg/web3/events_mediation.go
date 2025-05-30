package web3

import (
	"context"
	"fmt"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/mediation"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"
)

type MediationEventChannels struct {
	mediationRequestedChan chan *mediation.MediationMediationRequested
	mediationRequestedSubs []func(mediation.MediationMediationRequested)
}

func NewMediationEventChannels() *MediationEventChannels {
	return &MediationEventChannels{
		mediationRequestedChan: make(chan *mediation.MediationMediationRequested),
		mediationRequestedSubs: []func(mediation.MediationMediationRequested){},
	}
}

func (m *MediationEventChannels) Start(
	ctx context.Context,
	cm *system.CleanupManager,
	sdk *Web3SDK,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var mediationRequestedSub event.Subscription

	connectMediationRequestedSub := func() (event.Subscription, error) {
		log.Debug().
			Str("mediation->connect", "MediationRequested").
			Msgf("")
		return sdk.Contracts.Mediation.WatchMediationRequested(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			m.mediationRequestedChan,
		)
	}

	mediationRequestedSub, err = connectMediationRequestedSub()
	if err != nil {
		return err
	}
	cm.RegisterCallback(unsubscribeSub(mediationRequestedSub))

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-m.mediationRequestedChan:
			log.Debug().
				Str("mediation->event", "MediationRequested").
				Msgf("%+v", event)
			for _, handler := range m.mediationRequestedSubs {
				go handler(*event)
			}
		case err := <-mediationRequestedSub.Err():
			if err != nil {
				return fmt.Errorf("cancel by mediation MediationRequested event subscribe error %w", err)
			}
			return nil
		}
	}
}

func (m *MediationEventChannels) SubscribeMediationRequested(handler func(mediation.MediationMediationRequested)) {
	m.mediationRequestedSubs = append(m.mediationRequestedSubs, handler)
}
