package options

import (
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	options := solver.SolverOptions{
		Server: GetDefaultServerOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
	options.Web3.Service = system.SolverService
	return options
}

func AddSolverCliFlags(cmd *cobra.Command, options *solver.SolverOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServerCliFlags(cmd, &options.Server)
}

func CheckSolverOptions(options solver.SolverOptions) error {
	err := CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckServerOptions(options.Server)
	if err != nil {
		return err
	}
	return nil
}

func ProcessSolverOptions(options solver.SolverOptions) (solver.SolverOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckSolverOptions(options)
}
