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
resource provider options
*/

func GetDefaultResourceProviderOfferOptions() resourceprovider.ResourceProviderOfferOptions {
	return resourceprovider.ResourceProviderOfferOptions{
		// by default let's offer 1 CPU, 0 GPU and 1GB RAM
		Machine: data.Spec{
			CPU: GetDefaultServeOptionInt("OFFER_CPU", 1),    //nolint:gomnd
			GPU: GetDefaultServeOptionInt("OFFER_GPU", 0),    //nolint:gomnd
			RAM: GetDefaultServeOptionInt("OFFER_RAM", 1024), //nolint:gomnd
		},
		MachineCount: GetDefaultServeOptionInt("OFFER_MACHINES", 1), //nolint:gomnd
		Specs:        []data.Spec{},
		Modules:      GetDefaultServeOptionStringArray("OFFER_MODULES", []string{}),
	}
}

func ProcessResourceProviderOfferOptions(options resourceprovider.ResourceProviderOfferOptions) resourceprovider.ResourceProviderOfferOptions {
	// if there are no specs then populate with the single spec
	if len(options.Specs) == 0 {
		// loop the number of machines we want to offer
		for i := 0; i < options.MachineCount; i++ {
			options.Specs = append(options.Specs, options.Machine)
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
