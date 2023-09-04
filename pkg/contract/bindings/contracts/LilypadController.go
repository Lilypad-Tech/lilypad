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

// SharedStructsAgreement is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsAgreement struct {
	State                    uint8
	Mediator                 common.Address
	ResourceProviderAgreedAt *big.Int
	JobCreatorAgreedAt       *big.Int
	DealAgreedAt             *big.Int
	ResultsSubmittedAt       *big.Int
	ResultsAcceptedAt        *big.Int
	ResultsCheckedAt         *big.Int
	MediationAcceptedAt      *big.Int
	MediationRejectedAt      *big.Int
	TimeoutSubmitResultsAt   *big.Int
	TimeoutJudgeResultsAt    *big.Int
	TimeoutMediateResultsAt  *big.Int
}

// LilypadControllerMetaData contains all meta data concerning the LilypadController contract.
var LilypadControllerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"DealAgreed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"JobCreatorAgreed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"MediationAcceptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"MediationRejectResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"ResourceProviderAgreed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"ResultAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"ResultAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"ResultChecked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"TimeoutJudgeResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"TimeoutMediateResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"TimeoutSubmitResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"agree\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentsAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentsAddress\",\"type\":\"address\"}],\"name\":\"setPaymentsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadControllerMetaData.ABI instead.
var LilypadControllerABI = LilypadControllerMetaData.ABI

// LilypadController is an auto generated Go binding around an Ethereum contract.
type LilypadController struct {
	LilypadControllerCaller     // Read-only binding to the contract
	LilypadControllerTransactor // Write-only binding to the contract
	LilypadControllerFilterer   // Log filterer for contract events
}

// LilypadControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadControllerSession struct {
	Contract     *LilypadController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LilypadControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadControllerCallerSession struct {
	Contract *LilypadControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LilypadControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadControllerTransactorSession struct {
	Contract     *LilypadControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LilypadControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadControllerRaw struct {
	Contract *LilypadController // Generic contract binding to access the raw methods on
}

// LilypadControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadControllerCallerRaw struct {
	Contract *LilypadControllerCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadControllerTransactorRaw struct {
	Contract *LilypadControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadController creates a new instance of LilypadController, bound to a specific deployed contract.
func NewLilypadController(address common.Address, backend bind.ContractBackend) (*LilypadController, error) {
	contract, err := bindLilypadController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadController{LilypadControllerCaller: LilypadControllerCaller{contract: contract}, LilypadControllerTransactor: LilypadControllerTransactor{contract: contract}, LilypadControllerFilterer: LilypadControllerFilterer{contract: contract}}, nil
}

// NewLilypadControllerCaller creates a new read-only instance of LilypadController, bound to a specific deployed contract.
func NewLilypadControllerCaller(address common.Address, caller bind.ContractCaller) (*LilypadControllerCaller, error) {
	contract, err := bindLilypadController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerCaller{contract: contract}, nil
}

// NewLilypadControllerTransactor creates a new write-only instance of LilypadController, bound to a specific deployed contract.
func NewLilypadControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadControllerTransactor, error) {
	contract, err := bindLilypadController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerTransactor{contract: contract}, nil
}

// NewLilypadControllerFilterer creates a new log filterer instance of LilypadController, bound to a specific deployed contract.
func NewLilypadControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadControllerFilterer, error) {
	contract, err := bindLilypadController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerFilterer{contract: contract}, nil
}

// bindLilypadController binds a generic wrapper to an already deployed contract.
func bindLilypadController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadController *LilypadControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadController.Contract.LilypadControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadController *LilypadControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadController.Contract.LilypadControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadController *LilypadControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadController.Contract.LilypadControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadController *LilypadControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadController *LilypadControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadController *LilypadControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadController.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadController *LilypadControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadController *LilypadControllerSession) Owner() (common.Address, error) {
	return _LilypadController.Contract.Owner(&_LilypadController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadController *LilypadControllerCallerSession) Owner() (common.Address, error) {
	return _LilypadController.Contract.Owner(&_LilypadController.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.AcceptResult(&_LilypadController.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactorSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.AcceptResult(&_LilypadController.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns()
func (_LilypadController *LilypadControllerTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns()
func (_LilypadController *LilypadControllerSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.AddResult(&_LilypadController.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns()
func (_LilypadController *LilypadControllerTransactorSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.AddResult(&_LilypadController.TransactOpts, dealId, resultsId, instructionCount)
}

// Agree is a paid mutator transaction binding the contract method 0xb3857a88.
//
// Solidity: function agree(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadController *LilypadControllerTransactor) Agree(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "agree", dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// Agree is a paid mutator transaction binding the contract method 0xb3857a88.
//
// Solidity: function agree(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadController *LilypadControllerSession) Agree(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.Agree(&_LilypadController.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// Agree is a paid mutator transaction binding the contract method 0xb3857a88.
//
// Solidity: function agree(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadController *LilypadControllerTransactorSession) Agree(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.Agree(&_LilypadController.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadController *LilypadControllerTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadController *LilypadControllerSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.CheckResult(&_LilypadController.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadController *LilypadControllerTransactorSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.CheckResult(&_LilypadController.TransactOpts, dealId, mediator)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _storageAddress, address _paymentsAddress) returns()
func (_LilypadController *LilypadControllerTransactor) Initialize(opts *bind.TransactOpts, _storageAddress common.Address, _paymentsAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "initialize", _storageAddress, _paymentsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _storageAddress, address _paymentsAddress) returns()
func (_LilypadController *LilypadControllerSession) Initialize(_storageAddress common.Address, _paymentsAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.Initialize(&_LilypadController.TransactOpts, _storageAddress, _paymentsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _storageAddress, address _paymentsAddress) returns()
func (_LilypadController *LilypadControllerTransactorSession) Initialize(_storageAddress common.Address, _paymentsAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.Initialize(&_LilypadController.TransactOpts, _storageAddress, _paymentsAddress)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.MediationAcceptResult(&_LilypadController.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.MediationAcceptResult(&_LilypadController.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.MediationRejectResult(&_LilypadController.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.MediationRejectResult(&_LilypadController.TransactOpts, dealId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadController *LilypadControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadController *LilypadControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadController.Contract.RenounceOwnership(&_LilypadController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadController *LilypadControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadController.Contract.RenounceOwnership(&_LilypadController.TransactOpts)
}

// SetPaymentsAddress is a paid mutator transaction binding the contract method 0x640e570f.
//
// Solidity: function setPaymentsAddress(address _paymentsAddress) returns()
func (_LilypadController *LilypadControllerTransactor) SetPaymentsAddress(opts *bind.TransactOpts, _paymentsAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "setPaymentsAddress", _paymentsAddress)
}

// SetPaymentsAddress is a paid mutator transaction binding the contract method 0x640e570f.
//
// Solidity: function setPaymentsAddress(address _paymentsAddress) returns()
func (_LilypadController *LilypadControllerSession) SetPaymentsAddress(_paymentsAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.SetPaymentsAddress(&_LilypadController.TransactOpts, _paymentsAddress)
}

// SetPaymentsAddress is a paid mutator transaction binding the contract method 0x640e570f.
//
// Solidity: function setPaymentsAddress(address _paymentsAddress) returns()
func (_LilypadController *LilypadControllerTransactorSession) SetPaymentsAddress(_paymentsAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.SetPaymentsAddress(&_LilypadController.TransactOpts, _paymentsAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(address _storageAddress) returns()
func (_LilypadController *LilypadControllerTransactor) SetStorageAddress(opts *bind.TransactOpts, _storageAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "setStorageAddress", _storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(address _storageAddress) returns()
func (_LilypadController *LilypadControllerSession) SetStorageAddress(_storageAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.SetStorageAddress(&_LilypadController.TransactOpts, _storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(address _storageAddress) returns()
func (_LilypadController *LilypadControllerTransactorSession) SetStorageAddress(_storageAddress common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.SetStorageAddress(&_LilypadController.TransactOpts, _storageAddress)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.TimeoutJudgeResult(&_LilypadController.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactorSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.TimeoutJudgeResult(&_LilypadController.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.TimeoutMediateResult(&_LilypadController.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactorSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.TimeoutMediateResult(&_LilypadController.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.TimeoutSubmitResult(&_LilypadController.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadController *LilypadControllerTransactorSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadController.Contract.TimeoutSubmitResult(&_LilypadController.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadController *LilypadControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LilypadController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadController *LilypadControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.TransferOwnership(&_LilypadController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadController *LilypadControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadController.Contract.TransferOwnership(&_LilypadController.TransactOpts, newOwner)
}

// LilypadControllerDealAgreedIterator is returned from FilterDealAgreed and is used to iterate over the raw logs and unpacked data for DealAgreed events raised by the LilypadController contract.
type LilypadControllerDealAgreedIterator struct {
	Event *LilypadControllerDealAgreed // Event containing the contract specifics and raw log

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
func (it *LilypadControllerDealAgreedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerDealAgreed)
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
		it.Event = new(LilypadControllerDealAgreed)
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
func (it *LilypadControllerDealAgreedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerDealAgreedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerDealAgreed represents a DealAgreed event raised by the LilypadController contract.
type LilypadControllerDealAgreed struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealAgreed is a free log retrieval operation binding the contract event 0x5d44329b191ed1b94788d02cec3bed307a47286c27dfee6592827a8619d1129b.
//
// Solidity: event DealAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterDealAgreed(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerDealAgreedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "DealAgreed", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerDealAgreedIterator{contract: _LilypadController.contract, event: "DealAgreed", logs: logs, sub: sub}, nil
}

// WatchDealAgreed is a free log subscription operation binding the contract event 0x5d44329b191ed1b94788d02cec3bed307a47286c27dfee6592827a8619d1129b.
//
// Solidity: event DealAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchDealAgreed(opts *bind.WatchOpts, sink chan<- *LilypadControllerDealAgreed, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "DealAgreed", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerDealAgreed)
				if err := _LilypadController.contract.UnpackLog(event, "DealAgreed", log); err != nil {
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

// ParseDealAgreed is a log parse operation binding the contract event 0x5d44329b191ed1b94788d02cec3bed307a47286c27dfee6592827a8619d1129b.
//
// Solidity: event DealAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseDealAgreed(log types.Log) (*LilypadControllerDealAgreed, error) {
	event := new(LilypadControllerDealAgreed)
	if err := _LilypadController.contract.UnpackLog(event, "DealAgreed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadController contract.
type LilypadControllerInitializedIterator struct {
	Event *LilypadControllerInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerInitialized)
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
		it.Event = new(LilypadControllerInitialized)
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
func (it *LilypadControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerInitialized represents a Initialized event raised by the LilypadController contract.
type LilypadControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadController *LilypadControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadControllerInitializedIterator, error) {

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadControllerInitializedIterator{contract: _LilypadController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadController *LilypadControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerInitialized)
				if err := _LilypadController.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadController *LilypadControllerFilterer) ParseInitialized(log types.Log) (*LilypadControllerInitialized, error) {
	event := new(LilypadControllerInitialized)
	if err := _LilypadController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerJobCreatorAgreedIterator is returned from FilterJobCreatorAgreed and is used to iterate over the raw logs and unpacked data for JobCreatorAgreed events raised by the LilypadController contract.
type LilypadControllerJobCreatorAgreedIterator struct {
	Event *LilypadControllerJobCreatorAgreed // Event containing the contract specifics and raw log

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
func (it *LilypadControllerJobCreatorAgreedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerJobCreatorAgreed)
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
		it.Event = new(LilypadControllerJobCreatorAgreed)
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
func (it *LilypadControllerJobCreatorAgreedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerJobCreatorAgreedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerJobCreatorAgreed represents a JobCreatorAgreed event raised by the LilypadController contract.
type LilypadControllerJobCreatorAgreed struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJobCreatorAgreed is a free log retrieval operation binding the contract event 0x8e431afeb922a853c88906102d5cdd0f5d38dac2bd14ae62e32a1f6592b69c14.
//
// Solidity: event JobCreatorAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterJobCreatorAgreed(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerJobCreatorAgreedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "JobCreatorAgreed", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerJobCreatorAgreedIterator{contract: _LilypadController.contract, event: "JobCreatorAgreed", logs: logs, sub: sub}, nil
}

// WatchJobCreatorAgreed is a free log subscription operation binding the contract event 0x8e431afeb922a853c88906102d5cdd0f5d38dac2bd14ae62e32a1f6592b69c14.
//
// Solidity: event JobCreatorAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchJobCreatorAgreed(opts *bind.WatchOpts, sink chan<- *LilypadControllerJobCreatorAgreed, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "JobCreatorAgreed", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerJobCreatorAgreed)
				if err := _LilypadController.contract.UnpackLog(event, "JobCreatorAgreed", log); err != nil {
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

// ParseJobCreatorAgreed is a log parse operation binding the contract event 0x8e431afeb922a853c88906102d5cdd0f5d38dac2bd14ae62e32a1f6592b69c14.
//
// Solidity: event JobCreatorAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseJobCreatorAgreed(log types.Log) (*LilypadControllerJobCreatorAgreed, error) {
	event := new(LilypadControllerJobCreatorAgreed)
	if err := _LilypadController.contract.UnpackLog(event, "JobCreatorAgreed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerMediationAcceptResultIterator is returned from FilterMediationAcceptResult and is used to iterate over the raw logs and unpacked data for MediationAcceptResult events raised by the LilypadController contract.
type LilypadControllerMediationAcceptResultIterator struct {
	Event *LilypadControllerMediationAcceptResult // Event containing the contract specifics and raw log

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
func (it *LilypadControllerMediationAcceptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerMediationAcceptResult)
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
		it.Event = new(LilypadControllerMediationAcceptResult)
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
func (it *LilypadControllerMediationAcceptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerMediationAcceptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerMediationAcceptResult represents a MediationAcceptResult event raised by the LilypadController contract.
type LilypadControllerMediationAcceptResult struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMediationAcceptResult is a free log retrieval operation binding the contract event 0x6869c4fbe3662f8188805783ad034002c0317a267b33e571b21803e0a82c222b.
//
// Solidity: event MediationAcceptResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterMediationAcceptResult(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerMediationAcceptResultIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "MediationAcceptResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerMediationAcceptResultIterator{contract: _LilypadController.contract, event: "MediationAcceptResult", logs: logs, sub: sub}, nil
}

// WatchMediationAcceptResult is a free log subscription operation binding the contract event 0x6869c4fbe3662f8188805783ad034002c0317a267b33e571b21803e0a82c222b.
//
// Solidity: event MediationAcceptResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchMediationAcceptResult(opts *bind.WatchOpts, sink chan<- *LilypadControllerMediationAcceptResult, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "MediationAcceptResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerMediationAcceptResult)
				if err := _LilypadController.contract.UnpackLog(event, "MediationAcceptResult", log); err != nil {
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

// ParseMediationAcceptResult is a log parse operation binding the contract event 0x6869c4fbe3662f8188805783ad034002c0317a267b33e571b21803e0a82c222b.
//
// Solidity: event MediationAcceptResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseMediationAcceptResult(log types.Log) (*LilypadControllerMediationAcceptResult, error) {
	event := new(LilypadControllerMediationAcceptResult)
	if err := _LilypadController.contract.UnpackLog(event, "MediationAcceptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerMediationRejectResultIterator is returned from FilterMediationRejectResult and is used to iterate over the raw logs and unpacked data for MediationRejectResult events raised by the LilypadController contract.
type LilypadControllerMediationRejectResultIterator struct {
	Event *LilypadControllerMediationRejectResult // Event containing the contract specifics and raw log

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
func (it *LilypadControllerMediationRejectResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerMediationRejectResult)
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
		it.Event = new(LilypadControllerMediationRejectResult)
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
func (it *LilypadControllerMediationRejectResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerMediationRejectResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerMediationRejectResult represents a MediationRejectResult event raised by the LilypadController contract.
type LilypadControllerMediationRejectResult struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMediationRejectResult is a free log retrieval operation binding the contract event 0x2be5a64d41d363eaaa7b5195947407dde33371bff56ead257396804137b11d35.
//
// Solidity: event MediationRejectResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterMediationRejectResult(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerMediationRejectResultIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "MediationRejectResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerMediationRejectResultIterator{contract: _LilypadController.contract, event: "MediationRejectResult", logs: logs, sub: sub}, nil
}

// WatchMediationRejectResult is a free log subscription operation binding the contract event 0x2be5a64d41d363eaaa7b5195947407dde33371bff56ead257396804137b11d35.
//
// Solidity: event MediationRejectResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchMediationRejectResult(opts *bind.WatchOpts, sink chan<- *LilypadControllerMediationRejectResult, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "MediationRejectResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerMediationRejectResult)
				if err := _LilypadController.contract.UnpackLog(event, "MediationRejectResult", log); err != nil {
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

// ParseMediationRejectResult is a log parse operation binding the contract event 0x2be5a64d41d363eaaa7b5195947407dde33371bff56ead257396804137b11d35.
//
// Solidity: event MediationRejectResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseMediationRejectResult(log types.Log) (*LilypadControllerMediationRejectResult, error) {
	event := new(LilypadControllerMediationRejectResult)
	if err := _LilypadController.contract.UnpackLog(event, "MediationRejectResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LilypadController contract.
type LilypadControllerOwnershipTransferredIterator struct {
	Event *LilypadControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LilypadControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerOwnershipTransferred)
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
		it.Event = new(LilypadControllerOwnershipTransferred)
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
func (it *LilypadControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerOwnershipTransferred represents a OwnershipTransferred event raised by the LilypadController contract.
type LilypadControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadController *LilypadControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LilypadControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerOwnershipTransferredIterator{contract: _LilypadController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadController *LilypadControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LilypadControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerOwnershipTransferred)
				if err := _LilypadController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LilypadController *LilypadControllerFilterer) ParseOwnershipTransferred(log types.Log) (*LilypadControllerOwnershipTransferred, error) {
	event := new(LilypadControllerOwnershipTransferred)
	if err := _LilypadController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerResourceProviderAgreedIterator is returned from FilterResourceProviderAgreed and is used to iterate over the raw logs and unpacked data for ResourceProviderAgreed events raised by the LilypadController contract.
type LilypadControllerResourceProviderAgreedIterator struct {
	Event *LilypadControllerResourceProviderAgreed // Event containing the contract specifics and raw log

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
func (it *LilypadControllerResourceProviderAgreedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerResourceProviderAgreed)
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
		it.Event = new(LilypadControllerResourceProviderAgreed)
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
func (it *LilypadControllerResourceProviderAgreedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerResourceProviderAgreedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerResourceProviderAgreed represents a ResourceProviderAgreed event raised by the LilypadController contract.
type LilypadControllerResourceProviderAgreed struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResourceProviderAgreed is a free log retrieval operation binding the contract event 0x26a8f34da4d873e21cd45615776e6b843f82381ff094e72c5145f7a39466158f.
//
// Solidity: event ResourceProviderAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterResourceProviderAgreed(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerResourceProviderAgreedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "ResourceProviderAgreed", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerResourceProviderAgreedIterator{contract: _LilypadController.contract, event: "ResourceProviderAgreed", logs: logs, sub: sub}, nil
}

// WatchResourceProviderAgreed is a free log subscription operation binding the contract event 0x26a8f34da4d873e21cd45615776e6b843f82381ff094e72c5145f7a39466158f.
//
// Solidity: event ResourceProviderAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchResourceProviderAgreed(opts *bind.WatchOpts, sink chan<- *LilypadControllerResourceProviderAgreed, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "ResourceProviderAgreed", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerResourceProviderAgreed)
				if err := _LilypadController.contract.UnpackLog(event, "ResourceProviderAgreed", log); err != nil {
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

// ParseResourceProviderAgreed is a log parse operation binding the contract event 0x26a8f34da4d873e21cd45615776e6b843f82381ff094e72c5145f7a39466158f.
//
// Solidity: event ResourceProviderAgreed(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseResourceProviderAgreed(log types.Log) (*LilypadControllerResourceProviderAgreed, error) {
	event := new(LilypadControllerResourceProviderAgreed)
	if err := _LilypadController.contract.UnpackLog(event, "ResourceProviderAgreed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerResultAcceptedIterator is returned from FilterResultAccepted and is used to iterate over the raw logs and unpacked data for ResultAccepted events raised by the LilypadController contract.
type LilypadControllerResultAcceptedIterator struct {
	Event *LilypadControllerResultAccepted // Event containing the contract specifics and raw log

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
func (it *LilypadControllerResultAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerResultAccepted)
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
		it.Event = new(LilypadControllerResultAccepted)
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
func (it *LilypadControllerResultAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerResultAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerResultAccepted represents a ResultAccepted event raised by the LilypadController contract.
type LilypadControllerResultAccepted struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResultAccepted is a free log retrieval operation binding the contract event 0x5bdb4aa0a0647d4311e04dd653def3a11ed762cab8cd7eb1465046687b0018a3.
//
// Solidity: event ResultAccepted(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterResultAccepted(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerResultAcceptedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "ResultAccepted", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerResultAcceptedIterator{contract: _LilypadController.contract, event: "ResultAccepted", logs: logs, sub: sub}, nil
}

// WatchResultAccepted is a free log subscription operation binding the contract event 0x5bdb4aa0a0647d4311e04dd653def3a11ed762cab8cd7eb1465046687b0018a3.
//
// Solidity: event ResultAccepted(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchResultAccepted(opts *bind.WatchOpts, sink chan<- *LilypadControllerResultAccepted, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "ResultAccepted", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerResultAccepted)
				if err := _LilypadController.contract.UnpackLog(event, "ResultAccepted", log); err != nil {
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

// ParseResultAccepted is a log parse operation binding the contract event 0x5bdb4aa0a0647d4311e04dd653def3a11ed762cab8cd7eb1465046687b0018a3.
//
// Solidity: event ResultAccepted(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseResultAccepted(log types.Log) (*LilypadControllerResultAccepted, error) {
	event := new(LilypadControllerResultAccepted)
	if err := _LilypadController.contract.UnpackLog(event, "ResultAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerResultAddedIterator is returned from FilterResultAdded and is used to iterate over the raw logs and unpacked data for ResultAdded events raised by the LilypadController contract.
type LilypadControllerResultAddedIterator struct {
	Event *LilypadControllerResultAdded // Event containing the contract specifics and raw log

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
func (it *LilypadControllerResultAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerResultAdded)
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
		it.Event = new(LilypadControllerResultAdded)
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
func (it *LilypadControllerResultAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerResultAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerResultAdded represents a ResultAdded event raised by the LilypadController contract.
type LilypadControllerResultAdded struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResultAdded is a free log retrieval operation binding the contract event 0x92304b2ca1800e2e7b33e9d20a5da5822623e9a982f53a496597185a1abaaf35.
//
// Solidity: event ResultAdded(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterResultAdded(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerResultAddedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "ResultAdded", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerResultAddedIterator{contract: _LilypadController.contract, event: "ResultAdded", logs: logs, sub: sub}, nil
}

// WatchResultAdded is a free log subscription operation binding the contract event 0x92304b2ca1800e2e7b33e9d20a5da5822623e9a982f53a496597185a1abaaf35.
//
// Solidity: event ResultAdded(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchResultAdded(opts *bind.WatchOpts, sink chan<- *LilypadControllerResultAdded, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "ResultAdded", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerResultAdded)
				if err := _LilypadController.contract.UnpackLog(event, "ResultAdded", log); err != nil {
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

// ParseResultAdded is a log parse operation binding the contract event 0x92304b2ca1800e2e7b33e9d20a5da5822623e9a982f53a496597185a1abaaf35.
//
// Solidity: event ResultAdded(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseResultAdded(log types.Log) (*LilypadControllerResultAdded, error) {
	event := new(LilypadControllerResultAdded)
	if err := _LilypadController.contract.UnpackLog(event, "ResultAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerResultCheckedIterator is returned from FilterResultChecked and is used to iterate over the raw logs and unpacked data for ResultChecked events raised by the LilypadController contract.
type LilypadControllerResultCheckedIterator struct {
	Event *LilypadControllerResultChecked // Event containing the contract specifics and raw log

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
func (it *LilypadControllerResultCheckedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerResultChecked)
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
		it.Event = new(LilypadControllerResultChecked)
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
func (it *LilypadControllerResultCheckedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerResultCheckedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerResultChecked represents a ResultChecked event raised by the LilypadController contract.
type LilypadControllerResultChecked struct {
	DealId   *big.Int
	Mediator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResultChecked is a free log retrieval operation binding the contract event 0x074f78619102b3535ae08a4db24ea45c911a67d55d987fea03a839676c7b57e8.
//
// Solidity: event ResultChecked(uint256 indexed dealId, address indexed mediator)
func (_LilypadController *LilypadControllerFilterer) FilterResultChecked(opts *bind.FilterOpts, dealId []*big.Int, mediator []common.Address) (*LilypadControllerResultCheckedIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var mediatorRule []interface{}
	for _, mediatorItem := range mediator {
		mediatorRule = append(mediatorRule, mediatorItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "ResultChecked", dealIdRule, mediatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerResultCheckedIterator{contract: _LilypadController.contract, event: "ResultChecked", logs: logs, sub: sub}, nil
}

// WatchResultChecked is a free log subscription operation binding the contract event 0x074f78619102b3535ae08a4db24ea45c911a67d55d987fea03a839676c7b57e8.
//
// Solidity: event ResultChecked(uint256 indexed dealId, address indexed mediator)
func (_LilypadController *LilypadControllerFilterer) WatchResultChecked(opts *bind.WatchOpts, sink chan<- *LilypadControllerResultChecked, dealId []*big.Int, mediator []common.Address) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var mediatorRule []interface{}
	for _, mediatorItem := range mediator {
		mediatorRule = append(mediatorRule, mediatorItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "ResultChecked", dealIdRule, mediatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerResultChecked)
				if err := _LilypadController.contract.UnpackLog(event, "ResultChecked", log); err != nil {
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

// ParseResultChecked is a log parse operation binding the contract event 0x074f78619102b3535ae08a4db24ea45c911a67d55d987fea03a839676c7b57e8.
//
// Solidity: event ResultChecked(uint256 indexed dealId, address indexed mediator)
func (_LilypadController *LilypadControllerFilterer) ParseResultChecked(log types.Log) (*LilypadControllerResultChecked, error) {
	event := new(LilypadControllerResultChecked)
	if err := _LilypadController.contract.UnpackLog(event, "ResultChecked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerTimeoutJudgeResultIterator is returned from FilterTimeoutJudgeResult and is used to iterate over the raw logs and unpacked data for TimeoutJudgeResult events raised by the LilypadController contract.
type LilypadControllerTimeoutJudgeResultIterator struct {
	Event *LilypadControllerTimeoutJudgeResult // Event containing the contract specifics and raw log

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
func (it *LilypadControllerTimeoutJudgeResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerTimeoutJudgeResult)
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
		it.Event = new(LilypadControllerTimeoutJudgeResult)
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
func (it *LilypadControllerTimeoutJudgeResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerTimeoutJudgeResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerTimeoutJudgeResult represents a TimeoutJudgeResult event raised by the LilypadController contract.
type LilypadControllerTimeoutJudgeResult struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTimeoutJudgeResult is a free log retrieval operation binding the contract event 0x6d119d722b70ecb67111e4b5501842b0d67693c9f093672d604d39e3ae955504.
//
// Solidity: event TimeoutJudgeResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterTimeoutJudgeResult(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerTimeoutJudgeResultIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "TimeoutJudgeResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerTimeoutJudgeResultIterator{contract: _LilypadController.contract, event: "TimeoutJudgeResult", logs: logs, sub: sub}, nil
}

// WatchTimeoutJudgeResult is a free log subscription operation binding the contract event 0x6d119d722b70ecb67111e4b5501842b0d67693c9f093672d604d39e3ae955504.
//
// Solidity: event TimeoutJudgeResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchTimeoutJudgeResult(opts *bind.WatchOpts, sink chan<- *LilypadControllerTimeoutJudgeResult, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "TimeoutJudgeResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerTimeoutJudgeResult)
				if err := _LilypadController.contract.UnpackLog(event, "TimeoutJudgeResult", log); err != nil {
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

// ParseTimeoutJudgeResult is a log parse operation binding the contract event 0x6d119d722b70ecb67111e4b5501842b0d67693c9f093672d604d39e3ae955504.
//
// Solidity: event TimeoutJudgeResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseTimeoutJudgeResult(log types.Log) (*LilypadControllerTimeoutJudgeResult, error) {
	event := new(LilypadControllerTimeoutJudgeResult)
	if err := _LilypadController.contract.UnpackLog(event, "TimeoutJudgeResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerTimeoutMediateResultIterator is returned from FilterTimeoutMediateResult and is used to iterate over the raw logs and unpacked data for TimeoutMediateResult events raised by the LilypadController contract.
type LilypadControllerTimeoutMediateResultIterator struct {
	Event *LilypadControllerTimeoutMediateResult // Event containing the contract specifics and raw log

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
func (it *LilypadControllerTimeoutMediateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerTimeoutMediateResult)
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
		it.Event = new(LilypadControllerTimeoutMediateResult)
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
func (it *LilypadControllerTimeoutMediateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerTimeoutMediateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerTimeoutMediateResult represents a TimeoutMediateResult event raised by the LilypadController contract.
type LilypadControllerTimeoutMediateResult struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTimeoutMediateResult is a free log retrieval operation binding the contract event 0x89e0dd5728e1ea979c4292369726e7d04659e09cc314290d30caa1674be96f56.
//
// Solidity: event TimeoutMediateResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterTimeoutMediateResult(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerTimeoutMediateResultIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "TimeoutMediateResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerTimeoutMediateResultIterator{contract: _LilypadController.contract, event: "TimeoutMediateResult", logs: logs, sub: sub}, nil
}

// WatchTimeoutMediateResult is a free log subscription operation binding the contract event 0x89e0dd5728e1ea979c4292369726e7d04659e09cc314290d30caa1674be96f56.
//
// Solidity: event TimeoutMediateResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchTimeoutMediateResult(opts *bind.WatchOpts, sink chan<- *LilypadControllerTimeoutMediateResult, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "TimeoutMediateResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerTimeoutMediateResult)
				if err := _LilypadController.contract.UnpackLog(event, "TimeoutMediateResult", log); err != nil {
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

// ParseTimeoutMediateResult is a log parse operation binding the contract event 0x89e0dd5728e1ea979c4292369726e7d04659e09cc314290d30caa1674be96f56.
//
// Solidity: event TimeoutMediateResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseTimeoutMediateResult(log types.Log) (*LilypadControllerTimeoutMediateResult, error) {
	event := new(LilypadControllerTimeoutMediateResult)
	if err := _LilypadController.contract.UnpackLog(event, "TimeoutMediateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadControllerTimeoutSubmitResultIterator is returned from FilterTimeoutSubmitResult and is used to iterate over the raw logs and unpacked data for TimeoutSubmitResult events raised by the LilypadController contract.
type LilypadControllerTimeoutSubmitResultIterator struct {
	Event *LilypadControllerTimeoutSubmitResult // Event containing the contract specifics and raw log

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
func (it *LilypadControllerTimeoutSubmitResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadControllerTimeoutSubmitResult)
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
		it.Event = new(LilypadControllerTimeoutSubmitResult)
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
func (it *LilypadControllerTimeoutSubmitResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadControllerTimeoutSubmitResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadControllerTimeoutSubmitResult represents a TimeoutSubmitResult event raised by the LilypadController contract.
type LilypadControllerTimeoutSubmitResult struct {
	DealId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTimeoutSubmitResult is a free log retrieval operation binding the contract event 0xbc228c937af4086388ee0005a68438286f19bbaa17d16ca9fa1b4361a918ee42.
//
// Solidity: event TimeoutSubmitResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) FilterTimeoutSubmitResult(opts *bind.FilterOpts, dealId []*big.Int) (*LilypadControllerTimeoutSubmitResultIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.FilterLogs(opts, "TimeoutSubmitResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadControllerTimeoutSubmitResultIterator{contract: _LilypadController.contract, event: "TimeoutSubmitResult", logs: logs, sub: sub}, nil
}

// WatchTimeoutSubmitResult is a free log subscription operation binding the contract event 0xbc228c937af4086388ee0005a68438286f19bbaa17d16ca9fa1b4361a918ee42.
//
// Solidity: event TimeoutSubmitResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) WatchTimeoutSubmitResult(opts *bind.WatchOpts, sink chan<- *LilypadControllerTimeoutSubmitResult, dealId []*big.Int) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadController.contract.WatchLogs(opts, "TimeoutSubmitResult", dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadControllerTimeoutSubmitResult)
				if err := _LilypadController.contract.UnpackLog(event, "TimeoutSubmitResult", log); err != nil {
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

// ParseTimeoutSubmitResult is a log parse operation binding the contract event 0xbc228c937af4086388ee0005a68438286f19bbaa17d16ca9fa1b4361a918ee42.
//
// Solidity: event TimeoutSubmitResult(uint256 indexed dealId)
func (_LilypadController *LilypadControllerFilterer) ParseTimeoutSubmitResult(log types.Log) (*LilypadControllerTimeoutSubmitResult, error) {
	event := new(LilypadControllerTimeoutSubmitResult)
	if err := _LilypadController.contract.UnpackLog(event, "TimeoutSubmitResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
