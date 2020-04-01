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

// PriceABI is the input ABI used to generate the binding from.
const PriceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"price\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"PriceActivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"price\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"PriceVoid\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lagAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLagAddr\",\"type\":\"address\"}],\"name\":\"setLag\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lagAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"post\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PriceBin is the compiled bytecode used for deploying new contracts.
var PriceBin = "0x608060405234801561001057600080fd5b50610842806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636d4ce63c116100715780636d4ce63c146101325780638da5cb5b1461014c578063a035b1fe14610170578063ac4c25b214610178578063e34e380014610180578063fe0a55e1146101a6576100a9565b806302fb0c5e146100ae57806305c1f502146100ca5780630f15f4c0146100d4578063485cc955146100dc578063498b58281461010a575b600080fd5b6100b66101ae565b604080519115158252519081900360200190f35b6100d26101be565b005b6100d261029b565b6100d2600480360360408110156100f257600080fd5b506001600160a01b03813581169160200135166103e3565b6101126104d3565b604080519384529115156020840152151582820152519081900360600190f35b61013a6104f2565b60408051918252519081900360200190f35b610154610615565b604080516001600160a01b039092168252519081900360200190f35b61013a610624565b6100d261062a565b6100d26004803603602081101561019657600080fd5b50356001600160a01b03166106d9565b610154610744565b603554600160a81b900460ff1681565b6000806000603560009054906101000a90046001600160a01b03166001600160a01b031663498b58286040518163ffffffff1660e01b815260040160606040518083038186803b15801561021157600080fd5b505afa158015610225573d6000803e3d6000fd5b505050506040513d606081101561023b57600080fd5b5080516020820151604090920151909450909250905081610264576035805460ff60a81b191690555b6035805460ff60a01b1916600160a01b8315150217905582610294576035805460ff60a01b1916600160a01b1790555b5050603455565b6033546001600160a01b031633146102e45760405162461bcd60e51b81526004018080602001828103825260578152602001806107b76057913960600191505060405180910390fd5b603554600160a81b900460ff1615610343576040805162461bcd60e51b815260206004820152601f60248201527f50726963652e61637469766174653a2070726963652069732061637469766500604482015290519081900360640190fd5b6000603454116103845760405162461bcd60e51b815260040180806020018281038252602f81526020018061075a602f913960400191505060405180910390fd5b6035805460ff60a81b1916600160a81b90811791829055604080513381523060208201529190920460ff1615158183015290517f2e219088454a765156309a15f374957f610cefb0bab9c32ecf58446b3d94b8819181900360600190a1565b600054610100900460ff16806103fc57506103fc610753565b8061040a575060005460ff16155b6104455760405162461bcd60e51b815260040180806020018281038252602e815260200180610789602e913960400191505060405180910390fd5b600054610100900460ff16158015610470576000805460ff1961ff0019909116610100171660011790555b60358054603380546001600160a01b038781166001600160a01b031992831617909255600060345560ff60a01b1960ff60a81b19928716919093161716600160a81b1716600160a01b17905580156104ce576000805461ff00191690555b505050565b60345460355460ff600160a81b8204811691600160a01b900416909192565b603554600090600160a81b900460ff161515600114610558576040805162461bcd60e51b815260206004820152601e60248201527f50726963652e6765743a207072696365206973206e6f74206163746976650000604482015290519081900360640190fd5b6000603454116105af576040805162461bcd60e51b815260206004820152601860248201527f50726963652e6765743a20696e76616c69642070726963650000000000000000604482015290519081900360640190fd5b603554600160a01b900460ff161561060e576040805162461bcd60e51b815260206004820152601d60248201527f50726963652e6765743a2070726963652068617320616e206572726f72000000604482015290519081900360640190fd5b5060345490565b6033546001600160a01b031681565b60345481565b6033546001600160a01b031633146106735760405162461bcd60e51b81526004018080602001828103825260578152602001806107b76057913960600191505060405180910390fd5b60006034556035805461ffff60a01b1916600160a01b179081905560408051338152306020820152600160a81b90920460ff16151582820152517fa1777e5c3fd82f369568fbd467e2cf87e4aa58a75621857a7714235c8ff9073b9181900360600190a1565b6033546001600160a01b031633146107225760405162461bcd60e51b81526004018080602001828103825260578152602001806107b76057913960600191505060405180910390fd5b603580546001600160a01b0319166001600160a01b0392909216919091179055565b6035546001600160a01b031681565b303b159056fe50726963652e61637469766174653a207072696365206973206e6f7420696e206120636f7272656374207374617465436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a656450726963652e6f6e6c794f776e65723a20546865206d6573736167652073656e646572206973206e6f7420666f756e64206f7220646f6573206e6f7420686176652073756666696369656e74207065726d697373696f6ea265627a7a723158209fcda362e8420c476214553f7f1b276b7184f93075e955608e48354312a3fcfd64736f6c63430005100032"

// DeployPrice deploys a new Ethereum contract, binding an instance of Price to it.
func DeployPrice(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Price, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Price{PriceCaller: PriceCaller{contract: contract}, PriceTransactor: PriceTransactor{contract: contract}, PriceFilterer: PriceFilterer{contract: contract}}, nil
}

// Price is an auto generated Go binding around an Ethereum contract.
type Price struct {
	PriceCaller     // Read-only binding to the contract
	PriceTransactor // Write-only binding to the contract
	PriceFilterer   // Log filterer for contract events
}

// PriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceSession struct {
	Contract     *Price            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceCallerSession struct {
	Contract *PriceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceTransactorSession struct {
	Contract     *PriceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceRaw struct {
	Contract *Price // Generic contract binding to access the raw methods on
}

// PriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceCallerRaw struct {
	Contract *PriceCaller // Generic read-only contract binding to access the raw methods on
}

// PriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceTransactorRaw struct {
	Contract *PriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrice creates a new instance of Price, bound to a specific deployed contract.
func NewPrice(address common.Address, backend bind.ContractBackend) (*Price, error) {
	contract, err := bindPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Price{PriceCaller: PriceCaller{contract: contract}, PriceTransactor: PriceTransactor{contract: contract}, PriceFilterer: PriceFilterer{contract: contract}}, nil
}

// NewPriceCaller creates a new read-only instance of Price, bound to a specific deployed contract.
func NewPriceCaller(address common.Address, caller bind.ContractCaller) (*PriceCaller, error) {
	contract, err := bindPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceCaller{contract: contract}, nil
}

// NewPriceTransactor creates a new write-only instance of Price, bound to a specific deployed contract.
func NewPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceTransactor, error) {
	contract, err := bindPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceTransactor{contract: contract}, nil
}

// NewPriceFilterer creates a new log filterer instance of Price, bound to a specific deployed contract.
func NewPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFilterer, error) {
	contract, err := bindPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFilterer{contract: contract}, nil
}

// bindPrice binds a generic wrapper to an already deployed contract.
func bindPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Price *PriceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Price.Contract.PriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Price *PriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Price.Contract.PriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Price *PriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Price.Contract.PriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Price *PriceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Price.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Price *PriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Price.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Price *PriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Price.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Price *PriceCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Price.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Price *PriceSession) Active() (bool, error) {
	return _Price.Contract.Active(&_Price.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Price *PriceCallerSession) Active() (bool, error) {
	return _Price.Contract.Active(&_Price.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Price *PriceCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Price.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Price *PriceSession) Get() (*big.Int, error) {
	return _Price.Contract.Get(&_Price.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Price *PriceCallerSession) Get() (*big.Int, error) {
	return _Price.Contract.Get(&_Price.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Price *PriceCaller) GetWithError(opts *bind.CallOpts) (*big.Int, bool, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Price.contract.Call(opts, out, "getWithError")
	return *ret0, *ret1, *ret2, err
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Price *PriceSession) GetWithError() (*big.Int, bool, bool, error) {
	return _Price.Contract.GetWithError(&_Price.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Price *PriceCallerSession) GetWithError() (*big.Int, bool, bool, error) {
	return _Price.Contract.GetWithError(&_Price.CallOpts)
}

// LagAddr is a free data retrieval call binding the contract method 0xfe0a55e1.
//
// Solidity: function lagAddr() constant returns(address)
func (_Price *PriceCaller) LagAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Price.contract.Call(opts, out, "lagAddr")
	return *ret0, err
}

// LagAddr is a free data retrieval call binding the contract method 0xfe0a55e1.
//
// Solidity: function lagAddr() constant returns(address)
func (_Price *PriceSession) LagAddr() (common.Address, error) {
	return _Price.Contract.LagAddr(&_Price.CallOpts)
}

// LagAddr is a free data retrieval call binding the contract method 0xfe0a55e1.
//
// Solidity: function lagAddr() constant returns(address)
func (_Price *PriceCallerSession) LagAddr() (common.Address, error) {
	return _Price.Contract.LagAddr(&_Price.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Price *PriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Price.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Price *PriceSession) Owner() (common.Address, error) {
	return _Price.Contract.Owner(&_Price.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Price *PriceCallerSession) Owner() (common.Address, error) {
	return _Price.Contract.Owner(&_Price.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Price *PriceCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Price.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Price *PriceSession) Price() (*big.Int, error) {
	return _Price.Contract.Price(&_Price.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Price *PriceCallerSession) Price() (*big.Int, error) {
	return _Price.Contract.Price(&_Price.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Price *PriceTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Price.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Price *PriceSession) Activate() (*types.Transaction, error) {
	return _Price.Contract.Activate(&_Price.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Price *PriceTransactorSession) Activate() (*types.Transaction, error) {
	return _Price.Contract.Activate(&_Price.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _lagAddr) returns()
func (_Price *PriceTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _lagAddr common.Address) (*types.Transaction, error) {
	return _Price.contract.Transact(opts, "initialize", _owner, _lagAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _lagAddr) returns()
func (_Price *PriceSession) Initialize(_owner common.Address, _lagAddr common.Address) (*types.Transaction, error) {
	return _Price.Contract.Initialize(&_Price.TransactOpts, _owner, _lagAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _lagAddr) returns()
func (_Price *PriceTransactorSession) Initialize(_owner common.Address, _lagAddr common.Address) (*types.Transaction, error) {
	return _Price.Contract.Initialize(&_Price.TransactOpts, _owner, _lagAddr)
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns()
func (_Price *PriceTransactor) Post(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Price.contract.Transact(opts, "post")
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns()
func (_Price *PriceSession) Post() (*types.Transaction, error) {
	return _Price.Contract.Post(&_Price.TransactOpts)
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns()
func (_Price *PriceTransactorSession) Post() (*types.Transaction, error) {
	return _Price.Contract.Post(&_Price.TransactOpts)
}

// SetLag is a paid mutator transaction binding the contract method 0xe34e3800.
//
// Solidity: function setLag(address newLagAddr) returns()
func (_Price *PriceTransactor) SetLag(opts *bind.TransactOpts, newLagAddr common.Address) (*types.Transaction, error) {
	return _Price.contract.Transact(opts, "setLag", newLagAddr)
}

// SetLag is a paid mutator transaction binding the contract method 0xe34e3800.
//
// Solidity: function setLag(address newLagAddr) returns()
func (_Price *PriceSession) SetLag(newLagAddr common.Address) (*types.Transaction, error) {
	return _Price.Contract.SetLag(&_Price.TransactOpts, newLagAddr)
}

// SetLag is a paid mutator transaction binding the contract method 0xe34e3800.
//
// Solidity: function setLag(address newLagAddr) returns()
func (_Price *PriceTransactorSession) SetLag(newLagAddr common.Address) (*types.Transaction, error) {
	return _Price.Contract.SetLag(&_Price.TransactOpts, newLagAddr)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Price *PriceTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Price.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Price *PriceSession) Void() (*types.Transaction, error) {
	return _Price.Contract.Void(&_Price.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Price *PriceTransactorSession) Void() (*types.Transaction, error) {
	return _Price.Contract.Void(&_Price.TransactOpts)
}

// PricePriceActivateIterator is returned from FilterPriceActivate and is used to iterate over the raw logs and unpacked data for PriceActivate events raised by the Price contract.
type PricePriceActivateIterator struct {
	Event *PricePriceActivate // Event containing the contract specifics and raw log

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
func (it *PricePriceActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PricePriceActivate)
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
		it.Event = new(PricePriceActivate)
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
func (it *PricePriceActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PricePriceActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PricePriceActivate represents a PriceActivate event raised by the Price contract.
type PricePriceActivate struct {
	Caller   common.Address
	Price    common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceActivate is a free log retrieval operation binding the contract event 0x2e219088454a765156309a15f374957f610cefb0bab9c32ecf58446b3d94b881.
//
// Solidity: event PriceActivate(address caller, address price, bool isActive)
func (_Price *PriceFilterer) FilterPriceActivate(opts *bind.FilterOpts) (*PricePriceActivateIterator, error) {

	logs, sub, err := _Price.contract.FilterLogs(opts, "PriceActivate")
	if err != nil {
		return nil, err
	}
	return &PricePriceActivateIterator{contract: _Price.contract, event: "PriceActivate", logs: logs, sub: sub}, nil
}

// WatchPriceActivate is a free log subscription operation binding the contract event 0x2e219088454a765156309a15f374957f610cefb0bab9c32ecf58446b3d94b881.
//
// Solidity: event PriceActivate(address caller, address price, bool isActive)
func (_Price *PriceFilterer) WatchPriceActivate(opts *bind.WatchOpts, sink chan<- *PricePriceActivate) (event.Subscription, error) {

	logs, sub, err := _Price.contract.WatchLogs(opts, "PriceActivate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PricePriceActivate)
				if err := _Price.contract.UnpackLog(event, "PriceActivate", log); err != nil {
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

// ParsePriceActivate is a log parse operation binding the contract event 0x2e219088454a765156309a15f374957f610cefb0bab9c32ecf58446b3d94b881.
//
// Solidity: event PriceActivate(address caller, address price, bool isActive)
func (_Price *PriceFilterer) ParsePriceActivate(log types.Log) (*PricePriceActivate, error) {
	event := new(PricePriceActivate)
	if err := _Price.contract.UnpackLog(event, "PriceActivate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PricePriceVoidIterator is returned from FilterPriceVoid and is used to iterate over the raw logs and unpacked data for PriceVoid events raised by the Price contract.
type PricePriceVoidIterator struct {
	Event *PricePriceVoid // Event containing the contract specifics and raw log

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
func (it *PricePriceVoidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PricePriceVoid)
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
		it.Event = new(PricePriceVoid)
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
func (it *PricePriceVoidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PricePriceVoidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PricePriceVoid represents a PriceVoid event raised by the Price contract.
type PricePriceVoid struct {
	Caller   common.Address
	Price    common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceVoid is a free log retrieval operation binding the contract event 0xa1777e5c3fd82f369568fbd467e2cf87e4aa58a75621857a7714235c8ff9073b.
//
// Solidity: event PriceVoid(address caller, address price, bool isActive)
func (_Price *PriceFilterer) FilterPriceVoid(opts *bind.FilterOpts) (*PricePriceVoidIterator, error) {

	logs, sub, err := _Price.contract.FilterLogs(opts, "PriceVoid")
	if err != nil {
		return nil, err
	}
	return &PricePriceVoidIterator{contract: _Price.contract, event: "PriceVoid", logs: logs, sub: sub}, nil
}

// WatchPriceVoid is a free log subscription operation binding the contract event 0xa1777e5c3fd82f369568fbd467e2cf87e4aa58a75621857a7714235c8ff9073b.
//
// Solidity: event PriceVoid(address caller, address price, bool isActive)
func (_Price *PriceFilterer) WatchPriceVoid(opts *bind.WatchOpts, sink chan<- *PricePriceVoid) (event.Subscription, error) {

	logs, sub, err := _Price.contract.WatchLogs(opts, "PriceVoid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PricePriceVoid)
				if err := _Price.contract.UnpackLog(event, "PriceVoid", log); err != nil {
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

// ParsePriceVoid is a log parse operation binding the contract event 0xa1777e5c3fd82f369568fbd467e2cf87e4aa58a75621857a7714235c8ff9073b.
//
// Solidity: event PriceVoid(address caller, address price, bool isActive)
func (_Price *PriceFilterer) ParsePriceVoid(log types.Log) (*PricePriceVoid, error) {
	event := new(PricePriceVoid)
	if err := _Price.contract.UnpackLog(event, "PriceVoid", log); err != nil {
		return nil, err
	}
	return event, nil
}
