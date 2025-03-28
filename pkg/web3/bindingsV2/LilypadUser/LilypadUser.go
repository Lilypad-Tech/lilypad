// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypaduser

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

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress common.Address
	MetadataID  string
	Url         string
}

// LilypadUserMetaData contains all meta data concerning the LilypadUser contract.
var LilypadUserMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addRole\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.UserType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getControllerAccessControlRole\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getMinterAccessControlRole\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPauserAccessControlRole\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUser\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.User\",\"components\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVestingAccessControlRole\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.UserType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"insertUser\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.UserType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRole\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.UserType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateUserMetadata\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadUser__UserManagementEvent\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataID\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"role\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSharedStructs.UserType\"},{\"name\":\"operation\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSharedStructs.UserOperation\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadUser__RoleAlreadyAssigned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadUser__RoleNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadUser__RoleNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadUser__UserAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadUser__UserNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// LilypadUserABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadUserMetaData.ABI instead.
var LilypadUserABI = LilypadUserMetaData.ABI

// LilypadUser is an auto generated Go binding around an Ethereum contract.
type LilypadUser struct {
	LilypadUserCaller     // Read-only binding to the contract
	LilypadUserTransactor // Write-only binding to the contract
	LilypadUserFilterer   // Log filterer for contract events
}

// LilypadUserCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadUserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadUserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadUserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadUserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadUserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadUserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadUserSession struct {
	Contract     *LilypadUser      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LilypadUserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadUserCallerSession struct {
	Contract *LilypadUserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LilypadUserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadUserTransactorSession struct {
	Contract     *LilypadUserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LilypadUserRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadUserRaw struct {
	Contract *LilypadUser // Generic contract binding to access the raw methods on
}

// LilypadUserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadUserCallerRaw struct {
	Contract *LilypadUserCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadUserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadUserTransactorRaw struct {
	Contract *LilypadUserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadUser creates a new instance of LilypadUser, bound to a specific deployed contract.
func NewLilypadUser(address common.Address, backend bind.ContractBackend) (*LilypadUser, error) {
	contract, err := bindLilypadUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadUser{LilypadUserCaller: LilypadUserCaller{contract: contract}, LilypadUserTransactor: LilypadUserTransactor{contract: contract}, LilypadUserFilterer: LilypadUserFilterer{contract: contract}}, nil
}

// NewLilypadUserCaller creates a new read-only instance of LilypadUser, bound to a specific deployed contract.
func NewLilypadUserCaller(address common.Address, caller bind.ContractCaller) (*LilypadUserCaller, error) {
	contract, err := bindLilypadUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadUserCaller{contract: contract}, nil
}

// NewLilypadUserTransactor creates a new write-only instance of LilypadUser, bound to a specific deployed contract.
func NewLilypadUserTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadUserTransactor, error) {
	contract, err := bindLilypadUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadUserTransactor{contract: contract}, nil
}

// NewLilypadUserFilterer creates a new log filterer instance of LilypadUser, bound to a specific deployed contract.
func NewLilypadUserFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadUserFilterer, error) {
	contract, err := bindLilypadUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadUserFilterer{contract: contract}, nil
}

// bindLilypadUser binds a generic wrapper to an already deployed contract.
func bindLilypadUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadUserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadUser *LilypadUserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadUser.Contract.LilypadUserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadUser *LilypadUserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadUser.Contract.LilypadUserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadUser *LilypadUserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadUser.Contract.LilypadUserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadUser *LilypadUserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadUser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadUser *LilypadUserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadUser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadUser *LilypadUserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadUser.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadUser *LilypadUserCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadUser *LilypadUserSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadUser.Contract.DEFAULTADMINROLE(&_LilypadUser.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadUser *LilypadUserCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadUser.Contract.DEFAULTADMINROLE(&_LilypadUser.CallOpts)
}

// GetControllerAccessControlRole is a free data retrieval call binding the contract method 0xd2202f83.
//
// Solidity: function getControllerAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCaller) GetControllerAccessControlRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getControllerAccessControlRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetControllerAccessControlRole is a free data retrieval call binding the contract method 0xd2202f83.
//
// Solidity: function getControllerAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserSession) GetControllerAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetControllerAccessControlRole(&_LilypadUser.CallOpts)
}

// GetControllerAccessControlRole is a free data retrieval call binding the contract method 0xd2202f83.
//
// Solidity: function getControllerAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCallerSession) GetControllerAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetControllerAccessControlRole(&_LilypadUser.CallOpts)
}

// GetMinterAccessControlRole is a free data retrieval call binding the contract method 0xaa4a20b0.
//
// Solidity: function getMinterAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCaller) GetMinterAccessControlRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getMinterAccessControlRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMinterAccessControlRole is a free data retrieval call binding the contract method 0xaa4a20b0.
//
// Solidity: function getMinterAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserSession) GetMinterAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetMinterAccessControlRole(&_LilypadUser.CallOpts)
}

// GetMinterAccessControlRole is a free data retrieval call binding the contract method 0xaa4a20b0.
//
// Solidity: function getMinterAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCallerSession) GetMinterAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetMinterAccessControlRole(&_LilypadUser.CallOpts)
}

// GetPauserAccessControlRole is a free data retrieval call binding the contract method 0xf7248fa5.
//
// Solidity: function getPauserAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCaller) GetPauserAccessControlRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getPauserAccessControlRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPauserAccessControlRole is a free data retrieval call binding the contract method 0xf7248fa5.
//
// Solidity: function getPauserAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserSession) GetPauserAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetPauserAccessControlRole(&_LilypadUser.CallOpts)
}

// GetPauserAccessControlRole is a free data retrieval call binding the contract method 0xf7248fa5.
//
// Solidity: function getPauserAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCallerSession) GetPauserAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetPauserAccessControlRole(&_LilypadUser.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadUser *LilypadUserCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadUser *LilypadUserSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadUser.Contract.GetRoleAdmin(&_LilypadUser.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadUser *LilypadUserCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadUser.Contract.GetRoleAdmin(&_LilypadUser.CallOpts, role)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address walletAddress) view returns((address,string,string))
func (_LilypadUser *LilypadUserCaller) GetUser(opts *bind.CallOpts, walletAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getUser", walletAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address walletAddress) view returns((address,string,string))
func (_LilypadUser *LilypadUserSession) GetUser(walletAddress common.Address) (SharedStructsUser, error) {
	return _LilypadUser.Contract.GetUser(&_LilypadUser.CallOpts, walletAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address walletAddress) view returns((address,string,string))
func (_LilypadUser *LilypadUserCallerSession) GetUser(walletAddress common.Address) (SharedStructsUser, error) {
	return _LilypadUser.Contract.GetUser(&_LilypadUser.CallOpts, walletAddress)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_LilypadUser *LilypadUserCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_LilypadUser *LilypadUserSession) GetValidators() ([]common.Address, error) {
	return _LilypadUser.Contract.GetValidators(&_LilypadUser.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_LilypadUser *LilypadUserCallerSession) GetValidators() ([]common.Address, error) {
	return _LilypadUser.Contract.GetValidators(&_LilypadUser.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadUser *LilypadUserCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadUser *LilypadUserSession) GetVersion() (string, error) {
	return _LilypadUser.Contract.GetVersion(&_LilypadUser.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadUser *LilypadUserCallerSession) GetVersion() (string, error) {
	return _LilypadUser.Contract.GetVersion(&_LilypadUser.CallOpts)
}

// GetVestingAccessControlRole is a free data retrieval call binding the contract method 0x851ee3e5.
//
// Solidity: function getVestingAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCaller) GetVestingAccessControlRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "getVestingAccessControlRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVestingAccessControlRole is a free data retrieval call binding the contract method 0x851ee3e5.
//
// Solidity: function getVestingAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserSession) GetVestingAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetVestingAccessControlRole(&_LilypadUser.CallOpts)
}

// GetVestingAccessControlRole is a free data retrieval call binding the contract method 0x851ee3e5.
//
// Solidity: function getVestingAccessControlRole() pure returns(bytes32)
func (_LilypadUser *LilypadUserCallerSession) GetVestingAccessControlRole() ([32]byte, error) {
	return _LilypadUser.Contract.GetVestingAccessControlRole(&_LilypadUser.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadUser *LilypadUserCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadUser *LilypadUserSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadUser.Contract.HasRole(&_LilypadUser.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadUser *LilypadUserCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadUser.Contract.HasRole(&_LilypadUser.CallOpts, role, account)
}

// HasRole0 is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address walletAddress, uint8 role) view returns(bool)
func (_LilypadUser *LilypadUserCaller) HasRole0(opts *bind.CallOpts, walletAddress common.Address, role uint8) (bool, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "hasRole0", walletAddress, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole0 is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address walletAddress, uint8 role) view returns(bool)
func (_LilypadUser *LilypadUserSession) HasRole0(walletAddress common.Address, role uint8) (bool, error) {
	return _LilypadUser.Contract.HasRole0(&_LilypadUser.CallOpts, walletAddress, role)
}

// HasRole0 is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address walletAddress, uint8 role) view returns(bool)
func (_LilypadUser *LilypadUserCallerSession) HasRole0(walletAddress common.Address, role uint8) (bool, error) {
	return _LilypadUser.Contract.HasRole0(&_LilypadUser.CallOpts, walletAddress, role)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadUser *LilypadUserCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadUser *LilypadUserSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadUser.Contract.SupportsInterface(&_LilypadUser.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadUser *LilypadUserCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadUser.Contract.SupportsInterface(&_LilypadUser.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadUser *LilypadUserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadUser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadUser *LilypadUserSession) Version() (string, error) {
	return _LilypadUser.Contract.Version(&_LilypadUser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadUser *LilypadUserCallerSession) Version() (string, error) {
	return _LilypadUser.Contract.Version(&_LilypadUser.CallOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address walletAddress, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserTransactor) AddRole(opts *bind.TransactOpts, walletAddress common.Address, role uint8) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "addRole", walletAddress, role)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address walletAddress, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserSession) AddRole(walletAddress common.Address, role uint8) (*types.Transaction, error) {
	return _LilypadUser.Contract.AddRole(&_LilypadUser.TransactOpts, walletAddress, role)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address walletAddress, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserTransactorSession) AddRole(walletAddress common.Address, role uint8) (*types.Transaction, error) {
	return _LilypadUser.Contract.AddRole(&_LilypadUser.TransactOpts, walletAddress, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadUser *LilypadUserTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadUser *LilypadUserSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadUser.Contract.GrantRole(&_LilypadUser.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadUser *LilypadUserTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadUser.Contract.GrantRole(&_LilypadUser.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadUser *LilypadUserTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadUser *LilypadUserSession) Initialize() (*types.Transaction, error) {
	return _LilypadUser.Contract.Initialize(&_LilypadUser.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadUser *LilypadUserTransactorSession) Initialize() (*types.Transaction, error) {
	return _LilypadUser.Contract.Initialize(&_LilypadUser.TransactOpts)
}

// InsertUser is a paid mutator transaction binding the contract method 0xd8eaafd3.
//
// Solidity: function insertUser(address walletAddress, string metadataID, string url, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserTransactor) InsertUser(opts *bind.TransactOpts, walletAddress common.Address, metadataID string, url string, role uint8) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "insertUser", walletAddress, metadataID, url, role)
}

// InsertUser is a paid mutator transaction binding the contract method 0xd8eaafd3.
//
// Solidity: function insertUser(address walletAddress, string metadataID, string url, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserSession) InsertUser(walletAddress common.Address, metadataID string, url string, role uint8) (*types.Transaction, error) {
	return _LilypadUser.Contract.InsertUser(&_LilypadUser.TransactOpts, walletAddress, metadataID, url, role)
}

// InsertUser is a paid mutator transaction binding the contract method 0xd8eaafd3.
//
// Solidity: function insertUser(address walletAddress, string metadataID, string url, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserTransactorSession) InsertUser(walletAddress common.Address, metadataID string, url string, role uint8) (*types.Transaction, error) {
	return _LilypadUser.Contract.InsertUser(&_LilypadUser.TransactOpts, walletAddress, metadataID, url, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address walletAddress, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserTransactor) RemoveRole(opts *bind.TransactOpts, walletAddress common.Address, role uint8) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "removeRole", walletAddress, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address walletAddress, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserSession) RemoveRole(walletAddress common.Address, role uint8) (*types.Transaction, error) {
	return _LilypadUser.Contract.RemoveRole(&_LilypadUser.TransactOpts, walletAddress, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address walletAddress, uint8 role) returns(bool)
func (_LilypadUser *LilypadUserTransactorSession) RemoveRole(walletAddress common.Address, role uint8) (*types.Transaction, error) {
	return _LilypadUser.Contract.RemoveRole(&_LilypadUser.TransactOpts, walletAddress, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadUser *LilypadUserTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadUser *LilypadUserSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadUser.Contract.RenounceRole(&_LilypadUser.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadUser *LilypadUserTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadUser.Contract.RenounceRole(&_LilypadUser.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadUser *LilypadUserTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadUser *LilypadUserSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadUser.Contract.RevokeRole(&_LilypadUser.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadUser *LilypadUserTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadUser.Contract.RevokeRole(&_LilypadUser.TransactOpts, role, account)
}

// UpdateUserMetadata is a paid mutator transaction binding the contract method 0xd5d1c153.
//
// Solidity: function updateUserMetadata(address walletAddress, string metadataID, string url) returns(bool)
func (_LilypadUser *LilypadUserTransactor) UpdateUserMetadata(opts *bind.TransactOpts, walletAddress common.Address, metadataID string, url string) (*types.Transaction, error) {
	return _LilypadUser.contract.Transact(opts, "updateUserMetadata", walletAddress, metadataID, url)
}

// UpdateUserMetadata is a paid mutator transaction binding the contract method 0xd5d1c153.
//
// Solidity: function updateUserMetadata(address walletAddress, string metadataID, string url) returns(bool)
func (_LilypadUser *LilypadUserSession) UpdateUserMetadata(walletAddress common.Address, metadataID string, url string) (*types.Transaction, error) {
	return _LilypadUser.Contract.UpdateUserMetadata(&_LilypadUser.TransactOpts, walletAddress, metadataID, url)
}

// UpdateUserMetadata is a paid mutator transaction binding the contract method 0xd5d1c153.
//
// Solidity: function updateUserMetadata(address walletAddress, string metadataID, string url) returns(bool)
func (_LilypadUser *LilypadUserTransactorSession) UpdateUserMetadata(walletAddress common.Address, metadataID string, url string) (*types.Transaction, error) {
	return _LilypadUser.Contract.UpdateUserMetadata(&_LilypadUser.TransactOpts, walletAddress, metadataID, url)
}

// LilypadUserInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadUser contract.
type LilypadUserInitializedIterator struct {
	Event *LilypadUserInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadUserInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadUserInitialized)
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
		it.Event = new(LilypadUserInitialized)
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
func (it *LilypadUserInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadUserInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadUserInitialized represents a Initialized event raised by the LilypadUser contract.
type LilypadUserInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadUser *LilypadUserFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadUserInitializedIterator, error) {

	logs, sub, err := _LilypadUser.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadUserInitializedIterator{contract: _LilypadUser.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadUser *LilypadUserFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadUserInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadUser.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadUserInitialized)
				if err := _LilypadUser.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadUser *LilypadUserFilterer) ParseInitialized(log types.Log) (*LilypadUserInitialized, error) {
	event := new(LilypadUserInitialized)
	if err := _LilypadUser.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadUserLilypadUserUserManagementEventIterator is returned from FilterLilypadUserUserManagementEvent and is used to iterate over the raw logs and unpacked data for LilypadUserUserManagementEvent events raised by the LilypadUser contract.
type LilypadUserLilypadUserUserManagementEventIterator struct {
	Event *LilypadUserLilypadUserUserManagementEvent // Event containing the contract specifics and raw log

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
func (it *LilypadUserLilypadUserUserManagementEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadUserLilypadUserUserManagementEvent)
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
		it.Event = new(LilypadUserLilypadUserUserManagementEvent)
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
func (it *LilypadUserLilypadUserUserManagementEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadUserLilypadUserUserManagementEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadUserLilypadUserUserManagementEvent represents a LilypadUserUserManagementEvent event raised by the LilypadUser contract.
type LilypadUserLilypadUserUserManagementEvent struct {
	WalletAddress common.Address
	MetadataID    string
	Url           string
	Role          uint8
	Operation     uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLilypadUserUserManagementEvent is a free log retrieval operation binding the contract event 0x0aa890b2c73b3b643a7048ee9fb3dba2c873a74f489fbba5bc8a71a46b74d933.
//
// Solidity: event LilypadUser__UserManagementEvent(address indexed walletAddress, string metadataID, string url, uint8 role, uint8 operation)
func (_LilypadUser *LilypadUserFilterer) FilterLilypadUserUserManagementEvent(opts *bind.FilterOpts, walletAddress []common.Address) (*LilypadUserLilypadUserUserManagementEventIterator, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}

	logs, sub, err := _LilypadUser.contract.FilterLogs(opts, "LilypadUser__UserManagementEvent", walletAddressRule)
	if err != nil {
		return nil, err
	}
	return &LilypadUserLilypadUserUserManagementEventIterator{contract: _LilypadUser.contract, event: "LilypadUser__UserManagementEvent", logs: logs, sub: sub}, nil
}

// WatchLilypadUserUserManagementEvent is a free log subscription operation binding the contract event 0x0aa890b2c73b3b643a7048ee9fb3dba2c873a74f489fbba5bc8a71a46b74d933.
//
// Solidity: event LilypadUser__UserManagementEvent(address indexed walletAddress, string metadataID, string url, uint8 role, uint8 operation)
func (_LilypadUser *LilypadUserFilterer) WatchLilypadUserUserManagementEvent(opts *bind.WatchOpts, sink chan<- *LilypadUserLilypadUserUserManagementEvent, walletAddress []common.Address) (event.Subscription, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}

	logs, sub, err := _LilypadUser.contract.WatchLogs(opts, "LilypadUser__UserManagementEvent", walletAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadUserLilypadUserUserManagementEvent)
				if err := _LilypadUser.contract.UnpackLog(event, "LilypadUser__UserManagementEvent", log); err != nil {
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

// ParseLilypadUserUserManagementEvent is a log parse operation binding the contract event 0x0aa890b2c73b3b643a7048ee9fb3dba2c873a74f489fbba5bc8a71a46b74d933.
//
// Solidity: event LilypadUser__UserManagementEvent(address indexed walletAddress, string metadataID, string url, uint8 role, uint8 operation)
func (_LilypadUser *LilypadUserFilterer) ParseLilypadUserUserManagementEvent(log types.Log) (*LilypadUserLilypadUserUserManagementEvent, error) {
	event := new(LilypadUserLilypadUserUserManagementEvent)
	if err := _LilypadUser.contract.UnpackLog(event, "LilypadUser__UserManagementEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadUserRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadUser contract.
type LilypadUserRoleAdminChangedIterator struct {
	Event *LilypadUserRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadUserRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadUserRoleAdminChanged)
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
		it.Event = new(LilypadUserRoleAdminChanged)
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
func (it *LilypadUserRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadUserRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadUserRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadUser contract.
type LilypadUserRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadUser *LilypadUserFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadUserRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LilypadUser.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadUserRoleAdminChangedIterator{contract: _LilypadUser.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadUser *LilypadUserFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadUserRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LilypadUser.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadUserRoleAdminChanged)
				if err := _LilypadUser.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LilypadUser *LilypadUserFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadUserRoleAdminChanged, error) {
	event := new(LilypadUserRoleAdminChanged)
	if err := _LilypadUser.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadUserRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadUser contract.
type LilypadUserRoleGrantedIterator struct {
	Event *LilypadUserRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadUserRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadUserRoleGranted)
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
		it.Event = new(LilypadUserRoleGranted)
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
func (it *LilypadUserRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadUserRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadUserRoleGranted represents a RoleGranted event raised by the LilypadUser contract.
type LilypadUserRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadUser *LilypadUserFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadUserRoleGrantedIterator, error) {

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

	logs, sub, err := _LilypadUser.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadUserRoleGrantedIterator{contract: _LilypadUser.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadUser *LilypadUserFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadUserRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadUser.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadUserRoleGranted)
				if err := _LilypadUser.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LilypadUser *LilypadUserFilterer) ParseRoleGranted(log types.Log) (*LilypadUserRoleGranted, error) {
	event := new(LilypadUserRoleGranted)
	if err := _LilypadUser.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadUserRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadUser contract.
type LilypadUserRoleRevokedIterator struct {
	Event *LilypadUserRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadUserRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadUserRoleRevoked)
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
		it.Event = new(LilypadUserRoleRevoked)
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
func (it *LilypadUserRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadUserRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadUserRoleRevoked represents a RoleRevoked event raised by the LilypadUser contract.
type LilypadUserRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadUser *LilypadUserFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadUserRoleRevokedIterator, error) {

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

	logs, sub, err := _LilypadUser.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadUserRoleRevokedIterator{contract: _LilypadUser.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadUser *LilypadUserFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadUserRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadUser.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadUserRoleRevoked)
				if err := _LilypadUser.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LilypadUser *LilypadUserFilterer) ParseRoleRevoked(log types.Log) (*LilypadUserRoleRevoked, error) {
	event := new(LilypadUserRoleRevoked)
	if err := _LilypadUser.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
