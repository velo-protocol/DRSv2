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

// HeartABI is the input ABI used to generate the binding from.
const HeartABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collateralAssets\",\"outputs\":[{\"internalType\":\"contractICollateralAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collateralRatios\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collectedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditIssuanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeeders\",\"outputs\":[{\"internalType\":\"contractIPF\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"reserveFreeze\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveManager\",\"outputs\":[{\"internalType\":\"contractIRM\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stableCredits\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableCreditsLL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"llSize\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedPartners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReserveManager\",\"type\":\"address\"}],\"name\":\"setReserveManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserveManager\",\"outputs\":[{\"internalType\":\"contractIRM\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"newSeconds\",\"type\":\"uint32\"}],\"name\":\"setReserveFreeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getReserveFreeze\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDrsAddress\",\"type\":\"address\"}],\"name\":\"setDrsAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDrsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setCollateralAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollateralAsset\",\"outputs\":[{\"internalType\":\"contractICollateralAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setCollateralRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setCreditIssuanceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreditIssuanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTrustedPartner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isTrustedPartner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceFeeders\",\"type\":\"address\"}],\"name\":\"setPriceFeeders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceFeeders\",\"outputs\":[{\"internalType\":\"contractIPF\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"}],\"name\":\"collectFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollectedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"newStableCredit\",\"type\":\"address\"}],\"name\":\"addStableCredit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stableCreditId\",\"type\":\"bytes32\"}],\"name\":\"getStableCreditById\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecentStableCredit\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stableCreditId\",\"type\":\"bytes32\"}],\"name\":\"getNextStableCredit\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStableCreditCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setAllowedLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"}],\"name\":\"isLinkAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Heart is an auto generated Go binding around an Ethereum contract.
type Heart struct {
	HeartCaller     // Read-only binding to the contract
	HeartTransactor // Write-only binding to the contract
	HeartFilterer   // Log filterer for contract events
}

// HeartCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeartCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeartTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeartTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeartFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeartFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeartSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeartSession struct {
	Contract     *Heart            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeartCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeartCallerSession struct {
	Contract *HeartCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HeartTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeartTransactorSession struct {
	Contract     *HeartTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeartRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeartRaw struct {
	Contract *Heart // Generic contract binding to access the raw methods on
}

// HeartCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeartCallerRaw struct {
	Contract *HeartCaller // Generic read-only contract binding to access the raw methods on
}

// HeartTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeartTransactorRaw struct {
	Contract *HeartTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeart creates a new instance of Heart, bound to a specific deployed contract.
func NewHeart(address common.Address, backend bind.ContractBackend) (*Heart, error) {
	contract, err := bindHeart(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Heart{HeartCaller: HeartCaller{contract: contract}, HeartTransactor: HeartTransactor{contract: contract}, HeartFilterer: HeartFilterer{contract: contract}}, nil
}

// NewHeartCaller creates a new read-only instance of Heart, bound to a specific deployed contract.
func NewHeartCaller(address common.Address, caller bind.ContractCaller) (*HeartCaller, error) {
	contract, err := bindHeart(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeartCaller{contract: contract}, nil
}

// NewHeartTransactor creates a new write-only instance of Heart, bound to a specific deployed contract.
func NewHeartTransactor(address common.Address, transactor bind.ContractTransactor) (*HeartTransactor, error) {
	contract, err := bindHeart(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeartTransactor{contract: contract}, nil
}

// NewHeartFilterer creates a new log filterer instance of Heart, bound to a specific deployed contract.
func NewHeartFilterer(address common.Address, filterer bind.ContractFilterer) (*HeartFilterer, error) {
	contract, err := bindHeart(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeartFilterer{contract: contract}, nil
}

// bindHeart binds a generic wrapper to an already deployed contract.
func bindHeart(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeartABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heart *HeartRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Heart.Contract.HeartCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heart *HeartRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heart.Contract.HeartTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heart *HeartRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heart.Contract.HeartTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heart *HeartCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Heart.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heart *HeartTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heart.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heart *HeartTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heart.Contract.contract.Transact(opts, method, params...)
}

// CollateralAssets is a free data retrieval call binding the contract method 0x0f09f97c.
//
// Solidity: function collateralAssets(bytes32 ) constant returns(address)
func (_Heart *HeartCaller) CollateralAssets(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "collateralAssets", arg0)
	return *ret0, err
}

// CollateralAssets is a free data retrieval call binding the contract method 0x0f09f97c.
//
// Solidity: function collateralAssets(bytes32 ) constant returns(address)
func (_Heart *HeartSession) CollateralAssets(arg0 [32]byte) (common.Address, error) {
	return _Heart.Contract.CollateralAssets(&_Heart.CallOpts, arg0)
}

// CollateralAssets is a free data retrieval call binding the contract method 0x0f09f97c.
//
// Solidity: function collateralAssets(bytes32 ) constant returns(address)
func (_Heart *HeartCallerSession) CollateralAssets(arg0 [32]byte) (common.Address, error) {
	return _Heart.Contract.CollateralAssets(&_Heart.CallOpts, arg0)
}

// CollateralRatios is a free data retrieval call binding the contract method 0x7c0fa9a7.
//
// Solidity: function collateralRatios(bytes32 ) constant returns(uint256)
func (_Heart *HeartCaller) CollateralRatios(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "collateralRatios", arg0)
	return *ret0, err
}

// CollateralRatios is a free data retrieval call binding the contract method 0x7c0fa9a7.
//
// Solidity: function collateralRatios(bytes32 ) constant returns(uint256)
func (_Heart *HeartSession) CollateralRatios(arg0 [32]byte) (*big.Int, error) {
	return _Heart.Contract.CollateralRatios(&_Heart.CallOpts, arg0)
}

// CollateralRatios is a free data retrieval call binding the contract method 0x7c0fa9a7.
//
// Solidity: function collateralRatios(bytes32 ) constant returns(uint256)
func (_Heart *HeartCallerSession) CollateralRatios(arg0 [32]byte) (*big.Int, error) {
	return _Heart.Contract.CollateralRatios(&_Heart.CallOpts, arg0)
}

// CollectedFee is a free data retrieval call binding the contract method 0xd0d4212e.
//
// Solidity: function collectedFee(bytes32 ) constant returns(uint256)
func (_Heart *HeartCaller) CollectedFee(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "collectedFee", arg0)
	return *ret0, err
}

// CollectedFee is a free data retrieval call binding the contract method 0xd0d4212e.
//
// Solidity: function collectedFee(bytes32 ) constant returns(uint256)
func (_Heart *HeartSession) CollectedFee(arg0 [32]byte) (*big.Int, error) {
	return _Heart.Contract.CollectedFee(&_Heart.CallOpts, arg0)
}

// CollectedFee is a free data retrieval call binding the contract method 0xd0d4212e.
//
// Solidity: function collectedFee(bytes32 ) constant returns(uint256)
func (_Heart *HeartCallerSession) CollectedFee(arg0 [32]byte) (*big.Int, error) {
	return _Heart.Contract.CollectedFee(&_Heart.CallOpts, arg0)
}

// CreditIssuanceFee is a free data retrieval call binding the contract method 0x1f6f2072.
//
// Solidity: function creditIssuanceFee() constant returns(uint256)
func (_Heart *HeartCaller) CreditIssuanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "creditIssuanceFee")
	return *ret0, err
}

// CreditIssuanceFee is a free data retrieval call binding the contract method 0x1f6f2072.
//
// Solidity: function creditIssuanceFee() constant returns(uint256)
func (_Heart *HeartSession) CreditIssuanceFee() (*big.Int, error) {
	return _Heart.Contract.CreditIssuanceFee(&_Heart.CallOpts)
}

// CreditIssuanceFee is a free data retrieval call binding the contract method 0x1f6f2072.
//
// Solidity: function creditIssuanceFee() constant returns(uint256)
func (_Heart *HeartCallerSession) CreditIssuanceFee() (*big.Int, error) {
	return _Heart.Contract.CreditIssuanceFee(&_Heart.CallOpts)
}

// DrsAddr is a free data retrieval call binding the contract method 0x716ec8ee.
//
// Solidity: function drsAddr() constant returns(address)
func (_Heart *HeartCaller) DrsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "drsAddr")
	return *ret0, err
}

// DrsAddr is a free data retrieval call binding the contract method 0x716ec8ee.
//
// Solidity: function drsAddr() constant returns(address)
func (_Heart *HeartSession) DrsAddr() (common.Address, error) {
	return _Heart.Contract.DrsAddr(&_Heart.CallOpts)
}

// DrsAddr is a free data retrieval call binding the contract method 0x716ec8ee.
//
// Solidity: function drsAddr() constant returns(address)
func (_Heart *HeartCallerSession) DrsAddr() (common.Address, error) {
	return _Heart.Contract.DrsAddr(&_Heart.CallOpts)
}

// GetCollateralAsset is a free data retrieval call binding the contract method 0xd476f04b.
//
// Solidity: function getCollateralAsset(bytes32 assetCode) constant returns(address)
func (_Heart *HeartCaller) GetCollateralAsset(opts *bind.CallOpts, assetCode [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getCollateralAsset", assetCode)
	return *ret0, err
}

// GetCollateralAsset is a free data retrieval call binding the contract method 0xd476f04b.
//
// Solidity: function getCollateralAsset(bytes32 assetCode) constant returns(address)
func (_Heart *HeartSession) GetCollateralAsset(assetCode [32]byte) (common.Address, error) {
	return _Heart.Contract.GetCollateralAsset(&_Heart.CallOpts, assetCode)
}

// GetCollateralAsset is a free data retrieval call binding the contract method 0xd476f04b.
//
// Solidity: function getCollateralAsset(bytes32 assetCode) constant returns(address)
func (_Heart *HeartCallerSession) GetCollateralAsset(assetCode [32]byte) (common.Address, error) {
	return _Heart.Contract.GetCollateralAsset(&_Heart.CallOpts, assetCode)
}

// GetCollateralRatio is a free data retrieval call binding the contract method 0x0fb79eb2.
//
// Solidity: function getCollateralRatio(bytes32 assetCode) constant returns(uint256)
func (_Heart *HeartCaller) GetCollateralRatio(opts *bind.CallOpts, assetCode [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getCollateralRatio", assetCode)
	return *ret0, err
}

// GetCollateralRatio is a free data retrieval call binding the contract method 0x0fb79eb2.
//
// Solidity: function getCollateralRatio(bytes32 assetCode) constant returns(uint256)
func (_Heart *HeartSession) GetCollateralRatio(assetCode [32]byte) (*big.Int, error) {
	return _Heart.Contract.GetCollateralRatio(&_Heart.CallOpts, assetCode)
}

// GetCollateralRatio is a free data retrieval call binding the contract method 0x0fb79eb2.
//
// Solidity: function getCollateralRatio(bytes32 assetCode) constant returns(uint256)
func (_Heart *HeartCallerSession) GetCollateralRatio(assetCode [32]byte) (*big.Int, error) {
	return _Heart.Contract.GetCollateralRatio(&_Heart.CallOpts, assetCode)
}

// GetCollectedFee is a free data retrieval call binding the contract method 0xf9af9ce0.
//
// Solidity: function getCollectedFee(bytes32 collateralAssetCode) constant returns(uint256)
func (_Heart *HeartCaller) GetCollectedFee(opts *bind.CallOpts, collateralAssetCode [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getCollectedFee", collateralAssetCode)
	return *ret0, err
}

// GetCollectedFee is a free data retrieval call binding the contract method 0xf9af9ce0.
//
// Solidity: function getCollectedFee(bytes32 collateralAssetCode) constant returns(uint256)
func (_Heart *HeartSession) GetCollectedFee(collateralAssetCode [32]byte) (*big.Int, error) {
	return _Heart.Contract.GetCollectedFee(&_Heart.CallOpts, collateralAssetCode)
}

// GetCollectedFee is a free data retrieval call binding the contract method 0xf9af9ce0.
//
// Solidity: function getCollectedFee(bytes32 collateralAssetCode) constant returns(uint256)
func (_Heart *HeartCallerSession) GetCollectedFee(collateralAssetCode [32]byte) (*big.Int, error) {
	return _Heart.Contract.GetCollectedFee(&_Heart.CallOpts, collateralAssetCode)
}

// GetCreditIssuanceFee is a free data retrieval call binding the contract method 0x016b3a11.
//
// Solidity: function getCreditIssuanceFee() constant returns(uint256)
func (_Heart *HeartCaller) GetCreditIssuanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getCreditIssuanceFee")
	return *ret0, err
}

// GetCreditIssuanceFee is a free data retrieval call binding the contract method 0x016b3a11.
//
// Solidity: function getCreditIssuanceFee() constant returns(uint256)
func (_Heart *HeartSession) GetCreditIssuanceFee() (*big.Int, error) {
	return _Heart.Contract.GetCreditIssuanceFee(&_Heart.CallOpts)
}

// GetCreditIssuanceFee is a free data retrieval call binding the contract method 0x016b3a11.
//
// Solidity: function getCreditIssuanceFee() constant returns(uint256)
func (_Heart *HeartCallerSession) GetCreditIssuanceFee() (*big.Int, error) {
	return _Heart.Contract.GetCreditIssuanceFee(&_Heart.CallOpts)
}

// GetDrsAddress is a free data retrieval call binding the contract method 0x70fa9a6d.
//
// Solidity: function getDrsAddress() constant returns(address)
func (_Heart *HeartCaller) GetDrsAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getDrsAddress")
	return *ret0, err
}

// GetDrsAddress is a free data retrieval call binding the contract method 0x70fa9a6d.
//
// Solidity: function getDrsAddress() constant returns(address)
func (_Heart *HeartSession) GetDrsAddress() (common.Address, error) {
	return _Heart.Contract.GetDrsAddress(&_Heart.CallOpts)
}

// GetDrsAddress is a free data retrieval call binding the contract method 0x70fa9a6d.
//
// Solidity: function getDrsAddress() constant returns(address)
func (_Heart *HeartCallerSession) GetDrsAddress() (common.Address, error) {
	return _Heart.Contract.GetDrsAddress(&_Heart.CallOpts)
}

// GetNextStableCredit is a free data retrieval call binding the contract method 0xbb31146e.
//
// Solidity: function getNextStableCredit(bytes32 stableCreditId) constant returns(address)
func (_Heart *HeartCaller) GetNextStableCredit(opts *bind.CallOpts, stableCreditId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getNextStableCredit", stableCreditId)
	return *ret0, err
}

// GetNextStableCredit is a free data retrieval call binding the contract method 0xbb31146e.
//
// Solidity: function getNextStableCredit(bytes32 stableCreditId) constant returns(address)
func (_Heart *HeartSession) GetNextStableCredit(stableCreditId [32]byte) (common.Address, error) {
	return _Heart.Contract.GetNextStableCredit(&_Heart.CallOpts, stableCreditId)
}

// GetNextStableCredit is a free data retrieval call binding the contract method 0xbb31146e.
//
// Solidity: function getNextStableCredit(bytes32 stableCreditId) constant returns(address)
func (_Heart *HeartCallerSession) GetNextStableCredit(stableCreditId [32]byte) (common.Address, error) {
	return _Heart.Contract.GetNextStableCredit(&_Heart.CallOpts, stableCreditId)
}

// GetPriceFeeders is a free data retrieval call binding the contract method 0x794e3db0.
//
// Solidity: function getPriceFeeders() constant returns(address)
func (_Heart *HeartCaller) GetPriceFeeders(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getPriceFeeders")
	return *ret0, err
}

// GetPriceFeeders is a free data retrieval call binding the contract method 0x794e3db0.
//
// Solidity: function getPriceFeeders() constant returns(address)
func (_Heart *HeartSession) GetPriceFeeders() (common.Address, error) {
	return _Heart.Contract.GetPriceFeeders(&_Heart.CallOpts)
}

// GetPriceFeeders is a free data retrieval call binding the contract method 0x794e3db0.
//
// Solidity: function getPriceFeeders() constant returns(address)
func (_Heart *HeartCallerSession) GetPriceFeeders() (common.Address, error) {
	return _Heart.Contract.GetPriceFeeders(&_Heart.CallOpts)
}

// GetRecentStableCredit is a free data retrieval call binding the contract method 0x33aec2fd.
//
// Solidity: function getRecentStableCredit() constant returns(address)
func (_Heart *HeartCaller) GetRecentStableCredit(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getRecentStableCredit")
	return *ret0, err
}

// GetRecentStableCredit is a free data retrieval call binding the contract method 0x33aec2fd.
//
// Solidity: function getRecentStableCredit() constant returns(address)
func (_Heart *HeartSession) GetRecentStableCredit() (common.Address, error) {
	return _Heart.Contract.GetRecentStableCredit(&_Heart.CallOpts)
}

// GetRecentStableCredit is a free data retrieval call binding the contract method 0x33aec2fd.
//
// Solidity: function getRecentStableCredit() constant returns(address)
func (_Heart *HeartCallerSession) GetRecentStableCredit() (common.Address, error) {
	return _Heart.Contract.GetRecentStableCredit(&_Heart.CallOpts)
}

// GetReserveFreeze is a free data retrieval call binding the contract method 0xbd09b85d.
//
// Solidity: function getReserveFreeze(bytes32 assetCode) constant returns(uint32)
func (_Heart *HeartCaller) GetReserveFreeze(opts *bind.CallOpts, assetCode [32]byte) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getReserveFreeze", assetCode)
	return *ret0, err
}

// GetReserveFreeze is a free data retrieval call binding the contract method 0xbd09b85d.
//
// Solidity: function getReserveFreeze(bytes32 assetCode) constant returns(uint32)
func (_Heart *HeartSession) GetReserveFreeze(assetCode [32]byte) (uint32, error) {
	return _Heart.Contract.GetReserveFreeze(&_Heart.CallOpts, assetCode)
}

// GetReserveFreeze is a free data retrieval call binding the contract method 0xbd09b85d.
//
// Solidity: function getReserveFreeze(bytes32 assetCode) constant returns(uint32)
func (_Heart *HeartCallerSession) GetReserveFreeze(assetCode [32]byte) (uint32, error) {
	return _Heart.Contract.GetReserveFreeze(&_Heart.CallOpts, assetCode)
}

// GetReserveManager is a free data retrieval call binding the contract method 0x81c210f8.
//
// Solidity: function getReserveManager() constant returns(address)
func (_Heart *HeartCaller) GetReserveManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getReserveManager")
	return *ret0, err
}

// GetReserveManager is a free data retrieval call binding the contract method 0x81c210f8.
//
// Solidity: function getReserveManager() constant returns(address)
func (_Heart *HeartSession) GetReserveManager() (common.Address, error) {
	return _Heart.Contract.GetReserveManager(&_Heart.CallOpts)
}

// GetReserveManager is a free data retrieval call binding the contract method 0x81c210f8.
//
// Solidity: function getReserveManager() constant returns(address)
func (_Heart *HeartCallerSession) GetReserveManager() (common.Address, error) {
	return _Heart.Contract.GetReserveManager(&_Heart.CallOpts)
}

// GetStableCreditById is a free data retrieval call binding the contract method 0xf1e476a6.
//
// Solidity: function getStableCreditById(bytes32 stableCreditId) constant returns(address)
func (_Heart *HeartCaller) GetStableCreditById(opts *bind.CallOpts, stableCreditId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getStableCreditById", stableCreditId)
	return *ret0, err
}

// GetStableCreditById is a free data retrieval call binding the contract method 0xf1e476a6.
//
// Solidity: function getStableCreditById(bytes32 stableCreditId) constant returns(address)
func (_Heart *HeartSession) GetStableCreditById(stableCreditId [32]byte) (common.Address, error) {
	return _Heart.Contract.GetStableCreditById(&_Heart.CallOpts, stableCreditId)
}

// GetStableCreditById is a free data retrieval call binding the contract method 0xf1e476a6.
//
// Solidity: function getStableCreditById(bytes32 stableCreditId) constant returns(address)
func (_Heart *HeartCallerSession) GetStableCreditById(stableCreditId [32]byte) (common.Address, error) {
	return _Heart.Contract.GetStableCreditById(&_Heart.CallOpts, stableCreditId)
}

// GetStableCreditCount is a free data retrieval call binding the contract method 0x573ae02a.
//
// Solidity: function getStableCreditCount() constant returns(uint8)
func (_Heart *HeartCaller) GetStableCreditCount(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getStableCreditCount")
	return *ret0, err
}

// GetStableCreditCount is a free data retrieval call binding the contract method 0x573ae02a.
//
// Solidity: function getStableCreditCount() constant returns(uint8)
func (_Heart *HeartSession) GetStableCreditCount() (uint8, error) {
	return _Heart.Contract.GetStableCreditCount(&_Heart.CallOpts)
}

// GetStableCreditCount is a free data retrieval call binding the contract method 0x573ae02a.
//
// Solidity: function getStableCreditCount() constant returns(uint8)
func (_Heart *HeartCallerSession) GetStableCreditCount() (uint8, error) {
	return _Heart.Contract.GetStableCreditCount(&_Heart.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x70ec42e8.
//
// Solidity: function governor(address ) constant returns(bool)
func (_Heart *HeartCaller) Governor(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "governor", arg0)
	return *ret0, err
}

// Governor is a free data retrieval call binding the contract method 0x70ec42e8.
//
// Solidity: function governor(address ) constant returns(bool)
func (_Heart *HeartSession) Governor(arg0 common.Address) (bool, error) {
	return _Heart.Contract.Governor(&_Heart.CallOpts, arg0)
}

// Governor is a free data retrieval call binding the contract method 0x70ec42e8.
//
// Solidity: function governor(address ) constant returns(bool)
func (_Heart *HeartCallerSession) Governor(arg0 common.Address) (bool, error) {
	return _Heart.Contract.Governor(&_Heart.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address addr) constant returns(bool)
func (_Heart *HeartCaller) IsGovernor(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "isGovernor", addr)
	return *ret0, err
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address addr) constant returns(bool)
func (_Heart *HeartSession) IsGovernor(addr common.Address) (bool, error) {
	return _Heart.Contract.IsGovernor(&_Heart.CallOpts, addr)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address addr) constant returns(bool)
func (_Heart *HeartCallerSession) IsGovernor(addr common.Address) (bool, error) {
	return _Heart.Contract.IsGovernor(&_Heart.CallOpts, addr)
}

// IsLinkAllowed is a free data retrieval call binding the contract method 0xc084770d.
//
// Solidity: function isLinkAllowed(bytes32 linkId) constant returns(bool)
func (_Heart *HeartCaller) IsLinkAllowed(opts *bind.CallOpts, linkId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "isLinkAllowed", linkId)
	return *ret0, err
}

// IsLinkAllowed is a free data retrieval call binding the contract method 0xc084770d.
//
// Solidity: function isLinkAllowed(bytes32 linkId) constant returns(bool)
func (_Heart *HeartSession) IsLinkAllowed(linkId [32]byte) (bool, error) {
	return _Heart.Contract.IsLinkAllowed(&_Heart.CallOpts, linkId)
}

// IsLinkAllowed is a free data retrieval call binding the contract method 0xc084770d.
//
// Solidity: function isLinkAllowed(bytes32 linkId) constant returns(bool)
func (_Heart *HeartCallerSession) IsLinkAllowed(linkId [32]byte) (bool, error) {
	return _Heart.Contract.IsLinkAllowed(&_Heart.CallOpts, linkId)
}

// IsTrustedPartner is a free data retrieval call binding the contract method 0x449c3b74.
//
// Solidity: function isTrustedPartner(address addr) constant returns(bool)
func (_Heart *HeartCaller) IsTrustedPartner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "isTrustedPartner", addr)
	return *ret0, err
}

// IsTrustedPartner is a free data retrieval call binding the contract method 0x449c3b74.
//
// Solidity: function isTrustedPartner(address addr) constant returns(bool)
func (_Heart *HeartSession) IsTrustedPartner(addr common.Address) (bool, error) {
	return _Heart.Contract.IsTrustedPartner(&_Heart.CallOpts, addr)
}

// IsTrustedPartner is a free data retrieval call binding the contract method 0x449c3b74.
//
// Solidity: function isTrustedPartner(address addr) constant returns(bool)
func (_Heart *HeartCallerSession) IsTrustedPartner(addr common.Address) (bool, error) {
	return _Heart.Contract.IsTrustedPartner(&_Heart.CallOpts, addr)
}

// PriceFeeders is a free data retrieval call binding the contract method 0xadd36f68.
//
// Solidity: function priceFeeders() constant returns(address)
func (_Heart *HeartCaller) PriceFeeders(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "priceFeeders")
	return *ret0, err
}

// PriceFeeders is a free data retrieval call binding the contract method 0xadd36f68.
//
// Solidity: function priceFeeders() constant returns(address)
func (_Heart *HeartSession) PriceFeeders() (common.Address, error) {
	return _Heart.Contract.PriceFeeders(&_Heart.CallOpts)
}

// PriceFeeders is a free data retrieval call binding the contract method 0xadd36f68.
//
// Solidity: function priceFeeders() constant returns(address)
func (_Heart *HeartCallerSession) PriceFeeders() (common.Address, error) {
	return _Heart.Contract.PriceFeeders(&_Heart.CallOpts)
}

// ReserveFreeze is a free data retrieval call binding the contract method 0x78f4dce8.
//
// Solidity: function reserveFreeze(bytes32 ) constant returns(uint32)
func (_Heart *HeartCaller) ReserveFreeze(opts *bind.CallOpts, arg0 [32]byte) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "reserveFreeze", arg0)
	return *ret0, err
}

// ReserveFreeze is a free data retrieval call binding the contract method 0x78f4dce8.
//
// Solidity: function reserveFreeze(bytes32 ) constant returns(uint32)
func (_Heart *HeartSession) ReserveFreeze(arg0 [32]byte) (uint32, error) {
	return _Heart.Contract.ReserveFreeze(&_Heart.CallOpts, arg0)
}

// ReserveFreeze is a free data retrieval call binding the contract method 0x78f4dce8.
//
// Solidity: function reserveFreeze(bytes32 ) constant returns(uint32)
func (_Heart *HeartCallerSession) ReserveFreeze(arg0 [32]byte) (uint32, error) {
	return _Heart.Contract.ReserveFreeze(&_Heart.CallOpts, arg0)
}

// ReserveManager is a free data retrieval call binding the contract method 0xbb004abc.
//
// Solidity: function reserveManager() constant returns(address)
func (_Heart *HeartCaller) ReserveManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "reserveManager")
	return *ret0, err
}

// ReserveManager is a free data retrieval call binding the contract method 0xbb004abc.
//
// Solidity: function reserveManager() constant returns(address)
func (_Heart *HeartSession) ReserveManager() (common.Address, error) {
	return _Heart.Contract.ReserveManager(&_Heart.CallOpts)
}

// ReserveManager is a free data retrieval call binding the contract method 0xbb004abc.
//
// Solidity: function reserveManager() constant returns(address)
func (_Heart *HeartCallerSession) ReserveManager() (common.Address, error) {
	return _Heart.Contract.ReserveManager(&_Heart.CallOpts)
}

// StableCredits is a free data retrieval call binding the contract method 0x22042f9f.
//
// Solidity: function stableCredits(bytes32 ) constant returns(address)
func (_Heart *HeartCaller) StableCredits(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "stableCredits", arg0)
	return *ret0, err
}

// StableCredits is a free data retrieval call binding the contract method 0x22042f9f.
//
// Solidity: function stableCredits(bytes32 ) constant returns(address)
func (_Heart *HeartSession) StableCredits(arg0 [32]byte) (common.Address, error) {
	return _Heart.Contract.StableCredits(&_Heart.CallOpts, arg0)
}

// StableCredits is a free data retrieval call binding the contract method 0x22042f9f.
//
// Solidity: function stableCredits(bytes32 ) constant returns(address)
func (_Heart *HeartCallerSession) StableCredits(arg0 [32]byte) (common.Address, error) {
	return _Heart.Contract.StableCredits(&_Heart.CallOpts, arg0)
}

// StableCreditsLL is a free data retrieval call binding the contract method 0x8019155e.
//
// Solidity: function stableCreditsLL() constant returns(uint8 llSize)
func (_Heart *HeartCaller) StableCreditsLL(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "stableCreditsLL")
	return *ret0, err
}

// StableCreditsLL is a free data retrieval call binding the contract method 0x8019155e.
//
// Solidity: function stableCreditsLL() constant returns(uint8 llSize)
func (_Heart *HeartSession) StableCreditsLL() (uint8, error) {
	return _Heart.Contract.StableCreditsLL(&_Heart.CallOpts)
}

// StableCreditsLL is a free data retrieval call binding the contract method 0x8019155e.
//
// Solidity: function stableCreditsLL() constant returns(uint8 llSize)
func (_Heart *HeartCallerSession) StableCreditsLL() (uint8, error) {
	return _Heart.Contract.StableCreditsLL(&_Heart.CallOpts)
}

// TrustedPartners is a free data retrieval call binding the contract method 0x29fc9ec8.
//
// Solidity: function trustedPartners(address ) constant returns(bool)
func (_Heart *HeartCaller) TrustedPartners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "trustedPartners", arg0)
	return *ret0, err
}

// TrustedPartners is a free data retrieval call binding the contract method 0x29fc9ec8.
//
// Solidity: function trustedPartners(address ) constant returns(bool)
func (_Heart *HeartSession) TrustedPartners(arg0 common.Address) (bool, error) {
	return _Heart.Contract.TrustedPartners(&_Heart.CallOpts, arg0)
}

// TrustedPartners is a free data retrieval call binding the contract method 0x29fc9ec8.
//
// Solidity: function trustedPartners(address ) constant returns(bool)
func (_Heart *HeartCallerSession) TrustedPartners(arg0 common.Address) (bool, error) {
	return _Heart.Contract.TrustedPartners(&_Heart.CallOpts, arg0)
}

// AddStableCredit is a paid mutator transaction binding the contract method 0xbb33674a.
//
// Solidity: function addStableCredit(address newStableCredit) returns()
func (_Heart *HeartTransactor) AddStableCredit(opts *bind.TransactOpts, newStableCredit common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "addStableCredit", newStableCredit)
}

// AddStableCredit is a paid mutator transaction binding the contract method 0xbb33674a.
//
// Solidity: function addStableCredit(address newStableCredit) returns()
func (_Heart *HeartSession) AddStableCredit(newStableCredit common.Address) (*types.Transaction, error) {
	return _Heart.Contract.AddStableCredit(&_Heart.TransactOpts, newStableCredit)
}

// AddStableCredit is a paid mutator transaction binding the contract method 0xbb33674a.
//
// Solidity: function addStableCredit(address newStableCredit) returns()
func (_Heart *HeartTransactorSession) AddStableCredit(newStableCredit common.Address) (*types.Transaction, error) {
	return _Heart.Contract.AddStableCredit(&_Heart.TransactOpts, newStableCredit)
}

// CollectFee is a paid mutator transaction binding the contract method 0x0171af56.
//
// Solidity: function collectFee(uint256 fee, bytes32 collateralAssetCode) returns()
func (_Heart *HeartTransactor) CollectFee(opts *bind.TransactOpts, fee *big.Int, collateralAssetCode [32]byte) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "collectFee", fee, collateralAssetCode)
}

// CollectFee is a paid mutator transaction binding the contract method 0x0171af56.
//
// Solidity: function collectFee(uint256 fee, bytes32 collateralAssetCode) returns()
func (_Heart *HeartSession) CollectFee(fee *big.Int, collateralAssetCode [32]byte) (*types.Transaction, error) {
	return _Heart.Contract.CollectFee(&_Heart.TransactOpts, fee, collateralAssetCode)
}

// CollectFee is a paid mutator transaction binding the contract method 0x0171af56.
//
// Solidity: function collectFee(uint256 fee, bytes32 collateralAssetCode) returns()
func (_Heart *HeartTransactorSession) CollectFee(fee *big.Int, collateralAssetCode [32]byte) (*types.Transaction, error) {
	return _Heart.Contract.CollectFee(&_Heart.TransactOpts, fee, collateralAssetCode)
}

// SetAllowedLink is a paid mutator transaction binding the contract method 0xe402da4e.
//
// Solidity: function setAllowedLink(bytes32 linkId, bool enable) returns()
func (_Heart *HeartTransactor) SetAllowedLink(opts *bind.TransactOpts, linkId [32]byte, enable bool) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setAllowedLink", linkId, enable)
}

// SetAllowedLink is a paid mutator transaction binding the contract method 0xe402da4e.
//
// Solidity: function setAllowedLink(bytes32 linkId, bool enable) returns()
func (_Heart *HeartSession) SetAllowedLink(linkId [32]byte, enable bool) (*types.Transaction, error) {
	return _Heart.Contract.SetAllowedLink(&_Heart.TransactOpts, linkId, enable)
}

// SetAllowedLink is a paid mutator transaction binding the contract method 0xe402da4e.
//
// Solidity: function setAllowedLink(bytes32 linkId, bool enable) returns()
func (_Heart *HeartTransactorSession) SetAllowedLink(linkId [32]byte, enable bool) (*types.Transaction, error) {
	return _Heart.Contract.SetAllowedLink(&_Heart.TransactOpts, linkId, enable)
}

// SetCollateralAsset is a paid mutator transaction binding the contract method 0x88bed352.
//
// Solidity: function setCollateralAsset(bytes32 assetCode, address addr, uint256 ratio) returns()
func (_Heart *HeartTransactor) SetCollateralAsset(opts *bind.TransactOpts, assetCode [32]byte, addr common.Address, ratio *big.Int) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setCollateralAsset", assetCode, addr, ratio)
}

// SetCollateralAsset is a paid mutator transaction binding the contract method 0x88bed352.
//
// Solidity: function setCollateralAsset(bytes32 assetCode, address addr, uint256 ratio) returns()
func (_Heart *HeartSession) SetCollateralAsset(assetCode [32]byte, addr common.Address, ratio *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.SetCollateralAsset(&_Heart.TransactOpts, assetCode, addr, ratio)
}

// SetCollateralAsset is a paid mutator transaction binding the contract method 0x88bed352.
//
// Solidity: function setCollateralAsset(bytes32 assetCode, address addr, uint256 ratio) returns()
func (_Heart *HeartTransactorSession) SetCollateralAsset(assetCode [32]byte, addr common.Address, ratio *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.SetCollateralAsset(&_Heart.TransactOpts, assetCode, addr, ratio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0xdb4d96d1.
//
// Solidity: function setCollateralRatio(bytes32 assetCode, uint256 ratio) returns()
func (_Heart *HeartTransactor) SetCollateralRatio(opts *bind.TransactOpts, assetCode [32]byte, ratio *big.Int) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setCollateralRatio", assetCode, ratio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0xdb4d96d1.
//
// Solidity: function setCollateralRatio(bytes32 assetCode, uint256 ratio) returns()
func (_Heart *HeartSession) SetCollateralRatio(assetCode [32]byte, ratio *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.SetCollateralRatio(&_Heart.TransactOpts, assetCode, ratio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0xdb4d96d1.
//
// Solidity: function setCollateralRatio(bytes32 assetCode, uint256 ratio) returns()
func (_Heart *HeartTransactorSession) SetCollateralRatio(assetCode [32]byte, ratio *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.SetCollateralRatio(&_Heart.TransactOpts, assetCode, ratio)
}

// SetCreditIssuanceFee is a paid mutator transaction binding the contract method 0xaedb08d4.
//
// Solidity: function setCreditIssuanceFee(uint256 newFee) returns()
func (_Heart *HeartTransactor) SetCreditIssuanceFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setCreditIssuanceFee", newFee)
}

// SetCreditIssuanceFee is a paid mutator transaction binding the contract method 0xaedb08d4.
//
// Solidity: function setCreditIssuanceFee(uint256 newFee) returns()
func (_Heart *HeartSession) SetCreditIssuanceFee(newFee *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.SetCreditIssuanceFee(&_Heart.TransactOpts, newFee)
}

// SetCreditIssuanceFee is a paid mutator transaction binding the contract method 0xaedb08d4.
//
// Solidity: function setCreditIssuanceFee(uint256 newFee) returns()
func (_Heart *HeartTransactorSession) SetCreditIssuanceFee(newFee *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.SetCreditIssuanceFee(&_Heart.TransactOpts, newFee)
}

// SetDrsAddress is a paid mutator transaction binding the contract method 0xe12b39f5.
//
// Solidity: function setDrsAddress(address newDrsAddress) returns()
func (_Heart *HeartTransactor) SetDrsAddress(opts *bind.TransactOpts, newDrsAddress common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setDrsAddress", newDrsAddress)
}

// SetDrsAddress is a paid mutator transaction binding the contract method 0xe12b39f5.
//
// Solidity: function setDrsAddress(address newDrsAddress) returns()
func (_Heart *HeartSession) SetDrsAddress(newDrsAddress common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetDrsAddress(&_Heart.TransactOpts, newDrsAddress)
}

// SetDrsAddress is a paid mutator transaction binding the contract method 0xe12b39f5.
//
// Solidity: function setDrsAddress(address newDrsAddress) returns()
func (_Heart *HeartTransactorSession) SetDrsAddress(newDrsAddress common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetDrsAddress(&_Heart.TransactOpts, newDrsAddress)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address addr) returns()
func (_Heart *HeartTransactor) SetGovernor(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setGovernor", addr)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address addr) returns()
func (_Heart *HeartSession) SetGovernor(addr common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetGovernor(&_Heart.TransactOpts, addr)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address addr) returns()
func (_Heart *HeartTransactorSession) SetGovernor(addr common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetGovernor(&_Heart.TransactOpts, addr)
}

// SetPriceFeeders is a paid mutator transaction binding the contract method 0x1070ff5d.
//
// Solidity: function setPriceFeeders(address newPriceFeeders) returns()
func (_Heart *HeartTransactor) SetPriceFeeders(opts *bind.TransactOpts, newPriceFeeders common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setPriceFeeders", newPriceFeeders)
}

// SetPriceFeeders is a paid mutator transaction binding the contract method 0x1070ff5d.
//
// Solidity: function setPriceFeeders(address newPriceFeeders) returns()
func (_Heart *HeartSession) SetPriceFeeders(newPriceFeeders common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetPriceFeeders(&_Heart.TransactOpts, newPriceFeeders)
}

// SetPriceFeeders is a paid mutator transaction binding the contract method 0x1070ff5d.
//
// Solidity: function setPriceFeeders(address newPriceFeeders) returns()
func (_Heart *HeartTransactorSession) SetPriceFeeders(newPriceFeeders common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetPriceFeeders(&_Heart.TransactOpts, newPriceFeeders)
}

// SetReserveFreeze is a paid mutator transaction binding the contract method 0xf3685b7f.
//
// Solidity: function setReserveFreeze(bytes32 assetCode, uint32 newSeconds) returns()
func (_Heart *HeartTransactor) SetReserveFreeze(opts *bind.TransactOpts, assetCode [32]byte, newSeconds uint32) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setReserveFreeze", assetCode, newSeconds)
}

// SetReserveFreeze is a paid mutator transaction binding the contract method 0xf3685b7f.
//
// Solidity: function setReserveFreeze(bytes32 assetCode, uint32 newSeconds) returns()
func (_Heart *HeartSession) SetReserveFreeze(assetCode [32]byte, newSeconds uint32) (*types.Transaction, error) {
	return _Heart.Contract.SetReserveFreeze(&_Heart.TransactOpts, assetCode, newSeconds)
}

// SetReserveFreeze is a paid mutator transaction binding the contract method 0xf3685b7f.
//
// Solidity: function setReserveFreeze(bytes32 assetCode, uint32 newSeconds) returns()
func (_Heart *HeartTransactorSession) SetReserveFreeze(assetCode [32]byte, newSeconds uint32) (*types.Transaction, error) {
	return _Heart.Contract.SetReserveFreeze(&_Heart.TransactOpts, assetCode, newSeconds)
}

// SetReserveManager is a paid mutator transaction binding the contract method 0xdf7da754.
//
// Solidity: function setReserveManager(address newReserveManager) returns()
func (_Heart *HeartTransactor) SetReserveManager(opts *bind.TransactOpts, newReserveManager common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setReserveManager", newReserveManager)
}

// SetReserveManager is a paid mutator transaction binding the contract method 0xdf7da754.
//
// Solidity: function setReserveManager(address newReserveManager) returns()
func (_Heart *HeartSession) SetReserveManager(newReserveManager common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetReserveManager(&_Heart.TransactOpts, newReserveManager)
}

// SetReserveManager is a paid mutator transaction binding the contract method 0xdf7da754.
//
// Solidity: function setReserveManager(address newReserveManager) returns()
func (_Heart *HeartTransactorSession) SetReserveManager(newReserveManager common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetReserveManager(&_Heart.TransactOpts, newReserveManager)
}

// SetTrustedPartner is a paid mutator transaction binding the contract method 0x68698406.
//
// Solidity: function setTrustedPartner(address addr) returns()
func (_Heart *HeartTransactor) SetTrustedPartner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "setTrustedPartner", addr)
}

// SetTrustedPartner is a paid mutator transaction binding the contract method 0x68698406.
//
// Solidity: function setTrustedPartner(address addr) returns()
func (_Heart *HeartSession) SetTrustedPartner(addr common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetTrustedPartner(&_Heart.TransactOpts, addr)
}

// SetTrustedPartner is a paid mutator transaction binding the contract method 0x68698406.
//
// Solidity: function setTrustedPartner(address addr) returns()
func (_Heart *HeartTransactorSession) SetTrustedPartner(addr common.Address) (*types.Transaction, error) {
	return _Heart.Contract.SetTrustedPartner(&_Heart.TransactOpts, addr)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xd77da9a2.
//
// Solidity: function withdrawFee(bytes32 collateralAssetCode, uint256 amount) returns()
func (_Heart *HeartTransactor) WithdrawFee(opts *bind.TransactOpts, collateralAssetCode [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "withdrawFee", collateralAssetCode, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xd77da9a2.
//
// Solidity: function withdrawFee(bytes32 collateralAssetCode, uint256 amount) returns()
func (_Heart *HeartSession) WithdrawFee(collateralAssetCode [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.WithdrawFee(&_Heart.TransactOpts, collateralAssetCode, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xd77da9a2.
//
// Solidity: function withdrawFee(bytes32 collateralAssetCode, uint256 amount) returns()
func (_Heart *HeartTransactorSession) WithdrawFee(collateralAssetCode [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Heart.Contract.WithdrawFee(&_Heart.TransactOpts, collateralAssetCode, amount)
}
