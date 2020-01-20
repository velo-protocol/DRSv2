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

// DigitalReserveSystemABI is the input ABI used to generate the binding from.
const DigitalReserveSystemABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"heartAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableCreditAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAssetAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peggedCurrency\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"peggedValue\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"Setup\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"heart\",\"outputs\":[{\"internalType\":\"contractIHeart\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peggedCurrency\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"peggedValue\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"netCollateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"}],\"name\":\"mintFromCollateralAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"}],\"name\":\"mintFromStableCreditAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stableCreditAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"}],\"name\":\"rebalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"}],\"name\":\"getExchange\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetCode\",\"type\":\"string\"}],\"name\":\"collateralOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DigitalReserveSystem is an auto generated Go binding around an Ethereum contract.
type DigitalReserveSystem struct {
	DigitalReserveSystemCaller     // Read-only binding to the contract
	DigitalReserveSystemTransactor // Write-only binding to the contract
	DigitalReserveSystemFilterer   // Log filterer for contract events
}

// DigitalReserveSystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type DigitalReserveSystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalReserveSystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DigitalReserveSystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalReserveSystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DigitalReserveSystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalReserveSystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DigitalReserveSystemSession struct {
	Contract     *DigitalReserveSystem // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DigitalReserveSystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DigitalReserveSystemCallerSession struct {
	Contract *DigitalReserveSystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DigitalReserveSystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DigitalReserveSystemTransactorSession struct {
	Contract     *DigitalReserveSystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DigitalReserveSystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type DigitalReserveSystemRaw struct {
	Contract *DigitalReserveSystem // Generic contract binding to access the raw methods on
}

// DigitalReserveSystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DigitalReserveSystemCallerRaw struct {
	Contract *DigitalReserveSystemCaller // Generic read-only contract binding to access the raw methods on
}

// DigitalReserveSystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DigitalReserveSystemTransactorRaw struct {
	Contract *DigitalReserveSystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDigitalReserveSystem creates a new instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystem(address common.Address, backend bind.ContractBackend) (*DigitalReserveSystem, error) {
	contract, err := bindDigitalReserveSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystem{DigitalReserveSystemCaller: DigitalReserveSystemCaller{contract: contract}, DigitalReserveSystemTransactor: DigitalReserveSystemTransactor{contract: contract}, DigitalReserveSystemFilterer: DigitalReserveSystemFilterer{contract: contract}}, nil
}

// NewDigitalReserveSystemCaller creates a new read-only instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystemCaller(address common.Address, caller bind.ContractCaller) (*DigitalReserveSystemCaller, error) {
	contract, err := bindDigitalReserveSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemCaller{contract: contract}, nil
}

// NewDigitalReserveSystemTransactor creates a new write-only instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*DigitalReserveSystemTransactor, error) {
	contract, err := bindDigitalReserveSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemTransactor{contract: contract}, nil
}

// NewDigitalReserveSystemFilterer creates a new log filterer instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*DigitalReserveSystemFilterer, error) {
	contract, err := bindDigitalReserveSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemFilterer{contract: contract}, nil
}

// bindDigitalReserveSystem binds a generic wrapper to an already deployed contract.
func bindDigitalReserveSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DigitalReserveSystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigitalReserveSystem *DigitalReserveSystemRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DigitalReserveSystem.Contract.DigitalReserveSystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigitalReserveSystem *DigitalReserveSystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.DigitalReserveSystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigitalReserveSystem *DigitalReserveSystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.DigitalReserveSystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigitalReserveSystem *DigitalReserveSystemCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DigitalReserveSystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigitalReserveSystem *DigitalReserveSystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigitalReserveSystem *DigitalReserveSystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.contract.Transact(opts, method, params...)
}

// CollateralOf is a free data retrieval call binding the contract method 0x47d5ff3d.
//
// Solidity: function collateralOf(address creditOwner, string assetCode) constant returns(uint256, address)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) CollateralOf(opts *bind.CallOpts, creditOwner common.Address, assetCode string) (*big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DigitalReserveSystem.contract.Call(opts, out, "collateralOf", creditOwner, assetCode)
	return *ret0, *ret1, err
}

// CollateralOf is a free data retrieval call binding the contract method 0x47d5ff3d.
//
// Solidity: function collateralOf(address creditOwner, string assetCode) constant returns(uint256, address)
func (_DigitalReserveSystem *DigitalReserveSystemSession) CollateralOf(creditOwner common.Address, assetCode string) (*big.Int, common.Address, error) {
	return _DigitalReserveSystem.Contract.CollateralOf(&_DigitalReserveSystem.CallOpts, creditOwner, assetCode)
}

// CollateralOf is a free data retrieval call binding the contract method 0x47d5ff3d.
//
// Solidity: function collateralOf(address creditOwner, string assetCode) constant returns(uint256, address)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) CollateralOf(creditOwner common.Address, assetCode string) (*big.Int, common.Address, error) {
	return _DigitalReserveSystem.Contract.CollateralOf(&_DigitalReserveSystem.CallOpts, creditOwner, assetCode)
}

// GetExchange is a free data retrieval call binding the contract method 0xff4b2032.
//
// Solidity: function getExchange(string assetCode) constant returns(string, bytes32, uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) GetExchange(opts *bind.CallOpts, assetCode string) (string, [32]byte, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _DigitalReserveSystem.contract.Call(opts, out, "getExchange", assetCode)
	return *ret0, *ret1, *ret2, err
}

// GetExchange is a free data retrieval call binding the contract method 0xff4b2032.
//
// Solidity: function getExchange(string assetCode) constant returns(string, bytes32, uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) GetExchange(assetCode string) (string, [32]byte, *big.Int, error) {
	return _DigitalReserveSystem.Contract.GetExchange(&_DigitalReserveSystem.CallOpts, assetCode)
}

// GetExchange is a free data retrieval call binding the contract method 0xff4b2032.
//
// Solidity: function getExchange(string assetCode) constant returns(string, bytes32, uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) GetExchange(assetCode string) (string, [32]byte, *big.Int, error) {
	return _DigitalReserveSystem.Contract.GetExchange(&_DigitalReserveSystem.CallOpts, assetCode)
}

// Heart is a free data retrieval call binding the contract method 0xf58d1c94.
//
// Solidity: function heart() constant returns(address)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) Heart(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "heart")
	return *ret0, err
}

// Heart is a free data retrieval call binding the contract method 0xf58d1c94.
//
// Solidity: function heart() constant returns(address)
func (_DigitalReserveSystem *DigitalReserveSystemSession) Heart() (common.Address, error) {
	return _DigitalReserveSystem.Contract.Heart(&_DigitalReserveSystem.CallOpts)
}

// Heart is a free data retrieval call binding the contract method 0xf58d1c94.
//
// Solidity: function heart() constant returns(address)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) Heart() (common.Address, error) {
	return _DigitalReserveSystem.Contract.Heart(&_DigitalReserveSystem.CallOpts)
}

// MintFromCollateralAmount is a paid mutator transaction binding the contract method 0xf15821f7.
//
// Solidity: function mintFromCollateralAmount(uint256 netCollateralAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) MintFromCollateralAmount(opts *bind.TransactOpts, netCollateralAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "mintFromCollateralAmount", netCollateralAmount, assetCode)
}

// MintFromCollateralAmount is a paid mutator transaction binding the contract method 0xf15821f7.
//
// Solidity: function mintFromCollateralAmount(uint256 netCollateralAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemSession) MintFromCollateralAmount(netCollateralAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.MintFromCollateralAmount(&_DigitalReserveSystem.TransactOpts, netCollateralAmount, assetCode)
}

// MintFromCollateralAmount is a paid mutator transaction binding the contract method 0xf15821f7.
//
// Solidity: function mintFromCollateralAmount(uint256 netCollateralAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) MintFromCollateralAmount(netCollateralAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.MintFromCollateralAmount(&_DigitalReserveSystem.TransactOpts, netCollateralAmount, assetCode)
}

// MintFromStableCreditAmount is a paid mutator transaction binding the contract method 0x99bbfd72.
//
// Solidity: function mintFromStableCreditAmount(uint256 mintAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) MintFromStableCreditAmount(opts *bind.TransactOpts, mintAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "mintFromStableCreditAmount", mintAmount, assetCode)
}

// MintFromStableCreditAmount is a paid mutator transaction binding the contract method 0x99bbfd72.
//
// Solidity: function mintFromStableCreditAmount(uint256 mintAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemSession) MintFromStableCreditAmount(mintAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.MintFromStableCreditAmount(&_DigitalReserveSystem.TransactOpts, mintAmount, assetCode)
}

// MintFromStableCreditAmount is a paid mutator transaction binding the contract method 0x99bbfd72.
//
// Solidity: function mintFromStableCreditAmount(uint256 mintAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) MintFromStableCreditAmount(mintAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.MintFromStableCreditAmount(&_DigitalReserveSystem.TransactOpts, mintAmount, assetCode)
}

// Rebalance is a paid mutator transaction binding the contract method 0x41212e9e.
//
// Solidity: function rebalance(string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) Rebalance(opts *bind.TransactOpts, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "rebalance", assetCode)
}

// Rebalance is a paid mutator transaction binding the contract method 0x41212e9e.
//
// Solidity: function rebalance(string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemSession) Rebalance(assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Rebalance(&_DigitalReserveSystem.TransactOpts, assetCode)
}

// Rebalance is a paid mutator transaction binding the contract method 0x41212e9e.
//
// Solidity: function rebalance(string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) Rebalance(assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Rebalance(&_DigitalReserveSystem.TransactOpts, assetCode)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 stableCreditAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) Redeem(opts *bind.TransactOpts, stableCreditAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "redeem", stableCreditAmount, assetCode)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 stableCreditAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemSession) Redeem(stableCreditAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Redeem(&_DigitalReserveSystem.TransactOpts, stableCreditAmount, assetCode)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 stableCreditAmount, string assetCode) returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) Redeem(stableCreditAmount *big.Int, assetCode string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Redeem(&_DigitalReserveSystem.TransactOpts, stableCreditAmount, assetCode)
}

// Setup is a paid mutator transaction binding the contract method 0xcbbaa9a0.
//
// Solidity: function setup(bytes32 collateralAssetCode, bytes32 peggedCurrency, string assetCode, uint256 peggedValue) returns(string, address)
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) Setup(opts *bind.TransactOpts, collateralAssetCode [32]byte, peggedCurrency [32]byte, assetCode string, peggedValue *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "setup", collateralAssetCode, peggedCurrency, assetCode, peggedValue)
}

// Setup is a paid mutator transaction binding the contract method 0xcbbaa9a0.
//
// Solidity: function setup(bytes32 collateralAssetCode, bytes32 peggedCurrency, string assetCode, uint256 peggedValue) returns(string, address)
func (_DigitalReserveSystem *DigitalReserveSystemSession) Setup(collateralAssetCode [32]byte, peggedCurrency [32]byte, assetCode string, peggedValue *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Setup(&_DigitalReserveSystem.TransactOpts, collateralAssetCode, peggedCurrency, assetCode, peggedValue)
}

// Setup is a paid mutator transaction binding the contract method 0xcbbaa9a0.
//
// Solidity: function setup(bytes32 collateralAssetCode, bytes32 peggedCurrency, string assetCode, uint256 peggedValue) returns(string, address)
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) Setup(collateralAssetCode [32]byte, peggedCurrency [32]byte, assetCode string, peggedValue *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Setup(&_DigitalReserveSystem.TransactOpts, collateralAssetCode, peggedCurrency, assetCode, peggedValue)
}

// DigitalReserveSystemMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DigitalReserveSystem contract.
type DigitalReserveSystemMintIterator struct {
	Event *DigitalReserveSystemMint // Event containing the contract specifics and raw log

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
func (it *DigitalReserveSystemMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalReserveSystemMint)
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
		it.Event = new(DigitalReserveSystemMint)
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
func (it *DigitalReserveSystemMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalReserveSystemMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalReserveSystemMint represents a Mint event raised by the DigitalReserveSystem contract.
type DigitalReserveSystemMint struct {
	AssetCode           string
	MintAmount          *big.Int
	AssetAddress        common.Address
	CollateralAssetCode [32]byte
	CollateralAmount    *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xddd98b18c74c43aa2894d989d78dde03a92ddf9e04a53b5655fdfdb73ddd06dc.
//
// Solidity: event Mint(string assetCode, uint256 mintAmount, address indexed assetAddress, bytes32 indexed collateralAssetCode, uint256 collateralAmount)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) FilterMint(opts *bind.FilterOpts, assetAddress []common.Address, collateralAssetCode [][32]byte) (*DigitalReserveSystemMintIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var collateralAssetCodeRule []interface{}
	for _, collateralAssetCodeItem := range collateralAssetCode {
		collateralAssetCodeRule = append(collateralAssetCodeRule, collateralAssetCodeItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.FilterLogs(opts, "Mint", assetAddressRule, collateralAssetCodeRule)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemMintIterator{contract: _DigitalReserveSystem.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xddd98b18c74c43aa2894d989d78dde03a92ddf9e04a53b5655fdfdb73ddd06dc.
//
// Solidity: event Mint(string assetCode, uint256 mintAmount, address indexed assetAddress, bytes32 indexed collateralAssetCode, uint256 collateralAmount)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DigitalReserveSystemMint, assetAddress []common.Address, collateralAssetCode [][32]byte) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var collateralAssetCodeRule []interface{}
	for _, collateralAssetCodeItem := range collateralAssetCode {
		collateralAssetCodeRule = append(collateralAssetCodeRule, collateralAssetCodeItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.WatchLogs(opts, "Mint", assetAddressRule, collateralAssetCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalReserveSystemMint)
				if err := _DigitalReserveSystem.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xddd98b18c74c43aa2894d989d78dde03a92ddf9e04a53b5655fdfdb73ddd06dc.
//
// Solidity: event Mint(string assetCode, uint256 mintAmount, address indexed assetAddress, bytes32 indexed collateralAssetCode, uint256 collateralAmount)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) ParseMint(log types.Log) (*DigitalReserveSystemMint, error) {
	event := new(DigitalReserveSystemMint)
	if err := _DigitalReserveSystem.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DigitalReserveSystemRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the DigitalReserveSystem contract.
type DigitalReserveSystemRedeemIterator struct {
	Event *DigitalReserveSystemRedeem // Event containing the contract specifics and raw log

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
func (it *DigitalReserveSystemRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalReserveSystemRedeem)
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
		it.Event = new(DigitalReserveSystemRedeem)
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
func (it *DigitalReserveSystemRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalReserveSystemRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalReserveSystemRedeem represents a Redeem event raised by the DigitalReserveSystem contract.
type DigitalReserveSystemRedeem struct {
	AssetCode              string
	StableCreditAmount     *big.Int
	CollateralAssetAddress common.Address
	CollateralAssetCode    [32]byte
	CollateralAmount       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xa10c2c65339d202e362f405266f364b6d67e3d643f28be672e5629015930e94e.
//
// Solidity: event Redeem(string assetCode, uint256 stableCreditAmount, address indexed collateralAssetAddress, bytes32 indexed collateralAssetCode, uint256 collateralAmount)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) FilterRedeem(opts *bind.FilterOpts, collateralAssetAddress []common.Address, collateralAssetCode [][32]byte) (*DigitalReserveSystemRedeemIterator, error) {

	var collateralAssetAddressRule []interface{}
	for _, collateralAssetAddressItem := range collateralAssetAddress {
		collateralAssetAddressRule = append(collateralAssetAddressRule, collateralAssetAddressItem)
	}
	var collateralAssetCodeRule []interface{}
	for _, collateralAssetCodeItem := range collateralAssetCode {
		collateralAssetCodeRule = append(collateralAssetCodeRule, collateralAssetCodeItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.FilterLogs(opts, "Redeem", collateralAssetAddressRule, collateralAssetCodeRule)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemRedeemIterator{contract: _DigitalReserveSystem.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xa10c2c65339d202e362f405266f364b6d67e3d643f28be672e5629015930e94e.
//
// Solidity: event Redeem(string assetCode, uint256 stableCreditAmount, address indexed collateralAssetAddress, bytes32 indexed collateralAssetCode, uint256 collateralAmount)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *DigitalReserveSystemRedeem, collateralAssetAddress []common.Address, collateralAssetCode [][32]byte) (event.Subscription, error) {

	var collateralAssetAddressRule []interface{}
	for _, collateralAssetAddressItem := range collateralAssetAddress {
		collateralAssetAddressRule = append(collateralAssetAddressRule, collateralAssetAddressItem)
	}
	var collateralAssetCodeRule []interface{}
	for _, collateralAssetCodeItem := range collateralAssetCode {
		collateralAssetCodeRule = append(collateralAssetCodeRule, collateralAssetCodeItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.WatchLogs(opts, "Redeem", collateralAssetAddressRule, collateralAssetCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalReserveSystemRedeem)
				if err := _DigitalReserveSystem.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xa10c2c65339d202e362f405266f364b6d67e3d643f28be672e5629015930e94e.
//
// Solidity: event Redeem(string assetCode, uint256 stableCreditAmount, address indexed collateralAssetAddress, bytes32 indexed collateralAssetCode, uint256 collateralAmount)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) ParseRedeem(log types.Log) (*DigitalReserveSystemRedeem, error) {
	event := new(DigitalReserveSystemRedeem)
	if err := _DigitalReserveSystem.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DigitalReserveSystemSetupIterator is returned from FilterSetup and is used to iterate over the raw logs and unpacked data for Setup events raised by the DigitalReserveSystem contract.
type DigitalReserveSystemSetupIterator struct {
	Event *DigitalReserveSystemSetup // Event containing the contract specifics and raw log

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
func (it *DigitalReserveSystemSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalReserveSystemSetup)
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
		it.Event = new(DigitalReserveSystemSetup)
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
func (it *DigitalReserveSystemSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalReserveSystemSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalReserveSystemSetup represents a Setup event raised by the DigitalReserveSystem contract.
type DigitalReserveSystemSetup struct {
	AssetCode           string
	PeggedCurrency      [32]byte
	PeggedValue         *big.Int
	CollateralAssetCode [32]byte
	AssetAddress        common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetup is a free log retrieval operation binding the contract event 0xcab34b46abfe0198ee940493abe1167b55efc2b24ff7c43c988360d197c962df.
//
// Solidity: event Setup(string assetCode, bytes32 peggedCurrency, uint256 peggedValue, bytes32 indexed collateralAssetCode, address assetAddress)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) FilterSetup(opts *bind.FilterOpts, collateralAssetCode [][32]byte) (*DigitalReserveSystemSetupIterator, error) {

	var collateralAssetCodeRule []interface{}
	for _, collateralAssetCodeItem := range collateralAssetCode {
		collateralAssetCodeRule = append(collateralAssetCodeRule, collateralAssetCodeItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.FilterLogs(opts, "Setup", collateralAssetCodeRule)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemSetupIterator{contract: _DigitalReserveSystem.contract, event: "Setup", logs: logs, sub: sub}, nil
}

// WatchSetup is a free log subscription operation binding the contract event 0xcab34b46abfe0198ee940493abe1167b55efc2b24ff7c43c988360d197c962df.
//
// Solidity: event Setup(string assetCode, bytes32 peggedCurrency, uint256 peggedValue, bytes32 indexed collateralAssetCode, address assetAddress)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) WatchSetup(opts *bind.WatchOpts, sink chan<- *DigitalReserveSystemSetup, collateralAssetCode [][32]byte) (event.Subscription, error) {

	var collateralAssetCodeRule []interface{}
	for _, collateralAssetCodeItem := range collateralAssetCode {
		collateralAssetCodeRule = append(collateralAssetCodeRule, collateralAssetCodeItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.WatchLogs(opts, "Setup", collateralAssetCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalReserveSystemSetup)
				if err := _DigitalReserveSystem.contract.UnpackLog(event, "Setup", log); err != nil {
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

// ParseSetup is a log parse operation binding the contract event 0xcab34b46abfe0198ee940493abe1167b55efc2b24ff7c43c988360d197c962df.
//
// Solidity: event Setup(string assetCode, bytes32 peggedCurrency, uint256 peggedValue, bytes32 indexed collateralAssetCode, address assetAddress)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) ParseSetup(log types.Log) (*DigitalReserveSystemSetup, error) {
	event := new(DigitalReserveSystemSetup)
	if err := _DigitalReserveSystem.contract.UnpackLog(event, "Setup", log); err != nil {
		return nil, err
	}
	return event, nil
}
