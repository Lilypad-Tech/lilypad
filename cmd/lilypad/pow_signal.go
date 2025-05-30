package lilypad

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/options"
	optionsfactory "github.com/Lilypad-Tech/lilypad/v2/pkg/options"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/trace/noop"
)

func newPowSignalCmd() *cobra.Command {
	options := optionsfactory.NewPowSignalOptions()

	powSignalCmd := &cobra.Command{
		Use:     "pow-signal",
		Short:   "Send a pow signal to smart contract.",
		Long:    "Send a pow signal to smart contract.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			network, err := cmd.Flags().GetString("network")
			if err != nil {
				log.Error().Err(err).Msg("Failed to retrieve 'network' flag")
				return err
			}

			options, err := optionsfactory.ProcessPowSignalOptions(options, network)
			if err != nil {
				log.Error().Err(err).Msg("Failed to process PowSignal options")
				return err
			}
			cmd.SilenceUsage = true

			return runPowSignal(cmd, options)
		},
	}

	optionsfactory.AddPowSignalCliFlags(powSignalCmd, &options)

	return powSignalCmd
}

func runPowSignal(cmd *cobra.Command, options options.PowSignalOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	noopTracer := noop.NewTracerProvider().Tracer(system.GetOTelServiceName(system.DefaultService))
	system.SetupGlobalLogger(system.DefaultService, nil)

	web3SDK, err := web3.NewContractSDK(commandCtx.Ctx, options.Web3, noopTracer)
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize Web3 SDK")
		return err
	}

	_, err = web3SDK.SendPowSignal(commandCtx.Ctx)
	if err != nil {
		return err
	}
	log.Info().Msg("Send pow signal successful.")
	return nil
}
