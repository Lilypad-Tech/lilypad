package web3

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

type TokenEventChannels struct {
	transferChan chan *token.TokenTransfer
	transferSubs []func(token.TokenTransfer)
}

func NewTokenEventChannels() *TokenEventChannels {
	return &TokenEventChannels{
		transferChan: make(chan *token.TokenTransfer),
	}
}

func (t *TokenEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	transferSub, err := sdk.Contracts.Token.WatchTransfer(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		t.transferChan,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		transferSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-t.transferChan:
			log.Info().Msgf("token event -> transfer: %+v", event)
			for _, handler := range t.transferSubs {
				go handler(*event)
			}
		case err := <-transferSub.Err():
			return err
		}
	}
}

func (t *TokenEventChannels) SubscribeTransfer(handler func(token.TokenTransfer)) {
	t.transferSubs = append(t.transferSubs, handler)
}