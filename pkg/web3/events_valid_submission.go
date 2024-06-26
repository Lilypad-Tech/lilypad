package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/rs/zerolog/log"
)

type ValidPOWSubmittedEventChannels struct {
	powValidPOWSubmittedChan chan *pow.PowValidPOWSubmitted
	powValidPOWSubmittedSubs []func(pow.PowValidPOWSubmitted)
}

func NewValidPOWSubmittedEventChannels() *ValidPOWSubmittedEventChannels {
	return &ValidPOWSubmittedEventChannels{
		powValidPOWSubmittedChan: make(chan *pow.PowValidPOWSubmitted),
		powValidPOWSubmittedSubs: []func(pow.PowValidPOWSubmitted){},
	}
}

func (s *ValidPOWSubmittedEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var powValidPOWSubmittedSub event.Subscription

	connectPowValidPOWSubmittedSub := func() (event.Subscription, error) {
		log.Debug().
			Str("pow->connect", "PowValidPOWSubmitted").
			Msgf("")
		return sdk.Contracts.Pow.WatchValidPOWSubmitted(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.powValidPOWSubmittedChan,
			[]common.Address{},
		)
	} //todo change in future

	powValidPOWSubmittedSub, err = connectPowValidPOWSubmittedSub()
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		powValidPOWSubmittedSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-s.powValidPOWSubmittedChan:
			log.Debug().
				Str("pow->event", "PowValidPOWSubmitted").
				Msgf("%+v", event)
			for _, handler := range s.powValidPOWSubmittedSubs {
				go handler(*event)
			}
		case <-powValidPOWSubmittedSub.Err():
			powValidPOWSubmittedSub.Unsubscribe()
			powValidPOWSubmittedSub, err = connectPowValidPOWSubmittedSub()
			if err != nil {
				return err
			}
		}
	}
}

func (t *ValidPOWSubmittedEventChannels) SubscribePowValidPOWSubmitted(handler func(pow.PowValidPOWSubmitted)) {
	t.powValidPOWSubmittedSubs = append(t.powValidPOWSubmittedSubs, handler)
}
