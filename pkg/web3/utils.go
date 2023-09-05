package web3

import "fmt"

func checkOptions(options Web3Options) error {
	// check each of the Web3Options for empty values and return and error
	// if there are any found - otherwise return nil for error
	if options.RpcURL == "" {
		return fmt.Errorf("RpcURL is empty")
	}
	if options.PrivateKey == "" {
		return fmt.Errorf("PrivateKey is empty")
	}
	if options.ChainID == 0 {
		return fmt.Errorf("ChainID is empty")
	}
	if options.ControllerAddress == "" {
		return fmt.Errorf("ControllerAddress is empty")
	}
	if options.PaymentsAddress == "" {
		return fmt.Errorf("PaymentsAddress is empty")
	}
	if options.StorageAddress == "" {
		return fmt.Errorf("StorageAddress is empty")
	}
	if options.TokenAddress == "" {
		return fmt.Errorf("TokenAddress is empty")
	}
	return nil
}
