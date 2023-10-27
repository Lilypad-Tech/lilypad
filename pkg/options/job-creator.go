package options

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewJobCreatorOptions() jobcreator.JobCreatorOptions {
	options := jobcreator.JobCreatorOptions{
		Offer:     GetDefaultJobCreatorOfferOptions(),
		Web3:      GetDefaultWeb3Options(),
		Mediation: GetDefaultJobCreatorMediationOptions(),
	}
	options.Web3.Service = system.JobCreatorService
	return options
}

func GetDefaultJobCreatorMediationOptions() jobcreator.JobCreatorMediationOptions {
	return jobcreator.JobCreatorMediationOptions{
		CheckResultsPercentage: GetDefaultServeOptionInt("MEDIATION_CHANCE", 0),
	}
}

func GetDefaultJobCreatorOfferOptions() jobcreator.JobCreatorOfferOptions {
	return jobcreator.JobCreatorOfferOptions{
		Module: GetDefaultModuleOptions(),
		// this is the default pricing mode for an JC
		Mode:     GetDefaultPricingMode(data.MarketPrice),
		Pricing:  GetDefaultPricingOptions(),
		Timeouts: GetDefaultTimeoutOptions(),
		Inputs:   map[string]string{},
		Services: GetDefaultServicesOptions(),
	}
}

func AddJobCreatorMediationCliFlags(cmd *cobra.Command, mediationOptions *jobcreator.JobCreatorMediationOptions) {
	cmd.PersistentFlags().IntVar(
		&mediationOptions.CheckResultsPercentage,
		"mediation-chance",
		mediationOptions.CheckResultsPercentage,
		"The percentage chance we will check results",
	)
}

func AddJobCreatorOfferCliFlags(cmd *cobra.Command, offerOptions *jobcreator.JobCreatorOfferOptions) {
	// add the inputs that we will merge into the module template file
	cmd.PersistentFlags().StringToStringVarP(&offerOptions.Inputs, "input", "i", offerOptions.Inputs, "Input key-value pairs")

	AddPricingModeCliFlags(cmd, &offerOptions.Mode)
	AddPricingCliFlags(cmd, &offerOptions.Pricing)
	AddTimeoutCliFlags(cmd, &offerOptions.Timeouts)
	AddModuleCliFlags(cmd, &offerOptions.Module)
	AddServicesCliFlags(cmd, &offerOptions.Services)
}

func AddJobCreatorCliFlags(cmd *cobra.Command, options *jobcreator.JobCreatorOptions) {
	AddJobCreatorMediationCliFlags(cmd, &options.Mediation)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddJobCreatorOfferCliFlags(cmd, &options.Offer)
}

func CheckJobCreatorOptions(options jobcreator.JobCreatorOptions) error {
	err := CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckModuleOptions(options.Offer.Module)
	if err != nil {
		return err
	}
	err = CheckServicesOptions(options.Offer.Services)
	if err != nil {
		return err
	}

	if options.Mediation.CheckResultsPercentage < 0 || options.Mediation.CheckResultsPercentage > 100 {
		return fmt.Errorf("mediation-chance must be between 0 and 100")
	}

	return nil
}

func ProcessJobCreatorOptions(options jobcreator.JobCreatorOptions, args []string) (jobcreator.JobCreatorOptions, error) {
	name := ""
	if len(args) == 1 {
		name = args[0]
	}

	if name != "" {
		options.Offer.Module.Name = name
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

func ProcessOnChainJobCreatorOptions(options jobcreator.JobCreatorOptions, args []string) (jobcreator.JobCreatorOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options

	err = CheckWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	err = CheckServicesOptions(options.Offer.Services)
	if err != nil {
		return options, err
	}

	options.Mediation.CheckResultsPercentage = 0

	return options, nil
}
