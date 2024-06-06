package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider"
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider/powsignal"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewResourceProviderOptions() resourceprovider.ResourceProviderOptions {
	options := resourceprovider.ResourceProviderOptions{
		Bacalhau: GetDefaultBacalhauOptions(),
		Offers:   GetDefaultResourceProviderOfferOptions(),
		Web3:     GetDefaultWeb3Options(),
		Pow:      GetDefaultResourceProviderPowOptions(),
	}
	options.Web3.Service = system.ResourceProviderService
	return options
}

func GetDefaultResourceProviderPowOptions() resourceprovider.ResourceProviderPowOptions {
	return resourceprovider.ResourceProviderPowOptions{
		EnablePow: GetDefaultServeOptionBool("ENABLE_POW", false),
	}
}

func GetDefaultResourceProviderOfferOptions() resourceprovider.ResourceProviderOfferOptions {
	return resourceprovider.ResourceProviderOfferOptions{
		// by default let's offer 1 CPU, 0 GPU and 1GB RAM
		OfferSpec: data.MachineSpec{
			CPU: GetDefaultServeOptionInt("OFFER_CPU", 1000), //nolint:gomnd
			GPU: GetDefaultServeOptionInt("OFFER_GPU", 0),    //nolint:gomnd
			RAM: GetDefaultServeOptionInt("OFFER_RAM", 1024), //nolint:gomnd
		},
		OfferCount: GetDefaultServeOptionInt("OFFER_COUNT", 1), //nolint:gomnd
		// this can be populated by a config file
		Specs: []data.MachineSpec{},
		// if an RP wants to only run certain modules they list them here
		// XXX SECURITY: enforce that they are specified with specific git hashes!
		Modules: GetDefaultServeOptionStringArray("OFFER_MODULES", []string{}),
		// this is the default pricing mode for an RP
		Mode: GetDefaultPricingMode(data.FixedPrice),
		// this is the default pricing for a module unless it has a specific price
		DefaultPricing:  GetDefaultPricingOptions(),
		DefaultTimeouts: GetDefaultTimeoutOptions(),
		// allows an RP to list specific prices for each module
		ModulePricing:  map[string]data.DealPricing{},
		ModuleTimeouts: map[string]data.DealTimeouts{},
		Services:       GetDefaultServicesOptions(),
	}
}

func AddResourceProviderOfferCliFlags(cmd *cobra.Command, offerOptions *resourceprovider.ResourceProviderOfferOptions) {
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.CPU, "offer-cpu", offerOptions.OfferSpec.CPU,
		`How many milli-cpus to offer the network (OFFER_CPU).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.GPU, "offer-gpu", offerOptions.OfferSpec.GPU,
		`How many milli-gpus to offer the network (OFFER_GPU).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.RAM, "offer-ram", offerOptions.OfferSpec.RAM,
		`How many megabytes of RAM to offer the network (OFFER_RAM).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferCount, "offer-count", offerOptions.OfferCount,
		`How many machines will we offer using the cpu, ram and gpu settings (OFFER_COUNT).`,
	)
	cmd.PersistentFlags().StringArrayVar(
		&offerOptions.Modules, "offer-modules", offerOptions.Modules,
		`The modules you are willing to run (OFFER_MODULES).`,
	)
	AddPricingModeCliFlags(cmd, &offerOptions.Mode)
	AddPricingCliFlags(cmd, &offerOptions.DefaultPricing)
	AddTimeoutCliFlags(cmd, &offerOptions.DefaultTimeouts)
	AddServicesCliFlags(cmd, &offerOptions.Services)
}

func AddResourceProviderPowCliFlags(cmd *cobra.Command, options *resourceprovider.ResourceProviderPowOptions) {
	cmd.PersistentFlags().BoolVar(
		&options.EnablePow, "enable-pow", options.EnablePow,
		`Start pow mining (ENABLE_POW)`,
	)
}

func AddResourceProviderCliFlags(cmd *cobra.Command, options *resourceprovider.ResourceProviderOptions) {
	AddBacalhauCliFlags(cmd, &options.Bacalhau)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddResourceProviderOfferCliFlags(cmd, &options.Offers)
	AddResourceProviderPowCliFlags(cmd, &options.Pow)
}

func AddPowSignalCliFlags(cmd *cobra.Command, options *powsignal.PowSignalOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
}

func CheckResourceProviderOfferOptions(options resourceprovider.ResourceProviderOfferOptions) error {
	// loop over all specs and add up the total number of cpus
	totalCPU := 0
	for _, spec := range options.Specs {
		totalCPU += spec.CPU
	}

	if totalCPU <= 0 {
		return fmt.Errorf("OFFER_CPU cannot be zero")
	}

	// do the same for memory
	totalRAM := 0
	for _, spec := range options.Specs {
		totalRAM += spec.RAM
	}

	if totalRAM <= 0 {
		return fmt.Errorf("OFFER_RAM cannot be zero")
	}

	return nil
}

func CheckResourceProviderOptions(options resourceprovider.ResourceProviderOptions) error {
	err := CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckResourceProviderOfferOptions(options.Offers)
	if err != nil {
		return err
	}
	err = CheckServicesOptions(options.Offers.Services)
	if err != nil {
		return err
	}
	err = CheckBacalhauOptions(options.Bacalhau)
	if err != nil {
		return err
	}
	return nil
}

func ProcessResourceProviderOfferOptions(options resourceprovider.ResourceProviderOfferOptions, network string) (resourceprovider.ResourceProviderOfferOptions, error) {
	newServicesOptions, err := ProcessServicesOptions(options.Services, network)
	if err != nil {
		return options, err
	}
	options.Services = newServicesOptions

	// if there are no specs then populate with the single spec
	if len(options.Specs) == 0 {
		// loop the number of machines we want to offer
		for i := 0; i < options.OfferCount; i++ {
			options.Specs = append(options.Specs, options.OfferSpec)
		}
	}
	return options, nil
}

func ProcessResourceProviderOptions(options resourceprovider.ResourceProviderOptions, network string) (resourceprovider.ResourceProviderOptions, error) {
	newOfferOptions, err := ProcessResourceProviderOfferOptions(options.Offers, network)
	if err != nil {
		return options, err
	}
	options.Offers = newOfferOptions
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckResourceProviderOptions(options)
}
