package lilypad

import (
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	memorystore "github.com/bacalhau-project/lilypad/pkg/solver/store/memory"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newSolverCmd() *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	solverCmd := &cobra.Command{
		Use:     "solver",
		Short:   "Start the lilypad solver service.",
		Long:    "Start the lilypad solver service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runSolver(cmd, options)
		},
	}

	optionsfactory.AddServerCliFlags(solverCmd, options.Server)
	optionsfactory.AddWeb3CliFlags(solverCmd, options.Web3)

	return solverCmd
}

func runSolver(cmd *cobra.Command, options solver.SolverOptions) error {
	err := optionsfactory.CheckWeb3Options(options.Web3, false)
	if err != nil {
		return err
	}
	err = optionsfactory.CheckServerOptions(options.Server)
	if err != nil {
		return err
	}
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	solverStore, err := memorystore.NewSolverStoreMemory()
	if err != nil {
		return err
	}

	solverService, err := solver.NewSolver(options, solverStore, web3SDK)
	if err != nil {
		return err
	}

	solverErrors := solverService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-solverErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
