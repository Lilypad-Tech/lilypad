package lilypad

import (
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/spf13/cobra"
)

func newAllowlistCmd() *cobra.Command {
	options := optionsfactory.NewResourceProviderOptions()

	resourceProviderCmd := &cobra.Command{
		Use:     "resource-provider",
		Short:   "Start the lilypad resource-provider service.",
		Long:    "Start the lilypad resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			optionsfactory.CheckDeprecation(options.Offers.Services, options.Web3)

			allowlist, _ := cmd.Flags().GetString("network")
			//newAllowlistCmd, _ := cmd.Flags().GetBool("enable-allowlist")
			options, err := optionsfactory.ProcessResourceProviderOptions(options, allowlist)
			if err != nil {
				return err
			}
			return runResourceProvider(cmd, options)
		},
	}

	optionsfactory.AddResourceProviderCliFlags(resourceProviderCmd, &options)

	return resourceProviderCmd
}
