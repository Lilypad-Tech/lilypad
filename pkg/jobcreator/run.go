package jobcreator

import (
	"errors"
	"fmt"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type RunJobResults struct {
	JobOffer data.JobOfferContainer
	Result   data.Result
}

func RunJob(
	ctx *system.CommandContext,
	options JobCreatorOptions,
	tracer trace.Tracer,
	eventSub JobOfferSubscriber,
) (*RunJobResults, error) {
	web3SDK, err := web3.NewContractSDK(ctx.Ctx, options.Web3, tracer)
	if err != nil {
		return nil, err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := NewJobCreator(options, web3SDK, tracer)
	if err != nil {
		return nil, err
	}

	jobCreatorErrors := jobCreatorService.Start(ctx.Ctx, ctx.Cm)

	// let's process our options into an actual job offer
	// this will also validate the module we are asking for
	offer, err := jobCreatorService.GetJobOfferFromOptions(options.Offer)
	if err != nil {
		return nil, err
	}

	jobOfferID, err := data.GetJobOfferID(offer)
	if err != nil {
		return nil, err
	}
	log.Info().Str("cid", jobOfferID).Msg("sending job offer")

	// wait a short period because we've just started the job creator service
	time.Sleep(100 * time.Millisecond)

	// Start run job trace
	c, span := jobCreatorService.controller.tracer.Start(ctx.Ctx, "run_job",
		trace.WithAttributes(
			attribute.String("job_offer.job_creator", offer.JobCreator),
			attribute.String("job_offer.module.repo", offer.Module.Repo),
			attribute.String("job_offer.module.hash", offer.Module.Hash),
			attribute.String("job_offer.mode", string(offer.Mode)),
		))
	ctx.Ctx = c
	defer span.End()

	updateChan := make(chan data.JobOfferContainer)

	// Set up the subscription BEFORE adding the job offer
	cleanup := jobCreatorService.controller.SubscribeToJobOfferUpdates(func(evOffer data.JobOfferContainer) {
		span.AddEvent("job_offer_update",
			trace.WithAttributes(attribute.String("job_offer_container.state", data.GetAgreementStateString(evOffer.State))))
		updateChan <- evOffer

		// Additionally call the provided eventSub if one was passed in
		if eventSub != nil {
			eventSub(evOffer)
		}
	})

	// Ensure we clean up the subscription when we're done
	defer cleanup()

	// Add the job offer
	span.AddEvent("add_job_offer.start")
	jobOfferContainer, err := jobCreatorService.AddJobOffer(offer)
	if err != nil {
		jobCreatorService.controller.log.Error("failed to add job offer", err)
		span.SetStatus(codes.Error, "failed to add job offer")
		span.RecordError(err)
		return nil, err
	}
	jobCreatorService.controller.log.Debug("job offer container ID", jobOfferContainer.ID)
	span.AddEvent("add_job_offer.done",
		trace.WithAttributes(
			attribute.String("job_offer_container.deal_id", jobOfferContainer.DealID),
			attribute.String("job_offer_container.state", data.GetAgreementStateString(jobOfferContainer.State)),
		))
	span.SetAttributes(attribute.String("job_offer.id", jobOfferContainer.JobOffer.ID),
		attribute.String("deal.id", jobOfferContainer.DealID))

	var finalJobOffer data.JobOfferContainer

	// now we wait on the state of the job
waitloop:
	for {
		select {
		// this means the job was cancelled
		case err := <-jobCreatorErrors:
			span.SetStatus(codes.Error, "job cancelled")
			span.RecordError(err)
			return nil, err
		case <-ctx.Ctx.Done():
			err = errors.New("job cancelled by closed context")
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
			return nil, err
		case receivedOffer := <-updateChan:
			// Ensure we only process updates for our specific job
			if receivedOffer.JobOffer.ID != jobOfferContainer.JobOffer.ID {
				jobCreatorService.controller.log.Debug("Ignoring update for different job", map[string]interface{}{
					"received": receivedOffer.JobOffer.ID,
					"expected": jobOfferContainer.JobOffer.ID,
				})
				continue
			}

			finalJobOffer = receivedOffer
			if data.IsTerminalAgreementState(finalJobOffer.State) {
				break waitloop
			}
		}
	}

	// Check if our job was cancelled
	if finalJobOffer.State == data.GetAgreementStateIndex("JobOfferCancelled") {
		span.SetStatus(codes.Error, "job cancelled")
		span.RecordError(err)
		return nil, fmt.Errorf("job was cancelled")
	}

	// Check if our job timed out
	if finalJobOffer.State == data.GetAgreementStateIndex("JobTimedOut") {
		span.SetStatus(codes.Error, "job timed out")
		span.RecordError(err)
		return nil, fmt.Errorf("job timed out")
	}

	span.AddEvent("get_result.start")
	result, err := jobCreatorService.GetResult(finalJobOffer.DealID)
	if err != nil {
		return nil, err
	}
	span.AddEvent("get_result.done", trace.WithAttributes(attribute.String("result.deal_id", result.DealID)))

	return &RunJobResults{
		JobOffer: finalJobOffer,
		Result:   result,
	}, nil
}
