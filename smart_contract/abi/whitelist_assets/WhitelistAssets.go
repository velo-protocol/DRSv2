// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
const RolesBin = `0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582095f5f3cbbab6335e70c50e317d17e5cbe28700ec66aab4522a2b01c50b8b481d0029`

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// WhitelistAdminRoleABI is the input ABI used to generate the binding from.
const WhitelistAdminRoleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"}]"

// WhitelistAdminRoleBin is the compiled bytecode used for deploying new contracts.
const WhitelistAdminRoleBin = `0x`

// DeployWhitelistAdminRole deploys a new Ethereum contract, binding an instance of WhitelistAdminRole to it.
func DeployWhitelistAdminRole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitelistAdminRole, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAdminRoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistAdminRoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistAdminRole{WhitelistAdminRoleCaller: WhitelistAdminRoleCaller{contract: contract}, WhitelistAdminRoleTransactor: WhitelistAdminRoleTransactor{contract: contract}, WhitelistAdminRoleFilterer: WhitelistAdminRoleFilterer{contract: contract}}, nil
}

// WhitelistAdminRole is an auto generated Go binding around an Ethereum contract.
type WhitelistAdminRole struct {
	WhitelistAdminRoleCaller     // Read-only binding to the contract
	WhitelistAdminRoleTransactor // Write-only binding to the contract
	WhitelistAdminRoleFilterer   // Log filterer for contract events
}

// WhitelistAdminRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistAdminRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistAdminRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistAdminRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistAdminRoleSession struct {
	Contract     *WhitelistAdminRole // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WhitelistAdminRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistAdminRoleCallerSession struct {
	Contract *WhitelistAdminRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WhitelistAdminRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistAdminRoleTransactorSession struct {
	Contract     *WhitelistAdminRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WhitelistAdminRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistAdminRoleRaw struct {
	Contract *WhitelistAdminRole // Generic contract binding to access the raw methods on
}

// WhitelistAdminRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistAdminRoleCallerRaw struct {
	Contract *WhitelistAdminRoleCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistAdminRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistAdminRoleTransactorRaw struct {
	Contract *WhitelistAdminRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistAdminRole creates a new instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRole(address common.Address, backend bind.ContractBackend) (*WhitelistAdminRole, error) {
	contract, err := bindWhitelistAdminRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRole{WhitelistAdminRoleCaller: WhitelistAdminRoleCaller{contract: contract}, WhitelistAdminRoleTransactor: WhitelistAdminRoleTransactor{contract: contract}, WhitelistAdminRoleFilterer: WhitelistAdminRoleFilterer{contract: contract}}, nil
}

// NewWhitelistAdminRoleCaller creates a new read-only instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleCaller(address common.Address, caller bind.ContractCaller) (*WhitelistAdminRoleCaller, error) {
	contract, err := bindWhitelistAdminRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleCaller{contract: contract}, nil
}

// NewWhitelistAdminRoleTransactor creates a new write-only instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistAdminRoleTransactor, error) {
	contract, err := bindWhitelistAdminRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleTransactor{contract: contract}, nil
}

// NewWhitelistAdminRoleFilterer creates a new log filterer instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistAdminRoleFilterer, error) {
	contract, err := bindWhitelistAdminRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleFilterer{contract: contract}, nil
}

// bindWhitelistAdminRole binds a generic wrapper to an already deployed contract.
func bindWhitelistAdminRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAdminRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAdminRole *WhitelistAdminRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAdminRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistAdminRole.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAdminRole.Contract.IsWhitelistAdmin(&_WhitelistAdminRole.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAdminRole.Contract.IsWhitelistAdmin(&_WhitelistAdminRole.CallOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.AddWhitelistAdmin(&_WhitelistAdminRole.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.AddWhitelistAdmin(&_WhitelistAdminRole.TransactOpts, account)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.RenounceWhitelistAdmin(&_WhitelistAdminRole.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.RenounceWhitelistAdmin(&_WhitelistAdminRole.TransactOpts)
}

// WhitelistAdminRoleWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminAddedIterator struct {
	Event *WhitelistAdminRoleWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAdminRoleWhitelistAdminAdded)
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
		it.Event = new(WhitelistAdminRoleWhitelistAdminAdded)
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
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAdminRoleWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*WhitelistAdminRoleWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleWhitelistAdminAddedIterator{contract: _WhitelistAdminRole.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *WhitelistAdminRoleWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAdminRoleWhitelistAdminAdded)
				if err := _WhitelistAdminRole.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// WhitelistAdminRoleWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminRemovedIterator struct {
	Event *WhitelistAdminRoleWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAdminRoleWhitelistAdminRemoved)
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
		it.Event = new(WhitelistAdminRoleWhitelistAdminRemoved)
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
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAdminRoleWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*WhitelistAdminRoleWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleWhitelistAdminRemovedIterator{contract: _WhitelistAdminRole.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistAdminRoleWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAdminRoleWhitelistAdminRemoved)
				if err := _WhitelistAdminRole.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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

// WhitelistAssetsABI is the input ABI used to generate the binding from.
const WhitelistAssetsABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetName\",\"type\":\"bytes32\"}],\"name\":\"removeAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetName\",\"type\":\"bytes32\"}],\"name\":\"creditAddr\",\"outputs\":[{\"name\":\"assetAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetName\",\"type\":\"bytes32\"},{\"name\":\"assetContract\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"assetName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"assetAddr\",\"type\":\"address\"}],\"name\":\"AddedAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"assetName\",\"type\":\"bytes32\"}],\"name\":\"RemovedAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"}]"

// WhitelistAssetsBin is the compiled bytecode used for deploying new contracts.
const WhitelistAssetsBin = `0x60806040526100133361001860201b60201c565b61018b565b61003081600061006760201b61052b1790919060201c565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b610077828261010860201b60201c565b156100e357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216151561016b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806108206022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6106868061019a6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639fda5b661161005b5780639fda5b66146100cf578063bb5f747b14610108578063cc4f488914610142578063ce13f0fb1461015f5761007d565b80634c5a628c146100825780637362d9c81461008c578063837592db146100b2575b600080fd5b61008a61018b565b005b61008a600480360360208110156100a257600080fd5b50356001600160a01b0316610196565b61008a600480360360208110156100c857600080fd5b50356101eb565b6100ec600480360360208110156100e557600080fd5b503561027a565b604080516001600160a01b039092168252519081900360200190f35b61012e6004803603602081101561011e57600080fd5b50356001600160a01b0316610295565b604080519115158252519081900360200190f35b6100ec6004803603602081101561015857600080fd5b50356102ad565b61008a6004803603604081101561017557600080fd5b50803590602001356001600160a01b0316610321565b610194336103c3565b565b61019f33610295565b15156101df57604051600160e51b62461bcd02815260040180806020018281038252604081526020018061061b6040913960400191505060405180910390fd5b6101e88161040b565b50565b6101f433610295565b151561023457604051600160e51b62461bcd02815260040180806020018281038252604081526020018061061b6040913960400191505060405180910390fd5b60008181526001602052604080822080546001600160a01b03191690555182917f049af0f3c7a43e6327801380e09c382492f82b053cac998a237033b85897a0c991a250565b6001602052600090815260409020546001600160a01b031681565b60006102a7818363ffffffff61045316565b92915050565b6000818152600160205260408120546001600160a01b0316151561030557604051600160e51b62461bcd0281526004018080602001828103825260288152602001806105b06028913960400191505060405180910390fd5b506000908152600160205260409020546001600160a01b031690565b61032a33610295565b151561036a57604051600160e51b62461bcd02815260040180806020018281038252604081526020018061061b6040913960400191505060405180910390fd5b60008281526001602052604080822080546001600160a01b0319166001600160a01b0385169081179091559051909184917f5492b9bf82f630e2eb238760ff6780a294716174735cf447450f3d91074d306b9190a35050565b6103d460008263ffffffff6104bf16565b6040516001600160a01b038216907f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16590600090a250565b61041c60008263ffffffff61052b16565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b60006001600160a01b038216151561049f57604051600160e51b62461bcd0281526004018080602001828103825260228152602001806105f96022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6104c98282610453565b151561050957604051600160e51b62461bcd0281526004018080602001828103825260218152602001806105d86021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b6105358282610453565b1561058a5760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff1916600117905556fe57686974656c6973744173736574733a207468652061737365742069732070726f68696269746564526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65526f6c65733a206163636f756e7420697320746865207a65726f206164647265737357686974656c69737441646d696e526f6c653a2063616c6c657220646f6573206e6f742068617665207468652057686974656c69737441646d696e20726f6c65a165627a7a72305820e230696fbb3bb2fdc787c6a4ededd791fbf82c451ed4d3b24466831ae471fe350029526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373`

// DeployWhitelistAssets deploys a new Ethereum contract, binding an instance of WhitelistAssets to it.
func DeployWhitelistAssets(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitelistAssets, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAssetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistAssetsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistAssets{WhitelistAssetsCaller: WhitelistAssetsCaller{contract: contract}, WhitelistAssetsTransactor: WhitelistAssetsTransactor{contract: contract}, WhitelistAssetsFilterer: WhitelistAssetsFilterer{contract: contract}}, nil
}

// WhitelistAssets is an auto generated Go binding around an Ethereum contract.
type WhitelistAssets struct {
	WhitelistAssetsCaller     // Read-only binding to the contract
	WhitelistAssetsTransactor // Write-only binding to the contract
	WhitelistAssetsFilterer   // Log filterer for contract events
}

// WhitelistAssetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistAssetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAssetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistAssetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAssetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistAssetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAssetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistAssetsSession struct {
	Contract     *WhitelistAssets  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistAssetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistAssetsCallerSession struct {
	Contract *WhitelistAssetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WhitelistAssetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistAssetsTransactorSession struct {
	Contract     *WhitelistAssetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WhitelistAssetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistAssetsRaw struct {
	Contract *WhitelistAssets // Generic contract binding to access the raw methods on
}

// WhitelistAssetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistAssetsCallerRaw struct {
	Contract *WhitelistAssetsCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistAssetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistAssetsTransactorRaw struct {
	Contract *WhitelistAssetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistAssets creates a new instance of WhitelistAssets, bound to a specific deployed contract.
func NewWhitelistAssets(address common.Address, backend bind.ContractBackend) (*WhitelistAssets, error) {
	contract, err := bindWhitelistAssets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssets{WhitelistAssetsCaller: WhitelistAssetsCaller{contract: contract}, WhitelistAssetsTransactor: WhitelistAssetsTransactor{contract: contract}, WhitelistAssetsFilterer: WhitelistAssetsFilterer{contract: contract}}, nil
}

// NewWhitelistAssetsCaller creates a new read-only instance of WhitelistAssets, bound to a specific deployed contract.
func NewWhitelistAssetsCaller(address common.Address, caller bind.ContractCaller) (*WhitelistAssetsCaller, error) {
	contract, err := bindWhitelistAssets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsCaller{contract: contract}, nil
}

// NewWhitelistAssetsTransactor creates a new write-only instance of WhitelistAssets, bound to a specific deployed contract.
func NewWhitelistAssetsTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistAssetsTransactor, error) {
	contract, err := bindWhitelistAssets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsTransactor{contract: contract}, nil
}

// NewWhitelistAssetsFilterer creates a new log filterer instance of WhitelistAssets, bound to a specific deployed contract.
func NewWhitelistAssetsFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistAssetsFilterer, error) {
	contract, err := bindWhitelistAssets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsFilterer{contract: contract}, nil
}

// bindWhitelistAssets binds a generic wrapper to an already deployed contract.
func bindWhitelistAssets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAssetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAssets *WhitelistAssetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAssets.Contract.WhitelistAssetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAssets *WhitelistAssetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.WhitelistAssetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAssets *WhitelistAssetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.WhitelistAssetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAssets *WhitelistAssetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAssets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAssets *WhitelistAssetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAssets *WhitelistAssetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) constant returns(address)
func (_WhitelistAssets *WhitelistAssetsCaller) Assets(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistAssets.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) constant returns(address)
func (_WhitelistAssets *WhitelistAssetsSession) Assets(arg0 [32]byte) (common.Address, error) {
	return _WhitelistAssets.Contract.Assets(&_WhitelistAssets.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) constant returns(address)
func (_WhitelistAssets *WhitelistAssetsCallerSession) Assets(arg0 [32]byte) (common.Address, error) {
	return _WhitelistAssets.Contract.Assets(&_WhitelistAssets.CallOpts, arg0)
}

// CreditAddr is a free data retrieval call binding the contract method 0xcc4f4889.
//
// Solidity: function creditAddr(bytes32 assetName) constant returns(address assetAddr)
func (_WhitelistAssets *WhitelistAssetsCaller) CreditAddr(opts *bind.CallOpts, assetName [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistAssets.contract.Call(opts, out, "creditAddr", assetName)
	return *ret0, err
}

// CreditAddr is a free data retrieval call binding the contract method 0xcc4f4889.
//
// Solidity: function creditAddr(bytes32 assetName) constant returns(address assetAddr)
func (_WhitelistAssets *WhitelistAssetsSession) CreditAddr(assetName [32]byte) (common.Address, error) {
	return _WhitelistAssets.Contract.CreditAddr(&_WhitelistAssets.CallOpts, assetName)
}

// CreditAddr is a free data retrieval call binding the contract method 0xcc4f4889.
//
// Solidity: function creditAddr(bytes32 assetName) constant returns(address assetAddr)
func (_WhitelistAssets *WhitelistAssetsCallerSession) CreditAddr(assetName [32]byte) (common.Address, error) {
	return _WhitelistAssets.Contract.CreditAddr(&_WhitelistAssets.CallOpts, assetName)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAssets *WhitelistAssetsCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistAssets.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAssets *WhitelistAssetsSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAssets.Contract.IsWhitelistAdmin(&_WhitelistAssets.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAssets *WhitelistAssetsCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAssets.Contract.IsWhitelistAdmin(&_WhitelistAssets.CallOpts, account)
}

// AddAsset is a paid mutator transaction binding the contract method 0xce13f0fb.
//
// Solidity: function addAsset(bytes32 assetName, address assetContract) returns()
func (_WhitelistAssets *WhitelistAssetsTransactor) AddAsset(opts *bind.TransactOpts, assetName [32]byte, assetContract common.Address) (*types.Transaction, error) {
	return _WhitelistAssets.contract.Transact(opts, "addAsset", assetName, assetContract)
}

// AddAsset is a paid mutator transaction binding the contract method 0xce13f0fb.
//
// Solidity: function addAsset(bytes32 assetName, address assetContract) returns()
func (_WhitelistAssets *WhitelistAssetsSession) AddAsset(assetName [32]byte, assetContract common.Address) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.AddAsset(&_WhitelistAssets.TransactOpts, assetName, assetContract)
}

// AddAsset is a paid mutator transaction binding the contract method 0xce13f0fb.
//
// Solidity: function addAsset(bytes32 assetName, address assetContract) returns()
func (_WhitelistAssets *WhitelistAssetsTransactorSession) AddAsset(assetName [32]byte, assetContract common.Address) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.AddAsset(&_WhitelistAssets.TransactOpts, assetName, assetContract)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAssets *WhitelistAssetsTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistAssets.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAssets *WhitelistAssetsSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.AddWhitelistAdmin(&_WhitelistAssets.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAssets *WhitelistAssetsTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.AddWhitelistAdmin(&_WhitelistAssets.TransactOpts, account)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x837592db.
//
// Solidity: function removeAsset(bytes32 assetName) returns()
func (_WhitelistAssets *WhitelistAssetsTransactor) RemoveAsset(opts *bind.TransactOpts, assetName [32]byte) (*types.Transaction, error) {
	return _WhitelistAssets.contract.Transact(opts, "removeAsset", assetName)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x837592db.
//
// Solidity: function removeAsset(bytes32 assetName) returns()
func (_WhitelistAssets *WhitelistAssetsSession) RemoveAsset(assetName [32]byte) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.RemoveAsset(&_WhitelistAssets.TransactOpts, assetName)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x837592db.
//
// Solidity: function removeAsset(bytes32 assetName) returns()
func (_WhitelistAssets *WhitelistAssetsTransactorSession) RemoveAsset(assetName [32]byte) (*types.Transaction, error) {
	return _WhitelistAssets.Contract.RemoveAsset(&_WhitelistAssets.TransactOpts, assetName)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAssets *WhitelistAssetsTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAssets.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAssets *WhitelistAssetsSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAssets.Contract.RenounceWhitelistAdmin(&_WhitelistAssets.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAssets *WhitelistAssetsTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAssets.Contract.RenounceWhitelistAdmin(&_WhitelistAssets.TransactOpts)
}

// WhitelistAssetsAddedAssetIterator is returned from FilterAddedAsset and is used to iterate over the raw logs and unpacked data for AddedAsset events raised by the WhitelistAssets contract.
type WhitelistAssetsAddedAssetIterator struct {
	Event *WhitelistAssetsAddedAsset // Event containing the contract specifics and raw log

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
func (it *WhitelistAssetsAddedAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAssetsAddedAsset)
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
		it.Event = new(WhitelistAssetsAddedAsset)
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
func (it *WhitelistAssetsAddedAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAssetsAddedAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAssetsAddedAsset represents a AddedAsset event raised by the WhitelistAssets contract.
type WhitelistAssetsAddedAsset struct {
	AssetName [32]byte
	AssetAddr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedAsset is a free log retrieval operation binding the contract event 0x5492b9bf82f630e2eb238760ff6780a294716174735cf447450f3d91074d306b.
//
// Solidity: event AddedAsset(bytes32 indexed assetName, address indexed assetAddr)
func (_WhitelistAssets *WhitelistAssetsFilterer) FilterAddedAsset(opts *bind.FilterOpts, assetName [][32]byte, assetAddr []common.Address) (*WhitelistAssetsAddedAssetIterator, error) {

	var assetNameRule []interface{}
	for _, assetNameItem := range assetName {
		assetNameRule = append(assetNameRule, assetNameItem)
	}
	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}

	logs, sub, err := _WhitelistAssets.contract.FilterLogs(opts, "AddedAsset", assetNameRule, assetAddrRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsAddedAssetIterator{contract: _WhitelistAssets.contract, event: "AddedAsset", logs: logs, sub: sub}, nil
}

// WatchAddedAsset is a free log subscription operation binding the contract event 0x5492b9bf82f630e2eb238760ff6780a294716174735cf447450f3d91074d306b.
//
// Solidity: event AddedAsset(bytes32 indexed assetName, address indexed assetAddr)
func (_WhitelistAssets *WhitelistAssetsFilterer) WatchAddedAsset(opts *bind.WatchOpts, sink chan<- *WhitelistAssetsAddedAsset, assetName [][32]byte, assetAddr []common.Address) (event.Subscription, error) {

	var assetNameRule []interface{}
	for _, assetNameItem := range assetName {
		assetNameRule = append(assetNameRule, assetNameItem)
	}
	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}

	logs, sub, err := _WhitelistAssets.contract.WatchLogs(opts, "AddedAsset", assetNameRule, assetAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAssetsAddedAsset)
				if err := _WhitelistAssets.contract.UnpackLog(event, "AddedAsset", log); err != nil {
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

// WhitelistAssetsRemovedAssetIterator is returned from FilterRemovedAsset and is used to iterate over the raw logs and unpacked data for RemovedAsset events raised by the WhitelistAssets contract.
type WhitelistAssetsRemovedAssetIterator struct {
	Event *WhitelistAssetsRemovedAsset // Event containing the contract specifics and raw log

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
func (it *WhitelistAssetsRemovedAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAssetsRemovedAsset)
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
		it.Event = new(WhitelistAssetsRemovedAsset)
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
func (it *WhitelistAssetsRemovedAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAssetsRemovedAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAssetsRemovedAsset represents a RemovedAsset event raised by the WhitelistAssets contract.
type WhitelistAssetsRemovedAsset struct {
	AssetName [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedAsset is a free log retrieval operation binding the contract event 0x049af0f3c7a43e6327801380e09c382492f82b053cac998a237033b85897a0c9.
//
// Solidity: event RemovedAsset(bytes32 indexed assetName)
func (_WhitelistAssets *WhitelistAssetsFilterer) FilterRemovedAsset(opts *bind.FilterOpts, assetName [][32]byte) (*WhitelistAssetsRemovedAssetIterator, error) {

	var assetNameRule []interface{}
	for _, assetNameItem := range assetName {
		assetNameRule = append(assetNameRule, assetNameItem)
	}

	logs, sub, err := _WhitelistAssets.contract.FilterLogs(opts, "RemovedAsset", assetNameRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsRemovedAssetIterator{contract: _WhitelistAssets.contract, event: "RemovedAsset", logs: logs, sub: sub}, nil
}

// WatchRemovedAsset is a free log subscription operation binding the contract event 0x049af0f3c7a43e6327801380e09c382492f82b053cac998a237033b85897a0c9.
//
// Solidity: event RemovedAsset(bytes32 indexed assetName)
func (_WhitelistAssets *WhitelistAssetsFilterer) WatchRemovedAsset(opts *bind.WatchOpts, sink chan<- *WhitelistAssetsRemovedAsset, assetName [][32]byte) (event.Subscription, error) {

	var assetNameRule []interface{}
	for _, assetNameItem := range assetName {
		assetNameRule = append(assetNameRule, assetNameItem)
	}

	logs, sub, err := _WhitelistAssets.contract.WatchLogs(opts, "RemovedAsset", assetNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAssetsRemovedAsset)
				if err := _WhitelistAssets.contract.UnpackLog(event, "RemovedAsset", log); err != nil {
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

// WhitelistAssetsWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the WhitelistAssets contract.
type WhitelistAssetsWhitelistAdminAddedIterator struct {
	Event *WhitelistAssetsWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *WhitelistAssetsWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAssetsWhitelistAdminAdded)
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
		it.Event = new(WhitelistAssetsWhitelistAdminAdded)
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
func (it *WhitelistAssetsWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAssetsWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAssetsWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the WhitelistAssets contract.
type WhitelistAssetsWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_WhitelistAssets *WhitelistAssetsFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*WhitelistAssetsWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAssets.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsWhitelistAdminAddedIterator{contract: _WhitelistAssets.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_WhitelistAssets *WhitelistAssetsFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *WhitelistAssetsWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAssets.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAssetsWhitelistAdminAdded)
				if err := _WhitelistAssets.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// WhitelistAssetsWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the WhitelistAssets contract.
type WhitelistAssetsWhitelistAdminRemovedIterator struct {
	Event *WhitelistAssetsWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistAssetsWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAssetsWhitelistAdminRemoved)
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
		it.Event = new(WhitelistAssetsWhitelistAdminRemoved)
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
func (it *WhitelistAssetsWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAssetsWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAssetsWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the WhitelistAssets contract.
type WhitelistAssetsWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_WhitelistAssets *WhitelistAssetsFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*WhitelistAssetsWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAssets.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAssetsWhitelistAdminRemovedIterator{contract: _WhitelistAssets.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_WhitelistAssets *WhitelistAssetsFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistAssetsWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAssets.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAssetsWhitelistAdminRemoved)
				if err := _WhitelistAssets.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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
