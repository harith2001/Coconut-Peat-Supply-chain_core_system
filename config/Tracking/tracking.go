// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tracking

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

// TrackingMetaData contains all meta data concerning the Tracking contract.
var TrackingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"shipmentId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ShipmentCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_shipmentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_completedStep\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_acceptedCount\",\"type\":\"uint256\"}],\"name\":\"createShipment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_shipmentId\",\"type\":\"string\"}],\"name\":\"getShipment\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"shipments\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"shipmentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"completedStep\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"acceptedCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TrackingABI is the input ABI used to generate the binding from.
// Deprecated: Use TrackingMetaData.ABI instead.
var TrackingABI = TrackingMetaData.ABI

// Tracking is an auto generated Go binding around an Ethereum contract.
type Tracking struct {
	TrackingCaller     // Read-only binding to the contract
	TrackingTransactor // Write-only binding to the contract
	TrackingFilterer   // Log filterer for contract events
}

// TrackingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrackingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrackingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrackingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrackingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrackingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrackingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrackingSession struct {
	Contract     *Tracking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrackingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrackingCallerSession struct {
	Contract *TrackingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TrackingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrackingTransactorSession struct {
	Contract     *TrackingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TrackingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrackingRaw struct {
	Contract *Tracking // Generic contract binding to access the raw methods on
}

// TrackingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrackingCallerRaw struct {
	Contract *TrackingCaller // Generic read-only contract binding to access the raw methods on
}

// TrackingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrackingTransactorRaw struct {
	Contract *TrackingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTracking creates a new instance of Tracking, bound to a specific deployed contract.
func NewTracking(address common.Address, backend bind.ContractBackend) (*Tracking, error) {
	contract, err := bindTracking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tracking{TrackingCaller: TrackingCaller{contract: contract}, TrackingTransactor: TrackingTransactor{contract: contract}, TrackingFilterer: TrackingFilterer{contract: contract}}, nil
}

// NewTrackingCaller creates a new read-only instance of Tracking, bound to a specific deployed contract.
func NewTrackingCaller(address common.Address, caller bind.ContractCaller) (*TrackingCaller, error) {
	contract, err := bindTracking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrackingCaller{contract: contract}, nil
}

// NewTrackingTransactor creates a new write-only instance of Tracking, bound to a specific deployed contract.
func NewTrackingTransactor(address common.Address, transactor bind.ContractTransactor) (*TrackingTransactor, error) {
	contract, err := bindTracking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrackingTransactor{contract: contract}, nil
}

// NewTrackingFilterer creates a new log filterer instance of Tracking, bound to a specific deployed contract.
func NewTrackingFilterer(address common.Address, filterer bind.ContractFilterer) (*TrackingFilterer, error) {
	contract, err := bindTracking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrackingFilterer{contract: contract}, nil
}

// bindTracking binds a generic wrapper to an already deployed contract.
func bindTracking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrackingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tracking *TrackingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tracking.Contract.TrackingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tracking *TrackingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tracking.Contract.TrackingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tracking *TrackingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tracking.Contract.TrackingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tracking *TrackingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tracking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tracking *TrackingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tracking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tracking *TrackingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tracking.Contract.contract.Transact(opts, method, params...)
}

// GetShipment is a free data retrieval call binding the contract method 0xda34f3d5.
//
// Solidity: function getShipment(string _shipmentId) view returns(string, address, address, string, uint256)
func (_Tracking *TrackingCaller) GetShipment(opts *bind.CallOpts, _shipmentId string) (string, common.Address, common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "getShipment", _shipmentId)

	if err != nil {
		return *new(string), *new(common.Address), *new(common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetShipment is a free data retrieval call binding the contract method 0xda34f3d5.
//
// Solidity: function getShipment(string _shipmentId) view returns(string, address, address, string, uint256)
func (_Tracking *TrackingSession) GetShipment(_shipmentId string) (string, common.Address, common.Address, string, *big.Int, error) {
	return _Tracking.Contract.GetShipment(&_Tracking.CallOpts, _shipmentId)
}

// GetShipment is a free data retrieval call binding the contract method 0xda34f3d5.
//
// Solidity: function getShipment(string _shipmentId) view returns(string, address, address, string, uint256)
func (_Tracking *TrackingCallerSession) GetShipment(_shipmentId string) (string, common.Address, common.Address, string, *big.Int, error) {
	return _Tracking.Contract.GetShipment(&_Tracking.CallOpts, _shipmentId)
}

// Shipments is a free data retrieval call binding the contract method 0x5ceaef5b.
//
// Solidity: function shipments(string ) view returns(string shipmentId, address owner, address receiver, string completedStep, uint256 acceptedCount, bool exists)
func (_Tracking *TrackingCaller) Shipments(opts *bind.CallOpts, arg0 string) (struct {
	ShipmentId    string
	Owner         common.Address
	Receiver      common.Address
	CompletedStep string
	AcceptedCount *big.Int
	Exists        bool
}, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "shipments", arg0)

	outstruct := new(struct {
		ShipmentId    string
		Owner         common.Address
		Receiver      common.Address
		CompletedStep string
		AcceptedCount *big.Int
		Exists        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ShipmentId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CompletedStep = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.AcceptedCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Shipments is a free data retrieval call binding the contract method 0x5ceaef5b.
//
// Solidity: function shipments(string ) view returns(string shipmentId, address owner, address receiver, string completedStep, uint256 acceptedCount, bool exists)
func (_Tracking *TrackingSession) Shipments(arg0 string) (struct {
	ShipmentId    string
	Owner         common.Address
	Receiver      common.Address
	CompletedStep string
	AcceptedCount *big.Int
	Exists        bool
}, error) {
	return _Tracking.Contract.Shipments(&_Tracking.CallOpts, arg0)
}

// Shipments is a free data retrieval call binding the contract method 0x5ceaef5b.
//
// Solidity: function shipments(string ) view returns(string shipmentId, address owner, address receiver, string completedStep, uint256 acceptedCount, bool exists)
func (_Tracking *TrackingCallerSession) Shipments(arg0 string) (struct {
	ShipmentId    string
	Owner         common.Address
	Receiver      common.Address
	CompletedStep string
	AcceptedCount *big.Int
	Exists        bool
}, error) {
	return _Tracking.Contract.Shipments(&_Tracking.CallOpts, arg0)
}

// CreateShipment is a paid mutator transaction binding the contract method 0x17e551a2.
//
// Solidity: function createShipment(string _shipmentId, address _receiver, string _completedStep, uint256 _acceptedCount) returns()
func (_Tracking *TrackingTransactor) CreateShipment(opts *bind.TransactOpts, _shipmentId string, _receiver common.Address, _completedStep string, _acceptedCount *big.Int) (*types.Transaction, error) {
	return _Tracking.contract.Transact(opts, "createShipment", _shipmentId, _receiver, _completedStep, _acceptedCount)
}

// CreateShipment is a paid mutator transaction binding the contract method 0x17e551a2.
//
// Solidity: function createShipment(string _shipmentId, address _receiver, string _completedStep, uint256 _acceptedCount) returns()
func (_Tracking *TrackingSession) CreateShipment(_shipmentId string, _receiver common.Address, _completedStep string, _acceptedCount *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.CreateShipment(&_Tracking.TransactOpts, _shipmentId, _receiver, _completedStep, _acceptedCount)
}

// CreateShipment is a paid mutator transaction binding the contract method 0x17e551a2.
//
// Solidity: function createShipment(string _shipmentId, address _receiver, string _completedStep, uint256 _acceptedCount) returns()
func (_Tracking *TrackingTransactorSession) CreateShipment(_shipmentId string, _receiver common.Address, _completedStep string, _acceptedCount *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.CreateShipment(&_Tracking.TransactOpts, _shipmentId, _receiver, _completedStep, _acceptedCount)
}

// TrackingShipmentCreatedIterator is returned from FilterShipmentCreated and is used to iterate over the raw logs and unpacked data for ShipmentCreated events raised by the Tracking contract.
type TrackingShipmentCreatedIterator struct {
	Event *TrackingShipmentCreated // Event containing the contract specifics and raw log

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
func (it *TrackingShipmentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrackingShipmentCreated)
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
		it.Event = new(TrackingShipmentCreated)
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
func (it *TrackingShipmentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrackingShipmentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrackingShipmentCreated represents a ShipmentCreated event raised by the Tracking contract.
type TrackingShipmentCreated struct {
	ShipmentId string
	Owner      common.Address
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterShipmentCreated is a free log retrieval operation binding the contract event 0xa2e1697fc418aa0092973731742924954c645da2e7c8e9431707e4458b83ec80.
//
// Solidity: event ShipmentCreated(string shipmentId, address indexed owner, address indexed receiver)
func (_Tracking *TrackingFilterer) FilterShipmentCreated(opts *bind.FilterOpts, owner []common.Address, receiver []common.Address) (*TrackingShipmentCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.FilterLogs(opts, "ShipmentCreated", ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrackingShipmentCreatedIterator{contract: _Tracking.contract, event: "ShipmentCreated", logs: logs, sub: sub}, nil
}

// WatchShipmentCreated is a free log subscription operation binding the contract event 0xa2e1697fc418aa0092973731742924954c645da2e7c8e9431707e4458b83ec80.
//
// Solidity: event ShipmentCreated(string shipmentId, address indexed owner, address indexed receiver)
func (_Tracking *TrackingFilterer) WatchShipmentCreated(opts *bind.WatchOpts, sink chan<- *TrackingShipmentCreated, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.WatchLogs(opts, "ShipmentCreated", ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrackingShipmentCreated)
				if err := _Tracking.contract.UnpackLog(event, "ShipmentCreated", log); err != nil {
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

// ParseShipmentCreated is a log parse operation binding the contract event 0xa2e1697fc418aa0092973731742924954c645da2e7c8e9431707e4458b83ec80.
//
// Solidity: event ShipmentCreated(string shipmentId, address indexed owner, address indexed receiver)
func (_Tracking *TrackingFilterer) ParseShipmentCreated(log types.Log) (*TrackingShipmentCreated, error) {
	event := new(TrackingShipmentCreated)
	if err := _Tracking.contract.UnpackLog(event, "ShipmentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
