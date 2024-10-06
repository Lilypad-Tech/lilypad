package web3_test

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

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

	noopTracer := noop.NewTracerProvider().Tracer(system.GetOTelServiceName(system.DefaultService))
	sdk, err := web3.NewContractSDK(context.Background(), options, noopTracer)
	if err != nil {
		return nil, err
	}

	return sdk, nil
}

func TestGetBalance(t *testing.T) {
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

func TestGetBlockNumber(t *testing.T) {
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
