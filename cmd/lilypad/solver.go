package lilypad

import (
	"os"
	"os/signal"

	"github.com/bacalhau-project/lilypad/pkg/solver"
	memorystore "github.com/bacalhau-project/lilypad/pkg/solver/store/memory"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	return solver.SolverOptions{
		Server: getDefaultServerOptions(),
		Web3:   getDefaultWeb3Options(),
	}
}

func newSolverCmd() *cobra.Command {
	options := NewSolverOptions()

	solverCmd := &cobra.Command{
		Use:     "solver",
		Short:   "Start the lilypad solver service.",
		Long:    "Start the lilypad solver service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runSolver(cmd, options)
		},
	}

	addServerCliFlags(solverCmd, options.Server)
	addWeb3CliFlags(solverCmd, options.Web3)

	return solverCmd
}

func runSolver(cmd *cobra.Command, options solver.SolverOptions) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	cm := system.NewCleanupManager()
	defer cm.Cleanup(cmd.Context())
	ctx := cmd.Context()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	solverStore, err := memorystore.NewSolverStoreMemory()
	if err != nil {
		return err
	}

	solver, err := solver.NewSolver(options, solverStore, web3SDK)
	if err != nil {
		return err
	}

	err = solver.Start(ctx, cm)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
