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

// ILilypadTokenMetaData contains all meta data concerning the ILilypadToken contract.
var ILilypadTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"escrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refundEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILilypadTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ILilypadTokenMetaData.ABI instead.
var ILilypadTokenABI = ILilypadTokenMetaData.ABI

// ILilypadToken is an auto generated Go binding around an Ethereum contract.
type ILilypadToken struct {
	ILilypadTokenCaller     // Read-only binding to the contract
	ILilypadTokenTransactor // Write-only binding to the contract
	ILilypadTokenFilterer   // Log filterer for contract events
}

// ILilypadTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILilypadTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILilypadTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILilypadTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILilypadTokenSession struct {
	Contract     *ILilypadToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILilypadTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILilypadTokenCallerSession struct {
	Contract *ILilypadTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ILilypadTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILilypadTokenTransactorSession struct {
	Contract     *ILilypadTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ILilypadTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILilypadTokenRaw struct {
	Contract *ILilypadToken // Generic contract binding to access the raw methods on
}

// ILilypadTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILilypadTokenCallerRaw struct {
	Contract *ILilypadTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ILilypadTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILilypadTokenTransactorRaw struct {
	Contract *ILilypadTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILilypadToken creates a new instance of ILilypadToken, bound to a specific deployed contract.
func NewILilypadToken(address common.Address, backend bind.ContractBackend) (*ILilypadToken, error) {
	contract, err := bindILilypadToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILilypadToken{ILilypadTokenCaller: ILilypadTokenCaller{contract: contract}, ILilypadTokenTransactor: ILilypadTokenTransactor{contract: contract}, ILilypadTokenFilterer: ILilypadTokenFilterer{contract: contract}}, nil
}

// NewILilypadTokenCaller creates a new read-only instance of ILilypadToken, bound to a specific deployed contract.
func NewILilypadTokenCaller(address common.Address, caller bind.ContractCaller) (*ILilypadTokenCaller, error) {
	contract, err := bindILilypadToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILilypadTokenCaller{contract: contract}, nil
}

// NewILilypadTokenTransactor creates a new write-only instance of ILilypadToken, bound to a specific deployed contract.
func NewILilypadTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ILilypadTokenTransactor, error) {
	contract, err := bindILilypadToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILilypadTokenTransactor{contract: contract}, nil
}

// NewILilypadTokenFilterer creates a new log filterer instance of ILilypadToken, bound to a specific deployed contract.
func NewILilypadTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ILilypadTokenFilterer, error) {
	contract, err := bindILilypadToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILilypadTokenFilterer{contract: contract}, nil
}

// bindILilypadToken binds a generic wrapper to an already deployed contract.
func bindILilypadToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILilypadTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILilypadToken *ILilypadTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILilypadToken.Contract.ILilypadTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILilypadToken *ILilypadTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILilypadToken.Contract.ILilypadTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILilypadToken *ILilypadTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILilypadToken.Contract.ILilypadTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILilypadToken *ILilypadTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILilypadToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILilypadToken *ILilypadTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILilypadToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILilypadToken *ILilypadTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILilypadToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ILilypadToken *ILilypadTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILilypadToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ILilypadToken *ILilypadTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ILilypadToken.Contract.Allowance(&_ILilypadToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ILilypadToken *ILilypadTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ILilypadToken.Contract.Allowance(&_ILilypadToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ILilypadToken *ILilypadTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILilypadToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ILilypadToken *ILilypadTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ILilypadToken.Contract.BalanceOf(&_ILilypadToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ILilypadToken *ILilypadTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ILilypadToken.Contract.BalanceOf(&_ILilypadToken.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ILilypadToken *ILilypadTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILilypadToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ILilypadToken *ILilypadTokenSession) TotalSupply() (*big.Int, error) {
	return _ILilypadToken.Contract.TotalSupply(&_ILilypadToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ILilypadToken *ILilypadTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ILilypadToken.Contract.TotalSupply(&_ILilypadToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.Approve(&_ILilypadToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.Approve(&_ILilypadToken.TransactOpts, spender, amount)
}

// EscrowBalanceOf is a paid mutator transaction binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) returns(uint256)
func (_ILilypadToken *ILilypadTokenTransactor) EscrowBalanceOf(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "escrowBalanceOf", _address)
}

// EscrowBalanceOf is a paid mutator transaction binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) returns(uint256)
func (_ILilypadToken *ILilypadTokenSession) EscrowBalanceOf(_address common.Address) (*types.Transaction, error) {
	return _ILilypadToken.Contract.EscrowBalanceOf(&_ILilypadToken.TransactOpts, _address)
}

// EscrowBalanceOf is a paid mutator transaction binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) returns(uint256)
func (_ILilypadToken *ILilypadTokenTransactorSession) EscrowBalanceOf(_address common.Address) (*types.Transaction, error) {
	return _ILilypadToken.Contract.EscrowBalanceOf(&_ILilypadToken.TransactOpts, _address)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) PayEscrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "payEscrow", amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.PayEscrow(&_ILilypadToken.TransactOpts, amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.PayEscrow(&_ILilypadToken.TransactOpts, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) PayJob(opts *bind.TransactOpts, fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "payJob", fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.PayJob(&_ILilypadToken.TransactOpts, fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.PayJob(&_ILilypadToken.TransactOpts, fromAddress, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) RefundEscrow(opts *bind.TransactOpts, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "refundEscrow", toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.RefundEscrow(&_ILilypadToken.TransactOpts, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.RefundEscrow(&_ILilypadToken.TransactOpts, toAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) SlashEscrow(opts *bind.TransactOpts, slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "slashEscrow", slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.SlashEscrow(&_ILilypadToken.TransactOpts, slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.SlashEscrow(&_ILilypadToken.TransactOpts, slashedAddress, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.Transfer(&_ILilypadToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.Transfer(&_ILilypadToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.TransferFrom(&_ILilypadToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ILilypadToken *ILilypadTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILilypadToken.Contract.TransferFrom(&_ILilypadToken.TransactOpts, from, to, amount)
}

// ILilypadTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ILilypadToken contract.
type ILilypadTokenApprovalIterator struct {
	Event *ILilypadTokenApproval // Event containing the contract specifics and raw log

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
func (it *ILilypadTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILilypadTokenApproval)
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
		it.Event = new(ILilypadTokenApproval)
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
func (it *ILilypadTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILilypadTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILilypadTokenApproval represents a Approval event raised by the ILilypadToken contract.
type ILilypadTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ILilypadToken *ILilypadTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ILilypadTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ILilypadToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ILilypadTokenApprovalIterator{contract: _ILilypadToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ILilypadToken *ILilypadTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ILilypadTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ILilypadToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILilypadTokenApproval)
				if err := _ILilypadToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ILilypadToken *ILilypadTokenFilterer) ParseApproval(log types.Log) (*ILilypadTokenApproval, error) {
	event := new(ILilypadTokenApproval)
	if err := _ILilypadToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILilypadTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ILilypadToken contract.
type ILilypadTokenTransferIterator struct {
	Event *ILilypadTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ILilypadTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILilypadTokenTransfer)
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
		it.Event = new(ILilypadTokenTransfer)
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
func (it *ILilypadTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILilypadTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILilypadTokenTransfer represents a Transfer event raised by the ILilypadToken contract.
type ILilypadTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ILilypadToken *ILilypadTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ILilypadTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILilypadToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILilypadTokenTransferIterator{contract: _ILilypadToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ILilypadToken *ILilypadTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ILilypadTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILilypadToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILilypadTokenTransfer)
				if err := _ILilypadToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ILilypadToken *ILilypadTokenFilterer) ParseTransfer(log types.Log) (*ILilypadTokenTransfer, error) {
	event := new(ILilypadTokenTransfer)
	if err := _ILilypadToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
