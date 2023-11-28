package lilypad

import (
	"fmt"

	"github.com/spf13/cobra"

	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

const VERSION = ""

func newVersionCmd() *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	versionCmd := &cobra.Command{
		Use:     "solver",
		Short:   "Start the lilypad solver service.",
		Long:    "Start the lilypad solver service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := optionsfactory.ProcessSolverOptions(options)
			if err != nil {
				return err
			}
			return runVersion(cmd, options)
		},
	}

	optionsfactory.AddSolverCliFlags(versionCmd, &options)

	return versionCmd
}

func runVersion(cmd *cobra.Command, options solver.SolverOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	fmt.Printf("Lilypad: %s\n", VERSION)

	return nil
}
