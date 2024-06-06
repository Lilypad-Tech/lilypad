package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/powtoken"
	"github.com/rs/zerolog/log"
)

type PowEventChannels struct {
	newPowRoundChan chan *powtoken.PowtokenNewPowRound
	newPowRoundSubs []func(powtoken.PowtokenNewPowRound)
}

func NewPowEventChannels() *PowEventChannels {
	return &PowEventChannels{
		newPowRoundChan: make(chan *powtoken.PowtokenNewPowRound),
		newPowRoundSubs: []func(powtoken.PowtokenNewPowRound){},
	}
}

func (s *PowEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var newPowRoundSub event.Subscription

	connectnewPowRoundSub := func() (event.Subscription, error) {
		log.Debug().
			Str("jobcreator->connect", "newPowRound").
			Msgf("")
		return sdk.Contracts.PowToken.WatchNewPowRound(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.newPowRoundChan,
		)
	}

	newPowRoundSub, err = connectnewPowRoundSub()
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		newPowRoundSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-s.newPowRoundChan:
			log.Debug().
				Str("pow->event", "PowtokenNewPowRound").
				Msgf("%+v", event)
			for _, handler := range s.newPowRoundSubs {
				go handler(*event)
			}
		case <-newPowRoundSub.Err():
			newPowRoundSub.Unsubscribe()
			newPowRoundSub, err = connectnewPowRoundSub()
			if err != nil {
				return err
			}
		}
	}
}

func (t *PowEventChannels) SubscribenewPowRound(handler func(powtoken.PowtokenNewPowRound)) {
	t.newPowRoundSubs = append(t.newPowRoundSubs, handler)
}
