package jobcreator

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type JobCreatorMediationOptions struct {
	// out of 100 chance we will check results
	CheckResultsPercentage int
}

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
	Services data.ServiceConfig
}

type JobCreatorOptions struct {
	Mediation JobCreatorMediationOptions
	Offer     JobCreatorOfferOptions
	Web3      web3.Web3Options
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
	controller, err := NewJobCreatorController(options, web3SDK)
	if err != nil {
		return nil, err
	}
	jc := &JobCreator{
		controller: controller,
		options:    options,
		web3SDK:    web3SDK,
	}
	return jc, nil
}

func (jobCreator *JobCreator) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return jobCreator.controller.Start(ctx, cm)
}

func (jobCreator *JobCreator) GetJobOfferFromOptions(options JobCreatorOfferOptions) (data.JobOffer, error) {
	return getJobOfferFromOptions(options, jobCreator.web3SDK.GetAddress().String())
}

// adds the job offer to the solver
func (jobCreator *JobCreator) AddJobOffer(offer data.JobOffer) (data.JobOfferContainer, error) {
	return jobCreator.controller.AddJobOffer(offer)
}

func (jobCreator *JobCreator) SubscribeToJobOfferUpdates(sub JobOfferSubscriber) {
	jobCreator.controller.SubscribeToJobOfferUpdates(sub)
}

func (jobCreator *JobCreator) GetResult(dealId string) (data.Result, error) {
	return jobCreator.controller.solverClient.GetResult(dealId)
}
