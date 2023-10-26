package web3

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"
)

type StorageEventChannels struct {
	dealStateChangeChan chan *storage.StorageDealStateChange
	dealStateChangeSubs []func(storage.StorageDealStateChange)
}

func NewStorageEventChannels() *StorageEventChannels {
	return &StorageEventChannels{
		dealStateChangeChan: make(chan *storage.StorageDealStateChange),
		dealStateChangeSubs: []func(storage.StorageDealStateChange){},
	}
}

func (s *StorageEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var dealStateChangeSub event.Subscription

	connectDealStateChangeSub := func() (event.Subscription, error) {
		log.Debug().
			Str("storage->connect", "DealStateChange").
			Msgf("")
		return sdk.Contracts.Storage.WatchDealStateChange(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.dealStateChangeChan,
		)
	}

	dealStateChangeSub, err = connectDealStateChangeSub()
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
			log.Debug().
				Str("storage->event", "DealStateChange").
				Msgf("%+v", event)
			for _, handler := range s.dealStateChangeSubs {
				go handler(*event)
			}
		case err := <-dealStateChangeSub.Err():
			dealStateChangeSub.Unsubscribe()
			dealStateChangeSub, err = connectDealStateChangeSub()
			if err != nil {
				return err
			}
		}
	}
}

func (t *StorageEventChannels) SubscribeDealStateChange(handler func(storage.StorageDealStateChange)) {
	t.dealStateChangeSubs = append(t.dealStateChangeSubs, handler)
}
