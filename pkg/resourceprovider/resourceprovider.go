package resourceprovider

import (
	"context"
	"errors"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/executor/bacalhau"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/powtoken"
	"github.com/rs/zerolog/log"
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

// this configures the resource pow we will keep track of
type ResourceProviderPowOptions struct {
	EnablePow bool
}

type ResourceProviderOptions struct {
	Bacalhau bacalhau.BacalhauExecutorOptions
	Offers   ResourceProviderOfferOptions
	Web3     web3.Web3Options
	Pow      ResourceProviderPowOptions
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

func (resourceProvider *ResourceProvider) StartMineLoop(ctx context.Context, cm *system.CleanupManager) error {
	walletAddress := resourceProvider.web3SDK.GetAddress()
	nodeId, err := resourceProvider.controller.executor.Id()
	if err != nil {
		return err
	}

	tasks := make(chan task)
	resourceProvider.controller.web3Events.Pow.SubscribenewPowRound(func(newPowRound powtoken.PowtokenNewPowRound) {
		_, challenge, err := resourceProvider.web3SDK.GetGenerateChallenge(ctx, nodeId)
		if err != nil {
			log.Err(err).Msgf("unable to fetch challenge")
			return
		}
		tasks <- task{
			challenge:  challenge.Challenge,
			difficulty: challenge.Difficult,
		}

	})

	w := NewWorker(nodeId, walletAddress)
	for {
		select {
		case <-ctx.Done():
			return errors.New("cancel by context")
		case task := <-tasks:
			w.Reset()
			go func() {
				nonce, err := w.Solve(ctx, &task)
				if err != nil {
					log.Err(err).Msgf("solve challenge nonce fail")
					return
				}

				_, submission, err := resourceProvider.web3SDK.SubmitWork(ctx, nonce, nodeId)
				if err != nil {
					log.Err(err).Msgf("submit work fail")
					return
				}
				log.Info().Str("address", submission.WalletAddress.String()).
					Str("nodeid", submission.NodeId).
					Str("Nonce", submission.Nonce.String()).
					Msgf("miner success")
			}()
		}
	}
}
