package lilypad

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

/*
command line processing
*/
func getCommandLineExecutable() string {
	return os.Args[0]
}

func getDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func getDefaultServeOptionInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

/*
Telemetry
*/
func configureTelemetry(ctx context.Context,
	service system.Service,
	network string,
	options system.TelemetryOptions,
	metricsOptions *system.MetricsOptions,
	web3Options web3.Web3Options,
) (*system.Telemetry, error) {
	privateKey, err := web3.ParsePrivateKey(web3Options.PrivateKey)
	if err != nil {
		return nil, err
	}
	address := web3.GetAddress(privateKey)

	tc := system.TelemetryConfig{
		TelemetryURL:   options.URL,
		TelemetryToken: options.Token,
		Enabled:        !options.Disable,
		Service:        service,
		Network:        network,
		Address:        address.String(),
		GPU:            system.GetGPUInfo(),
	}

	var mc system.MetricsConfig
	if metricsOptions != nil {
		mc = system.MetricsConfig{
			MetricsURL:   metricsOptions.URL,
			MetricsToken: metricsOptions.Token,
			Enabled:      metricsOptions.Enable,
		}
	} else {
		mc = system.MetricsConfig{
			MetricsURL:   "",
			MetricsToken: "",
			Enabled:      false,
		}
	}

	telemetry, err := system.SetupOTelSDK(ctx, tc, mc)

	return &telemetry, err
}

/*
useful tools
*/
func FatalErrorHandler(cmd *cobra.Command, msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		cmd.Print(msg)
	}
	os.Exit(code)
}
