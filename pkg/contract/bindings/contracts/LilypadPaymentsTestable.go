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

// LilypadPaymentsTestableMetaData contains all meta data concerning the LilypadPaymentsTestable contract.
var LilypadPaymentsTestableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentDirection\",\"name\":\"direction\",\"type\":\"uint8\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadPaymentsTestableABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadPaymentsTestableMetaData.ABI instead.
var LilypadPaymentsTestableABI = LilypadPaymentsTestableMetaData.ABI

// LilypadPaymentsTestable is an auto generated Go binding around an Ethereum contract.
type LilypadPaymentsTestable struct {
	LilypadPaymentsTestableCaller     // Read-only binding to the contract
	LilypadPaymentsTestableTransactor // Write-only binding to the contract
	LilypadPaymentsTestableFilterer   // Log filterer for contract events
}

// LilypadPaymentsTestableCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadPaymentsTestableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadPaymentsTestableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadPaymentsTestableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadPaymentsTestableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadPaymentsTestableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadPaymentsTestableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadPaymentsTestableSession struct {
	Contract     *LilypadPaymentsTestable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LilypadPaymentsTestableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadPaymentsTestableCallerSession struct {
	Contract *LilypadPaymentsTestableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// LilypadPaymentsTestableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadPaymentsTestableTransactorSession struct {
	Contract     *LilypadPaymentsTestableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// LilypadPaymentsTestableRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadPaymentsTestableRaw struct {
	Contract *LilypadPaymentsTestable // Generic contract binding to access the raw methods on
}

// LilypadPaymentsTestableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadPaymentsTestableCallerRaw struct {
	Contract *LilypadPaymentsTestableCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadPaymentsTestableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadPaymentsTestableTransactorRaw struct {
	Contract *LilypadPaymentsTestableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadPaymentsTestable creates a new instance of LilypadPaymentsTestable, bound to a specific deployed contract.
func NewLilypadPaymentsTestable(address common.Address, backend bind.ContractBackend) (*LilypadPaymentsTestable, error) {
	contract, err := bindLilypadPaymentsTestable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestable{LilypadPaymentsTestableCaller: LilypadPaymentsTestableCaller{contract: contract}, LilypadPaymentsTestableTransactor: LilypadPaymentsTestableTransactor{contract: contract}, LilypadPaymentsTestableFilterer: LilypadPaymentsTestableFilterer{contract: contract}}, nil
}

// NewLilypadPaymentsTestableCaller creates a new read-only instance of LilypadPaymentsTestable, bound to a specific deployed contract.
func NewLilypadPaymentsTestableCaller(address common.Address, caller bind.ContractCaller) (*LilypadPaymentsTestableCaller, error) {
	contract, err := bindLilypadPaymentsTestable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestableCaller{contract: contract}, nil
}

// NewLilypadPaymentsTestableTransactor creates a new write-only instance of LilypadPaymentsTestable, bound to a specific deployed contract.
func NewLilypadPaymentsTestableTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadPaymentsTestableTransactor, error) {
	contract, err := bindLilypadPaymentsTestable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestableTransactor{contract: contract}, nil
}

// NewLilypadPaymentsTestableFilterer creates a new log filterer instance of LilypadPaymentsTestable, bound to a specific deployed contract.
func NewLilypadPaymentsTestableFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadPaymentsTestableFilterer, error) {
	contract, err := bindLilypadPaymentsTestable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestableFilterer{contract: contract}, nil
}

// bindLilypadPaymentsTestable binds a generic wrapper to an already deployed contract.
func bindLilypadPaymentsTestable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadPaymentsTestableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadPaymentsTestable *LilypadPaymentsTestableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadPaymentsTestable.Contract.LilypadPaymentsTestableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadPaymentsTestable *LilypadPaymentsTestableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.LilypadPaymentsTestableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadPaymentsTestable *LilypadPaymentsTestableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.LilypadPaymentsTestableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadPaymentsTestable *LilypadPaymentsTestableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadPaymentsTestable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadPaymentsTestable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) Owner() (common.Address, error) {
	return _LilypadPaymentsTestable.Contract.Owner(&_LilypadPaymentsTestable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableCallerSession) Owner() (common.Address, error) {
	return _LilypadPaymentsTestable.Contract.Owner(&_LilypadPaymentsTestable.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "acceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) AcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AcceptResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) AcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AcceptResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "addResult", dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) AddResult(dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AddResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) AddResult(dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AddResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "agreeJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) AgreeJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AgreeJobCreator(&_LilypadPaymentsTestable.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) AgreeJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AgreeJobCreator(&_LilypadPaymentsTestable.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "agreeResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) AgreeResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AgreeResourceProvider(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) AgreeResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.AgreeResourceProvider(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "checkResult", dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) CheckResult(dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.CheckResult(&_LilypadPaymentsTestable.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) CheckResult(dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.CheckResult(&_LilypadPaymentsTestable.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.DisableChangeControllerAddress(&_LilypadPaymentsTestable.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.DisableChangeControllerAddress(&_LilypadPaymentsTestable.TransactOpts)
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) DisableChangeTokenAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "disableChangeTokenAddress")
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) DisableChangeTokenAddress() (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.DisableChangeTokenAddress(&_LilypadPaymentsTestable.TransactOpts)
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) DisableChangeTokenAddress() (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.DisableChangeTokenAddress(&_LilypadPaymentsTestable.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.Initialize(&_LilypadPaymentsTestable.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.Initialize(&_LilypadPaymentsTestable.TransactOpts, _tokenAddress)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "mediationAcceptResult", dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) MediationAcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.MediationAcceptResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) MediationAcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.MediationAcceptResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "mediationRejectResult", dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) MediationRejectResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.MediationRejectResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) MediationRejectResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.MediationRejectResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.RenounceOwnership(&_LilypadPaymentsTestable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.RenounceOwnership(&_LilypadPaymentsTestable.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.SetControllerAddress(&_LilypadPaymentsTestable.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.SetControllerAddress(&_LilypadPaymentsTestable.TransactOpts, _controllerAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.SetTokenAddress(&_LilypadPaymentsTestable.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.SetTokenAddress(&_LilypadPaymentsTestable.TransactOpts, _tokenAddress)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x840e9d4e.
//
// Solidity: function timeoutJudgeResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) TimeoutJudgeResults(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "timeoutJudgeResults", dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x840e9d4e.
//
// Solidity: function timeoutJudgeResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) TimeoutJudgeResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TimeoutJudgeResults(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x840e9d4e.
//
// Solidity: function timeoutJudgeResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) TimeoutJudgeResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TimeoutJudgeResults(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "timeoutMediateResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) TimeoutMediateResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TimeoutMediateResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) TimeoutMediateResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TimeoutMediateResult(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x8eed1e9c.
//
// Solidity: function timeoutSubmitResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) TimeoutSubmitResults(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "timeoutSubmitResults", dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x8eed1e9c.
//
// Solidity: function timeoutSubmitResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) TimeoutSubmitResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TimeoutSubmitResults(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x8eed1e9c.
//
// Solidity: function timeoutSubmitResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) TimeoutSubmitResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TimeoutSubmitResults(&_LilypadPaymentsTestable.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TransferOwnership(&_LilypadPaymentsTestable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadPaymentsTestable *LilypadPaymentsTestableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadPaymentsTestable.Contract.TransferOwnership(&_LilypadPaymentsTestable.TransactOpts, newOwner)
}

// LilypadPaymentsTestableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadPaymentsTestable contract.
type LilypadPaymentsTestableInitializedIterator struct {
	Event *LilypadPaymentsTestableInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentsTestableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentsTestableInitialized)
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
		it.Event = new(LilypadPaymentsTestableInitialized)
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
func (it *LilypadPaymentsTestableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentsTestableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentsTestableInitialized represents a Initialized event raised by the LilypadPaymentsTestable contract.
type LilypadPaymentsTestableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadPaymentsTestableInitializedIterator, error) {

	logs, sub, err := _LilypadPaymentsTestable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestableInitializedIterator{contract: _LilypadPaymentsTestable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadPaymentsTestableInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadPaymentsTestable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentsTestableInitialized)
				if err := _LilypadPaymentsTestable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) ParseInitialized(log types.Log) (*LilypadPaymentsTestableInitialized, error) {
	event := new(LilypadPaymentsTestableInitialized)
	if err := _LilypadPaymentsTestable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentsTestableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LilypadPaymentsTestable contract.
type LilypadPaymentsTestableOwnershipTransferredIterator struct {
	Event *LilypadPaymentsTestableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentsTestableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentsTestableOwnershipTransferred)
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
		it.Event = new(LilypadPaymentsTestableOwnershipTransferred)
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
func (it *LilypadPaymentsTestableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentsTestableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentsTestableOwnershipTransferred represents a OwnershipTransferred event raised by the LilypadPaymentsTestable contract.
type LilypadPaymentsTestableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LilypadPaymentsTestableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadPaymentsTestable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestableOwnershipTransferredIterator{contract: _LilypadPaymentsTestable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LilypadPaymentsTestableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadPaymentsTestable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentsTestableOwnershipTransferred)
				if err := _LilypadPaymentsTestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) ParseOwnershipTransferred(log types.Log) (*LilypadPaymentsTestableOwnershipTransferred, error) {
	event := new(LilypadPaymentsTestableOwnershipTransferred)
	if err := _LilypadPaymentsTestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentsTestablePaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the LilypadPaymentsTestable contract.
type LilypadPaymentsTestablePaymentIterator struct {
	Event *LilypadPaymentsTestablePayment // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentsTestablePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentsTestablePayment)
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
		it.Event = new(LilypadPaymentsTestablePayment)
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
func (it *LilypadPaymentsTestablePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentsTestablePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentsTestablePayment represents a Payment event raised by the LilypadPaymentsTestable contract.
type LilypadPaymentsTestablePayment struct {
	DealId    *big.Int
	Payee     common.Address
	Amount    *big.Int
	Reason    uint8
	Direction uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0xb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563.
//
// Solidity: event Payment(uint256 indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) FilterPayment(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadPaymentsTestablePaymentIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadPaymentsTestable.contract.FilterLogs(opts, "Payment", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentsTestablePaymentIterator{contract: _LilypadPaymentsTestable.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0xb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563.
//
// Solidity: event Payment(uint256 indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *LilypadPaymentsTestablePayment, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadPaymentsTestable.contract.WatchLogs(opts, "Payment", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentsTestablePayment)
				if err := _LilypadPaymentsTestable.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0xb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563.
//
// Solidity: event Payment(uint256 indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_LilypadPaymentsTestable *LilypadPaymentsTestableFilterer) ParsePayment(log types.Log) (*LilypadPaymentsTestablePayment, error) {
	event := new(LilypadPaymentsTestablePayment)
	if err := _LilypadPaymentsTestable.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
