package lilypad

import (
	"github.com/lilypad-tech/lilypad/pkg/jobcreator"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/trace/noop"
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
			options, err := optionsfactory.ProcessOnChainJobCreatorOptions(options, args, network)
			if err != nil {
				return err
			}
			return runJobCreator(cmd, options)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(solverCmd, &options)

	return solverCmd
}

func runJobCreator(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	noopTracer := noop.NewTracerProvider().Tracer(system.GetOTelServiceName(system.JobCreatorService))
	web3SDK, err := web3.NewContractSDK(commandCtx.Ctx, options.Web3, noopTracer)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobcreator.NewOnChainJobCreator(options, web3SDK)
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
