package resourceprovider

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/executor"
	"github.com/bacalhau-project/lilypad/pkg/executor/bacalhau"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

// this configures the resource offers we will keep track of
type ResourceProviderOfferOptions struct {
	// if we are configuring a single machine then
	// these values are populated by the flags
	OfferSpec data.MachineSpec
	// we can dupliate the single spec to create a list of specs
	OfferCount int
	// this represents how many machines we will keep
	// offering to the network
	// we can configure this with a config file
	// to start with we will just add --cpu --gpu and --ram flags
	// to the resource provider CLI which constrains them to a single machine
	Specs []data.MachineSpec
	// the list of modules we are willing to run
	// an empty list means anything
	Modules []string

	// this will normally be FixedPrice for RP's
	Mode data.PricingMode

	// the default pricing for this resource provider
	// for all modules that don't have a specific price
	DefaultPricing  data.DealPricing
	DefaultTimeouts data.DealTimeouts

	// allow different pricing for different modules
	ModulePricing  map[string]data.DealPricing
	ModuleTimeouts map[string]data.DealTimeouts

	// which mediators and directories this RP will trust
	Services data.ServiceConfig
}

type ResourceProviderOptions struct {
	Bacalhau bacalhau.BacalhauExecutorOptions
	Offers   ResourceProviderOfferOptions
	Web3     web3.Web3Options
}

type ResourceProvider struct {
	web3SDK    *web3.Web3SDK
	options    ResourceProviderOptions
	controller *ResourceProviderController
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
) (*ResourceProvider, error) {
	controller, err := NewResourceProviderController(options, web3SDK, executor)
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
