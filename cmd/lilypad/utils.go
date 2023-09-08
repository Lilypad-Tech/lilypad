package lilypad

import (
	"os"
	"strconv"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

/*
command line processing
*/
func getCommandLineExecutable() string {
	return os.Args[0]
}

func getDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func getDefaultServeOptionInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

/*
useful tools
*/
func FatalErrorHandler(cmd *cobra.Command, msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		cmd.Print(msg)
	}
	os.Exit(code)
}

/*
server options
*/
func getDefaultServerOptions() server.ServerOptions {
	return server.ServerOptions{
		URL:  getDefaultServeOptionString("SERVER_URL", ""),
		Host: getDefaultServeOptionString("SERVER_HOST", "0.0.0.0"),
		Port: getDefaultServeOptionInt("SERVER_PORT", 8080), //nolint:gomnd
	}
}

func addServerCliFlags(cmd *cobra.Command, serverOptions server.ServerOptions) {
	cmd.PersistentFlags().StringVar(
		&serverOptions.URL, "server-url", serverOptions.URL,
		`The URL the api server is listening on (SERVER_URL).`,
	)
	cmd.PersistentFlags().StringVar(
		&serverOptions.Host, "server-host", serverOptions.Host,
		`The host to bind the api server to (SERVER_HOST).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.Port, "server-port", serverOptions.Port,
		`The port to bind the api server to (SERVER_PORT).`,
	)
}

/*
web3 options
*/
func getDefaultWeb3Options() web3.Web3Options {
	return web3.Web3Options{
		RpcURL:            getDefaultServeOptionString("WEB3_RPC_URL", ""),
		PrivateKey:        getDefaultServeOptionString("WEB3_PRIVATE_KEY", ""),
		ChainID:           getDefaultServeOptionInt("WEB3_CHAIN_ID", 1337), //nolint:gomnd
		ControllerAddress: getDefaultServeOptionString("WEB3_CONTROLLER_ADDRESS", ""),
		PaymentsAddress:   getDefaultServeOptionString("WEB3_PAYMENTS_ADDRESS", ""),
		StorageAddress:    getDefaultServeOptionString("WEB3_STORAGE_ADDRESS", ""),
		TokenAddress:      getDefaultServeOptionString("WEB3_TOKEN_ADDRESS", ""),
	}
}

func addWeb3CliFlags(cmd *cobra.Command, web3Options web3.Web3Options) {
	cmd.PersistentFlags().StringVar(
		&web3Options.RpcURL, "web3-rpc-url", web3Options.RpcURL,
		`The URL of the web3 RPC server (WEB3_RPC_URL).`,
	)
	cmd.PersistentFlags().StringVar(
		&web3Options.PrivateKey, "web3-private-key", web3Options.PrivateKey,
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
		&web3Options.TokenAddress, "web3-token-address", web3Options.TokenAddress,
		`The address of the token contract (WEB3_TOKEN_ADDRESS).`,
	)
}
