package web3

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

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

func GetPublicKey(privateKey *ecdsa.PrivateKey) ecdsa.PublicKey {
	return privateKey.PublicKey
}

func GetAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(message)
	return crypto.Sign(hash.Bytes(), privateKey)
}

func GetAddressFromSignedMessage(message []byte, sig []byte) (common.Address, error) {
	hash := crypto.Keccak256Hash(message)
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		return common.Address{}, err
	}
	recoveredPubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*recoveredPubKey), nil
}

func ParsePrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.Replace(privateKey, "0x", "", 1))
}
