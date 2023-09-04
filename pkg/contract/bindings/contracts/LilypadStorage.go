// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SharedStructsAgreement is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsAgreement struct {
	State                    uint8
	Mediator                 common.Address
	ResourceProviderAgreedAt *big.Int
	JobCreatorAgreedAt       *big.Int
	DealAgreedAt             *big.Int
	ResultsSubmittedAt       *big.Int
	ResultsAcceptedAt        *big.Int
	ResultsCheckedAt         *big.Int
	MediationAcceptedAt      *big.Int
	MediationRejectedAt      *big.Int
	TimeoutSubmitResultsAt   *big.Int
	TimeoutJudgeResultsAt    *big.Int
	TimeoutMediateResultsAt  *big.Int
}

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId                    *big.Int
	ResourceProvider          common.Address
	JobCreator                common.Address
	InstructionPrice          *big.Int
	Timeout                   *big.Int
	TimeoutCollateral         *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	DealId           *big.Int
	ResultsId        *big.Int
	InstructionCount *big.Int
}

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress        common.Address
	MetadataCID        *big.Int
	Url                string
	Roles              []uint8
	TrustedMediators   []common.Address
	TrustedDirectories []common.Address
}

// LilypadStorageMetaData contains all meta data concerning the LilypadStorage contract.
var LilypadStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadStorageMetaData.ABI instead.
var LilypadStorageABI = LilypadStorageMetaData.ABI

// LilypadStorage is an auto generated Go binding around an Ethereum contract.
type LilypadStorage struct {
	LilypadStorageCaller     // Read-only binding to the contract
	LilypadStorageTransactor // Write-only binding to the contract
	LilypadStorageFilterer   // Log filterer for contract events
}

// LilypadStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadStorageSession struct {
	Contract     *LilypadStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LilypadStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadStorageCallerSession struct {
	Contract *LilypadStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LilypadStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadStorageTransactorSession struct {
	Contract     *LilypadStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LilypadStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadStorageRaw struct {
	Contract *LilypadStorage // Generic contract binding to access the raw methods on
}

// LilypadStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadStorageCallerRaw struct {
	Contract *LilypadStorageCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadStorageTransactorRaw struct {
	Contract *LilypadStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadStorage creates a new instance of LilypadStorage, bound to a specific deployed contract.
func NewLilypadStorage(address common.Address, backend bind.ContractBackend) (*LilypadStorage, error) {
	contract, err := bindLilypadStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadStorage{LilypadStorageCaller: LilypadStorageCaller{contract: contract}, LilypadStorageTransactor: LilypadStorageTransactor{contract: contract}, LilypadStorageFilterer: LilypadStorageFilterer{contract: contract}}, nil
}

// NewLilypadStorageCaller creates a new read-only instance of LilypadStorage, bound to a specific deployed contract.
func NewLilypadStorageCaller(address common.Address, caller bind.ContractCaller) (*LilypadStorageCaller, error) {
	contract, err := bindLilypadStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageCaller{contract: contract}, nil
}

// NewLilypadStorageTransactor creates a new write-only instance of LilypadStorage, bound to a specific deployed contract.
func NewLilypadStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadStorageTransactor, error) {
	contract, err := bindLilypadStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTransactor{contract: contract}, nil
}

// NewLilypadStorageFilterer creates a new log filterer instance of LilypadStorage, bound to a specific deployed contract.
func NewLilypadStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadStorageFilterer, error) {
	contract, err := bindLilypadStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageFilterer{contract: contract}, nil
}

// bindLilypadStorage binds a generic wrapper to an already deployed contract.
func bindLilypadStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadStorage *LilypadStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadStorage.Contract.LilypadStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadStorage *LilypadStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorage.Contract.LilypadStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadStorage *LilypadStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadStorage.Contract.LilypadStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadStorage *LilypadStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadStorage *LilypadStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadStorage *LilypadStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageCaller) GetAgreement(opts *bind.CallOpts, dealId *big.Int) (SharedStructsAgreement, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getAgreement", dealId)

	if err != nil {
		return *new(SharedStructsAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsAgreement)).(*SharedStructsAgreement)

	return out0, err

}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _LilypadStorage.Contract.GetAgreement(&_LilypadStorage.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageCallerSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _LilypadStorage.Contract.GetAgreement(&_LilypadStorage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageCaller) GetDeal(opts *bind.CallOpts, dealId *big.Int) (SharedStructsDeal, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _LilypadStorage.Contract.GetDeal(&_LilypadStorage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageCallerSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _LilypadStorage.Contract.GetDeal(&_LilypadStorage.CallOpts, dealId)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_LilypadStorage *LilypadStorageCaller) GetDealsForParty(opts *bind.CallOpts, party common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getDealsForParty", party)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_LilypadStorage *LilypadStorageSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _LilypadStorage.Contract.GetDealsForParty(&_LilypadStorage.CallOpts, party)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_LilypadStorage *LilypadStorageCallerSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _LilypadStorage.Contract.GetDealsForParty(&_LilypadStorage.CallOpts, party)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_LilypadStorage *LilypadStorageCaller) GetJobCost(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getJobCost", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_LilypadStorage *LilypadStorageSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorage.Contract.GetJobCost(&_LilypadStorage.CallOpts, dealId)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_LilypadStorage *LilypadStorageCallerSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorage.Contract.GetJobCost(&_LilypadStorage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageCaller) GetResult(opts *bind.CallOpts, dealId *big.Int) (SharedStructsResult, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getResult", dealId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _LilypadStorage.Contract.GetResult(&_LilypadStorage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageCallerSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _LilypadStorage.Contract.GetResult(&_LilypadStorage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_LilypadStorage *LilypadStorageCaller) GetResultsCollateral(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getResultsCollateral", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_LilypadStorage *LilypadStorageSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorage.Contract.GetResultsCollateral(&_LilypadStorage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_LilypadStorage *LilypadStorageCallerSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorage.Contract.GetResultsCollateral(&_LilypadStorage.CallOpts, dealId)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorage *LilypadStorageCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getUser", userAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorage *LilypadStorageSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _LilypadStorage.Contract.GetUser(&_LilypadStorage.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorage *LilypadStorageCallerSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _LilypadStorage.Contract.GetUser(&_LilypadStorage.CallOpts, userAddress)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_LilypadStorage *LilypadStorageCaller) HasDeal(opts *bind.CallOpts, dealId *big.Int) (bool, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "hasDeal", dealId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_LilypadStorage *LilypadStorageSession) HasDeal(dealId *big.Int) (bool, error) {
	return _LilypadStorage.Contract.HasDeal(&_LilypadStorage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_LilypadStorage *LilypadStorageCallerSession) HasDeal(dealId *big.Int) (bool, error) {
	return _LilypadStorage.Contract.HasDeal(&_LilypadStorage.CallOpts, dealId)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_LilypadStorage *LilypadStorageCaller) IsState(opts *bind.CallOpts, dealId *big.Int, state uint8) (bool, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "isState", dealId, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_LilypadStorage *LilypadStorageSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _LilypadStorage.Contract.IsState(&_LilypadStorage.CallOpts, dealId, state)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_LilypadStorage *LilypadStorageCallerSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _LilypadStorage.Contract.IsState(&_LilypadStorage.CallOpts, dealId, state)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadStorage *LilypadStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadStorage *LilypadStorageSession) Owner() (common.Address, error) {
	return _LilypadStorage.Contract.Owner(&_LilypadStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadStorage *LilypadStorageCallerSession) Owner() (common.Address, error) {
	return _LilypadStorage.Contract.Owner(&_LilypadStorage.CallOpts)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_LilypadStorage *LilypadStorageCaller) ShowUsersInList(opts *bind.CallOpts, serviceType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "showUsersInList", serviceType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_LilypadStorage *LilypadStorageSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _LilypadStorage.Contract.ShowUsersInList(&_LilypadStorage.CallOpts, serviceType)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_LilypadStorage *LilypadStorageCallerSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _LilypadStorage.Contract.ShowUsersInList(&_LilypadStorage.CallOpts, serviceType)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AcceptResult(&_LilypadStorage.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AcceptResult(&_LilypadStorage.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AddResult(&_LilypadStorage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactorSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AddResult(&_LilypadStorage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_LilypadStorage *LilypadStorageTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_LilypadStorage *LilypadStorageSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AddUserToList(&_LilypadStorage.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AddUserToList(&_LilypadStorage.TransactOpts, serviceType)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AgreeJobCreator(&_LilypadStorage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactorSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AgreeJobCreator(&_LilypadStorage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AgreeResourceProvider(&_LilypadStorage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactorSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.AgreeResourceProvider(&_LilypadStorage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadStorage *LilypadStorageTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadStorage *LilypadStorageSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.CheckResult(&_LilypadStorage.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.CheckResult(&_LilypadStorage.TransactOpts, dealId, mediator)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadStorage *LilypadStorageTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadStorage *LilypadStorageSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadStorage.Contract.DisableChangeControllerAddress(&_LilypadStorage.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadStorage *LilypadStorageTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadStorage.Contract.DisableChangeControllerAddress(&_LilypadStorage.TransactOpts)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactor) EnsureDeal(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "ensureDeal", dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.EnsureDeal(&_LilypadStorage.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorage *LilypadStorageTransactorSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.EnsureDeal(&_LilypadStorage.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadStorage *LilypadStorageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadStorage *LilypadStorageSession) Initialize() (*types.Transaction, error) {
	return _LilypadStorage.Contract.Initialize(&_LilypadStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadStorage *LilypadStorageTransactorSession) Initialize() (*types.Transaction, error) {
	return _LilypadStorage.Contract.Initialize(&_LilypadStorage.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.MediationAcceptResult(&_LilypadStorage.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.MediationAcceptResult(&_LilypadStorage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.MediationRejectResult(&_LilypadStorage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.MediationRejectResult(&_LilypadStorage.TransactOpts, dealId)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_LilypadStorage *LilypadStorageTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_LilypadStorage *LilypadStorageSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.RemoveUserFromList(&_LilypadStorage.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.RemoveUserFromList(&_LilypadStorage.TransactOpts, serviceType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadStorage *LilypadStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadStorage *LilypadStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadStorage.Contract.RenounceOwnership(&_LilypadStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadStorage *LilypadStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadStorage.Contract.RenounceOwnership(&_LilypadStorage.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadStorage *LilypadStorageTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadStorage *LilypadStorageSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SetControllerAddress(&_LilypadStorage.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SetControllerAddress(&_LilypadStorage.TransactOpts, _controllerAddress)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TimeoutJudgeResult(&_LilypadStorage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TimeoutJudgeResult(&_LilypadStorage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TimeoutMediateResult(&_LilypadStorage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TimeoutMediateResult(&_LilypadStorage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TimeoutSubmitResult(&_LilypadStorage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TimeoutSubmitResult(&_LilypadStorage.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadStorage *LilypadStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadStorage *LilypadStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TransferOwnership(&_LilypadStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.TransferOwnership(&_LilypadStorage.TransactOpts, newOwner)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorage *LilypadStorageTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "updateUser", metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorage *LilypadStorageSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.UpdateUser(&_LilypadStorage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorage *LilypadStorageTransactorSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.UpdateUser(&_LilypadStorage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// LilypadStorageDealStateChangeIterator is returned from FilterDealStateChange and is used to iterate over the raw logs and unpacked data for DealStateChange events raised by the LilypadStorage contract.
type LilypadStorageDealStateChangeIterator struct {
	Event *LilypadStorageDealStateChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LilypadStorageDealStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageDealStateChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LilypadStorageDealStateChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LilypadStorageDealStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageDealStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageDealStateChange represents a DealStateChange event raised by the LilypadStorage contract.
type LilypadStorageDealStateChange struct {
	DealId *big.Int
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_LilypadStorage *LilypadStorageFilterer) FilterDealStateChange(opts *bind.FilterOpts, dealId []*big.Int, state []uint8) (*LilypadStorageDealStateChangeIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageDealStateChangeIterator{contract: _LilypadStorage.contract, event: "DealStateChange", logs: logs, sub: sub}, nil
}

// WatchDealStateChange is a free log subscription operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_LilypadStorage *LilypadStorageFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *LilypadStorageDealStateChange, dealId []*big.Int, state []uint8) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageDealStateChange)
				if err := _LilypadStorage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDealStateChange is a log parse operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_LilypadStorage *LilypadStorageFilterer) ParseDealStateChange(log types.Log) (*LilypadStorageDealStateChange, error) {
	event := new(LilypadStorageDealStateChange)
	if err := _LilypadStorage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadStorage contract.
type LilypadStorageInitializedIterator struct {
	Event *LilypadStorageInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LilypadStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LilypadStorageInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LilypadStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageInitialized represents a Initialized event raised by the LilypadStorage contract.
type LilypadStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadStorage *LilypadStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadStorageInitializedIterator, error) {

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadStorageInitializedIterator{contract: _LilypadStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadStorage *LilypadStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageInitialized)
				if err := _LilypadStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadStorage *LilypadStorageFilterer) ParseInitialized(log types.Log) (*LilypadStorageInitialized, error) {
	event := new(LilypadStorageInitialized)
	if err := _LilypadStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LilypadStorage contract.
type LilypadStorageOwnershipTransferredIterator struct {
	Event *LilypadStorageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LilypadStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LilypadStorageOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LilypadStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageOwnershipTransferred represents a OwnershipTransferred event raised by the LilypadStorage contract.
type LilypadStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadStorage *LilypadStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LilypadStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageOwnershipTransferredIterator{contract: _LilypadStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadStorage *LilypadStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LilypadStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageOwnershipTransferred)
				if err := _LilypadStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadStorage *LilypadStorageFilterer) ParseOwnershipTransferred(log types.Log) (*LilypadStorageOwnershipTransferred, error) {
	event := new(LilypadStorageOwnershipTransferred)
	if err := _LilypadStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
