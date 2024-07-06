package lilypad

import (
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

type AllowlistOptions struct {
	AllowlistCmd bool
}

func newAllowlistCmd() *cobra.Command {
	options := optionsfactory.NewResourceProviderOptions()

	newAllowlistCmd := &cobra.Command{
		Use:   "resource-provider",
		Short: "Allowlist checker",
		Long:  "Allowlist checker to limit the containers that can be executed",
		// This RunE function is where the command is exectued
		RunE: func(cmd *cobra.Command, _ []string) error {
			optionsfactory.CheckDeprecation(options.Offers.Services, options.Web3)

			// Now you need to retrieve the allowlist flag value

			allowlist, _ := cmd.Flags().GetBool("enable-allowlist")

			//options, err := optionsfactory.NewAllowlistOptions().AllowlistCmd(options, allowlist)
			//if err != nil {
			//		return err
			//	}
			return runResourceProvider(cmd, options)
		},
	}

	optionsfactory.AddResourceProviderCliFlags(newAllowlistCmd, &options)

	return newAllowlistCmd
}

func AllowlistToggle(cmd *cobra.Command, options AllowlistOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	resourceprovider.AllowlistCmd(options.AllowlistCmd)
	return nil
}
