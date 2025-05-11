// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dlt

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

// RequestContractRequest is an auto generated low-level Go binding around an user-defined struct.
type RequestContractRequest struct {
	Id     *big.Int
	Status uint8
}

// RequestContractMetaData contains all meta data concerning the RequestContract contract.
var RequestContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ApprovedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"SavedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approvedAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"approvedBy\",\"type\":\"address\"}],\"name\":\"ApproveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"QueryRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumRequestContract.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structRequestContract.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumRequestContract.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structRequestContract.Request\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"SaveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061030b8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80636789c2de146100435780637e1168d5146100585780639a7c930014610081575b5f5ffd5b6100566100513660046101cd565b610094565b005b61006b610066366004610232565b610112565b604051610078919061025d565b60405180910390f35b61005661008f366004610294565b610178565b5f602082810182815283518352908290526040909120825181559051600180830180548594939260ff19909116908360048111156100d4576100d4610249565b021790555050604051600181527fd45b452d98457e457b349d94bc8ae6692f8ffeb0893a2f050eeefcef2cea3478915060200160405180910390a150565b604080518082019091525f80825260208201525f82815260208181526040918290208251808401909352805483526001810154909183019060ff16600481111561015e5761015e610249565b600481111561016f5761016f610249565b90525092915050565b5f83815260208181526040918290206001808201805460ff1916821790559251928352917f5a82565f9da83bdbfb0b014c0c0d7de93d152b4b2e53e71840aae92af371b868910160405180910390a150505050565b5f60408284031280156101de575f5ffd5b506040805190810167ffffffffffffffff8111828210171561020e57634e487b7160e01b5f52604160045260245ffd5b60405282358152602083013560058110610226575f5ffd5b60208201529392505050565b5f60208284031215610242575f5ffd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b81518152602082015160408201906005811061028757634e487b7160e01b5f52602160045260245ffd5b8060208401525092915050565b5f5f5f606084860312156102a6575f5ffd5b833592506020840135915060408401356001600160a01b03811681146102ca575f5ffd5b80915050925092509256fea26469706673582212208421ed06286dd44b431a927e2a56cde3d6151cef40b97eaafb30b3eed8dd538264736f6c634300081d0033",
}

// RequestContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestContractMetaData.ABI instead.
var RequestContractABI = RequestContractMetaData.ABI

// RequestContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestContractMetaData.Bin instead.
var RequestContractBin = RequestContractMetaData.Bin

// DeployRequestContract deploys a new Ethereum contract, binding an instance of RequestContract to it.
func DeployRequestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestContract, error) {
	parsed, err := RequestContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestContract{RequestContractCaller: RequestContractCaller{contract: contract}, RequestContractTransactor: RequestContractTransactor{contract: contract}, RequestContractFilterer: RequestContractFilterer{contract: contract}}, nil
}

// RequestContract is an auto generated Go binding around an Ethereum contract.
type RequestContract struct {
	RequestContractCaller     // Read-only binding to the contract
	RequestContractTransactor // Write-only binding to the contract
	RequestContractFilterer   // Log filterer for contract events
}

// RequestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestContractSession struct {
	Contract     *RequestContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestContractCallerSession struct {
	Contract *RequestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RequestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestContractTransactorSession struct {
	Contract     *RequestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RequestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestContractRaw struct {
	Contract *RequestContract // Generic contract binding to access the raw methods on
}

// RequestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestContractCallerRaw struct {
	Contract *RequestContractCaller // Generic read-only contract binding to access the raw methods on
}

// RequestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestContractTransactorRaw struct {
	Contract *RequestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestContract creates a new instance of RequestContract, bound to a specific deployed contract.
func NewRequestContract(address common.Address, backend bind.ContractBackend) (*RequestContract, error) {
	contract, err := bindRequestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestContract{RequestContractCaller: RequestContractCaller{contract: contract}, RequestContractTransactor: RequestContractTransactor{contract: contract}, RequestContractFilterer: RequestContractFilterer{contract: contract}}, nil
}

// NewRequestContractCaller creates a new read-only instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractCaller(address common.Address, caller bind.ContractCaller) (*RequestContractCaller, error) {
	contract, err := bindRequestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestContractCaller{contract: contract}, nil
}

// NewRequestContractTransactor creates a new write-only instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestContractTransactor, error) {
	contract, err := bindRequestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestContractTransactor{contract: contract}, nil
}

// NewRequestContractFilterer creates a new log filterer instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestContractFilterer, error) {
	contract, err := bindRequestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestContractFilterer{contract: contract}, nil
}

// bindRequestContract binds a generic wrapper to an already deployed contract.
func bindRequestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequestContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestContract *RequestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestContract.Contract.RequestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestContract *RequestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.Contract.RequestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestContract *RequestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestContract.Contract.RequestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestContract *RequestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestContract *RequestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestContract *RequestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestContract.Contract.contract.Transact(opts, method, params...)
}

// QueryRequest is a free data retrieval call binding the contract method 0x7e1168d5.
//
// Solidity: function QueryRequest(uint256 id) view returns((uint256,uint8))
func (_RequestContract *RequestContractCaller) QueryRequest(opts *bind.CallOpts, id *big.Int) (RequestContractRequest, error) {
	var out []interface{}
	err := _RequestContract.contract.Call(opts, &out, "QueryRequest", id)

	if err != nil {
		return *new(RequestContractRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(RequestContractRequest)).(*RequestContractRequest)

	return out0, err

}

// QueryRequest is a free data retrieval call binding the contract method 0x7e1168d5.
//
// Solidity: function QueryRequest(uint256 id) view returns((uint256,uint8))
func (_RequestContract *RequestContractSession) QueryRequest(id *big.Int) (RequestContractRequest, error) {
	return _RequestContract.Contract.QueryRequest(&_RequestContract.CallOpts, id)
}

// QueryRequest is a free data retrieval call binding the contract method 0x7e1168d5.
//
// Solidity: function QueryRequest(uint256 id) view returns((uint256,uint8))
func (_RequestContract *RequestContractCallerSession) QueryRequest(id *big.Int) (RequestContractRequest, error) {
	return _RequestContract.Contract.QueryRequest(&_RequestContract.CallOpts, id)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x9a7c9300.
//
// Solidity: function ApproveRequest(uint256 id, uint256 approvedAt, address approvedBy) returns()
func (_RequestContract *RequestContractTransactor) ApproveRequest(opts *bind.TransactOpts, id *big.Int, approvedAt *big.Int, approvedBy common.Address) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "ApproveRequest", id, approvedAt, approvedBy)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x9a7c9300.
//
// Solidity: function ApproveRequest(uint256 id, uint256 approvedAt, address approvedBy) returns()
func (_RequestContract *RequestContractSession) ApproveRequest(id *big.Int, approvedAt *big.Int, approvedBy common.Address) (*types.Transaction, error) {
	return _RequestContract.Contract.ApproveRequest(&_RequestContract.TransactOpts, id, approvedAt, approvedBy)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x9a7c9300.
//
// Solidity: function ApproveRequest(uint256 id, uint256 approvedAt, address approvedBy) returns()
func (_RequestContract *RequestContractTransactorSession) ApproveRequest(id *big.Int, approvedAt *big.Int, approvedBy common.Address) (*types.Transaction, error) {
	return _RequestContract.Contract.ApproveRequest(&_RequestContract.TransactOpts, id, approvedAt, approvedBy)
}

// SaveRequest is a paid mutator transaction binding the contract method 0x6789c2de.
//
// Solidity: function SaveRequest((uint256,uint8) _request) returns()
func (_RequestContract *RequestContractTransactor) SaveRequest(opts *bind.TransactOpts, _request RequestContractRequest) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "SaveRequest", _request)
}

// SaveRequest is a paid mutator transaction binding the contract method 0x6789c2de.
//
// Solidity: function SaveRequest((uint256,uint8) _request) returns()
func (_RequestContract *RequestContractSession) SaveRequest(_request RequestContractRequest) (*types.Transaction, error) {
	return _RequestContract.Contract.SaveRequest(&_RequestContract.TransactOpts, _request)
}

// SaveRequest is a paid mutator transaction binding the contract method 0x6789c2de.
//
// Solidity: function SaveRequest((uint256,uint8) _request) returns()
func (_RequestContract *RequestContractTransactorSession) SaveRequest(_request RequestContractRequest) (*types.Transaction, error) {
	return _RequestContract.Contract.SaveRequest(&_RequestContract.TransactOpts, _request)
}

// RequestContractApprovedEventIterator is returned from FilterApprovedEvent and is used to iterate over the raw logs and unpacked data for ApprovedEvent events raised by the RequestContract contract.
type RequestContractApprovedEventIterator struct {
	Event *RequestContractApprovedEvent // Event containing the contract specifics and raw log

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
func (it *RequestContractApprovedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestContractApprovedEvent)
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
		it.Event = new(RequestContractApprovedEvent)
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
func (it *RequestContractApprovedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestContractApprovedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestContractApprovedEvent represents a ApprovedEvent event raised by the RequestContract contract.
type RequestContractApprovedEvent struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApprovedEvent is a free log retrieval operation binding the contract event 0x5a82565f9da83bdbfb0b014c0c0d7de93d152b4b2e53e71840aae92af371b868.
//
// Solidity: event ApprovedEvent(bool status)
func (_RequestContract *RequestContractFilterer) FilterApprovedEvent(opts *bind.FilterOpts) (*RequestContractApprovedEventIterator, error) {

	logs, sub, err := _RequestContract.contract.FilterLogs(opts, "ApprovedEvent")
	if err != nil {
		return nil, err
	}
	return &RequestContractApprovedEventIterator{contract: _RequestContract.contract, event: "ApprovedEvent", logs: logs, sub: sub}, nil
}

// WatchApprovedEvent is a free log subscription operation binding the contract event 0x5a82565f9da83bdbfb0b014c0c0d7de93d152b4b2e53e71840aae92af371b868.
//
// Solidity: event ApprovedEvent(bool status)
func (_RequestContract *RequestContractFilterer) WatchApprovedEvent(opts *bind.WatchOpts, sink chan<- *RequestContractApprovedEvent) (event.Subscription, error) {

	logs, sub, err := _RequestContract.contract.WatchLogs(opts, "ApprovedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestContractApprovedEvent)
				if err := _RequestContract.contract.UnpackLog(event, "ApprovedEvent", log); err != nil {
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

// ParseApprovedEvent is a log parse operation binding the contract event 0x5a82565f9da83bdbfb0b014c0c0d7de93d152b4b2e53e71840aae92af371b868.
//
// Solidity: event ApprovedEvent(bool status)
func (_RequestContract *RequestContractFilterer) ParseApprovedEvent(log types.Log) (*RequestContractApprovedEvent, error) {
	event := new(RequestContractApprovedEvent)
	if err := _RequestContract.contract.UnpackLog(event, "ApprovedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestContractSavedEventIterator is returned from FilterSavedEvent and is used to iterate over the raw logs and unpacked data for SavedEvent events raised by the RequestContract contract.
type RequestContractSavedEventIterator struct {
	Event *RequestContractSavedEvent // Event containing the contract specifics and raw log

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
func (it *RequestContractSavedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestContractSavedEvent)
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
		it.Event = new(RequestContractSavedEvent)
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
func (it *RequestContractSavedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestContractSavedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestContractSavedEvent represents a SavedEvent event raised by the RequestContract contract.
type RequestContractSavedEvent struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSavedEvent is a free log retrieval operation binding the contract event 0xd45b452d98457e457b349d94bc8ae6692f8ffeb0893a2f050eeefcef2cea3478.
//
// Solidity: event SavedEvent(bool status)
func (_RequestContract *RequestContractFilterer) FilterSavedEvent(opts *bind.FilterOpts) (*RequestContractSavedEventIterator, error) {

	logs, sub, err := _RequestContract.contract.FilterLogs(opts, "SavedEvent")
	if err != nil {
		return nil, err
	}
	return &RequestContractSavedEventIterator{contract: _RequestContract.contract, event: "SavedEvent", logs: logs, sub: sub}, nil
}

// WatchSavedEvent is a free log subscription operation binding the contract event 0xd45b452d98457e457b349d94bc8ae6692f8ffeb0893a2f050eeefcef2cea3478.
//
// Solidity: event SavedEvent(bool status)
func (_RequestContract *RequestContractFilterer) WatchSavedEvent(opts *bind.WatchOpts, sink chan<- *RequestContractSavedEvent) (event.Subscription, error) {

	logs, sub, err := _RequestContract.contract.WatchLogs(opts, "SavedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestContractSavedEvent)
				if err := _RequestContract.contract.UnpackLog(event, "SavedEvent", log); err != nil {
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

// ParseSavedEvent is a log parse operation binding the contract event 0xd45b452d98457e457b349d94bc8ae6692f8ffeb0893a2f050eeefcef2cea3478.
//
// Solidity: event SavedEvent(bool status)
func (_RequestContract *RequestContractFilterer) ParseSavedEvent(log types.Log) (*RequestContractSavedEvent, error) {
	event := new(RequestContractSavedEvent)
	if err := _RequestContract.contract.UnpackLog(event, "SavedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
