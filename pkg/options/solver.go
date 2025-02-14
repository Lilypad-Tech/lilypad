package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/solver"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	options := solver.SolverOptions{
		Server:    GetDefaultServerOptions(),
		Store:     GetDefaultStoreOptions(),
		Web3:      GetDefaultWeb3Options(),
		Services:  GetDefaultServicesOptions(),
		Telemetry: GetDefaultTelemetryOptions(),
		Metrics:   GetDefaultMetricsOptions(),
		// Timeout in seconds with a default of ten minutes
		JobOfferTimeout: GetDefaultServeOptionInt("JOB_OFFER_TIMEOUT", 600),
	}
	options.Web3.Service = system.SolverService
	return options
}

func AddSolverCliFlags(cmd *cobra.Command, options *solver.SolverOptions) {
	AddServerCliFlags(cmd, &options.Server)
	AddStoreCliFlags(cmd, &options.Store)
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServicesCliFlags(cmd, &options.Services)
	AddTelemetryCliFlags(cmd, &options.Telemetry)
	AddMetricsCliFlags(cmd, &options.Metrics)
	cmd.PersistentFlags().IntVar(
		&options.JobOfferTimeout, "job-offer-timeout", options.JobOfferTimeout,
		`The global timeout for job offers in seconds (JOB_OFFER_TIMEOUT).`,
	)
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
	err = CheckTelemetryOptions(options.Telemetry)
	if err != nil {
		return err
	}
	err = CheckMetricsOptions(options.Metrics)
	if err != nil {
		return err
	}
	if options.JobOfferTimeout <= 0 {
		return fmt.Errorf("JOB_OFFER_TIMEOUT must be greater than zero")
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
