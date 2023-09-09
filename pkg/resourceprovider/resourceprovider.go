package resourceprovider

import (
	"context"
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type ResourceProviderOptions struct {
	Web3 web3.Web3Options
}

type ResourceProvider struct {
	web3SDK    *web3.Web3SDK
	controller *ResourceProviderController
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
) (*ResourceProvider, error) {
	if options.Web3.SolverAddress == "" {
		return nil, fmt.Errorf("--web3-solver-address or WEB3_SOLVER_ADDRESS is empty")
	}
	controller, err := NewResourceProviderController(web3SDK)
	if err != nil {
		return nil, err
	}
	solver := &ResourceProvider{
		controller: controller,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (resourceProvider *ResourceProvider) Start(ctx context.Context, cm *system.CleanupManager) error {
	return resourceProvider.controller.Start(ctx, cm)
}
