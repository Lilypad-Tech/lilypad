package resourceprovider

import (
	"context"
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

// this configures the resource offers we will keep track of
type ResourceProviderOfferOptions struct {
	// if we are configuring a single machine then
	// these values are populated by the flags
	Machine data.Spec
	// we can dupliate the single spec to create a list of specs
	MachineCount int
	// this represents how many machines we will keep
	// offering to the network
	// we can configure this with a config file
	// to start with we will just add --cpu --gpu and --ram flags
	// to the resource provider CLI which constrains them to a single machine
	Specs []data.Spec
	// the list of modules we are willing to run
	// an empty list means anything
	Modules []string
}

type ResourceProviderOptions struct {
	Offers ResourceProviderOfferOptions
	Web3   web3.Web3Options
}

type ResourceProvider struct {
	web3SDK    *web3.Web3SDK
	options    ResourceProviderOptions
	controller *ResourceProviderController
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
) (*ResourceProvider, error) {
	if options.Web3.SolverAddress == "" {
		return nil, fmt.Errorf("--web3-solver-address or WEB3_SOLVER_ADDRESS is empty")
	}
	controller, err := NewResourceProviderController(options, web3SDK)
	if err != nil {
		return nil, err
	}
	solver := &ResourceProvider{
		controller: controller,
		options:    options,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (resourceProvider *ResourceProvider) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return resourceProvider.controller.Start(ctx, cm)
}
