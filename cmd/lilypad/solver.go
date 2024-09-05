package lilypad

import (
	"fmt"

	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/solver"
	memorystore "github.com/lilypad-tech/lilypad/pkg/solver/store/memory"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
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
			return runSolver(cmd, options, network)
		},
	}

	optionsfactory.AddSolverCliFlags(solverCmd, &options)

	return solverCmd
}

func runSolver(cmd *cobra.Command, options solver.SolverOptions, network string) error {
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

	tc := system.TelemetryConfig{
		TelemetryURL:   options.Telemetry.URL,
		TelemetryToken: options.Telemetry.Token,
		Enabled:        !options.Telemetry.Disable,
		Service:        system.SolverService,
		Network:        network,
		Address:        web3SDK.GetAddress().String(),
		GPU:            system.GetGPUInfo(),
	}
	telemetry, err := system.SetupOTelSDK(commandCtx.Ctx, tc)
	if err != nil {
		fmt.Printf("failed to setup opentelemetry: %s", err)
	}
	commandCtx.Cm.RegisterCallbackWithContext(telemetry.Shutdown)

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
