package lilypad

import (
	"context"
	"fmt"
	"os"

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
		Long:  fmt.Sprintf("Lilypad: %s \nCommit: %s \n", version, commitSHA),
	}

	var network string
	RootCmd.PersistentFlags().StringVarP(&network, "network", "n", "testnet", "Sets a target network configuration")

	var powSignalCmd bool
	RootCmd.PersistentFlags().BoolVarP(&powSignalCmd, "pow-signal", "", false, "Send a pow signal to smart contract")

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
