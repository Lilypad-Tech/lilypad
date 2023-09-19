package lilypad

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()

	runCmd := &cobra.Command{
		Use:     "run",
		Short:   "Run a job on the Lilypad network.",
		Long:    "Run a job on the Lilypad network.",
		Example: "run cowsay v0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := optionsfactory.ProcessJobCreatorOptions(options, args)
			if err != nil {
				return err
			}
			return runJob(cmd, options)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(runCmd, &options)

	return runCmd
}

func runJob(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()
	result, err := jobcreator.RunJob(commandCtx, options)
	fmt.Printf("result --------------------------------------\n")
	spew.Dump(result)
	return err
}
