package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

func GetDefaultServicesOptions() data.ServiceConfig {
	return data.ServiceConfig{
		Solver:   GetDefaultServeOptionString("SERVICE_SOLVER", ""),
		Mediator: GetDefaultServeOptionStringArray("SERVICE_MEDIATORS", []string{}),
		APIHost:  GetDefaultServeOptionString("API_HOST", ""),
	}
}

func AddServicesCliFlags(cmd *cobra.Command, servicesConfig *data.ServiceConfig) {
	cmd.PersistentFlags().StringVar(
		&servicesConfig.Solver, "service-solver", servicesConfig.Solver,
		`The solver to connect to (SERVICE_SOLVER)`,
	)
	cmd.PersistentFlags().StringSliceVar(
		&servicesConfig.Mediator, "service-mediators", servicesConfig.Mediator,
		`The mediators we trust (SERVICE_MEDIATORS)`,
	)
	cmd.PersistentFlags().StringVar(
		&servicesConfig.APIHost, "api-host", servicesConfig.APIHost,
		`The api host to connect to (API_HOST)`,
	)
}

func ProcessServicesOptions(options data.ServiceConfig, network string) (data.ServiceConfig, error) {
	config, err := getConfig(network)
	if err != nil {
		return options, err
	}

	// Apply configs when environment variables or command line options are not used
	if options.Solver == "" {
		options.Solver = config.ServiceConfig.Solver
	}
	if len(options.Mediator) == 0 {
		options.Mediator = config.ServiceConfig.Mediator
	}
	if options.APIHost == "" {
		options.APIHost = config.ServiceConfig.APIHost
	}

	return options, nil
}

func CheckServicesOptions(options data.ServiceConfig) error {
	if options.Solver == "" {
		return fmt.Errorf("No solver service specified - please use SERVICE_SOLVER or --service-solver")
	}
	if len(options.Mediator) == 0 {
		return fmt.Errorf("No mediators services specified - please use SERVICE_MEDIATORS or --service-mediators")
	}
	return nil

}
