package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

func GetDefaultServicesOptions() data.ServiceConfig {
	return data.ServiceConfig{
		Solver:   GetDefaultServeOptionString("SERVICE_SOLVER", "0x346d811cbb883548252418121f5bb0371eb07049"),
		Mediator: GetDefaultServeOptionStringArray("SERVICE_MEDIATORS", []string{"0xc66b9b74e307f30e7af79c03fee6ceb8b1ced997"}),
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
