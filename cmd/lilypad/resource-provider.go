package lilypad

import (
	"os"
	"os/signal"

	"github.com/bacalhau-project/lilypad/pkg/resourceprovider"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewResourceProviderOptions() resourceprovider.ResourceProviderOptions {
	return resourceprovider.ResourceProviderOptions{
		Web3: getDefaultWeb3Options(),
	}
}

func newResourceProviderCmd() *cobra.Command {
	options := NewResourceProviderOptions()

	solverCmd := &cobra.Command{
		Use:     "resource-provider",
		Short:   "Start the lilypad resource-provider service.",
		Long:    "Start the lilypad resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runResourceProvider(cmd, options)
		},
	}

	addWeb3CliFlags(solverCmd, options.Web3)

	return solverCmd
}

func runResourceProvider(cmd *cobra.Command, options resourceprovider.ResourceProviderOptions) error {
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

	solver, err := resourceprovider.NewResourceProvider(options, web3SDK)
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
