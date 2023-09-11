package lilypad

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/davecgh/go-spew/spew"
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
			return runJobCreator(cmd, options)
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

	jobCreatorService, err := jobcreator.NewJobCreator(options, web3SDK)
	if err != nil {
		return err
	}

	jobCreatorErrors := jobCreatorService.Start(commandCtx.Ctx, commandCtx.Cm)

	fmt.Printf("HERE --------------------------------------\n")
	spew.Dump("HERE")
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
