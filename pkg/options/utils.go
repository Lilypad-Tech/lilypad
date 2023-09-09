package options

import (
	"os"
	"strconv"

	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

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

/*
web3 options
*/
func GetDefaultWeb3Options() web3.Web3Options {
	return web3.Web3Options{
		RpcURL:            GetDefaultServeOptionString("WEB3_RPC_URL", ""),
		PrivateKey:        GetDefaultServeOptionString("WEB3_PRIVATE_KEY", ""),
		ChainID:           GetDefaultServeOptionInt("WEB3_CHAIN_ID", 1337), //nolint:gomnd
		ControllerAddress: GetDefaultServeOptionString("WEB3_CONTROLLER_ADDRESS", ""),
		PaymentsAddress:   GetDefaultServeOptionString("WEB3_PAYMENTS_ADDRESS", ""),
		StorageAddress:    GetDefaultServeOptionString("WEB3_STORAGE_ADDRESS", ""),
		TokenAddress:      GetDefaultServeOptionString("WEB3_TOKEN_ADDRESS", ""),
		SolverAddress:     GetDefaultServeOptionString("WEB3_SOLVER_ADDRESS", ""),
	}
}
