// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// MuleRegistryMetaData contains all meta data concerning the MuleRegistry contract.
var MuleRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"senderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiverHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MuleReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountReports\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"flaggedAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountHash\",\"type\":\"bytes32\"}],\"name\":\"getReportCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountHash\",\"type\":\"bytes32\"}],\"name\":\"isAccountFlagged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_transactionId\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_senderHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiverHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"reportMule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MuleRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MuleRegistryMetaData.ABI instead.
var MuleRegistryABI = MuleRegistryMetaData.ABI

// MuleRegistry is an auto generated Go binding around an Ethereum contract.
type MuleRegistry struct {
	MuleRegistryCaller     // Read-only binding to the contract
	MuleRegistryTransactor // Write-only binding to the contract
	MuleRegistryFilterer   // Log filterer for contract events
}

// MuleRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuleRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuleRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuleRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuleRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuleRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuleRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuleRegistrySession struct {
	Contract     *MuleRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuleRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuleRegistryCallerSession struct {
	Contract *MuleRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MuleRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuleRegistryTransactorSession struct {
	Contract     *MuleRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MuleRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuleRegistryRaw struct {
	Contract *MuleRegistry // Generic contract binding to access the raw methods on
}

// MuleRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuleRegistryCallerRaw struct {
	Contract *MuleRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MuleRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuleRegistryTransactorRaw struct {
	Contract *MuleRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuleRegistry creates a new instance of MuleRegistry, bound to a specific deployed contract.
func NewMuleRegistry(address common.Address, backend bind.ContractBackend) (*MuleRegistry, error) {
	contract, err := bindMuleRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuleRegistry{MuleRegistryCaller: MuleRegistryCaller{contract: contract}, MuleRegistryTransactor: MuleRegistryTransactor{contract: contract}, MuleRegistryFilterer: MuleRegistryFilterer{contract: contract}}, nil
}

// NewMuleRegistryCaller creates a new read-only instance of MuleRegistry, bound to a specific deployed contract.
func NewMuleRegistryCaller(address common.Address, caller bind.ContractCaller) (*MuleRegistryCaller, error) {
	contract, err := bindMuleRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuleRegistryCaller{contract: contract}, nil
}

// NewMuleRegistryTransactor creates a new write-only instance of MuleRegistry, bound to a specific deployed contract.
func NewMuleRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MuleRegistryTransactor, error) {
	contract, err := bindMuleRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuleRegistryTransactor{contract: contract}, nil
}

// NewMuleRegistryFilterer creates a new log filterer instance of MuleRegistry, bound to a specific deployed contract.
func NewMuleRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MuleRegistryFilterer, error) {
	contract, err := bindMuleRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuleRegistryFilterer{contract: contract}, nil
}

// bindMuleRegistry binds a generic wrapper to an already deployed contract.
func bindMuleRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuleRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuleRegistry *MuleRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuleRegistry.Contract.MuleRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuleRegistry *MuleRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuleRegistry.Contract.MuleRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuleRegistry *MuleRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuleRegistry.Contract.MuleRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuleRegistry *MuleRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuleRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuleRegistry *MuleRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuleRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuleRegistry *MuleRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuleRegistry.Contract.contract.Transact(opts, method, params...)
}

// AccountReports is a free data retrieval call binding the contract method 0x5b0c2bd4.
//
// Solidity: function accountReports(bytes32 , uint256 ) view returns(string transactionId, uint256 amount, string location, uint256 timestamp, string reason, address reporter)
func (_MuleRegistry *MuleRegistryCaller) AccountReports(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	TransactionId string
	Amount        *big.Int
	Location      string
	Timestamp     *big.Int
	Reason        string
	Reporter      common.Address
}, error) {
	var out []interface{}
	err := _MuleRegistry.contract.Call(opts, &out, "accountReports", arg0, arg1)

	outstruct := new(struct {
		TransactionId string
		Amount        *big.Int
		Location      string
		Timestamp     *big.Int
		Reason        string
		Reporter      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransactionId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Reason = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Reporter = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AccountReports is a free data retrieval call binding the contract method 0x5b0c2bd4.
//
// Solidity: function accountReports(bytes32 , uint256 ) view returns(string transactionId, uint256 amount, string location, uint256 timestamp, string reason, address reporter)
func (_MuleRegistry *MuleRegistrySession) AccountReports(arg0 [32]byte, arg1 *big.Int) (struct {
	TransactionId string
	Amount        *big.Int
	Location      string
	Timestamp     *big.Int
	Reason        string
	Reporter      common.Address
}, error) {
	return _MuleRegistry.Contract.AccountReports(&_MuleRegistry.CallOpts, arg0, arg1)
}

// AccountReports is a free data retrieval call binding the contract method 0x5b0c2bd4.
//
// Solidity: function accountReports(bytes32 , uint256 ) view returns(string transactionId, uint256 amount, string location, uint256 timestamp, string reason, address reporter)
func (_MuleRegistry *MuleRegistryCallerSession) AccountReports(arg0 [32]byte, arg1 *big.Int) (struct {
	TransactionId string
	Amount        *big.Int
	Location      string
	Timestamp     *big.Int
	Reason        string
	Reporter      common.Address
}, error) {
	return _MuleRegistry.Contract.AccountReports(&_MuleRegistry.CallOpts, arg0, arg1)
}

// FlaggedAccounts is a free data retrieval call binding the contract method 0x4a0aae58.
//
// Solidity: function flaggedAccounts(bytes32 ) view returns(bool)
func (_MuleRegistry *MuleRegistryCaller) FlaggedAccounts(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MuleRegistry.contract.Call(opts, &out, "flaggedAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FlaggedAccounts is a free data retrieval call binding the contract method 0x4a0aae58.
//
// Solidity: function flaggedAccounts(bytes32 ) view returns(bool)
func (_MuleRegistry *MuleRegistrySession) FlaggedAccounts(arg0 [32]byte) (bool, error) {
	return _MuleRegistry.Contract.FlaggedAccounts(&_MuleRegistry.CallOpts, arg0)
}

// FlaggedAccounts is a free data retrieval call binding the contract method 0x4a0aae58.
//
// Solidity: function flaggedAccounts(bytes32 ) view returns(bool)
func (_MuleRegistry *MuleRegistryCallerSession) FlaggedAccounts(arg0 [32]byte) (bool, error) {
	return _MuleRegistry.Contract.FlaggedAccounts(&_MuleRegistry.CallOpts, arg0)
}

// GetReportCount is a free data retrieval call binding the contract method 0x580317eb.
//
// Solidity: function getReportCount(bytes32 _accountHash) view returns(uint256)
func (_MuleRegistry *MuleRegistryCaller) GetReportCount(opts *bind.CallOpts, _accountHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MuleRegistry.contract.Call(opts, &out, "getReportCount", _accountHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportCount is a free data retrieval call binding the contract method 0x580317eb.
//
// Solidity: function getReportCount(bytes32 _accountHash) view returns(uint256)
func (_MuleRegistry *MuleRegistrySession) GetReportCount(_accountHash [32]byte) (*big.Int, error) {
	return _MuleRegistry.Contract.GetReportCount(&_MuleRegistry.CallOpts, _accountHash)
}

// GetReportCount is a free data retrieval call binding the contract method 0x580317eb.
//
// Solidity: function getReportCount(bytes32 _accountHash) view returns(uint256)
func (_MuleRegistry *MuleRegistryCallerSession) GetReportCount(_accountHash [32]byte) (*big.Int, error) {
	return _MuleRegistry.Contract.GetReportCount(&_MuleRegistry.CallOpts, _accountHash)
}

// IsAccountFlagged is a free data retrieval call binding the contract method 0x675ef4bb.
//
// Solidity: function isAccountFlagged(bytes32 _accountHash) view returns(bool)
func (_MuleRegistry *MuleRegistryCaller) IsAccountFlagged(opts *bind.CallOpts, _accountHash [32]byte) (bool, error) {
	var out []interface{}
	err := _MuleRegistry.contract.Call(opts, &out, "isAccountFlagged", _accountHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountFlagged is a free data retrieval call binding the contract method 0x675ef4bb.
//
// Solidity: function isAccountFlagged(bytes32 _accountHash) view returns(bool)
func (_MuleRegistry *MuleRegistrySession) IsAccountFlagged(_accountHash [32]byte) (bool, error) {
	return _MuleRegistry.Contract.IsAccountFlagged(&_MuleRegistry.CallOpts, _accountHash)
}

// IsAccountFlagged is a free data retrieval call binding the contract method 0x675ef4bb.
//
// Solidity: function isAccountFlagged(bytes32 _accountHash) view returns(bool)
func (_MuleRegistry *MuleRegistryCallerSession) IsAccountFlagged(_accountHash [32]byte) (bool, error) {
	return _MuleRegistry.Contract.IsAccountFlagged(&_MuleRegistry.CallOpts, _accountHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MuleRegistry *MuleRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuleRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MuleRegistry *MuleRegistrySession) Owner() (common.Address, error) {
	return _MuleRegistry.Contract.Owner(&_MuleRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MuleRegistry *MuleRegistryCallerSession) Owner() (common.Address, error) {
	return _MuleRegistry.Contract.Owner(&_MuleRegistry.CallOpts)
}

// ReportMule is a paid mutator transaction binding the contract method 0xda1f0d4c.
//
// Solidity: function reportMule(string _transactionId, bytes32 _senderHash, bytes32 _receiverHash, uint256 _amount, string _location, string _reason) returns()
func (_MuleRegistry *MuleRegistryTransactor) ReportMule(opts *bind.TransactOpts, _transactionId string, _senderHash [32]byte, _receiverHash [32]byte, _amount *big.Int, _location string, _reason string) (*types.Transaction, error) {
	return _MuleRegistry.contract.Transact(opts, "reportMule", _transactionId, _senderHash, _receiverHash, _amount, _location, _reason)
}

// ReportMule is a paid mutator transaction binding the contract method 0xda1f0d4c.
//
// Solidity: function reportMule(string _transactionId, bytes32 _senderHash, bytes32 _receiverHash, uint256 _amount, string _location, string _reason) returns()
func (_MuleRegistry *MuleRegistrySession) ReportMule(_transactionId string, _senderHash [32]byte, _receiverHash [32]byte, _amount *big.Int, _location string, _reason string) (*types.Transaction, error) {
	return _MuleRegistry.Contract.ReportMule(&_MuleRegistry.TransactOpts, _transactionId, _senderHash, _receiverHash, _amount, _location, _reason)
}

// ReportMule is a paid mutator transaction binding the contract method 0xda1f0d4c.
//
// Solidity: function reportMule(string _transactionId, bytes32 _senderHash, bytes32 _receiverHash, uint256 _amount, string _location, string _reason) returns()
func (_MuleRegistry *MuleRegistryTransactorSession) ReportMule(_transactionId string, _senderHash [32]byte, _receiverHash [32]byte, _amount *big.Int, _location string, _reason string) (*types.Transaction, error) {
	return _MuleRegistry.Contract.ReportMule(&_MuleRegistry.TransactOpts, _transactionId, _senderHash, _receiverHash, _amount, _location, _reason)
}

// MuleRegistryMuleReportedIterator is returned from FilterMuleReported and is used to iterate over the raw logs and unpacked data for MuleReported events raised by the MuleRegistry contract.
type MuleRegistryMuleReportedIterator struct {
	Event *MuleRegistryMuleReported // Event containing the contract specifics and raw log

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
func (it *MuleRegistryMuleReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuleRegistryMuleReported)
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
		it.Event = new(MuleRegistryMuleReported)
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
func (it *MuleRegistryMuleReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuleRegistryMuleReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuleRegistryMuleReported represents a MuleReported event raised by the MuleRegistry contract.
type MuleRegistryMuleReported struct {
	SenderHash   [32]byte
	ReceiverHash [32]byte
	Amount       *big.Int
	Reason       string
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMuleReported is a free log retrieval operation binding the contract event 0x83836cc0d166876e68904bae129ac8f5aa73eb9a2ea197f1cf947e3871b10988.
//
// Solidity: event MuleReported(bytes32 indexed senderHash, bytes32 indexed receiverHash, uint256 amount, string reason, uint256 timestamp)
func (_MuleRegistry *MuleRegistryFilterer) FilterMuleReported(opts *bind.FilterOpts, senderHash [][32]byte, receiverHash [][32]byte) (*MuleRegistryMuleReportedIterator, error) {

	var senderHashRule []interface{}
	for _, senderHashItem := range senderHash {
		senderHashRule = append(senderHashRule, senderHashItem)
	}
	var receiverHashRule []interface{}
	for _, receiverHashItem := range receiverHash {
		receiverHashRule = append(receiverHashRule, receiverHashItem)
	}

	logs, sub, err := _MuleRegistry.contract.FilterLogs(opts, "MuleReported", senderHashRule, receiverHashRule)
	if err != nil {
		return nil, err
	}
	return &MuleRegistryMuleReportedIterator{contract: _MuleRegistry.contract, event: "MuleReported", logs: logs, sub: sub}, nil
}

// WatchMuleReported is a free log subscription operation binding the contract event 0x83836cc0d166876e68904bae129ac8f5aa73eb9a2ea197f1cf947e3871b10988.
//
// Solidity: event MuleReported(bytes32 indexed senderHash, bytes32 indexed receiverHash, uint256 amount, string reason, uint256 timestamp)
func (_MuleRegistry *MuleRegistryFilterer) WatchMuleReported(opts *bind.WatchOpts, sink chan<- *MuleRegistryMuleReported, senderHash [][32]byte, receiverHash [][32]byte) (event.Subscription, error) {

	var senderHashRule []interface{}
	for _, senderHashItem := range senderHash {
		senderHashRule = append(senderHashRule, senderHashItem)
	}
	var receiverHashRule []interface{}
	for _, receiverHashItem := range receiverHash {
		receiverHashRule = append(receiverHashRule, receiverHashItem)
	}

	logs, sub, err := _MuleRegistry.contract.WatchLogs(opts, "MuleReported", senderHashRule, receiverHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuleRegistryMuleReported)
				if err := _MuleRegistry.contract.UnpackLog(event, "MuleReported", log); err != nil {
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

// ParseMuleReported is a log parse operation binding the contract event 0x83836cc0d166876e68904bae129ac8f5aa73eb9a2ea197f1cf947e3871b10988.
//
// Solidity: event MuleReported(bytes32 indexed senderHash, bytes32 indexed receiverHash, uint256 amount, string reason, uint256 timestamp)
func (_MuleRegistry *MuleRegistryFilterer) ParseMuleReported(log types.Log) (*MuleRegistryMuleReported, error) {
	event := new(MuleRegistryMuleReported)
	if err := _MuleRegistry.contract.UnpackLog(event, "MuleReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
