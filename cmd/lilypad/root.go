package lilypad

import (
	"context"
	"fmt"
	"os"

	"github.com/lilypad-tech/lilypad/v2/pkg/system"
	"github.com/spf13/cobra"
)

var Fatal = FatalErrorHandler

//FIXME: why @Kai?
//func init() { //nolint:gochecknoinits
//	NewRootCmd()
//}

func NewRootCmd() *cobra.Command {
	RootCmd := &cobra.Command{
		Use:   getCommandLineExecutable(),
		Short: "Lilypad",
		Long:  fmt.Sprintf("Lilypad: %s \nCommit: %s \n", system.Version, system.CommitSHA),
	}

	var network string
	var lilynext bool
	RootCmd.PersistentFlags().StringVarP(&network, "network", "n", "testnet", "Sets a target network configuration")
	RootCmd.PersistentFlags().BoolVar(&lilynext, "lilynext", false, "Use the new Lilypad protocol")

	RootCmd.AddCommand(newSolverCmd())
	RootCmd.AddCommand(newResourceProviderCmd())
	RootCmd.AddCommand(newPowSignalCmd())
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
