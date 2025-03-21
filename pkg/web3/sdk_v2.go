package web3

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog"
	"math/big"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadContractRegistry"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadModuleDirectory"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadPaymentEngine"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadProxy"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadStorage"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadToken"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadTokenomics"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindingsV2/LilypadUser"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var logger = system.GetLogger(system.Web3Service)

// These are the Go binding wrappers for the new version of the protocol
type ContractsV2 struct {
	LilypadPaymentEngine    *lilypadpaymentengine.LilypadPaymentEngine
	LilypadProxy            *lilypadproxy.LilypadProxy
	LilypadModuleDirectory  *lilypadmoduledirectory.LilypadModuleDirectory
	LilypadContractRegistry *lilypadcontractregistry.LilypadContractRegistry
	LilypadStorage          *lilypadstorage.LilypadStorage
	LilypadUser             *lilypaduser.LilypadUser
	LilypadTokenomics       *lilypadtokenomics.LilypadTokenomics
	LilypadToken            *lilypadtoken.LilypadToken
}

type Web3SDKV2 struct {
	Options      Web3Options
	PrivateKey   *ecdsa.PrivateKey
	Client       *ethclient.Client
	CallOpts     *bind.CallOpts
	TransactOpts *bind.TransactOpts
	Contracts    *ContractsV2
	Log          *zerolog.Logger
}

// Creates the new Go-binding wrapper modules for the new protocol
func NewContractsV2(
	options Web3Options,
	client *ethclient.Client,
	callOpts *bind.CallOpts,
) (*ContractsV2, error) {
	lilypadProxy, err := lilypadproxy.NewLilypadProxy(common.HexToAddress(options.LilypadProxyAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadModuleDirectory, err := lilypadmoduledirectory.NewLilypadModuleDirectory(common.HexToAddress(options.LilypadModuleDirectoryAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadPaymentEngine, err := lilypadpaymentengine.NewLilypadPaymentEngine(common.HexToAddress(options.LilypadPaymentEngineAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadContractRegistry, err := lilypadcontractregistry.NewLilypadContractRegistry(common.HexToAddress(options.LilypadContractRegistryAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadStorage, err := lilypadstorage.NewLilypadStorage(common.HexToAddress(options.LilypadStorageAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadUser, err := lilypaduser.NewLilypadUser(common.HexToAddress(options.LilypadUserAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadTokenomics, err := lilypadtokenomics.NewLilypadTokenomics(common.HexToAddress(options.LilypadTokenomicsAddress), client)
	if err != nil {
		return nil, err
	}

	lilypadToken, err := lilypadtoken.NewLilypadToken(common.HexToAddress(options.LilypadL2TokenAddress), client)
	if err != nil {
		return nil, err
	}

	return &ContractsV2{
		LilypadPaymentEngine:    lilypadPaymentEngine,
		LilypadProxy:            lilypadProxy,
		LilypadContractRegistry: lilypadContractRegistry,
		LilypadModuleDirectory:  lilypadModuleDirectory,
		LilypadStorage:          lilypadStorage,
		LilypadToken:            lilypadToken,
		LilypadUser:             lilypadUser,
		LilypadTokenomics:       lilypadTokenomics,
	}, nil
}

func NewContractSDKV2(ctx context.Context, options Web3Options, tracer trace.Tracer) (*Web3SDKV2, error) {
	displayOpts := options
	displayOpts.PrivateKey = "*********"
	logger.Debug().Msgf("NewContractSDK: %+v", displayOpts)

	client, err := getEthClient(ctx, options, tracer)
	if err != nil {
		return nil, err
	}

	privateKey, err := ParsePrivateKey(options.PrivateKey)
	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(options.ChainID)))
	if err != nil {
		return nil, err
	}
	contracts, err := NewContractsV2(options, client, callOpts)
	if err != nil {
		return nil, err
	}

	web3SDK := &Web3SDKV2{
		PrivateKey:   privateKey,
		Options:      options,
		Client:       client,
		CallOpts:     callOpts,
		TransactOpts: transactOpts,
		Contracts:    contracts,
		Log:          logger,
	}

	return web3SDK, nil
}

func (sdk *Web3SDKV2) getEthClient(ctx context.Context, options Web3Options, tracer trace.Tracer) (*ethclient.Client, error) {
	ctx, span := tracer.Start(ctx, "get_ethclient", trace.WithAttributes(attribute.Int("web3.chain_id", options.ChainID)))
	defer span.End()

	rpcs := strings.Split(options.RpcURL, ",")
	var err error
	var parsedURL *url.URL
	var client *ethclient.Client
	for _, u := range rpcs {
		parsedURL, err = url.Parse(u)
		if err != nil {
			logger.Warn().Msgf("Unable to parse web3 RPC URL: %v", err)
			span.RecordError(errors.New("Unable to parse web3 RPC URL"))
			continue
		}

		span.AddEvent("ethclient.dial", trace.WithAttributes(attribute.String("web3.rpc_url", parsedURL.Host)))
		client, err = ethclient.Dial(u)
		if err != nil {
			logger.Warn().Msgf("Failed to connect to %s: %v", parsedURL.Host, err)
			span.RecordError(fmt.Errorf("Failed to connect to %s", parsedURL.Host))
			continue
		} else {
			logger.Info().Msgf("Connected to %s", parsedURL.Host)
			span.AddEvent("ethclient.connected")
			break
		}
	}
	if client == nil {
		span.SetStatus(codes.Error, "Failed to connect with web3 RPC URL")
		return nil, errors.New("Failed to connect to a web3 RPC provider")
	}

	return client, nil
}

func (sdk *Web3SDKV2) getBlockNumber() (uint64, error) {
	var blockNumberHex string
	err := sdk.Client.Client().Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		logger.Error().Msgf("error for getBlockNumber: %s", err.Error())
		return 0, err
	}
	blockNumberHex = strings.TrimPrefix(blockNumberHex, "0x")
	return strconv.ParseUint(blockNumberHex, 16, 64)
}

func (sdk *Web3SDKV2) WaitTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, sdk.Client, tx)
}

func (sdk *Web3SDKV2) GetAddress() common.Address {
	return crypto.PubkeyToAddress(GetPublicKey(sdk.PrivateKey))
}

func (sdk *Web3SDKV2) GetBalance(address string) (*big.Int, error) {
	// Convert the string address to common.Address
	ethAddress := common.HexToAddress(address)

	// Get the balance using the converted address
	balance, err := sdk.Client.BalanceAt(context.Background(), ethAddress, nil)
	if err != nil {
		logger.Error().Msgf("error for GetBalance: %s", err.Error())
		return nil, err
	}
	return balance, nil
}
