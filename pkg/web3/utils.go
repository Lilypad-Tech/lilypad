package web3

import "fmt"

func checkOptions(options Web3Options) error {
	// check each of the Web3Options for empty values and return and error
	// if there are any found - otherwise return nil for error
	if options.RpcURL == "" {
		return fmt.Errorf("--web3-rpc-url or WEB3_RPC_URL is empty")
	}
	if options.PrivateKey == "" {
		return fmt.Errorf("--web3-private-key or WEB3_PRIVATE_KEY is empty")
	}
	if options.ChainID == 0 {
		return fmt.Errorf("--web3-chain-id or WEB3_CHAIN_ID is empty")
	}
	if options.ControllerAddress == "" {
		return fmt.Errorf("--web3-controller-address or WEB3_CONTROLLER_ADDRESS is empty")
	}
	if options.PaymentsAddress == "" {
		return fmt.Errorf("--web3-payments-address or WEB3_PAYMENTS_ADDRESS is empty")
	}
	if options.StorageAddress == "" {
		return fmt.Errorf("--web3-storage-address or WEB3_STORAGE_ADDRESS is empty")
	}
	if options.TokenAddress == "" {
		return fmt.Errorf("--web3-token-address or WEB3_TOKEN_ADDRESS is empty")
	}
	return nil
}
