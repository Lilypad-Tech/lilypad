package lilypad

import (
	"os"
	"strconv"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/resourceprovider"
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
func addServerCliFlags(cmd *cobra.Command, serverOptions http.ServerOptions) {
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

	cmd.PersistentFlags().StringVar(
		&web3Options.TokenAddress, "web3-solver-address", web3Options.SolverAddress,
		`The address of the solver service (WEB3_SOLVER_ADDRESS).`,
	)
}

func addResourceProviderOfferCliFlags(cmd *cobra.Command, offerOptions resourceprovider.ResourceProviderOfferOptions) {
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.CPU, "cpu", offerOptions.OfferSpec.CPU,
		`How many milli-cpus to offer the network.`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.GPU, "gpu", offerOptions.OfferSpec.GPU,
		`How many milli-gpus to offer the network.`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.RAM, "ram", offerOptions.OfferSpec.RAM,
		`How many megabytes of RAM to offer the network.`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferCount, "ram", offerOptions.OfferCount,
		`How many machines will we offer using the cpu, ram and gpu settings.`,
	)
	cmd.PersistentFlags().StringArrayVar(
		&offerOptions.Modules, "modules", offerOptions.Modules,
		`The modules you are willing to run.`,
	)
}
