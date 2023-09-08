package resourceprovider

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type ResourceProviderOptions struct {
	Web3 web3.Web3Options
}

type ResourceProvider struct {
	web3SDK    *web3.ContractSDK
	controller *ResourceProviderController
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.ContractSDK,
) (*ResourceProvider, error) {
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
