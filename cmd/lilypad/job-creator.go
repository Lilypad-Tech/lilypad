package lilypad

import (
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newJobCreatorCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()

	jobCreatorCmd := &cobra.Command{
		Use:     "job-creator",
		Short:   "Start the lilypad resource-provider service.",
		Long:    "Start the lilypad resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			newWeb3Options, err := optionsfactory.ProcessWeb3Options(options.Web3)
			if err != nil {
				return err
			}
			options.Web3 = newWeb3Options

			newOfferOptions, err := optionsfactory.ProcessJobCreatorOfferOptions(options.Offer)
			if err != nil {
				return err
			}
			options.Offer = newOfferOptions

			// check the options
			err = optionsfactory.CheckWeb3Options(options.Web3, true)
			if err != nil {
				return err
			}
			err = optionsfactory.CheckJobCreatorOfferOptions(options.Offer)
			if err != nil {
				return err
			}
			return runJobCreator(cmd, options)
		},
	}

	optionsfactory.AddWeb3CliFlags(jobCreatorCmd, options.Web3)
	optionsfactory.AddJobCreatorOfferCliFlags(jobCreatorCmd, options.Offer)

	return jobCreatorCmd
}

func runJobCreator(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	jobCreatorService, err := jobcreator.NewJobCreator(options, web3SDK)
	if err != nil {
		return err
	}

	jobCreatorErrors := jobCreatorService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-jobCreatorErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
