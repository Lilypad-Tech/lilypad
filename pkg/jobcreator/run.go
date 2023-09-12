package jobcreator

import (
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
)

func RunJob(ctx *system.CommandContext, options JobCreatorOptions) error {
	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := NewJobCreator(options, web3SDK)
	if err != nil {
		return err
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
		return err
	}

	_, err = jobCreatorService.AddJobOffer(offer)
	if err != nil {
		return err
	}

	return nil
}
