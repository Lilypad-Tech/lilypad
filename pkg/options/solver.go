package options

import (
	"github.com/lilypad-tech/lilypad/pkg/solver"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	options := solver.SolverOptions{
		Server:    GetDefaultServerOptions(),
		Web3:      GetDefaultWeb3Options(),
		Services:  GetDefaultServicesOptions(),
		Telemetry: GetDefaultTelemetryOptions(),
		Metrics:   GetDefaultMetricsOptions(),
	}
	options.Web3.Service = system.SolverService
	return options
}

func AddSolverCliFlags(cmd *cobra.Command, options *solver.SolverOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServerCliFlags(cmd, &options.Server)
	AddServicesCliFlags(cmd, &options.Services)
	AddTelemetryCliFlags(cmd, &options.Telemetry)
	AddMetricsCliFlags(cmd, &options.Metrics)
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
	err = CheckTelemetryOptions(options.Telemetry)
	if err != nil {
		return err
	}
	err = CheckMetricsOptions(options.Metrics)
	if err != nil {
		return err
	}
	return nil
}

func ProcessSolverOptions(options solver.SolverOptions, network string) (solver.SolverOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	newTelemetryOptions, err := ProcessTelemetryOptions(options.Telemetry, network)
	if err != nil {
		return options, err
	}
	options.Telemetry = newTelemetryOptions
	return options, CheckSolverOptions(options)
}
