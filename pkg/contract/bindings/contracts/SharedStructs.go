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

// SharedStructsMetaData contains all meta data concerning the SharedStructs contract.
var SharedStructsMetaData = &bind.MetaData{
	ABI: "[]",
}

// SharedStructsABI is the input ABI used to generate the binding from.
// Deprecated: Use SharedStructsMetaData.ABI instead.
var SharedStructsABI = SharedStructsMetaData.ABI

// SharedStructs is an auto generated Go binding around an Ethereum contract.
type SharedStructs struct {
	SharedStructsCaller     // Read-only binding to the contract
	SharedStructsTransactor // Write-only binding to the contract
	SharedStructsFilterer   // Log filterer for contract events
}

// SharedStructsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SharedStructsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharedStructsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SharedStructsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharedStructsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SharedStructsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharedStructsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SharedStructsSession struct {
	Contract     *SharedStructs    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SharedStructsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SharedStructsCallerSession struct {
	Contract *SharedStructsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SharedStructsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SharedStructsTransactorSession struct {
	Contract     *SharedStructsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SharedStructsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SharedStructsRaw struct {
	Contract *SharedStructs // Generic contract binding to access the raw methods on
}

// SharedStructsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SharedStructsCallerRaw struct {
	Contract *SharedStructsCaller // Generic read-only contract binding to access the raw methods on
}

// SharedStructsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SharedStructsTransactorRaw struct {
	Contract *SharedStructsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSharedStructs creates a new instance of SharedStructs, bound to a specific deployed contract.
func NewSharedStructs(address common.Address, backend bind.ContractBackend) (*SharedStructs, error) {
	contract, err := bindSharedStructs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SharedStructs{SharedStructsCaller: SharedStructsCaller{contract: contract}, SharedStructsTransactor: SharedStructsTransactor{contract: contract}, SharedStructsFilterer: SharedStructsFilterer{contract: contract}}, nil
}

// NewSharedStructsCaller creates a new read-only instance of SharedStructs, bound to a specific deployed contract.
func NewSharedStructsCaller(address common.Address, caller bind.ContractCaller) (*SharedStructsCaller, error) {
	contract, err := bindSharedStructs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SharedStructsCaller{contract: contract}, nil
}

// NewSharedStructsTransactor creates a new write-only instance of SharedStructs, bound to a specific deployed contract.
func NewSharedStructsTransactor(address common.Address, transactor bind.ContractTransactor) (*SharedStructsTransactor, error) {
	contract, err := bindSharedStructs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SharedStructsTransactor{contract: contract}, nil
}

// NewSharedStructsFilterer creates a new log filterer instance of SharedStructs, bound to a specific deployed contract.
func NewSharedStructsFilterer(address common.Address, filterer bind.ContractFilterer) (*SharedStructsFilterer, error) {
	contract, err := bindSharedStructs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SharedStructsFilterer{contract: contract}, nil
}

// bindSharedStructs binds a generic wrapper to an already deployed contract.
func bindSharedStructs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SharedStructsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SharedStructs *SharedStructsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SharedStructs.Contract.SharedStructsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SharedStructs *SharedStructsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SharedStructs.Contract.SharedStructsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SharedStructs *SharedStructsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SharedStructs.Contract.SharedStructsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SharedStructs *SharedStructsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SharedStructs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SharedStructs *SharedStructsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SharedStructs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SharedStructs *SharedStructsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SharedStructs.Contract.contract.Transact(opts, method, params...)
}
