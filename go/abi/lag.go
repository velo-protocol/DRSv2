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

// LagABI is the input ABI used to generate the binding from.
const LagABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lag\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"LagActivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lag\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"LagVoid\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"medianizerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumPeriodLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeLastUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_medianizerAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMedianizerAddr\",\"type\":\"address\"}],\"name\":\"setMedianizer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"newMinimumPeriod\",\"type\":\"int256\"}],\"name\":\"setMinimumPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isMinimumPeriodPass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"post\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextWithError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LagBin is the compiled bytecode used for deploying new contracts.
var LagBin = "0x608060405234801561001057600080fd5b50610c27806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063498b5828116100a25780638da5cb5b116100715780638da5cb5b1461020f578063ac4c25b214610217578063bb07f8dd1461021f578063e7baf8dd14610227578063f4bae7d51461024d5761010b565b8063498b5828146101ef57806351b42b00146101f757806367a493f9146101ff5780636d4ce63c146102075761010b565b806326f25c04116100de57806326f25c0414610162578063418819da1461018a578063457c85eb146101a7578063485cc955146101c15761010b565b806302fb0c5e1461011057806305c1f5021461012c5780630f15f4c0146101345780632482c6d11461013e575b600080fd5b610118610255565b604080519115158252519081900360200190f35b610118610265565b61013c6103c6565b005b61014661051a565b604080516001600160a01b039092168252519081900360200190f35b61016a610529565b604080519384529115156020840152151582820152519081900360600190f35b61013c600480360360208110156101a057600080fd5b5035610545565b6101af610614565b60408051918252519081900360200190f35b61013c600480360360408110156101d757600080fd5b506001600160a01b038135811691602001351661061a565b61016a610754565b61013c610770565b6101186107c8565b6101af6107f0565b61014661089a565b61013c6108a9565b6101af610976565b61013c6004803603602081101561023d57600080fd5b50356001600160a01b031661097c565b6101af6109e7565b603354600160a01b900460ff1681565b600061026f6107c8565b6102aa5760405162461bcd60e51b815260040180806020018281038252603d815260200180610bb6603d913960400191505060405180910390fd5b6000806000603460009054906101000a90046001600160a01b03166001600160a01b031663498b58286040518163ffffffff1660e01b815260040160606040518083038186803b1580156102fd57600080fd5b505afa158015610311573d6000803e3d6000fd5b505050506040513d606081101561032757600080fd5b508051602082015160409092015190945090925090508161035b5750506033805460ff60a01b1916905550600190506103c3565b801580156103695750600083115b156103bb576037805460395560388054603a805460ff8316151560ff1991821617909155604080518082019091528781528515156020909101819052938790551690911790556103b76109ed565b603b555b600193505050505b90565b6033546001600160a01b0316331461040f5760405162461bcd60e51b8152600401808060200182810382526055815260200180610a596055913960600191505060405180910390fd5b603354600160a01b900460ff161561046e576040805162461bcd60e51b815260206004820152601b60248201527f4c61672e61637469766174653a206c6167206973206163746976650000000000604482015290519081900360640190fd5b60395415801590610480575060375415155b6104bb5760405162461bcd60e51b815260040180806020018281038252602d815260200180610b89602d913960400191505060405180910390fd5b6033805460ff60a01b1916600160a01b90811791829055604080513381523060208201529190920460ff1615158183015290517f7f62a36327a11de3be66fdb622fb3631e71b9145fd53295d111869c49f8bbe389181900360600190a1565b6034546001600160a01b031681565b60375460335460385460ff600160a01b90920482169116909192565b6033546001600160a01b0316331461058e5760405162461bcd60e51b8152600401808060200182810382526055815260200180610a596055913960600191505060405180910390fd5b60008112156105ce5760405162461bcd60e51b8152600401808060200182810382526041815260200180610ad06041913960600191505060405180910390fd5b60365481111561060f5760405162461bcd60e51b815260040180806020018281038252604a815260200180610b11604a913960600191505060405180910390fd5b603555565b60355481565b600054610100900460ff168061063357506106336109f1565b80610641575060005460ff16155b61067c5760405162461bcd60e51b815260040180806020018281038252602e815260200180610b5b602e913960400191505060405180910390fd5b600054610100900460ff161580156106a7576000805460ff1961ff0019909116610100171660011790555b603480546001600160a01b038085166001600160a01b0319928316179092556033805461038460355562278d0060365560ff60a01b1993871692169190911791909116600160a01b1790556040805180820182526000808252600160209283018190526039829055603a805460ff19908116831790915584518086019095528285529390920182905260375560388054909216179055801561074f576000805461ff00191690555b505050565b603954603354603a5460ff600160a01b90920482169116909192565b6033546001600160a01b031633146107b95760405162461bcd60e51b8152600401808060200182810382526055815260200180610a596055913960600191505060405180910390fd5b6033805460ff60a01b19169055565b60006107e1603554603b546109f790919063ffffffff16565b6107e96109ed565b1015905090565b603a5460009060ff16156108355760405162461bcd60e51b8152600401808060200182810382526022815260200180610aae6022913960400191505060405180910390fd5b603354600160a01b900460ff16610893576040805162461bcd60e51b815260206004820152601860248201527f4c61672e6765743a206c616720697320696e6163746976650000000000000000604482015290519081900360640190fd5b5060395490565b6033546001600160a01b031681565b6033546001600160a01b031633146108f25760405162461bcd60e51b8152600401808060200182810382526055815260200180610a596055913960600191505060405180910390fd5b604080518082018252600080825260016020928301819052603991909155603a805460ff191690911790556033805460ff60a01b19169081905582513381523092810192909252600160a01b900460ff1615158183015290517e7c5635357dbfc3e0e64195a2d3cb69d46a53fe985e127be4d6aa371bc6809e9181900360600190a1565b603b5481565b6033546001600160a01b031633146109c55760405162461bcd60e51b8152600401808060200182810382526055815260200180610a596055913960600191505060405180910390fd5b603480546001600160a01b0319166001600160a01b0392909216919091179055565b60365481565b4290565b303b1590565b600082820183811015610a51576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe4c61672e6f6e6c794f776e65723a20546865206d6573736167652073656e646572206973206e6f7420666f756e64206f7220646f6573206e6f7420686176652073756666696369656e74207065726d697373696f6e4c61672e6765743a2063757272656e7450726963652068617320616e206572726f724c61672e7365744d696e696d756d506572696f643a206d696e696d756d506572696f642076616c7565206d757374206e6f74206265206c657373207468616e20304c61672e7365744d696e696d756d506572696f643a206d696e696d756d506572696f642076616c7565206d757374206e6f742062652067726561746572207468616e2032353932303030436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65644c61672e61637469766174653a207072696365206973206e6f7420696e206120636f72726563742073746174654c61672e706f73743a206d696e696d756d20706572696f6420666f722074686520706f73742066756e6374696f6e20686173206e6f7420706173736564a265627a7a723158200d8d06c49e93ec3a9cc0e85fb975c80879884c94a21eb309eaba742a2f30d6f464736f6c63430005100032"

// DeployLag deploys a new Ethereum contract, binding an instance of Lag to it.
func DeployLag(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lag, error) {
	parsed, err := abi.JSON(strings.NewReader(LagABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LagBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lag{LagCaller: LagCaller{contract: contract}, LagTransactor: LagTransactor{contract: contract}, LagFilterer: LagFilterer{contract: contract}}, nil
}

// Lag is an auto generated Go binding around an Ethereum contract.
type Lag struct {
	LagCaller     // Read-only binding to the contract
	LagTransactor // Write-only binding to the contract
	LagFilterer   // Log filterer for contract events
}

// LagCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagSession struct {
	Contract     *Lag              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LagCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagCallerSession struct {
	Contract *LagCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LagTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagTransactorSession struct {
	Contract     *LagTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LagRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagRaw struct {
	Contract *Lag // Generic contract binding to access the raw methods on
}

// LagCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagCallerRaw struct {
	Contract *LagCaller // Generic read-only contract binding to access the raw methods on
}

// LagTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagTransactorRaw struct {
	Contract *LagTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLag creates a new instance of Lag, bound to a specific deployed contract.
func NewLag(address common.Address, backend bind.ContractBackend) (*Lag, error) {
	contract, err := bindLag(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lag{LagCaller: LagCaller{contract: contract}, LagTransactor: LagTransactor{contract: contract}, LagFilterer: LagFilterer{contract: contract}}, nil
}

// NewLagCaller creates a new read-only instance of Lag, bound to a specific deployed contract.
func NewLagCaller(address common.Address, caller bind.ContractCaller) (*LagCaller, error) {
	contract, err := bindLag(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagCaller{contract: contract}, nil
}

// NewLagTransactor creates a new write-only instance of Lag, bound to a specific deployed contract.
func NewLagTransactor(address common.Address, transactor bind.ContractTransactor) (*LagTransactor, error) {
	contract, err := bindLag(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagTransactor{contract: contract}, nil
}

// NewLagFilterer creates a new log filterer instance of Lag, bound to a specific deployed contract.
func NewLagFilterer(address common.Address, filterer bind.ContractFilterer) (*LagFilterer, error) {
	contract, err := bindLag(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagFilterer{contract: contract}, nil
}

// bindLag binds a generic wrapper to an already deployed contract.
func bindLag(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LagABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lag *LagRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lag.Contract.LagCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lag *LagRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lag.Contract.LagTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lag *LagRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lag.Contract.LagTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lag *LagCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lag.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lag *LagTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lag.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lag *LagTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lag.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Lag *LagCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Lag *LagSession) Active() (bool, error) {
	return _Lag.Contract.Active(&_Lag.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Lag *LagCallerSession) Active() (bool, error) {
	return _Lag.Contract.Active(&_Lag.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Lag *LagCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Lag *LagSession) Get() (*big.Int, error) {
	return _Lag.Contract.Get(&_Lag.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Lag *LagCallerSession) Get() (*big.Int, error) {
	return _Lag.Contract.Get(&_Lag.CallOpts)
}

// GetNextWithError is a free data retrieval call binding the contract method 0x26f25c04.
//
// Solidity: function getNextWithError() constant returns(uint256, bool, bool)
func (_Lag *LagCaller) GetNextWithError(opts *bind.CallOpts) (*big.Int, bool, bool, error) {
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
	err := _Lag.contract.Call(opts, out, "getNextWithError")
	return *ret0, *ret1, *ret2, err
}

// GetNextWithError is a free data retrieval call binding the contract method 0x26f25c04.
//
// Solidity: function getNextWithError() constant returns(uint256, bool, bool)
func (_Lag *LagSession) GetNextWithError() (*big.Int, bool, bool, error) {
	return _Lag.Contract.GetNextWithError(&_Lag.CallOpts)
}

// GetNextWithError is a free data retrieval call binding the contract method 0x26f25c04.
//
// Solidity: function getNextWithError() constant returns(uint256, bool, bool)
func (_Lag *LagCallerSession) GetNextWithError() (*big.Int, bool, bool, error) {
	return _Lag.Contract.GetNextWithError(&_Lag.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Lag *LagCaller) GetWithError(opts *bind.CallOpts) (*big.Int, bool, bool, error) {
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
	err := _Lag.contract.Call(opts, out, "getWithError")
	return *ret0, *ret1, *ret2, err
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Lag *LagSession) GetWithError() (*big.Int, bool, bool, error) {
	return _Lag.Contract.GetWithError(&_Lag.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Lag *LagCallerSession) GetWithError() (*big.Int, bool, bool, error) {
	return _Lag.Contract.GetWithError(&_Lag.CallOpts)
}

// IsMinimumPeriodPass is a free data retrieval call binding the contract method 0x67a493f9.
//
// Solidity: function isMinimumPeriodPass() constant returns(bool)
func (_Lag *LagCaller) IsMinimumPeriodPass(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "isMinimumPeriodPass")
	return *ret0, err
}

// IsMinimumPeriodPass is a free data retrieval call binding the contract method 0x67a493f9.
//
// Solidity: function isMinimumPeriodPass() constant returns(bool)
func (_Lag *LagSession) IsMinimumPeriodPass() (bool, error) {
	return _Lag.Contract.IsMinimumPeriodPass(&_Lag.CallOpts)
}

// IsMinimumPeriodPass is a free data retrieval call binding the contract method 0x67a493f9.
//
// Solidity: function isMinimumPeriodPass() constant returns(bool)
func (_Lag *LagCallerSession) IsMinimumPeriodPass() (bool, error) {
	return _Lag.Contract.IsMinimumPeriodPass(&_Lag.CallOpts)
}

// MedianizerAddr is a free data retrieval call binding the contract method 0x2482c6d1.
//
// Solidity: function medianizerAddr() constant returns(address)
func (_Lag *LagCaller) MedianizerAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "medianizerAddr")
	return *ret0, err
}

// MedianizerAddr is a free data retrieval call binding the contract method 0x2482c6d1.
//
// Solidity: function medianizerAddr() constant returns(address)
func (_Lag *LagSession) MedianizerAddr() (common.Address, error) {
	return _Lag.Contract.MedianizerAddr(&_Lag.CallOpts)
}

// MedianizerAddr is a free data retrieval call binding the contract method 0x2482c6d1.
//
// Solidity: function medianizerAddr() constant returns(address)
func (_Lag *LagCallerSession) MedianizerAddr() (common.Address, error) {
	return _Lag.Contract.MedianizerAddr(&_Lag.CallOpts)
}

// MinimumPeriod is a free data retrieval call binding the contract method 0x457c85eb.
//
// Solidity: function minimumPeriod() constant returns(uint256)
func (_Lag *LagCaller) MinimumPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "minimumPeriod")
	return *ret0, err
}

// MinimumPeriod is a free data retrieval call binding the contract method 0x457c85eb.
//
// Solidity: function minimumPeriod() constant returns(uint256)
func (_Lag *LagSession) MinimumPeriod() (*big.Int, error) {
	return _Lag.Contract.MinimumPeriod(&_Lag.CallOpts)
}

// MinimumPeriod is a free data retrieval call binding the contract method 0x457c85eb.
//
// Solidity: function minimumPeriod() constant returns(uint256)
func (_Lag *LagCallerSession) MinimumPeriod() (*big.Int, error) {
	return _Lag.Contract.MinimumPeriod(&_Lag.CallOpts)
}

// MinimumPeriodLimit is a free data retrieval call binding the contract method 0xf4bae7d5.
//
// Solidity: function minimumPeriodLimit() constant returns(uint256)
func (_Lag *LagCaller) MinimumPeriodLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "minimumPeriodLimit")
	return *ret0, err
}

// MinimumPeriodLimit is a free data retrieval call binding the contract method 0xf4bae7d5.
//
// Solidity: function minimumPeriodLimit() constant returns(uint256)
func (_Lag *LagSession) MinimumPeriodLimit() (*big.Int, error) {
	return _Lag.Contract.MinimumPeriodLimit(&_Lag.CallOpts)
}

// MinimumPeriodLimit is a free data retrieval call binding the contract method 0xf4bae7d5.
//
// Solidity: function minimumPeriodLimit() constant returns(uint256)
func (_Lag *LagCallerSession) MinimumPeriodLimit() (*big.Int, error) {
	return _Lag.Contract.MinimumPeriodLimit(&_Lag.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lag *LagCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lag *LagSession) Owner() (common.Address, error) {
	return _Lag.Contract.Owner(&_Lag.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lag *LagCallerSession) Owner() (common.Address, error) {
	return _Lag.Contract.Owner(&_Lag.CallOpts)
}

// TimeLastUpdate is a free data retrieval call binding the contract method 0xbb07f8dd.
//
// Solidity: function timeLastUpdate() constant returns(uint256)
func (_Lag *LagCaller) TimeLastUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lag.contract.Call(opts, out, "timeLastUpdate")
	return *ret0, err
}

// TimeLastUpdate is a free data retrieval call binding the contract method 0xbb07f8dd.
//
// Solidity: function timeLastUpdate() constant returns(uint256)
func (_Lag *LagSession) TimeLastUpdate() (*big.Int, error) {
	return _Lag.Contract.TimeLastUpdate(&_Lag.CallOpts)
}

// TimeLastUpdate is a free data retrieval call binding the contract method 0xbb07f8dd.
//
// Solidity: function timeLastUpdate() constant returns(uint256)
func (_Lag *LagCallerSession) TimeLastUpdate() (*big.Int, error) {
	return _Lag.Contract.TimeLastUpdate(&_Lag.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Lag *LagTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Lag *LagSession) Activate() (*types.Transaction, error) {
	return _Lag.Contract.Activate(&_Lag.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Lag *LagTransactorSession) Activate() (*types.Transaction, error) {
	return _Lag.Contract.Activate(&_Lag.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Lag *LagTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Lag *LagSession) Deactivate() (*types.Transaction, error) {
	return _Lag.Contract.Deactivate(&_Lag.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Lag *LagTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Lag.Contract.Deactivate(&_Lag.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _medianizerAddr) returns()
func (_Lag *LagTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _medianizerAddr common.Address) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "initialize", _owner, _medianizerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _medianizerAddr) returns()
func (_Lag *LagSession) Initialize(_owner common.Address, _medianizerAddr common.Address) (*types.Transaction, error) {
	return _Lag.Contract.Initialize(&_Lag.TransactOpts, _owner, _medianizerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _medianizerAddr) returns()
func (_Lag *LagTransactorSession) Initialize(_owner common.Address, _medianizerAddr common.Address) (*types.Transaction, error) {
	return _Lag.Contract.Initialize(&_Lag.TransactOpts, _owner, _medianizerAddr)
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns(bool)
func (_Lag *LagTransactor) Post(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "post")
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns(bool)
func (_Lag *LagSession) Post() (*types.Transaction, error) {
	return _Lag.Contract.Post(&_Lag.TransactOpts)
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns(bool)
func (_Lag *LagTransactorSession) Post() (*types.Transaction, error) {
	return _Lag.Contract.Post(&_Lag.TransactOpts)
}

// SetMedianizer is a paid mutator transaction binding the contract method 0xe7baf8dd.
//
// Solidity: function setMedianizer(address newMedianizerAddr) returns()
func (_Lag *LagTransactor) SetMedianizer(opts *bind.TransactOpts, newMedianizerAddr common.Address) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "setMedianizer", newMedianizerAddr)
}

// SetMedianizer is a paid mutator transaction binding the contract method 0xe7baf8dd.
//
// Solidity: function setMedianizer(address newMedianizerAddr) returns()
func (_Lag *LagSession) SetMedianizer(newMedianizerAddr common.Address) (*types.Transaction, error) {
	return _Lag.Contract.SetMedianizer(&_Lag.TransactOpts, newMedianizerAddr)
}

// SetMedianizer is a paid mutator transaction binding the contract method 0xe7baf8dd.
//
// Solidity: function setMedianizer(address newMedianizerAddr) returns()
func (_Lag *LagTransactorSession) SetMedianizer(newMedianizerAddr common.Address) (*types.Transaction, error) {
	return _Lag.Contract.SetMedianizer(&_Lag.TransactOpts, newMedianizerAddr)
}

// SetMinimumPeriod is a paid mutator transaction binding the contract method 0x418819da.
//
// Solidity: function setMinimumPeriod(int256 newMinimumPeriod) returns()
func (_Lag *LagTransactor) SetMinimumPeriod(opts *bind.TransactOpts, newMinimumPeriod *big.Int) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "setMinimumPeriod", newMinimumPeriod)
}

// SetMinimumPeriod is a paid mutator transaction binding the contract method 0x418819da.
//
// Solidity: function setMinimumPeriod(int256 newMinimumPeriod) returns()
func (_Lag *LagSession) SetMinimumPeriod(newMinimumPeriod *big.Int) (*types.Transaction, error) {
	return _Lag.Contract.SetMinimumPeriod(&_Lag.TransactOpts, newMinimumPeriod)
}

// SetMinimumPeriod is a paid mutator transaction binding the contract method 0x418819da.
//
// Solidity: function setMinimumPeriod(int256 newMinimumPeriod) returns()
func (_Lag *LagTransactorSession) SetMinimumPeriod(newMinimumPeriod *big.Int) (*types.Transaction, error) {
	return _Lag.Contract.SetMinimumPeriod(&_Lag.TransactOpts, newMinimumPeriod)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Lag *LagTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lag.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Lag *LagSession) Void() (*types.Transaction, error) {
	return _Lag.Contract.Void(&_Lag.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Lag *LagTransactorSession) Void() (*types.Transaction, error) {
	return _Lag.Contract.Void(&_Lag.TransactOpts)
}

// LagLagActivateIterator is returned from FilterLagActivate and is used to iterate over the raw logs and unpacked data for LagActivate events raised by the Lag contract.
type LagLagActivateIterator struct {
	Event *LagLagActivate // Event containing the contract specifics and raw log

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
func (it *LagLagActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagLagActivate)
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
		it.Event = new(LagLagActivate)
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
func (it *LagLagActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagLagActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagLagActivate represents a LagActivate event raised by the Lag contract.
type LagLagActivate struct {
	Caller   common.Address
	Lag      common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLagActivate is a free log retrieval operation binding the contract event 0x7f62a36327a11de3be66fdb622fb3631e71b9145fd53295d111869c49f8bbe38.
//
// Solidity: event LagActivate(address caller, address lag, bool isActive)
func (_Lag *LagFilterer) FilterLagActivate(opts *bind.FilterOpts) (*LagLagActivateIterator, error) {

	logs, sub, err := _Lag.contract.FilterLogs(opts, "LagActivate")
	if err != nil {
		return nil, err
	}
	return &LagLagActivateIterator{contract: _Lag.contract, event: "LagActivate", logs: logs, sub: sub}, nil
}

// WatchLagActivate is a free log subscription operation binding the contract event 0x7f62a36327a11de3be66fdb622fb3631e71b9145fd53295d111869c49f8bbe38.
//
// Solidity: event LagActivate(address caller, address lag, bool isActive)
func (_Lag *LagFilterer) WatchLagActivate(opts *bind.WatchOpts, sink chan<- *LagLagActivate) (event.Subscription, error) {

	logs, sub, err := _Lag.contract.WatchLogs(opts, "LagActivate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagLagActivate)
				if err := _Lag.contract.UnpackLog(event, "LagActivate", log); err != nil {
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

// ParseLagActivate is a log parse operation binding the contract event 0x7f62a36327a11de3be66fdb622fb3631e71b9145fd53295d111869c49f8bbe38.
//
// Solidity: event LagActivate(address caller, address lag, bool isActive)
func (_Lag *LagFilterer) ParseLagActivate(log types.Log) (*LagLagActivate, error) {
	event := new(LagLagActivate)
	if err := _Lag.contract.UnpackLog(event, "LagActivate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LagLagVoidIterator is returned from FilterLagVoid and is used to iterate over the raw logs and unpacked data for LagVoid events raised by the Lag contract.
type LagLagVoidIterator struct {
	Event *LagLagVoid // Event containing the contract specifics and raw log

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
func (it *LagLagVoidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagLagVoid)
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
		it.Event = new(LagLagVoid)
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
func (it *LagLagVoidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagLagVoidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagLagVoid represents a LagVoid event raised by the Lag contract.
type LagLagVoid struct {
	Caller   common.Address
	Lag      common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLagVoid is a free log retrieval operation binding the contract event 0x007c5635357dbfc3e0e64195a2d3cb69d46a53fe985e127be4d6aa371bc6809e.
//
// Solidity: event LagVoid(address caller, address lag, bool isActive)
func (_Lag *LagFilterer) FilterLagVoid(opts *bind.FilterOpts) (*LagLagVoidIterator, error) {

	logs, sub, err := _Lag.contract.FilterLogs(opts, "LagVoid")
	if err != nil {
		return nil, err
	}
	return &LagLagVoidIterator{contract: _Lag.contract, event: "LagVoid", logs: logs, sub: sub}, nil
}

// WatchLagVoid is a free log subscription operation binding the contract event 0x007c5635357dbfc3e0e64195a2d3cb69d46a53fe985e127be4d6aa371bc6809e.
//
// Solidity: event LagVoid(address caller, address lag, bool isActive)
func (_Lag *LagFilterer) WatchLagVoid(opts *bind.WatchOpts, sink chan<- *LagLagVoid) (event.Subscription, error) {

	logs, sub, err := _Lag.contract.WatchLogs(opts, "LagVoid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagLagVoid)
				if err := _Lag.contract.UnpackLog(event, "LagVoid", log); err != nil {
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

// ParseLagVoid is a log parse operation binding the contract event 0x007c5635357dbfc3e0e64195a2d3cb69d46a53fe985e127be4d6aa371bc6809e.
//
// Solidity: event LagVoid(address caller, address lag, bool isActive)
func (_Lag *LagFilterer) ParseLagVoid(log types.Log) (*LagLagVoid, error) {
	event := new(LagLagVoid)
	if err := _Lag.contract.UnpackLog(event, "LagVoid", log); err != nil {
		return nil, err
	}
	return event, nil
}
