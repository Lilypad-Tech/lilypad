//go:build unit

package web3_test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"github.com/BurntSushi/toml"
	"github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"go.opentelemetry.io/otel/trace/noop"
)

func getConfig() (*options.Config, error) {
	var config options.Config

	config_toml, err := os.ReadFile("../options/configs/testnet.toml") // Changed fs.ReadFile to os.ReadFile

	_, err = toml.Decode(string(config_toml), &config)
	if err != nil {
		return nil, errors.New("unable to parse config file")
	}

	return &config, nil
}

func CreateTestWeb3SDK() (*web3.Web3SDK, error) {
	// Load environment variables from .local.dev

	config, err := getConfig()
	if err != nil {
		log.Fatalf("Error loading config file")
	}

	options := web3.Web3Options{
		RpcURL:            config.Web3.RpcURL,
		ChainID:           config.Web3.ChainID,
		PrivateKey:        config.Web3.PrivateKey,
		ControllerAddress: config.Web3.ControllerAddress,
		PaymentsAddress:   config.Web3.PaymentsAddress,
		StorageAddress:    config.Web3.StorageAddress,
		UsersAddress:      config.Web3.UsersAddress,
		TokenAddress:      config.Web3.TokenAddress,
		MediationAddress:  config.Web3.MediationAddress,
		JobCreatorAddress: config.Web3.JobCreatorAddress,
	}

	fmt.Println("options: ", options)

	noopTracer := noop.NewTracerProvider().Tracer(system.GetOTelServiceName(system.DefaultService))
	sdk, err := web3.NewContractSDK(context.Background(), options, noopTracer, nil)
	if err != nil {
		return nil, err
	}

	return sdk, nil
}

func TestGetBalance(t *testing.T) {
	t.Skip("Issue: https://github.com/Lilypad-Tech/lilypad/issues/385")

	sdk, err := CreateTestWeb3SDK()
	if err != nil {
		t.Fatalf("Failed to create Web3SDK: %v", err)
	}
	balance, err := sdk.GetBalance("0x0000000000000000000000000000000000000000")
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}

	t.Logf("Balance: %d\n", balance)
}

func TestGetLPBalance(t *testing.T) {
	t.Skip("Issue: https://github.com/Lilypad-Tech/lilypad/issues/385")

	sdk, err := CreateTestWeb3SDK()
	if err != nil {
		t.Fatalf("Failed to create Web3SDK: %v", err)
	}
	balance, err := sdk.GetLPBalance("0x9162B48910E12079c089477DeF4384312f0a6E00") // faucet address
	if err != nil {
		t.Fatalf("Failed to get LP balance: %v", err)
	}
	t.Logf("LP Balance: %d\n", balance)
}

func TestGetLPBalanceNoBalance(t *testing.T) {
	t.Skip("Issue: https://github.com/Lilypad-Tech/lilypad/issues/385")

	sdk, err := CreateTestWeb3SDK()
	noBalanceInt := big.NewInt(0)
	if err != nil {
		t.Fatalf("Failed to create Web3SDK: %v", err)
	}

	// generate new address with no balance
	newAddress := generateNewAddressWoLp()

	t.Logf("newAddress to check no LP balance: %s\n", newAddress)

	balance, err := sdk.GetLPBalance(newAddress) // randomly generated address (100% no LP balance)
	if err != nil {
		t.Fatalf("Failed to get LP balance: %v", err)
	}

	if balance.Cmp(noBalanceInt) > 0 {
		t.Fatalf("Balance should be nil")
	}

	t.Logf("LP Balance: %d\n", balance)
}

func TestGetBlockNumber(t *testing.T) {
	t.Skip("Issue: https://github.com/Lilypad-Tech/lilypad/issues/385")

	sdk, err := CreateTestWeb3SDK()
	if err != nil {
		t.Fatalf("Failed to create Web3SDK: %v", err)
	}
	var blockNumberHex string
	err = sdk.Client.Client().Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		t.Fatalf("error for getBlockNumber: %s", err.Error())
	}

	t.Logf("Block number: %s\n", blockNumberHex)
}

func generateNewAddressWoLp() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	return address

}
