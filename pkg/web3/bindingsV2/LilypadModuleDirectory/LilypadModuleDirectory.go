// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypadmoduledirectory

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

// SharedStructsModule is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsModule struct {
	ModuleOwner common.Address
	ModuleName  string
	ModuleUrl   string
}

// LilypadModuleDirectoryMetaData contains all meta data concerning the LilypadModuleDirectory contract.
var LilypadModuleDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveTransfer\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLilypadUser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOwnedModules\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structSharedStructs.Module[]\",\"components\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_lilypadUser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTransferApproved\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"purchaser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerModuleCreator\",\"inputs\":[{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerModuleForCreator\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeTransferApproval\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadUser\",\"inputs\":[{\"name\":\"_lilypadUser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferModuleOwnership\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateModuleName\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"newModuleName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateModuleUrl\",\"inputs\":[{\"name\":\"moduleOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"newModuleUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__LilypadUserUpdated\",\"inputs\":[{\"name\":\"lilypadUser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleCreatorRegistered\",\"inputs\":[{\"name\":\"moduleCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleNameUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldModuleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newModuleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleRegistered\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleTransferApproved\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"purchaser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleTransferRevoked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"revokedFrom\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleTransferred\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"moduleUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadModuleDirectory__ModuleUrlUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"moduleName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newModuleUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__EmptyModuleName\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__EmptyModuleUrl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__InvalidAddressForLilypadUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__ModuleAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__ModuleCreatorAlreadyExists\",\"inputs\":[{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__ModuleNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__NotController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__NotModuleOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__SameOwnerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__TransferNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadModuleDirectory__ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// LilypadModuleDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadModuleDirectoryMetaData.ABI instead.
var LilypadModuleDirectoryABI = LilypadModuleDirectoryMetaData.ABI

// LilypadModuleDirectory is an auto generated Go binding around an Ethereum contract.
type LilypadModuleDirectory struct {
	LilypadModuleDirectoryCaller     // Read-only binding to the contract
	LilypadModuleDirectoryTransactor // Write-only binding to the contract
	LilypadModuleDirectoryFilterer   // Log filterer for contract events
}

// LilypadModuleDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadModuleDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadModuleDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadModuleDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadModuleDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadModuleDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadModuleDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadModuleDirectorySession struct {
	Contract     *LilypadModuleDirectory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LilypadModuleDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadModuleDirectoryCallerSession struct {
	Contract *LilypadModuleDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// LilypadModuleDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadModuleDirectoryTransactorSession struct {
	Contract     *LilypadModuleDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// LilypadModuleDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadModuleDirectoryRaw struct {
	Contract *LilypadModuleDirectory // Generic contract binding to access the raw methods on
}

// LilypadModuleDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadModuleDirectoryCallerRaw struct {
	Contract *LilypadModuleDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadModuleDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadModuleDirectoryTransactorRaw struct {
	Contract *LilypadModuleDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadModuleDirectory creates a new instance of LilypadModuleDirectory, bound to a specific deployed contract.
func NewLilypadModuleDirectory(address common.Address, backend bind.ContractBackend) (*LilypadModuleDirectory, error) {
	contract, err := bindLilypadModuleDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectory{LilypadModuleDirectoryCaller: LilypadModuleDirectoryCaller{contract: contract}, LilypadModuleDirectoryTransactor: LilypadModuleDirectoryTransactor{contract: contract}, LilypadModuleDirectoryFilterer: LilypadModuleDirectoryFilterer{contract: contract}}, nil
}

// NewLilypadModuleDirectoryCaller creates a new read-only instance of LilypadModuleDirectory, bound to a specific deployed contract.
func NewLilypadModuleDirectoryCaller(address common.Address, caller bind.ContractCaller) (*LilypadModuleDirectoryCaller, error) {
	contract, err := bindLilypadModuleDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryCaller{contract: contract}, nil
}

// NewLilypadModuleDirectoryTransactor creates a new write-only instance of LilypadModuleDirectory, bound to a specific deployed contract.
func NewLilypadModuleDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadModuleDirectoryTransactor, error) {
	contract, err := bindLilypadModuleDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryTransactor{contract: contract}, nil
}

// NewLilypadModuleDirectoryFilterer creates a new log filterer instance of LilypadModuleDirectory, bound to a specific deployed contract.
func NewLilypadModuleDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadModuleDirectoryFilterer, error) {
	contract, err := bindLilypadModuleDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryFilterer{contract: contract}, nil
}

// bindLilypadModuleDirectory binds a generic wrapper to an already deployed contract.
func bindLilypadModuleDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadModuleDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadModuleDirectory *LilypadModuleDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadModuleDirectory.Contract.LilypadModuleDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadModuleDirectory *LilypadModuleDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.LilypadModuleDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadModuleDirectory *LilypadModuleDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.LilypadModuleDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadModuleDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadModuleDirectory.Contract.DEFAULTADMINROLE(&_LilypadModuleDirectory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadModuleDirectory.Contract.DEFAULTADMINROLE(&_LilypadModuleDirectory.CallOpts)
}

// GetLilypadUser is a free data retrieval call binding the contract method 0x14dd1d2e.
//
// Solidity: function getLilypadUser() view returns(address)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) GetLilypadUser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "getLilypadUser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLilypadUser is a free data retrieval call binding the contract method 0x14dd1d2e.
//
// Solidity: function getLilypadUser() view returns(address)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) GetLilypadUser() (common.Address, error) {
	return _LilypadModuleDirectory.Contract.GetLilypadUser(&_LilypadModuleDirectory.CallOpts)
}

// GetLilypadUser is a free data retrieval call binding the contract method 0x14dd1d2e.
//
// Solidity: function getLilypadUser() view returns(address)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) GetLilypadUser() (common.Address, error) {
	return _LilypadModuleDirectory.Contract.GetLilypadUser(&_LilypadModuleDirectory.CallOpts)
}

// GetOwnedModules is a free data retrieval call binding the contract method 0x6fabef35.
//
// Solidity: function getOwnedModules(address moduleOwner) view returns((address,string,string)[])
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) GetOwnedModules(opts *bind.CallOpts, moduleOwner common.Address) ([]SharedStructsModule, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "getOwnedModules", moduleOwner)

	if err != nil {
		return *new([]SharedStructsModule), err
	}

	out0 := *abi.ConvertType(out[0], new([]SharedStructsModule)).(*[]SharedStructsModule)

	return out0, err

}

// GetOwnedModules is a free data retrieval call binding the contract method 0x6fabef35.
//
// Solidity: function getOwnedModules(address moduleOwner) view returns((address,string,string)[])
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) GetOwnedModules(moduleOwner common.Address) ([]SharedStructsModule, error) {
	return _LilypadModuleDirectory.Contract.GetOwnedModules(&_LilypadModuleDirectory.CallOpts, moduleOwner)
}

// GetOwnedModules is a free data retrieval call binding the contract method 0x6fabef35.
//
// Solidity: function getOwnedModules(address moduleOwner) view returns((address,string,string)[])
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) GetOwnedModules(moduleOwner common.Address) ([]SharedStructsModule, error) {
	return _LilypadModuleDirectory.Contract.GetOwnedModules(&_LilypadModuleDirectory.CallOpts, moduleOwner)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadModuleDirectory.Contract.GetRoleAdmin(&_LilypadModuleDirectory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadModuleDirectory.Contract.GetRoleAdmin(&_LilypadModuleDirectory.CallOpts, role)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) GetVersion() (string, error) {
	return _LilypadModuleDirectory.Contract.GetVersion(&_LilypadModuleDirectory.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) GetVersion() (string, error) {
	return _LilypadModuleDirectory.Contract.GetVersion(&_LilypadModuleDirectory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadModuleDirectory.Contract.HasRole(&_LilypadModuleDirectory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadModuleDirectory.Contract.HasRole(&_LilypadModuleDirectory.CallOpts, role, account)
}

// IsTransferApproved is a free data retrieval call binding the contract method 0xcafebf56.
//
// Solidity: function isTransferApproved(address moduleOwner, string moduleName, address purchaser) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) IsTransferApproved(opts *bind.CallOpts, moduleOwner common.Address, moduleName string, purchaser common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "isTransferApproved", moduleOwner, moduleName, purchaser)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferApproved is a free data retrieval call binding the contract method 0xcafebf56.
//
// Solidity: function isTransferApproved(address moduleOwner, string moduleName, address purchaser) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) IsTransferApproved(moduleOwner common.Address, moduleName string, purchaser common.Address) (bool, error) {
	return _LilypadModuleDirectory.Contract.IsTransferApproved(&_LilypadModuleDirectory.CallOpts, moduleOwner, moduleName, purchaser)
}

// IsTransferApproved is a free data retrieval call binding the contract method 0xcafebf56.
//
// Solidity: function isTransferApproved(address moduleOwner, string moduleName, address purchaser) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) IsTransferApproved(moduleOwner common.Address, moduleName string, purchaser common.Address) (bool, error) {
	return _LilypadModuleDirectory.Contract.IsTransferApproved(&_LilypadModuleDirectory.CallOpts, moduleOwner, moduleName, purchaser)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadModuleDirectory.Contract.SupportsInterface(&_LilypadModuleDirectory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadModuleDirectory.Contract.SupportsInterface(&_LilypadModuleDirectory.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadModuleDirectory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) Version() (string, error) {
	return _LilypadModuleDirectory.Contract.Version(&_LilypadModuleDirectory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadModuleDirectory *LilypadModuleDirectoryCallerSession) Version() (string, error) {
	return _LilypadModuleDirectory.Contract.Version(&_LilypadModuleDirectory.CallOpts)
}

// ApproveTransfer is a paid mutator transaction binding the contract method 0x3c189c9c.
//
// Solidity: function approveTransfer(address moduleOwner, address newOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) ApproveTransfer(opts *bind.TransactOpts, moduleOwner common.Address, newOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "approveTransfer", moduleOwner, newOwner, moduleName, moduleUrl)
}

// ApproveTransfer is a paid mutator transaction binding the contract method 0x3c189c9c.
//
// Solidity: function approveTransfer(address moduleOwner, address newOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) ApproveTransfer(moduleOwner common.Address, newOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.ApproveTransfer(&_LilypadModuleDirectory.TransactOpts, moduleOwner, newOwner, moduleName, moduleUrl)
}

// ApproveTransfer is a paid mutator transaction binding the contract method 0x3c189c9c.
//
// Solidity: function approveTransfer(address moduleOwner, address newOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) ApproveTransfer(moduleOwner common.Address, newOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.ApproveTransfer(&_LilypadModuleDirectory.TransactOpts, moduleOwner, newOwner, moduleName, moduleUrl)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.GrantRole(&_LilypadModuleDirectory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.GrantRole(&_LilypadModuleDirectory.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _lilypadUser) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) Initialize(opts *bind.TransactOpts, _lilypadUser common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "initialize", _lilypadUser)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _lilypadUser) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) Initialize(_lilypadUser common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.Initialize(&_LilypadModuleDirectory.TransactOpts, _lilypadUser)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _lilypadUser) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) Initialize(_lilypadUser common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.Initialize(&_LilypadModuleDirectory.TransactOpts, _lilypadUser)
}

// RegisterModuleCreator is a paid mutator transaction binding the contract method 0x334cc8bb.
//
// Solidity: function registerModuleCreator(address moduleCreator) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) RegisterModuleCreator(opts *bind.TransactOpts, moduleCreator common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "registerModuleCreator", moduleCreator)
}

// RegisterModuleCreator is a paid mutator transaction binding the contract method 0x334cc8bb.
//
// Solidity: function registerModuleCreator(address moduleCreator) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) RegisterModuleCreator(moduleCreator common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RegisterModuleCreator(&_LilypadModuleDirectory.TransactOpts, moduleCreator)
}

// RegisterModuleCreator is a paid mutator transaction binding the contract method 0x334cc8bb.
//
// Solidity: function registerModuleCreator(address moduleCreator) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) RegisterModuleCreator(moduleCreator common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RegisterModuleCreator(&_LilypadModuleDirectory.TransactOpts, moduleCreator)
}

// RegisterModuleForCreator is a paid mutator transaction binding the contract method 0x0ba123c5.
//
// Solidity: function registerModuleForCreator(address moduleOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) RegisterModuleForCreator(opts *bind.TransactOpts, moduleOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "registerModuleForCreator", moduleOwner, moduleName, moduleUrl)
}

// RegisterModuleForCreator is a paid mutator transaction binding the contract method 0x0ba123c5.
//
// Solidity: function registerModuleForCreator(address moduleOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) RegisterModuleForCreator(moduleOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RegisterModuleForCreator(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName, moduleUrl)
}

// RegisterModuleForCreator is a paid mutator transaction binding the contract method 0x0ba123c5.
//
// Solidity: function registerModuleForCreator(address moduleOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) RegisterModuleForCreator(moduleOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RegisterModuleForCreator(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName, moduleUrl)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RenounceRole(&_LilypadModuleDirectory.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RenounceRole(&_LilypadModuleDirectory.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RevokeRole(&_LilypadModuleDirectory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RevokeRole(&_LilypadModuleDirectory.TransactOpts, role, account)
}

// RevokeTransferApproval is a paid mutator transaction binding the contract method 0xf842f717.
//
// Solidity: function revokeTransferApproval(address moduleOwner, string moduleName) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) RevokeTransferApproval(opts *bind.TransactOpts, moduleOwner common.Address, moduleName string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "revokeTransferApproval", moduleOwner, moduleName)
}

// RevokeTransferApproval is a paid mutator transaction binding the contract method 0xf842f717.
//
// Solidity: function revokeTransferApproval(address moduleOwner, string moduleName) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) RevokeTransferApproval(moduleOwner common.Address, moduleName string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RevokeTransferApproval(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName)
}

// RevokeTransferApproval is a paid mutator transaction binding the contract method 0xf842f717.
//
// Solidity: function revokeTransferApproval(address moduleOwner, string moduleName) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) RevokeTransferApproval(moduleOwner common.Address, moduleName string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.RevokeTransferApproval(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName)
}

// SetLilypadUser is a paid mutator transaction binding the contract method 0xf30e1383.
//
// Solidity: function setLilypadUser(address _lilypadUser) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) SetLilypadUser(opts *bind.TransactOpts, _lilypadUser common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "setLilypadUser", _lilypadUser)
}

// SetLilypadUser is a paid mutator transaction binding the contract method 0xf30e1383.
//
// Solidity: function setLilypadUser(address _lilypadUser) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) SetLilypadUser(_lilypadUser common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.SetLilypadUser(&_LilypadModuleDirectory.TransactOpts, _lilypadUser)
}

// SetLilypadUser is a paid mutator transaction binding the contract method 0xf30e1383.
//
// Solidity: function setLilypadUser(address _lilypadUser) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) SetLilypadUser(_lilypadUser common.Address) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.SetLilypadUser(&_LilypadModuleDirectory.TransactOpts, _lilypadUser)
}

// TransferModuleOwnership is a paid mutator transaction binding the contract method 0xfc7d0fb8.
//
// Solidity: function transferModuleOwnership(address moduleOwner, address newOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) TransferModuleOwnership(opts *bind.TransactOpts, moduleOwner common.Address, newOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "transferModuleOwnership", moduleOwner, newOwner, moduleName, moduleUrl)
}

// TransferModuleOwnership is a paid mutator transaction binding the contract method 0xfc7d0fb8.
//
// Solidity: function transferModuleOwnership(address moduleOwner, address newOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) TransferModuleOwnership(moduleOwner common.Address, newOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.TransferModuleOwnership(&_LilypadModuleDirectory.TransactOpts, moduleOwner, newOwner, moduleName, moduleUrl)
}

// TransferModuleOwnership is a paid mutator transaction binding the contract method 0xfc7d0fb8.
//
// Solidity: function transferModuleOwnership(address moduleOwner, address newOwner, string moduleName, string moduleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) TransferModuleOwnership(moduleOwner common.Address, newOwner common.Address, moduleName string, moduleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.TransferModuleOwnership(&_LilypadModuleDirectory.TransactOpts, moduleOwner, newOwner, moduleName, moduleUrl)
}

// UpdateModuleName is a paid mutator transaction binding the contract method 0x2718e7ee.
//
// Solidity: function updateModuleName(address moduleOwner, string moduleName, string newModuleName) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) UpdateModuleName(opts *bind.TransactOpts, moduleOwner common.Address, moduleName string, newModuleName string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "updateModuleName", moduleOwner, moduleName, newModuleName)
}

// UpdateModuleName is a paid mutator transaction binding the contract method 0x2718e7ee.
//
// Solidity: function updateModuleName(address moduleOwner, string moduleName, string newModuleName) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) UpdateModuleName(moduleOwner common.Address, moduleName string, newModuleName string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.UpdateModuleName(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName, newModuleName)
}

// UpdateModuleName is a paid mutator transaction binding the contract method 0x2718e7ee.
//
// Solidity: function updateModuleName(address moduleOwner, string moduleName, string newModuleName) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) UpdateModuleName(moduleOwner common.Address, moduleName string, newModuleName string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.UpdateModuleName(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName, newModuleName)
}

// UpdateModuleUrl is a paid mutator transaction binding the contract method 0x45fe5e00.
//
// Solidity: function updateModuleUrl(address moduleOwner, string moduleName, string newModuleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactor) UpdateModuleUrl(opts *bind.TransactOpts, moduleOwner common.Address, moduleName string, newModuleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.contract.Transact(opts, "updateModuleUrl", moduleOwner, moduleName, newModuleUrl)
}

// UpdateModuleUrl is a paid mutator transaction binding the contract method 0x45fe5e00.
//
// Solidity: function updateModuleUrl(address moduleOwner, string moduleName, string newModuleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectorySession) UpdateModuleUrl(moduleOwner common.Address, moduleName string, newModuleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.UpdateModuleUrl(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName, newModuleUrl)
}

// UpdateModuleUrl is a paid mutator transaction binding the contract method 0x45fe5e00.
//
// Solidity: function updateModuleUrl(address moduleOwner, string moduleName, string newModuleUrl) returns(bool)
func (_LilypadModuleDirectory *LilypadModuleDirectoryTransactorSession) UpdateModuleUrl(moduleOwner common.Address, moduleName string, newModuleUrl string) (*types.Transaction, error) {
	return _LilypadModuleDirectory.Contract.UpdateModuleUrl(&_LilypadModuleDirectory.TransactOpts, moduleOwner, moduleName, newModuleUrl)
}

// LilypadModuleDirectoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryInitializedIterator struct {
	Event *LilypadModuleDirectoryInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryInitialized)
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
		it.Event = new(LilypadModuleDirectoryInitialized)
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
func (it *LilypadModuleDirectoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryInitialized represents a Initialized event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadModuleDirectoryInitializedIterator, error) {

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryInitializedIterator{contract: _LilypadModuleDirectory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryInitialized)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseInitialized(log types.Log) (*LilypadModuleDirectoryInitialized, error) {
	event := new(LilypadModuleDirectoryInitialized)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator is returned from FilterLilypadModuleDirectoryLilypadUserUpdated and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryLilypadUserUpdated events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated represents a LilypadModuleDirectoryLilypadUserUpdated event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated struct {
	LilypadUser common.Address
	Caller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryLilypadUserUpdated is a free log retrieval operation binding the contract event 0xe6c2ee42f256f2bb1bc0bbe151f8c780426cfcb8d702a13c776e9e57bd0f6762.
//
// Solidity: event LilypadModuleDirectory__LilypadUserUpdated(address indexed lilypadUser, address indexed caller)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryLilypadUserUpdated(opts *bind.FilterOpts, lilypadUser []common.Address, caller []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator, error) {

	var lilypadUserRule []interface{}
	for _, lilypadUserItem := range lilypadUser {
		lilypadUserRule = append(lilypadUserRule, lilypadUserItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__LilypadUserUpdated", lilypadUserRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdatedIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__LilypadUserUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryLilypadUserUpdated is a free log subscription operation binding the contract event 0xe6c2ee42f256f2bb1bc0bbe151f8c780426cfcb8d702a13c776e9e57bd0f6762.
//
// Solidity: event LilypadModuleDirectory__LilypadUserUpdated(address indexed lilypadUser, address indexed caller)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryLilypadUserUpdated(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated, lilypadUser []common.Address, caller []common.Address) (event.Subscription, error) {

	var lilypadUserRule []interface{}
	for _, lilypadUserItem := range lilypadUser {
		lilypadUserRule = append(lilypadUserRule, lilypadUserItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__LilypadUserUpdated", lilypadUserRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__LilypadUserUpdated", log); err != nil {
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

// ParseLilypadModuleDirectoryLilypadUserUpdated is a log parse operation binding the contract event 0xe6c2ee42f256f2bb1bc0bbe151f8c780426cfcb8d702a13c776e9e57bd0f6762.
//
// Solidity: event LilypadModuleDirectory__LilypadUserUpdated(address indexed lilypadUser, address indexed caller)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryLilypadUserUpdated(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryLilypadUserUpdated)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__LilypadUserUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator is returned from FilterLilypadModuleDirectoryModuleCreatorRegistered and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleCreatorRegistered events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered represents a LilypadModuleDirectoryModuleCreatorRegistered event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered struct {
	ModuleCreator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleCreatorRegistered is a free log retrieval operation binding the contract event 0x48391147531cad00caac8f4787b4e7b2b9993d44110a2ff1419ae33812ce27e0.
//
// Solidity: event LilypadModuleDirectory__ModuleCreatorRegistered(address indexed moduleCreator)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleCreatorRegistered(opts *bind.FilterOpts, moduleCreator []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator, error) {

	var moduleCreatorRule []interface{}
	for _, moduleCreatorItem := range moduleCreator {
		moduleCreatorRule = append(moduleCreatorRule, moduleCreatorItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleCreatorRegistered", moduleCreatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegisteredIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleCreatorRegistered", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleCreatorRegistered is a free log subscription operation binding the contract event 0x48391147531cad00caac8f4787b4e7b2b9993d44110a2ff1419ae33812ce27e0.
//
// Solidity: event LilypadModuleDirectory__ModuleCreatorRegistered(address indexed moduleCreator)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleCreatorRegistered(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered, moduleCreator []common.Address) (event.Subscription, error) {

	var moduleCreatorRule []interface{}
	for _, moduleCreatorItem := range moduleCreator {
		moduleCreatorRule = append(moduleCreatorRule, moduleCreatorItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleCreatorRegistered", moduleCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleCreatorRegistered", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleCreatorRegistered is a log parse operation binding the contract event 0x48391147531cad00caac8f4787b4e7b2b9993d44110a2ff1419ae33812ce27e0.
//
// Solidity: event LilypadModuleDirectory__ModuleCreatorRegistered(address indexed moduleCreator)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleCreatorRegistered(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleCreatorRegistered)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleCreatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator is returned from FilterLilypadModuleDirectoryModuleNameUpdated and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleNameUpdated events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated represents a LilypadModuleDirectoryModuleNameUpdated event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated struct {
	Owner         common.Address
	OldModuleName string
	NewModuleName string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleNameUpdated is a free log retrieval operation binding the contract event 0x64d5584871d4e36223536000f2a6fcdf9777c12e31e88671f032d0ba5b34d27e.
//
// Solidity: event LilypadModuleDirectory__ModuleNameUpdated(address indexed owner, string oldModuleName, string newModuleName)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleNameUpdated(opts *bind.FilterOpts, owner []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleNameUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdatedIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleNameUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleNameUpdated is a free log subscription operation binding the contract event 0x64d5584871d4e36223536000f2a6fcdf9777c12e31e88671f032d0ba5b34d27e.
//
// Solidity: event LilypadModuleDirectory__ModuleNameUpdated(address indexed owner, string oldModuleName, string newModuleName)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleNameUpdated(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleNameUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleNameUpdated", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleNameUpdated is a log parse operation binding the contract event 0x64d5584871d4e36223536000f2a6fcdf9777c12e31e88671f032d0ba5b34d27e.
//
// Solidity: event LilypadModuleDirectory__ModuleNameUpdated(address indexed owner, string oldModuleName, string newModuleName)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleNameUpdated(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleNameUpdated)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator is returned from FilterLilypadModuleDirectoryModuleRegistered and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleRegistered events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered represents a LilypadModuleDirectoryModuleRegistered event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered struct {
	Owner      common.Address
	ModuleName string
	ModuleUrl  string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleRegistered is a free log retrieval operation binding the contract event 0xba4f6fb1694d21cafaecdc84bf64dcd60229b4b1e93c88325dc875fb3a7ac3b3.
//
// Solidity: event LilypadModuleDirectory__ModuleRegistered(address indexed owner, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleRegistered(opts *bind.FilterOpts, owner []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleRegisteredIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleRegistered", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleRegistered is a free log subscription operation binding the contract event 0xba4f6fb1694d21cafaecdc84bf64dcd60229b4b1e93c88325dc875fb3a7ac3b3.
//
// Solidity: event LilypadModuleDirectory__ModuleRegistered(address indexed owner, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleRegistered(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleRegistered", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleRegistered is a log parse operation binding the contract event 0xba4f6fb1694d21cafaecdc84bf64dcd60229b4b1e93c88325dc875fb3a7ac3b3.
//
// Solidity: event LilypadModuleDirectory__ModuleRegistered(address indexed owner, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleRegistered(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleRegistered)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator is returned from FilterLilypadModuleDirectoryModuleTransferApproved and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleTransferApproved events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved represents a LilypadModuleDirectoryModuleTransferApproved event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved struct {
	Owner      common.Address
	Purchaser  common.Address
	ModuleName string
	ModuleUrl  string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleTransferApproved is a free log retrieval operation binding the contract event 0xbce1bd154986092afc92370dcd98072885be653515b462a9bc5d9217a824cfe6.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferApproved(address indexed owner, address indexed purchaser, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleTransferApproved(opts *bind.FilterOpts, owner []common.Address, purchaser []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleTransferApproved", ownerRule, purchaserRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApprovedIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleTransferApproved", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleTransferApproved is a free log subscription operation binding the contract event 0xbce1bd154986092afc92370dcd98072885be653515b462a9bc5d9217a824cfe6.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferApproved(address indexed owner, address indexed purchaser, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleTransferApproved(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved, owner []common.Address, purchaser []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleTransferApproved", ownerRule, purchaserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleTransferApproved", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleTransferApproved is a log parse operation binding the contract event 0xbce1bd154986092afc92370dcd98072885be653515b462a9bc5d9217a824cfe6.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferApproved(address indexed owner, address indexed purchaser, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleTransferApproved(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferApproved)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleTransferApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator is returned from FilterLilypadModuleDirectoryModuleTransferRevoked and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleTransferRevoked events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked represents a LilypadModuleDirectoryModuleTransferRevoked event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked struct {
	Owner       common.Address
	RevokedFrom common.Address
	ModuleName  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleTransferRevoked is a free log retrieval operation binding the contract event 0xf6078793aacb510d74514d4bbed67a3b0be79fa862f5a8dab6fee8f7d1d9612d.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferRevoked(address indexed owner, address indexed revokedFrom, string moduleName)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleTransferRevoked(opts *bind.FilterOpts, owner []common.Address, revokedFrom []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var revokedFromRule []interface{}
	for _, revokedFromItem := range revokedFrom {
		revokedFromRule = append(revokedFromRule, revokedFromItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleTransferRevoked", ownerRule, revokedFromRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevokedIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleTransferRevoked", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleTransferRevoked is a free log subscription operation binding the contract event 0xf6078793aacb510d74514d4bbed67a3b0be79fa862f5a8dab6fee8f7d1d9612d.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferRevoked(address indexed owner, address indexed revokedFrom, string moduleName)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleTransferRevoked(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked, owner []common.Address, revokedFrom []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var revokedFromRule []interface{}
	for _, revokedFromItem := range revokedFrom {
		revokedFromRule = append(revokedFromRule, revokedFromItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleTransferRevoked", ownerRule, revokedFromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleTransferRevoked", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleTransferRevoked is a log parse operation binding the contract event 0xf6078793aacb510d74514d4bbed67a3b0be79fa862f5a8dab6fee8f7d1d9612d.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferRevoked(address indexed owner, address indexed revokedFrom, string moduleName)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleTransferRevoked(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferRevoked)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleTransferRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator is returned from FilterLilypadModuleDirectoryModuleTransferred and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleTransferred events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred represents a LilypadModuleDirectoryModuleTransferred event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred struct {
	NewOwner      common.Address
	PreviousOwner common.Address
	ModuleName    string
	ModuleUrl     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleTransferred is a free log retrieval operation binding the contract event 0xe056b787162beedd119835ebe75422503530c7dd16da5304c3893551b180d2cd.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferred(address indexed newOwner, address indexed previousOwner, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleTransferred(opts *bind.FilterOpts, newOwner []common.Address, previousOwner []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleTransferred", newOwnerRule, previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferredIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleTransferred", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleTransferred is a free log subscription operation binding the contract event 0xe056b787162beedd119835ebe75422503530c7dd16da5304c3893551b180d2cd.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferred(address indexed newOwner, address indexed previousOwner, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleTransferred(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred, newOwner []common.Address, previousOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleTransferred", newOwnerRule, previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleTransferred", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleTransferred is a log parse operation binding the contract event 0xe056b787162beedd119835ebe75422503530c7dd16da5304c3893551b180d2cd.
//
// Solidity: event LilypadModuleDirectory__ModuleTransferred(address indexed newOwner, address indexed previousOwner, string moduleName, string moduleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleTransferred(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleTransferred)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator is returned from FilterLilypadModuleDirectoryModuleUrlUpdated and is used to iterate over the raw logs and unpacked data for LilypadModuleDirectoryModuleUrlUpdated events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator struct {
	Event *LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated)
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
		it.Event = new(LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated)
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
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated represents a LilypadModuleDirectoryModuleUrlUpdated event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated struct {
	Owner        common.Address
	ModuleName   string
	NewModuleUrl string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLilypadModuleDirectoryModuleUrlUpdated is a free log retrieval operation binding the contract event 0xd002b7b619e16e3e9e93faff3fda6a96a8420d0a7e48f23e69745507a9971c62.
//
// Solidity: event LilypadModuleDirectory__ModuleUrlUpdated(address indexed owner, string moduleName, string newModuleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterLilypadModuleDirectoryModuleUrlUpdated(opts *bind.FilterOpts, owner []common.Address) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "LilypadModuleDirectory__ModuleUrlUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdatedIterator{contract: _LilypadModuleDirectory.contract, event: "LilypadModuleDirectory__ModuleUrlUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadModuleDirectoryModuleUrlUpdated is a free log subscription operation binding the contract event 0xd002b7b619e16e3e9e93faff3fda6a96a8420d0a7e48f23e69745507a9971c62.
//
// Solidity: event LilypadModuleDirectory__ModuleUrlUpdated(address indexed owner, string moduleName, string newModuleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchLilypadModuleDirectoryModuleUrlUpdated(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "LilypadModuleDirectory__ModuleUrlUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleUrlUpdated", log); err != nil {
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

// ParseLilypadModuleDirectoryModuleUrlUpdated is a log parse operation binding the contract event 0xd002b7b619e16e3e9e93faff3fda6a96a8420d0a7e48f23e69745507a9971c62.
//
// Solidity: event LilypadModuleDirectory__ModuleUrlUpdated(address indexed owner, string moduleName, string newModuleUrl)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseLilypadModuleDirectoryModuleUrlUpdated(log types.Log) (*LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated, error) {
	event := new(LilypadModuleDirectoryLilypadModuleDirectoryModuleUrlUpdated)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "LilypadModuleDirectory__ModuleUrlUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryRoleAdminChangedIterator struct {
	Event *LilypadModuleDirectoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryRoleAdminChanged)
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
		it.Event = new(LilypadModuleDirectoryRoleAdminChanged)
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
func (it *LilypadModuleDirectoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadModuleDirectoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryRoleAdminChangedIterator{contract: _LilypadModuleDirectory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryRoleAdminChanged)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadModuleDirectoryRoleAdminChanged, error) {
	event := new(LilypadModuleDirectoryRoleAdminChanged)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryRoleGrantedIterator struct {
	Event *LilypadModuleDirectoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryRoleGranted)
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
		it.Event = new(LilypadModuleDirectoryRoleGranted)
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
func (it *LilypadModuleDirectoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryRoleGranted represents a RoleGranted event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadModuleDirectoryRoleGrantedIterator, error) {

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

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryRoleGrantedIterator{contract: _LilypadModuleDirectory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryRoleGranted)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseRoleGranted(log types.Log) (*LilypadModuleDirectoryRoleGranted, error) {
	event := new(LilypadModuleDirectoryRoleGranted)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadModuleDirectoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryRoleRevokedIterator struct {
	Event *LilypadModuleDirectoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadModuleDirectoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadModuleDirectoryRoleRevoked)
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
		it.Event = new(LilypadModuleDirectoryRoleRevoked)
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
func (it *LilypadModuleDirectoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadModuleDirectoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadModuleDirectoryRoleRevoked represents a RoleRevoked event raised by the LilypadModuleDirectory contract.
type LilypadModuleDirectoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadModuleDirectoryRoleRevokedIterator, error) {

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

	logs, sub, err := _LilypadModuleDirectory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadModuleDirectoryRoleRevokedIterator{contract: _LilypadModuleDirectory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadModuleDirectoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadModuleDirectory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadModuleDirectoryRoleRevoked)
				if err := _LilypadModuleDirectory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LilypadModuleDirectory *LilypadModuleDirectoryFilterer) ParseRoleRevoked(log types.Log) (*LilypadModuleDirectoryRoleRevoked, error) {
	event := new(LilypadModuleDirectoryRoleRevoked)
	if err := _LilypadModuleDirectory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
