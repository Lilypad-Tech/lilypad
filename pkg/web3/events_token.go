package web3

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TokenEventChannels struct {
	Transfer chan *token.TokenTransfer
}

func NewTokenEventChannels() (*TokenEventChannels, error) {
	return &TokenEventChannels{
		Transfer: make(chan *token.TokenTransfer),
	}, nil
}

func (tokenEventChannels *TokenEventChannels) Listen(ctx context.Context, sdk *ContractSDK) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}
	sub, err := sdk.Contracts.Token.WatchTransfer(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		tokenEventChannels.Transfer,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		return err
	}
	<-ctx.Done()
	sub.Unsubscribe()
	return nil
}
