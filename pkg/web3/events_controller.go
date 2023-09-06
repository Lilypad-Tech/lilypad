package web3

import (
	"context"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/controller"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ControllerEventChannels struct {
	ResourceProviderAgreed chan *controller.ControllerResourceProviderAgreed
	JobCreatorAgreed       chan *controller.ControllerJobCreatorAgreed
	DealAgreed             chan *controller.ControllerDealAgreed
	ResultAdded            chan *controller.ControllerResultAdded
	ResultAccepted         chan *controller.ControllerResultAccepted
	ResultChecked          chan *controller.ControllerResultChecked
	MediationAcceptResult  chan *controller.ControllerMediationAcceptResult
	MediationRejectResult  chan *controller.ControllerMediationRejectResult
	TimeoutSubmitResult    chan *controller.ControllerTimeoutSubmitResult
	TimeoutJudgeResult     chan *controller.ControllerTimeoutJudgeResult
	TimeoutMediateResult   chan *controller.ControllerTimeoutMediateResult
}

func NewControllerEventChannels() (*ControllerEventChannels, error) {
	return &ControllerEventChannels{}, nil
}

func (controllerEventChannels *ControllerEventChannels) Listen(ctx context.Context, sdk *ContractSDK) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}
	resourceProviderAgreedSub, err := sdk.Contracts.Controller.WatchResourceProviderAgreed(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.ResourceProviderAgreed,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	jobCreatorAgreedSub, err := sdk.Contracts.Controller.WatchJobCreatorAgreed(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.JobCreatorAgreed,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	dealAgreedSub, err := sdk.Contracts.Controller.WatchDealAgreed(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.DealAgreed,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	resultAddedSub, err := sdk.Contracts.Controller.WatchResultAdded(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.ResultAdded,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	resultAcceptedSub, err := sdk.Contracts.Controller.WatchResultAccepted(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.ResultAccepted,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	resultCheckedSub, err := sdk.Contracts.Controller.WatchResultChecked(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.ResultChecked,
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		return err
	}
	mediationAcceptResultSub, err := sdk.Contracts.Controller.WatchMediationAcceptResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.MediationAcceptResult,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	mediationRejectResultSub, err := sdk.Contracts.Controller.WatchMediationRejectResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.MediationRejectResult,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	timeoutSubmitResultSub, err := sdk.Contracts.Controller.WatchTimeoutSubmitResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.TimeoutSubmitResult,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	timeoutJudgeResultSub, err := sdk.Contracts.Controller.WatchTimeoutJudgeResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.TimeoutJudgeResult,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	timeoutMediateResultSub, err := sdk.Contracts.Controller.WatchTimeoutMediateResult(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		controllerEventChannels.TimeoutMediateResult,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
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
	return nil
}
