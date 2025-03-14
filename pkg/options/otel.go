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

func GetDefaultLogsOptions() system.LogsOptions {
	return system.LogsOptions{
		Enable: GetDefaultServeOptionBool("ENABLE_OTEL_LOGS", false),
		URL:    GetDefaultServeOptionString("LOGS_URL", ""),
		Token:  GetDefaultServeOptionString("LOGS_TOKEN", ""),
		// Default values from: https://pkg.go.dev/go.opentelemetry.io/otel/sdk/log@v0.10.0#BatchProcessorOption
		ExportBufferSize:   GetDefaultServeOptionInt("LOGS_EXPORT_BUFFER_SIZE", 1),
		ExportInterval:     GetDefaultServeOptionInt("LOGS_EXPORT_INTERVAL_SECONDS", 1),
		ExportMaxBatchSize: GetDefaultServeOptionInt("LOGS_EXPORT_MAX_BATCH_SIZE", 512),
		ExportTimeout:      GetDefaultServeOptionInt("LOGS_EXPORT_TIMEOUT_SECONDS", 30),
		MaxQueueSize:       GetDefaultServeOptionInt("LOGS_MAX_QUEUE_SIZE", 2048),
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

func AddLogsCliFlags(cmd *cobra.Command, logsOptions *system.LogsOptions) {
	cmd.PersistentFlags().BoolVar(
		&logsOptions.Enable, "enable-otel-logs", logsOptions.Enable,
		`Enable OpenTelemetry logs (ENABLE_OTEL_LOGS)`,
	)
	cmd.PersistentFlags().StringVar(
		&logsOptions.URL, "logs-url", logsOptions.URL,
		`The logs endpoint to connect to (LOGS_URL)`,
	)
	cmd.PersistentFlags().StringVar(
		&logsOptions.Token, "logs-token", logsOptions.Token,
		`The token to auth with the logs service (LOGS_TOKEN)`,
	)
	cmd.PersistentFlags().IntVar(
		&logsOptions.ExportBufferSize, "logs-export-buffer-size", logsOptions.ExportBufferSize,
		`The log batches stored by the exporter (LOGS_EXPORT_BUFFER_SIZE)`,
	)
	cmd.PersistentFlags().IntVar(
		&logsOptions.ExportInterval, "logs-export-interval-seconds", logsOptions.ExportInterval,
		`The logs export interval (LOGS_EXPORT_INTERVAL_SECONDS)`,
	)
	cmd.PersistentFlags().IntVar(
		&logsOptions.ExportMaxBatchSize, "logs-export-max-batch-size", logsOptions.ExportMaxBatchSize,
		`The logs export max records batch size (LOGS_EXPORT_MAX_BATCH_SIZE)`,
	)
	cmd.PersistentFlags().IntVar(
		&logsOptions.ExportTimeout, "logs-export-timeout-seconds", logsOptions.ExportTimeout,
		`The logs export timeout (LOGS_EXPORT_TIMEOUT_SECONDS)`,
	)
	cmd.PersistentFlags().IntVar(
		&logsOptions.MaxQueueSize, "logs-max-queue-size", logsOptions.MaxQueueSize,
		`The logs exporter max records queue size (LOGS_MAX_QUEUE_SIZE)`,
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
	if options.Enable && len(options.URL) == 0 {
		return fmt.Errorf("No metrics endpoint specified - please use METRICS_URL or --metrics-url")
	}

	return nil
}

func CheckLogsOptions(options system.LogsOptions) error {
	if options.Enable && len(options.URL) == 0 {
		return fmt.Errorf("No logging endpoint specified - please use LOGS_URL or --logs-url")
	}

	return nil
}
