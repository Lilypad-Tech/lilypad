package lilypad

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()

	runCmd := &cobra.Command{
		Use:     "run",
		Short:   "Run a job on the Lilypad network.",
		Long:    "Run a job on the Lilypad network.",
		Example: "run cowsay v0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			version := args[1]

			newWeb3Options, err := optionsfactory.ProcessWeb3Options(options.Web3)
			if err != nil {
				return err
			}
			options.Web3 = newWeb3Options

			// map the options
			newOfferOptions, err := optionsfactory.ProcessJobCreatorOfferOptions(options.Offer)
			if err != nil {
				return err
			}
			options.Offer = newOfferOptions

			if name != "" {
				options.Offer.Module.Name = name
			}

			if version != "" {
				options.Offer.Module.Version = version
			}

			// check the options
			err = optionsfactory.CheckWeb3Options(options.Web3, true)
			if err != nil {
				return err
			}
			err = optionsfactory.CheckJobCreatorOfferOptions(options.Offer)
			if err != nil {
				return err
			}
			return runJob(cmd, options)
		},
	}

	optionsfactory.AddWeb3CliFlags(runCmd, options.Web3)
	optionsfactory.AddJobCreatorOfferCliFlags(runCmd, options.Offer)

	return runCmd
}

func runJob(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobcreator.NewJobCreator(options, web3SDK)
	if err != nil {
		return err
	}

	// let's process our options into an actual job offer
	// this will also validate the module we are asking for
	offer, err := jobCreatorService.GetJobOfferFromOptions(options.Offer)
	if err != nil {
		return err
	}

	jobCreatorErrors := jobCreatorService.Start(commandCtx.Ctx, commandCtx.Cm)

	// start the error channels in a goroutine
	// because we want to block on the actual job we are running
	go func() {
		for {
			select {
			case err = <-jobCreatorErrors:
				log.Error().Err(err).Msg("error in job creator")
				commandCtx.Cleanup()
				return
			case <-commandCtx.Ctx.Done():
				return
			}
		}
	}()

	fmt.Printf(" --------------------------------------\n")
	spew.Dump(offer)

	return nil
}
