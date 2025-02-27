package lilypad

import (
	"fmt"

	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/solver"
	"github.com/lilypad-tech/lilypad/pkg/solver/stats"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	db "github.com/lilypad-tech/lilypad/pkg/solver/store/db"
	memorystore "github.com/lilypad-tech/lilypad/pkg/solver/store/memory"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	zerolog "github.com/rs/zerolog/log"
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
			network, _ := cmd.Flags().GetString("network")
			options, err := optionsfactory.ProcessSolverOptions(options, network)
			if err != nil {
				return err
			}
			cmd.SilenceUsage = true

			return runSolver(cmd, options, network)
		},
	}

	optionsfactory.AddSolverCliFlags(solverCmd, &options)

	return solverCmd
}

func runSolver(cmd *cobra.Command, options solver.SolverOptions, network string) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	telemetry, err := configureTelemetry(commandCtx.Ctx, system.SolverService, network, options.Telemetry, &options.Metrics, options.Web3)
	if err != nil {
		zerolog.Warn().Msgf("failed to setup opentelemetry: %s", err)
	}
	commandCtx.Cm.RegisterCallbackWithContext(telemetry.Shutdown)
	tracer := telemetry.TracerProvider.Tracer(system.GetOTelServiceName(system.SolverService))
	meter := telemetry.MeterProvider.Meter(system.GetOTelServiceName(system.SolverService))
	log := system.NewServiceLogger(system.SolverService)

	unregisterMetrics, err := system.NewMetrics(meter)
	if err != nil {
		zerolog.Warn().Msgf("failed to start system metrics: %s", err)
	}
	commandCtx.Cm.RegisterCallback(unregisterMetrics)

	web3SDK, err := web3.NewContractSDK(commandCtx.Ctx, options.Web3, tracer)
	if err != nil {
		return err
	}

	solverStore, err := getSolverStore(options.Store)
	if err != nil {
		return err
	}

	stats, err := stats.NewStats(system.SolverService, options.Stats, options.Web3, web3SDK)
	if err != nil {
		return err
	}

	solverService, err := solver.NewSolver(options, solverStore, web3SDK, stats, tracer, meter, log)
	if err != nil {
		return err
	}

	solverErrors := solverService.Start(commandCtx.Ctx, commandCtx.Cm, telemetry.TracerProvider)

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

func getSolverStore(options store.StoreOptions) (store.SolverStore, error) {
	var solverStore store.SolverStore
	var err error

	switch options.Type {
	case "database":
		solverStore, err = db.NewSolverStoreDatabase(options.ConnStr, options.GormLogLevel)
		if err != nil {
			return nil, err
		}
	case "memory":
		solverStore, err = memorystore.NewSolverStoreMemory()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("expected solver store type database or memory, but received: %s", options.Type)
	}

	return solverStore, nil
}
