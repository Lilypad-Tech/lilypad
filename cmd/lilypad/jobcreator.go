package lilypad

import (
	"github.com/lilypad-tech/lilypad/pkg/jobcreator"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newJobCreatorCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()

	solverCmd := &cobra.Command{
		Use:     "jobcreator",
		Short:   "Start the lilypad job creator service.",
		Long:    "Start the lilypad job creator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {

			network, _ := cmd.Flags().GetString("network")
			lilynext, _ := cmd.Flags().GetBool("lilynext")
			options, err := optionsfactory.ProcessOnChainJobCreatorOptions(options, args, network)
			if err != nil {
				return err
			}
			return runJobCreator(cmd, options, network, lilynext)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(solverCmd, &options)

	return solverCmd
}

func runJobCreator(cmd *cobra.Command, options jobcreator.JobCreatorOptions, network string, lilynext bool) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	if lilynext {
		log.Info().Msg("🍃 Running the new lilypad protocol")
	}

	telemetry, err := configureTelemetry(commandCtx.Ctx, system.JobCreatorService, network, options.Telemetry, nil, nil, options.Web3)
	if err != nil {
		log.Warn().Msgf("failed to setup opentelemetry: %s", err)
	}
	commandCtx.Cm.RegisterCallbackWithContext(telemetry.Shutdown)
	tracer := telemetry.TracerProvider.Tracer(system.GetOTelServiceName(system.JobCreatorService))
	system.SetupGlobalLogger(system.JobCreatorService, nil)

	web3SDK, err := web3.NewContractSDK(commandCtx.Ctx, options.Web3, tracer)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobcreator.NewOnChainJobCreator(options, web3SDK, tracer)
	if err != nil {
		return err
	}

	jobCreatorErrors := jobCreatorService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-jobCreatorErrors:
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
