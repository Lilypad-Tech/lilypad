package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func GetDefaultTelemetryOptions() system.TelemetryOptions {
	return system.TelemetryOptions{
		URL:     GetDefaultServeOptionString("TELEMETRY_URL", ""),
		Token:   GetDefaultServeOptionString("TELEMETRY_TOKEN", ""),
		Disable: GetDefaultServeOptionBool("DISABLE_TELEMETRY", false),
	}
}

func GetDefaultMetricsOptions() system.MetricsOptions {
	return system.MetricsOptions{
		URL:    GetDefaultServeOptionString("METRICS_URL", ""),
		Token:  GetDefaultServeOptionString("METRICS_TOKEN", ""),
		Enable: GetDefaultServeOptionBool("ENABLE_METRICS", false),
	}
}

func AddTelemetryCliFlags(cmd *cobra.Command, telemetryOptions *system.TelemetryOptions) {
	cmd.PersistentFlags().StringVar(
		&telemetryOptions.URL, "telemetry-url", telemetryOptions.URL,
		`The telemetry endpoint to connect to (TELEMETRY_URL)`,
	)
	cmd.PersistentFlags().StringVar(
		&telemetryOptions.Token, "telemetry-token", telemetryOptions.Token,
		`The token to auth with telemetry service (TELEMETRY_TOKEN)`,
	)
	cmd.PersistentFlags().BoolVar(
		&telemetryOptions.Disable, "disable-telemetry", telemetryOptions.Disable,
		`Disable telemetry (DISABLE_TELEMETRY)`,
	)
}

func AddMetricsCliFlags(cmd *cobra.Command, metricsOptions *system.MetricsOptions) {
	cmd.PersistentFlags().StringVar(
		&metricsOptions.URL, "metrics-url", metricsOptions.URL,
		`The metrics endpoint to connect to (METRICS_URL)`,
	)
	cmd.PersistentFlags().StringVar(
		&metricsOptions.Token, "metrics-token", metricsOptions.Token,
		`The token to auth with the metrics service (METRICS_TOKEN)`,
	)
	cmd.PersistentFlags().BoolVar(
		&metricsOptions.Enable, "enable-metrics", metricsOptions.Enable,
		`Enable metrics (ENABLE_METRICS)`,
	)
}

func ProcessTelemetryOptions(options system.TelemetryOptions, network string) (system.TelemetryOptions, error) {
	config, err := getConfig(network)
	if err != nil {
		log.Error().Msgf("failed to load config for network %s: ", err)
		return options, err
	}

	// Apply configs when environment variables or command line options are not used
	if options.URL == "" {
		options.URL = config.TelemetryOptions.URL
	}
	if options.Token == "" {
		options.Token = config.TelemetryOptions.Token
	}

	return options, nil
}

func CheckTelemetryOptions(options system.TelemetryOptions) error {
	if len(options.URL) == 0 {
		return fmt.Errorf("No telemetry endpoint specified - please use TELEMETRY_URL or --telemetry-url")
	}
	if len(options.Token) == 0 {
		return fmt.Errorf("No telemetry token specified - please use TELEMETRY_TOKEN or --telemetry-token")
	}
	return nil
}

func CheckMetricsOptions(options system.MetricsOptions) error {
	if options.Enable {
		if len(options.URL) == 0 {
			return fmt.Errorf("No metrics endpoint specified - please use METRICS_URL or --metrics-url")
		}
	}

	return nil
}
