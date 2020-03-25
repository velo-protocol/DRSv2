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

// PriceFeedersABI is the input ABI used to generate the binding from.
const PriceFeedersABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetFiat\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"medianPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"priceFeeders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"setAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"}],\"name\":\"addAssetFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getAssetFiat\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAssetFiat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"priceFeederAddr\",\"type\":\"address\"}],\"name\":\"addPriceFeeder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"}],\"name\":\"getPriceFeeders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pfId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePriceFeeder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"}],\"name\":\"getMedianPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceFeedersBin is the compiled bytecode used for deploying new contracts.
var PriceFeedersBin = "0x60806040526100266100186001600160e01b0361002b16565b6001600160e01b0361002f16565b6101a3565b3390565b61004781600061007e60201b610e7b1790919060201c565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b61009182826001600160e01b0361012216565b156100fd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216610183576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806112b06022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6110fe806101b26000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80634c5a628c1161009757806389dbe24b1161006657806389dbe24b14610315578063bb5f747b14610332578063c966a39f14610358578063d8d4f09a1461039757610100565b80634c5a628c146102955780635aeb8ea91461029d57806360846bc6146102d25780637362d9c8146102ef57610100565b806321f73ab0116100d357806321f73ab01461020057806333e954571461021d578063367ae0171461024057806343cb01011461027257610100565b806315b9de67146101055780631610c6a61461013c578063213b9d9e1461016a57806321dad76214610193575b600080fd5b6101286004803603604081101561011b57600080fd5b50803590602001356103b4565b604080519115158252519081900360200190f35b6101686004803603604081101561015257600080fd5b50803590602001356001600160a01b03166104e7565b005b6101686004803603606081101561018057600080fd5b508035906020810135906040013561055b565b6101b0600480360360208110156101a957600080fd5b50356106cb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101ec5781810151838201526020016101d4565b505050509050019250505060405180910390f35b6101b06004803603602081101561021657600080fd5b503561072d565b6101286004803603604081101561023357600080fd5b5080359060200135610798565b6101686004803603606081101561025657600080fd5b50803590602081013590604001356001600160a01b0316610893565b6101686004803603604081101561028857600080fd5b5080359060200135610a34565b610168610b40565b6102c0600480360360408110156102b357600080fd5b5080359060200135610b52565b60408051918252519081900360200190f35b6102c0600480360360208110156102e857600080fd5b5035610b80565b6101686004803603602081101561030557600080fd5b50356001600160a01b0316610b92565b6102c06004803603602081101561032b57600080fd5b5035610be4565b6101286004803603602081101561034857600080fd5b50356001600160a01b0316610bf6565b61037b6004803603604081101561036e57600080fd5b5080359060200135610c08565b604080516001600160a01b039092168252519081900360200190f35b6102c0600480360360208110156103ad57600080fd5b5035610c3d565b60006103c66103c1610c4f565b610bf6565b6104015760405162461bcd60e51b81526004018080602001828103825260408152602001806110086040913960400191505060405180910390fd5b600083815260036020526040902054821061041e575060006104e1565b815b600084815260036020526040902054600019018110156104bb57600084815260036020526040902080546001830190811061045757fe5b60009182526020808320909101548683526003909152604090912080546001600160a01b03909216918390811061048a57fe5b600091825260209091200180546001600160a01b0319166001600160a01b0392909216919091179055600101610420565b5060008381526003602052604090208054906104db906000198301610efc565b50600190505b92915050565b6104f26103c1610c4f565b61052d5760405162461bcd60e51b81526004018080602001828103825260408152602001806110086040913960400191505060405180910390fd5b60009182526001602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b60408051602080820186905281830185905282518083038401815260609092019092528051910120600061058f8233610c54565b9050600061059d8686610cc6565b6000878152600160205260409020549091506001600160a01b03166105f35760405162461bcd60e51b81526004018080602001828103825260298152602001806110a16029913960400191505060405180910390fd5b8060001914156106345760405162461bcd60e51b8152600401808060200182810382526028815260200180610fbe6028913960400191505060405180910390fd5b8160001914156106755760405162461bcd60e51b8152600401808060200182810382526035815260200180610f446035913960400191505060405180910390fd5b5050604080516020808201969096528082019490945233606090811b9085015280516054818603018152607490940181528351938501939093206000908152600485528381208390559081526005909352912055565b60008181526002602090815260409182902080548351818402810184019094528084526060939283018282801561072157602002820191906000526020600020905b81548152602001906001019080831161070d575b50505050509050919050565b60008181526003602090815260409182902080548351818402810184019094528084526060939283018282801561072157602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161076f5750505050509050919050565b60006107a56103c1610c4f565b6107e05760405162461bcd60e51b81526004018080602001828103825260408152602001806110086040913960400191505060405180910390fd5b60008381526002602052604090205482106107fd575060006104e1565b815b6000848152600260205260409020546000190181101561087357600084815260026020526040902080546001830190811061083657fe5b906000526020600020015460026000868152602001908152602001600020828154811061085f57fe5b6000918252602090912001556001016107ff565b5060008381526002602052604090208054906104db906000198301610efc565b61089e6103c1610c4f565b6108d95760405162461bcd60e51b81526004018080602001828103825260408152602001806110086040913960400191505060405180910390fd5b60408051602080820186905281830185905282518083038401815260609092019092528051910120600061090d8585610cc6565b9050600061091b8385610c54565b6000878152600160205260409020549091506001600160a01b03166109715760405162461bcd60e51b81526004018080602001828103825260298152602001806110a16029913960400191505060405180910390fd5b8160001914156109b25760405162461bcd60e51b8152600401808060200182810382526028815260200180610fbe6028913960400191505060405180910390fd5b80600019146109f25760405162461bcd60e51b815260040180806020018281038252602c815260200180611075602c913960400191505060405180910390fd5b505060009081526003602090815260408220805460018101825590835291200180546001600160a01b0319166001600160a01b03929092169190911790555050565b610a3f6103c1610c4f565b610a7a5760405162461bcd60e51b81526004018080602001828103825260408152602001806110086040913960400191505060405180910390fd5b6000610a868383610cc6565b6000848152600160205260409020549091506001600160a01b0316610adc5760405162461bcd60e51b815260040180806020018281038252602d815260200180611048602d913960400191505060405180910390fd5b8060001914610b1c5760405162461bcd60e51b8152600401808060200182810382526024815260200180610f9a6024913960400191505060405180910390fd5b50600091825260026020908152604083208054600181018255908452922090910155565b610b50610b4b610c4f565b610d1d565b565b60026020528160005260406000208181548110610b6b57fe5b90600052602060002001600091509150505481565b60046020526000908152604090205481565b610b9d6103c1610c4f565b610bd85760405162461bcd60e51b81526004018080602001828103825260408152602001806110086040913960400191505060405180910390fd5b610be181610d65565b50565b60009081526005602052604090205490565b60006104e1818363ffffffff610dad16565b60036020528160005260406000208181548110610c2157fe5b6000918252602090912001546001600160a01b03169150829050565b60056020526000908152604090205481565b335b90565b6000600019815b600085815260036020526040902054811015610cbe57600085815260036020526040902080546001600160a01b038616919083908110610c9757fe5b6000918252602090912001546001600160a01b03161415610cb6578091505b600101610c5b565b509392505050565b6000600019815b600085815260026020526040902054811015610cbe576000858152600260205260409020805485919083908110610d0057fe5b90600052602060002001541415610d15578091505b600101610ccd565b610d2e60008263ffffffff610e1416565b6040516001600160a01b038216907f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16590600090a250565b610d7660008263ffffffff610e7b16565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b60006001600160a01b038216610df45760405162461bcd60e51b8152600401808060200182810382526022815260200180610fe66022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610e1e8282610dad565b610e595760405162461bcd60e51b8152600401808060200182810382526021815260200180610f796021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b610e858282610dad565b15610ed7576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b815481835581811115610f2057600083815260209020610f20918101908301610f25565b505050565b610c5191905b80821115610f3f5760008155600101610f2b565b509056fe6d73672e73656e64657220646964206e6f74206861766520616e20617574686f7269747920746f2073657420746865207072696365526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65617373657420686173206265656e206c696e6b6564207769746820746869732066696174617373657420686173206e6f74206265656e206c696e6b6564207769746820746869732066696174526f6c65733a206163636f756e7420697320746865207a65726f206164647265737357686974656c69737441646d696e526f6c653a2063616c6c657220646f6573206e6f742068617665207468652057686974656c69737441646d696e20726f6c656173736574436f646520686173206e6f74206265656e20616464656420746f207468652061737365744c697374746869732070726963652066656564657220686173206265656e20616464656420746f20746865206c6973746173736574436f646520686173206e6f74206265656e20616464656420746f2061737365744c697374a265627a7a723158209a57035fe9997374980ddc0e0c124984582ad56e297fd9b07bad82e3c2036bdc64736f6c634300050c0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployPriceFeeders deploys a new Ethereum contract, binding an instance of PriceFeeders to it.
func DeployPriceFeeders(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeeders, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeedersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceFeedersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeeders{PriceFeedersCaller: PriceFeedersCaller{contract: contract}, PriceFeedersTransactor: PriceFeedersTransactor{contract: contract}, PriceFeedersFilterer: PriceFeedersFilterer{contract: contract}}, nil
}

// PriceFeeders is an auto generated Go binding around an Ethereum contract.
type PriceFeeders struct {
	PriceFeedersCaller     // Read-only binding to the contract
	PriceFeedersTransactor // Write-only binding to the contract
	PriceFeedersFilterer   // Log filterer for contract events
}

// PriceFeedersCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedersSession struct {
	Contract     *PriceFeeders     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedersCallerSession struct {
	Contract *PriceFeedersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PriceFeedersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedersTransactorSession struct {
	Contract     *PriceFeedersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PriceFeedersRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedersRaw struct {
	Contract *PriceFeeders // Generic contract binding to access the raw methods on
}

// PriceFeedersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedersCallerRaw struct {
	Contract *PriceFeedersCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedersTransactorRaw struct {
	Contract *PriceFeedersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeeders creates a new instance of PriceFeeders, bound to a specific deployed contract.
func NewPriceFeeders(address common.Address, backend bind.ContractBackend) (*PriceFeeders, error) {
	contract, err := bindPriceFeeders(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeeders{PriceFeedersCaller: PriceFeedersCaller{contract: contract}, PriceFeedersTransactor: PriceFeedersTransactor{contract: contract}, PriceFeedersFilterer: PriceFeedersFilterer{contract: contract}}, nil
}

// NewPriceFeedersCaller creates a new read-only instance of PriceFeeders, bound to a specific deployed contract.
func NewPriceFeedersCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedersCaller, error) {
	contract, err := bindPriceFeeders(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedersCaller{contract: contract}, nil
}

// NewPriceFeedersTransactor creates a new write-only instance of PriceFeeders, bound to a specific deployed contract.
func NewPriceFeedersTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedersTransactor, error) {
	contract, err := bindPriceFeeders(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedersTransactor{contract: contract}, nil
}

// NewPriceFeedersFilterer creates a new log filterer instance of PriceFeeders, bound to a specific deployed contract.
func NewPriceFeedersFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedersFilterer, error) {
	contract, err := bindPriceFeeders(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedersFilterer{contract: contract}, nil
}

// bindPriceFeeders binds a generic wrapper to an already deployed contract.
func bindPriceFeeders(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeedersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeeders *PriceFeedersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeeders.Contract.PriceFeedersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeeders *PriceFeedersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeders.Contract.PriceFeedersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeeders *PriceFeedersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeeders.Contract.PriceFeedersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeeders *PriceFeedersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeeders.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeeders *PriceFeedersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeders.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeeders *PriceFeedersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeeders.Contract.contract.Transact(opts, method, params...)
}

// AssetFiat is a free data retrieval call binding the contract method 0x5aeb8ea9.
//
// Solidity: function assetFiat(bytes32 , uint256 ) constant returns(bytes32)
func (_PriceFeeders *PriceFeedersCaller) AssetFiat(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "assetFiat", arg0, arg1)
	return *ret0, err
}

// AssetFiat is a free data retrieval call binding the contract method 0x5aeb8ea9.
//
// Solidity: function assetFiat(bytes32 , uint256 ) constant returns(bytes32)
func (_PriceFeeders *PriceFeedersSession) AssetFiat(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _PriceFeeders.Contract.AssetFiat(&_PriceFeeders.CallOpts, arg0, arg1)
}

// AssetFiat is a free data retrieval call binding the contract method 0x5aeb8ea9.
//
// Solidity: function assetFiat(bytes32 , uint256 ) constant returns(bytes32)
func (_PriceFeeders *PriceFeedersCallerSession) AssetFiat(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _PriceFeeders.Contract.AssetFiat(&_PriceFeeders.CallOpts, arg0, arg1)
}

// GetAssetFiat is a free data retrieval call binding the contract method 0x21dad762.
//
// Solidity: function getAssetFiat(bytes32 assetCode) constant returns(bytes32[])
func (_PriceFeeders *PriceFeedersCaller) GetAssetFiat(opts *bind.CallOpts, assetCode [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "getAssetFiat", assetCode)
	return *ret0, err
}

// GetAssetFiat is a free data retrieval call binding the contract method 0x21dad762.
//
// Solidity: function getAssetFiat(bytes32 assetCode) constant returns(bytes32[])
func (_PriceFeeders *PriceFeedersSession) GetAssetFiat(assetCode [32]byte) ([][32]byte, error) {
	return _PriceFeeders.Contract.GetAssetFiat(&_PriceFeeders.CallOpts, assetCode)
}

// GetAssetFiat is a free data retrieval call binding the contract method 0x21dad762.
//
// Solidity: function getAssetFiat(bytes32 assetCode) constant returns(bytes32[])
func (_PriceFeeders *PriceFeedersCallerSession) GetAssetFiat(assetCode [32]byte) ([][32]byte, error) {
	return _PriceFeeders.Contract.GetAssetFiat(&_PriceFeeders.CallOpts, assetCode)
}

// GetMedianPrice is a free data retrieval call binding the contract method 0x89dbe24b.
//
// Solidity: function getMedianPrice(bytes32 linkId) constant returns(uint256)
func (_PriceFeeders *PriceFeedersCaller) GetMedianPrice(opts *bind.CallOpts, linkId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "getMedianPrice", linkId)
	return *ret0, err
}

// GetMedianPrice is a free data retrieval call binding the contract method 0x89dbe24b.
//
// Solidity: function getMedianPrice(bytes32 linkId) constant returns(uint256)
func (_PriceFeeders *PriceFeedersSession) GetMedianPrice(linkId [32]byte) (*big.Int, error) {
	return _PriceFeeders.Contract.GetMedianPrice(&_PriceFeeders.CallOpts, linkId)
}

// GetMedianPrice is a free data retrieval call binding the contract method 0x89dbe24b.
//
// Solidity: function getMedianPrice(bytes32 linkId) constant returns(uint256)
func (_PriceFeeders *PriceFeedersCallerSession) GetMedianPrice(linkId [32]byte) (*big.Int, error) {
	return _PriceFeeders.Contract.GetMedianPrice(&_PriceFeeders.CallOpts, linkId)
}

// GetPriceFeeders is a free data retrieval call binding the contract method 0x21f73ab0.
//
// Solidity: function getPriceFeeders(bytes32 linkId) constant returns(address[])
func (_PriceFeeders *PriceFeedersCaller) GetPriceFeeders(opts *bind.CallOpts, linkId [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "getPriceFeeders", linkId)
	return *ret0, err
}

// GetPriceFeeders is a free data retrieval call binding the contract method 0x21f73ab0.
//
// Solidity: function getPriceFeeders(bytes32 linkId) constant returns(address[])
func (_PriceFeeders *PriceFeedersSession) GetPriceFeeders(linkId [32]byte) ([]common.Address, error) {
	return _PriceFeeders.Contract.GetPriceFeeders(&_PriceFeeders.CallOpts, linkId)
}

// GetPriceFeeders is a free data retrieval call binding the contract method 0x21f73ab0.
//
// Solidity: function getPriceFeeders(bytes32 linkId) constant returns(address[])
func (_PriceFeeders *PriceFeedersCallerSession) GetPriceFeeders(linkId [32]byte) ([]common.Address, error) {
	return _PriceFeeders.Contract.GetPriceFeeders(&_PriceFeeders.CallOpts, linkId)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_PriceFeeders *PriceFeedersCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_PriceFeeders *PriceFeedersSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _PriceFeeders.Contract.IsWhitelistAdmin(&_PriceFeeders.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_PriceFeeders *PriceFeedersCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _PriceFeeders.Contract.IsWhitelistAdmin(&_PriceFeeders.CallOpts, account)
}

// MedianPrices is a free data retrieval call binding the contract method 0xd8d4f09a.
//
// Solidity: function medianPrices(bytes32 ) constant returns(uint256)
func (_PriceFeeders *PriceFeedersCaller) MedianPrices(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "medianPrices", arg0)
	return *ret0, err
}

// MedianPrices is a free data retrieval call binding the contract method 0xd8d4f09a.
//
// Solidity: function medianPrices(bytes32 ) constant returns(uint256)
func (_PriceFeeders *PriceFeedersSession) MedianPrices(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeeders.Contract.MedianPrices(&_PriceFeeders.CallOpts, arg0)
}

// MedianPrices is a free data retrieval call binding the contract method 0xd8d4f09a.
//
// Solidity: function medianPrices(bytes32 ) constant returns(uint256)
func (_PriceFeeders *PriceFeedersCallerSession) MedianPrices(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeeders.Contract.MedianPrices(&_PriceFeeders.CallOpts, arg0)
}

// PriceFeeders is a free data retrieval call binding the contract method 0xc966a39f.
//
// Solidity: function priceFeeders(bytes32 , uint256 ) constant returns(address)
func (_PriceFeeders *PriceFeedersCaller) PriceFeeders(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "priceFeeders", arg0, arg1)
	return *ret0, err
}

// PriceFeeders is a free data retrieval call binding the contract method 0xc966a39f.
//
// Solidity: function priceFeeders(bytes32 , uint256 ) constant returns(address)
func (_PriceFeeders *PriceFeedersSession) PriceFeeders(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _PriceFeeders.Contract.PriceFeeders(&_PriceFeeders.CallOpts, arg0, arg1)
}

// PriceFeeders is a free data retrieval call binding the contract method 0xc966a39f.
//
// Solidity: function priceFeeders(bytes32 , uint256 ) constant returns(address)
func (_PriceFeeders *PriceFeedersCallerSession) PriceFeeders(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _PriceFeeders.Contract.PriceFeeders(&_PriceFeeders.CallOpts, arg0, arg1)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(uint256)
func (_PriceFeeders *PriceFeedersCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceFeeders.contract.Call(opts, out, "prices", arg0)
	return *ret0, err
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(uint256)
func (_PriceFeeders *PriceFeedersSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeeders.Contract.Prices(&_PriceFeeders.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(uint256)
func (_PriceFeeders *PriceFeedersCallerSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeeders.Contract.Prices(&_PriceFeeders.CallOpts, arg0)
}

// AddAssetFiat is a paid mutator transaction binding the contract method 0x43cb0101.
//
// Solidity: function addAssetFiat(bytes32 assetCode, bytes32 currency) returns()
func (_PriceFeeders *PriceFeedersTransactor) AddAssetFiat(opts *bind.TransactOpts, assetCode [32]byte, currency [32]byte) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "addAssetFiat", assetCode, currency)
}

// AddAssetFiat is a paid mutator transaction binding the contract method 0x43cb0101.
//
// Solidity: function addAssetFiat(bytes32 assetCode, bytes32 currency) returns()
func (_PriceFeeders *PriceFeedersSession) AddAssetFiat(assetCode [32]byte, currency [32]byte) (*types.Transaction, error) {
	return _PriceFeeders.Contract.AddAssetFiat(&_PriceFeeders.TransactOpts, assetCode, currency)
}

// AddAssetFiat is a paid mutator transaction binding the contract method 0x43cb0101.
//
// Solidity: function addAssetFiat(bytes32 assetCode, bytes32 currency) returns()
func (_PriceFeeders *PriceFeedersTransactorSession) AddAssetFiat(assetCode [32]byte, currency [32]byte) (*types.Transaction, error) {
	return _PriceFeeders.Contract.AddAssetFiat(&_PriceFeeders.TransactOpts, assetCode, currency)
}

// AddPriceFeeder is a paid mutator transaction binding the contract method 0x367ae017.
//
// Solidity: function addPriceFeeder(bytes32 assetCode, bytes32 currency, address priceFeederAddr) returns()
func (_PriceFeeders *PriceFeedersTransactor) AddPriceFeeder(opts *bind.TransactOpts, assetCode [32]byte, currency [32]byte, priceFeederAddr common.Address) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "addPriceFeeder", assetCode, currency, priceFeederAddr)
}

// AddPriceFeeder is a paid mutator transaction binding the contract method 0x367ae017.
//
// Solidity: function addPriceFeeder(bytes32 assetCode, bytes32 currency, address priceFeederAddr) returns()
func (_PriceFeeders *PriceFeedersSession) AddPriceFeeder(assetCode [32]byte, currency [32]byte, priceFeederAddr common.Address) (*types.Transaction, error) {
	return _PriceFeeders.Contract.AddPriceFeeder(&_PriceFeeders.TransactOpts, assetCode, currency, priceFeederAddr)
}

// AddPriceFeeder is a paid mutator transaction binding the contract method 0x367ae017.
//
// Solidity: function addPriceFeeder(bytes32 assetCode, bytes32 currency, address priceFeederAddr) returns()
func (_PriceFeeders *PriceFeedersTransactorSession) AddPriceFeeder(assetCode [32]byte, currency [32]byte, priceFeederAddr common.Address) (*types.Transaction, error) {
	return _PriceFeeders.Contract.AddPriceFeeder(&_PriceFeeders.TransactOpts, assetCode, currency, priceFeederAddr)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_PriceFeeders *PriceFeedersTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_PriceFeeders *PriceFeedersSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _PriceFeeders.Contract.AddWhitelistAdmin(&_PriceFeeders.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_PriceFeeders *PriceFeedersTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _PriceFeeders.Contract.AddWhitelistAdmin(&_PriceFeeders.TransactOpts, account)
}

// RemoveAssetFiat is a paid mutator transaction binding the contract method 0x33e95457.
//
// Solidity: function removeAssetFiat(bytes32 assetCode, uint256 index) returns(bool)
func (_PriceFeeders *PriceFeedersTransactor) RemoveAssetFiat(opts *bind.TransactOpts, assetCode [32]byte, index *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "removeAssetFiat", assetCode, index)
}

// RemoveAssetFiat is a paid mutator transaction binding the contract method 0x33e95457.
//
// Solidity: function removeAssetFiat(bytes32 assetCode, uint256 index) returns(bool)
func (_PriceFeeders *PriceFeedersSession) RemoveAssetFiat(assetCode [32]byte, index *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.Contract.RemoveAssetFiat(&_PriceFeeders.TransactOpts, assetCode, index)
}

// RemoveAssetFiat is a paid mutator transaction binding the contract method 0x33e95457.
//
// Solidity: function removeAssetFiat(bytes32 assetCode, uint256 index) returns(bool)
func (_PriceFeeders *PriceFeedersTransactorSession) RemoveAssetFiat(assetCode [32]byte, index *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.Contract.RemoveAssetFiat(&_PriceFeeders.TransactOpts, assetCode, index)
}

// RemovePriceFeeder is a paid mutator transaction binding the contract method 0x15b9de67.
//
// Solidity: function removePriceFeeder(bytes32 pfId, uint256 index) returns(bool)
func (_PriceFeeders *PriceFeedersTransactor) RemovePriceFeeder(opts *bind.TransactOpts, pfId [32]byte, index *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "removePriceFeeder", pfId, index)
}

// RemovePriceFeeder is a paid mutator transaction binding the contract method 0x15b9de67.
//
// Solidity: function removePriceFeeder(bytes32 pfId, uint256 index) returns(bool)
func (_PriceFeeders *PriceFeedersSession) RemovePriceFeeder(pfId [32]byte, index *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.Contract.RemovePriceFeeder(&_PriceFeeders.TransactOpts, pfId, index)
}

// RemovePriceFeeder is a paid mutator transaction binding the contract method 0x15b9de67.
//
// Solidity: function removePriceFeeder(bytes32 pfId, uint256 index) returns(bool)
func (_PriceFeeders *PriceFeedersTransactorSession) RemovePriceFeeder(pfId [32]byte, index *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.Contract.RemovePriceFeeder(&_PriceFeeders.TransactOpts, pfId, index)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_PriceFeeders *PriceFeedersTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_PriceFeeders *PriceFeedersSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _PriceFeeders.Contract.RenounceWhitelistAdmin(&_PriceFeeders.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_PriceFeeders *PriceFeedersTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _PriceFeeders.Contract.RenounceWhitelistAdmin(&_PriceFeeders.TransactOpts)
}

// SetAsset is a paid mutator transaction binding the contract method 0x1610c6a6.
//
// Solidity: function setAsset(bytes32 assetCode, address assetAddress) returns()
func (_PriceFeeders *PriceFeedersTransactor) SetAsset(opts *bind.TransactOpts, assetCode [32]byte, assetAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "setAsset", assetCode, assetAddress)
}

// SetAsset is a paid mutator transaction binding the contract method 0x1610c6a6.
//
// Solidity: function setAsset(bytes32 assetCode, address assetAddress) returns()
func (_PriceFeeders *PriceFeedersSession) SetAsset(assetCode [32]byte, assetAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeders.Contract.SetAsset(&_PriceFeeders.TransactOpts, assetCode, assetAddress)
}

// SetAsset is a paid mutator transaction binding the contract method 0x1610c6a6.
//
// Solidity: function setAsset(bytes32 assetCode, address assetAddress) returns()
func (_PriceFeeders *PriceFeedersTransactorSession) SetAsset(assetCode [32]byte, assetAddress common.Address) (*types.Transaction, error) {
	return _PriceFeeders.Contract.SetAsset(&_PriceFeeders.TransactOpts, assetCode, assetAddress)
}

// SetPrice is a paid mutator transaction binding the contract method 0x213b9d9e.
//
// Solidity: function setPrice(bytes32 assetCode, bytes32 currency, uint256 price) returns()
func (_PriceFeeders *PriceFeedersTransactor) SetPrice(opts *bind.TransactOpts, assetCode [32]byte, currency [32]byte, price *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.contract.Transact(opts, "setPrice", assetCode, currency, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x213b9d9e.
//
// Solidity: function setPrice(bytes32 assetCode, bytes32 currency, uint256 price) returns()
func (_PriceFeeders *PriceFeedersSession) SetPrice(assetCode [32]byte, currency [32]byte, price *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.Contract.SetPrice(&_PriceFeeders.TransactOpts, assetCode, currency, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x213b9d9e.
//
// Solidity: function setPrice(bytes32 assetCode, bytes32 currency, uint256 price) returns()
func (_PriceFeeders *PriceFeedersTransactorSession) SetPrice(assetCode [32]byte, currency [32]byte, price *big.Int) (*types.Transaction, error) {
	return _PriceFeeders.Contract.SetPrice(&_PriceFeeders.TransactOpts, assetCode, currency, price)
}

// PriceFeedersWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the PriceFeeders contract.
type PriceFeedersWhitelistAdminAddedIterator struct {
	Event *PriceFeedersWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *PriceFeedersWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedersWhitelistAdminAdded)
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
		it.Event = new(PriceFeedersWhitelistAdminAdded)
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
func (it *PriceFeedersWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedersWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedersWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the PriceFeeders contract.
type PriceFeedersWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_PriceFeeders *PriceFeedersFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*PriceFeedersWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceFeeders.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedersWhitelistAdminAddedIterator{contract: _PriceFeeders.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_PriceFeeders *PriceFeedersFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *PriceFeedersWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceFeeders.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedersWhitelistAdminAdded)
				if err := _PriceFeeders.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// ParseWhitelistAdminAdded is a log parse operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_PriceFeeders *PriceFeedersFilterer) ParseWhitelistAdminAdded(log types.Log) (*PriceFeedersWhitelistAdminAdded, error) {
	event := new(PriceFeedersWhitelistAdminAdded)
	if err := _PriceFeeders.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceFeedersWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the PriceFeeders contract.
type PriceFeedersWhitelistAdminRemovedIterator struct {
	Event *PriceFeedersWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PriceFeedersWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedersWhitelistAdminRemoved)
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
		it.Event = new(PriceFeedersWhitelistAdminRemoved)
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
func (it *PriceFeedersWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedersWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedersWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the PriceFeeders contract.
type PriceFeedersWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_PriceFeeders *PriceFeedersFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PriceFeedersWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceFeeders.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedersWhitelistAdminRemovedIterator{contract: _PriceFeeders.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_PriceFeeders *PriceFeedersFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *PriceFeedersWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceFeeders.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedersWhitelistAdminRemoved)
				if err := _PriceFeeders.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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

// ParseWhitelistAdminRemoved is a log parse operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_PriceFeeders *PriceFeedersFilterer) ParseWhitelistAdminRemoved(log types.Log) (*PriceFeedersWhitelistAdminRemoved, error) {
	event := new(PriceFeedersWhitelistAdminRemoved)
	if err := _PriceFeeders.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
