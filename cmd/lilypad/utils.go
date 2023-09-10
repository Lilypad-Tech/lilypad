package lilypad

import (
	"os"
	"strconv"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
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

/*
pricing options
*/
func addPricingCliFlags(cmd *cobra.Command, pricingConfig data.PricingConfig) {
	cmd.PersistentFlags().StringVar(
		&pricingConfig.InstructionPrice, "pricing-instruction-price", pricingConfig.InstructionPrice,
		`The price per instruction to offer (PRICING_INSTRUCTION_PRICE)`,
	)
	cmd.PersistentFlags().StringVar(
		&pricingConfig.Timeout, "pricing-timeout", pricingConfig.Timeout,
		`The timeout seconds (PRICING_TIMEOUT)`,
	)
	cmd.PersistentFlags().StringVar(
		&pricingConfig.TimeoutCollateral, "pricing-timeout-collateral", pricingConfig.TimeoutCollateral,
		`The timeout collateral (PRICING_TIMEOUT_COLLATERAL)`,
	)
	cmd.PersistentFlags().StringVar(
		&pricingConfig.PaymentCollateral, "pricing-payment-collateral", pricingConfig.PaymentCollateral,
		`The payment collateral (PRICING_PAYMENT_COLLATERAL)`,
	)
	cmd.PersistentFlags().StringVar(
		&pricingConfig.ResultsCollateralMultiple, "pricing-results-collateral-multiple", pricingConfig.ResultsCollateralMultiple,
		`The results collateral multiple (PRICING_RESULTS_COLLATERAL_MULTIPLE)`,
	)
	cmd.PersistentFlags().StringVar(
		&pricingConfig.MediationFee, "pricing-mediation-fee", pricingConfig.MediationFee,
		`The mediation fee (PRICING_MEDIATION_FEE)`,
	)
}

/*
module options
*/
func addModuleCliFlags(cmd *cobra.Command, moduleConfig data.Module) {
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Repo, "module-repo", moduleConfig.Repo,
		`The (http) git repo we can close (MODULE_REPO)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Hash, "module-hash", moduleConfig.Hash,
		`The hash of the repo we can checkout (MODULE_HASH)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Path, "module-path", moduleConfig.Path,
		`The path in the repo to the go template (MODULE_PATH)`,
	)
}

/*
resource provider options
*/
func addResourceProviderOfferCliFlags(cmd *cobra.Command, offerOptions resourceprovider.ResourceProviderOfferOptions) {
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.CPU, "offer-cpu", offerOptions.OfferSpec.CPU,
		`How many milli-cpus to offer the network (OFFER_CPU).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.GPU, "offer-gpu", offerOptions.OfferSpec.GPU,
		`How many milli-gpus to offer the network (OFFER_GPU).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferSpec.RAM, "offer-ram", offerOptions.OfferSpec.RAM,
		`How many megabytes of RAM to offer the network (OFFER_RAM).`,
	)
	cmd.PersistentFlags().IntVar(
		&offerOptions.OfferCount, "offer-count", offerOptions.OfferCount,
		`How many machines will we offer using the cpu, ram and gpu settings (OFFER_COUNT).`,
	)
	cmd.PersistentFlags().StringArrayVar(
		&offerOptions.Modules, "offer-modules", offerOptions.Modules,
		`The modules you are willing to run (OFFER_MODULES).`,
	)
	addPricingCliFlags(cmd, offerOptions.DefaultPricing)
}

/*
job creator options
*/
func addJobCreatorOfferCliFlags(cmd *cobra.Command, offerOptions jobcreator.JobCreatorOfferOptions) {
	cmd.PersistentFlags().BoolVar(
		&offerOptions.MarketOrder, "offer-market", offerOptions.MarketOrder,
		`Ignore the pricing config and pick the lowest priced resource offer.`,
	)
	addPricingCliFlags(cmd, offerOptions.Pricing)
}
