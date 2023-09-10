package options

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/directory"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	"github.com/bacalhau-project/lilypad/pkg/mediator"
	"github.com/bacalhau-project/lilypad/pkg/resourceprovider"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func NewSolverOptions() solver.SolverOptions {
	return solver.SolverOptions{
		Server: GetDefaultServerOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
}

func NewDirectoryOptions() directory.DirectoryOptions {
	return directory.DirectoryOptions{
		Server: GetDefaultServerOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
}

func NewJobCreatorOptions() jobcreator.JobCreatorOptions {
	return jobcreator.JobCreatorOptions{
		Web3: GetDefaultWeb3Options(),
	}
}

func NewMediatorOptions() mediator.MediatorOptions {
	return mediator.MediatorOptions{
		Web3: GetDefaultWeb3Options(),
	}
}

func NewResourceProviderOptions() resourceprovider.ResourceProviderOptions {
	return resourceprovider.ResourceProviderOptions{
		Offers: GetDefaultResourceProviderOfferOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
}

func GetDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func GetDefaultServeOptionStringArray(envName string, defaultValue []string) []string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return strings.Split(envValue, ",")
	}
	return defaultValue
}

func GetDefaultServeOptionInt(envName string, defaultValue int) int {
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
server options
*/
func GetDefaultServerOptions() http.ServerOptions {
	return http.ServerOptions{
		URL:  GetDefaultServeOptionString("SERVER_URL", ""),
		Host: GetDefaultServeOptionString("SERVER_HOST", "0.0.0.0"),
		Port: GetDefaultServeOptionInt("SERVER_PORT", 8080), //nolint:gomnd
	}
}

func AddServerCliFlags(cmd *cobra.Command, serverOptions http.ServerOptions) {
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

func CheckServerOptions(options http.ServerOptions) error {
	if options.URL == "" {
		return fmt.Errorf("SERVER_URL is required")
	}
	return nil
}

/*
web3 options
*/
func GetDefaultWeb3Options() web3.Web3Options {
	return web3.Web3Options{

		// core settings
		RpcURL:     GetDefaultServeOptionString("WEB3_RPC_URL", ""),
		PrivateKey: GetDefaultServeOptionString("WEB3_PRIVATE_KEY", ""),
		ChainID:    GetDefaultServeOptionInt("WEB3_CHAIN_ID", 1337), //nolint:gomnd

		// contract addresses
		ControllerAddress: GetDefaultServeOptionString("WEB3_CONTROLLER_ADDRESS", ""),
		PaymentsAddress:   GetDefaultServeOptionString("WEB3_PAYMENTS_ADDRESS", ""),
		StorageAddress:    GetDefaultServeOptionString("WEB3_STORAGE_ADDRESS", ""),
		TokenAddress:      GetDefaultServeOptionString("WEB3_TOKEN_ADDRESS", ""),

		// service addresses
		SolverAddress:    GetDefaultServeOptionString("WEB3_SOLVER_ADDRESS", ""),
		DirectoryAddress: GetDefaultServeOptionString("WEB3_DIRECTORY_ADDRESS", ""),
	}
}

func AddWeb3CliFlags(cmd *cobra.Command, web3Options web3.Web3Options) {
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

func CheckWeb3Options(options web3.Web3Options, checkForServices bool) error {

	// core settings
	if options.RpcURL == "" {
		return fmt.Errorf("WEB3_RPC_URL is required")
	}
	if options.PrivateKey == "" {
		return fmt.Errorf("WEB3_PRIVATE_KEY is required")
	}

	// contract addresses
	if options.ControllerAddress == "" {
		return fmt.Errorf("WEB3_CONTROLLER_ADDRESS is required")
	}
	if options.PaymentsAddress == "" {
		return fmt.Errorf("WEB3_PAYMENTS_ADDRESS is required")
	}
	if options.StorageAddress == "" {
		return fmt.Errorf("WEB3_STORAGE_ADDRESS is required")
	}
	if options.TokenAddress == "" {
		return fmt.Errorf("WEB3_TOKEN_ADDRESS is required")
	}

	if checkForServices {
		// service addresses
		if options.SolverAddress == "" {
			return fmt.Errorf("WEB3_SOLVER_ADDRESS is required")
		}
		if options.DirectoryAddress == "" {
			return fmt.Errorf("WEB3_DIRECTORY_ADDRESS is required")
		}
	}

	return nil
}

/*
pricing options
*/
func GetDefaultPricingOptions() data.PricingConfig {
	return data.PricingConfig{
		// let's make the default price 1 ether
		InstructionPrice: GetDefaultServeOptionString("PRICING_INSTRUCTION_PRICE", "1"),
		// 1 hour timeout
		Timeout: GetDefaultServeOptionString("PRICING_TIMEOUT", "3600"),
		// 1 ether for timeout collateral
		TimeoutCollateral: GetDefaultServeOptionString("PRICING_TIMEOUT_COLLATERAL", "1"),
		// 2 x ether for payment collateral (assuming modules that have a single instruction count)
		PaymentCollateral: GetDefaultServeOptionString("PRICING_PAYMENT_COLLATERAL", "2"),
		// 2 x results collateral multiple
		ResultsCollateralMultiple: GetDefaultServeOptionString("PRICING_RESULTS_COLLATERAL_MULTIPLE", "2"),
		// 1 ether for mediation fee
		MediationFee: GetDefaultServeOptionString("PRICING_MEDIATION_FEE", "1"),
	}
}

func AddPricingCliFlags(cmd *cobra.Command, pricingConfig data.PricingConfig) {
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
func GetDefaultModuleOptions() data.Module {
	return data.Module{
		// the repo we can clone from
		Repo: GetDefaultServeOptionString("MODULE_REPO", ""),
		// the hash to checkout the repo
		Hash: GetDefaultServeOptionString("MODULE_HASH", ""),
		// the path to the go template file
		Path: GetDefaultServeOptionString("MODULE_PATH", ""),
	}
}

func AddModuleCliFlags(cmd *cobra.Command, moduleConfig data.Module) {
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

func GetDefaultResourceProviderOfferOptions() resourceprovider.ResourceProviderOfferOptions {
	return resourceprovider.ResourceProviderOfferOptions{
		// by default let's offer 1 CPU, 0 GPU and 1GB RAM
		OfferSpec: data.Spec{
			CPU: GetDefaultServeOptionInt("OFFER_CPU", 1),    //nolint:gomnd
			GPU: GetDefaultServeOptionInt("OFFER_GPU", 0),    //nolint:gomnd
			RAM: GetDefaultServeOptionInt("OFFER_RAM", 1024), //nolint:gomnd
		},
		OfferCount:     GetDefaultServeOptionInt("OFFER_COUNT", 1), //nolint:gomnd
		Specs:          []data.Spec{},
		Modules:        GetDefaultServeOptionStringArray("OFFER_MODULES", []string{}),
		DefaultPricing: GetDefaultPricingOptions(),
		ModulePricing:  map[string]data.PricingConfig{},
	}
}

func AddResourceProviderOfferCliFlags(cmd *cobra.Command, offerOptions resourceprovider.ResourceProviderOfferOptions) {
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
	AddPricingCliFlags(cmd, offerOptions.DefaultPricing)
}

func ProcessResourceProviderOfferOptions(options resourceprovider.ResourceProviderOfferOptions) resourceprovider.ResourceProviderOfferOptions {
	// if there are no specs then populate with the single spec
	if len(options.Specs) == 0 {
		// loop the number of machines we want to offer
		for i := 0; i < options.OfferCount; i++ {
			options.Specs = append(options.Specs, options.OfferSpec)
		}
	}
	return options
}

func CheckResourceProviderOfferOptions(options resourceprovider.ResourceProviderOfferOptions) error {

	// loop over all specs and add up the total number of cpus
	totalCPU := 0
	for _, spec := range options.Specs {
		totalCPU += spec.CPU
	}

	if totalCPU <= 0 {
		return fmt.Errorf("OFFER_CPU cannot be zero")
	}

	// do the same for memory
	totalRAM := 0
	for _, spec := range options.Specs {
		totalRAM += spec.RAM
	}

	if totalRAM <= 0 {
		return fmt.Errorf("OFFER_RAM cannot be zero")
	}

	return nil
}

/*
job creator options
*/

func GetDefaultJobCreatorOfferOptions() jobcreator.JobCreatorOfferOptions {
	return jobcreator.JobCreatorOfferOptions{
		MarketOrder: true,
		Module:      GetDefaultModuleOptions(),
		Pricing:     GetDefaultPricingOptions(),
	}
}

func AddJobCreatorOfferCliFlags(cmd *cobra.Command, offerOptions jobcreator.JobCreatorOfferOptions) {
	cmd.PersistentFlags().BoolVar(
		&offerOptions.MarketOrder, "offer-market", offerOptions.MarketOrder,
		`Ignore the pricing config and pick the lowest priced resource offer.`,
	)
	AddPricingCliFlags(cmd, offerOptions.Pricing)
}
