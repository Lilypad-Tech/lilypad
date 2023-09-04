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

// LilypadTokenMetaData contains all meta data concerning the LilypadToken contract.
var LilypadTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"escrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refundEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadTokenMetaData.ABI instead.
var LilypadTokenABI = LilypadTokenMetaData.ABI

// LilypadToken is an auto generated Go binding around an Ethereum contract.
type LilypadToken struct {
	LilypadTokenCaller     // Read-only binding to the contract
	LilypadTokenTransactor // Write-only binding to the contract
	LilypadTokenFilterer   // Log filterer for contract events
}

// LilypadTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadTokenSession struct {
	Contract     *LilypadToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LilypadTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadTokenCallerSession struct {
	Contract *LilypadTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LilypadTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadTokenTransactorSession struct {
	Contract     *LilypadTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LilypadTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadTokenRaw struct {
	Contract *LilypadToken // Generic contract binding to access the raw methods on
}

// LilypadTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadTokenCallerRaw struct {
	Contract *LilypadTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadTokenTransactorRaw struct {
	Contract *LilypadTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadToken creates a new instance of LilypadToken, bound to a specific deployed contract.
func NewLilypadToken(address common.Address, backend bind.ContractBackend) (*LilypadToken, error) {
	contract, err := bindLilypadToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadToken{LilypadTokenCaller: LilypadTokenCaller{contract: contract}, LilypadTokenTransactor: LilypadTokenTransactor{contract: contract}, LilypadTokenFilterer: LilypadTokenFilterer{contract: contract}}, nil
}

// NewLilypadTokenCaller creates a new read-only instance of LilypadToken, bound to a specific deployed contract.
func NewLilypadTokenCaller(address common.Address, caller bind.ContractCaller) (*LilypadTokenCaller, error) {
	contract, err := bindLilypadToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenCaller{contract: contract}, nil
}

// NewLilypadTokenTransactor creates a new write-only instance of LilypadToken, bound to a specific deployed contract.
func NewLilypadTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadTokenTransactor, error) {
	contract, err := bindLilypadToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTransactor{contract: contract}, nil
}

// NewLilypadTokenFilterer creates a new log filterer instance of LilypadToken, bound to a specific deployed contract.
func NewLilypadTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadTokenFilterer, error) {
	contract, err := bindLilypadToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenFilterer{contract: contract}, nil
}

// bindLilypadToken binds a generic wrapper to an already deployed contract.
func bindLilypadToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadToken *LilypadTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadToken.Contract.LilypadTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadToken *LilypadTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadToken.Contract.LilypadTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadToken *LilypadTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadToken.Contract.LilypadTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadToken *LilypadTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadToken *LilypadTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadToken *LilypadTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LilypadToken *LilypadTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LilypadToken *LilypadTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LilypadToken.Contract.Allowance(&_LilypadToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LilypadToken *LilypadTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LilypadToken.Contract.Allowance(&_LilypadToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LilypadToken *LilypadTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LilypadToken *LilypadTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LilypadToken.Contract.BalanceOf(&_LilypadToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LilypadToken *LilypadTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LilypadToken.Contract.BalanceOf(&_LilypadToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LilypadToken *LilypadTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LilypadToken *LilypadTokenSession) Decimals() (uint8, error) {
	return _LilypadToken.Contract.Decimals(&_LilypadToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LilypadToken *LilypadTokenCallerSession) Decimals() (uint8, error) {
	return _LilypadToken.Contract.Decimals(&_LilypadToken.CallOpts)
}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_LilypadToken *LilypadTokenCaller) EscrowBalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "escrowBalanceOf", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_LilypadToken *LilypadTokenSession) EscrowBalanceOf(_address common.Address) (*big.Int, error) {
	return _LilypadToken.Contract.EscrowBalanceOf(&_LilypadToken.CallOpts, _address)
}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_LilypadToken *LilypadTokenCallerSession) EscrowBalanceOf(_address common.Address) (*big.Int, error) {
	return _LilypadToken.Contract.EscrowBalanceOf(&_LilypadToken.CallOpts, _address)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LilypadToken *LilypadTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LilypadToken *LilypadTokenSession) Name() (string, error) {
	return _LilypadToken.Contract.Name(&_LilypadToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LilypadToken *LilypadTokenCallerSession) Name() (string, error) {
	return _LilypadToken.Contract.Name(&_LilypadToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadToken *LilypadTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadToken *LilypadTokenSession) Owner() (common.Address, error) {
	return _LilypadToken.Contract.Owner(&_LilypadToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadToken *LilypadTokenCallerSession) Owner() (common.Address, error) {
	return _LilypadToken.Contract.Owner(&_LilypadToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LilypadToken *LilypadTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LilypadToken *LilypadTokenSession) Symbol() (string, error) {
	return _LilypadToken.Contract.Symbol(&_LilypadToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LilypadToken *LilypadTokenCallerSession) Symbol() (string, error) {
	return _LilypadToken.Contract.Symbol(&_LilypadToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LilypadToken *LilypadTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LilypadToken *LilypadTokenSession) TotalSupply() (*big.Int, error) {
	return _LilypadToken.Contract.TotalSupply(&_LilypadToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LilypadToken *LilypadTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LilypadToken.Contract.TotalSupply(&_LilypadToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.Approve(&_LilypadToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.Approve(&_LilypadToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LilypadToken *LilypadTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.DecreaseAllowance(&_LilypadToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.DecreaseAllowance(&_LilypadToken.TransactOpts, spender, subtractedValue)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadToken *LilypadTokenTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadToken *LilypadTokenSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadToken.Contract.DisableChangeControllerAddress(&_LilypadToken.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadToken *LilypadTokenTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadToken.Contract.DisableChangeControllerAddress(&_LilypadToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LilypadToken *LilypadTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.IncreaseAllowance(&_LilypadToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.IncreaseAllowance(&_LilypadToken.TransactOpts, spender, addedValue)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) PayEscrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "payEscrow", amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.PayEscrow(&_LilypadToken.TransactOpts, amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.PayEscrow(&_LilypadToken.TransactOpts, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) PayJob(opts *bind.TransactOpts, fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "payJob", fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.PayJob(&_LilypadToken.TransactOpts, fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.PayJob(&_LilypadToken.TransactOpts, fromAddress, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) RefundEscrow(opts *bind.TransactOpts, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "refundEscrow", toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.RefundEscrow(&_LilypadToken.TransactOpts, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.RefundEscrow(&_LilypadToken.TransactOpts, toAddress, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadToken *LilypadTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadToken *LilypadTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadToken.Contract.RenounceOwnership(&_LilypadToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadToken *LilypadTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadToken.Contract.RenounceOwnership(&_LilypadToken.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadToken *LilypadTokenTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadToken *LilypadTokenSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadToken.Contract.SetControllerAddress(&_LilypadToken.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadToken *LilypadTokenTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadToken.Contract.SetControllerAddress(&_LilypadToken.TransactOpts, _controllerAddress)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) SlashEscrow(opts *bind.TransactOpts, slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "slashEscrow", slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.SlashEscrow(&_LilypadToken.TransactOpts, slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.SlashEscrow(&_LilypadToken.TransactOpts, slashedAddress, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.Transfer(&_LilypadToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.Transfer(&_LilypadToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.TransferFrom(&_LilypadToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LilypadToken *LilypadTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LilypadToken.Contract.TransferFrom(&_LilypadToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadToken *LilypadTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LilypadToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadToken *LilypadTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadToken.Contract.TransferOwnership(&_LilypadToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadToken *LilypadTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadToken.Contract.TransferOwnership(&_LilypadToken.TransactOpts, newOwner)
}

// LilypadTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LilypadToken contract.
type LilypadTokenApprovalIterator struct {
	Event *LilypadTokenApproval // Event containing the contract specifics and raw log

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
func (it *LilypadTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenApproval)
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
		it.Event = new(LilypadTokenApproval)
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
func (it *LilypadTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenApproval represents a Approval event raised by the LilypadToken contract.
type LilypadTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LilypadToken *LilypadTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LilypadTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LilypadToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenApprovalIterator{contract: _LilypadToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LilypadToken *LilypadTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LilypadTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LilypadToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenApproval)
				if err := _LilypadToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LilypadToken *LilypadTokenFilterer) ParseApproval(log types.Log) (*LilypadTokenApproval, error) {
	event := new(LilypadTokenApproval)
	if err := _LilypadToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LilypadToken contract.
type LilypadTokenOwnershipTransferredIterator struct {
	Event *LilypadTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LilypadTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenOwnershipTransferred)
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
		it.Event = new(LilypadTokenOwnershipTransferred)
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
func (it *LilypadTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenOwnershipTransferred represents a OwnershipTransferred event raised by the LilypadToken contract.
type LilypadTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadToken *LilypadTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LilypadTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenOwnershipTransferredIterator{contract: _LilypadToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadToken *LilypadTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LilypadTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenOwnershipTransferred)
				if err := _LilypadToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LilypadToken *LilypadTokenFilterer) ParseOwnershipTransferred(log types.Log) (*LilypadTokenOwnershipTransferred, error) {
	event := new(LilypadTokenOwnershipTransferred)
	if err := _LilypadToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LilypadToken contract.
type LilypadTokenTransferIterator struct {
	Event *LilypadTokenTransfer // Event containing the contract specifics and raw log

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
func (it *LilypadTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadTokenTransfer)
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
		it.Event = new(LilypadTokenTransfer)
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
func (it *LilypadTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadTokenTransfer represents a Transfer event raised by the LilypadToken contract.
type LilypadTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LilypadToken *LilypadTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LilypadTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LilypadToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LilypadTokenTransferIterator{contract: _LilypadToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LilypadToken *LilypadTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LilypadTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LilypadToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadTokenTransfer)
				if err := _LilypadToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LilypadToken *LilypadTokenFilterer) ParseTransfer(log types.Log) (*LilypadTokenTransfer, error) {
	event := new(LilypadTokenTransfer)
	if err := _LilypadToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
