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

// LilypadTokenTestableMetaData contains all meta data concerning the LilypadTokenTestable contract.
var LilypadTokenTestableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"escrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refundEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadTokenTestableABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadTokenTestableMetaData.ABI instead.
var LilypadTokenTestableABI = LilypadTokenTestableMetaData.ABI

// LilypadTokenTestable is an auto generated Go binding around an Ethereum contract.
type LilypadTokenTestable struct {
	LilypadTokenTestableCaller     // Read-only binding to the contract
	LilypadTokenTestableTransactor // Write-only binding to the contract
	LilypadTokenTestableFilterer   // Log filterer for contract events
}

// LilypadTokenTestableCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadTokenTestableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenTestableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadTokenTestableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenTestableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadTokenTestableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenTestableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadTokenTestableSession struct {
	Contract     *LilypadTokenTestable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LilypadTokenTestableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadTokenTestableCallerSession struct {
	Contract *LilypadTokenTestableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// LilypadTokenTestableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadTokenTestableTransactorSession struct {
	Contract     *LilypadTokenTestableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// LilypadTokenTestableRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadTokenTestableRaw struct {
	Contract *LilypadTokenTestable // Generic contract binding to access the raw methods on
}

// LilypadTokenTestableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadTokenTestableCallerRaw struct {
	Contract *LilypadTokenTestableCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadTokenTestableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadTokenTestableTransactorRaw struct {
	Contract *LilypadTokenTestableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadTokenTestable creates a new instance of LilypadTokenTestable, bound to a specific deployed contract.
func NewLilypadTokenTestable(address common.Address, backend bind.ContractBackend) (*LilypadTokenTestable, error) {
	contract, err := bindLilypadTokenTestable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestable{LilypadTokenTestableCaller: LilypadTokenTestableCaller{contract: contract}, LilypadTokenTestableTransactor: LilypadTokenTestableTransactor{contract: contract}, LilypadTokenTestableFilterer: LilypadTokenTestableFilterer{contract: contract}}, nil
}

// NewLilypadTokenTestableCaller creates a new read-only instance of LilypadTokenTestable, bound to a specific deployed contract.
func NewLilypadTokenTestableCaller(address common.Address, caller bind.ContractCaller) (*LilypadTokenTestableCaller, error) {
	contract, err := bindLilypadTokenTestable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestableCaller{contract: contract}, nil
}

// NewLilypadTokenTestableTransactor creates a new write-only instance of LilypadTokenTestable, bound to a specific deployed contract.
func NewLilypadTokenTestableTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadTokenTestableTransactor, error) {
	contract, err := bindLilypadTokenTestable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestableTransactor{contract: contract}, nil
}

// NewLilypadTokenTestableFilterer creates a new log filterer instance of LilypadTokenTestable, bound to a specific deployed contract.
func NewLilypadTokenTestableFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadTokenTestableFilterer, error) {
	contract, err := bindLilypadTokenTestable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestableFilterer{contract: contract}, nil
}

// bindLilypadTokenTestable binds a generic wrapper to an already deployed contract.
func bindLilypadTokenTestable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadTokenTestableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadTokenTestable *LilypadTokenTestableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadTokenTestable.Contract.LilypadTokenTestableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadTokenTestable *LilypadTokenTestableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.LilypadTokenTestableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadTokenTestable *LilypadTokenTestableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.LilypadTokenTestableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadTokenTestable *LilypadTokenTestableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadTokenTestable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadTokenTestable *LilypadTokenTestableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadTokenTestable *LilypadTokenTestableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LilypadTokenTestable.Contract.Allowance(&_LilypadTokenTestable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LilypadTokenTestable.Contract.Allowance(&_LilypadTokenTestable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LilypadTokenTestable.Contract.BalanceOf(&_LilypadTokenTestable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LilypadTokenTestable.Contract.BalanceOf(&_LilypadTokenTestable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Decimals() (uint8, error) {
	return _LilypadTokenTestable.Contract.Decimals(&_LilypadTokenTestable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) Decimals() (uint8, error) {
	return _LilypadTokenTestable.Contract.Decimals(&_LilypadTokenTestable.CallOpts)
}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) EscrowBalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "escrowBalanceOf", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableSession) EscrowBalanceOf(_address common.Address) (*big.Int, error) {
	return _LilypadTokenTestable.Contract.EscrowBalanceOf(&_LilypadTokenTestable.CallOpts, _address)
}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) EscrowBalanceOf(_address common.Address) (*big.Int, error) {
	return _LilypadTokenTestable.Contract.EscrowBalanceOf(&_LilypadTokenTestable.CallOpts, _address)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Name() (string, error) {
	return _LilypadTokenTestable.Contract.Name(&_LilypadTokenTestable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) Name() (string, error) {
	return _LilypadTokenTestable.Contract.Name(&_LilypadTokenTestable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Owner() (common.Address, error) {
	return _LilypadTokenTestable.Contract.Owner(&_LilypadTokenTestable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) Owner() (common.Address, error) {
	return _LilypadTokenTestable.Contract.Owner(&_LilypadTokenTestable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Symbol() (string, error) {
	return _LilypadTokenTestable.Contract.Symbol(&_LilypadTokenTestable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) Symbol() (string, error) {
	return _LilypadTokenTestable.Contract.Symbol(&_LilypadTokenTestable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadTokenTestable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableSession) TotalSupply() (*big.Int, error) {
	return _LilypadTokenTestable.Contract.TotalSupply(&_LilypadTokenTestable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LilypadTokenTestable *LilypadTokenTestableCallerSession) TotalSupply() (*big.Int, error) {
	return _LilypadTokenTestable.Contract.TotalSupply(&_LilypadTokenTestable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.Approve(&_LilypadTokenTestable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.Approve(&_LilypadTokenTestable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.DecreaseAllowance(&_LilypadTokenTestable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.DecreaseAllowance(&_LilypadTokenTestable.TransactOpts, spender, subtractedValue)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadTokenTestable *LilypadTokenTestableSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.DisableChangeControllerAddress(&_LilypadTokenTestable.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.DisableChangeControllerAddress(&_LilypadTokenTestable.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.IncreaseAllowance(&_LilypadTokenTestable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.IncreaseAllowance(&_LilypadTokenTestable.TransactOpts, spender, addedValue)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) PayEscrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "payEscrow", amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.PayEscrow(&_LilypadTokenTestable.TransactOpts, amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.PayEscrow(&_LilypadTokenTestable.TransactOpts, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) PayJob(opts *bind.TransactOpts, fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "payJob", fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.PayJob(&_LilypadTokenTestable.TransactOpts, fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.PayJob(&_LilypadTokenTestable.TransactOpts, fromAddress, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) RefundEscrow(opts *bind.TransactOpts, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "refundEscrow", toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.RefundEscrow(&_LilypadTokenTestable.TransactOpts, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.RefundEscrow(&_LilypadTokenTestable.TransactOpts, toAddress, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadTokenTestable *LilypadTokenTestableSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.RenounceOwnership(&_LilypadTokenTestable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.RenounceOwnership(&_LilypadTokenTestable.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadTokenTestable *LilypadTokenTestableSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.SetControllerAddress(&_LilypadTokenTestable.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.SetControllerAddress(&_LilypadTokenTestable.TransactOpts, _controllerAddress)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) SlashEscrow(opts *bind.TransactOpts, slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "slashEscrow", slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.SlashEscrow(&_LilypadTokenTestable.TransactOpts, slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.SlashEscrow(&_LilypadTokenTestable.TransactOpts, slashedAddress, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.Transfer(&_LilypadTokenTestable.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.Transfer(&_LilypadTokenTestable.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.TransferFrom(&_LilypadTokenTestable.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.TransferFrom(&_LilypadTokenTestable.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LilypadTokenTestable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadTokenTestable *LilypadTokenTestableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.TransferOwnership(&_LilypadTokenTestable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadTokenTestable *LilypadTokenTestableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadTokenTestable.Contract.TransferOwnership(&_LilypadTokenTestable.TransactOpts, newOwner)
}

// LilypadTokenTestableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LilypadTokenTestable contract.
type LilypadTokenTestableApprovalIterator struct {
	Event *LilypadTokenTestableApproval // Event containing the contract specifics and raw log

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
func (it *LilypadTokenTestableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenTestableApproval)
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
		it.Event = new(LilypadTokenTestableApproval)
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
func (it *LilypadTokenTestableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenTestableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenTestableApproval represents a Approval event raised by the LilypadTokenTestable contract.
type LilypadTokenTestableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LilypadTokenTestableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LilypadTokenTestable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestableApprovalIterator{contract: _LilypadTokenTestable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LilypadTokenTestableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LilypadTokenTestable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenTestableApproval)
				if err := _LilypadTokenTestable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) ParseApproval(log types.Log) (*LilypadTokenTestableApproval, error) {
	event := new(LilypadTokenTestableApproval)
	if err := _LilypadTokenTestable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenTestableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LilypadTokenTestable contract.
type LilypadTokenTestableOwnershipTransferredIterator struct {
	Event *LilypadTokenTestableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LilypadTokenTestableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenTestableOwnershipTransferred)
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
		it.Event = new(LilypadTokenTestableOwnershipTransferred)
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
func (it *LilypadTokenTestableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenTestableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenTestableOwnershipTransferred represents a OwnershipTransferred event raised by the LilypadTokenTestable contract.
type LilypadTokenTestableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LilypadTokenTestableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadTokenTestable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestableOwnershipTransferredIterator{contract: _LilypadTokenTestable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LilypadTokenTestableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadTokenTestable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenTestableOwnershipTransferred)
				if err := _LilypadTokenTestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) ParseOwnershipTransferred(log types.Log) (*LilypadTokenTestableOwnershipTransferred, error) {
	event := new(LilypadTokenTestableOwnershipTransferred)
	if err := _LilypadTokenTestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenTestableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LilypadTokenTestable contract.
type LilypadTokenTestableTransferIterator struct {
	Event *LilypadTokenTestableTransfer // Event containing the contract specifics and raw log

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
func (it *LilypadTokenTestableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenTestableTransfer)
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
		it.Event = new(LilypadTokenTestableTransfer)
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
func (it *LilypadTokenTestableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenTestableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenTestableTransfer represents a Transfer event raised by the LilypadTokenTestable contract.
type LilypadTokenTestableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LilypadTokenTestableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LilypadTokenTestable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTestableTransferIterator{contract: _LilypadTokenTestable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LilypadTokenTestableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LilypadTokenTestable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenTestableTransfer)
				if err := _LilypadTokenTestable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LilypadTokenTestable *LilypadTokenTestableFilterer) ParseTransfer(log types.Log) (*LilypadTokenTestableTransfer, error) {
	event := new(LilypadTokenTestableTransfer)
	if err := _LilypadTokenTestable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
