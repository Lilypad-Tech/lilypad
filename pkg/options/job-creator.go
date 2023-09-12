package options

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	"github.com/spf13/cobra"
)

func NewJobCreatorOptions() jobcreator.JobCreatorOptions {
	return jobcreator.JobCreatorOptions{
		Offer: GetDefaultJobCreatorOfferOptions(),
		Web3:  GetDefaultWeb3Options(),
	}
}

func GetDefaultJobCreatorOfferOptions() jobcreator.JobCreatorOfferOptions {
	return jobcreator.JobCreatorOfferOptions{
		Module: GetDefaultModuleOptions(),
		// this is the default pricing mode for an JC
		Mode:           GetDefaultPricingMode(data.MarketPrice),
		Pricing:        GetDefaultPricingOptions(),
		Timeouts:       GetDefaultTimeoutOptions(),
		Inputs:         map[string]string{},
		TrustedParties: GetDefaultTrustedPartyOptions(),
	}
}

func AddJobCreatorOfferCliFlags(cmd *cobra.Command, offerOptions *jobcreator.JobCreatorOfferOptions) {
	// add the inputs that we will merge into the module template file
	cmd.PersistentFlags().StringToStringVarP(&offerOptions.Inputs, "input", "i", offerOptions.Inputs, "Input key-value pairs")

	AddPricingModeCliFlags(cmd, &offerOptions.Mode)
	AddPricingCliFlags(cmd, &offerOptions.Pricing)
	AddTimeoutCliFlags(cmd, &offerOptions.Timeouts)
	AddModuleCliFlags(cmd, &offerOptions.Module)
	AddTrustedPartyCliFlags(cmd, &offerOptions.TrustedParties)
}

func ProcessJobCreatorOfferOptions(options jobcreator.JobCreatorOfferOptions) (jobcreator.JobCreatorOfferOptions, error) {
	// parse the module flags
	moduleOptions, err := ProcessModuleOptions(options.Module)
	if err != nil {
		return options, err
	}
	options.Module = moduleOptions
	return options, nil
}

func CheckJobCreatorOfferOptions(options jobcreator.JobCreatorOfferOptions) error {
	err := CheckModuleOptions(options.Module)
	if err != nil {
		return err
	}
	return nil
}
