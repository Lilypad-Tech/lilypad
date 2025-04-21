package web3

import (
	"context"
	"fmt"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/pow"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"
)

type PowEventChannels struct {
	newPowRoundChan chan *pow.PowNewPowRound
	newPowRoundSubs []func(pow.PowNewPowRound)
}

func NewPowEventChannels() *PowEventChannels {
	return &PowEventChannels{
		newPowRoundChan: make(chan *pow.PowNewPowRound),
		newPowRoundSubs: []func(pow.PowNewPowRound){},
	}
}

func (s *PowEventChannels) Start(
	ctx context.Context,
	cm *system.CleanupManager,
	sdk *Web3SDK,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var newPowRoundSub event.Subscription

	connectnewPowRoundSub := func() (event.Subscription, error) {
		log.Debug().
			Str("pow->connect", "newPowRound").
			Msgf("start to watch new pow round")

		return sdk.Contracts.Pow.WatchNewPowRound(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.newPowRoundChan,
		)
	}

	newPowRoundSub, err = connectnewPowRoundSub()
	if err != nil {
		return err
	}
	cm.RegisterCallback(unsubscribeSub(newPowRoundSub))

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-s.newPowRoundChan:
			log.Debug().
				Str("pow->event", "PowNewPowRound").
				Msgf("%+v", event)

			for _, handler := range s.newPowRoundSubs {
				go handler(*event)
			}
		case err := <-newPowRoundSub.Err():
			if err != nil {
				return fmt.Errorf("cancel by pow newPowRound event subscribe error %w", err)
			}
			return nil
		}
	}
}

func (t *PowEventChannels) SubscribenewPowRound(handler func(pow.PowNewPowRound)) {
	t.newPowRoundSubs = append(t.newPowRoundSubs, handler)
}
