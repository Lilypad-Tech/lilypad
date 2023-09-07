package web3

import (
	"context"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type StorageEventChannels struct {
	dealStateChangeChan chan *storage.StorageDealStateChange
	dealStateChangeSubs []func(*storage.StorageDealStateChange)
}

func NewStorageEventChannels() (*StorageEventChannels, error) {
	return &StorageEventChannels{
		dealStateChangeChan: make(chan *storage.StorageDealStateChange),
	}, nil
}

func (s *StorageEventChannels) Start(
	sdk *ContractSDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	dealStateChangeSub, err := sdk.Contracts.Storage.WatchDealStateChange(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		s.dealStateChangeChan,
		[]*big.Int{},
		[]uint8{},
	)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		dealStateChangeSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-s.dealStateChangeChan:
			for _, handler := range s.dealStateChangeSubs {
				handler(event)
			}
		case err := <-dealStateChangeSub.Err():
			return err
		}
	}
}
