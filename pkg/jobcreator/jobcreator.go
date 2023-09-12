package jobcreator

import (
	"context"
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/module"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type JobCreatorOfferOptions struct {
	// the module that is wanting to be run
	// this contains the spec that is required to run the module
	Module data.ModuleConfig
	// the required spec hoisted from the module
	Spec data.MachineSpec
	// this is so clients can put limit orders for jobs
	// and the solver will match as soon as a resource offer
	// is added that matches the bid
	Pricing data.Pricing
	// the inputs to the module
	Inputs map[string]string
}

type JobCreatorOptions struct {
	Offer JobCreatorOfferOptions
	Web3  web3.Web3Options
}

type JobCreator struct {
	web3SDK    *web3.Web3SDK
	options    JobCreatorOptions
	controller *JobCreatorController
}

func NewJobCreator(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
) (*JobCreator, error) {
	if options.Web3.SolverAddress == "" {
		return nil, fmt.Errorf("--web3-solver-address or WEB3_SOLVER_ADDRESS is empty")
	}
	controller, err := NewJobCreatorController(options, web3SDK)
	if err != nil {
		return nil, err
	}
	solver := &JobCreator{
		controller: controller,
		options:    options,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (jobCreator *JobCreator) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return jobCreator.controller.Start(ctx, cm)
}

// this will load the module in the offer options
// and hoist the machine spec from the module into the offer
func (jobCreator *JobCreator) GetJobOfferFromOptions(options JobCreatorOfferOptions) (data.JobOffer, error) {
	// process the given module so we know what spec the job is asking for
	// this will also validate the module the user is asking for
	loadedModule, err := module.LoadModule(options.Module, options.Inputs)
	if err != nil {
		return data.JobOffer{}, fmt.Errorf("error loading module: %s", err.Error())
	}

	return data.JobOffer{
		JobCreator: jobCreator.web3SDK.GetAddress().String(),
		Module:     options.Module,
		Spec:       loadedModule.Machine,
		Inputs:     options.Inputs,
		// are will pay whatever the market is offering
		// TODO: put price limits here
		Pricing: data.Pricing{
			Mode: data.MarketPrice,
		},
	}, nil
}

func (JobCreator *JobCreator) AddOffer(ctx context.Context, cm *system.CleanupManager) chan error {
	return JobCreator.controller.Start(ctx, cm)
}
