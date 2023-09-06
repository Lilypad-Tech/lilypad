package web3

import (
	"context"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type StorageEventChannels struct {
	DealStateChange chan *storage.StorageDealStateChange
}

func NewStorageEventChannels() (*StorageEventChannels, error) {
	return &StorageEventChannels{}, nil
}

func (storageEventChannels *StorageEventChannels) Listen(ctx context.Context, sdk *ContractSDK) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}
	sub, err := sdk.Contracts.Storage.WatchDealStateChange(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		storageEventChannels.DealStateChange,
		[]*big.Int{},
		[]uint8{},
	)
	if err != nil {
		return err
	}
	<-ctx.Done()
	sub.Unsubscribe()
	return nil
}
