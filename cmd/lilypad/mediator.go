package lilypad

import (
	"github.com/bacalhau-project/lilypad/pkg/executor/bacalhau"
	"github.com/bacalhau-project/lilypad/pkg/mediator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
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
			options, err := optionsfactory.ProcessMediatorOptions(options)
			if err != nil {
				return err
			}
			return runMediator(cmd, options)
		},
	}

	optionsfactory.AddMediatorCliFlags(mediatorCmd, &options)

	return mediatorCmd
}

func runMediator(cmd *cobra.Command, options mediator.MediatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
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
