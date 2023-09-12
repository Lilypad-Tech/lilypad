package options

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	"github.com/spf13/cobra"
)

func NewJobCreatorOptions() jobcreator.JobCreatorOptions {
	return jobcreator.JobCreatorOptions{
		Web3: GetDefaultWeb3Options(),
	}
}

func GetDefaultJobCreatorOfferOptions() jobcreator.JobCreatorOfferOptions {
	return jobcreator.JobCreatorOfferOptions{
		Module:  GetDefaultModuleOptions(),
		Pricing: GetDefaultPricingOptions(data.MarketPrice),
	}
}

func AddJobCreatorOfferCliFlags(cmd *cobra.Command, offerOptions jobcreator.JobCreatorOfferOptions) {
	AddPricingCliFlags(cmd, offerOptions.Pricing)
	AddModuleCliFlags(cmd, offerOptions.Module)

	// add the inputs that we will merge into the module template file
	cmd.PersistentFlags().StringToStringVarP(&offerOptions.Inputs, "input", "i", map[string]string{}, "Input key-value pairs")
}

func ProcessJobCreatorOfferOptions(options jobcreator.JobCreatorOfferOptions) (jobcreator.JobCreatorOfferOptions, error) {
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
