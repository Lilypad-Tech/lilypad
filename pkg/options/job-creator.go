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

func AddJobCreatorCliFlags(cmd *cobra.Command, options *jobcreator.JobCreatorOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
	AddJobCreatorOfferCliFlags(cmd, &options.Offer)
}

func CheckJobCreatorOptions(options jobcreator.JobCreatorOptions) error {
	err := CheckWeb3Options(options.Web3, true)
	if err != nil {
		return err
	}
	err = CheckModuleOptions(options.Offer.Module)
	if err != nil {
		return err
	}
	return nil
}

func ProcessJobCreatorOptions(options jobcreator.JobCreatorOptions, args []string) (jobcreator.JobCreatorOptions, error) {

	name := ""
	version := ""
	if len(args) >= 2 {
		version = args[1]
	} else if len(args) == 1 {
		name = args[0]
	}

	if name != "" {
		options.Offer.Module.Name = name
	}

	if version != "" {
		options.Offer.Module.Version = version
	}

	moduleOptions, err := ProcessModuleOptions(options.Offer.Module)
	if err != nil {
		return options, err
	}
	options.Offer.Module = moduleOptions
	newWeb3Options, err := ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckJobCreatorOptions(options)
}
