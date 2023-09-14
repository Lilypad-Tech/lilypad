package web3

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strconv"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/controller"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/mediation"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/payments"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/users"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

// these are the go-binding wrappers for the various deployed contracts
type Contracts struct {
	Token      *token.Token
	Payments   *payments.Payments
	Storage    *storage.Storage
	Users      *users.Users
	Mediation  *mediation.Mediation
	Controller *controller.Controller
}

type Web3SDK struct {
	Service      system.Service
	Options      Web3Options
	PrivateKey   *ecdsa.PrivateKey
	Client       *ethclient.Client
	CallOpts     *bind.CallOpts
	TransactOpts *bind.TransactOpts
	Contracts    *Contracts
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

	if paymentsAddress == "" {
		loadedPaymentsAddress, err := controller.GetPaymentsAddress(callOpts)
		if err != nil {
			return nil, err
		}
		paymentsAddress = loadedPaymentsAddress.String()
		log.Debug().
			Str("load payments address", paymentsAddress).
			Msgf("")
	}

	payments, err := payments.NewPayments(common.HexToAddress(paymentsAddress), client)
	if err != nil {
		return nil, err
	}

	tokenAddress := options.TokenAddress

	if tokenAddress == "" {
		loadedTokenAddress, err := payments.GetTokenAddress(callOpts)
		if err != nil {
			return nil, err
		}
		tokenAddress = loadedTokenAddress.String()
		log.Debug().
			Str("load token address", tokenAddress).
			Msgf("")
	}

	token, err := token.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		return nil, err
	}

	storageAddress := options.StorageAddress

	if storageAddress == "" {
		loadedStorageAddress, err := controller.GetStorageAddress(callOpts)
		if err != nil {
			return nil, err
		}
		storageAddress = loadedStorageAddress.String()
		log.Debug().
			Str("load storage address", storageAddress).
			Msgf("")
	}

	storage, err := storage.NewStorage(common.HexToAddress(storageAddress), client)
	if err != nil {
		return nil, err
	}

	usersAddress := options.UsersAddress

	if usersAddress == "" {
		loadedUsersAddress, err := controller.GetUsersAddress(callOpts)
		if err != nil {
			return nil, err
		}
		usersAddress = loadedUsersAddress.String()
		log.Debug().
			Str("load users address", usersAddress).
			Msgf("")
	}

	users, err := users.NewUsers(common.HexToAddress(usersAddress), client)
	if err != nil {
		return nil, err
	}

	mediationAddress := options.MediationAddress

	if mediationAddress == "" {
		loadedMediationAddress, err := controller.GetMediationAddress(callOpts)
		if err != nil {
			return nil, err
		}
		mediationAddress = loadedMediationAddress.String()
		log.Debug().
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
		Mediation:  mediation,
		Controller: controller,
	}, nil
}

func NewContractSDK(options Web3Options) (*Web3SDK, error) {
	client, err := ethclient.Dial(options.RpcURL)
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
	return &Web3SDK{
		PrivateKey:   privateKey,
		Options:      options,
		Client:       client,
		CallOpts:     callOpts,
		TransactOpts: transactOpts,
		Contracts:    contracts,
	}, nil
}

func (sdk *Web3SDK) getBlockNumber() (uint64, error) {
	var blockNumberHex string
	err := sdk.Client.Client().Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		log.Error().Msgf("error for getBlockNumber: %s", err.Error())
		return 0, err
	}
	blockNumberHex = strings.TrimPrefix(blockNumberHex, "0x")
	return strconv.ParseUint(blockNumberHex, 16, 64)
}

func (sdk *Web3SDK) waitTx(tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(context.Background(), sdk.Client, tx)
}

func (sdk *Web3SDK) GetAddress() common.Address {
	return crypto.PubkeyToAddress(GetPublicKey(sdk.PrivateKey))
}
