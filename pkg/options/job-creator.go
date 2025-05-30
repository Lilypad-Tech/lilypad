package options

import (
	"fmt"
	"os"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/jobcreator"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/spf13/cobra"
)

func NewJobCreatorOptions() jobcreator.JobCreatorOptions {
	options := jobcreator.JobCreatorOptions{
		Offer:     GetDefaultJobCreatorOfferOptions(),
		Web3:      GetDefaultWeb3Options(),
		Mediation: GetDefaultJobCreatorMediationOptions(),
		Telemetry: GetDefaultTelemetryOptions(),
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
		Mode:       GetDefaultPricingMode(data.MarketPrice),
		Pricing:    GetDefaultPricingOptions(),
		Timeouts:   GetDefaultTimeoutOptions(),
		Inputs:     map[string]string{},
		InputsPath: GetDefaultServeOptionString("INPUTS_PATH", ""),
		Services:   GetDefaultServicesOptions(),
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
	cmd.PersistentFlags().StringVar(&offerOptions.InputsPath, "inputs-path", offerOptions.InputsPath, "Path to file inputs directory (INPUTS_PATH)")

	AddPricingModeCliFlags(cmd, &offerOptions.Mode)
	AddPricingCliFlags(cmd, &offerOptions.Pricing)
	AddTimeoutCliFlags(cmd, &offerOptions.Timeouts)
	AddModuleCliFlags(cmd, &offerOptions.Module)
	AddServicesCliFlags(cmd, &offerOptions.Services)
	AddTargetCliFlags(cmd, &offerOptions.Target)
}

func AddJobCreatorCliFlags(cmd *cobra.Command, options *jobcreator.JobCreatorOptions) {
	AddJobCreatorMediationCliFlags(cmd, &options.Mediation)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddJobCreatorOfferCliFlags(cmd, &options.Offer)
	AddTelemetryCliFlags(cmd, &options.Telemetry)
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
	err = CheckTelemetryOptions(options.Telemetry)
	if err != nil {
		return err
	}

	if options.Offer.InputsPath != "" {
		fileInfo, err := os.Stat(options.Offer.InputsPath)
		if err != nil {
			return fmt.Errorf("inputs path error: %w", err)
		}
		if !fileInfo.IsDir() {
			return fmt.Errorf("inputs path must be a directory: %s", options.Offer.InputsPath)
		}
	}

	if options.Mediation.CheckResultsPercentage < 0 || options.Mediation.CheckResultsPercentage > 100 {
		return fmt.Errorf("mediation-chance must be between 0 and 100")
	}

	return nil
}

func ProcessJobCreatorOptions(options jobcreator.JobCreatorOptions, args []string, network string) (jobcreator.JobCreatorOptions, error) {
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

	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options

	newServicesOptions, err := ProcessServicesOptions(options.Offer.Services, network)
	if err != nil {
		return options, err
	}
	options.Offer.Services = newServicesOptions

	newTargetOptions, err := ProcessTargetOptions(options.Offer.Target)
	if err != nil {
		return options, err
	}
	options.Offer.Target = newTargetOptions

	newTelemetryOptions, err := ProcessTelemetryOptions(options.Telemetry, network)
	if err != nil {
		return options, err
	}
	options.Telemetry = newTelemetryOptions

	return options, CheckJobCreatorOptions(options)
}

func ProcessOnChainJobCreatorOptions(options jobcreator.JobCreatorOptions, args []string, network string) (jobcreator.JobCreatorOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options

	newServicesOptions, err := ProcessServicesOptions(options.Offer.Services, network)
	if err != nil {
		return options, err
	}
	options.Offer.Services = newServicesOptions

	newTelemetryOptions, err := ProcessTelemetryOptions(options.Telemetry, network)
	if err != nil {
		return options, err
	}
	options.Telemetry = newTelemetryOptions

	err = CheckWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	err = CheckServicesOptions(options.Offer.Services)
	if err != nil {
		return options, err
	}
	err = CheckTelemetryOptions(options.Telemetry)
	if err != nil {
		return options, err
	}

	options.Mediation.CheckResultsPercentage = 0

	return options, nil
}
