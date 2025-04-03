package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/v2/pkg/mediator"
	"github.com/lilypad-tech/lilypad/v2/pkg/system"
	"github.com/spf13/cobra"
)

func NewMediatorOptions() mediator.MediatorOptions {
	options := mediator.MediatorOptions{
		Bacalhau: GetDefaultBacalhauOptions(),
		Web3:     GetDefaultWeb3Options(),
		Services: GetDefaultServicesOptions(),
		IPFS:     GetDefaultIPFSOptions(),
	}
	options.Web3.Service = system.MediatorService
	return options
}

func AddMediatorCliFlags(cmd *cobra.Command, options *mediator.MediatorOptions) {
	AddBacalhauCliFlags(cmd, &options.Bacalhau)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServicesCliFlags(cmd, &options.Services)
	AddIPFSCliFlags(cmd, &options.IPFS)
}

func CheckMediatorOptions(options mediator.MediatorOptions) error {
	err := CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckBacalhauOptions(options.Bacalhau)
	if err != nil {
		return err
	}
	err = CheckIPFSOptions(options.IPFS)
	if err != nil {
		return err
	}
	// only check the solver because we are the mediator
	if options.Services.Solver == "" {
		return fmt.Errorf("No solver service specified - please use SERVICE_SOLVER or --service-solver")
	}
	return nil
}

func ProcessMediatorOptions(options mediator.MediatorOptions, network string) (mediator.MediatorOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	newServicesOptions, err := ProcessServicesOptions(options.Services, network)
	if err != nil {
		return options, err
	}
	options.Services = newServicesOptions
	newIPFSOptions, err := ProcessIPFSOptions(options.IPFS, network)
	if err != nil {
		return options, err
	}
	options.IPFS = newIPFSOptions

	return options, CheckMediatorOptions(options)
}
