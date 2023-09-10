package lilypad

import (
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/resourceprovider"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newResourceProviderCmd() *cobra.Command {
	options := optionsfactory.NewResourceProviderOptions()

	resourceProviderCmd := &cobra.Command{
		Use:     "resource-provider",
		Short:   "Start the lilypad resource-provider service.",
		Long:    "Start the lilypad resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options.Offers = optionsfactory.ProcessResourceProviderOfferOptions(options.Offers)
			return runResourceProvider(cmd, options)
		},
	}

	addWeb3CliFlags(resourceProviderCmd, options.Web3)
	addResourceProviderOfferCliFlags(resourceProviderCmd, options.Offers)

	return resourceProviderCmd
}

func runResourceProvider(cmd *cobra.Command, options resourceprovider.ResourceProviderOptions) error {
	err := optionsfactory.CheckWeb3Options(options.Web3, true)
	if err != nil {
		return err
	}
	err = optionsfactory.CheckResourceProviderOfferOptions(options.Offers)
	if err != nil {
		return err
	}
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	resourceProviderService, err := resourceprovider.NewResourceProvider(options, web3SDK)
	if err != nil {
		return err
	}

	errChan := resourceProviderService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-errChan:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
