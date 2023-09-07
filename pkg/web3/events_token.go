package web3

import (
	"context"
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
)

type TokenEventChannels struct {
	transferChan chan *token.TokenTransfer
	transferSubs []func(*token.TokenTransfer)
}

func NewTokenEventChannels() (*TokenEventChannels, error) {
	return &TokenEventChannels{
		transferChan: make(chan *token.TokenTransfer),
	}, nil
}

func (t *TokenEventChannels) Start(ctx context.Context, sdk *ContractSDK) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	fmt.Printf("starting token event channels at block %d\n", blockNumber)

	return nil

	// transferSub, err := sdk.Contracts.Token.WatchTransfer(
	// 	&bind.WatchOpts{Start: &blockNumber, Context: ctx},
	// 	t.transferChan,
	// 	[]common.Address{},
	// 	[]common.Address{},
	// )
	// if err != nil {
	// 	return err
	// }

	// go func() {
	// 	<-ctx.Done()
	// 	transferSub.Unsubscribe()
	// }()

	// for {
	// 	select {
	// 	case event := <-t.transferChan:
	// 		for _, handler := range t.transferSubs {
	// 			handler(event)
	// 		}
	// 	case err := <-transferSub.Err():
	// 		return err
	// 	}
	// }
}

func (t *TokenEventChannels) SubscribeTransfer(handler func(*token.TokenTransfer)) {
	t.transferSubs = append(t.transferSubs, handler)
}
