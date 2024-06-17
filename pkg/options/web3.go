package options

import (
	"fmt"
	"os"

	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func GetDefaultWeb3Options() web3.Web3Options {
	return web3.Web3Options{

		// core settings
		RpcURL:     GetDefaultServeOptionString("WEB3_RPC_URL", ""),
		PrivateKey: GetDefaultServeOptionString("WEB3_PRIVATE_KEY", ""),
		ChainID:    GetDefaultServeOptionInt("WEB3_CHAIN_ID", 0), //nolint:gomnd

		// contract addresses
		ControllerAddress: GetDefaultServeOptionString("WEB3_CONTROLLER_ADDRESS", ""),
		PaymentsAddress:   GetDefaultServeOptionString("WEB3_PAYMENTS_ADDRESS", ""),
		StorageAddress:    GetDefaultServeOptionString("WEB3_STORAGE_ADDRESS", ""),
		UsersAddress:      GetDefaultServeOptionString("WEB3_USERS_ADDRESS", ""),
		TokenAddress:      GetDefaultServeOptionString("WEB3_TOKEN_ADDRESS", ""),
		MediationAddress:  GetDefaultServeOptionString("WEB3_MEDIATION_ADDRESS", ""),
		JobCreatorAddress: GetDefaultServeOptionString("WEB3_JOBCREATOR_ADDRESS", ""),
		PowAddress:        GetDefaultServeOptionString("WEB3_POW_ADDRESS", ""),

		// misc
		Service: system.DefaultService,
	}
}

func AddWeb3CliFlags(cmd *cobra.Command, web3Options *web3.Web3Options) {
	cmd.PersistentFlags().StringVar(
		&web3Options.RpcURL, "web3-rpc-url", web3Options.RpcURL,
		`The URL of the web3 RPC server (WEB3_RPC_URL).`,
	)

	// don't use the env as the default here because otherwise it will show when --help is used
	// instead we inject the env value into the options after boot if needed
	cmd.PersistentFlags().StringVar(
		&web3Options.PrivateKey, "web3-private-key", "",
		`The private key to use for signing web3 transactions (WEB3_PRIVATE_KEY).`,
	)
	cmd.PersistentFlags().IntVar(
		&web3Options.ChainID, "web3-chain-id", web3Options.ChainID,
		`The chain id for the web3 RPC server (WEB3_CHAIN_ID).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.ControllerAddress, "web3-controller-address", web3Options.ControllerAddress,
		`The address of the controller contract (WEB3_CONTROLLER_ADDRESS).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.PaymentsAddress, "web3-payments-address", web3Options.PaymentsAddress,
		`The address of the payments contract (WEB3_PAYMENTS_ADDRESS).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.StorageAddress, "web3-storage-address", web3Options.StorageAddress,
		`The address of the storage contract (WEB3_STORAGE_ADDRESS).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.UsersAddress, "web3-users-address", web3Options.UsersAddress,
		`The address of the users contract (WEB3_USERS_ADDRESS).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.TokenAddress, "web3-token-address", web3Options.TokenAddress,
		`The address of the token contract (WEB3_TOKEN_ADDRESS).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.PowAddress, "web3-pow-address", web3Options.PowAddress,
		`The address of the pow contract (WEB3_POW_ADDRESS).`,
	)
}

func CheckWeb3Options(options web3.Web3Options) error {
	// core settings
	if options.RpcURL == "" {
		return fmt.Errorf("WEB3_RPC_URL is required")
	}
	if options.PrivateKey == "" {
		return fmt.Errorf("WEB3_PRIVATE_KEY is required")
	}

	// this is the only address we actually need
	// we can load the rest of the addresses from the controller address if needed
	if options.ControllerAddress == "" {
		return fmt.Errorf("WEB3_CONTROLLER_ADDRESS is required")
	}

	return nil
}

func ProcessWeb3Options(options web3.Web3Options, network string) (web3.Web3Options, error) {
	config, err := getConfig(network)
	if err != nil {
		return options, err
	}

	// Apply configs when environment variables or command line options are not used
	if options.RpcURL == "" {
		options.RpcURL = config.Web3.RpcURL
	}
	if options.ChainID == 0 {
		options.ChainID = config.Web3.ChainID
	}
	if options.ControllerAddress == "" {
		options.ControllerAddress = config.Web3.ControllerAddress
	}
	if options.PaymentsAddress == "" {
		options.PaymentsAddress = config.Web3.PaymentsAddress
	}
	if options.StorageAddress == "" {
		options.StorageAddress = config.Web3.StorageAddress
	}
	if options.UsersAddress == "" {
		options.UsersAddress = config.Web3.UsersAddress
	}
	if options.MediationAddress == "" {
		options.MediationAddress = config.Web3.MediationAddress
	}
	if options.JobCreatorAddress == "" {
		options.JobCreatorAddress = config.Web3.JobCreatorAddress
	}
	if options.TokenAddress == "" {
		options.TokenAddress = config.Web3.TokenAddress
	}

	if options.PrivateKey == "" {
		options.PrivateKey = os.Getenv("WEB3_PRIVATE_KEY")
	}
	return options, nil
}
