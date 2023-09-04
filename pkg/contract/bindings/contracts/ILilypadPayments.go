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

// ILilypadPaymentsMetaData contains all meta data concerning the ILilypadPayments contract.
var ILilypadPaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getEscrowBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILilypadPaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use ILilypadPaymentsMetaData.ABI instead.
var ILilypadPaymentsABI = ILilypadPaymentsMetaData.ABI

// ILilypadPayments is an auto generated Go binding around an Ethereum contract.
type ILilypadPayments struct {
	ILilypadPaymentsCaller     // Read-only binding to the contract
	ILilypadPaymentsTransactor // Write-only binding to the contract
	ILilypadPaymentsFilterer   // Log filterer for contract events
}

// ILilypadPaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILilypadPaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadPaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILilypadPaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadPaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILilypadPaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadPaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILilypadPaymentsSession struct {
	Contract     *ILilypadPayments // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILilypadPaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILilypadPaymentsCallerSession struct {
	Contract *ILilypadPaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ILilypadPaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILilypadPaymentsTransactorSession struct {
	Contract     *ILilypadPaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ILilypadPaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILilypadPaymentsRaw struct {
	Contract *ILilypadPayments // Generic contract binding to access the raw methods on
}

// ILilypadPaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILilypadPaymentsCallerRaw struct {
	Contract *ILilypadPaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// ILilypadPaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILilypadPaymentsTransactorRaw struct {
	Contract *ILilypadPaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILilypadPayments creates a new instance of ILilypadPayments, bound to a specific deployed contract.
func NewILilypadPayments(address common.Address, backend bind.ContractBackend) (*ILilypadPayments, error) {
	contract, err := bindILilypadPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILilypadPayments{ILilypadPaymentsCaller: ILilypadPaymentsCaller{contract: contract}, ILilypadPaymentsTransactor: ILilypadPaymentsTransactor{contract: contract}, ILilypadPaymentsFilterer: ILilypadPaymentsFilterer{contract: contract}}, nil
}

// NewILilypadPaymentsCaller creates a new read-only instance of ILilypadPayments, bound to a specific deployed contract.
func NewILilypadPaymentsCaller(address common.Address, caller bind.ContractCaller) (*ILilypadPaymentsCaller, error) {
	contract, err := bindILilypadPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILilypadPaymentsCaller{contract: contract}, nil
}

// NewILilypadPaymentsTransactor creates a new write-only instance of ILilypadPayments, bound to a specific deployed contract.
func NewILilypadPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*ILilypadPaymentsTransactor, error) {
	contract, err := bindILilypadPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILilypadPaymentsTransactor{contract: contract}, nil
}

// NewILilypadPaymentsFilterer creates a new log filterer instance of ILilypadPayments, bound to a specific deployed contract.
func NewILilypadPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*ILilypadPaymentsFilterer, error) {
	contract, err := bindILilypadPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILilypadPaymentsFilterer{contract: contract}, nil
}

// bindILilypadPayments binds a generic wrapper to an already deployed contract.
func bindILilypadPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILilypadPaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILilypadPayments *ILilypadPaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILilypadPayments.Contract.ILilypadPaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILilypadPayments *ILilypadPaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.ILilypadPaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILilypadPayments *ILilypadPaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.ILilypadPaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILilypadPayments *ILilypadPaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILilypadPayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILilypadPayments *ILilypadPaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILilypadPayments *ILilypadPaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.contract.Transact(opts, method, params...)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "acceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) AcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AcceptResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) AcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AcceptResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "addResult", dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) AddResult(dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AddResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) AddResult(dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AddResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "agreeJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) AgreeJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AgreeJobCreator(&_ILilypadPayments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) AgreeJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AgreeJobCreator(&_ILilypadPayments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "agreeResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) AgreeResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AgreeResourceProvider(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) AgreeResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.AgreeResourceProvider(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "checkResult", dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) CheckResult(dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.CheckResult(&_ILilypadPayments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) CheckResult(dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.CheckResult(&_ILilypadPayments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// GetEscrowBalance is a paid mutator transaction binding the contract method 0x6374b11b.
//
// Solidity: function getEscrowBalance(address _tokenAddress) returns(uint256)
func (_ILilypadPayments *ILilypadPaymentsTransactor) GetEscrowBalance(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "getEscrowBalance", _tokenAddress)
}

// GetEscrowBalance is a paid mutator transaction binding the contract method 0x6374b11b.
//
// Solidity: function getEscrowBalance(address _tokenAddress) returns(uint256)
func (_ILilypadPayments *ILilypadPaymentsSession) GetEscrowBalance(_tokenAddress common.Address) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.GetEscrowBalance(&_ILilypadPayments.TransactOpts, _tokenAddress)
}

// GetEscrowBalance is a paid mutator transaction binding the contract method 0x6374b11b.
//
// Solidity: function getEscrowBalance(address _tokenAddress) returns(uint256)
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) GetEscrowBalance(_tokenAddress common.Address) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.GetEscrowBalance(&_ILilypadPayments.TransactOpts, _tokenAddress)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "mediationAcceptResult", dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) MediationAcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.MediationAcceptResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) MediationAcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.MediationAcceptResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "mediationRejectResult", dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) MediationRejectResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.MediationRejectResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) MediationRejectResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.MediationRejectResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.SetTokenAddress(&_ILilypadPayments.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.SetTokenAddress(&_ILilypadPayments.TransactOpts, _tokenAddress)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x4e511a47.
//
// Solidity: function timeoutJudgeResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "timeoutJudgeResult", dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x4e511a47.
//
// Solidity: function timeoutJudgeResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) TimeoutJudgeResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.TimeoutJudgeResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x4e511a47.
//
// Solidity: function timeoutJudgeResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) TimeoutJudgeResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.TimeoutJudgeResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "timeoutMediateResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) TimeoutMediateResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.TimeoutMediateResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) TimeoutMediateResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.TimeoutMediateResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0xeb7e9394.
//
// Solidity: function timeoutSubmitResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.contract.Transact(opts, "timeoutSubmitResult", dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0xeb7e9394.
//
// Solidity: function timeoutSubmitResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsSession) TimeoutSubmitResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.TimeoutSubmitResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0xeb7e9394.
//
// Solidity: function timeoutSubmitResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_ILilypadPayments *ILilypadPaymentsTransactorSession) TimeoutSubmitResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _ILilypadPayments.Contract.TimeoutSubmitResult(&_ILilypadPayments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}
