// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EventsABI is the input ABI used to generate the binding from.
const EventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accuserBondReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"embalmerBondReward\",\"type\":\"uint256\"}],\"name\":\"AccuseArchaeologist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"BurySarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"CancelSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cleaner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cleanerBondReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"embalmerBondReward\",\"type\":\"uint256\"}],\"name\":\"CleanUpSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"archaeologistPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"embalmer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipientPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"name\":\"CreateSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sarcophagusContract\",\"type\":\"address\"}],\"name\":\"Creation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"}],\"name\":\"RegisterArchaeologist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"name\":\"RewrapSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"privatekey\",\"type\":\"bytes32\"}],\"name\":\"UnwrapSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedBond\",\"type\":\"uint256\"}],\"name\":\"UpdateArchaeologist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"}],\"name\":\"UpdateArchaeologistPublicKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"}],\"name\":\"UpdateSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnBond\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFreeBond\",\"type\":\"event\"}]"

// Events is an auto generated Go binding around an Ethereum contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// EventsAccuseArchaeologistIterator is returned from FilterAccuseArchaeologist and is used to iterate over the raw logs and unpacked data for AccuseArchaeologist events raised by the Events contract.
type EventsAccuseArchaeologistIterator struct {
	Event *EventsAccuseArchaeologist // Event containing the contract specifics and raw log

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
func (it *EventsAccuseArchaeologistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsAccuseArchaeologist)
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
		it.Event = new(EventsAccuseArchaeologist)
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
func (it *EventsAccuseArchaeologistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsAccuseArchaeologistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsAccuseArchaeologist represents a AccuseArchaeologist event raised by the Events contract.
type EventsAccuseArchaeologist struct {
	AssetDoubleHash    [32]byte
	Accuser            common.Address
	AccuserBondReward  *big.Int
	EmbalmerBondReward *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAccuseArchaeologist is a free log retrieval operation binding the contract event 0x6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd4.
//
// Solidity: event AccuseArchaeologist(bytes32 indexed assetDoubleHash, address indexed accuser, uint256 accuserBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) FilterAccuseArchaeologist(opts *bind.FilterOpts, assetDoubleHash [][32]byte, accuser []common.Address) (*EventsAccuseArchaeologistIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "AccuseArchaeologist", assetDoubleHashRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return &EventsAccuseArchaeologistIterator{contract: _Events.contract, event: "AccuseArchaeologist", logs: logs, sub: sub}, nil
}

// WatchAccuseArchaeologist is a free log subscription operation binding the contract event 0x6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd4.
//
// Solidity: event AccuseArchaeologist(bytes32 indexed assetDoubleHash, address indexed accuser, uint256 accuserBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) WatchAccuseArchaeologist(opts *bind.WatchOpts, sink chan<- *EventsAccuseArchaeologist, assetDoubleHash [][32]byte, accuser []common.Address) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "AccuseArchaeologist", assetDoubleHashRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsAccuseArchaeologist)
				if err := _Events.contract.UnpackLog(event, "AccuseArchaeologist", log); err != nil {
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

// ParseAccuseArchaeologist is a log parse operation binding the contract event 0x6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd4.
//
// Solidity: event AccuseArchaeologist(bytes32 indexed assetDoubleHash, address indexed accuser, uint256 accuserBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) ParseAccuseArchaeologist(log types.Log) (*EventsAccuseArchaeologist, error) {
	event := new(EventsAccuseArchaeologist)
	if err := _Events.contract.UnpackLog(event, "AccuseArchaeologist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsBurySarcophagusIterator is returned from FilterBurySarcophagus and is used to iterate over the raw logs and unpacked data for BurySarcophagus events raised by the Events contract.
type EventsBurySarcophagusIterator struct {
	Event *EventsBurySarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsBurySarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsBurySarcophagus)
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
		it.Event = new(EventsBurySarcophagus)
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
func (it *EventsBurySarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsBurySarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsBurySarcophagus represents a BurySarcophagus event raised by the Events contract.
type EventsBurySarcophagus struct {
	AssetDoubleHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurySarcophagus is a free log retrieval operation binding the contract event 0x3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a.
//
// Solidity: event BurySarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) FilterBurySarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsBurySarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "BurySarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsBurySarcophagusIterator{contract: _Events.contract, event: "BurySarcophagus", logs: logs, sub: sub}, nil
}

// WatchBurySarcophagus is a free log subscription operation binding the contract event 0x3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a.
//
// Solidity: event BurySarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) WatchBurySarcophagus(opts *bind.WatchOpts, sink chan<- *EventsBurySarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "BurySarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsBurySarcophagus)
				if err := _Events.contract.UnpackLog(event, "BurySarcophagus", log); err != nil {
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

// ParseBurySarcophagus is a log parse operation binding the contract event 0x3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a.
//
// Solidity: event BurySarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) ParseBurySarcophagus(log types.Log) (*EventsBurySarcophagus, error) {
	event := new(EventsBurySarcophagus)
	if err := _Events.contract.UnpackLog(event, "BurySarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCancelSarcophagusIterator is returned from FilterCancelSarcophagus and is used to iterate over the raw logs and unpacked data for CancelSarcophagus events raised by the Events contract.
type EventsCancelSarcophagusIterator struct {
	Event *EventsCancelSarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsCancelSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCancelSarcophagus)
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
		it.Event = new(EventsCancelSarcophagus)
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
func (it *EventsCancelSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCancelSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCancelSarcophagus represents a CancelSarcophagus event raised by the Events contract.
type EventsCancelSarcophagus struct {
	AssetDoubleHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelSarcophagus is a free log retrieval operation binding the contract event 0x829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e7750.
//
// Solidity: event CancelSarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) FilterCancelSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsCancelSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CancelSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsCancelSarcophagusIterator{contract: _Events.contract, event: "CancelSarcophagus", logs: logs, sub: sub}, nil
}

// WatchCancelSarcophagus is a free log subscription operation binding the contract event 0x829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e7750.
//
// Solidity: event CancelSarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) WatchCancelSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsCancelSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CancelSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCancelSarcophagus)
				if err := _Events.contract.UnpackLog(event, "CancelSarcophagus", log); err != nil {
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

// ParseCancelSarcophagus is a log parse operation binding the contract event 0x829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e7750.
//
// Solidity: event CancelSarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) ParseCancelSarcophagus(log types.Log) (*EventsCancelSarcophagus, error) {
	event := new(EventsCancelSarcophagus)
	if err := _Events.contract.UnpackLog(event, "CancelSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCleanUpSarcophagusIterator is returned from FilterCleanUpSarcophagus and is used to iterate over the raw logs and unpacked data for CleanUpSarcophagus events raised by the Events contract.
type EventsCleanUpSarcophagusIterator struct {
	Event *EventsCleanUpSarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsCleanUpSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCleanUpSarcophagus)
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
		it.Event = new(EventsCleanUpSarcophagus)
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
func (it *EventsCleanUpSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCleanUpSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCleanUpSarcophagus represents a CleanUpSarcophagus event raised by the Events contract.
type EventsCleanUpSarcophagus struct {
	AssetDoubleHash    [32]byte
	Cleaner            common.Address
	CleanerBondReward  *big.Int
	EmbalmerBondReward *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCleanUpSarcophagus is a free log retrieval operation binding the contract event 0x106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb9.
//
// Solidity: event CleanUpSarcophagus(bytes32 indexed assetDoubleHash, address indexed cleaner, uint256 cleanerBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) FilterCleanUpSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte, cleaner []common.Address) (*EventsCleanUpSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var cleanerRule []interface{}
	for _, cleanerItem := range cleaner {
		cleanerRule = append(cleanerRule, cleanerItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CleanUpSarcophagus", assetDoubleHashRule, cleanerRule)
	if err != nil {
		return nil, err
	}
	return &EventsCleanUpSarcophagusIterator{contract: _Events.contract, event: "CleanUpSarcophagus", logs: logs, sub: sub}, nil
}

// WatchCleanUpSarcophagus is a free log subscription operation binding the contract event 0x106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb9.
//
// Solidity: event CleanUpSarcophagus(bytes32 indexed assetDoubleHash, address indexed cleaner, uint256 cleanerBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) WatchCleanUpSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsCleanUpSarcophagus, assetDoubleHash [][32]byte, cleaner []common.Address) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var cleanerRule []interface{}
	for _, cleanerItem := range cleaner {
		cleanerRule = append(cleanerRule, cleanerItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CleanUpSarcophagus", assetDoubleHashRule, cleanerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCleanUpSarcophagus)
				if err := _Events.contract.UnpackLog(event, "CleanUpSarcophagus", log); err != nil {
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

// ParseCleanUpSarcophagus is a log parse operation binding the contract event 0x106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb9.
//
// Solidity: event CleanUpSarcophagus(bytes32 indexed assetDoubleHash, address indexed cleaner, uint256 cleanerBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) ParseCleanUpSarcophagus(log types.Log) (*EventsCleanUpSarcophagus, error) {
	event := new(EventsCleanUpSarcophagus)
	if err := _Events.contract.UnpackLog(event, "CleanUpSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCreateSarcophagusIterator is returned from FilterCreateSarcophagus and is used to iterate over the raw logs and unpacked data for CreateSarcophagus events raised by the Events contract.
type EventsCreateSarcophagusIterator struct {
	Event *EventsCreateSarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsCreateSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCreateSarcophagus)
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
		it.Event = new(EventsCreateSarcophagus)
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
func (it *EventsCreateSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCreateSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCreateSarcophagus represents a CreateSarcophagus event raised by the Events contract.
type EventsCreateSarcophagus struct {
	AssetDoubleHash        [32]byte
	Archaeologist          common.Address
	ArchaeologistPublicKey []byte
	Embalmer               common.Address
	Name                   string
	ResurrectionTime       *big.Int
	ResurrectionWindow     *big.Int
	StorageFee             *big.Int
	DiggingFee             *big.Int
	Bounty                 *big.Int
	RecipientPublicKey     []byte
	CursedBond             *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCreateSarcophagus is a free log retrieval operation binding the contract event 0x61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f850.
//
// Solidity: event CreateSarcophagus(bytes32 indexed assetDoubleHash, address indexed archaeologist, bytes archaeologistPublicKey, address embalmer, string name, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes recipientPublicKey, uint256 cursedBond)
func (_Events *EventsFilterer) FilterCreateSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte, archaeologist []common.Address) (*EventsCreateSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CreateSarcophagus", assetDoubleHashRule, archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsCreateSarcophagusIterator{contract: _Events.contract, event: "CreateSarcophagus", logs: logs, sub: sub}, nil
}

// WatchCreateSarcophagus is a free log subscription operation binding the contract event 0x61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f850.
//
// Solidity: event CreateSarcophagus(bytes32 indexed assetDoubleHash, address indexed archaeologist, bytes archaeologistPublicKey, address embalmer, string name, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes recipientPublicKey, uint256 cursedBond)
func (_Events *EventsFilterer) WatchCreateSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsCreateSarcophagus, assetDoubleHash [][32]byte, archaeologist []common.Address) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CreateSarcophagus", assetDoubleHashRule, archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCreateSarcophagus)
				if err := _Events.contract.UnpackLog(event, "CreateSarcophagus", log); err != nil {
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

// ParseCreateSarcophagus is a log parse operation binding the contract event 0x61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f850.
//
// Solidity: event CreateSarcophagus(bytes32 indexed assetDoubleHash, address indexed archaeologist, bytes archaeologistPublicKey, address embalmer, string name, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes recipientPublicKey, uint256 cursedBond)
func (_Events *EventsFilterer) ParseCreateSarcophagus(log types.Log) (*EventsCreateSarcophagus, error) {
	event := new(EventsCreateSarcophagus)
	if err := _Events.contract.UnpackLog(event, "CreateSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCreationIterator is returned from FilterCreation and is used to iterate over the raw logs and unpacked data for Creation events raised by the Events contract.
type EventsCreationIterator struct {
	Event *EventsCreation // Event containing the contract specifics and raw log

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
func (it *EventsCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCreation)
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
		it.Event = new(EventsCreation)
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
func (it *EventsCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCreation represents a Creation event raised by the Events contract.
type EventsCreation struct {
	SarcophagusContract common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCreation is a free log retrieval operation binding the contract event 0x2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c.
//
// Solidity: event Creation(address sarcophagusContract)
func (_Events *EventsFilterer) FilterCreation(opts *bind.FilterOpts) (*EventsCreationIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "Creation")
	if err != nil {
		return nil, err
	}
	return &EventsCreationIterator{contract: _Events.contract, event: "Creation", logs: logs, sub: sub}, nil
}

// WatchCreation is a free log subscription operation binding the contract event 0x2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c.
//
// Solidity: event Creation(address sarcophagusContract)
func (_Events *EventsFilterer) WatchCreation(opts *bind.WatchOpts, sink chan<- *EventsCreation) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "Creation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCreation)
				if err := _Events.contract.UnpackLog(event, "Creation", log); err != nil {
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

// ParseCreation is a log parse operation binding the contract event 0x2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c.
//
// Solidity: event Creation(address sarcophagusContract)
func (_Events *EventsFilterer) ParseCreation(log types.Log) (*EventsCreation, error) {
	event := new(EventsCreation)
	if err := _Events.contract.UnpackLog(event, "Creation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsRegisterArchaeologistIterator is returned from FilterRegisterArchaeologist and is used to iterate over the raw logs and unpacked data for RegisterArchaeologist events raised by the Events contract.
type EventsRegisterArchaeologistIterator struct {
	Event *EventsRegisterArchaeologist // Event containing the contract specifics and raw log

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
func (it *EventsRegisterArchaeologistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsRegisterArchaeologist)
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
		it.Event = new(EventsRegisterArchaeologist)
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
func (it *EventsRegisterArchaeologistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsRegisterArchaeologistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsRegisterArchaeologist represents a RegisterArchaeologist event raised by the Events contract.
type EventsRegisterArchaeologist struct {
	Archaeologist           common.Address
	CurrentPublicKey        []byte
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	Bond                    *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRegisterArchaeologist is a free log retrieval operation binding the contract event 0x6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f.
//
// Solidity: event RegisterArchaeologist(address indexed archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 bond)
func (_Events *EventsFilterer) FilterRegisterArchaeologist(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsRegisterArchaeologistIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "RegisterArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsRegisterArchaeologistIterator{contract: _Events.contract, event: "RegisterArchaeologist", logs: logs, sub: sub}, nil
}

// WatchRegisterArchaeologist is a free log subscription operation binding the contract event 0x6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f.
//
// Solidity: event RegisterArchaeologist(address indexed archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 bond)
func (_Events *EventsFilterer) WatchRegisterArchaeologist(opts *bind.WatchOpts, sink chan<- *EventsRegisterArchaeologist, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "RegisterArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsRegisterArchaeologist)
				if err := _Events.contract.UnpackLog(event, "RegisterArchaeologist", log); err != nil {
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

// ParseRegisterArchaeologist is a log parse operation binding the contract event 0x6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f.
//
// Solidity: event RegisterArchaeologist(address indexed archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 bond)
func (_Events *EventsFilterer) ParseRegisterArchaeologist(log types.Log) (*EventsRegisterArchaeologist, error) {
	event := new(EventsRegisterArchaeologist)
	if err := _Events.contract.UnpackLog(event, "RegisterArchaeologist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsRewrapSarcophagusIterator is returned from FilterRewrapSarcophagus and is used to iterate over the raw logs and unpacked data for RewrapSarcophagus events raised by the Events contract.
type EventsRewrapSarcophagusIterator struct {
	Event *EventsRewrapSarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsRewrapSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsRewrapSarcophagus)
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
		it.Event = new(EventsRewrapSarcophagus)
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
func (it *EventsRewrapSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsRewrapSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsRewrapSarcophagus represents a RewrapSarcophagus event raised by the Events contract.
type EventsRewrapSarcophagus struct {
	AssetId            string
	AssetDoubleHash    [32]byte
	ResurrectionTime   *big.Int
	ResurrectionWindow *big.Int
	DiggingFee         *big.Int
	Bounty             *big.Int
	CursedBond         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewrapSarcophagus is a free log retrieval operation binding the contract event 0x17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46.
//
// Solidity: event RewrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 diggingFee, uint256 bounty, uint256 cursedBond)
func (_Events *EventsFilterer) FilterRewrapSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsRewrapSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "RewrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsRewrapSarcophagusIterator{contract: _Events.contract, event: "RewrapSarcophagus", logs: logs, sub: sub}, nil
}

// WatchRewrapSarcophagus is a free log subscription operation binding the contract event 0x17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46.
//
// Solidity: event RewrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 diggingFee, uint256 bounty, uint256 cursedBond)
func (_Events *EventsFilterer) WatchRewrapSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsRewrapSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "RewrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsRewrapSarcophagus)
				if err := _Events.contract.UnpackLog(event, "RewrapSarcophagus", log); err != nil {
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

// ParseRewrapSarcophagus is a log parse operation binding the contract event 0x17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46.
//
// Solidity: event RewrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 diggingFee, uint256 bounty, uint256 cursedBond)
func (_Events *EventsFilterer) ParseRewrapSarcophagus(log types.Log) (*EventsRewrapSarcophagus, error) {
	event := new(EventsRewrapSarcophagus)
	if err := _Events.contract.UnpackLog(event, "RewrapSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUnwrapSarcophagusIterator is returned from FilterUnwrapSarcophagus and is used to iterate over the raw logs and unpacked data for UnwrapSarcophagus events raised by the Events contract.
type EventsUnwrapSarcophagusIterator struct {
	Event *EventsUnwrapSarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsUnwrapSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUnwrapSarcophagus)
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
		it.Event = new(EventsUnwrapSarcophagus)
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
func (it *EventsUnwrapSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUnwrapSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUnwrapSarcophagus represents a UnwrapSarcophagus event raised by the Events contract.
type EventsUnwrapSarcophagus struct {
	AssetId         string
	AssetDoubleHash [32]byte
	SingleHash      []byte
	Privatekey      [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnwrapSarcophagus is a free log retrieval operation binding the contract event 0x8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c788.
//
// Solidity: event UnwrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, bytes singleHash, bytes32 privatekey)
func (_Events *EventsFilterer) FilterUnwrapSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsUnwrapSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UnwrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsUnwrapSarcophagusIterator{contract: _Events.contract, event: "UnwrapSarcophagus", logs: logs, sub: sub}, nil
}

// WatchUnwrapSarcophagus is a free log subscription operation binding the contract event 0x8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c788.
//
// Solidity: event UnwrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, bytes singleHash, bytes32 privatekey)
func (_Events *EventsFilterer) WatchUnwrapSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsUnwrapSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UnwrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUnwrapSarcophagus)
				if err := _Events.contract.UnpackLog(event, "UnwrapSarcophagus", log); err != nil {
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

// ParseUnwrapSarcophagus is a log parse operation binding the contract event 0x8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c788.
//
// Solidity: event UnwrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, bytes singleHash, bytes32 privatekey)
func (_Events *EventsFilterer) ParseUnwrapSarcophagus(log types.Log) (*EventsUnwrapSarcophagus, error) {
	event := new(EventsUnwrapSarcophagus)
	if err := _Events.contract.UnpackLog(event, "UnwrapSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUpdateArchaeologistIterator is returned from FilterUpdateArchaeologist and is used to iterate over the raw logs and unpacked data for UpdateArchaeologist events raised by the Events contract.
type EventsUpdateArchaeologistIterator struct {
	Event *EventsUpdateArchaeologist // Event containing the contract specifics and raw log

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
func (it *EventsUpdateArchaeologistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUpdateArchaeologist)
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
		it.Event = new(EventsUpdateArchaeologist)
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
func (it *EventsUpdateArchaeologistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUpdateArchaeologistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUpdateArchaeologist represents a UpdateArchaeologist event raised by the Events contract.
type EventsUpdateArchaeologist struct {
	Archaeologist           common.Address
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	AddedBond               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdateArchaeologist is a free log retrieval operation binding the contract event 0x8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f.
//
// Solidity: event UpdateArchaeologist(address indexed archaeologist, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 addedBond)
func (_Events *EventsFilterer) FilterUpdateArchaeologist(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsUpdateArchaeologistIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UpdateArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsUpdateArchaeologistIterator{contract: _Events.contract, event: "UpdateArchaeologist", logs: logs, sub: sub}, nil
}

// WatchUpdateArchaeologist is a free log subscription operation binding the contract event 0x8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f.
//
// Solidity: event UpdateArchaeologist(address indexed archaeologist, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 addedBond)
func (_Events *EventsFilterer) WatchUpdateArchaeologist(opts *bind.WatchOpts, sink chan<- *EventsUpdateArchaeologist, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UpdateArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUpdateArchaeologist)
				if err := _Events.contract.UnpackLog(event, "UpdateArchaeologist", log); err != nil {
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

// ParseUpdateArchaeologist is a log parse operation binding the contract event 0x8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f.
//
// Solidity: event UpdateArchaeologist(address indexed archaeologist, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 addedBond)
func (_Events *EventsFilterer) ParseUpdateArchaeologist(log types.Log) (*EventsUpdateArchaeologist, error) {
	event := new(EventsUpdateArchaeologist)
	if err := _Events.contract.UnpackLog(event, "UpdateArchaeologist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUpdateArchaeologistPublicKeyIterator is returned from FilterUpdateArchaeologistPublicKey and is used to iterate over the raw logs and unpacked data for UpdateArchaeologistPublicKey events raised by the Events contract.
type EventsUpdateArchaeologistPublicKeyIterator struct {
	Event *EventsUpdateArchaeologistPublicKey // Event containing the contract specifics and raw log

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
func (it *EventsUpdateArchaeologistPublicKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUpdateArchaeologistPublicKey)
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
		it.Event = new(EventsUpdateArchaeologistPublicKey)
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
func (it *EventsUpdateArchaeologistPublicKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUpdateArchaeologistPublicKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUpdateArchaeologistPublicKey represents a UpdateArchaeologistPublicKey event raised by the Events contract.
type EventsUpdateArchaeologistPublicKey struct {
	Archaeologist    common.Address
	CurrentPublicKey []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateArchaeologistPublicKey is a free log retrieval operation binding the contract event 0x6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17.
//
// Solidity: event UpdateArchaeologistPublicKey(address indexed archaeologist, bytes currentPublicKey)
func (_Events *EventsFilterer) FilterUpdateArchaeologistPublicKey(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsUpdateArchaeologistPublicKeyIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UpdateArchaeologistPublicKey", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsUpdateArchaeologistPublicKeyIterator{contract: _Events.contract, event: "UpdateArchaeologistPublicKey", logs: logs, sub: sub}, nil
}

// WatchUpdateArchaeologistPublicKey is a free log subscription operation binding the contract event 0x6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17.
//
// Solidity: event UpdateArchaeologistPublicKey(address indexed archaeologist, bytes currentPublicKey)
func (_Events *EventsFilterer) WatchUpdateArchaeologistPublicKey(opts *bind.WatchOpts, sink chan<- *EventsUpdateArchaeologistPublicKey, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UpdateArchaeologistPublicKey", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUpdateArchaeologistPublicKey)
				if err := _Events.contract.UnpackLog(event, "UpdateArchaeologistPublicKey", log); err != nil {
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

// ParseUpdateArchaeologistPublicKey is a log parse operation binding the contract event 0x6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17.
//
// Solidity: event UpdateArchaeologistPublicKey(address indexed archaeologist, bytes currentPublicKey)
func (_Events *EventsFilterer) ParseUpdateArchaeologistPublicKey(log types.Log) (*EventsUpdateArchaeologistPublicKey, error) {
	event := new(EventsUpdateArchaeologistPublicKey)
	if err := _Events.contract.UnpackLog(event, "UpdateArchaeologistPublicKey", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUpdateSarcophagusIterator is returned from FilterUpdateSarcophagus and is used to iterate over the raw logs and unpacked data for UpdateSarcophagus events raised by the Events contract.
type EventsUpdateSarcophagusIterator struct {
	Event *EventsUpdateSarcophagus // Event containing the contract specifics and raw log

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
func (it *EventsUpdateSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUpdateSarcophagus)
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
		it.Event = new(EventsUpdateSarcophagus)
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
func (it *EventsUpdateSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUpdateSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUpdateSarcophagus represents a UpdateSarcophagus event raised by the Events contract.
type EventsUpdateSarcophagus struct {
	AssetDoubleHash [32]byte
	AssetId         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateSarcophagus is a free log retrieval operation binding the contract event 0x2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212.
//
// Solidity: event UpdateSarcophagus(bytes32 indexed assetDoubleHash, string assetId)
func (_Events *EventsFilterer) FilterUpdateSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsUpdateSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UpdateSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsUpdateSarcophagusIterator{contract: _Events.contract, event: "UpdateSarcophagus", logs: logs, sub: sub}, nil
}

// WatchUpdateSarcophagus is a free log subscription operation binding the contract event 0x2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212.
//
// Solidity: event UpdateSarcophagus(bytes32 indexed assetDoubleHash, string assetId)
func (_Events *EventsFilterer) WatchUpdateSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsUpdateSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UpdateSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUpdateSarcophagus)
				if err := _Events.contract.UnpackLog(event, "UpdateSarcophagus", log); err != nil {
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

// ParseUpdateSarcophagus is a log parse operation binding the contract event 0x2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212.
//
// Solidity: event UpdateSarcophagus(bytes32 indexed assetDoubleHash, string assetId)
func (_Events *EventsFilterer) ParseUpdateSarcophagus(log types.Log) (*EventsUpdateSarcophagus, error) {
	event := new(EventsUpdateSarcophagus)
	if err := _Events.contract.UnpackLog(event, "UpdateSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsWithdrawalFreeBondIterator is returned from FilterWithdrawalFreeBond and is used to iterate over the raw logs and unpacked data for WithdrawalFreeBond events raised by the Events contract.
type EventsWithdrawalFreeBondIterator struct {
	Event *EventsWithdrawalFreeBond // Event containing the contract specifics and raw log

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
func (it *EventsWithdrawalFreeBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsWithdrawalFreeBond)
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
		it.Event = new(EventsWithdrawalFreeBond)
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
func (it *EventsWithdrawalFreeBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsWithdrawalFreeBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsWithdrawalFreeBond represents a WithdrawalFreeBond event raised by the Events contract.
type EventsWithdrawalFreeBond struct {
	Archaeologist common.Address
	WithdrawnBond *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFreeBond is a free log retrieval operation binding the contract event 0x8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f.
//
// Solidity: event WithdrawalFreeBond(address indexed archaeologist, uint256 withdrawnBond)
func (_Events *EventsFilterer) FilterWithdrawalFreeBond(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsWithdrawalFreeBondIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "WithdrawalFreeBond", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsWithdrawalFreeBondIterator{contract: _Events.contract, event: "WithdrawalFreeBond", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFreeBond is a free log subscription operation binding the contract event 0x8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f.
//
// Solidity: event WithdrawalFreeBond(address indexed archaeologist, uint256 withdrawnBond)
func (_Events *EventsFilterer) WatchWithdrawalFreeBond(opts *bind.WatchOpts, sink chan<- *EventsWithdrawalFreeBond, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "WithdrawalFreeBond", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsWithdrawalFreeBond)
				if err := _Events.contract.UnpackLog(event, "WithdrawalFreeBond", log); err != nil {
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

// ParseWithdrawalFreeBond is a log parse operation binding the contract event 0x8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f.
//
// Solidity: event WithdrawalFreeBond(address indexed archaeologist, uint256 withdrawnBond)
func (_Events *EventsFilterer) ParseWithdrawalFreeBond(log types.Log) (*EventsWithdrawalFreeBond, error) {
	event := new(EventsWithdrawalFreeBond)
	if err := _Events.contract.UnpackLog(event, "WithdrawalFreeBond", log); err != nil {
		return nil, err
	}
	return event, nil
}
