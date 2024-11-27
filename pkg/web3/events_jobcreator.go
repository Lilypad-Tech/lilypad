package web3

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/jobcreator"
	"github.com/rs/zerolog/log"
)

type JobCreatorEventChannels struct {
	jobAddedChan chan *jobcreator.JobcreatorJobAdded
	jobAddedSubs []func(jobcreator.JobcreatorJobAdded)
}

func NewJobCreatorEventChannels() *JobCreatorEventChannels {
	return &JobCreatorEventChannels{
		jobAddedChan: make(chan *jobcreator.JobcreatorJobAdded),
		jobAddedSubs: []func(jobcreator.JobcreatorJobAdded){},
	}
}

func (s *JobCreatorEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var jobAddedSub event.Subscription

	connectJobAddedSub := func() (event.Subscription, error) {
		log.Debug().
			Str("jobcreator->connect", "JobAdded").
			Msgf("")
		return sdk.Contracts.JobCreator.WatchJobAdded(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.jobAddedChan,
		)
	}

	jobAddedSub, err = connectJobAddedSub()
	if err != nil {
		return err
	}
	cm.RegisterCallback(unsubscribeSub(jobAddedSub))

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-s.jobAddedChan:
			log.Debug().
				Str("storage->event", "JobAdded").
				Msgf("%+v", event)
			for _, handler := range s.jobAddedSubs {
				go handler(*event)
			}
		case err := <-jobAddedSub.Err():
			if err != nil {
				return fmt.Errorf("cancel by job JobAdded event subscribe error %w", err)
			}
			return nil
		}
	}
}

func (t *JobCreatorEventChannels) SubscribeJobAdded(handler func(jobcreator.JobcreatorJobAdded)) {
	t.jobAddedSubs = append(t.jobAddedSubs, handler)
}
