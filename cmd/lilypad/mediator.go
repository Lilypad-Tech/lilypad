package lilypad

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/executor/bacalhau"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/mediator"
	optionsfactory "github.com/Lilypad-Tech/lilypad/v2/pkg/options"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/trace/noop"
)

func newMediatorCmd() *cobra.Command {
	options := optionsfactory.NewMediatorOptions()

	mediatorCmd := &cobra.Command{
		Use:     "mediator",
		Short:   "Start the lilypad mediator service.",
		Long:    "Start the lilypad mediator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			network, _ := cmd.Flags().GetString("network")
			options, err := optionsfactory.ProcessMediatorOptions(options, network)
			if err != nil {
				return err
			}
			cmd.SilenceUsage = true

			return runMediator(cmd, options)
		},
	}

	optionsfactory.AddMediatorCliFlags(mediatorCmd, &options)

	return mediatorCmd
}

func runMediator(cmd *cobra.Command, options mediator.MediatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	noopTracer := noop.NewTracerProvider().Tracer(system.GetOTelServiceName(system.MediatorService))
	system.SetupGlobalLogger(system.MediatorService, nil)

	web3SDK, err := web3.NewContractSDK(commandCtx.Ctx, options.Web3, noopTracer)
	if err != nil {
		return err
	}

	executor, err := bacalhau.NewBacalhauExecutor(options.Bacalhau)
	if err != nil {
		return err
	}

	mediatorService, err := mediator.NewMediator(options, web3SDK, executor)
	if err != nil {
		return err
	}

	log.Debug().Msgf("Starting mediator service.")
	mediatorErrors := mediatorService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-mediatorErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
