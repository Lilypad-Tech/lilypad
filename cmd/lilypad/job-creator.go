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

	solverCmd := &cobra.Command{
		Use:     "job-creator",
		Short:   "Start the lilypad resource-provider service.",
		Long:    "Start the lilypad resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runJobCreator(cmd, options)
		},
	}

	addWeb3CliFlags(solverCmd, options.Web3)

	return solverCmd
}

func runJobCreator(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	err := optionsfactory.CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	solver, err := jobcreator.NewJobCreator(options, web3SDK)
	if err != nil {
		return err
	}

	err = solver.Start(commandCtx.Ctx, commandCtx.Cm)
	if err != nil {
		return err
	}

	<-commandCtx.Ctx.Done()
	return nil
}
