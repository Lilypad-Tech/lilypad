package jobcreator

import (
	"context"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"go.opentelemetry.io/otel/trace"
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
	// which node(s) (if any) to target
	Target data.TargetConfig
}

type JobCreatorOptions struct {
	Mediation JobCreatorMediationOptions
	Offer     JobCreatorOfferOptions
	Web3      web3.Web3Options
	Telemetry system.TelemetryOptions
}

type JobCreator interface {
	Start(ctx context.Context, cm *system.CleanupManager) chan error
	GetJobOfferFromOptions(options JobCreatorOfferOptions) (data.JobOffer, error)
	AddJobOffer(offer data.JobOffer) (data.JobOfferContainer, error)
	SubscribeToJobOfferUpdates(sub JobOfferSubscriber) func()
	SubscribeToJobOfferUpdatesWithFilter(sub JobOfferSubscriber, jobID string) func()
	GetResult(dealId string) (data.Result, error)
}

type BasicJobCreator struct {
	controller *JobCreatorController
	web3SDK    *web3.Web3SDK
}

func NewJobCreator(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
	tracer trace.Tracer,
) (*BasicJobCreator, error) {
	controller, err := NewJobCreatorController(options, web3SDK, tracer)
	if err != nil {
		return nil, err
	}
	return &BasicJobCreator{
		controller: controller,
		web3SDK:    web3SDK,
	}, nil
}

func (jobCreator *BasicJobCreator) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return jobCreator.controller.Start(ctx, cm)
}

func (jobCreator *BasicJobCreator) GetJobOfferFromOptions(options JobCreatorOfferOptions) (data.JobOffer, error) {
	return getJobOfferFromOptions(options, jobCreator.web3SDK.GetAddress().String())
}

// adds the job offer to the solver
func (jobCreator *BasicJobCreator) AddJobOffer(offer data.JobOffer) (data.JobOfferContainer, error) {
	return jobCreator.controller.AddJobOffer(offer)
}

func (jobCreator *BasicJobCreator) SubscribeToJobOfferUpdates(sub JobOfferSubscriber) func() {
	return jobCreator.controller.SubscribeToJobOfferUpdates(sub)
}

func (jobCreator *BasicJobCreator) SubscribeToJobOfferUpdatesWithFilter(sub JobOfferSubscriber, jobID string) func() {
	return jobCreator.controller.SubscribeToJobOfferUpdatesWithFilter(sub, jobID)
}

func (jobCreator *BasicJobCreator) GetResult(dealId string) (data.Result, error) {
	return jobCreator.controller.solverClient.GetResult(dealId)
}
