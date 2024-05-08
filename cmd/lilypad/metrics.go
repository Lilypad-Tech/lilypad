package lilypad

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/lilymetrics"
	"github.com/spf13/cobra"
)

func newMetricsCmd() *cobra.Command {

	envCmd := &cobra.Command{
		Use:     "metrics",
		Short:   "Launch Metrics.",
		Long:    "Launch Metrics.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			// options, err := optionsfactory.ProcessSolverOptions(options)
			// if err != nil {
			// 	return err
			// }
			// return runSolver(cmd, options)
			// fmt.Println("updating env")
			// metrics.runMetrics()
			// runMetrics()
			// runMediator()
			if len(args) > 0 {
				fmt.Println("First argument:", args[0])
				fmt.Println(args[0])
				if args[0] == "migrate" {
					// lilymetrics.MigrateUp()
				}
			} else {
				lilymetrics.RunMetrics()
			}

			lilymetrics.RunMetrics()
			// message := greetings.Hello("Gladys")
			return nil
		},
	}

	// optionsfactory.AddSolverCliFlags(solverCmd, &options)

	return envCmd
}
