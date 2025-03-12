// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypadproxy

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

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId           string
	JobCreator       common.Address
	ResourceProvider common.Address
	ModuleCreator    common.Address
	Solver           common.Address
	JobOfferCID      string
	ResourceOfferCID string
	Status           uint8
	Timestamp        *big.Int
	PaymentStructure SharedStructsDealPaymentStructure
}

// SharedStructsDealPaymentStructure is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPaymentStructure struct {
	JobCreatorSolverFee       *big.Int
	ResourceProviderSolverFee *big.Int
	NetworkCongestionFee      *big.Int
	ModuleCreatorFee          *big.Int
	PriceOfJobWithoutFees     *big.Int
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	ResultId  string
	DealId    string
	ResultCID string
	Status    uint8
	Timestamp *big.Int
}

// LilypadProxyMetaData contains all meta data concerning the LilypadProxy contract.
var LilypadProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptJobPayment\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptResourceProviderCollateral\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeal\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Deal\",\"components\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentStructure\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.DealPaymentStructure\",\"components\":[{\"name\":\"jobCreatorSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkCongestionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleCreatorFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceOfJobWithoutFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowBalance\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumResourceProviderCollateralAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentEngineAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResult\",\"inputs\":[{\"name\":\"_resultId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Result\",\"components\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getl2LilypadTokenAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_storageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_paymentEngineAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDeal\",\"inputs\":[{\"name\":\"deal\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Deal\",\"components\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentStructure\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.DealPaymentStructure\",\"components\":[{\"name\":\"jobCreatorSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkCongestionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleCreatorFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceOfJobWithoutFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL2LilypadTokenContract\",\"inputs\":[{\"name\":\"_l2LilypadTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentEngineContract\",\"inputs\":[{\"name\":\"_paymentEngineAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setResult\",\"inputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Result\",\"components\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStorageContract\",\"inputs\":[{\"name\":\"_storageAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserContract\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadProxy__JobCreatorEscrowPayment\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadProxy__JobCreatorInserted\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadProxy__ResourceProviderCollateralPayment\",\"inputs\":[{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadProxy__ResourceProviderInserted\",\"inputs\":[{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__DealFailedToLockup\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__DealFailedToSave\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__NotAuthorizedToGetResult\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__NotEnoughAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__ResultFailedToSave\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__ZeroAmountNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__acceptJobPayment__NotJobCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadProxy__acceptResourceProviderCollateral__NotResourceProvider\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// LilypadProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadProxyMetaData.ABI instead.
var LilypadProxyABI = LilypadProxyMetaData.ABI

// LilypadProxy is an auto generated Go binding around an Ethereum contract.
type LilypadProxy struct {
	LilypadProxyCaller     // Read-only binding to the contract
	LilypadProxyTransactor // Write-only binding to the contract
	LilypadProxyFilterer   // Log filterer for contract events
}

// LilypadProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadProxySession struct {
	Contract     *LilypadProxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LilypadProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadProxyCallerSession struct {
	Contract *LilypadProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LilypadProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadProxyTransactorSession struct {
	Contract     *LilypadProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LilypadProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadProxyRaw struct {
	Contract *LilypadProxy // Generic contract binding to access the raw methods on
}

// LilypadProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadProxyCallerRaw struct {
	Contract *LilypadProxyCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadProxyTransactorRaw struct {
	Contract *LilypadProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadProxy creates a new instance of LilypadProxy, bound to a specific deployed contract.
func NewLilypadProxy(address common.Address, backend bind.ContractBackend) (*LilypadProxy, error) {
	contract, err := bindLilypadProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadProxy{LilypadProxyCaller: LilypadProxyCaller{contract: contract}, LilypadProxyTransactor: LilypadProxyTransactor{contract: contract}, LilypadProxyFilterer: LilypadProxyFilterer{contract: contract}}, nil
}

// NewLilypadProxyCaller creates a new read-only instance of LilypadProxy, bound to a specific deployed contract.
func NewLilypadProxyCaller(address common.Address, caller bind.ContractCaller) (*LilypadProxyCaller, error) {
	contract, err := bindLilypadProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyCaller{contract: contract}, nil
}

// NewLilypadProxyTransactor creates a new write-only instance of LilypadProxy, bound to a specific deployed contract.
func NewLilypadProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadProxyTransactor, error) {
	contract, err := bindLilypadProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyTransactor{contract: contract}, nil
}

// NewLilypadProxyFilterer creates a new log filterer instance of LilypadProxy, bound to a specific deployed contract.
func NewLilypadProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadProxyFilterer, error) {
	contract, err := bindLilypadProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyFilterer{contract: contract}, nil
}

// bindLilypadProxy binds a generic wrapper to an already deployed contract.
func bindLilypadProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadProxy *LilypadProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadProxy.Contract.LilypadProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadProxy *LilypadProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadProxy.Contract.LilypadProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadProxy *LilypadProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadProxy.Contract.LilypadProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadProxy *LilypadProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadProxy *LilypadProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadProxy *LilypadProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadProxy.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadProxy *LilypadProxyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadProxy *LilypadProxySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadProxy.Contract.DEFAULTADMINROLE(&_LilypadProxy.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadProxy *LilypadProxyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadProxy.Contract.DEFAULTADMINROLE(&_LilypadProxy.CallOpts)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)))
func (_LilypadProxy *LilypadProxyCaller) GetDeal(opts *bind.CallOpts, dealId string) (SharedStructsDeal, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)))
func (_LilypadProxy *LilypadProxySession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _LilypadProxy.Contract.GetDeal(&_LilypadProxy.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)))
func (_LilypadProxy *LilypadProxyCallerSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _LilypadProxy.Contract.GetDeal(&_LilypadProxy.CallOpts, dealId)
}

// GetEscrowBalance is a free data retrieval call binding the contract method 0x6374b11b.
//
// Solidity: function getEscrowBalance(address _address) view returns(uint256)
func (_LilypadProxy *LilypadProxyCaller) GetEscrowBalance(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getEscrowBalance", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEscrowBalance is a free data retrieval call binding the contract method 0x6374b11b.
//
// Solidity: function getEscrowBalance(address _address) view returns(uint256)
func (_LilypadProxy *LilypadProxySession) GetEscrowBalance(_address common.Address) (*big.Int, error) {
	return _LilypadProxy.Contract.GetEscrowBalance(&_LilypadProxy.CallOpts, _address)
}

// GetEscrowBalance is a free data retrieval call binding the contract method 0x6374b11b.
//
// Solidity: function getEscrowBalance(address _address) view returns(uint256)
func (_LilypadProxy *LilypadProxyCallerSession) GetEscrowBalance(_address common.Address) (*big.Int, error) {
	return _LilypadProxy.Contract.GetEscrowBalance(&_LilypadProxy.CallOpts, _address)
}

// GetMinimumResourceProviderCollateralAmount is a free data retrieval call binding the contract method 0x457d7171.
//
// Solidity: function getMinimumResourceProviderCollateralAmount() view returns(uint256)
func (_LilypadProxy *LilypadProxyCaller) GetMinimumResourceProviderCollateralAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getMinimumResourceProviderCollateralAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumResourceProviderCollateralAmount is a free data retrieval call binding the contract method 0x457d7171.
//
// Solidity: function getMinimumResourceProviderCollateralAmount() view returns(uint256)
func (_LilypadProxy *LilypadProxySession) GetMinimumResourceProviderCollateralAmount() (*big.Int, error) {
	return _LilypadProxy.Contract.GetMinimumResourceProviderCollateralAmount(&_LilypadProxy.CallOpts)
}

// GetMinimumResourceProviderCollateralAmount is a free data retrieval call binding the contract method 0x457d7171.
//
// Solidity: function getMinimumResourceProviderCollateralAmount() view returns(uint256)
func (_LilypadProxy *LilypadProxyCallerSession) GetMinimumResourceProviderCollateralAmount() (*big.Int, error) {
	return _LilypadProxy.Contract.GetMinimumResourceProviderCollateralAmount(&_LilypadProxy.CallOpts)
}

// GetPaymentEngineAddress is a free data retrieval call binding the contract method 0xf932d450.
//
// Solidity: function getPaymentEngineAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCaller) GetPaymentEngineAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getPaymentEngineAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPaymentEngineAddress is a free data retrieval call binding the contract method 0xf932d450.
//
// Solidity: function getPaymentEngineAddress() view returns(address)
func (_LilypadProxy *LilypadProxySession) GetPaymentEngineAddress() (common.Address, error) {
	return _LilypadProxy.Contract.GetPaymentEngineAddress(&_LilypadProxy.CallOpts)
}

// GetPaymentEngineAddress is a free data retrieval call binding the contract method 0xf932d450.
//
// Solidity: function getPaymentEngineAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCallerSession) GetPaymentEngineAddress() (common.Address, error) {
	return _LilypadProxy.Contract.GetPaymentEngineAddress(&_LilypadProxy.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string _resultId) view returns((string,string,string,uint8,uint256))
func (_LilypadProxy *LilypadProxyCaller) GetResult(opts *bind.CallOpts, _resultId string) (SharedStructsResult, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getResult", _resultId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string _resultId) view returns((string,string,string,uint8,uint256))
func (_LilypadProxy *LilypadProxySession) GetResult(_resultId string) (SharedStructsResult, error) {
	return _LilypadProxy.Contract.GetResult(&_LilypadProxy.CallOpts, _resultId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string _resultId) view returns((string,string,string,uint8,uint256))
func (_LilypadProxy *LilypadProxyCallerSession) GetResult(_resultId string) (SharedStructsResult, error) {
	return _LilypadProxy.Contract.GetResult(&_LilypadProxy.CallOpts, _resultId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadProxy *LilypadProxyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadProxy *LilypadProxySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadProxy.Contract.GetRoleAdmin(&_LilypadProxy.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadProxy *LilypadProxyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadProxy.Contract.GetRoleAdmin(&_LilypadProxy.CallOpts, role)
}

// GetStorageAddress is a free data retrieval call binding the contract method 0x393a4d34.
//
// Solidity: function getStorageAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCaller) GetStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getStorageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStorageAddress is a free data retrieval call binding the contract method 0x393a4d34.
//
// Solidity: function getStorageAddress() view returns(address)
func (_LilypadProxy *LilypadProxySession) GetStorageAddress() (common.Address, error) {
	return _LilypadProxy.Contract.GetStorageAddress(&_LilypadProxy.CallOpts)
}

// GetStorageAddress is a free data retrieval call binding the contract method 0x393a4d34.
//
// Solidity: function getStorageAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCallerSession) GetStorageAddress() (common.Address, error) {
	return _LilypadProxy.Contract.GetStorageAddress(&_LilypadProxy.CallOpts)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xebc3eee8.
//
// Solidity: function getUserAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCaller) GetUserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getUserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserAddress is a free data retrieval call binding the contract method 0xebc3eee8.
//
// Solidity: function getUserAddress() view returns(address)
func (_LilypadProxy *LilypadProxySession) GetUserAddress() (common.Address, error) {
	return _LilypadProxy.Contract.GetUserAddress(&_LilypadProxy.CallOpts)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xebc3eee8.
//
// Solidity: function getUserAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCallerSession) GetUserAddress() (common.Address, error) {
	return _LilypadProxy.Contract.GetUserAddress(&_LilypadProxy.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadProxy *LilypadProxyCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadProxy *LilypadProxySession) GetVersion() (string, error) {
	return _LilypadProxy.Contract.GetVersion(&_LilypadProxy.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_LilypadProxy *LilypadProxyCallerSession) GetVersion() (string, error) {
	return _LilypadProxy.Contract.GetVersion(&_LilypadProxy.CallOpts)
}

// Getl2LilypadTokenAddress is a free data retrieval call binding the contract method 0x336b4fb4.
//
// Solidity: function getl2LilypadTokenAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCaller) Getl2LilypadTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "getl2LilypadTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Getl2LilypadTokenAddress is a free data retrieval call binding the contract method 0x336b4fb4.
//
// Solidity: function getl2LilypadTokenAddress() view returns(address)
func (_LilypadProxy *LilypadProxySession) Getl2LilypadTokenAddress() (common.Address, error) {
	return _LilypadProxy.Contract.Getl2LilypadTokenAddress(&_LilypadProxy.CallOpts)
}

// Getl2LilypadTokenAddress is a free data retrieval call binding the contract method 0x336b4fb4.
//
// Solidity: function getl2LilypadTokenAddress() view returns(address)
func (_LilypadProxy *LilypadProxyCallerSession) Getl2LilypadTokenAddress() (common.Address, error) {
	return _LilypadProxy.Contract.Getl2LilypadTokenAddress(&_LilypadProxy.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadProxy *LilypadProxyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadProxy *LilypadProxySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadProxy.Contract.HasRole(&_LilypadProxy.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadProxy *LilypadProxyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadProxy.Contract.HasRole(&_LilypadProxy.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadProxy *LilypadProxyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadProxy *LilypadProxySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadProxy.Contract.SupportsInterface(&_LilypadProxy.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadProxy *LilypadProxyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadProxy.Contract.SupportsInterface(&_LilypadProxy.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadProxy *LilypadProxyCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadProxy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadProxy *LilypadProxySession) Version() (string, error) {
	return _LilypadProxy.Contract.Version(&_LilypadProxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadProxy *LilypadProxyCallerSession) Version() (string, error) {
	return _LilypadProxy.Contract.Version(&_LilypadProxy.CallOpts)
}

// AcceptJobPayment is a paid mutator transaction binding the contract method 0x00e8d4d0.
//
// Solidity: function acceptJobPayment(uint256 _amount) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) AcceptJobPayment(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "acceptJobPayment", _amount)
}

// AcceptJobPayment is a paid mutator transaction binding the contract method 0x00e8d4d0.
//
// Solidity: function acceptJobPayment(uint256 _amount) returns(bool)
func (_LilypadProxy *LilypadProxySession) AcceptJobPayment(_amount *big.Int) (*types.Transaction, error) {
	return _LilypadProxy.Contract.AcceptJobPayment(&_LilypadProxy.TransactOpts, _amount)
}

// AcceptJobPayment is a paid mutator transaction binding the contract method 0x00e8d4d0.
//
// Solidity: function acceptJobPayment(uint256 _amount) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) AcceptJobPayment(_amount *big.Int) (*types.Transaction, error) {
	return _LilypadProxy.Contract.AcceptJobPayment(&_LilypadProxy.TransactOpts, _amount)
}

// AcceptResourceProviderCollateral is a paid mutator transaction binding the contract method 0xa983cb75.
//
// Solidity: function acceptResourceProviderCollateral(uint256 _amount) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) AcceptResourceProviderCollateral(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "acceptResourceProviderCollateral", _amount)
}

// AcceptResourceProviderCollateral is a paid mutator transaction binding the contract method 0xa983cb75.
//
// Solidity: function acceptResourceProviderCollateral(uint256 _amount) returns(bool)
func (_LilypadProxy *LilypadProxySession) AcceptResourceProviderCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _LilypadProxy.Contract.AcceptResourceProviderCollateral(&_LilypadProxy.TransactOpts, _amount)
}

// AcceptResourceProviderCollateral is a paid mutator transaction binding the contract method 0xa983cb75.
//
// Solidity: function acceptResourceProviderCollateral(uint256 _amount) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) AcceptResourceProviderCollateral(_amount *big.Int) (*types.Transaction, error) {
	return _LilypadProxy.Contract.AcceptResourceProviderCollateral(&_LilypadProxy.TransactOpts, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadProxy *LilypadProxyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadProxy *LilypadProxySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.GrantRole(&_LilypadProxy.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadProxy *LilypadProxyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.GrantRole(&_LilypadProxy.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _storageAddress, address _paymentEngineAddress, address _userAddress, address _tokenAddress) returns()
func (_LilypadProxy *LilypadProxyTransactor) Initialize(opts *bind.TransactOpts, _storageAddress common.Address, _paymentEngineAddress common.Address, _userAddress common.Address, _tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "initialize", _storageAddress, _paymentEngineAddress, _userAddress, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _storageAddress, address _paymentEngineAddress, address _userAddress, address _tokenAddress) returns()
func (_LilypadProxy *LilypadProxySession) Initialize(_storageAddress common.Address, _paymentEngineAddress common.Address, _userAddress common.Address, _tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.Initialize(&_LilypadProxy.TransactOpts, _storageAddress, _paymentEngineAddress, _userAddress, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _storageAddress, address _paymentEngineAddress, address _userAddress, address _tokenAddress) returns()
func (_LilypadProxy *LilypadProxyTransactorSession) Initialize(_storageAddress common.Address, _paymentEngineAddress common.Address, _userAddress common.Address, _tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.Initialize(&_LilypadProxy.TransactOpts, _storageAddress, _paymentEngineAddress, _userAddress, _tokenAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadProxy *LilypadProxyTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadProxy *LilypadProxySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.RenounceRole(&_LilypadProxy.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadProxy *LilypadProxyTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.RenounceRole(&_LilypadProxy.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadProxy *LilypadProxyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadProxy *LilypadProxySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.RevokeRole(&_LilypadProxy.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadProxy *LilypadProxyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.RevokeRole(&_LilypadProxy.TransactOpts, role, account)
}

// SetDeal is a paid mutator transaction binding the contract method 0xd6fae778.
//
// Solidity: function setDeal((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) deal) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) SetDeal(opts *bind.TransactOpts, deal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "setDeal", deal)
}

// SetDeal is a paid mutator transaction binding the contract method 0xd6fae778.
//
// Solidity: function setDeal((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) deal) returns(bool)
func (_LilypadProxy *LilypadProxySession) SetDeal(deal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetDeal(&_LilypadProxy.TransactOpts, deal)
}

// SetDeal is a paid mutator transaction binding the contract method 0xd6fae778.
//
// Solidity: function setDeal((string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) deal) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) SetDeal(deal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetDeal(&_LilypadProxy.TransactOpts, deal)
}

// SetL2LilypadTokenContract is a paid mutator transaction binding the contract method 0x8f39c9b5.
//
// Solidity: function setL2LilypadTokenContract(address _l2LilypadTokenAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) SetL2LilypadTokenContract(opts *bind.TransactOpts, _l2LilypadTokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "setL2LilypadTokenContract", _l2LilypadTokenAddress)
}

// SetL2LilypadTokenContract is a paid mutator transaction binding the contract method 0x8f39c9b5.
//
// Solidity: function setL2LilypadTokenContract(address _l2LilypadTokenAddress) returns(bool)
func (_LilypadProxy *LilypadProxySession) SetL2LilypadTokenContract(_l2LilypadTokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetL2LilypadTokenContract(&_LilypadProxy.TransactOpts, _l2LilypadTokenAddress)
}

// SetL2LilypadTokenContract is a paid mutator transaction binding the contract method 0x8f39c9b5.
//
// Solidity: function setL2LilypadTokenContract(address _l2LilypadTokenAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) SetL2LilypadTokenContract(_l2LilypadTokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetL2LilypadTokenContract(&_LilypadProxy.TransactOpts, _l2LilypadTokenAddress)
}

// SetPaymentEngineContract is a paid mutator transaction binding the contract method 0x099639c2.
//
// Solidity: function setPaymentEngineContract(address _paymentEngineAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) SetPaymentEngineContract(opts *bind.TransactOpts, _paymentEngineAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "setPaymentEngineContract", _paymentEngineAddress)
}

// SetPaymentEngineContract is a paid mutator transaction binding the contract method 0x099639c2.
//
// Solidity: function setPaymentEngineContract(address _paymentEngineAddress) returns(bool)
func (_LilypadProxy *LilypadProxySession) SetPaymentEngineContract(_paymentEngineAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetPaymentEngineContract(&_LilypadProxy.TransactOpts, _paymentEngineAddress)
}

// SetPaymentEngineContract is a paid mutator transaction binding the contract method 0x099639c2.
//
// Solidity: function setPaymentEngineContract(address _paymentEngineAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) SetPaymentEngineContract(_paymentEngineAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetPaymentEngineContract(&_LilypadProxy.TransactOpts, _paymentEngineAddress)
}

// SetResult is a paid mutator transaction binding the contract method 0x9d3d78d0.
//
// Solidity: function setResult((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) SetResult(opts *bind.TransactOpts, result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "setResult", result)
}

// SetResult is a paid mutator transaction binding the contract method 0x9d3d78d0.
//
// Solidity: function setResult((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadProxy *LilypadProxySession) SetResult(result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetResult(&_LilypadProxy.TransactOpts, result)
}

// SetResult is a paid mutator transaction binding the contract method 0x9d3d78d0.
//
// Solidity: function setResult((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) SetResult(result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetResult(&_LilypadProxy.TransactOpts, result)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address _storageAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) SetStorageContract(opts *bind.TransactOpts, _storageAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "setStorageContract", _storageAddress)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address _storageAddress) returns(bool)
func (_LilypadProxy *LilypadProxySession) SetStorageContract(_storageAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetStorageContract(&_LilypadProxy.TransactOpts, _storageAddress)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address _storageAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) SetStorageContract(_storageAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetStorageContract(&_LilypadProxy.TransactOpts, _storageAddress)
}

// SetUserContract is a paid mutator transaction binding the contract method 0xa1206d4e.
//
// Solidity: function setUserContract(address _userAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactor) SetUserContract(opts *bind.TransactOpts, _userAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.contract.Transact(opts, "setUserContract", _userAddress)
}

// SetUserContract is a paid mutator transaction binding the contract method 0xa1206d4e.
//
// Solidity: function setUserContract(address _userAddress) returns(bool)
func (_LilypadProxy *LilypadProxySession) SetUserContract(_userAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetUserContract(&_LilypadProxy.TransactOpts, _userAddress)
}

// SetUserContract is a paid mutator transaction binding the contract method 0xa1206d4e.
//
// Solidity: function setUserContract(address _userAddress) returns(bool)
func (_LilypadProxy *LilypadProxyTransactorSession) SetUserContract(_userAddress common.Address) (*types.Transaction, error) {
	return _LilypadProxy.Contract.SetUserContract(&_LilypadProxy.TransactOpts, _userAddress)
}

// LilypadProxyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadProxy contract.
type LilypadProxyInitializedIterator struct {
	Event *LilypadProxyInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadProxyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyInitialized)
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
		it.Event = new(LilypadProxyInitialized)
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
func (it *LilypadProxyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyInitialized represents a Initialized event raised by the LilypadProxy contract.
type LilypadProxyInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadProxy *LilypadProxyFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadProxyInitializedIterator, error) {

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadProxyInitializedIterator{contract: _LilypadProxy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadProxy *LilypadProxyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadProxyInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyInitialized)
				if err := _LilypadProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LilypadProxy *LilypadProxyFilterer) ParseInitialized(log types.Log) (*LilypadProxyInitialized, error) {
	event := new(LilypadProxyInitialized)
	if err := _LilypadProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator is returned from FilterLilypadProxyJobCreatorEscrowPayment and is used to iterate over the raw logs and unpacked data for LilypadProxyJobCreatorEscrowPayment events raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator struct {
	Event *LilypadProxyLilypadProxyJobCreatorEscrowPayment // Event containing the contract specifics and raw log

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
func (it *LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyLilypadProxyJobCreatorEscrowPayment)
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
		it.Event = new(LilypadProxyLilypadProxyJobCreatorEscrowPayment)
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
func (it *LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyLilypadProxyJobCreatorEscrowPayment represents a LilypadProxyJobCreatorEscrowPayment event raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyJobCreatorEscrowPayment struct {
	JobCreator common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLilypadProxyJobCreatorEscrowPayment is a free log retrieval operation binding the contract event 0x7cc822bca116608144e9e02e9421260014477ffacfba53e62226c835aea12e38.
//
// Solidity: event LilypadProxy__JobCreatorEscrowPayment(address indexed jobCreator, uint256 amount)
func (_LilypadProxy *LilypadProxyFilterer) FilterLilypadProxyJobCreatorEscrowPayment(opts *bind.FilterOpts, jobCreator []common.Address) (*LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "LilypadProxy__JobCreatorEscrowPayment", jobCreatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyLilypadProxyJobCreatorEscrowPaymentIterator{contract: _LilypadProxy.contract, event: "LilypadProxy__JobCreatorEscrowPayment", logs: logs, sub: sub}, nil
}

// WatchLilypadProxyJobCreatorEscrowPayment is a free log subscription operation binding the contract event 0x7cc822bca116608144e9e02e9421260014477ffacfba53e62226c835aea12e38.
//
// Solidity: event LilypadProxy__JobCreatorEscrowPayment(address indexed jobCreator, uint256 amount)
func (_LilypadProxy *LilypadProxyFilterer) WatchLilypadProxyJobCreatorEscrowPayment(opts *bind.WatchOpts, sink chan<- *LilypadProxyLilypadProxyJobCreatorEscrowPayment, jobCreator []common.Address) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "LilypadProxy__JobCreatorEscrowPayment", jobCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyLilypadProxyJobCreatorEscrowPayment)
				if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__JobCreatorEscrowPayment", log); err != nil {
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

// ParseLilypadProxyJobCreatorEscrowPayment is a log parse operation binding the contract event 0x7cc822bca116608144e9e02e9421260014477ffacfba53e62226c835aea12e38.
//
// Solidity: event LilypadProxy__JobCreatorEscrowPayment(address indexed jobCreator, uint256 amount)
func (_LilypadProxy *LilypadProxyFilterer) ParseLilypadProxyJobCreatorEscrowPayment(log types.Log) (*LilypadProxyLilypadProxyJobCreatorEscrowPayment, error) {
	event := new(LilypadProxyLilypadProxyJobCreatorEscrowPayment)
	if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__JobCreatorEscrowPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyLilypadProxyJobCreatorInsertedIterator is returned from FilterLilypadProxyJobCreatorInserted and is used to iterate over the raw logs and unpacked data for LilypadProxyJobCreatorInserted events raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyJobCreatorInsertedIterator struct {
	Event *LilypadProxyLilypadProxyJobCreatorInserted // Event containing the contract specifics and raw log

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
func (it *LilypadProxyLilypadProxyJobCreatorInsertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyLilypadProxyJobCreatorInserted)
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
		it.Event = new(LilypadProxyLilypadProxyJobCreatorInserted)
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
func (it *LilypadProxyLilypadProxyJobCreatorInsertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyLilypadProxyJobCreatorInsertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyLilypadProxyJobCreatorInserted represents a LilypadProxyJobCreatorInserted event raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyJobCreatorInserted struct {
	JobCreator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLilypadProxyJobCreatorInserted is a free log retrieval operation binding the contract event 0xa03c73bf251171248c56d7e56efb4dff8caabd2cc76aa6362911fd344d2cef45.
//
// Solidity: event LilypadProxy__JobCreatorInserted(address indexed jobCreator)
func (_LilypadProxy *LilypadProxyFilterer) FilterLilypadProxyJobCreatorInserted(opts *bind.FilterOpts, jobCreator []common.Address) (*LilypadProxyLilypadProxyJobCreatorInsertedIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "LilypadProxy__JobCreatorInserted", jobCreatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyLilypadProxyJobCreatorInsertedIterator{contract: _LilypadProxy.contract, event: "LilypadProxy__JobCreatorInserted", logs: logs, sub: sub}, nil
}

// WatchLilypadProxyJobCreatorInserted is a free log subscription operation binding the contract event 0xa03c73bf251171248c56d7e56efb4dff8caabd2cc76aa6362911fd344d2cef45.
//
// Solidity: event LilypadProxy__JobCreatorInserted(address indexed jobCreator)
func (_LilypadProxy *LilypadProxyFilterer) WatchLilypadProxyJobCreatorInserted(opts *bind.WatchOpts, sink chan<- *LilypadProxyLilypadProxyJobCreatorInserted, jobCreator []common.Address) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "LilypadProxy__JobCreatorInserted", jobCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyLilypadProxyJobCreatorInserted)
				if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__JobCreatorInserted", log); err != nil {
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

// ParseLilypadProxyJobCreatorInserted is a log parse operation binding the contract event 0xa03c73bf251171248c56d7e56efb4dff8caabd2cc76aa6362911fd344d2cef45.
//
// Solidity: event LilypadProxy__JobCreatorInserted(address indexed jobCreator)
func (_LilypadProxy *LilypadProxyFilterer) ParseLilypadProxyJobCreatorInserted(log types.Log) (*LilypadProxyLilypadProxyJobCreatorInserted, error) {
	event := new(LilypadProxyLilypadProxyJobCreatorInserted)
	if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__JobCreatorInserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator is returned from FilterLilypadProxyResourceProviderCollateralPayment and is used to iterate over the raw logs and unpacked data for LilypadProxyResourceProviderCollateralPayment events raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator struct {
	Event *LilypadProxyLilypadProxyResourceProviderCollateralPayment // Event containing the contract specifics and raw log

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
func (it *LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyLilypadProxyResourceProviderCollateralPayment)
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
		it.Event = new(LilypadProxyLilypadProxyResourceProviderCollateralPayment)
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
func (it *LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyLilypadProxyResourceProviderCollateralPayment represents a LilypadProxyResourceProviderCollateralPayment event raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyResourceProviderCollateralPayment struct {
	ResourceProvider common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadProxyResourceProviderCollateralPayment is a free log retrieval operation binding the contract event 0x9a73f16cd498fba2fb90f676ce8eed075770707ee29bac8e6dec321cde49b583.
//
// Solidity: event LilypadProxy__ResourceProviderCollateralPayment(address indexed resourceProvider, uint256 amount)
func (_LilypadProxy *LilypadProxyFilterer) FilterLilypadProxyResourceProviderCollateralPayment(opts *bind.FilterOpts, resourceProvider []common.Address) (*LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator, error) {

	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "LilypadProxy__ResourceProviderCollateralPayment", resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyLilypadProxyResourceProviderCollateralPaymentIterator{contract: _LilypadProxy.contract, event: "LilypadProxy__ResourceProviderCollateralPayment", logs: logs, sub: sub}, nil
}

// WatchLilypadProxyResourceProviderCollateralPayment is a free log subscription operation binding the contract event 0x9a73f16cd498fba2fb90f676ce8eed075770707ee29bac8e6dec321cde49b583.
//
// Solidity: event LilypadProxy__ResourceProviderCollateralPayment(address indexed resourceProvider, uint256 amount)
func (_LilypadProxy *LilypadProxyFilterer) WatchLilypadProxyResourceProviderCollateralPayment(opts *bind.WatchOpts, sink chan<- *LilypadProxyLilypadProxyResourceProviderCollateralPayment, resourceProvider []common.Address) (event.Subscription, error) {

	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "LilypadProxy__ResourceProviderCollateralPayment", resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyLilypadProxyResourceProviderCollateralPayment)
				if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__ResourceProviderCollateralPayment", log); err != nil {
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

// ParseLilypadProxyResourceProviderCollateralPayment is a log parse operation binding the contract event 0x9a73f16cd498fba2fb90f676ce8eed075770707ee29bac8e6dec321cde49b583.
//
// Solidity: event LilypadProxy__ResourceProviderCollateralPayment(address indexed resourceProvider, uint256 amount)
func (_LilypadProxy *LilypadProxyFilterer) ParseLilypadProxyResourceProviderCollateralPayment(log types.Log) (*LilypadProxyLilypadProxyResourceProviderCollateralPayment, error) {
	event := new(LilypadProxyLilypadProxyResourceProviderCollateralPayment)
	if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__ResourceProviderCollateralPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyLilypadProxyResourceProviderInsertedIterator is returned from FilterLilypadProxyResourceProviderInserted and is used to iterate over the raw logs and unpacked data for LilypadProxyResourceProviderInserted events raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyResourceProviderInsertedIterator struct {
	Event *LilypadProxyLilypadProxyResourceProviderInserted // Event containing the contract specifics and raw log

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
func (it *LilypadProxyLilypadProxyResourceProviderInsertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyLilypadProxyResourceProviderInserted)
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
		it.Event = new(LilypadProxyLilypadProxyResourceProviderInserted)
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
func (it *LilypadProxyLilypadProxyResourceProviderInsertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyLilypadProxyResourceProviderInsertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyLilypadProxyResourceProviderInserted represents a LilypadProxyResourceProviderInserted event raised by the LilypadProxy contract.
type LilypadProxyLilypadProxyResourceProviderInserted struct {
	ResourceProvider common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadProxyResourceProviderInserted is a free log retrieval operation binding the contract event 0x8364be879e1ecb73f6498bb5ec867cfc192860269a248a9a90aa82ae1576052e.
//
// Solidity: event LilypadProxy__ResourceProviderInserted(address indexed resourceProvider)
func (_LilypadProxy *LilypadProxyFilterer) FilterLilypadProxyResourceProviderInserted(opts *bind.FilterOpts, resourceProvider []common.Address) (*LilypadProxyLilypadProxyResourceProviderInsertedIterator, error) {

	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "LilypadProxy__ResourceProviderInserted", resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyLilypadProxyResourceProviderInsertedIterator{contract: _LilypadProxy.contract, event: "LilypadProxy__ResourceProviderInserted", logs: logs, sub: sub}, nil
}

// WatchLilypadProxyResourceProviderInserted is a free log subscription operation binding the contract event 0x8364be879e1ecb73f6498bb5ec867cfc192860269a248a9a90aa82ae1576052e.
//
// Solidity: event LilypadProxy__ResourceProviderInserted(address indexed resourceProvider)
func (_LilypadProxy *LilypadProxyFilterer) WatchLilypadProxyResourceProviderInserted(opts *bind.WatchOpts, sink chan<- *LilypadProxyLilypadProxyResourceProviderInserted, resourceProvider []common.Address) (event.Subscription, error) {

	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "LilypadProxy__ResourceProviderInserted", resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyLilypadProxyResourceProviderInserted)
				if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__ResourceProviderInserted", log); err != nil {
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

// ParseLilypadProxyResourceProviderInserted is a log parse operation binding the contract event 0x8364be879e1ecb73f6498bb5ec867cfc192860269a248a9a90aa82ae1576052e.
//
// Solidity: event LilypadProxy__ResourceProviderInserted(address indexed resourceProvider)
func (_LilypadProxy *LilypadProxyFilterer) ParseLilypadProxyResourceProviderInserted(log types.Log) (*LilypadProxyLilypadProxyResourceProviderInserted, error) {
	event := new(LilypadProxyLilypadProxyResourceProviderInserted)
	if err := _LilypadProxy.contract.UnpackLog(event, "LilypadProxy__ResourceProviderInserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadProxy contract.
type LilypadProxyRoleAdminChangedIterator struct {
	Event *LilypadProxyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadProxyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyRoleAdminChanged)
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
		it.Event = new(LilypadProxyRoleAdminChanged)
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
func (it *LilypadProxyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadProxy contract.
type LilypadProxyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadProxy *LilypadProxyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadProxyRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyRoleAdminChangedIterator{contract: _LilypadProxy.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadProxy *LilypadProxyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadProxyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyRoleAdminChanged)
				if err := _LilypadProxy.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LilypadProxy *LilypadProxyFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadProxyRoleAdminChanged, error) {
	event := new(LilypadProxyRoleAdminChanged)
	if err := _LilypadProxy.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadProxy contract.
type LilypadProxyRoleGrantedIterator struct {
	Event *LilypadProxyRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadProxyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyRoleGranted)
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
		it.Event = new(LilypadProxyRoleGranted)
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
func (it *LilypadProxyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyRoleGranted represents a RoleGranted event raised by the LilypadProxy contract.
type LilypadProxyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadProxy *LilypadProxyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadProxyRoleGrantedIterator, error) {

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

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyRoleGrantedIterator{contract: _LilypadProxy.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadProxy *LilypadProxyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadProxyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyRoleGranted)
				if err := _LilypadProxy.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LilypadProxy *LilypadProxyFilterer) ParseRoleGranted(log types.Log) (*LilypadProxyRoleGranted, error) {
	event := new(LilypadProxyRoleGranted)
	if err := _LilypadProxy.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadProxyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadProxy contract.
type LilypadProxyRoleRevokedIterator struct {
	Event *LilypadProxyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadProxyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadProxyRoleRevoked)
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
		it.Event = new(LilypadProxyRoleRevoked)
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
func (it *LilypadProxyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadProxyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadProxyRoleRevoked represents a RoleRevoked event raised by the LilypadProxy contract.
type LilypadProxyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadProxy *LilypadProxyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadProxyRoleRevokedIterator, error) {

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

	logs, sub, err := _LilypadProxy.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadProxyRoleRevokedIterator{contract: _LilypadProxy.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadProxy *LilypadProxyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadProxyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LilypadProxy.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadProxyRoleRevoked)
				if err := _LilypadProxy.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LilypadProxy *LilypadProxyFilterer) ParseRoleRevoked(log types.Log) (*LilypadProxyRoleRevoked, error) {
	event := new(LilypadProxyRoleRevoked)
	if err := _LilypadProxy.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
