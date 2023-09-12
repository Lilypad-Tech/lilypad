package lilypad

import (
	"github.com/bacalhau-project/lilypad/pkg/directory"
	memorystore "github.com/bacalhau-project/lilypad/pkg/directory/store/memory"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newDirectoryCmd() *cobra.Command {
	options := optionsfactory.NewDirectoryOptions()

	directoryCmd := &cobra.Command{
		Use:     "directory",
		Short:   "Start the lilypad directory service.",
		Long:    "Start the lilypad directory service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			newWeb3Options, err := optionsfactory.ProcessWeb3Options(options.Web3)
			if err != nil {
				return err
			}
			options.Web3 = newWeb3Options
			return runDirectory(cmd, options)
		},
	}

	optionsfactory.AddServerCliFlags(directoryCmd, options.Server)
	optionsfactory.AddWeb3CliFlags(directoryCmd, options.Web3)

	return directoryCmd
}

func runDirectory(cmd *cobra.Command, options directory.DirectoryOptions) error {
	err := optionsfactory.CheckWeb3Options(options.Web3, false)
	if err != nil {
		return err
	}
	err = optionsfactory.CheckServerOptions(options.Server)
	if err != nil {
		return err
	}
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

	directoryService, err := directory.NewDirectory(options, directoryStore, web3SDK)
	if err != nil {
		return err
	}

	directoryErrors := directoryService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-directoryErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
