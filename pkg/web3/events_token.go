package web3

import (
	"context"
	"fmt"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"
)

type TokenEventChannels struct {
	transferChan chan *token.TokenTransfer
	transferSubs []func(token.TokenTransfer)
}

func NewTokenEventChannels() *TokenEventChannels {
	return &TokenEventChannels{
		transferChan: make(chan *token.TokenTransfer),
		transferSubs: []func(token.TokenTransfer){},
	}
}

func (t *TokenEventChannels) Start(
	ctx context.Context,
	cm *system.CleanupManager,
	sdk *Web3SDK,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var transferSub event.Subscription

	connectTransferSub := func() (event.Subscription, error) {
		log.Debug().
			Str("token->connect", "Transfer").
			Msgf("")
		return sdk.Contracts.Token.WatchTransfer(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			t.transferChan,
			[]common.Address{},
			[]common.Address{},
		)
	}

	transferSub, err = connectTransferSub()
	if err != nil {
		return err
	}
	cm.RegisterCallback(unsubscribeSub(transferSub))

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-t.transferChan:
			log.Debug().
				Str("token->event", "Transfer").
				Msgf("%+v", event)
			for _, handler := range t.transferSubs {
				go handler(*event)
			}
		case err := <-transferSub.Err():
			if err != nil {
				return fmt.Errorf("cancel by token Transfer event subscribe error %w", err)
			}
			return nil
		}
	}
}

func (t *TokenEventChannels) SubscribeTransfer(handler func(token.TokenTransfer)) {
	t.transferSubs = append(t.transferSubs, handler)
}
