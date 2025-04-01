package web3

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/lilypad-tech/lilypad/v2/pkg/system"
	"github.com/rs/zerolog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/controller"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/jobcreator"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/mediation"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/payments"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/pow"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/storage"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/token"
	"github.com/lilypad-tech/lilypad/v2/pkg/web3/bindings/users"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// LP token ABI
const erc20ABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256","internalType": "uint256"}],"type":"function"}]`

var sdkLogger = system.GetLogger(system.Web3Service)

// these are the go-binding wrappers for the various deployed contracts
type Contracts struct {
	Token      *token.Token
	Payments   *payments.Payments
	Storage    *storage.Storage
	Users      *users.Users
	JobCreator *jobcreator.Jobcreator
	Mediation  *mediation.Mediation
	Controller *controller.Controller
	Pow        *pow.Pow
}

type Web3SDK struct {
	Options      Web3Options
	PrivateKey   *ecdsa.PrivateKey
	Client       *ethclient.Client
	CallOpts     *bind.CallOpts
	TransactOpts *bind.TransactOpts
	Contracts    *Contracts
	Log          *zerolog.Logger
}

func NewContracts(
	options Web3Options,
	client *ethclient.Client,
	callOpts *bind.CallOpts,
) (*Contracts, error) {
	controller, err := controller.NewController(common.HexToAddress(options.ControllerAddress), client)
	if err != nil {
		return nil, err
	}

	paymentsAddress := options.PaymentsAddress
	sdkLogger.Debug().Msgf("paymentsAddress: %s", paymentsAddress)
	if paymentsAddress == "" {
		loadedPaymentsAddress, err := controller.GetPaymentsAddress(callOpts)
		if err != nil {
			return nil, err
		}
		paymentsAddress = loadedPaymentsAddress.String()
		sdkLogger.Debug().
			Str("load payments address", paymentsAddress).
			Msgf("")
	}
	payments, err := payments.NewPayments(common.HexToAddress(paymentsAddress), client)
	if err != nil {
		return nil, err
	}

	powAddress := options.PowAddress
	sdkLogger.Debug().Msgf("PowAddress: %s", powAddress)
	if powAddress == "" {
		loadedPowAddress, err := controller.GetPowAddress(callOpts)
		if err != nil {
			return nil, err
		}
		powAddress = loadedPowAddress.String()
		sdkLogger.Debug().
			Str("load pow address", powAddress).
			Msgf("")
	}

	pow, err := pow.NewPow(common.HexToAddress(powAddress), client)
	if err != nil {
		return nil, err
	}

	tokenAddress := options.TokenAddress
	sdkLogger.Debug().Msgf("TokenAddress: %s", tokenAddress)
	if tokenAddress == "" {
		loadedTokenAddress, err := payments.GetTokenAddress(callOpts)
		if err != nil {
			return nil, err
		}
		tokenAddress = loadedTokenAddress.String()
		sdkLogger.Debug().
			Str("load token address", tokenAddress).
			Msgf("")
	}

	token, err := token.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		return nil, err
	}

	storageAddress := options.StorageAddress
	sdkLogger.Debug().Msgf("StorageAddress: %s", storageAddress)
	if storageAddress == "" {
		loadedStorageAddress, err := controller.GetStorageAddress(callOpts)
		if err != nil {
			return nil, err
		}
		storageAddress = loadedStorageAddress.String()
		sdkLogger.Debug().
			Str("load storage address", storageAddress).
			Msgf("")
	}

	storage, err := storage.NewStorage(common.HexToAddress(storageAddress), client)
	if err != nil {
		return nil, err
	}

	usersAddress := options.UsersAddress
	sdkLogger.Debug().Msgf("UsersAddress: %s", usersAddress)
	if usersAddress == "" {
		loadedUsersAddress, err := controller.GetUsersAddress(callOpts)
		if err != nil {
			return nil, err
		}
		usersAddress = loadedUsersAddress.String()
		sdkLogger.Debug().
			Str("load users address", usersAddress).
			Msgf("")
	}

	users, err := users.NewUsers(common.HexToAddress(usersAddress), client)
	if err != nil {
		return nil, err
	}

	jobcreatorAddress := options.JobCreatorAddress
	sdkLogger.Debug().Msgf("JobCreatorAddress: %s", jobcreatorAddress)
	if jobcreatorAddress == "" {
		loadedJobCreatorAddress, err := controller.GetJobCreatorAddress(callOpts)
		if err != nil {
			return nil, err
		}
		jobcreatorAddress = loadedJobCreatorAddress.String()
		sdkLogger.Debug().
			Str("load jobcreator address", jobcreatorAddress).
			Msgf("")
	}

	jobCreator, err := jobcreator.NewJobcreator(common.HexToAddress(jobcreatorAddress), client)
	if err != nil {
		return nil, err
	}

	mediationAddress := options.MediationAddress
	sdkLogger.Debug().Msgf("MediationAddress: %s", mediationAddress)
	if mediationAddress == "" {
		loadedMediationAddress, err := controller.GetMediationAddress(callOpts)
		if err != nil {
			return nil, err
		}
		mediationAddress = loadedMediationAddress.String()
		sdkLogger.Debug().
			Str("load mediation address", mediationAddress).
			Msgf("")
	}

	mediation, err := mediation.NewMediation(common.HexToAddress(mediationAddress), client)
	if err != nil {
		return nil, err
	}

	return &Contracts{
		Token:      token,
		Payments:   payments,
		Storage:    storage,
		Users:      users,
		JobCreator: jobCreator,
		Mediation:  mediation,
		Controller: controller,
		Pow:        pow,
	}, nil
}

func NewContractSDK(ctx context.Context, options Web3Options, tracer trace.Tracer) (*Web3SDK, error) {
	displayOpts := options
	displayOpts.PrivateKey = "*********"
	sdkLogger.Debug().Msgf("NewContractSDK: %+v", displayOpts)

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
	contracts, err := NewContracts(options, client, callOpts)
	if err != nil {
		return nil, err
	}

	web3SDK := &Web3SDK{
		PrivateKey:   privateKey,
		Options:      options,
		Client:       client,
		CallOpts:     callOpts,
		TransactOpts: transactOpts,
		Contracts:    contracts,
		Log:          logger,
	}
	sdkLogger.Info().Msgf("Public Address: %s", web3SDK.GetAddress())

	return web3SDK, nil
}

func getEthClient(ctx context.Context, options Web3Options, tracer trace.Tracer) (*ethclient.Client, error) {
	ctx, span := tracer.Start(ctx, "get_ethclient", trace.WithAttributes(attribute.Int("web3.chain_id", options.ChainID)))
	defer span.End()

	rpcs := strings.Split(options.RpcURL, ",")
	var err error
	var parsedURL *url.URL
	var client *ethclient.Client
	for _, u := range rpcs {
		parsedURL, err = url.Parse(u)
		if err != nil {
			sdkLogger.Warn().Msgf("Unable to parse web3 RPC URL: %v", err)
			span.RecordError(errors.New("Unable to parse web3 RPC URL"))
			continue
		}

		span.AddEvent("ethclient.dial", trace.WithAttributes(attribute.String("web3.rpc_url", parsedURL.Host)))
		client, err = ethclient.Dial(u)
		if err != nil {
			sdkLogger.Warn().Msgf("Failed to connect to %s: %v", parsedURL.Host, err)
			span.RecordError(fmt.Errorf("Failed to connect to %s", parsedURL.Host))
			continue
		} else {
			sdkLogger.Info().Msgf("Connected to %s", parsedURL.Host)
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

func (sdk *Web3SDK) getBlockNumber() (uint64, error) {
	var blockNumberHex string
	err := sdk.Client.Client().Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		sdkLogger.Error().Msgf("error for getBlockNumber: %s", err.Error())
		return 0, err
	}
	blockNumberHex = strings.TrimPrefix(blockNumberHex, "0x")
	return strconv.ParseUint(blockNumberHex, 16, 64)
}

func (sdk *Web3SDK) WaitTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, sdk.Client, tx)
}

func (sdk *Web3SDK) GetAddress() common.Address {
	return crypto.PubkeyToAddress(GetPublicKey(sdk.PrivateKey))
}

func (sdk *Web3SDK) GetBalance(address string) (*big.Int, error) {
	// Convert the string address to common.Address
	ethAddress := common.HexToAddress(address)

	// Get the balance using the converted address
	balance, err := sdk.Client.BalanceAt(context.Background(), ethAddress, nil)
	if err != nil {
		sdkLogger.Error().Msgf("error for GetBalance: %s", err.Error())
		return nil, err
	}
	return balance, nil
}

func (sdk *Web3SDK) GetLPBalance(address string) (*big.Int, error) {
	// Convert the string address to common.Address
	client := sdk.Client
	ethAddress := common.HexToAddress(address)
	lpToken := common.HexToAddress(sdk.Options.TokenAddress) // LP Token Address

	// Get the balance using the converted address
	erc20ABIObj, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		sdkLogger.Error().Msgf("error parsing ABI: %s", err)
		return nil, err
	}
	tokenContract := bind.NewBoundContract(lpToken, erc20ABIObj, client, client, client)

	var out []interface{}
	err = tokenContract.Call(nil, &out, "balanceOf", ethAddress)
	if err != nil {
		sdkLogger.Error().Msgf("error calling balanceOf: %s", err)
		return nil, err
	}

	lpBalance := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)

	if err != nil {
		sdkLogger.Error().Msgf("error for GetBalance: %s", err.Error())
		return nil, err
	}
	return &lpBalance, nil
}
