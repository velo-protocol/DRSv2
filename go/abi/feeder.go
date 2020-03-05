// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vabi

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

// FeederABI is the input ABI used to generate the binding from.
const FeederABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fiatCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_collateralCode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curTimestamp\",\"type\":\"uint256\"}],\"name\":\"PriceSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valueTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"post\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeederBin is the compiled bytecode used for deploying new contracts.
var FeederBin = "0x608060405234801561001057600080fd5b506040516105283803806105288339818101604052606081101561003357600080fd5b50805160208201516040909201516003805460ff19166001179055600080546001600160a01b039093166001600160a01b0319909316929092179091556004919091556005556104a0806100886000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806364bc52041161007157806364bc52041461011e5780636d4ce63c146101265780638da5cb5b1461012e578063a3907d7114610152578063caa5863c1461015a578063dc9eaf3814610177576100a9565b806302fb0c5e146100ae5780632f2770db146100ca5780633fa4f245146100d4578063498b5828146100ee5780635587c43614610116575b600080fd5b6100b661017f565b604080519115158252519081900360200190f35b6100d2610188565b005b6100dc6101dd565b60408051918252519081900360200190f35b6100f66101e3565b604080519384526020840192909252151582820152519081900360600190f35b6100dc6101f5565b6100dc6101fb565b6100dc610201565b6101366102b7565b604080516001600160a01b039092168252519081900360200190f35b6100d26102c6565b6100d26004803603602081101561017057600080fd5b503561031e565b6100dc610400565b60035460ff1681565b6000546001600160a01b031633146101d15760405162461bcd60e51b815260040180806020018281038252603a815260200180610407603a913960400191505060405180910390fd5b6003805460ff19169055565b60015481565b60015460025460035460ff1615909192565b60055481565b60045481565b60008060015411610259576040805162461bcd60e51b815260206004820152601f60248201527f4665656465722e6765743a2076616c7565206e6f7420617661696c61626c6500604482015290519081900360640190fd5b60035460ff166102b0576040805162461bcd60e51b815260206004820152601f60248201527f4665656465722e6765743a20616374697665206d757374206265207472756500604482015290519081900360640190fd5b5060015490565b6000546001600160a01b031681565b6000546001600160a01b0316331461030f5760405162461bcd60e51b815260040180806020018281038252603a815260200180610407603a913960400191505060405180910390fd5b6003805460ff19166001179055565b6000546001600160a01b031633146103675760405162461bcd60e51b815260040180806020018281038252603a815260200180610407603a913960400191505060405180910390fd5b600154600254826103a95760405162461bcd60e51b815260040180806020018281038252602b815260200180610441602b913960400191505060405180910390fd5b600183905542600281905560408051848152602081018490528082018690526060810192909252517f409a3cd8da9fd5287dcd055b527b58b542d9e4efcb0c72732ed3f0046b77d8689181900360800190a1505050565b6002548156fe4665656465722e6f6e6c794f776e65723a2063616c6c6572206d75737420626520616e206f776e6572206f66207468697320636f6e74726163744665656465722e7365743a206e657756616c7565206d7573742062652067726561746572207468616e2030a265627a7a723158201bc2098842696384611ee0ec4e47db6403690eb5617f167f2c3ac49e8596f15c64736f6c63430005100032"

// DeployFeeder deploys a new Ethereum contract, binding an instance of Feeder to it.
func DeployFeeder(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _fiatCode [32]byte, _collateralCode [32]byte) (common.Address, *types.Transaction, *Feeder, error) {
	parsed, err := abi.JSON(strings.NewReader(FeederABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeederBin), backend, _owner, _fiatCode, _collateralCode)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Feeder{FeederCaller: FeederCaller{contract: contract}, FeederTransactor: FeederTransactor{contract: contract}, FeederFilterer: FeederFilterer{contract: contract}}, nil
}

// Feeder is an auto generated Go binding around an Ethereum contract.
type Feeder struct {
	FeederCaller     // Read-only binding to the contract
	FeederTransactor // Write-only binding to the contract
	FeederFilterer   // Log filterer for contract events
}

// FeederCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeederCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeederTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeederTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeederFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeederFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeederSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeederSession struct {
	Contract     *Feeder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeederCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeederCallerSession struct {
	Contract *FeederCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FeederTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeederTransactorSession struct {
	Contract     *FeederTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeederRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeederRaw struct {
	Contract *Feeder // Generic contract binding to access the raw methods on
}

// FeederCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeederCallerRaw struct {
	Contract *FeederCaller // Generic read-only contract binding to access the raw methods on
}

// FeederTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeederTransactorRaw struct {
	Contract *FeederTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeder creates a new instance of Feeder, bound to a specific deployed contract.
func NewFeeder(address common.Address, backend bind.ContractBackend) (*Feeder, error) {
	contract, err := bindFeeder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feeder{FeederCaller: FeederCaller{contract: contract}, FeederTransactor: FeederTransactor{contract: contract}, FeederFilterer: FeederFilterer{contract: contract}}, nil
}

// NewFeederCaller creates a new read-only instance of Feeder, bound to a specific deployed contract.
func NewFeederCaller(address common.Address, caller bind.ContractCaller) (*FeederCaller, error) {
	contract, err := bindFeeder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeederCaller{contract: contract}, nil
}

// NewFeederTransactor creates a new write-only instance of Feeder, bound to a specific deployed contract.
func NewFeederTransactor(address common.Address, transactor bind.ContractTransactor) (*FeederTransactor, error) {
	contract, err := bindFeeder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeederTransactor{contract: contract}, nil
}

// NewFeederFilterer creates a new log filterer instance of Feeder, bound to a specific deployed contract.
func NewFeederFilterer(address common.Address, filterer bind.ContractFilterer) (*FeederFilterer, error) {
	contract, err := bindFeeder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeederFilterer{contract: contract}, nil
}

// bindFeeder binds a generic wrapper to an already deployed contract.
func bindFeeder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeederABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feeder *FeederRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Feeder.Contract.FeederCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feeder *FeederRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feeder.Contract.FeederTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feeder *FeederRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feeder.Contract.FeederTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feeder *FeederCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Feeder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feeder *FeederTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feeder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feeder *FeederTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feeder.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Feeder *FeederCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Feeder *FeederSession) Active() (bool, error) {
	return _Feeder.Contract.Active(&_Feeder.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Feeder *FeederCallerSession) Active() (bool, error) {
	return _Feeder.Contract.Active(&_Feeder.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Feeder *FeederCaller) CollateralCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "collateralCode")
	return *ret0, err
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Feeder *FeederSession) CollateralCode() ([32]byte, error) {
	return _Feeder.Contract.CollateralCode(&_Feeder.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Feeder *FeederCallerSession) CollateralCode() ([32]byte, error) {
	return _Feeder.Contract.CollateralCode(&_Feeder.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Feeder *FeederCaller) FiatCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "fiatCode")
	return *ret0, err
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Feeder *FeederSession) FiatCode() ([32]byte, error) {
	return _Feeder.Contract.FiatCode(&_Feeder.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Feeder *FeederCallerSession) FiatCode() ([32]byte, error) {
	return _Feeder.Contract.FiatCode(&_Feeder.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Feeder *FeederCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Feeder *FeederSession) Get() (*big.Int, error) {
	return _Feeder.Contract.Get(&_Feeder.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Feeder *FeederCallerSession) Get() (*big.Int, error) {
	return _Feeder.Contract.Get(&_Feeder.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, uint256, bool)
func (_Feeder *FeederCaller) GetWithError(opts *bind.CallOpts) (*big.Int, *big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Feeder.contract.Call(opts, out, "getWithError")
	return *ret0, *ret1, *ret2, err
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, uint256, bool)
func (_Feeder *FeederSession) GetWithError() (*big.Int, *big.Int, bool, error) {
	return _Feeder.Contract.GetWithError(&_Feeder.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, uint256, bool)
func (_Feeder *FeederCallerSession) GetWithError() (*big.Int, *big.Int, bool, error) {
	return _Feeder.Contract.GetWithError(&_Feeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feeder *FeederCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feeder *FeederSession) Owner() (common.Address, error) {
	return _Feeder.Contract.Owner(&_Feeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feeder *FeederCallerSession) Owner() (common.Address, error) {
	return _Feeder.Contract.Owner(&_Feeder.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(uint256)
func (_Feeder *FeederCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "value")
	return *ret0, err
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(uint256)
func (_Feeder *FeederSession) Value() (*big.Int, error) {
	return _Feeder.Contract.Value(&_Feeder.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(uint256)
func (_Feeder *FeederCallerSession) Value() (*big.Int, error) {
	return _Feeder.Contract.Value(&_Feeder.CallOpts)
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() constant returns(uint256)
func (_Feeder *FeederCaller) ValueTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "valueTimestamp")
	return *ret0, err
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() constant returns(uint256)
func (_Feeder *FeederSession) ValueTimestamp() (*big.Int, error) {
	return _Feeder.Contract.ValueTimestamp(&_Feeder.CallOpts)
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() constant returns(uint256)
func (_Feeder *FeederCallerSession) ValueTimestamp() (*big.Int, error) {
	return _Feeder.Contract.ValueTimestamp(&_Feeder.CallOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_Feeder *FeederTransactor) Disable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "disable")
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_Feeder *FeederSession) Disable() (*types.Transaction, error) {
	return _Feeder.Contract.Disable(&_Feeder.TransactOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_Feeder *FeederTransactorSession) Disable() (*types.Transaction, error) {
	return _Feeder.Contract.Disable(&_Feeder.TransactOpts)
}

// Enable is a paid mutator transaction binding the contract method 0xa3907d71.
//
// Solidity: function enable() returns()
func (_Feeder *FeederTransactor) Enable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "enable")
}

// Enable is a paid mutator transaction binding the contract method 0xa3907d71.
//
// Solidity: function enable() returns()
func (_Feeder *FeederSession) Enable() (*types.Transaction, error) {
	return _Feeder.Contract.Enable(&_Feeder.TransactOpts)
}

// Enable is a paid mutator transaction binding the contract method 0xa3907d71.
//
// Solidity: function enable() returns()
func (_Feeder *FeederTransactorSession) Enable() (*types.Transaction, error) {
	return _Feeder.Contract.Enable(&_Feeder.TransactOpts)
}

// Post is a paid mutator transaction binding the contract method 0xcaa5863c.
//
// Solidity: function post(uint256 newValue) returns()
func (_Feeder *FeederTransactor) Post(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "post", newValue)
}

// Post is a paid mutator transaction binding the contract method 0xcaa5863c.
//
// Solidity: function post(uint256 newValue) returns()
func (_Feeder *FeederSession) Post(newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.Post(&_Feeder.TransactOpts, newValue)
}

// Post is a paid mutator transaction binding the contract method 0xcaa5863c.
//
// Solidity: function post(uint256 newValue) returns()
func (_Feeder *FeederTransactorSession) Post(newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.Post(&_Feeder.TransactOpts, newValue)
}

// FeederPriceSetIterator is returned from FilterPriceSet and is used to iterate over the raw logs and unpacked data for PriceSet events raised by the Feeder contract.
type FeederPriceSetIterator struct {
	Event *FeederPriceSet // Event containing the contract specifics and raw log

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
func (it *FeederPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederPriceSet)
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
		it.Event = new(FeederPriceSet)
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
func (it *FeederPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederPriceSet represents a PriceSet event raised by the Feeder contract.
type FeederPriceSet struct {
	PrevValue     *big.Int
	PrevTimestamp *big.Int
	CurPrice      *big.Int
	CurTimestamp  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceSet is a free log retrieval operation binding the contract event 0x409a3cd8da9fd5287dcd055b527b58b542d9e4efcb0c72732ed3f0046b77d868.
//
// Solidity: event PriceSet(uint256 prevValue, uint256 prevTimestamp, uint256 curPrice, uint256 curTimestamp)
func (_Feeder *FeederFilterer) FilterPriceSet(opts *bind.FilterOpts) (*FeederPriceSetIterator, error) {

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "PriceSet")
	if err != nil {
		return nil, err
	}
	return &FeederPriceSetIterator{contract: _Feeder.contract, event: "PriceSet", logs: logs, sub: sub}, nil
}

// WatchPriceSet is a free log subscription operation binding the contract event 0x409a3cd8da9fd5287dcd055b527b58b542d9e4efcb0c72732ed3f0046b77d868.
//
// Solidity: event PriceSet(uint256 prevValue, uint256 prevTimestamp, uint256 curPrice, uint256 curTimestamp)
func (_Feeder *FeederFilterer) WatchPriceSet(opts *bind.WatchOpts, sink chan<- *FeederPriceSet) (event.Subscription, error) {

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "PriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederPriceSet)
				if err := _Feeder.contract.UnpackLog(event, "PriceSet", log); err != nil {
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

// ParsePriceSet is a log parse operation binding the contract event 0x409a3cd8da9fd5287dcd055b527b58b542d9e4efcb0c72732ed3f0046b77d868.
//
// Solidity: event PriceSet(uint256 prevValue, uint256 prevTimestamp, uint256 curPrice, uint256 curTimestamp)
func (_Feeder *FeederFilterer) ParsePriceSet(log types.Log) (*FeederPriceSet, error) {
	event := new(FeederPriceSet)
	if err := _Feeder.contract.UnpackLog(event, "PriceSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
