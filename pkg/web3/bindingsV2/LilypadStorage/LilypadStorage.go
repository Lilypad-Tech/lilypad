// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypadstorage

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

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId           string
	JobCreator       common.Address
	ResourceProvider common.Address
	ModuleCreator    common.Address
	Solver           common.Address
	JobOfferCID      string
	ResourceOfferCID string
	Status           uint8
	Timestamp        *big.Int
	PaymentStructure SharedStructsDealPaymentStructure
}

// SharedStructsDealPaymentStructure is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPaymentStructure struct {
	JobCreatorSolverFee       *big.Int
	ResourceProviderSolverFee *big.Int
	NetworkCongestionFee      *big.Int
	ModuleCreatorFee          *big.Int
	PriceOfJobWithoutFees     *big.Int
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	ResultId  string
	DealId    string
	ResultCID string
	Status    uint8
	Timestamp *big.Int
}

// SharedStructsValidationResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsValidationResult struct {
	ValidationResultId string
	ResultId           string
	ValidationCID      string
	Status             uint8
	Timestamp          *big.Int
	Validator          common.Address
}

// LilypadStorageMetaData contains all meta data concerning the LilypadStorage contract.
var LilypadStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeDealStatus\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeResultStatus\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeValidationResultStatus\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkDealStatus\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkResultStatus\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkValidationResultStatus\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeal\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Deal\",\"components\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentStructure\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.DealPaymentStructure\",\"components\":[{\"name\":\"jobCreatorSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkCongestionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleCreatorFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceOfJobWithoutFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResult\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Result\",\"components\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidationResult\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.ValidationResult\",\"components\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validationCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveDeal\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deal\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Deal\",\"components\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentStructure\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.DealPaymentStructure\",\"components\":[{\"name\":\"jobCreatorSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkCongestionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleCreatorFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceOfJobWithoutFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveResult\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Result\",\"components\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saveValidationResult\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validationResult\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.ValidationResult\",\"components\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validationCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadStorage__DealSaved\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadStorage__DealStatusChanged\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSharedStructs.DealStatusEnum\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadStorage__ResultSaved\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadStorage__ResultStatusChanged\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSharedStructs.ResultStatusEnum\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadStorage__ValidationResultSaved\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"resultId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadStorage__ValidationResultStatusChanged\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__DealNotFound\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"LilypadStorage__EmptyCID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__EmptyDealId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__EmptyResultId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__EmptyValidationResultId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__InvalidJobCreatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__InvalidModuleCreatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__InvalidResourceProviderAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__InvalidSolverAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__InvalidValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__ResultNotFound\",\"inputs\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"LilypadStorage__SameAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadStorage__ValidationResultNotFound\",\"inputs\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"LilypadStorage__ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadStorage *LilypadStorageCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadStorage *LilypadStorageSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadStorage.Contract.DEFAULTADMINROLE(&_LilypadStorage.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadStorage *LilypadStorageCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadStorage.Contract.DEFAULTADMINROLE(&_LilypadStorage.CallOpts)
}

// CheckDealStatus is a free data retrieval call binding the contract method 0xd32c6323.
//
// Solidity: function checkDealStatus(string dealId) view returns(uint8)
func (_LilypadStorage *LilypadStorageCaller) CheckDealStatus(opts *bind.CallOpts, dealId string) (uint8, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "checkDealStatus", dealId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckDealStatus is a free data retrieval call binding the contract method 0xd32c6323.
//
// Solidity: function checkDealStatus(string dealId) view returns(uint8)
func (_LilypadStorage *LilypadStorageSession) CheckDealStatus(dealId string) (uint8, error) {
	return _LilypadStorage.Contract.CheckDealStatus(&_LilypadStorage.CallOpts, dealId)
}

// CheckDealStatus is a free data retrieval call binding the contract method 0xd32c6323.
//
// Solidity: function checkDealStatus(string dealId) view returns(uint8)
func (_LilypadStorage *LilypadStorageCallerSession) CheckDealStatus(dealId string) (uint8, error) {
	return _LilypadStorage.Contract.CheckDealStatus(&_LilypadStorage.CallOpts, dealId)
}

// CheckResultStatus is a free data retrieval call binding the contract method 0x137836e8.
//
// Solidity: function checkResultStatus(string resultId) view returns(uint8)
func (_LilypadStorage *LilypadStorageCaller) CheckResultStatus(opts *bind.CallOpts, resultId string) (uint8, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "checkResultStatus", resultId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckResultStatus is a free data retrieval call binding the contract method 0x137836e8.
//
// Solidity: function checkResultStatus(string resultId) view returns(uint8)
func (_LilypadStorage *LilypadStorageSession) CheckResultStatus(resultId string) (uint8, error) {
	return _LilypadStorage.Contract.CheckResultStatus(&_LilypadStorage.CallOpts, resultId)
}

// CheckResultStatus is a free data retrieval call binding the contract method 0x137836e8.
//
// Solidity: function checkResultStatus(string resultId) view returns(uint8)
func (_LilypadStorage *LilypadStorageCallerSession) CheckResultStatus(resultId string) (uint8, error) {
	return _LilypadStorage.Contract.CheckResultStatus(&_LilypadStorage.CallOpts, resultId)
}

// CheckValidationResultStatus is a free data retrieval call binding the contract method 0x98217c66.
//
// Solidity: function checkValidationResultStatus(string validationResultId) view returns(uint8)
func (_LilypadStorage *LilypadStorageCaller) CheckValidationResultStatus(opts *bind.CallOpts, validationResultId string) (uint8, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "checkValidationResultStatus", validationResultId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckValidationResultStatus is a free data retrieval call binding the contract method 0x98217c66.
//
// Solidity: function checkValidationResultStatus(string validationResultId) view returns(uint8)
func (_LilypadStorage *LilypadStorageSession) CheckValidationResultStatus(validationResultId string) (uint8, error) {
	return _LilypadStorage.Contract.CheckValidationResultStatus(&_LilypadStorage.CallOpts, validationResultId)
}

// CheckValidationResultStatus is a free data retrieval call binding the contract method 0x98217c66.
//
// Solidity: function checkValidationResultStatus(string validationResultId) view returns(uint8)
func (_LilypadStorage *LilypadStorageCallerSession) CheckValidationResultStatus(validationResultId string) (uint8, error) {
	return _LilypadStorage.Contract.CheckValidationResultStatus(&_LilypadStorage.CallOpts, validationResultId)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)))
func (_LilypadStorage *LilypadStorageCaller) GetDeal(opts *bind.CallOpts, dealId string) (SharedStructsDeal, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)))
func (_LilypadStorage *LilypadStorageSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _LilypadStorage.Contract.GetDeal(&_LilypadStorage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)))
func (_LilypadStorage *LilypadStorageCallerSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _LilypadStorage.Contract.GetDeal(&_LilypadStorage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string resultId) view returns((string,string,string,uint8,uint256))
func (_LilypadStorage *LilypadStorageCaller) GetResult(opts *bind.CallOpts, resultId string) (SharedStructsResult, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getResult", resultId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string resultId) view returns((string,string,string,uint8,uint256))
func (_LilypadStorage *LilypadStorageSession) GetResult(resultId string) (SharedStructsResult, error) {
	return _LilypadStorage.Contract.GetResult(&_LilypadStorage.CallOpts, resultId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string resultId) view returns((string,string,string,uint8,uint256))
func (_LilypadStorage *LilypadStorageCallerSession) GetResult(resultId string) (SharedStructsResult, error) {
	return _LilypadStorage.Contract.GetResult(&_LilypadStorage.CallOpts, resultId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadStorage *LilypadStorageCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadStorage *LilypadStorageSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadStorage.Contract.GetRoleAdmin(&_LilypadStorage.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadStorage *LilypadStorageCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadStorage.Contract.GetRoleAdmin(&_LilypadStorage.CallOpts, role)
}

// GetValidationResult is a free data retrieval call binding the contract method 0xd629490c.
//
// Solidity: function getValidationResult(string validationResultId) view returns((string,string,string,uint8,uint256,address))
func (_LilypadStorage *LilypadStorageCaller) GetValidationResult(opts *bind.CallOpts, validationResultId string) (SharedStructsValidationResult, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getValidationResult", validationResultId)

	if err != nil {
		return *new(SharedStructsValidationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsValidationResult)).(*SharedStructsValidationResult)

	return out0, err

}

// GetValidationResult is a free data retrieval call binding the contract method 0xd629490c.
//
// Solidity: function getValidationResult(string validationResultId) view returns((string,string,string,uint8,uint256,address))
func (_LilypadStorage *LilypadStorageSession) GetValidationResult(validationResultId string) (SharedStructsValidationResult, error) {
	return _LilypadStorage.Contract.GetValidationResult(&_LilypadStorage.CallOpts, validationResultId)
}

// GetValidationResult is a free data retrieval call binding the contract method 0xd629490c.
//
// Solidity: function getValidationResult(string validationResultId) view returns((string,string,string,uint8,uint256,address))
func (_LilypadStorage *LilypadStorageCallerSession) GetValidationResult(validationResultId string) (SharedStructsValidationResult, error) {
	return _LilypadStorage.Contract.GetValidationResult(&_LilypadStorage.CallOpts, validationResultId)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadStorage *LilypadStorageCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadStorage *LilypadStorageSession) GetVersion() (string, error) {
	return _LilypadStorage.Contract.GetVersion(&_LilypadStorage.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadStorage *LilypadStorageCallerSession) GetVersion() (string, error) {
	return _LilypadStorage.Contract.GetVersion(&_LilypadStorage.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadStorage *LilypadStorageCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadStorage *LilypadStorageSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadStorage.Contract.HasRole(&_LilypadStorage.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadStorage *LilypadStorageCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadStorage.Contract.HasRole(&_LilypadStorage.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadStorage *LilypadStorageCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadStorage *LilypadStorageSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadStorage.Contract.SupportsInterface(&_LilypadStorage.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadStorage *LilypadStorageCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadStorage.Contract.SupportsInterface(&_LilypadStorage.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadStorage *LilypadStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadStorage *LilypadStorageSession) Version() (string, error) {
	return _LilypadStorage.Contract.Version(&_LilypadStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadStorage *LilypadStorageCallerSession) Version() (string, error) {
	return _LilypadStorage.Contract.Version(&_LilypadStorage.CallOpts)
}

// ChangeDealStatus is a paid mutator transaction binding the contract method 0x0dd52b0f.
//
// Solidity: function changeDealStatus(string dealId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageTransactor) ChangeDealStatus(opts *bind.TransactOpts, dealId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "changeDealStatus", dealId, status)
}

// ChangeDealStatus is a paid mutator transaction binding the contract method 0x0dd52b0f.
//
// Solidity: function changeDealStatus(string dealId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageSession) ChangeDealStatus(dealId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.ChangeDealStatus(&_LilypadStorage.TransactOpts, dealId, status)
}

// ChangeDealStatus is a paid mutator transaction binding the contract method 0x0dd52b0f.
//
// Solidity: function changeDealStatus(string dealId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageTransactorSession) ChangeDealStatus(dealId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.ChangeDealStatus(&_LilypadStorage.TransactOpts, dealId, status)
}

// ChangeResultStatus is a paid mutator transaction binding the contract method 0xfd823a07.
//
// Solidity: function changeResultStatus(string resultId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageTransactor) ChangeResultStatus(opts *bind.TransactOpts, resultId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "changeResultStatus", resultId, status)
}

// ChangeResultStatus is a paid mutator transaction binding the contract method 0xfd823a07.
//
// Solidity: function changeResultStatus(string resultId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageSession) ChangeResultStatus(resultId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.ChangeResultStatus(&_LilypadStorage.TransactOpts, resultId, status)
}

// ChangeResultStatus is a paid mutator transaction binding the contract method 0xfd823a07.
//
// Solidity: function changeResultStatus(string resultId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageTransactorSession) ChangeResultStatus(resultId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.ChangeResultStatus(&_LilypadStorage.TransactOpts, resultId, status)
}

// ChangeValidationResultStatus is a paid mutator transaction binding the contract method 0xe7e14ffa.
//
// Solidity: function changeValidationResultStatus(string validationResultId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageTransactor) ChangeValidationResultStatus(opts *bind.TransactOpts, validationResultId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "changeValidationResultStatus", validationResultId, status)
}

// ChangeValidationResultStatus is a paid mutator transaction binding the contract method 0xe7e14ffa.
//
// Solidity: function changeValidationResultStatus(string validationResultId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageSession) ChangeValidationResultStatus(validationResultId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.ChangeValidationResultStatus(&_LilypadStorage.TransactOpts, validationResultId, status)
}

// ChangeValidationResultStatus is a paid mutator transaction binding the contract method 0xe7e14ffa.
//
// Solidity: function changeValidationResultStatus(string validationResultId, uint8 status) returns(bool)
func (_LilypadStorage *LilypadStorageTransactorSession) ChangeValidationResultStatus(validationResultId string, status uint8) (*types.Transaction, error) {
	return _LilypadStorage.Contract.ChangeValidationResultStatus(&_LilypadStorage.TransactOpts, validationResultId, status)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadStorage *LilypadStorageTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadStorage *LilypadStorageSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.GrantRole(&_LilypadStorage.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.GrantRole(&_LilypadStorage.TransactOpts, role, account)
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

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadStorage *LilypadStorageTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadStorage *LilypadStorageSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.RenounceRole(&_LilypadStorage.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.RenounceRole(&_LilypadStorage.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadStorage *LilypadStorageTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadStorage *LilypadStorageSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.RevokeRole(&_LilypadStorage.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadStorage *LilypadStorageTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadStorage.Contract.RevokeRole(&_LilypadStorage.TransactOpts, role, account)
}

// SaveDeal is a paid mutator transaction binding the contract method 0xcf7c125a.
//
// Solidity: function saveDeal(string dealId, (string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) deal) returns(bool)
func (_LilypadStorage *LilypadStorageTransactor) SaveDeal(opts *bind.TransactOpts, dealId string, deal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "saveDeal", dealId, deal)
}

// SaveDeal is a paid mutator transaction binding the contract method 0xcf7c125a.
//
// Solidity: function saveDeal(string dealId, (string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) deal) returns(bool)
func (_LilypadStorage *LilypadStorageSession) SaveDeal(dealId string, deal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SaveDeal(&_LilypadStorage.TransactOpts, dealId, deal)
}

// SaveDeal is a paid mutator transaction binding the contract method 0xcf7c125a.
//
// Solidity: function saveDeal(string dealId, (string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) deal) returns(bool)
func (_LilypadStorage *LilypadStorageTransactorSession) SaveDeal(dealId string, deal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SaveDeal(&_LilypadStorage.TransactOpts, dealId, deal)
}

// SaveResult is a paid mutator transaction binding the contract method 0x5747505a.
//
// Solidity: function saveResult(string resultId, (string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadStorage *LilypadStorageTransactor) SaveResult(opts *bind.TransactOpts, resultId string, result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "saveResult", resultId, result)
}

// SaveResult is a paid mutator transaction binding the contract method 0x5747505a.
//
// Solidity: function saveResult(string resultId, (string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadStorage *LilypadStorageSession) SaveResult(resultId string, result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SaveResult(&_LilypadStorage.TransactOpts, resultId, result)
}

// SaveResult is a paid mutator transaction binding the contract method 0x5747505a.
//
// Solidity: function saveResult(string resultId, (string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadStorage *LilypadStorageTransactorSession) SaveResult(resultId string, result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SaveResult(&_LilypadStorage.TransactOpts, resultId, result)
}

// SaveValidationResult is a paid mutator transaction binding the contract method 0x0825355e.
//
// Solidity: function saveValidationResult(string validationResultId, (string,string,string,uint8,uint256,address) validationResult) returns(bool)
func (_LilypadStorage *LilypadStorageTransactor) SaveValidationResult(opts *bind.TransactOpts, validationResultId string, validationResult SharedStructsValidationResult) (*types.Transaction, error) {
	return _LilypadStorage.contract.Transact(opts, "saveValidationResult", validationResultId, validationResult)
}

// SaveValidationResult is a paid mutator transaction binding the contract method 0x0825355e.
//
// Solidity: function saveValidationResult(string validationResultId, (string,string,string,uint8,uint256,address) validationResult) returns(bool)
func (_LilypadStorage *LilypadStorageSession) SaveValidationResult(validationResultId string, validationResult SharedStructsValidationResult) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SaveValidationResult(&_LilypadStorage.TransactOpts, validationResultId, validationResult)
}

// SaveValidationResult is a paid mutator transaction binding the contract method 0x0825355e.
//
// Solidity: function saveValidationResult(string validationResultId, (string,string,string,uint8,uint256,address) validationResult) returns(bool)
func (_LilypadStorage *LilypadStorageTransactorSession) SaveValidationResult(validationResultId string, validationResult SharedStructsValidationResult) (*types.Transaction, error) {
	return _LilypadStorage.Contract.SaveValidationResult(&_LilypadStorage.TransactOpts, validationResultId, validationResult)
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
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadStorage *LilypadStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadStorageInitializedIterator, error) {

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadStorageInitializedIterator{contract: _LilypadStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadStorage *LilypadStorageFilterer) ParseInitialized(log types.Log) (*LilypadStorageInitialized, error) {
	event := new(LilypadStorageInitialized)
	if err := _LilypadStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageLilypadStorageDealSavedIterator is returned from FilterLilypadStorageDealSaved and is used to iterate over the raw logs and unpacked data for LilypadStorageDealSaved events raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageDealSavedIterator struct {
	Event *LilypadStorageLilypadStorageDealSaved // Event containing the contract specifics and raw log

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
func (it *LilypadStorageLilypadStorageDealSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageLilypadStorageDealSaved)
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
		it.Event = new(LilypadStorageLilypadStorageDealSaved)
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
func (it *LilypadStorageLilypadStorageDealSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageLilypadStorageDealSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageLilypadStorageDealSaved represents a LilypadStorageDealSaved event raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageDealSaved struct {
	DealId           common.Hash
	JobCreator       common.Address
	ResourceProvider common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadStorageDealSaved is a free log retrieval operation binding the contract event 0xbb10347021eed142c482ff1963cc12e084311ab2420965d56bb48368d56c470a.
//
// Solidity: event LilypadStorage__DealSaved(string indexed dealId, address indexed jobCreator, address indexed resourceProvider)
func (_LilypadStorage *LilypadStorageFilterer) FilterLilypadStorageDealSaved(opts *bind.FilterOpts, dealId []string, jobCreator []common.Address, resourceProvider []common.Address) (*LilypadStorageLilypadStorageDealSavedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "LilypadStorage__DealSaved", dealIdRule, jobCreatorRule, resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageLilypadStorageDealSavedIterator{contract: _LilypadStorage.contract, event: "LilypadStorage__DealSaved", logs: logs, sub: sub}, nil
}

// WatchLilypadStorageDealSaved is a free log subscription operation binding the contract event 0xbb10347021eed142c482ff1963cc12e084311ab2420965d56bb48368d56c470a.
//
// Solidity: event LilypadStorage__DealSaved(string indexed dealId, address indexed jobCreator, address indexed resourceProvider)
func (_LilypadStorage *LilypadStorageFilterer) WatchLilypadStorageDealSaved(opts *bind.WatchOpts, sink chan<- *LilypadStorageLilypadStorageDealSaved, dealId []string, jobCreator []common.Address, resourceProvider []common.Address) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "LilypadStorage__DealSaved", dealIdRule, jobCreatorRule, resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageLilypadStorageDealSaved)
				if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__DealSaved", log); err != nil {
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

// ParseLilypadStorageDealSaved is a log parse operation binding the contract event 0xbb10347021eed142c482ff1963cc12e084311ab2420965d56bb48368d56c470a.
//
// Solidity: event LilypadStorage__DealSaved(string indexed dealId, address indexed jobCreator, address indexed resourceProvider)
func (_LilypadStorage *LilypadStorageFilterer) ParseLilypadStorageDealSaved(log types.Log) (*LilypadStorageLilypadStorageDealSaved, error) {
	event := new(LilypadStorageLilypadStorageDealSaved)
	if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__DealSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageLilypadStorageDealStatusChangedIterator is returned from FilterLilypadStorageDealStatusChanged and is used to iterate over the raw logs and unpacked data for LilypadStorageDealStatusChanged events raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageDealStatusChangedIterator struct {
	Event *LilypadStorageLilypadStorageDealStatusChanged // Event containing the contract specifics and raw log

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
func (it *LilypadStorageLilypadStorageDealStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageLilypadStorageDealStatusChanged)
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
		it.Event = new(LilypadStorageLilypadStorageDealStatusChanged)
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
func (it *LilypadStorageLilypadStorageDealStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageLilypadStorageDealStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageLilypadStorageDealStatusChanged represents a LilypadStorageDealStatusChanged event raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageDealStatusChanged struct {
	DealId common.Hash
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLilypadStorageDealStatusChanged is a free log retrieval operation binding the contract event 0x7e532b0625528a1b9bec2db501b49d27efabb6cae47ace7c861be3c50eca0f0c.
//
// Solidity: event LilypadStorage__DealStatusChanged(string indexed dealId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) FilterLilypadStorageDealStatusChanged(opts *bind.FilterOpts, dealId []string) (*LilypadStorageLilypadStorageDealStatusChangedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "LilypadStorage__DealStatusChanged", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageLilypadStorageDealStatusChangedIterator{contract: _LilypadStorage.contract, event: "LilypadStorage__DealStatusChanged", logs: logs, sub: sub}, nil
}

// WatchLilypadStorageDealStatusChanged is a free log subscription operation binding the contract event 0x7e532b0625528a1b9bec2db501b49d27efabb6cae47ace7c861be3c50eca0f0c.
//
// Solidity: event LilypadStorage__DealStatusChanged(string indexed dealId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) WatchLilypadStorageDealStatusChanged(opts *bind.WatchOpts, sink chan<- *LilypadStorageLilypadStorageDealStatusChanged, dealId []string) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "LilypadStorage__DealStatusChanged", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageLilypadStorageDealStatusChanged)
				if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__DealStatusChanged", log); err != nil {
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

// ParseLilypadStorageDealStatusChanged is a log parse operation binding the contract event 0x7e532b0625528a1b9bec2db501b49d27efabb6cae47ace7c861be3c50eca0f0c.
//
// Solidity: event LilypadStorage__DealStatusChanged(string indexed dealId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) ParseLilypadStorageDealStatusChanged(log types.Log) (*LilypadStorageLilypadStorageDealStatusChanged, error) {
	event := new(LilypadStorageLilypadStorageDealStatusChanged)
	if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__DealStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageLilypadStorageResultSavedIterator is returned from FilterLilypadStorageResultSaved and is used to iterate over the raw logs and unpacked data for LilypadStorageResultSaved events raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageResultSavedIterator struct {
	Event *LilypadStorageLilypadStorageResultSaved // Event containing the contract specifics and raw log

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
func (it *LilypadStorageLilypadStorageResultSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageLilypadStorageResultSaved)
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
		it.Event = new(LilypadStorageLilypadStorageResultSaved)
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
func (it *LilypadStorageLilypadStorageResultSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageLilypadStorageResultSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageLilypadStorageResultSaved represents a LilypadStorageResultSaved event raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageResultSaved struct {
	ResultId common.Hash
	DealId   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLilypadStorageResultSaved is a free log retrieval operation binding the contract event 0xca9d6e446353f560d04e3a4764ae74917e7a347db93588e048c8c5d1baeeaa2b.
//
// Solidity: event LilypadStorage__ResultSaved(string indexed resultId, string dealId)
func (_LilypadStorage *LilypadStorageFilterer) FilterLilypadStorageResultSaved(opts *bind.FilterOpts, resultId []string) (*LilypadStorageLilypadStorageResultSavedIterator, error) {

	var resultIdRule []interface{}
	for _, resultIdItem := range resultId {
		resultIdRule = append(resultIdRule, resultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "LilypadStorage__ResultSaved", resultIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageLilypadStorageResultSavedIterator{contract: _LilypadStorage.contract, event: "LilypadStorage__ResultSaved", logs: logs, sub: sub}, nil
}

// WatchLilypadStorageResultSaved is a free log subscription operation binding the contract event 0xca9d6e446353f560d04e3a4764ae74917e7a347db93588e048c8c5d1baeeaa2b.
//
// Solidity: event LilypadStorage__ResultSaved(string indexed resultId, string dealId)
func (_LilypadStorage *LilypadStorageFilterer) WatchLilypadStorageResultSaved(opts *bind.WatchOpts, sink chan<- *LilypadStorageLilypadStorageResultSaved, resultId []string) (event.Subscription, error) {

	var resultIdRule []interface{}
	for _, resultIdItem := range resultId {
		resultIdRule = append(resultIdRule, resultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "LilypadStorage__ResultSaved", resultIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageLilypadStorageResultSaved)
				if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ResultSaved", log); err != nil {
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

// ParseLilypadStorageResultSaved is a log parse operation binding the contract event 0xca9d6e446353f560d04e3a4764ae74917e7a347db93588e048c8c5d1baeeaa2b.
//
// Solidity: event LilypadStorage__ResultSaved(string indexed resultId, string dealId)
func (_LilypadStorage *LilypadStorageFilterer) ParseLilypadStorageResultSaved(log types.Log) (*LilypadStorageLilypadStorageResultSaved, error) {
	event := new(LilypadStorageLilypadStorageResultSaved)
	if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ResultSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageLilypadStorageResultStatusChangedIterator is returned from FilterLilypadStorageResultStatusChanged and is used to iterate over the raw logs and unpacked data for LilypadStorageResultStatusChanged events raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageResultStatusChangedIterator struct {
	Event *LilypadStorageLilypadStorageResultStatusChanged // Event containing the contract specifics and raw log

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
func (it *LilypadStorageLilypadStorageResultStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageLilypadStorageResultStatusChanged)
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
		it.Event = new(LilypadStorageLilypadStorageResultStatusChanged)
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
func (it *LilypadStorageLilypadStorageResultStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageLilypadStorageResultStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageLilypadStorageResultStatusChanged represents a LilypadStorageResultStatusChanged event raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageResultStatusChanged struct {
	ResultId common.Hash
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLilypadStorageResultStatusChanged is a free log retrieval operation binding the contract event 0x0c5e4c648275611d9807d3bac6a7b2cbfcdb3663a3e93e131fa0ee4badbd4459.
//
// Solidity: event LilypadStorage__ResultStatusChanged(string indexed resultId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) FilterLilypadStorageResultStatusChanged(opts *bind.FilterOpts, resultId []string) (*LilypadStorageLilypadStorageResultStatusChangedIterator, error) {

	var resultIdRule []interface{}
	for _, resultIdItem := range resultId {
		resultIdRule = append(resultIdRule, resultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "LilypadStorage__ResultStatusChanged", resultIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageLilypadStorageResultStatusChangedIterator{contract: _LilypadStorage.contract, event: "LilypadStorage__ResultStatusChanged", logs: logs, sub: sub}, nil
}

// WatchLilypadStorageResultStatusChanged is a free log subscription operation binding the contract event 0x0c5e4c648275611d9807d3bac6a7b2cbfcdb3663a3e93e131fa0ee4badbd4459.
//
// Solidity: event LilypadStorage__ResultStatusChanged(string indexed resultId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) WatchLilypadStorageResultStatusChanged(opts *bind.WatchOpts, sink chan<- *LilypadStorageLilypadStorageResultStatusChanged, resultId []string) (event.Subscription, error) {

	var resultIdRule []interface{}
	for _, resultIdItem := range resultId {
		resultIdRule = append(resultIdRule, resultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "LilypadStorage__ResultStatusChanged", resultIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageLilypadStorageResultStatusChanged)
				if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ResultStatusChanged", log); err != nil {
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

// ParseLilypadStorageResultStatusChanged is a log parse operation binding the contract event 0x0c5e4c648275611d9807d3bac6a7b2cbfcdb3663a3e93e131fa0ee4badbd4459.
//
// Solidity: event LilypadStorage__ResultStatusChanged(string indexed resultId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) ParseLilypadStorageResultStatusChanged(log types.Log) (*LilypadStorageLilypadStorageResultStatusChanged, error) {
	event := new(LilypadStorageLilypadStorageResultStatusChanged)
	if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ResultStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageLilypadStorageValidationResultSavedIterator is returned from FilterLilypadStorageValidationResultSaved and is used to iterate over the raw logs and unpacked data for LilypadStorageValidationResultSaved events raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageValidationResultSavedIterator struct {
	Event *LilypadStorageLilypadStorageValidationResultSaved // Event containing the contract specifics and raw log

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
func (it *LilypadStorageLilypadStorageValidationResultSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageLilypadStorageValidationResultSaved)
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
		it.Event = new(LilypadStorageLilypadStorageValidationResultSaved)
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
func (it *LilypadStorageLilypadStorageValidationResultSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageLilypadStorageValidationResultSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageLilypadStorageValidationResultSaved represents a LilypadStorageValidationResultSaved event raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageValidationResultSaved struct {
	ValidationResultId common.Hash
	ResultId           string
	Validator          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLilypadStorageValidationResultSaved is a free log retrieval operation binding the contract event 0xeafb76e6bce6b88c68a784a25bf7f9acbb786ebb6435f9c5917da8cf977646e9.
//
// Solidity: event LilypadStorage__ValidationResultSaved(string indexed validationResultId, string resultId, address validator)
func (_LilypadStorage *LilypadStorageFilterer) FilterLilypadStorageValidationResultSaved(opts *bind.FilterOpts, validationResultId []string) (*LilypadStorageLilypadStorageValidationResultSavedIterator, error) {

	var validationResultIdRule []interface{}
	for _, validationResultIdItem := range validationResultId {
		validationResultIdRule = append(validationResultIdRule, validationResultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "LilypadStorage__ValidationResultSaved", validationResultIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageLilypadStorageValidationResultSavedIterator{contract: _LilypadStorage.contract, event: "LilypadStorage__ValidationResultSaved", logs: logs, sub: sub}, nil
}

// WatchLilypadStorageValidationResultSaved is a free log subscription operation binding the contract event 0xeafb76e6bce6b88c68a784a25bf7f9acbb786ebb6435f9c5917da8cf977646e9.
//
// Solidity: event LilypadStorage__ValidationResultSaved(string indexed validationResultId, string resultId, address validator)
func (_LilypadStorage *LilypadStorageFilterer) WatchLilypadStorageValidationResultSaved(opts *bind.WatchOpts, sink chan<- *LilypadStorageLilypadStorageValidationResultSaved, validationResultId []string) (event.Subscription, error) {

	var validationResultIdRule []interface{}
	for _, validationResultIdItem := range validationResultId {
		validationResultIdRule = append(validationResultIdRule, validationResultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "LilypadStorage__ValidationResultSaved", validationResultIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageLilypadStorageValidationResultSaved)
				if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ValidationResultSaved", log); err != nil {
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

// ParseLilypadStorageValidationResultSaved is a log parse operation binding the contract event 0xeafb76e6bce6b88c68a784a25bf7f9acbb786ebb6435f9c5917da8cf977646e9.
//
// Solidity: event LilypadStorage__ValidationResultSaved(string indexed validationResultId, string resultId, address validator)
func (_LilypadStorage *LilypadStorageFilterer) ParseLilypadStorageValidationResultSaved(log types.Log) (*LilypadStorageLilypadStorageValidationResultSaved, error) {
	event := new(LilypadStorageLilypadStorageValidationResultSaved)
	if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ValidationResultSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageLilypadStorageValidationResultStatusChangedIterator is returned from FilterLilypadStorageValidationResultStatusChanged and is used to iterate over the raw logs and unpacked data for LilypadStorageValidationResultStatusChanged events raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageValidationResultStatusChangedIterator struct {
	Event *LilypadStorageLilypadStorageValidationResultStatusChanged // Event containing the contract specifics and raw log

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
func (it *LilypadStorageLilypadStorageValidationResultStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageLilypadStorageValidationResultStatusChanged)
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
		it.Event = new(LilypadStorageLilypadStorageValidationResultStatusChanged)
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
func (it *LilypadStorageLilypadStorageValidationResultStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageLilypadStorageValidationResultStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageLilypadStorageValidationResultStatusChanged represents a LilypadStorageValidationResultStatusChanged event raised by the LilypadStorage contract.
type LilypadStorageLilypadStorageValidationResultStatusChanged struct {
	ValidationResultId common.Hash
	Status             uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLilypadStorageValidationResultStatusChanged is a free log retrieval operation binding the contract event 0x3d781fb7a561711bbbf33d528fe804ec4693ce3ec3e2ee5e8fd9a381ca22f37e.
//
// Solidity: event LilypadStorage__ValidationResultStatusChanged(string indexed validationResultId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) FilterLilypadStorageValidationResultStatusChanged(opts *bind.FilterOpts, validationResultId []string) (*LilypadStorageLilypadStorageValidationResultStatusChangedIterator, error) {

	var validationResultIdRule []interface{}
	for _, validationResultIdItem := range validationResultId {
		validationResultIdRule = append(validationResultIdRule, validationResultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "LilypadStorage__ValidationResultStatusChanged", validationResultIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageLilypadStorageValidationResultStatusChangedIterator{contract: _LilypadStorage.contract, event: "LilypadStorage__ValidationResultStatusChanged", logs: logs, sub: sub}, nil
}

// WatchLilypadStorageValidationResultStatusChanged is a free log subscription operation binding the contract event 0x3d781fb7a561711bbbf33d528fe804ec4693ce3ec3e2ee5e8fd9a381ca22f37e.
//
// Solidity: event LilypadStorage__ValidationResultStatusChanged(string indexed validationResultId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) WatchLilypadStorageValidationResultStatusChanged(opts *bind.WatchOpts, sink chan<- *LilypadStorageLilypadStorageValidationResultStatusChanged, validationResultId []string) (event.Subscription, error) {

	var validationResultIdRule []interface{}
	for _, validationResultIdItem := range validationResultId {
		validationResultIdRule = append(validationResultIdRule, validationResultIdItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "LilypadStorage__ValidationResultStatusChanged", validationResultIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageLilypadStorageValidationResultStatusChanged)
				if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ValidationResultStatusChanged", log); err != nil {
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

// ParseLilypadStorageValidationResultStatusChanged is a log parse operation binding the contract event 0x3d781fb7a561711bbbf33d528fe804ec4693ce3ec3e2ee5e8fd9a381ca22f37e.
//
// Solidity: event LilypadStorage__ValidationResultStatusChanged(string indexed validationResultId, uint8 status)
func (_LilypadStorage *LilypadStorageFilterer) ParseLilypadStorageValidationResultStatusChanged(log types.Log) (*LilypadStorageLilypadStorageValidationResultStatusChanged, error) {
	event := new(LilypadStorageLilypadStorageValidationResultStatusChanged)
	if err := _LilypadStorage.contract.UnpackLog(event, "LilypadStorage__ValidationResultStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadStorage contract.
type LilypadStorageRoleAdminChangedIterator struct {
	Event *LilypadStorageRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadStorageRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageRoleAdminChanged)
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
		it.Event = new(LilypadStorageRoleAdminChanged)
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
func (it *LilypadStorageRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadStorage contract.
type LilypadStorageRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadStorage *LilypadStorageFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadStorageRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageRoleAdminChangedIterator{contract: _LilypadStorage.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadStorage *LilypadStorageFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadStorageRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageRoleAdminChanged)
				if err := _LilypadStorage.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadStorage *LilypadStorageFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadStorageRoleAdminChanged, error) {
	event := new(LilypadStorageRoleAdminChanged)
	if err := _LilypadStorage.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadStorage contract.
type LilypadStorageRoleGrantedIterator struct {
	Event *LilypadStorageRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadStorageRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageRoleGranted)
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
		it.Event = new(LilypadStorageRoleGranted)
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
func (it *LilypadStorageRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageRoleGranted represents a RoleGranted event raised by the LilypadStorage contract.
type LilypadStorageRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadStorage *LilypadStorageFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadStorageRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageRoleGrantedIterator{contract: _LilypadStorage.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadStorage *LilypadStorageFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadStorageRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageRoleGranted)
				if err := _LilypadStorage.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadStorage *LilypadStorageFilterer) ParseRoleGranted(log types.Log) (*LilypadStorageRoleGranted, error) {
	event := new(LilypadStorageRoleGranted)
	if err := _LilypadStorage.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadStorage contract.
type LilypadStorageRoleRevokedIterator struct {
	Event *LilypadStorageRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadStorageRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageRoleRevoked)
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
		it.Event = new(LilypadStorageRoleRevoked)
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
func (it *LilypadStorageRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageRoleRevoked represents a RoleRevoked event raised by the LilypadStorage contract.
type LilypadStorageRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadStorage *LilypadStorageFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadStorageRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadStorage.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageRoleRevokedIterator{contract: _LilypadStorage.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadStorage *LilypadStorageFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadStorageRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadStorage.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageRoleRevoked)
				if err := _LilypadStorage.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadStorage *LilypadStorageFilterer) ParseRoleRevoked(log types.Log) (*LilypadStorageRoleRevoked, error) {
	event := new(LilypadStorageRoleRevoked)
	if err := _LilypadStorage.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
