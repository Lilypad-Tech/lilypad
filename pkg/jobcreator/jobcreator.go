package jobcreator

import (
	"context"
	"fmt"
	"time"

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
	// this will normally be MarketPrice for JC's
	Mode data.PricingMode
	// this is so clients can put limit orders for jobs
	// and the solver will match as soon as a resource offer
	// is added that matches the bid
	Pricing data.DealPricing
	// the timeouts we are offering with the deal
	Timeouts data.DealTimeouts
	// the inputs to the module
	Inputs map[string]string
	// which mediators and directories this RP will trust
	TrustedParties data.TrustedParties
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
		// assign CreatedAt to the current millisecond timestamp
		CreatedAt:      int(time.Now().UnixNano() / int64(time.Millisecond)),
		JobCreator:     jobCreator.web3SDK.GetAddress().String(),
		Module:         options.Module,
		Spec:           loadedModule.Machine,
		Inputs:         options.Inputs,
		Mode:           options.Mode,
		Pricing:        options.Pricing,
		Timeouts:       options.Timeouts,
		TrustedParties: options.TrustedParties,
	}, nil
}

// adds the job offer to the solver
func (jobCreator *JobCreator) AddJobOffer(offer data.JobOffer) (data.JobOffer, error) {
	return jobCreator.controller.AddJobOffer(offer)
}
