package lilypad

import (
	"github.com/lilypad-tech/lilypad/pkg/options"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newPowSignalCmd() *cobra.Command {
	options := optionsfactory.NewPowSignalOptions()

	powSignalCmd := &cobra.Command{
		Use:     "pow-signal",
		Short:   "Send a pow signal to smart contract.",
		Long:    "Send a pow signal to smart contract.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			network, _ := cmd.Flags().GetString("network")

			options, err := optionsfactory.ProcessPowSignalOptions(options, network)
			if err != nil {
				return err
			}
			return runPowSignal(cmd, options)
		},
	}

	optionsfactory.AddPowSignalCliFlags(powSignalCmd, &options)

	return powSignalCmd
}

func runPowSignal(cmd *cobra.Command, options options.PowSignalOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	_, err = web3SDK.SendPowSignal(commandCtx.Ctx)
	if err != nil {
		return err
	}
	log.Info().Msgf("send pow signal successful.")
	return nil
}
