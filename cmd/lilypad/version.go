package lilypad

import (
	"fmt"

	"github.com/spf13/cobra"

	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/system"
)

const goBinaryURL = "https://github.com/lilypad-tech/lilypad/releases/"

func newVersionCmd() *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Get the lilypad version",
		Long:    "Get the lilypad version",
		Example: "lilypad version",
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

	if system.Version == "" {
		fmt.Printf("version not found: download the latest binary from %s", goBinaryURL)
		// unnecessary help shows up when returned as error, so shortciruting here
		return nil
	}

	fmt.Printf("Lilypad: %s\n", system.Version)
	fmt.Printf("Commit: %s\n", system.CommitSHA)

	// TODO: suggest updating to the latest version if the current version is not the latest version

	return nil
}

//func init() {
//	if VERSION == "" {
//		VERSION = "v2" //TODO: @release, FIX: L41
//	}
//}
