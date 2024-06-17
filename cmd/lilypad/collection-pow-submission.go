package lilypad

import (
	"github.com/lilypad-tech/lilypad/pkg/collectpowsubmission"
	"github.com/lilypad-tech/lilypad/pkg/options"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newCollectionPowSubmissionCmd() *cobra.Command {
	options := optionsfactory.NewCollectPowSubmissionOptions()

	powSignalCmd := &cobra.Command{
		Use:     "collection-pow-submission",
		Short:   "Collect pow submission data to database.",
		Long:    "Collect pow submission data to database.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			network, _ := cmd.Flags().GetString("network")

			options, err := optionsfactory.ProcessCollectPowSubmissionOptions(options, network)
			if err != nil {
				return err
			}
			return runCollectionPowSubmission(cmd, options)
		},
	}

	optionsfactory.AddCollectPowSubmissionCliFlags(powSignalCmd, &options)

	return powSignalCmd
}

func runCollectionPowSubmission(cmd *cobra.Command, options options.CollectPowSubmissionOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	err = collectpowsubmission.StartCollectPowSubmission(commandCtx.Ctx, web3SDK, options.PGConnectionString, commandCtx.Cm)
	if err != nil {
		return err
	}
	return nil
}
