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

	resourceProviderCmd := &cobra.Command{
		Use: "resource-provider",
		RunE: func(cmd *cobra.Command, _ []string) error {
			optionsfactory.CheckDeprecation(options.Offers.Services, options.Web3)

			//	allowlist, _ := cmd.Flags().GetBool("enable-allowlist")

			options, err := optionsfactory.NewAllowlistOptions().AllowlistCmd(options, allowlist)
			if err != nil {
				return err
			}
			return runResourceProvider(cmd, options)
		},
	}

	optionsfactory.AddResourceProviderCliFlags(resourceProviderCmd, &options)

	return resourceProviderCmd
}

func AllowlistToggle(cmd *cobra.Command, options AllowlistOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	resourceprovider.AllowlistCmd(options.AllowlistCmd)
	return nil
}
