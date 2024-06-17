package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

type CollectPowSubmissionOptions struct {
	Web3               web3.Web3Options
	PGConnectionString string
}

func NewCollectPowSubmissionOptions() CollectPowSubmissionOptions {
	options := CollectPowSubmissionOptions{
		Web3:               GetDefaultWeb3Options(),
		PGConnectionString: GetDefaultServeOptionString("PG_CONNECTION_STRING", ""),
	}
	return options
}

func ProcessCollectPowSubmissionOptions(options CollectPowSubmissionOptions, network string) (CollectPowSubmissionOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckCollectPowSubmissionOptions(options)
}

func CheckCollectPowSubmissionOptions(options CollectPowSubmissionOptions) error {
	if options.Web3.RpcURL == "" {
		return fmt.Errorf("WEB3_RPC_URL is required")
	}

	if options.Web3.ControllerAddress == "" {
		return fmt.Errorf("WEB3_CONTROLLER_ADDRESS is required")
	}

	if options.PGConnectionString == "" {
		return fmt.Errorf("PG_CONNECTION_STRING is required")
	}
	return nil
}

func AddCollectPowSubmissionCliFlags(cmd *cobra.Command, options *CollectPowSubmissionOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)

	cmd.PersistentFlags().StringVar(
		&options.PGConnectionString, "pg-connection-string", options.PGConnectionString,
		`The Connection string  of the PG database (PG_CONNECTION_STRING).`,
	)
}
