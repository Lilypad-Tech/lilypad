package lilypad

import (
	"github.com/bacalhau-project/lilypad/pkg/mediator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newMediatorCmd() *cobra.Command {
	options := optionsfactory.NewMediatorOptions()

	solverCmd := &cobra.Command{
		Use:     "mediator",
		Short:   "Start the lilypad mediator service.",
		Long:    "Start the lilypad mediator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runMediator(cmd, options)
		},
	}

	addWeb3CliFlags(solverCmd, options.Web3)

	return solverCmd
}

func runMediator(cmd *cobra.Command, options mediator.MediatorOptions) error {
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

	solver, err := mediator.NewMediator(options, web3SDK)
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
