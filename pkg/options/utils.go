package options

import (
	"fmt"
	"os"
	"strconv"

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
		Web3: GetDefaultWeb3Options(),
	}
}

func GetDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
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
