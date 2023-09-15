package options

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/mediator"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewMediatorOptions() mediator.MediatorOptions {
	options := mediator.MediatorOptions{
		Bacalhau: GetDefaultBacalhauOptions(),
		Web3:     GetDefaultWeb3Options(),
		Services: GetDefaultServicesOptions(),
	}
	options.Web3.Service = system.MediatorService
	return options
}

func AddMediatorCliFlags(cmd *cobra.Command, options *mediator.MediatorOptions) {
	AddBacalhauCliFlags(cmd, &options.Bacalhau)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServicesCliFlags(cmd, &options.Services)
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
	// only check the solver because we are the mediator
	if options.Services.Solver == "" {
		return fmt.Errorf("No solver service specified - please use SERVICE_SOLVER or --service-solver")
	}
	return nil
}

func ProcessMediatorOptions(options mediator.MediatorOptions) (mediator.MediatorOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckMediatorOptions(options)
}
