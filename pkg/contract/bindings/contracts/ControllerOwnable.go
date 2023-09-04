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

// ControllerOwnableMetaData contains all meta data concerning the ControllerOwnable contract.
var ControllerOwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ControllerOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerOwnableMetaData.ABI instead.
var ControllerOwnableABI = ControllerOwnableMetaData.ABI

// ControllerOwnable is an auto generated Go binding around an Ethereum contract.
type ControllerOwnable struct {
	ControllerOwnableCaller     // Read-only binding to the contract
	ControllerOwnableTransactor // Write-only binding to the contract
	ControllerOwnableFilterer   // Log filterer for contract events
}

// ControllerOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerOwnableSession struct {
	Contract     *ControllerOwnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ControllerOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerOwnableCallerSession struct {
	Contract *ControllerOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ControllerOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerOwnableTransactorSession struct {
	Contract     *ControllerOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ControllerOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerOwnableRaw struct {
	Contract *ControllerOwnable // Generic contract binding to access the raw methods on
}

// ControllerOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerOwnableCallerRaw struct {
	Contract *ControllerOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerOwnableTransactorRaw struct {
	Contract *ControllerOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControllerOwnable creates a new instance of ControllerOwnable, bound to a specific deployed contract.
func NewControllerOwnable(address common.Address, backend bind.ContractBackend) (*ControllerOwnable, error) {
	contract, err := bindControllerOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnable{ControllerOwnableCaller: ControllerOwnableCaller{contract: contract}, ControllerOwnableTransactor: ControllerOwnableTransactor{contract: contract}, ControllerOwnableFilterer: ControllerOwnableFilterer{contract: contract}}, nil
}

// NewControllerOwnableCaller creates a new read-only instance of ControllerOwnable, bound to a specific deployed contract.
func NewControllerOwnableCaller(address common.Address, caller bind.ContractCaller) (*ControllerOwnableCaller, error) {
	contract, err := bindControllerOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnableCaller{contract: contract}, nil
}

// NewControllerOwnableTransactor creates a new write-only instance of ControllerOwnable, bound to a specific deployed contract.
func NewControllerOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerOwnableTransactor, error) {
	contract, err := bindControllerOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnableTransactor{contract: contract}, nil
}

// NewControllerOwnableFilterer creates a new log filterer instance of ControllerOwnable, bound to a specific deployed contract.
func NewControllerOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerOwnableFilterer, error) {
	contract, err := bindControllerOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnableFilterer{contract: contract}, nil
}

// bindControllerOwnable binds a generic wrapper to an already deployed contract.
func bindControllerOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerOwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControllerOwnable *ControllerOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ControllerOwnable.Contract.ControllerOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControllerOwnable *ControllerOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.ControllerOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControllerOwnable *ControllerOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.ControllerOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ControllerOwnable *ControllerOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ControllerOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ControllerOwnable *ControllerOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ControllerOwnable *ControllerOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ControllerOwnable *ControllerOwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ControllerOwnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ControllerOwnable *ControllerOwnableSession) Owner() (common.Address, error) {
	return _ControllerOwnable.Contract.Owner(&_ControllerOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ControllerOwnable *ControllerOwnableCallerSession) Owner() (common.Address, error) {
	return _ControllerOwnable.Contract.Owner(&_ControllerOwnable.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_ControllerOwnable *ControllerOwnableTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ControllerOwnable.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_ControllerOwnable *ControllerOwnableSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _ControllerOwnable.Contract.DisableChangeControllerAddress(&_ControllerOwnable.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_ControllerOwnable *ControllerOwnableTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _ControllerOwnable.Contract.DisableChangeControllerAddress(&_ControllerOwnable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ControllerOwnable *ControllerOwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ControllerOwnable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ControllerOwnable *ControllerOwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _ControllerOwnable.Contract.RenounceOwnership(&_ControllerOwnable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ControllerOwnable *ControllerOwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ControllerOwnable.Contract.RenounceOwnership(&_ControllerOwnable.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_ControllerOwnable *ControllerOwnableTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _ControllerOwnable.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_ControllerOwnable *ControllerOwnableSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.SetControllerAddress(&_ControllerOwnable.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_ControllerOwnable *ControllerOwnableTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.SetControllerAddress(&_ControllerOwnable.TransactOpts, _controllerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ControllerOwnable *ControllerOwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ControllerOwnable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ControllerOwnable *ControllerOwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.TransferOwnership(&_ControllerOwnable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ControllerOwnable *ControllerOwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ControllerOwnable.Contract.TransferOwnership(&_ControllerOwnable.TransactOpts, newOwner)
}

// ControllerOwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ControllerOwnable contract.
type ControllerOwnableOwnershipTransferredIterator struct {
	Event *ControllerOwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerOwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerOwnableOwnershipTransferred)
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
		it.Event = new(ControllerOwnableOwnershipTransferred)
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
func (it *ControllerOwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerOwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerOwnableOwnershipTransferred represents a OwnershipTransferred event raised by the ControllerOwnable contract.
type ControllerOwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerOwnable *ControllerOwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerOwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ControllerOwnable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnableOwnershipTransferredIterator{contract: _ControllerOwnable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ControllerOwnable *ControllerOwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerOwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ControllerOwnable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerOwnableOwnershipTransferred)
				if err := _ControllerOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ControllerOwnable *ControllerOwnableFilterer) ParseOwnershipTransferred(log types.Log) (*ControllerOwnableOwnershipTransferred, error) {
	event := new(ControllerOwnableOwnershipTransferred)
	if err := _ControllerOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
