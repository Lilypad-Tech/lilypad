package mediator

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type MediatorOptions struct {
	Web3 web3.Web3Options
}

type Mediator struct {
	web3SDK    *web3.ContractSDK
	controller *MediatorController
}

func NewMediator(
	options MediatorOptions,
	web3SDK *web3.ContractSDK,
) (*Mediator, error) {
	controller, err := NewMediatorController(web3SDK)
	if err != nil {
		return nil, err
	}
	solver := &Mediator{
		controller: controller,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (mediator *Mediator) Start(ctx context.Context, cm *system.CleanupManager) error {
	return mediator.controller.Start(ctx, cm)
}
