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

	mediatorCmd := &cobra.Command{
		Use:     "mediator",
		Short:   "Start the lilypad mediator service.",
		Long:    "Start the lilypad mediator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			newWeb3Options, err := optionsfactory.ProcessWeb3Options(options.Web3)
			if err != nil {
				return err
			}
			options.Web3 = newWeb3Options
			err = optionsfactory.CheckWeb3Options(options.Web3, false)
			if err != nil {
				return err
			}
			return runMediator(cmd, options)
		},
	}

	optionsfactory.AddWeb3CliFlags(mediatorCmd, &options.Web3)

	return mediatorCmd
}

func runMediator(cmd *cobra.Command, options mediator.MediatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	mediatorService, err := mediator.NewMediator(options, web3SDK)
	if err != nil {
		return err
	}

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
