package web3

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/storage"
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
	ctx context.Context,
	cm *system.CleanupManager,
	sdk *Web3SDK,
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
	cm.RegisterCallback(unsubscribeSub(dealStateChangeSub))

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-s.dealStateChangeChan:
			log.Debug().
				Str("storage->event", "DealStateChange").
				Msgf("%+v", event)
			for _, handler := range s.dealStateChangeSubs {
				go handler(*event)
			}
		case err := <-dealStateChangeSub.Err():
			if err != nil {
				return fmt.Errorf("cancel by storage DealStateChange event subscribe error %w", err)
			}
			return nil
		}
	}
}

func (t *StorageEventChannels) SubscribeDealStateChange(handler func(storage.StorageDealStateChange)) {
	t.dealStateChangeSubs = append(t.dealStateChangeSubs, handler)
}
