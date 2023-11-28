package lilypad

import (
	"fmt"

	"github.com/spf13/cobra"

	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

// VERSION use `go build -ldflags="-X lilypad.VERSION=vX.Y.Z" `
var VERSION string

// COMMIT_SHA use `go build -ldflags="-X lilypad.COMMIT_SHA=$(git rev-parse HEAD)" `
var COMMIT_SHA string

const GO_BINARY_URL = "https://github.com/bacalhau-project/lilypad/releases/"

func newVersionCmd() *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	versionCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"-v"},
		Short:   "Get the lilypad version",
		Long:    "Get the lilypad version",
		Example: "lilypad version, lilypad -v",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runVersion(cmd)
		},
	}

	optionsfactory.AddSolverCliFlags(versionCmd, &options)

	return versionCmd
}

func runVersion(cmd *cobra.Command) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	if VERSION == "" {
		fmt.Printf("version not found: download the latest binary from %s", GO_BINARY_URL)
		// unnecessary help shows up when returned as error, so shortciruting here
		return nil
	}

	fmt.Printf("Lilypad: %s\n", VERSION)
	fmt.Printf("Commit: %s\n", COMMIT_SHA)

	// TODO: suggest updating to the latest version if the current version is not the latest version

	return nil
}