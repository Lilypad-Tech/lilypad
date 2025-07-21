package options

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	options := solver.SolverOptions{
		Server:       GetDefaultServerOptions(),
		Store:        GetDefaultStoreOptions(),
		Web3:         GetDefaultWeb3Options(),
		Services:     GetDefaultServicesOptions(),
		Stats:        GetDefaultStatsOptions(),
		Telemetry:    GetDefaultTelemetryOptions(),
		Metrics:      GetDefaultMetricsOptions(),
		Logs:         GetDefaultLogsOptions(),
		Timeouts:     GetDefaultTimeoutOptions(),
		AdminService: GetDefaultAdminServiceOptions(),
	}
	options.Web3.Service = system.SolverService
	return options
}

func AddSolverCliFlags(cmd *cobra.Command, options *solver.SolverOptions) {
	AddServerCliFlags(cmd, &options.Server)
	AddStoreCliFlags(cmd, &options.Store)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServicesCliFlags(cmd, &options.Services)
	AddStatsCliFlags(cmd, &options.Stats)
	AddTelemetryCliFlags(cmd, &options.Telemetry)
	AddMetricsCliFlags(cmd, &options.Metrics)
	AddLogsCliFlags(cmd, &options.Logs)
	AddTimeoutCliFlags(cmd, &options.Timeouts)
}

func CheckSolverOptions(options solver.SolverOptions) error {
	err := CheckStoreOptions(options.Store)
	if err != nil {
		return err
	}
	err = CheckServerOptions(options.Server, options.Store.Type)
	if err != nil {
		return err
	}
	err = CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}
	err = CheckStatsOptions(options.Stats)
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
	err = CheckLogsOptions(options.Logs)
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
