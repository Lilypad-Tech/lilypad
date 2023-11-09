package options

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

func GetDefaultServicesOptions() data.ServiceConfig {
	return data.ServiceConfig{
		Solver:   GetDefaultServeOptionString("SERVICE_SOLVER", "0x16dfb253f7d1510498c849e7f3ce7a01384f2c2a"),
		Mediator: GetDefaultServeOptionStringArray("SERVICE_MEDIATORS", []string{"0x0e0f5b1402c15a2439eeb5b106b99b1523e5cdbc"}),
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
}

func ProcessServicesOptions(options data.ServiceConfig) (data.ServiceConfig, error) {
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
