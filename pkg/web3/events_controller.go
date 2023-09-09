package web3

import (
	"context"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/controller"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ControllerEventChannels struct {
	resourceProviderAgreedChan chan *controller.ControllerResourceProviderAgreed
	resourceProviderAgreedSubs []func(*controller.ControllerResourceProviderAgreed)

	jobCreatorAgreedChan chan *controller.ControllerJobCreatorAgreed
	jobCreatorAgreedSubs []func(*controller.ControllerJobCreatorAgreed)

	dealAgreedChan chan *controller.ControllerDealAgreed
	dealAgreedSubs []func(*controller.ControllerDealAgreed)

	resultAddedChan chan *controller.ControllerResultAdded
	resultAddedSubs []func(*controller.ControllerResultAdded)

	resultAcceptedChan chan *controller.ControllerResultAccepted
	resultAcceptedSubs []func(*controller.ControllerResultAccepted)

	resultCheckedChan chan *controller.ControllerResultChecked
	resultCheckedSubs []func(*controller.ControllerResultChecked)

	mediationAcceptResultChan chan *controller.ControllerMediationAcceptResult
	mediationAcceptResultSubs []func(*controller.ControllerMediationAcceptResult)

	mediationRejectResultChan chan *controller.ControllerMediationRejectResult
	mediationRejectResultSubs []func(*controller.ControllerMediationRejectResult)

	timeoutSubmitResultChan chan *controller.ControllerTimeoutSubmitResult
	timeoutSubmitResultSubs []func(*controller.ControllerTimeoutSubmitResult)

	timeoutJudgeResultChan chan *controller.ControllerTimeoutJudgeResult
	timeoutJudgeResultSubs []func(*controller.ControllerTimeoutJudgeResult)

	timeoutMediateResultChan chan *controller.ControllerTimeoutMediateResult
	timeoutMediateResultSubs []func(*controller.ControllerTimeoutMediateResult)
}

func NewControllerEventChannels() *ControllerEventChannels {
	return &ControllerEventChannels{
		resourceProviderAgreedChan: make(chan *controller.ControllerResourceProviderAgreed),
		jobCreatorAgreedChan:       make(chan *controller.ControllerJobCreatorAgreed),
		dealAgreedChan:             make(chan *controller.ControllerDealAgreed),
		resultAddedChan:            make(chan *controller.ControllerResultAdded),
		resultAcceptedChan:         make(chan *controller.ControllerResultAccepted),
		resultCheckedChan:          make(chan *controller.ControllerResultChecked),
		mediationAcceptResultChan:  make(chan *controller.ControllerMediationAcceptResult),
		mediationRejectResultChan:  make(chan *controller.ControllerMediationRejectResult),
		timeoutSubmitResultChan:    make(chan *controller.ControllerTimeoutSubmitResult),
		timeoutJudgeResultChan:     make(chan *controller.ControllerTimeoutJudgeResult),
		timeoutMediateResultChan:   make(chan *controller.ControllerTimeoutMediateResult),
	}
}

func (c *ControllerEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}
	resourceProviderAgreedSub, err := sdk.Contracts.Controller.WatchResourceProviderAgreed(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.resourceProviderAgreedChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	jobCreatorAgreedSub, err := sdk.Contracts.Controller.WatchJobCreatorAgreed(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.jobCreatorAgreedChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	dealAgreedSub, err := sdk.Contracts.Controller.WatchDealAgreed(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.dealAgreedChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	resultAddedSub, err := sdk.Contracts.Controller.WatchResultAdded(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.resultAddedChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	resultAcceptedSub, err := sdk.Contracts.Controller.WatchResultAccepted(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.resultAcceptedChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	resultCheckedSub, err := sdk.Contracts.Controller.WatchResultChecked(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.resultCheckedChan,
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		return err
	}
	mediationAcceptResultSub, err := sdk.Contracts.Controller.WatchMediationAcceptResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.mediationAcceptResultChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	mediationRejectResultSub, err := sdk.Contracts.Controller.WatchMediationRejectResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.mediationRejectResultChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	timeoutSubmitResultSub, err := sdk.Contracts.Controller.WatchTimeoutSubmitResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.timeoutSubmitResultChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	timeoutJudgeResultSub, err := sdk.Contracts.Controller.WatchTimeoutJudgeResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.timeoutJudgeResultChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	timeoutMediateResultSub, err := sdk.Contracts.Controller.WatchTimeoutMediateResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		c.timeoutMediateResultChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		resourceProviderAgreedSub.Unsubscribe()
		jobCreatorAgreedSub.Unsubscribe()
		dealAgreedSub.Unsubscribe()
		resultAddedSub.Unsubscribe()
		resultAcceptedSub.Unsubscribe()
		resultCheckedSub.Unsubscribe()
		mediationAcceptResultSub.Unsubscribe()
		mediationRejectResultSub.Unsubscribe()
		timeoutSubmitResultSub.Unsubscribe()
		timeoutJudgeResultSub.Unsubscribe()
		timeoutMediateResultSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-c.resourceProviderAgreedChan:
			for _, handler := range c.resourceProviderAgreedSubs {
				handler(event)
			}
		case err := <-resourceProviderAgreedSub.Err():
			return err

		case event := <-c.jobCreatorAgreedChan:
			for _, handler := range c.jobCreatorAgreedSubs {
				handler(event)
			}
		case err := <-jobCreatorAgreedSub.Err():
			return err

		case event := <-c.dealAgreedChan:
			for _, handler := range c.dealAgreedSubs {
				handler(event)
			}
		case err := <-dealAgreedSub.Err():
			return err

		case event := <-c.resultAddedChan:
			for _, handler := range c.resultAddedSubs {
				handler(event)
			}
		case err := <-resultAddedSub.Err():
			return err

		case event := <-c.resultAcceptedChan:
			for _, handler := range c.resultAcceptedSubs {
				handler(event)
			}
		case err := <-resultAcceptedSub.Err():
			return err

		case event := <-c.resultCheckedChan:
			for _, handler := range c.resultCheckedSubs {
				handler(event)
			}
		case err := <-resultCheckedSub.Err():
			return err

		case event := <-c.mediationAcceptResultChan:
			for _, handler := range c.mediationAcceptResultSubs {
				handler(event)
			}
		case err := <-mediationAcceptResultSub.Err():
			return err

		case event := <-c.mediationRejectResultChan:
			for _, handler := range c.mediationRejectResultSubs {
				handler(event)
			}
		case err := <-mediationRejectResultSub.Err():
			return err

		case event := <-c.timeoutSubmitResultChan:
			for _, handler := range c.timeoutSubmitResultSubs {
				handler(event)
			}
		case err := <-timeoutSubmitResultSub.Err():
			return err

		case event := <-c.timeoutJudgeResultChan:
			for _, handler := range c.timeoutJudgeResultSubs {
				handler(event)
			}
		case err := <-timeoutJudgeResultSub.Err():
			return err

		case event := <-c.timeoutMediateResultChan:
			for _, handler := range c.timeoutMediateResultSubs {
				handler(event)
			}
		case err := <-timeoutMediateResultSub.Err():
			return err
		}
	}
}
