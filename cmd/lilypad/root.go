package lilypad

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var Fatal = FatalErrorHandler

func init() { //nolint:gochecknoinits
	NewRootCmd()
}

func NewRootCmd() *cobra.Command {
	RootCmd := &cobra.Command{
		Use:   getCommandLineExecutable(),
		Short: "Lilypad",
		Long:  `Lilypad`,
	}
	RootCmd.AddCommand(newSolverCmd())
	RootCmd.AddCommand(newResourceProviderCmd())
	RootCmd.AddCommand(newRunCmd())
	RootCmd.AddCommand(newMediatorCmd())
	RootCmd.AddCommand(newJobCreatorCmd())
	RootCmd.AddCommand(newVersionCmd())
	return RootCmd
}

func Execute() {
	RootCmd := NewRootCmd()
	RootCmd.SetContext(context.Background())
	RootCmd.SetOutput(os.Stdout)
	if err := RootCmd.Execute(); err != nil {
		Fatal(RootCmd, err.Error(), 1)
	}
}
