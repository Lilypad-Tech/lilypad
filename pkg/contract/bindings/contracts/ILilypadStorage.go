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

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId                    *big.Int
	ResourceProvider          common.Address
	JobCreator                common.Address
	InstructionPrice          *big.Int
	Timeout                   *big.Int
	TimeoutCollateral         *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	DealId           *big.Int
	ResultsId        *big.Int
	InstructionCount *big.Int
}

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress        common.Address
	MetadataCID        *big.Int
	Url                string
	Roles              []uint8
	TrustedMediators   []common.Address
	TrustedDirectories []common.Address
}

// ILilypadStorageMetaData contains all meta data concerning the ILilypadStorage contract.
var ILilypadStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILilypadStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ILilypadStorageMetaData.ABI instead.
var ILilypadStorageABI = ILilypadStorageMetaData.ABI

// ILilypadStorage is an auto generated Go binding around an Ethereum contract.
type ILilypadStorage struct {
	ILilypadStorageCaller     // Read-only binding to the contract
	ILilypadStorageTransactor // Write-only binding to the contract
	ILilypadStorageFilterer   // Log filterer for contract events
}

// ILilypadStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILilypadStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILilypadStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILilypadStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILilypadStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILilypadStorageSession struct {
	Contract     *ILilypadStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILilypadStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILilypadStorageCallerSession struct {
	Contract *ILilypadStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ILilypadStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILilypadStorageTransactorSession struct {
	Contract     *ILilypadStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ILilypadStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILilypadStorageRaw struct {
	Contract *ILilypadStorage // Generic contract binding to access the raw methods on
}

// ILilypadStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILilypadStorageCallerRaw struct {
	Contract *ILilypadStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ILilypadStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILilypadStorageTransactorRaw struct {
	Contract *ILilypadStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILilypadStorage creates a new instance of ILilypadStorage, bound to a specific deployed contract.
func NewILilypadStorage(address common.Address, backend bind.ContractBackend) (*ILilypadStorage, error) {
	contract, err := bindILilypadStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILilypadStorage{ILilypadStorageCaller: ILilypadStorageCaller{contract: contract}, ILilypadStorageTransactor: ILilypadStorageTransactor{contract: contract}, ILilypadStorageFilterer: ILilypadStorageFilterer{contract: contract}}, nil
}

// NewILilypadStorageCaller creates a new read-only instance of ILilypadStorage, bound to a specific deployed contract.
func NewILilypadStorageCaller(address common.Address, caller bind.ContractCaller) (*ILilypadStorageCaller, error) {
	contract, err := bindILilypadStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILilypadStorageCaller{contract: contract}, nil
}

// NewILilypadStorageTransactor creates a new write-only instance of ILilypadStorage, bound to a specific deployed contract.
func NewILilypadStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ILilypadStorageTransactor, error) {
	contract, err := bindILilypadStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILilypadStorageTransactor{contract: contract}, nil
}

// NewILilypadStorageFilterer creates a new log filterer instance of ILilypadStorage, bound to a specific deployed contract.
func NewILilypadStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ILilypadStorageFilterer, error) {
	contract, err := bindILilypadStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILilypadStorageFilterer{contract: contract}, nil
}

// bindILilypadStorage binds a generic wrapper to an already deployed contract.
func bindILilypadStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILilypadStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILilypadStorage *ILilypadStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILilypadStorage.Contract.ILilypadStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILilypadStorage *ILilypadStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.ILilypadStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILilypadStorage *ILilypadStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.ILilypadStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILilypadStorage *ILilypadStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILilypadStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILilypadStorage *ILilypadStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILilypadStorage *ILilypadStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.contract.Transact(opts, method, params...)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AcceptResult(&_ILilypadStorage.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AcceptResult(&_ILilypadStorage.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AddResult(&_ILilypadStorage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AddResult(&_ILilypadStorage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_ILilypadStorage *ILilypadStorageSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AddUserToList(&_ILilypadStorage.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AddUserToList(&_ILilypadStorage.TransactOpts, serviceType)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AgreeJobCreator(&_ILilypadStorage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AgreeJobCreator(&_ILilypadStorage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AgreeResourceProvider(&_ILilypadStorage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.AgreeResourceProvider(&_ILilypadStorage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_ILilypadStorage *ILilypadStorageSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.CheckResult(&_ILilypadStorage.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.CheckResult(&_ILilypadStorage.TransactOpts, dealId, mediator)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 jobCollateral, uint256 resultsCollateral, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) EnsureDeal(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, jobCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "ensureDeal", dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, jobCollateral, resultsCollateral, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 jobCollateral, uint256 resultsCollateral, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, jobCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.EnsureDeal(&_ILilypadStorage.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, jobCollateral, resultsCollateral, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 jobCollateral, uint256 resultsCollateral, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, jobCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.EnsureDeal(&_ILilypadStorage.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, jobCollateral, resultsCollateral, mediationFee)
}

// GetAgreement is a paid mutator transaction binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) GetAgreement(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getAgreement", dealId)
}

// GetAgreement is a paid mutator transaction binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) GetAgreement(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetAgreement(&_ILilypadStorage.TransactOpts, dealId)
}

// GetAgreement is a paid mutator transaction binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetAgreement(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetAgreement(&_ILilypadStorage.TransactOpts, dealId)
}

// GetDeal is a paid mutator transaction binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) GetDeal(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getDeal", dealId)
}

// GetDeal is a paid mutator transaction binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) GetDeal(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetDeal(&_ILilypadStorage.TransactOpts, dealId)
}

// GetDeal is a paid mutator transaction binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetDeal(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetDeal(&_ILilypadStorage.TransactOpts, dealId)
}

// GetDealsForParty is a paid mutator transaction binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) returns(uint256[])
func (_ILilypadStorage *ILilypadStorageTransactor) GetDealsForParty(opts *bind.TransactOpts, party common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getDealsForParty", party)
}

// GetDealsForParty is a paid mutator transaction binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) returns(uint256[])
func (_ILilypadStorage *ILilypadStorageSession) GetDealsForParty(party common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetDealsForParty(&_ILilypadStorage.TransactOpts, party)
}

// GetDealsForParty is a paid mutator transaction binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) returns(uint256[])
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetDealsForParty(party common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetDealsForParty(&_ILilypadStorage.TransactOpts, party)
}

// GetJobCost is a paid mutator transaction binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) returns(uint256)
func (_ILilypadStorage *ILilypadStorageTransactor) GetJobCost(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getJobCost", dealId)
}

// GetJobCost is a paid mutator transaction binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) returns(uint256)
func (_ILilypadStorage *ILilypadStorageSession) GetJobCost(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetJobCost(&_ILilypadStorage.TransactOpts, dealId)
}

// GetJobCost is a paid mutator transaction binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) returns(uint256)
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetJobCost(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetJobCost(&_ILilypadStorage.TransactOpts, dealId)
}

// GetResult is a paid mutator transaction binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) returns((uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactor) GetResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getResult", dealId)
}

// GetResult is a paid mutator transaction binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) returns((uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageSession) GetResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetResult(&_ILilypadStorage.TransactOpts, dealId)
}

// GetResult is a paid mutator transaction binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) returns((uint256,uint256,uint256))
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetResult(&_ILilypadStorage.TransactOpts, dealId)
}

// GetResultsCollateral is a paid mutator transaction binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) returns(uint256)
func (_ILilypadStorage *ILilypadStorageTransactor) GetResultsCollateral(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getResultsCollateral", dealId)
}

// GetResultsCollateral is a paid mutator transaction binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) returns(uint256)
func (_ILilypadStorage *ILilypadStorageSession) GetResultsCollateral(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetResultsCollateral(&_ILilypadStorage.TransactOpts, dealId)
}

// GetResultsCollateral is a paid mutator transaction binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) returns(uint256)
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetResultsCollateral(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetResultsCollateral(&_ILilypadStorage.TransactOpts, dealId)
}

// GetUser is a paid mutator transaction binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) returns((address,uint256,string,uint8[],address[],address[]))
func (_ILilypadStorage *ILilypadStorageTransactor) GetUser(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "getUser", userAddress)
}

// GetUser is a paid mutator transaction binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) returns((address,uint256,string,uint8[],address[],address[]))
func (_ILilypadStorage *ILilypadStorageSession) GetUser(userAddress common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetUser(&_ILilypadStorage.TransactOpts, userAddress)
}

// GetUser is a paid mutator transaction binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) returns((address,uint256,string,uint8[],address[],address[]))
func (_ILilypadStorage *ILilypadStorageTransactorSession) GetUser(userAddress common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.GetUser(&_ILilypadStorage.TransactOpts, userAddress)
}

// HasDeal is a paid mutator transaction binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) returns(bool)
func (_ILilypadStorage *ILilypadStorageTransactor) HasDeal(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "hasDeal", dealId)
}

// HasDeal is a paid mutator transaction binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) returns(bool)
func (_ILilypadStorage *ILilypadStorageSession) HasDeal(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.HasDeal(&_ILilypadStorage.TransactOpts, dealId)
}

// HasDeal is a paid mutator transaction binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) returns(bool)
func (_ILilypadStorage *ILilypadStorageTransactorSession) HasDeal(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.HasDeal(&_ILilypadStorage.TransactOpts, dealId)
}

// IsState is a paid mutator transaction binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) returns(bool)
func (_ILilypadStorage *ILilypadStorageTransactor) IsState(opts *bind.TransactOpts, dealId *big.Int, state uint8) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "isState", dealId, state)
}

// IsState is a paid mutator transaction binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) returns(bool)
func (_ILilypadStorage *ILilypadStorageSession) IsState(dealId *big.Int, state uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.IsState(&_ILilypadStorage.TransactOpts, dealId, state)
}

// IsState is a paid mutator transaction binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) returns(bool)
func (_ILilypadStorage *ILilypadStorageTransactorSession) IsState(dealId *big.Int, state uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.IsState(&_ILilypadStorage.TransactOpts, dealId, state)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.MediationAcceptResult(&_ILilypadStorage.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.MediationAcceptResult(&_ILilypadStorage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.MediationRejectResult(&_ILilypadStorage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.MediationRejectResult(&_ILilypadStorage.TransactOpts, dealId)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_ILilypadStorage *ILilypadStorageSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.RemoveUserFromList(&_ILilypadStorage.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.RemoveUserFromList(&_ILilypadStorage.TransactOpts, serviceType)
}

// ShowUsersInList is a paid mutator transaction binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) returns(address[])
func (_ILilypadStorage *ILilypadStorageTransactor) ShowUsersInList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "showUsersInList", serviceType)
}

// ShowUsersInList is a paid mutator transaction binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) returns(address[])
func (_ILilypadStorage *ILilypadStorageSession) ShowUsersInList(serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.ShowUsersInList(&_ILilypadStorage.TransactOpts, serviceType)
}

// ShowUsersInList is a paid mutator transaction binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) returns(address[])
func (_ILilypadStorage *ILilypadStorageTransactorSession) ShowUsersInList(serviceType uint8) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.ShowUsersInList(&_ILilypadStorage.TransactOpts, serviceType)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.TimeoutJudgeResult(&_ILilypadStorage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.TimeoutJudgeResult(&_ILilypadStorage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.TimeoutMediateResult(&_ILilypadStorage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.TimeoutMediateResult(&_ILilypadStorage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.TimeoutSubmitResult(&_ILilypadStorage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_ILilypadStorage *ILilypadStorageTransactorSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.TimeoutSubmitResult(&_ILilypadStorage.TransactOpts, dealId)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_ILilypadStorage *ILilypadStorageTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.contract.Transact(opts, "updateUser", metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_ILilypadStorage *ILilypadStorageSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.UpdateUser(&_ILilypadStorage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_ILilypadStorage *ILilypadStorageTransactorSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _ILilypadStorage.Contract.UpdateUser(&_ILilypadStorage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}
