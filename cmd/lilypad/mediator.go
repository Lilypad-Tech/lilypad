package lilypad

import (
	"github.com/bacalhau-project/lilypad/pkg/mediator"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func NewMediatorOptions() mediator.MediatorOptions {
	return mediator.MediatorOptions{
		Web3: getDefaultWeb3Options(),
	}
}

func newMediatorCmd() *cobra.Command {
	options := NewMediatorOptions()

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
