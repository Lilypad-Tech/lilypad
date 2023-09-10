package options

import (
	"fmt"
	"math/big"
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

func GetDefaultServeOptionBigInt(envName string, defaultValue big.Int) big.Int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return *big.NewInt(int64(i))
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
func GetDefaultPricingOptions() data.Pricing {
	return data.Pricing{
		// let's make the default price 1 ether
		InstructionPrice: GetDefaultServeOptionBigInt("PRICING_INSTRUCTION_PRICE", *web3.EtherToWei(1)),
		// 1 hour timeout
		Timeout: GetDefaultServeOptionBigInt("PRICING_TIMEOUT", *big.NewInt(60 * 60)),
		// 1 ether for timeout collateral
		TimeoutCollateral: GetDefaultServeOptionBigInt("PRICING_TIMEOUT_COLLATERAL", *web3.EtherToWei(1)),
		// 2 x ether for payment collateral (assuming modules that have a single instruction count)
		PaymentCollateral: GetDefaultServeOptionBigInt("PRICING_PAYMENT_COLLATERAL", *web3.EtherToWei(2)),
		// 2 x results collateral multiple
		ResultsCollateralMultiple: GetDefaultServeOptionBigInt("PRICING_RESULTS_COLLATERAL_MULTIPLE", *big.NewInt(2)),
		// 1 ether for mediation fee
		MediationFee: GetDefaultServeOptionBigInt("PRICING_MEDIATION_FEE", *web3.EtherToWei(1)),
	}
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
		ModulePricing:  map[string]data.Pricing{},
	}
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
