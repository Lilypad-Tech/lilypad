// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypadtokenomics

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

// LilypadTokenomicsMetaData contains all meta data concerning the LilypadTokenomics contract.
var LilypadTokenomicsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"m\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"p\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"p1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"p2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"p3\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resourceProviderActiveEscrowScaler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setM\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setP\",\"inputs\":[{\"name\":\"_p\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPvalues\",\"inputs\":[{\"name\":\"_p1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_p2\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_p3\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setResourceProviderActiveEscrowScaler\",\"inputs\":[{\"name\":\"_resourceProviderActiveEscrowScaler\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVValues\",\"inputs\":[{\"name\":\"_v1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_v2\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"v1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"v2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadTokenomics__TokenomicsParameterUpdated\",\"inputs\":[{\"name\":\"parameter\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadTokenomics__MValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadTokenomics__PValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadTokenomics__ParametersMustSumToTenThousand\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadTokenomics__V1MustBeGreaterThanV2\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadTokenomics__ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// LilypadTokenomicsABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadTokenomicsMetaData.ABI instead.
var LilypadTokenomicsABI = LilypadTokenomicsMetaData.ABI

// LilypadTokenomics is an auto generated Go binding around an Ethereum contract.
type LilypadTokenomics struct {
	LilypadTokenomicsCaller     // Read-only binding to the contract
	LilypadTokenomicsTransactor // Write-only binding to the contract
	LilypadTokenomicsFilterer   // Log filterer for contract events
}

// LilypadTokenomicsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadTokenomicsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenomicsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadTokenomicsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenomicsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadTokenomicsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenomicsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadTokenomicsSession struct {
	Contract     *LilypadTokenomics // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LilypadTokenomicsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadTokenomicsCallerSession struct {
	Contract *LilypadTokenomicsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LilypadTokenomicsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadTokenomicsTransactorSession struct {
	Contract     *LilypadTokenomicsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LilypadTokenomicsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadTokenomicsRaw struct {
	Contract *LilypadTokenomics // Generic contract binding to access the raw methods on
}

// LilypadTokenomicsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadTokenomicsCallerRaw struct {
	Contract *LilypadTokenomicsCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadTokenomicsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadTokenomicsTransactorRaw struct {
	Contract *LilypadTokenomicsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadTokenomics creates a new instance of LilypadTokenomics, bound to a specific deployed contract.
func NewLilypadTokenomics(address common.Address, backend bind.ContractBackend) (*LilypadTokenomics, error) {
	contract, err := bindLilypadTokenomics(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomics{LilypadTokenomicsCaller: LilypadTokenomicsCaller{contract: contract}, LilypadTokenomicsTransactor: LilypadTokenomicsTransactor{contract: contract}, LilypadTokenomicsFilterer: LilypadTokenomicsFilterer{contract: contract}}, nil
}

// NewLilypadTokenomicsCaller creates a new read-only instance of LilypadTokenomics, bound to a specific deployed contract.
func NewLilypadTokenomicsCaller(address common.Address, caller bind.ContractCaller) (*LilypadTokenomicsCaller, error) {
	contract, err := bindLilypadTokenomics(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsCaller{contract: contract}, nil
}

// NewLilypadTokenomicsTransactor creates a new write-only instance of LilypadTokenomics, bound to a specific deployed contract.
func NewLilypadTokenomicsTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadTokenomicsTransactor, error) {
	contract, err := bindLilypadTokenomics(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsTransactor{contract: contract}, nil
}

// NewLilypadTokenomicsFilterer creates a new log filterer instance of LilypadTokenomics, bound to a specific deployed contract.
func NewLilypadTokenomicsFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadTokenomicsFilterer, error) {
	contract, err := bindLilypadTokenomics(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsFilterer{contract: contract}, nil
}

// bindLilypadTokenomics binds a generic wrapper to an already deployed contract.
func bindLilypadTokenomics(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadTokenomicsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadTokenomics *LilypadTokenomicsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadTokenomics.Contract.LilypadTokenomicsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadTokenomics *LilypadTokenomicsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.LilypadTokenomicsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadTokenomics *LilypadTokenomicsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.LilypadTokenomicsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadTokenomics *LilypadTokenomicsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadTokenomics.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadTokenomics *LilypadTokenomicsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadTokenomics *LilypadTokenomicsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadTokenomics *LilypadTokenomicsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadTokenomics *LilypadTokenomicsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadTokenomics.Contract.DEFAULTADMINROLE(&_LilypadTokenomics.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadTokenomics.Contract.DEFAULTADMINROLE(&_LilypadTokenomics.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadTokenomics *LilypadTokenomicsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadTokenomics *LilypadTokenomicsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadTokenomics.Contract.GetRoleAdmin(&_LilypadTokenomics.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadTokenomics.Contract.GetRoleAdmin(&_LilypadTokenomics.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadTokenomics *LilypadTokenomicsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadTokenomics *LilypadTokenomicsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadTokenomics.Contract.HasRole(&_LilypadTokenomics.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadTokenomics.Contract.HasRole(&_LilypadTokenomics.CallOpts, role, account)
}

// M is a free data retrieval call binding the contract method 0x5a2ee019.
//
// Solidity: function m() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) M(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "m")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// M is a free data retrieval call binding the contract method 0x5a2ee019.
//
// Solidity: function m() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) M() (*big.Int, error) {
	return _LilypadTokenomics.Contract.M(&_LilypadTokenomics.CallOpts)
}

// M is a free data retrieval call binding the contract method 0x5a2ee019.
//
// Solidity: function m() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) M() (*big.Int, error) {
	return _LilypadTokenomics.Contract.M(&_LilypadTokenomics.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x9ae8886a.
//
// Solidity: function p() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "p")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x9ae8886a.
//
// Solidity: function p() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) P() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P(&_LilypadTokenomics.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x9ae8886a.
//
// Solidity: function p() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) P() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P(&_LilypadTokenomics.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) P1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "p1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) P1() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P1(&_LilypadTokenomics.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) P1() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P1(&_LilypadTokenomics.CallOpts)
}

// P2 is a free data retrieval call binding the contract method 0x81d01ed3.
//
// Solidity: function p2() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) P2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "p2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P2 is a free data retrieval call binding the contract method 0x81d01ed3.
//
// Solidity: function p2() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) P2() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P2(&_LilypadTokenomics.CallOpts)
}

// P2 is a free data retrieval call binding the contract method 0x81d01ed3.
//
// Solidity: function p2() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) P2() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P2(&_LilypadTokenomics.CallOpts)
}

// P3 is a free data retrieval call binding the contract method 0x6e219667.
//
// Solidity: function p3() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) P3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "p3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P3 is a free data retrieval call binding the contract method 0x6e219667.
//
// Solidity: function p3() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) P3() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P3(&_LilypadTokenomics.CallOpts)
}

// P3 is a free data retrieval call binding the contract method 0x6e219667.
//
// Solidity: function p3() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) P3() (*big.Int, error) {
	return _LilypadTokenomics.Contract.P3(&_LilypadTokenomics.CallOpts)
}

// ResourceProviderActiveEscrowScaler is a free data retrieval call binding the contract method 0x368fab37.
//
// Solidity: function resourceProviderActiveEscrowScaler() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) ResourceProviderActiveEscrowScaler(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "resourceProviderActiveEscrowScaler")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResourceProviderActiveEscrowScaler is a free data retrieval call binding the contract method 0x368fab37.
//
// Solidity: function resourceProviderActiveEscrowScaler() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) ResourceProviderActiveEscrowScaler() (*big.Int, error) {
	return _LilypadTokenomics.Contract.ResourceProviderActiveEscrowScaler(&_LilypadTokenomics.CallOpts)
}

// ResourceProviderActiveEscrowScaler is a free data retrieval call binding the contract method 0x368fab37.
//
// Solidity: function resourceProviderActiveEscrowScaler() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) ResourceProviderActiveEscrowScaler() (*big.Int, error) {
	return _LilypadTokenomics.Contract.ResourceProviderActiveEscrowScaler(&_LilypadTokenomics.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadTokenomics *LilypadTokenomicsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadTokenomics *LilypadTokenomicsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadTokenomics.Contract.SupportsInterface(&_LilypadTokenomics.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadTokenomics.Contract.SupportsInterface(&_LilypadTokenomics.CallOpts, interfaceId)
}

// V1 is a free data retrieval call binding the contract method 0x6854171d.
//
// Solidity: function v1() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) V1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "v1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// V1 is a free data retrieval call binding the contract method 0x6854171d.
//
// Solidity: function v1() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) V1() (*big.Int, error) {
	return _LilypadTokenomics.Contract.V1(&_LilypadTokenomics.CallOpts)
}

// V1 is a free data retrieval call binding the contract method 0x6854171d.
//
// Solidity: function v1() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) V1() (*big.Int, error) {
	return _LilypadTokenomics.Contract.V1(&_LilypadTokenomics.CallOpts)
}

// V2 is a free data retrieval call binding the contract method 0xf3acae3a.
//
// Solidity: function v2() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCaller) V2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "v2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// V2 is a free data retrieval call binding the contract method 0xf3acae3a.
//
// Solidity: function v2() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsSession) V2() (*big.Int, error) {
	return _LilypadTokenomics.Contract.V2(&_LilypadTokenomics.CallOpts)
}

// V2 is a free data retrieval call binding the contract method 0xf3acae3a.
//
// Solidity: function v2() view returns(uint256)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) V2() (*big.Int, error) {
	return _LilypadTokenomics.Contract.V2(&_LilypadTokenomics.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadTokenomics *LilypadTokenomicsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadTokenomics.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadTokenomics *LilypadTokenomicsSession) Version() (string, error) {
	return _LilypadTokenomics.Contract.Version(&_LilypadTokenomics.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadTokenomics *LilypadTokenomicsCallerSession) Version() (string, error) {
	return _LilypadTokenomics.Contract.Version(&_LilypadTokenomics.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.GrantRole(&_LilypadTokenomics.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.GrantRole(&_LilypadTokenomics.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) Initialize() (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.Initialize(&_LilypadTokenomics.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) Initialize() (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.Initialize(&_LilypadTokenomics.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.RenounceRole(&_LilypadTokenomics.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.RenounceRole(&_LilypadTokenomics.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.RevokeRole(&_LilypadTokenomics.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.RevokeRole(&_LilypadTokenomics.TransactOpts, role, account)
}

// SetM is a paid mutator transaction binding the contract method 0x54221253.
//
// Solidity: function setM(uint256 _m) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) SetM(opts *bind.TransactOpts, _m *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "setM", _m)
}

// SetM is a paid mutator transaction binding the contract method 0x54221253.
//
// Solidity: function setM(uint256 _m) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) SetM(_m *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetM(&_LilypadTokenomics.TransactOpts, _m)
}

// SetM is a paid mutator transaction binding the contract method 0x54221253.
//
// Solidity: function setM(uint256 _m) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) SetM(_m *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetM(&_LilypadTokenomics.TransactOpts, _m)
}

// SetP is a paid mutator transaction binding the contract method 0xc95473db.
//
// Solidity: function setP(uint256 _p) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) SetP(opts *bind.TransactOpts, _p *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "setP", _p)
}

// SetP is a paid mutator transaction binding the contract method 0xc95473db.
//
// Solidity: function setP(uint256 _p) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) SetP(_p *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetP(&_LilypadTokenomics.TransactOpts, _p)
}

// SetP is a paid mutator transaction binding the contract method 0xc95473db.
//
// Solidity: function setP(uint256 _p) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) SetP(_p *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetP(&_LilypadTokenomics.TransactOpts, _p)
}

// SetPvalues is a paid mutator transaction binding the contract method 0x5d0bc7bb.
//
// Solidity: function setPvalues(uint256 _p1, uint256 _p2, uint256 _p3) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) SetPvalues(opts *bind.TransactOpts, _p1 *big.Int, _p2 *big.Int, _p3 *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "setPvalues", _p1, _p2, _p3)
}

// SetPvalues is a paid mutator transaction binding the contract method 0x5d0bc7bb.
//
// Solidity: function setPvalues(uint256 _p1, uint256 _p2, uint256 _p3) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) SetPvalues(_p1 *big.Int, _p2 *big.Int, _p3 *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetPvalues(&_LilypadTokenomics.TransactOpts, _p1, _p2, _p3)
}

// SetPvalues is a paid mutator transaction binding the contract method 0x5d0bc7bb.
//
// Solidity: function setPvalues(uint256 _p1, uint256 _p2, uint256 _p3) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) SetPvalues(_p1 *big.Int, _p2 *big.Int, _p3 *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetPvalues(&_LilypadTokenomics.TransactOpts, _p1, _p2, _p3)
}

// SetResourceProviderActiveEscrowScaler is a paid mutator transaction binding the contract method 0x91340a16.
//
// Solidity: function setResourceProviderActiveEscrowScaler(uint256 _resourceProviderActiveEscrowScaler) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) SetResourceProviderActiveEscrowScaler(opts *bind.TransactOpts, _resourceProviderActiveEscrowScaler *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "setResourceProviderActiveEscrowScaler", _resourceProviderActiveEscrowScaler)
}

// SetResourceProviderActiveEscrowScaler is a paid mutator transaction binding the contract method 0x91340a16.
//
// Solidity: function setResourceProviderActiveEscrowScaler(uint256 _resourceProviderActiveEscrowScaler) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) SetResourceProviderActiveEscrowScaler(_resourceProviderActiveEscrowScaler *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetResourceProviderActiveEscrowScaler(&_LilypadTokenomics.TransactOpts, _resourceProviderActiveEscrowScaler)
}

// SetResourceProviderActiveEscrowScaler is a paid mutator transaction binding the contract method 0x91340a16.
//
// Solidity: function setResourceProviderActiveEscrowScaler(uint256 _resourceProviderActiveEscrowScaler) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) SetResourceProviderActiveEscrowScaler(_resourceProviderActiveEscrowScaler *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetResourceProviderActiveEscrowScaler(&_LilypadTokenomics.TransactOpts, _resourceProviderActiveEscrowScaler)
}

// SetVValues is a paid mutator transaction binding the contract method 0xa0608c8a.
//
// Solidity: function setVValues(uint256 _v1, uint256 _v2) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactor) SetVValues(opts *bind.TransactOpts, _v1 *big.Int, _v2 *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.contract.Transact(opts, "setVValues", _v1, _v2)
}

// SetVValues is a paid mutator transaction binding the contract method 0xa0608c8a.
//
// Solidity: function setVValues(uint256 _v1, uint256 _v2) returns()
func (_LilypadTokenomics *LilypadTokenomicsSession) SetVValues(_v1 *big.Int, _v2 *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetVValues(&_LilypadTokenomics.TransactOpts, _v1, _v2)
}

// SetVValues is a paid mutator transaction binding the contract method 0xa0608c8a.
//
// Solidity: function setVValues(uint256 _v1, uint256 _v2) returns()
func (_LilypadTokenomics *LilypadTokenomicsTransactorSession) SetVValues(_v1 *big.Int, _v2 *big.Int) (*types.Transaction, error) {
	return _LilypadTokenomics.Contract.SetVValues(&_LilypadTokenomics.TransactOpts, _v1, _v2)
}

// LilypadTokenomicsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadTokenomics contract.
type LilypadTokenomicsInitializedIterator struct {
	Event *LilypadTokenomicsInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadTokenomicsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenomicsInitialized)
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
		it.Event = new(LilypadTokenomicsInitialized)
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
func (it *LilypadTokenomicsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenomicsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenomicsInitialized represents a Initialized event raised by the LilypadTokenomics contract.
type LilypadTokenomicsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadTokenomicsInitializedIterator, error) {

	logs, sub, err := _LilypadTokenomics.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsInitializedIterator{contract: _LilypadTokenomics.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadTokenomicsInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadTokenomics.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenomicsInitialized)
				if err := _LilypadTokenomics.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadTokenomics *LilypadTokenomicsFilterer) ParseInitialized(log types.Log) (*LilypadTokenomicsInitialized, error) {
	event := new(LilypadTokenomicsInitialized)
	if err := _LilypadTokenomics.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator is returned from FilterLilypadTokenomicsTokenomicsParameterUpdated and is used to iterate over the raw logs and unpacked data for LilypadTokenomicsTokenomicsParameterUpdated events raised by the LilypadTokenomics contract.
type LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator struct {
	Event *LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated)
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
		it.Event = new(LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated)
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
func (it *LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated represents a LilypadTokenomicsTokenomicsParameterUpdated event raised by the LilypadTokenomics contract.
type LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated struct {
	Parameter common.Hash
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLilypadTokenomicsTokenomicsParameterUpdated is a free log retrieval operation binding the contract event 0x4082466583d48b5fa1fa523fe6bbadb8e176575003edc907006aa227806f6114.
//
// Solidity: event LilypadTokenomics__TokenomicsParameterUpdated(string indexed parameter, uint256 value)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) FilterLilypadTokenomicsTokenomicsParameterUpdated(opts *bind.FilterOpts, parameter []string) (*LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _LilypadTokenomics.contract.FilterLogs(opts, "LilypadTokenomics__TokenomicsParameterUpdated", parameterRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdatedIterator{contract: _LilypadTokenomics.contract, event: "LilypadTokenomics__TokenomicsParameterUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadTokenomicsTokenomicsParameterUpdated is a free log subscription operation binding the contract event 0x4082466583d48b5fa1fa523fe6bbadb8e176575003edc907006aa227806f6114.
//
// Solidity: event LilypadTokenomics__TokenomicsParameterUpdated(string indexed parameter, uint256 value)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) WatchLilypadTokenomicsTokenomicsParameterUpdated(opts *bind.WatchOpts, sink chan<- *LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated, parameter []string) (event.Subscription, error) {

	var parameterRule []interface{}
	for _, parameterItem := range parameter {
		parameterRule = append(parameterRule, parameterItem)
	}

	logs, sub, err := _LilypadTokenomics.contract.WatchLogs(opts, "LilypadTokenomics__TokenomicsParameterUpdated", parameterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated)
				if err := _LilypadTokenomics.contract.UnpackLog(event, "LilypadTokenomics__TokenomicsParameterUpdated", log); err != nil {
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

// ParseLilypadTokenomicsTokenomicsParameterUpdated is a log parse operation binding the contract event 0x4082466583d48b5fa1fa523fe6bbadb8e176575003edc907006aa227806f6114.
//
// Solidity: event LilypadTokenomics__TokenomicsParameterUpdated(string indexed parameter, uint256 value)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) ParseLilypadTokenomicsTokenomicsParameterUpdated(log types.Log) (*LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated, error) {
	event := new(LilypadTokenomicsLilypadTokenomicsTokenomicsParameterUpdated)
	if err := _LilypadTokenomics.contract.UnpackLog(event, "LilypadTokenomics__TokenomicsParameterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenomicsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadTokenomics contract.
type LilypadTokenomicsRoleAdminChangedIterator struct {
	Event *LilypadTokenomicsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadTokenomicsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenomicsRoleAdminChanged)
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
		it.Event = new(LilypadTokenomicsRoleAdminChanged)
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
func (it *LilypadTokenomicsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenomicsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenomicsRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadTokenomics contract.
type LilypadTokenomicsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadTokenomicsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LilypadTokenomics.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsRoleAdminChangedIterator{contract: _LilypadTokenomics.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadTokenomicsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LilypadTokenomics.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenomicsRoleAdminChanged)
				if err := _LilypadTokenomics.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LilypadTokenomics *LilypadTokenomicsFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadTokenomicsRoleAdminChanged, error) {
	event := new(LilypadTokenomicsRoleAdminChanged)
	if err := _LilypadTokenomics.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenomicsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadTokenomics contract.
type LilypadTokenomicsRoleGrantedIterator struct {
	Event *LilypadTokenomicsRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadTokenomicsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenomicsRoleGranted)
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
		it.Event = new(LilypadTokenomicsRoleGranted)
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
func (it *LilypadTokenomicsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenomicsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenomicsRoleGranted represents a RoleGranted event raised by the LilypadTokenomics contract.
type LilypadTokenomicsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadTokenomicsRoleGrantedIterator, error) {

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

	logs, sub, err := _LilypadTokenomics.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsRoleGrantedIterator{contract: _LilypadTokenomics.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadTokenomicsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadTokenomics.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenomicsRoleGranted)
				if err := _LilypadTokenomics.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LilypadTokenomics *LilypadTokenomicsFilterer) ParseRoleGranted(log types.Log) (*LilypadTokenomicsRoleGranted, error) {
	event := new(LilypadTokenomicsRoleGranted)
	if err := _LilypadTokenomics.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenomicsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadTokenomics contract.
type LilypadTokenomicsRoleRevokedIterator struct {
	Event *LilypadTokenomicsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadTokenomicsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenomicsRoleRevoked)
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
		it.Event = new(LilypadTokenomicsRoleRevoked)
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
func (it *LilypadTokenomicsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenomicsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenomicsRoleRevoked represents a RoleRevoked event raised by the LilypadTokenomics contract.
type LilypadTokenomicsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadTokenomicsRoleRevokedIterator, error) {

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

	logs, sub, err := _LilypadTokenomics.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenomicsRoleRevokedIterator{contract: _LilypadTokenomics.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadTokenomics *LilypadTokenomicsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadTokenomicsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadTokenomics.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenomicsRoleRevoked)
				if err := _LilypadTokenomics.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LilypadTokenomics *LilypadTokenomicsFilterer) ParseRoleRevoked(log types.Log) (*LilypadTokenomicsRoleRevoked, error) {
	event := new(LilypadTokenomicsRoleRevoked)
	if err := _LilypadTokenomics.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
