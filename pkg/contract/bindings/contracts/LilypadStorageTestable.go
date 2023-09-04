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

// LilypadStorageTestableMetaData contains all meta data concerning the LilypadStorageTestable contract.
var LilypadStorageTestableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadStorageTestableABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadStorageTestableMetaData.ABI instead.
var LilypadStorageTestableABI = LilypadStorageTestableMetaData.ABI

// LilypadStorageTestable is an auto generated Go binding around an Ethereum contract.
type LilypadStorageTestable struct {
	LilypadStorageTestableCaller     // Read-only binding to the contract
	LilypadStorageTestableTransactor // Write-only binding to the contract
	LilypadStorageTestableFilterer   // Log filterer for contract events
}

// LilypadStorageTestableCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadStorageTestableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadStorageTestableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadStorageTestableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadStorageTestableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadStorageTestableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadStorageTestableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadStorageTestableSession struct {
	Contract     *LilypadStorageTestable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LilypadStorageTestableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadStorageTestableCallerSession struct {
	Contract *LilypadStorageTestableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// LilypadStorageTestableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadStorageTestableTransactorSession struct {
	Contract     *LilypadStorageTestableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// LilypadStorageTestableRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadStorageTestableRaw struct {
	Contract *LilypadStorageTestable // Generic contract binding to access the raw methods on
}

// LilypadStorageTestableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadStorageTestableCallerRaw struct {
	Contract *LilypadStorageTestableCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadStorageTestableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadStorageTestableTransactorRaw struct {
	Contract *LilypadStorageTestableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadStorageTestable creates a new instance of LilypadStorageTestable, bound to a specific deployed contract.
func NewLilypadStorageTestable(address common.Address, backend bind.ContractBackend) (*LilypadStorageTestable, error) {
	contract, err := bindLilypadStorageTestable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestable{LilypadStorageTestableCaller: LilypadStorageTestableCaller{contract: contract}, LilypadStorageTestableTransactor: LilypadStorageTestableTransactor{contract: contract}, LilypadStorageTestableFilterer: LilypadStorageTestableFilterer{contract: contract}}, nil
}

// NewLilypadStorageTestableCaller creates a new read-only instance of LilypadStorageTestable, bound to a specific deployed contract.
func NewLilypadStorageTestableCaller(address common.Address, caller bind.ContractCaller) (*LilypadStorageTestableCaller, error) {
	contract, err := bindLilypadStorageTestable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestableCaller{contract: contract}, nil
}

// NewLilypadStorageTestableTransactor creates a new write-only instance of LilypadStorageTestable, bound to a specific deployed contract.
func NewLilypadStorageTestableTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadStorageTestableTransactor, error) {
	contract, err := bindLilypadStorageTestable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestableTransactor{contract: contract}, nil
}

// NewLilypadStorageTestableFilterer creates a new log filterer instance of LilypadStorageTestable, bound to a specific deployed contract.
func NewLilypadStorageTestableFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadStorageTestableFilterer, error) {
	contract, err := bindLilypadStorageTestable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestableFilterer{contract: contract}, nil
}

// bindLilypadStorageTestable binds a generic wrapper to an already deployed contract.
func bindLilypadStorageTestable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadStorageTestableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadStorageTestable *LilypadStorageTestableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadStorageTestable.Contract.LilypadStorageTestableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadStorageTestable *LilypadStorageTestableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.LilypadStorageTestableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadStorageTestable *LilypadStorageTestableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.LilypadStorageTestableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadStorageTestable *LilypadStorageTestableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadStorageTestable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadStorageTestable *LilypadStorageTestableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadStorageTestable *LilypadStorageTestableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.contract.Transact(opts, method, params...)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetAgreement(opts *bind.CallOpts, dealId *big.Int) (SharedStructsAgreement, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getAgreement", dealId)

	if err != nil {
		return *new(SharedStructsAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsAgreement)).(*SharedStructsAgreement)

	return out0, err

}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _LilypadStorageTestable.Contract.GetAgreement(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _LilypadStorageTestable.Contract.GetAgreement(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetDeal(opts *bind.CallOpts, dealId *big.Int) (SharedStructsDeal, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _LilypadStorageTestable.Contract.GetDeal(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _LilypadStorageTestable.Contract.GetDeal(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetDealsForParty(opts *bind.CallOpts, party common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getDealsForParty", party)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _LilypadStorageTestable.Contract.GetDealsForParty(&_LilypadStorageTestable.CallOpts, party)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _LilypadStorageTestable.Contract.GetDealsForParty(&_LilypadStorageTestable.CallOpts, party)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetJobCost(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getJobCost", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorageTestable.Contract.GetJobCost(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorageTestable.Contract.GetJobCost(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetResult(opts *bind.CallOpts, dealId *big.Int) (SharedStructsResult, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getResult", dealId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _LilypadStorageTestable.Contract.GetResult(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _LilypadStorageTestable.Contract.GetResult(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetResultsCollateral(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getResultsCollateral", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorageTestable.Contract.GetResultsCollateral(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _LilypadStorageTestable.Contract.GetResultsCollateral(&_LilypadStorageTestable.CallOpts, dealId)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorageTestable *LilypadStorageTestableCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "getUser", userAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorageTestable *LilypadStorageTestableSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _LilypadStorageTestable.Contract.GetUser(&_LilypadStorageTestable.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _LilypadStorageTestable.Contract.GetUser(&_LilypadStorageTestable.CallOpts, userAddress)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_LilypadStorageTestable *LilypadStorageTestableCaller) HasDeal(opts *bind.CallOpts, dealId *big.Int) (bool, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "hasDeal", dealId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_LilypadStorageTestable *LilypadStorageTestableSession) HasDeal(dealId *big.Int) (bool, error) {
	return _LilypadStorageTestable.Contract.HasDeal(&_LilypadStorageTestable.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) HasDeal(dealId *big.Int) (bool, error) {
	return _LilypadStorageTestable.Contract.HasDeal(&_LilypadStorageTestable.CallOpts, dealId)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_LilypadStorageTestable *LilypadStorageTestableCaller) IsState(opts *bind.CallOpts, dealId *big.Int, state uint8) (bool, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "isState", dealId, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_LilypadStorageTestable *LilypadStorageTestableSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _LilypadStorageTestable.Contract.IsState(&_LilypadStorageTestable.CallOpts, dealId, state)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _LilypadStorageTestable.Contract.IsState(&_LilypadStorageTestable.CallOpts, dealId, state)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadStorageTestable *LilypadStorageTestableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadStorageTestable *LilypadStorageTestableSession) Owner() (common.Address, error) {
	return _LilypadStorageTestable.Contract.Owner(&_LilypadStorageTestable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) Owner() (common.Address, error) {
	return _LilypadStorageTestable.Contract.Owner(&_LilypadStorageTestable.CallOpts)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_LilypadStorageTestable *LilypadStorageTestableCaller) ShowUsersInList(opts *bind.CallOpts, serviceType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LilypadStorageTestable.contract.Call(opts, &out, "showUsersInList", serviceType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_LilypadStorageTestable *LilypadStorageTestableSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _LilypadStorageTestable.Contract.ShowUsersInList(&_LilypadStorageTestable.CallOpts, serviceType)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_LilypadStorageTestable *LilypadStorageTestableCallerSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _LilypadStorageTestable.Contract.ShowUsersInList(&_LilypadStorageTestable.CallOpts, serviceType)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AcceptResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AcceptResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AddResult(&_LilypadStorageTestable.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AddResult(&_LilypadStorageTestable.TransactOpts, dealId, resultsId, instructionCount)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AddUserToList(&_LilypadStorageTestable.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AddUserToList(&_LilypadStorageTestable.TransactOpts, serviceType)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AgreeJobCreator(&_LilypadStorageTestable.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AgreeJobCreator(&_LilypadStorageTestable.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AgreeResourceProvider(&_LilypadStorageTestable.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.AgreeResourceProvider(&_LilypadStorageTestable.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.CheckResult(&_LilypadStorageTestable.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.CheckResult(&_LilypadStorageTestable.TransactOpts, dealId, mediator)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.DisableChangeControllerAddress(&_LilypadStorageTestable.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.DisableChangeControllerAddress(&_LilypadStorageTestable.TransactOpts)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) EnsureDeal(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "ensureDeal", dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.EnsureDeal(&_LilypadStorageTestable.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.EnsureDeal(&_LilypadStorageTestable.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) Initialize() (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.Initialize(&_LilypadStorageTestable.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) Initialize() (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.Initialize(&_LilypadStorageTestable.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.MediationAcceptResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.MediationAcceptResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.MediationRejectResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.MediationRejectResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.RemoveUserFromList(&_LilypadStorageTestable.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.RemoveUserFromList(&_LilypadStorageTestable.TransactOpts, serviceType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.RenounceOwnership(&_LilypadStorageTestable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.RenounceOwnership(&_LilypadStorageTestable.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.SetControllerAddress(&_LilypadStorageTestable.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.SetControllerAddress(&_LilypadStorageTestable.TransactOpts, _controllerAddress)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TimeoutJudgeResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TimeoutJudgeResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TimeoutMediateResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TimeoutMediateResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TimeoutSubmitResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TimeoutSubmitResult(&_LilypadStorageTestable.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadStorageTestable *LilypadStorageTestableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TransferOwnership(&_LilypadStorageTestable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.TransferOwnership(&_LilypadStorageTestable.TransactOpts, newOwner)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorageTestable *LilypadStorageTestableTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.contract.Transact(opts, "updateUser", metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorageTestable *LilypadStorageTestableSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.UpdateUser(&_LilypadStorageTestable.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_LilypadStorageTestable *LilypadStorageTestableTransactorSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _LilypadStorageTestable.Contract.UpdateUser(&_LilypadStorageTestable.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// LilypadStorageTestableDealStateChangeIterator is returned from FilterDealStateChange and is used to iterate over the raw logs and unpacked data for DealStateChange events raised by the LilypadStorageTestable contract.
type LilypadStorageTestableDealStateChangeIterator struct {
	Event *LilypadStorageTestableDealStateChange // Event containing the contract specifics and raw log

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
func (it *LilypadStorageTestableDealStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageTestableDealStateChange)
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
		it.Event = new(LilypadStorageTestableDealStateChange)
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
func (it *LilypadStorageTestableDealStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageTestableDealStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageTestableDealStateChange represents a DealStateChange event raised by the LilypadStorageTestable contract.
type LilypadStorageTestableDealStateChange struct {
	DealId *big.Int
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) FilterDealStateChange(opts *bind.FilterOpts, dealId []*big.Int, state []uint8) (*LilypadStorageTestableDealStateChangeIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _LilypadStorageTestable.contract.FilterLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestableDealStateChangeIterator{contract: _LilypadStorageTestable.contract, event: "DealStateChange", logs: logs, sub: sub}, nil
}

// WatchDealStateChange is a free log subscription operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *LilypadStorageTestableDealStateChange, dealId []*big.Int, state []uint8) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _LilypadStorageTestable.contract.WatchLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageTestableDealStateChange)
				if err := _LilypadStorageTestable.contract.UnpackLog(event, "DealStateChange", log); err != nil {
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

// ParseDealStateChange is a log parse operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) ParseDealStateChange(log types.Log) (*LilypadStorageTestableDealStateChange, error) {
	event := new(LilypadStorageTestableDealStateChange)
	if err := _LilypadStorageTestable.contract.UnpackLog(event, "DealStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageTestableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadStorageTestable contract.
type LilypadStorageTestableInitializedIterator struct {
	Event *LilypadStorageTestableInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadStorageTestableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageTestableInitialized)
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
		it.Event = new(LilypadStorageTestableInitialized)
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
func (it *LilypadStorageTestableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageTestableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageTestableInitialized represents a Initialized event raised by the LilypadStorageTestable contract.
type LilypadStorageTestableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadStorageTestableInitializedIterator, error) {

	logs, sub, err := _LilypadStorageTestable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestableInitializedIterator{contract: _LilypadStorageTestable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadStorageTestableInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadStorageTestable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageTestableInitialized)
				if err := _LilypadStorageTestable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) ParseInitialized(log types.Log) (*LilypadStorageTestableInitialized, error) {
	event := new(LilypadStorageTestableInitialized)
	if err := _LilypadStorageTestable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadStorageTestableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LilypadStorageTestable contract.
type LilypadStorageTestableOwnershipTransferredIterator struct {
	Event *LilypadStorageTestableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LilypadStorageTestableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadStorageTestableOwnershipTransferred)
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
		it.Event = new(LilypadStorageTestableOwnershipTransferred)
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
func (it *LilypadStorageTestableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadStorageTestableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadStorageTestableOwnershipTransferred represents a OwnershipTransferred event raised by the LilypadStorageTestable contract.
type LilypadStorageTestableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LilypadStorageTestableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadStorageTestable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadStorageTestableOwnershipTransferredIterator{contract: _LilypadStorageTestable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LilypadStorageTestableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LilypadStorageTestable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadStorageTestableOwnershipTransferred)
				if err := _LilypadStorageTestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LilypadStorageTestable *LilypadStorageTestableFilterer) ParseOwnershipTransferred(log types.Log) (*LilypadStorageTestableOwnershipTransferred, error) {
	event := new(LilypadStorageTestableOwnershipTransferred)
	if err := _LilypadStorageTestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
