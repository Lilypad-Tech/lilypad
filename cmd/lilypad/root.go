package lilypad

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var Fatal = FatalErrorHandler

//FIXME: why @Kai?
//func init() { //nolint:gochecknoinits
//	NewRootCmd()
//}

func NewRootCmd() *cobra.Command {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	RootCmd := &cobra.Command{
		Use:   getCommandLineExecutable(),
		Short: "Lilypad",
		Long:  fmt.Sprintf("Lilypad: %s \nCommit: %s \n", VERSION, COMMIT_SHA),
	}
	RootCmd.AddCommand(newSolverCmd())
	RootCmd.AddCommand(newResourceProviderCmd())
	RootCmd.AddCommand(newRunCmd())
	RootCmd.AddCommand(newMediatorCmd())
	RootCmd.AddCommand(newJobCreatorCmd())
	RootCmd.AddCommand(newVersionCmd())
	RootCmd.AddCommand(newEnvCmd())
	RootCmd.AddCommand(newMetricsCmd())

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
