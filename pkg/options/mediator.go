package options

import (
	"github.com/bacalhau-project/lilypad/pkg/mediator"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewMediatorOptions() mediator.MediatorOptions {
	options := mediator.MediatorOptions{
		Web3: GetDefaultWeb3Options(),
	}
	options.Web3.Service = system.MediatorService
	return options
}

func AddMediatorCliFlags(cmd *cobra.Command, options *mediator.MediatorOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
}

func CheckMediatorOptions(options mediator.MediatorOptions) error {
	err := CheckWeb3Options(options.Web3)
	if err != nil {
		return err
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
