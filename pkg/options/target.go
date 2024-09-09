package options

import (
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

func AddTargetCliFlags(cmd *cobra.Command, targetConfig *data.TargetConfig) {
	cmd.PersistentFlags().StringVarP(
		&targetConfig.Address, "target", "t", targetConfig.Address,
		`The address to target (TARGET)`,
	)
}

func ProcessTargetOptions(options data.TargetConfig) (data.TargetConfig, error) {
	return options, nil
}
