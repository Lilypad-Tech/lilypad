package web3

import (
	"crypto/ecdsa"
	"math/big"
	"strconv"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/controller"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/payments"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

// these are the go-binding wrappers for the various deployed contracts
type Contracts struct {
	Token      *token.Token
	Payments   *payments.Payments
	Storage    *storage.Storage
	Controller *controller.Controller
}

type ContractSDK struct {
	Options    Web3Options
	PrivateKey *ecdsa.PrivateKey
	Client     *ethclient.Client
	Auth       *bind.TransactOpts
	Contracts  *Contracts
}

func NewContracts(options Web3Options, client *ethclient.Client) (*Contracts, error) {
	token, err := token.NewToken(common.HexToAddress(options.TokenAddress), client)
	if err != nil {
		return nil, err
	}
	payments, err := payments.NewPayments(common.HexToAddress(options.PaymentsAddress), client)
	if err != nil {
		return nil, err
	}
	storage, err := storage.NewStorage(common.HexToAddress(options.StorageAddress), client)
	if err != nil {
		return nil, err
	}
	controller, err := controller.NewController(common.HexToAddress(options.ControllerAddress), client)
	if err != nil {
		return nil, err
	}
	return &Contracts{
		Token:      token,
		Payments:   payments,
		Storage:    storage,
		Controller: controller,
	}, nil
}

func NewContractSDK(options Web3Options) (*ContractSDK, error) {
	optionsErr := checkOptions(options)
	if optionsErr != nil {
		return nil, optionsErr
	}
	client, err := ethclient.Dial(options.RpcURL)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(strings.Replace(options.PrivateKey, "0x", "", 1))
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(options.ChainID)))
	if err != nil {
		return nil, err
	}
	contracts, err := NewContracts(options, client)
	if err != nil {
		return nil, err
	}
	return &ContractSDK{
		PrivateKey: privateKey,
		Options:    options,
		Client:     client,
		Auth:       auth,
		Contracts:  contracts,
	}, nil
}

func (sdk *ContractSDK) getBlockNumber() (uint64, error) {
	var blockNumberHex string
	err := sdk.Client.Client().Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		log.Error().Msgf("error for getBlockNumber: %s", err.Error())
		return 0, err
	}
	blockNumberHex = strings.TrimPrefix(blockNumberHex, "0x")
	return strconv.ParseUint(blockNumberHex, 16, 64)
}

func (sdk *ContractSDK) GetAddress() common.Address {
	return crypto.PubkeyToAddress(GetPublicKey(sdk.PrivateKey))
}
