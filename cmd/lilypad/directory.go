package lilypad

import (
	"github.com/bacalhau-project/lilypad/pkg/directory"
	memorystore "github.com/bacalhau-project/lilypad/pkg/directory/store/memory"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func NewDirectoryOptions() directory.DirectoryOptions {
	return directory.DirectoryOptions{
		Server: getDefaultServerOptions(),
		Web3:   getDefaultWeb3Options(),
	}
}

func newDirectoryCmd() *cobra.Command {
	options := NewDirectoryOptions()

	solverCmd := &cobra.Command{
		Use:     "directory",
		Short:   "Start the lilypad directory service.",
		Long:    "Start the lilypad directory service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDirectory(cmd, options)
		},
	}

	addServerCliFlags(solverCmd, options.Server)
	addWeb3CliFlags(solverCmd, options.Web3)

	return solverCmd
}

func runDirectory(cmd *cobra.Command, options directory.DirectoryOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	directoryStore, err := memorystore.NewDirectoryStoreMemory()
	if err != nil {
		return err
	}

	solver, err := directory.NewDirectory(options, directoryStore, web3SDK)
	if err != nil {
		return err
	}

	err = solver.Start(commandCtx.Ctx, commandCtx.Cm)
	if err != nil {
		return err
	}

	<-commandCtx.Ctx.Done()
	return nil
}
