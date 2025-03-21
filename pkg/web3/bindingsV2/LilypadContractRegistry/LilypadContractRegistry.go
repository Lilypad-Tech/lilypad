// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypadcontractregistry

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

// LilypadContractRegistryMetaData contains all meta data concerning the LilypadContractRegistry contract.
var LilypadContractRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_l1LilypadTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2LilypadTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadUserAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadModuleDirectoryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadPaymentEngineAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadProxyAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadVestingAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1LilypadTokenAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2LilypadTokenAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lilypadModuleDirectoryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lilypadPaymentEngineAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lilypadProxyAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lilypadStorageAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lilypadUserAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lilypadVestingAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL2LilypadTokenAddress\",\"inputs\":[{\"name\":\"_l2LilypadTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadModuleDirectoryAddress\",\"inputs\":[{\"name\":\"_lilypadModuleDirectoryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadPaymentEngineAddress\",\"inputs\":[{\"name\":\"_lilypadPaymentEngineAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadProxyAddress\",\"inputs\":[{\"name\":\"_lilypadProxyAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadStorageAddress\",\"inputs\":[{\"name\":\"_lilypadStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadUserAddress\",\"inputs\":[{\"name\":\"_lilypadUserAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadVestingAddress\",\"inputs\":[{\"name\":\"_lilypadVestingAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadContractRegistry__ContractAddressSet\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilpadContractRegistry__NoZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadContractRegistry__NotController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// LilypadContractRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadContractRegistryMetaData.ABI instead.
var LilypadContractRegistryABI = LilypadContractRegistryMetaData.ABI

// LilypadContractRegistry is an auto generated Go binding around an Ethereum contract.
type LilypadContractRegistry struct {
	LilypadContractRegistryCaller     // Read-only binding to the contract
	LilypadContractRegistryTransactor // Write-only binding to the contract
	LilypadContractRegistryFilterer   // Log filterer for contract events
}

// LilypadContractRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadContractRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadContractRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadContractRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadContractRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadContractRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadContractRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadContractRegistrySession struct {
	Contract     *LilypadContractRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LilypadContractRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadContractRegistryCallerSession struct {
	Contract *LilypadContractRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// LilypadContractRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadContractRegistryTransactorSession struct {
	Contract     *LilypadContractRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// LilypadContractRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadContractRegistryRaw struct {
	Contract *LilypadContractRegistry // Generic contract binding to access the raw methods on
}

// LilypadContractRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadContractRegistryCallerRaw struct {
	Contract *LilypadContractRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadContractRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadContractRegistryTransactorRaw struct {
	Contract *LilypadContractRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadContractRegistry creates a new instance of LilypadContractRegistry, bound to a specific deployed contract.
func NewLilypadContractRegistry(address common.Address, backend bind.ContractBackend) (*LilypadContractRegistry, error) {
	contract, err := bindLilypadContractRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistry{LilypadContractRegistryCaller: LilypadContractRegistryCaller{contract: contract}, LilypadContractRegistryTransactor: LilypadContractRegistryTransactor{contract: contract}, LilypadContractRegistryFilterer: LilypadContractRegistryFilterer{contract: contract}}, nil
}

// NewLilypadContractRegistryCaller creates a new read-only instance of LilypadContractRegistry, bound to a specific deployed contract.
func NewLilypadContractRegistryCaller(address common.Address, caller bind.ContractCaller) (*LilypadContractRegistryCaller, error) {
	contract, err := bindLilypadContractRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryCaller{contract: contract}, nil
}

// NewLilypadContractRegistryTransactor creates a new write-only instance of LilypadContractRegistry, bound to a specific deployed contract.
func NewLilypadContractRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadContractRegistryTransactor, error) {
	contract, err := bindLilypadContractRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryTransactor{contract: contract}, nil
}

// NewLilypadContractRegistryFilterer creates a new log filterer instance of LilypadContractRegistry, bound to a specific deployed contract.
func NewLilypadContractRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadContractRegistryFilterer, error) {
	contract, err := bindLilypadContractRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryFilterer{contract: contract}, nil
}

// bindLilypadContractRegistry binds a generic wrapper to an already deployed contract.
func bindLilypadContractRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadContractRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadContractRegistry *LilypadContractRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadContractRegistry.Contract.LilypadContractRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadContractRegistry *LilypadContractRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.LilypadContractRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadContractRegistry *LilypadContractRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.LilypadContractRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadContractRegistry *LilypadContractRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadContractRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadContractRegistry *LilypadContractRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadContractRegistry *LilypadContractRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadContractRegistry *LilypadContractRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadContractRegistry.Contract.DEFAULTADMINROLE(&_LilypadContractRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadContractRegistry.Contract.DEFAULTADMINROLE(&_LilypadContractRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadContractRegistry *LilypadContractRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadContractRegistry.Contract.GetRoleAdmin(&_LilypadContractRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadContractRegistry.Contract.GetRoleAdmin(&_LilypadContractRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadContractRegistry *LilypadContractRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadContractRegistry.Contract.HasRole(&_LilypadContractRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadContractRegistry.Contract.HasRole(&_LilypadContractRegistry.CallOpts, role, account)
}

// L1LilypadTokenAddress is a free data retrieval call binding the contract method 0xdf502a43.
//
// Solidity: function l1LilypadTokenAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) L1LilypadTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "l1LilypadTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1LilypadTokenAddress is a free data retrieval call binding the contract method 0xdf502a43.
//
// Solidity: function l1LilypadTokenAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) L1LilypadTokenAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.L1LilypadTokenAddress(&_LilypadContractRegistry.CallOpts)
}

// L1LilypadTokenAddress is a free data retrieval call binding the contract method 0xdf502a43.
//
// Solidity: function l1LilypadTokenAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) L1LilypadTokenAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.L1LilypadTokenAddress(&_LilypadContractRegistry.CallOpts)
}

// L2LilypadTokenAddress is a free data retrieval call binding the contract method 0x8d416df3.
//
// Solidity: function l2LilypadTokenAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) L2LilypadTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "l2LilypadTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2LilypadTokenAddress is a free data retrieval call binding the contract method 0x8d416df3.
//
// Solidity: function l2LilypadTokenAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) L2LilypadTokenAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.L2LilypadTokenAddress(&_LilypadContractRegistry.CallOpts)
}

// L2LilypadTokenAddress is a free data retrieval call binding the contract method 0x8d416df3.
//
// Solidity: function l2LilypadTokenAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) L2LilypadTokenAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.L2LilypadTokenAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadModuleDirectoryAddress is a free data retrieval call binding the contract method 0xe528f9c8.
//
// Solidity: function lilypadModuleDirectoryAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) LilypadModuleDirectoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "lilypadModuleDirectoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LilypadModuleDirectoryAddress is a free data retrieval call binding the contract method 0xe528f9c8.
//
// Solidity: function lilypadModuleDirectoryAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) LilypadModuleDirectoryAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadModuleDirectoryAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadModuleDirectoryAddress is a free data retrieval call binding the contract method 0xe528f9c8.
//
// Solidity: function lilypadModuleDirectoryAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) LilypadModuleDirectoryAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadModuleDirectoryAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadPaymentEngineAddress is a free data retrieval call binding the contract method 0xea0b4f27.
//
// Solidity: function lilypadPaymentEngineAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) LilypadPaymentEngineAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "lilypadPaymentEngineAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LilypadPaymentEngineAddress is a free data retrieval call binding the contract method 0xea0b4f27.
//
// Solidity: function lilypadPaymentEngineAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) LilypadPaymentEngineAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadPaymentEngineAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadPaymentEngineAddress is a free data retrieval call binding the contract method 0xea0b4f27.
//
// Solidity: function lilypadPaymentEngineAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) LilypadPaymentEngineAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadPaymentEngineAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadProxyAddress is a free data retrieval call binding the contract method 0x38c477b9.
//
// Solidity: function lilypadProxyAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) LilypadProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "lilypadProxyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LilypadProxyAddress is a free data retrieval call binding the contract method 0x38c477b9.
//
// Solidity: function lilypadProxyAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) LilypadProxyAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadProxyAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadProxyAddress is a free data retrieval call binding the contract method 0x38c477b9.
//
// Solidity: function lilypadProxyAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) LilypadProxyAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadProxyAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadStorageAddress is a free data retrieval call binding the contract method 0x3832b2cd.
//
// Solidity: function lilypadStorageAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) LilypadStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "lilypadStorageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LilypadStorageAddress is a free data retrieval call binding the contract method 0x3832b2cd.
//
// Solidity: function lilypadStorageAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) LilypadStorageAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadStorageAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadStorageAddress is a free data retrieval call binding the contract method 0x3832b2cd.
//
// Solidity: function lilypadStorageAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) LilypadStorageAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadStorageAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadUserAddress is a free data retrieval call binding the contract method 0x01a7ef66.
//
// Solidity: function lilypadUserAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) LilypadUserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "lilypadUserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LilypadUserAddress is a free data retrieval call binding the contract method 0x01a7ef66.
//
// Solidity: function lilypadUserAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) LilypadUserAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadUserAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadUserAddress is a free data retrieval call binding the contract method 0x01a7ef66.
//
// Solidity: function lilypadUserAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) LilypadUserAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadUserAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadVestingAddress is a free data retrieval call binding the contract method 0x88240325.
//
// Solidity: function lilypadVestingAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) LilypadVestingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "lilypadVestingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LilypadVestingAddress is a free data retrieval call binding the contract method 0x88240325.
//
// Solidity: function lilypadVestingAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistrySession) LilypadVestingAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadVestingAddress(&_LilypadContractRegistry.CallOpts)
}

// LilypadVestingAddress is a free data retrieval call binding the contract method 0x88240325.
//
// Solidity: function lilypadVestingAddress() view returns(address)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) LilypadVestingAddress() (common.Address, error) {
	return _LilypadContractRegistry.Contract.LilypadVestingAddress(&_LilypadContractRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadContractRegistry *LilypadContractRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadContractRegistry.Contract.SupportsInterface(&_LilypadContractRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadContractRegistry.Contract.SupportsInterface(&_LilypadContractRegistry.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadContractRegistry *LilypadContractRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadContractRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadContractRegistry *LilypadContractRegistrySession) Version() (string, error) {
	return _LilypadContractRegistry.Contract.Version(&_LilypadContractRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadContractRegistry *LilypadContractRegistryCallerSession) Version() (string, error) {
	return _LilypadContractRegistry.Contract.Version(&_LilypadContractRegistry.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.GrantRole(&_LilypadContractRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.GrantRole(&_LilypadContractRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _l1LilypadTokenAddress, address _l2LilypadTokenAddress, address _lilypadUserAddress, address _lilypadModuleDirectoryAddress, address _lilypadStorageAddress, address _lilypadPaymentEngineAddress, address _lilypadProxyAddress, address _lilypadVestingAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) Initialize(opts *bind.TransactOpts, _l1LilypadTokenAddress common.Address, _l2LilypadTokenAddress common.Address, _lilypadUserAddress common.Address, _lilypadModuleDirectoryAddress common.Address, _lilypadStorageAddress common.Address, _lilypadPaymentEngineAddress common.Address, _lilypadProxyAddress common.Address, _lilypadVestingAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "initialize", _l1LilypadTokenAddress, _l2LilypadTokenAddress, _lilypadUserAddress, _lilypadModuleDirectoryAddress, _lilypadStorageAddress, _lilypadPaymentEngineAddress, _lilypadProxyAddress, _lilypadVestingAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _l1LilypadTokenAddress, address _l2LilypadTokenAddress, address _lilypadUserAddress, address _lilypadModuleDirectoryAddress, address _lilypadStorageAddress, address _lilypadPaymentEngineAddress, address _lilypadProxyAddress, address _lilypadVestingAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) Initialize(_l1LilypadTokenAddress common.Address, _l2LilypadTokenAddress common.Address, _lilypadUserAddress common.Address, _lilypadModuleDirectoryAddress common.Address, _lilypadStorageAddress common.Address, _lilypadPaymentEngineAddress common.Address, _lilypadProxyAddress common.Address, _lilypadVestingAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.Initialize(&_LilypadContractRegistry.TransactOpts, _l1LilypadTokenAddress, _l2LilypadTokenAddress, _lilypadUserAddress, _lilypadModuleDirectoryAddress, _lilypadStorageAddress, _lilypadPaymentEngineAddress, _lilypadProxyAddress, _lilypadVestingAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _l1LilypadTokenAddress, address _l2LilypadTokenAddress, address _lilypadUserAddress, address _lilypadModuleDirectoryAddress, address _lilypadStorageAddress, address _lilypadPaymentEngineAddress, address _lilypadProxyAddress, address _lilypadVestingAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) Initialize(_l1LilypadTokenAddress common.Address, _l2LilypadTokenAddress common.Address, _lilypadUserAddress common.Address, _lilypadModuleDirectoryAddress common.Address, _lilypadStorageAddress common.Address, _lilypadPaymentEngineAddress common.Address, _lilypadProxyAddress common.Address, _lilypadVestingAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.Initialize(&_LilypadContractRegistry.TransactOpts, _l1LilypadTokenAddress, _l2LilypadTokenAddress, _lilypadUserAddress, _lilypadModuleDirectoryAddress, _lilypadStorageAddress, _lilypadPaymentEngineAddress, _lilypadProxyAddress, _lilypadVestingAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.RenounceRole(&_LilypadContractRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.RenounceRole(&_LilypadContractRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.RevokeRole(&_LilypadContractRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.RevokeRole(&_LilypadContractRegistry.TransactOpts, role, account)
}

// SetL2LilypadTokenAddress is a paid mutator transaction binding the contract method 0xe8916cf1.
//
// Solidity: function setL2LilypadTokenAddress(address _l2LilypadTokenAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetL2LilypadTokenAddress(opts *bind.TransactOpts, _l2LilypadTokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setL2LilypadTokenAddress", _l2LilypadTokenAddress)
}

// SetL2LilypadTokenAddress is a paid mutator transaction binding the contract method 0xe8916cf1.
//
// Solidity: function setL2LilypadTokenAddress(address _l2LilypadTokenAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetL2LilypadTokenAddress(_l2LilypadTokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetL2LilypadTokenAddress(&_LilypadContractRegistry.TransactOpts, _l2LilypadTokenAddress)
}

// SetL2LilypadTokenAddress is a paid mutator transaction binding the contract method 0xe8916cf1.
//
// Solidity: function setL2LilypadTokenAddress(address _l2LilypadTokenAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetL2LilypadTokenAddress(_l2LilypadTokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetL2LilypadTokenAddress(&_LilypadContractRegistry.TransactOpts, _l2LilypadTokenAddress)
}

// SetLilypadModuleDirectoryAddress is a paid mutator transaction binding the contract method 0x6879dc92.
//
// Solidity: function setLilypadModuleDirectoryAddress(address _lilypadModuleDirectoryAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetLilypadModuleDirectoryAddress(opts *bind.TransactOpts, _lilypadModuleDirectoryAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setLilypadModuleDirectoryAddress", _lilypadModuleDirectoryAddress)
}

// SetLilypadModuleDirectoryAddress is a paid mutator transaction binding the contract method 0x6879dc92.
//
// Solidity: function setLilypadModuleDirectoryAddress(address _lilypadModuleDirectoryAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetLilypadModuleDirectoryAddress(_lilypadModuleDirectoryAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadModuleDirectoryAddress(&_LilypadContractRegistry.TransactOpts, _lilypadModuleDirectoryAddress)
}

// SetLilypadModuleDirectoryAddress is a paid mutator transaction binding the contract method 0x6879dc92.
//
// Solidity: function setLilypadModuleDirectoryAddress(address _lilypadModuleDirectoryAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetLilypadModuleDirectoryAddress(_lilypadModuleDirectoryAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadModuleDirectoryAddress(&_LilypadContractRegistry.TransactOpts, _lilypadModuleDirectoryAddress)
}

// SetLilypadPaymentEngineAddress is a paid mutator transaction binding the contract method 0x41f0cc40.
//
// Solidity: function setLilypadPaymentEngineAddress(address _lilypadPaymentEngineAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetLilypadPaymentEngineAddress(opts *bind.TransactOpts, _lilypadPaymentEngineAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setLilypadPaymentEngineAddress", _lilypadPaymentEngineAddress)
}

// SetLilypadPaymentEngineAddress is a paid mutator transaction binding the contract method 0x41f0cc40.
//
// Solidity: function setLilypadPaymentEngineAddress(address _lilypadPaymentEngineAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetLilypadPaymentEngineAddress(_lilypadPaymentEngineAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadPaymentEngineAddress(&_LilypadContractRegistry.TransactOpts, _lilypadPaymentEngineAddress)
}

// SetLilypadPaymentEngineAddress is a paid mutator transaction binding the contract method 0x41f0cc40.
//
// Solidity: function setLilypadPaymentEngineAddress(address _lilypadPaymentEngineAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetLilypadPaymentEngineAddress(_lilypadPaymentEngineAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadPaymentEngineAddress(&_LilypadContractRegistry.TransactOpts, _lilypadPaymentEngineAddress)
}

// SetLilypadProxyAddress is a paid mutator transaction binding the contract method 0xac863c08.
//
// Solidity: function setLilypadProxyAddress(address _lilypadProxyAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetLilypadProxyAddress(opts *bind.TransactOpts, _lilypadProxyAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setLilypadProxyAddress", _lilypadProxyAddress)
}

// SetLilypadProxyAddress is a paid mutator transaction binding the contract method 0xac863c08.
//
// Solidity: function setLilypadProxyAddress(address _lilypadProxyAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetLilypadProxyAddress(_lilypadProxyAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadProxyAddress(&_LilypadContractRegistry.TransactOpts, _lilypadProxyAddress)
}

// SetLilypadProxyAddress is a paid mutator transaction binding the contract method 0xac863c08.
//
// Solidity: function setLilypadProxyAddress(address _lilypadProxyAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetLilypadProxyAddress(_lilypadProxyAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadProxyAddress(&_LilypadContractRegistry.TransactOpts, _lilypadProxyAddress)
}

// SetLilypadStorageAddress is a paid mutator transaction binding the contract method 0xfcb15853.
//
// Solidity: function setLilypadStorageAddress(address _lilypadStorageAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetLilypadStorageAddress(opts *bind.TransactOpts, _lilypadStorageAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setLilypadStorageAddress", _lilypadStorageAddress)
}

// SetLilypadStorageAddress is a paid mutator transaction binding the contract method 0xfcb15853.
//
// Solidity: function setLilypadStorageAddress(address _lilypadStorageAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetLilypadStorageAddress(_lilypadStorageAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadStorageAddress(&_LilypadContractRegistry.TransactOpts, _lilypadStorageAddress)
}

// SetLilypadStorageAddress is a paid mutator transaction binding the contract method 0xfcb15853.
//
// Solidity: function setLilypadStorageAddress(address _lilypadStorageAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetLilypadStorageAddress(_lilypadStorageAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadStorageAddress(&_LilypadContractRegistry.TransactOpts, _lilypadStorageAddress)
}

// SetLilypadUserAddress is a paid mutator transaction binding the contract method 0x654e13af.
//
// Solidity: function setLilypadUserAddress(address _lilypadUserAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetLilypadUserAddress(opts *bind.TransactOpts, _lilypadUserAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setLilypadUserAddress", _lilypadUserAddress)
}

// SetLilypadUserAddress is a paid mutator transaction binding the contract method 0x654e13af.
//
// Solidity: function setLilypadUserAddress(address _lilypadUserAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetLilypadUserAddress(_lilypadUserAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadUserAddress(&_LilypadContractRegistry.TransactOpts, _lilypadUserAddress)
}

// SetLilypadUserAddress is a paid mutator transaction binding the contract method 0x654e13af.
//
// Solidity: function setLilypadUserAddress(address _lilypadUserAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetLilypadUserAddress(_lilypadUserAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadUserAddress(&_LilypadContractRegistry.TransactOpts, _lilypadUserAddress)
}

// SetLilypadVestingAddress is a paid mutator transaction binding the contract method 0xcaa442ca.
//
// Solidity: function setLilypadVestingAddress(address _lilypadVestingAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactor) SetLilypadVestingAddress(opts *bind.TransactOpts, _lilypadVestingAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.contract.Transact(opts, "setLilypadVestingAddress", _lilypadVestingAddress)
}

// SetLilypadVestingAddress is a paid mutator transaction binding the contract method 0xcaa442ca.
//
// Solidity: function setLilypadVestingAddress(address _lilypadVestingAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistrySession) SetLilypadVestingAddress(_lilypadVestingAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadVestingAddress(&_LilypadContractRegistry.TransactOpts, _lilypadVestingAddress)
}

// SetLilypadVestingAddress is a paid mutator transaction binding the contract method 0xcaa442ca.
//
// Solidity: function setLilypadVestingAddress(address _lilypadVestingAddress) returns()
func (_LilypadContractRegistry *LilypadContractRegistryTransactorSession) SetLilypadVestingAddress(_lilypadVestingAddress common.Address) (*types.Transaction, error) {
	return _LilypadContractRegistry.Contract.SetLilypadVestingAddress(&_LilypadContractRegistry.TransactOpts, _lilypadVestingAddress)
}

// LilypadContractRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadContractRegistry contract.
type LilypadContractRegistryInitializedIterator struct {
	Event *LilypadContractRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadContractRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadContractRegistryInitialized)
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
		it.Event = new(LilypadContractRegistryInitialized)
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
func (it *LilypadContractRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadContractRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadContractRegistryInitialized represents a Initialized event raised by the LilypadContractRegistry contract.
type LilypadContractRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadContractRegistryInitializedIterator, error) {

	logs, sub, err := _LilypadContractRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryInitializedIterator{contract: _LilypadContractRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadContractRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadContractRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadContractRegistryInitialized)
				if err := _LilypadContractRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) ParseInitialized(log types.Log) (*LilypadContractRegistryInitialized, error) {
	event := new(LilypadContractRegistryInitialized)
	if err := _LilypadContractRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator is returned from FilterLilypadContractRegistryContractAddressSet and is used to iterate over the raw logs and unpacked data for LilypadContractRegistryContractAddressSet events raised by the LilypadContractRegistry contract.
type LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator struct {
	Event *LilypadContractRegistryLilypadContractRegistryContractAddressSet // Event containing the contract specifics and raw log

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
func (it *LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadContractRegistryLilypadContractRegistryContractAddressSet)
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
		it.Event = new(LilypadContractRegistryLilypadContractRegistryContractAddressSet)
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
func (it *LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadContractRegistryLilypadContractRegistryContractAddressSet represents a LilypadContractRegistryContractAddressSet event raised by the LilypadContractRegistry contract.
type LilypadContractRegistryLilypadContractRegistryContractAddressSet struct {
	Name            string
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLilypadContractRegistryContractAddressSet is a free log retrieval operation binding the contract event 0x3c2dc18fa411ade1f2db27ebe21bed1261c284abeda18af5cd036716440c1981.
//
// Solidity: event LilypadContractRegistry__ContractAddressSet(string name, address indexed contractAddress)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) FilterLilypadContractRegistryContractAddressSet(opts *bind.FilterOpts, contractAddress []common.Address) (*LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _LilypadContractRegistry.contract.FilterLogs(opts, "LilypadContractRegistry__ContractAddressSet", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryLilypadContractRegistryContractAddressSetIterator{contract: _LilypadContractRegistry.contract, event: "LilypadContractRegistry__ContractAddressSet", logs: logs, sub: sub}, nil
}

// WatchLilypadContractRegistryContractAddressSet is a free log subscription operation binding the contract event 0x3c2dc18fa411ade1f2db27ebe21bed1261c284abeda18af5cd036716440c1981.
//
// Solidity: event LilypadContractRegistry__ContractAddressSet(string name, address indexed contractAddress)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) WatchLilypadContractRegistryContractAddressSet(opts *bind.WatchOpts, sink chan<- *LilypadContractRegistryLilypadContractRegistryContractAddressSet, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _LilypadContractRegistry.contract.WatchLogs(opts, "LilypadContractRegistry__ContractAddressSet", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadContractRegistryLilypadContractRegistryContractAddressSet)
				if err := _LilypadContractRegistry.contract.UnpackLog(event, "LilypadContractRegistry__ContractAddressSet", log); err != nil {
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

// ParseLilypadContractRegistryContractAddressSet is a log parse operation binding the contract event 0x3c2dc18fa411ade1f2db27ebe21bed1261c284abeda18af5cd036716440c1981.
//
// Solidity: event LilypadContractRegistry__ContractAddressSet(string name, address indexed contractAddress)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) ParseLilypadContractRegistryContractAddressSet(log types.Log) (*LilypadContractRegistryLilypadContractRegistryContractAddressSet, error) {
	event := new(LilypadContractRegistryLilypadContractRegistryContractAddressSet)
	if err := _LilypadContractRegistry.contract.UnpackLog(event, "LilypadContractRegistry__ContractAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadContractRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadContractRegistry contract.
type LilypadContractRegistryRoleAdminChangedIterator struct {
	Event *LilypadContractRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadContractRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadContractRegistryRoleAdminChanged)
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
		it.Event = new(LilypadContractRegistryRoleAdminChanged)
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
func (it *LilypadContractRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadContractRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadContractRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadContractRegistry contract.
type LilypadContractRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadContractRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LilypadContractRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryRoleAdminChangedIterator{contract: _LilypadContractRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadContractRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LilypadContractRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadContractRegistryRoleAdminChanged)
				if err := _LilypadContractRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadContractRegistryRoleAdminChanged, error) {
	event := new(LilypadContractRegistryRoleAdminChanged)
	if err := _LilypadContractRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadContractRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadContractRegistry contract.
type LilypadContractRegistryRoleGrantedIterator struct {
	Event *LilypadContractRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadContractRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadContractRegistryRoleGranted)
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
		it.Event = new(LilypadContractRegistryRoleGranted)
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
func (it *LilypadContractRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadContractRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadContractRegistryRoleGranted represents a RoleGranted event raised by the LilypadContractRegistry contract.
type LilypadContractRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadContractRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _LilypadContractRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryRoleGrantedIterator{contract: _LilypadContractRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadContractRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadContractRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadContractRegistryRoleGranted)
				if err := _LilypadContractRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) ParseRoleGranted(log types.Log) (*LilypadContractRegistryRoleGranted, error) {
	event := new(LilypadContractRegistryRoleGranted)
	if err := _LilypadContractRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadContractRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadContractRegistry contract.
type LilypadContractRegistryRoleRevokedIterator struct {
	Event *LilypadContractRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadContractRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadContractRegistryRoleRevoked)
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
		it.Event = new(LilypadContractRegistryRoleRevoked)
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
func (it *LilypadContractRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadContractRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadContractRegistryRoleRevoked represents a RoleRevoked event raised by the LilypadContractRegistry contract.
type LilypadContractRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadContractRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _LilypadContractRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadContractRegistryRoleRevokedIterator{contract: _LilypadContractRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadContractRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadContractRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadContractRegistryRoleRevoked)
				if err := _LilypadContractRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LilypadContractRegistry *LilypadContractRegistryFilterer) ParseRoleRevoked(log types.Log) (*LilypadContractRegistryRoleRevoked, error) {
	event := new(LilypadContractRegistryRoleRevoked)
	if err := _LilypadContractRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
