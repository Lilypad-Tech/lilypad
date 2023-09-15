package jobcreator

import (
	"fmt"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
)

func RunJob(ctx *system.CommandContext, options JobCreatorOptions) (*data.Result, error) {
	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return nil, err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := NewJobCreator(options, web3SDK)
	if err != nil {
		return nil, err
	}

	jobCreatorErrors := jobCreatorService.Start(ctx.Ctx, ctx.Cm)

	// start the error channels in a goroutine
	// because we want to block on the actual job we are running
	go func() {
		for {
			select {
			case err = <-jobCreatorErrors:
				log.Error().Err(err).Msg("error in job creator")
				ctx.Cleanup()
				return
			case <-ctx.Ctx.Done():
				return
			}
		}
	}()

	// let's process our options into an actual job offer
	// this will also validate the module we are asking for
	offer, err := jobCreatorService.GetJobOfferFromOptions(options.Offer)
	if err != nil {
		return nil, err
	}

	// wait a short period because we've just started the job creator service
	time.Sleep(100 * time.Millisecond)

	jobOfferContainer, err := jobCreatorService.AddJobOffer(offer)
	if err != nil {
		return nil, err
	}

	jobCreatorService.SubscribeToJobOffers(func(evOffer data.JobOfferContainer) {
		if evOffer.JobOffer.ID != jobOfferContainer.ID {
			return
		}

		updatedState := data.GetAgreementStateString(evOffer.State)

		fmt.Printf("evOffer --------------------------------------\n")
		spew.Dump(updatedState)

	})

	checkJob := func() (bool, error) {
		return false, nil
	}

	// now we wait on the state of the job
	for {
		select {
		// this means the job was cancelled
		case <-ctx.Ctx.Done():
			return nil, nil
		// keep looping on the job state
		case <-time.After(1 * time.Second):
			complete, err := checkJob()
			if err != nil {
				return nil, err
			}
			if complete {
				return nil, nil
			}
		}
	}
}
