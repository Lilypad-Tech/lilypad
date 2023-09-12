package options

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

func GetDefaultTrustedPartyOptions() data.TrustedParties {
	return data.TrustedParties{
		Mediator: GetDefaultServeOptionStringArray("TRUSTED_MEDIATORS", []string{}),
		// the shortcut version
		Directory: GetDefaultServeOptionStringArray("TRUSTED_DIRECTORIES", []string{}),
	}
}

func AddTrustedPartyCliFlags(cmd *cobra.Command, trustedPartyConfig *data.TrustedParties) {
	cmd.PersistentFlags().StringSliceVar(
		&trustedPartyConfig.Mediator, "trusted-mediators", trustedPartyConfig.Mediator,
		`The mediators we trust (TRUSTED_MEDIATORS)`,
	)
	cmd.PersistentFlags().StringSliceVar(
		&trustedPartyConfig.Directory, "trusted-directories", trustedPartyConfig.Directory,
		`The directory services we trust (TRUSTED_DIRECTORIES)`,
	)
}

// see if we have a shortcut and fill in the other values if we do
func ProcessTrustedPartyOptions(options data.TrustedParties) (data.TrustedParties, error) {
	return options, nil
}

func CheckTrustedPartyOptions(options data.TrustedParties) error {
	if len(options.Mediator) == 0 {
		return fmt.Errorf("No trusted mediators specified - please use TRUSTED_MEDIATORS or --trusted-mediators")
	}
	if len(options.Directory) == 0 {
		return fmt.Errorf("No trusted directories specified - please use TRUSTED_DIRECTORIES or --trusted-directories")
	}
	return nil
}
