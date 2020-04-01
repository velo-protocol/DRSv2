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

// MedianizerABI is the input ABI used to generate the binding from.
const MedianizerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"medianizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeder\",\"type\":\"address\"}],\"name\":\"FeederAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"medianizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeder\",\"type\":\"address\"}],\"name\":\"FeederRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"medianizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"MedianActivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"medianizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"MedianVoid\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeders\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"llSize\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFedPrices\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validityPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fiatCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_collateralCode\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"post\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newMinFedPrices\",\"type\":\"uint8\"}],\"name\":\"setMinFedPrices\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeder\",\"type\":\"address\"}],\"name\":\"addFeeder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeder\",\"type\":\"address\"}],\"name\":\"removeFeeder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidityPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"newValidityPeriod\",\"type\":\"int256\"}],\"name\":\"setValidityPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MedianizerBin is the compiled bytecode used for deploying new contracts.
var MedianizerBin = "0x608060405234801561001057600080fd5b506114ab806100206000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80635d967b9e116100b85780638e1207dc1161007c5780638e1207dc146102e1578063981761c514610307578063a035b1fe14610325578063ac4c25b21461032d578063d653bdc014610335578063e6f2ca271461035257610137565b80635d967b9e1461025b57806364bc52041461027b5780636d4ce63c1461028357806385f812c31461028b5780638da5cb5b146102bd57610137565b80631a43c338116100ff5780631a43c338146101aa5780631e6c3850146101cb57806341bf7fd9146101d3578063498b58281461022b5780635587c4361461025357610137565b806302fb0c5e1461013c57806305c1f502146101585780630e4e0c65146101625780630f15f4c01461017c578063155fb6ae14610184575b600080fd5b61014461035a565b604080519115158252519081900360200190f35b61016061036a565b005b61016a610383565b60408051918252519081900360200190f35b610160610389565b6101606004803603602081101561019a57600080fd5b50356001600160a01b03166104bb565b6101b2610560565b6040805192835290151560208301528051918290030190f35b61016a610868565b6101db61086e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102175781810151838201526020016101ff565b505050509050019250505060405180910390f35b61023361087f565b604080519384529115156020840152151582820152519081900360600190f35b61016a610897565b6101606004803603602081101561027157600080fd5b503560ff1661089d565b61016a61093b565b61016a610941565b610160600480360360608110156102a157600080fd5b506001600160a01b0381351690602081013590604001356109e8565b6102c5610ae6565b604080516001600160a01b039092168252519081900360200190f35b610160600480360360208110156102f757600080fd5b50356001600160a01b0316610af5565b61030f610c4d565b6040805160ff9092168252519081900360200190f35b61016a610c56565b610160610c5c565b6101606004803603602081101561034b57600080fd5b5035610d04565b61030f610dd2565b603554600160a01b900460ff1681565b6000610374610560565b50905061038081610ddb565b50565b603a5490565b6035546001600160a01b031633146103d25760405162461bcd60e51b815260040180806020018281038252605c8152602001806113ee605c913960600191505060405180910390fd5b603554600160a01b900460ff161561041b5760405162461bcd60e51b815260040180806020018281038252602d81526020018061144a602d913960400191505060405180910390fd5b60006036541161045c5760405162461bcd60e51b815260040180806020018281038252603f815260200180611381603f913960400191505060405180910390fd5b6035805460ff60a01b1916600160a01b90811791829055604080513381523060208201529190920460ff1615158183015290517f9137522953a20570987f56292aae83c97008172af9c07b27079c7e87918c05449181900360600190a1565b6035546001600160a01b031633146105045760405162461bcd60e51b815260040180806020018281038252605c8152602001806113ee605c913960600191505060405180910390fd5b61051560338263ffffffff610de016565b50604080513381523060208201526001600160a01b0383168183015290517fd2bb3386b022c08ee866ea92152a49a6f30f37aba23b60afe24e315a32e5ab289181900360600190a150565b60016000908152603460209081527f7fb70d3b53b50f6798636307bc089a811f848f8c80c93eb14575305275478318546033546040805160ff9092168083528085028301909401905283926001600160a01b03909216916060919080156105d1578160200160208202803883390190505b5090506000805b6001600160a01b038416600114610790576000806000866001600160a01b031663498b58286040518163ffffffff1660e01b815260040160606040518083038186803b15801561062757600080fd5b505afa15801561063b573d6000803e3d6000fd5b505050506040513d606081101561065157600080fd5b5080516020820151604090920151603a549195509193509091504203821015610678575060015b806107675760ff851615806106a65750856001860360ff168151811061069a57fe5b60200260200101518310155b156106cc5782868660ff16815181106106bb57fe5b602002602001018181525050610760565b60005b868160ff16815181106106de57fe5b602002602001015184106106f4576001016106cf565b855b8160ff168160ff16111561074157876001820360ff168151811061071657fe5b6020026020010151888260ff168151811061072d57fe5b6020908102919091010152600019016106f6565b5083878260ff168151811061075257fe5b602002602001018181525050505b6001909401935b5050506001600160a01b03938416600090815260346020526040902054909316926001016105d8565b5060395460ff90811690821610156107b357506000935060019250610864915050565b600060018216610836576000836001600260ff8616040360ff16815181106107d757fe5b6020026020010151905060008460028560ff16816107f157fe5b0460ff16815181106107ff57fe5b6020026020010151905061082d60026108218385610eba90919063ffffffff16565b9063ffffffff610f1b16565b9250505061085a565b82600260ff6000198501160460ff168151811061084f57fe5b602002602001015190505b9450600093505050505b9091565b603a5481565b606061087a6033610f5d565b905090565b60365460355460ff600160a01b909104168115909192565b60385481565b6035546001600160a01b031633146108e65760405162461bcd60e51b815260040180806020018281038252605c8152602001806113ee605c913960600191505060405180910390fd5b60ff81166109255760405162461bcd60e51b81526004018080602001828103825260238152602001806112c66023913960400191505060405180910390fd5b6039805460ff191660ff92909216919091179055565b60375481565b60008060365411610999576040805162461bcd60e51b815260206004820152601d60248201527f4d656469616e697a65722e6765743a20696e76616c6964207072696365000000604482015290519081900360640190fd5b603554600160a01b900460ff166109e15760405162461bcd60e51b81526004018080602001828103825260268152602001806112e96026913960400191505060405180910390fd5b5060365490565b600054610100900460ff1680610a015750610a01611012565b80610a0f575060005460ff16155b610a4a5760405162461bcd60e51b815260040180806020018281038252602e8152602001806113c0602e913960400191505060405180910390fd5b600054610100900460ff16158015610a75576000805460ff1961ff0019909116610100171660011790555b60358054600160a01b60ff60a01b19909116176001600160a01b0319166001600160a01b038616179055603783905560388290556039805460ff19166001179055610384603a5562278d00603b55610acd6033611018565b508015610ae0576000805461ff00191690555b50505050565b6035546001600160a01b031681565b6035546001600160a01b03163314610b3e5760405162461bcd60e51b815260040180806020018281038252605c8152602001806113ee605c913960600191505060405180910390fd5b6060610b4861086e565b905060016000805b8351811015610bb057838181518110610b6557fe5b60200260200101516001600160a01b0316856001600160a01b03161415610b8f5760019150610bb0565b838181518110610b9b57fe5b60209081029190910101519250600101610b50565b5080610bed5760405162461bcd60e51b815260040180806020018281038252602f81526020018061130f602f913960400191505060405180910390fd5b610bff6033858463ffffffff61105916565b50604080513381523060208201526001600160a01b0386168183015290517f2dc15dd92190df6fa5ffe9febdce6505f3fc1b4f3846044441a4e37ccecfe6fb9181900360600190a150505050565b60335460ff1681565b60365481565b6035546001600160a01b03163314610ca55760405162461bcd60e51b815260040180806020018281038252605c8152602001806113ee605c913960600191505060405180910390fd5b60006036556035805460ff60a01b19169081905560408051338152306020820152600160a01b90920460ff16151582820152517fc1c85319c7923477ac646408ce7d1255f85bf413c80258f2367de8f4fd51c9509181900360600190a1565b6035546001600160a01b03163314610d4d5760405162461bcd60e51b815260040180806020018281038252605c8152602001806113ee605c913960600191505060405180910390fd5b60008113610d8c5760405162461bcd60e51b815260040180806020018281038252604381526020018061133e6043913960600191505060405180910390fd5b603b54811115610dcd5760405162461bcd60e51b815260040180806020018281038252605781526020018061126f6057913960600191505060405180910390fd5b603a55565b60395460ff1681565b603655565b610de861125c565b610df28383611195565b15610e44576040805162461bcd60e51b815260206004820152601b60248201527f6164647220697320616c726561647920696e20746865206c6973740000000000604482015290519081900360640190fd5b50600160008181528382016020908152604080832080546001600160a01b039687168086528386208054989092166001600160a01b031998891617909155938590528054909516909217909355835460ff19811660ff918216909301811692909217938490558051928301905291909116815290565b600082820183811015610f14576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000610f1483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506111ba565b80546040805160ff909216808352602080820284010190915260609182918015610f91578160200160208202803883390190505b506001600081815290850160205260408120549192506001600160a01b03909116905b6001600160a01b0382166001146110095781838281518110610fd257fe5b6001600160a01b03928316602091820292909201810191909152928116600090815260018088019094526040902054169101610fb4565b50909392505050565b303b1590565b61102061125c565b5060016000818152828201602090815260409182902080546001600160a01b03191690931790925580519182019052905460ff16815290565b61106161125c565b61106b8484611195565b6110bc576040805162461bcd60e51b815260206004820152601860248201527f61646472206e6f742077686974656c6973746564207965740000000000000000604482015290519081900360640190fd5b6001600160a01b038281166000908152600186016020526040902054811690841614611124576040805162461bcd60e51b81526020600482015260126024820152713bb937b73390383932bb21b7b739bab6b2b960711b604482015290519081900360640190fd5b506001600160a01b0391821660008181526001850160209081526040808320805495871684528184208054969097166001600160a01b031996871617909655929091528354909216909255825460ff19811660ff918216600019018216179384905582519182019092529116815290565b6001600160a01b03818116600090815260018401602052604090205416151592915050565b600081836112465760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561120b5781810151838201526020016111f3565b50505050905090810190601f1680156112385780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161125257fe5b0495945050505050565b6040805160208101909152600081529056fe4d656469616e697a65722e73657456616c6964697479506572696f643a2076616c6964697479506572696f64206d757374206e6f742062652067726561746572207468616e206d617856616c6964697479506572696f644d6564207c206d696e466564507269636573206d757374206d6f7265207468616e20304d656469616e697a65722e6765743a206d656469616e697a657220697320696e6163746976654d656469616e697a65722e72656d6f76654665656465723a206164647265737320646f6573206e6f742065786973744d656469616e697a65722e73657456616c6964697479506572696f643a2076616c6964697479506572696f64206d7573742062652067726561746572207468616e20304d656469616e697a65722e61637469766174653a20746865206d656469616e697a657220646f6573206e6f74206861766520612076616c6964207072696365436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65644d656469616e697a65722e6f6e6c794f776e65723a20546865206d6573736167652073656e646572206973206e6f7420666f756e64206f7220646f6573206e6f7420686176652073756666696369656e74207065726d697373696f6e4d656469616e697a65722e61637469766174653a20746865206d656469616e697a657220697320616374697665a265627a7a7231582002f3335deb1954ef795916f4f81c36f4b55417d2aed0476bd5358bb49991560064736f6c63430005100032"

// DeployMedianizer deploys a new Ethereum contract, binding an instance of Medianizer to it.
func DeployMedianizer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Medianizer, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MedianizerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Medianizer{MedianizerCaller: MedianizerCaller{contract: contract}, MedianizerTransactor: MedianizerTransactor{contract: contract}, MedianizerFilterer: MedianizerFilterer{contract: contract}}, nil
}

// Medianizer is an auto generated Go binding around an Ethereum contract.
type Medianizer struct {
	MedianizerCaller     // Read-only binding to the contract
	MedianizerTransactor // Write-only binding to the contract
	MedianizerFilterer   // Log filterer for contract events
}

// MedianizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedianizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedianizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedianizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedianizerSession struct {
	Contract     *Medianizer       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedianizerCallerSession struct {
	Contract *MedianizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MedianizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedianizerTransactorSession struct {
	Contract     *MedianizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MedianizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedianizerRaw struct {
	Contract *Medianizer // Generic contract binding to access the raw methods on
}

// MedianizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedianizerCallerRaw struct {
	Contract *MedianizerCaller // Generic read-only contract binding to access the raw methods on
}

// MedianizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedianizerTransactorRaw struct {
	Contract *MedianizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedianizer creates a new instance of Medianizer, bound to a specific deployed contract.
func NewMedianizer(address common.Address, backend bind.ContractBackend) (*Medianizer, error) {
	contract, err := bindMedianizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Medianizer{MedianizerCaller: MedianizerCaller{contract: contract}, MedianizerTransactor: MedianizerTransactor{contract: contract}, MedianizerFilterer: MedianizerFilterer{contract: contract}}, nil
}

// NewMedianizerCaller creates a new read-only instance of Medianizer, bound to a specific deployed contract.
func NewMedianizerCaller(address common.Address, caller bind.ContractCaller) (*MedianizerCaller, error) {
	contract, err := bindMedianizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedianizerCaller{contract: contract}, nil
}

// NewMedianizerTransactor creates a new write-only instance of Medianizer, bound to a specific deployed contract.
func NewMedianizerTransactor(address common.Address, transactor bind.ContractTransactor) (*MedianizerTransactor, error) {
	contract, err := bindMedianizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedianizerTransactor{contract: contract}, nil
}

// NewMedianizerFilterer creates a new log filterer instance of Medianizer, bound to a specific deployed contract.
func NewMedianizerFilterer(address common.Address, filterer bind.ContractFilterer) (*MedianizerFilterer, error) {
	contract, err := bindMedianizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedianizerFilterer{contract: contract}, nil
}

// bindMedianizer binds a generic wrapper to an already deployed contract.
func bindMedianizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medianizer *MedianizerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Medianizer.Contract.MedianizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medianizer *MedianizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.Contract.MedianizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medianizer *MedianizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medianizer.Contract.MedianizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medianizer *MedianizerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Medianizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medianizer *MedianizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medianizer *MedianizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medianizer.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Medianizer *MedianizerCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Medianizer *MedianizerSession) Active() (bool, error) {
	return _Medianizer.Contract.Active(&_Medianizer.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Medianizer *MedianizerCallerSession) Active() (bool, error) {
	return _Medianizer.Contract.Active(&_Medianizer.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Medianizer *MedianizerCaller) CollateralCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "collateralCode")
	return *ret0, err
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Medianizer *MedianizerSession) CollateralCode() ([32]byte, error) {
	return _Medianizer.Contract.CollateralCode(&_Medianizer.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Medianizer *MedianizerCallerSession) CollateralCode() ([32]byte, error) {
	return _Medianizer.Contract.CollateralCode(&_Medianizer.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(uint256, bool)
func (_Medianizer *MedianizerCaller) Compute(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Medianizer.contract.Call(opts, out, "compute")
	return *ret0, *ret1, err
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(uint256, bool)
func (_Medianizer *MedianizerSession) Compute() (*big.Int, bool, error) {
	return _Medianizer.Contract.Compute(&_Medianizer.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(uint256, bool)
func (_Medianizer *MedianizerCallerSession) Compute() (*big.Int, bool, error) {
	return _Medianizer.Contract.Compute(&_Medianizer.CallOpts)
}

// Feeders is a free data retrieval call binding the contract method 0x981761c5.
//
// Solidity: function feeders() constant returns(uint8 llSize)
func (_Medianizer *MedianizerCaller) Feeders(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "feeders")
	return *ret0, err
}

// Feeders is a free data retrieval call binding the contract method 0x981761c5.
//
// Solidity: function feeders() constant returns(uint8 llSize)
func (_Medianizer *MedianizerSession) Feeders() (uint8, error) {
	return _Medianizer.Contract.Feeders(&_Medianizer.CallOpts)
}

// Feeders is a free data retrieval call binding the contract method 0x981761c5.
//
// Solidity: function feeders() constant returns(uint8 llSize)
func (_Medianizer *MedianizerCallerSession) Feeders() (uint8, error) {
	return _Medianizer.Contract.Feeders(&_Medianizer.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Medianizer *MedianizerCaller) FiatCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "fiatCode")
	return *ret0, err
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Medianizer *MedianizerSession) FiatCode() ([32]byte, error) {
	return _Medianizer.Contract.FiatCode(&_Medianizer.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Medianizer *MedianizerCallerSession) FiatCode() ([32]byte, error) {
	return _Medianizer.Contract.FiatCode(&_Medianizer.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Medianizer *MedianizerCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Medianizer *MedianizerSession) Get() (*big.Int, error) {
	return _Medianizer.Contract.Get(&_Medianizer.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(uint256)
func (_Medianizer *MedianizerCallerSession) Get() (*big.Int, error) {
	return _Medianizer.Contract.Get(&_Medianizer.CallOpts)
}

// GetFeeders is a free data retrieval call binding the contract method 0x41bf7fd9.
//
// Solidity: function getFeeders() constant returns(address[])
func (_Medianizer *MedianizerCaller) GetFeeders(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "getFeeders")
	return *ret0, err
}

// GetFeeders is a free data retrieval call binding the contract method 0x41bf7fd9.
//
// Solidity: function getFeeders() constant returns(address[])
func (_Medianizer *MedianizerSession) GetFeeders() ([]common.Address, error) {
	return _Medianizer.Contract.GetFeeders(&_Medianizer.CallOpts)
}

// GetFeeders is a free data retrieval call binding the contract method 0x41bf7fd9.
//
// Solidity: function getFeeders() constant returns(address[])
func (_Medianizer *MedianizerCallerSession) GetFeeders() ([]common.Address, error) {
	return _Medianizer.Contract.GetFeeders(&_Medianizer.CallOpts)
}

// GetValidityPeriod is a free data retrieval call binding the contract method 0x0e4e0c65.
//
// Solidity: function getValidityPeriod() constant returns(uint256)
func (_Medianizer *MedianizerCaller) GetValidityPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "getValidityPeriod")
	return *ret0, err
}

// GetValidityPeriod is a free data retrieval call binding the contract method 0x0e4e0c65.
//
// Solidity: function getValidityPeriod() constant returns(uint256)
func (_Medianizer *MedianizerSession) GetValidityPeriod() (*big.Int, error) {
	return _Medianizer.Contract.GetValidityPeriod(&_Medianizer.CallOpts)
}

// GetValidityPeriod is a free data retrieval call binding the contract method 0x0e4e0c65.
//
// Solidity: function getValidityPeriod() constant returns(uint256)
func (_Medianizer *MedianizerCallerSession) GetValidityPeriod() (*big.Int, error) {
	return _Medianizer.Contract.GetValidityPeriod(&_Medianizer.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Medianizer *MedianizerCaller) GetWithError(opts *bind.CallOpts) (*big.Int, bool, bool, error) {
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
	err := _Medianizer.contract.Call(opts, out, "getWithError")
	return *ret0, *ret1, *ret2, err
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Medianizer *MedianizerSession) GetWithError() (*big.Int, bool, bool, error) {
	return _Medianizer.Contract.GetWithError(&_Medianizer.CallOpts)
}

// GetWithError is a free data retrieval call binding the contract method 0x498b5828.
//
// Solidity: function getWithError() constant returns(uint256, bool, bool)
func (_Medianizer *MedianizerCallerSession) GetWithError() (*big.Int, bool, bool, error) {
	return _Medianizer.Contract.GetWithError(&_Medianizer.CallOpts)
}

// MinFedPrices is a free data retrieval call binding the contract method 0xe6f2ca27.
//
// Solidity: function minFedPrices() constant returns(uint8)
func (_Medianizer *MedianizerCaller) MinFedPrices(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "minFedPrices")
	return *ret0, err
}

// MinFedPrices is a free data retrieval call binding the contract method 0xe6f2ca27.
//
// Solidity: function minFedPrices() constant returns(uint8)
func (_Medianizer *MedianizerSession) MinFedPrices() (uint8, error) {
	return _Medianizer.Contract.MinFedPrices(&_Medianizer.CallOpts)
}

// MinFedPrices is a free data retrieval call binding the contract method 0xe6f2ca27.
//
// Solidity: function minFedPrices() constant returns(uint8)
func (_Medianizer *MedianizerCallerSession) MinFedPrices() (uint8, error) {
	return _Medianizer.Contract.MinFedPrices(&_Medianizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Medianizer *MedianizerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Medianizer *MedianizerSession) Owner() (common.Address, error) {
	return _Medianizer.Contract.Owner(&_Medianizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Medianizer *MedianizerCallerSession) Owner() (common.Address, error) {
	return _Medianizer.Contract.Owner(&_Medianizer.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Medianizer *MedianizerCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Medianizer *MedianizerSession) Price() (*big.Int, error) {
	return _Medianizer.Contract.Price(&_Medianizer.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Medianizer *MedianizerCallerSession) Price() (*big.Int, error) {
	return _Medianizer.Contract.Price(&_Medianizer.CallOpts)
}

// ValidityPeriod is a free data retrieval call binding the contract method 0x1e6c3850.
//
// Solidity: function validityPeriod() constant returns(uint256)
func (_Medianizer *MedianizerCaller) ValidityPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "validityPeriod")
	return *ret0, err
}

// ValidityPeriod is a free data retrieval call binding the contract method 0x1e6c3850.
//
// Solidity: function validityPeriod() constant returns(uint256)
func (_Medianizer *MedianizerSession) ValidityPeriod() (*big.Int, error) {
	return _Medianizer.Contract.ValidityPeriod(&_Medianizer.CallOpts)
}

// ValidityPeriod is a free data retrieval call binding the contract method 0x1e6c3850.
//
// Solidity: function validityPeriod() constant returns(uint256)
func (_Medianizer *MedianizerCallerSession) ValidityPeriod() (*big.Int, error) {
	return _Medianizer.Contract.ValidityPeriod(&_Medianizer.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Medianizer *MedianizerTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Medianizer *MedianizerSession) Activate() (*types.Transaction, error) {
	return _Medianizer.Contract.Activate(&_Medianizer.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Medianizer *MedianizerTransactorSession) Activate() (*types.Transaction, error) {
	return _Medianizer.Contract.Activate(&_Medianizer.TransactOpts)
}

// AddFeeder is a paid mutator transaction binding the contract method 0x155fb6ae.
//
// Solidity: function addFeeder(address feeder) returns()
func (_Medianizer *MedianizerTransactor) AddFeeder(opts *bind.TransactOpts, feeder common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "addFeeder", feeder)
}

// AddFeeder is a paid mutator transaction binding the contract method 0x155fb6ae.
//
// Solidity: function addFeeder(address feeder) returns()
func (_Medianizer *MedianizerSession) AddFeeder(feeder common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.AddFeeder(&_Medianizer.TransactOpts, feeder)
}

// AddFeeder is a paid mutator transaction binding the contract method 0x155fb6ae.
//
// Solidity: function addFeeder(address feeder) returns()
func (_Medianizer *MedianizerTransactorSession) AddFeeder(feeder common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.AddFeeder(&_Medianizer.TransactOpts, feeder)
}

// Initialize is a paid mutator transaction binding the contract method 0x85f812c3.
//
// Solidity: function initialize(address _owner, bytes32 _fiatCode, bytes32 _collateralCode) returns()
func (_Medianizer *MedianizerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _fiatCode [32]byte, _collateralCode [32]byte) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "initialize", _owner, _fiatCode, _collateralCode)
}

// Initialize is a paid mutator transaction binding the contract method 0x85f812c3.
//
// Solidity: function initialize(address _owner, bytes32 _fiatCode, bytes32 _collateralCode) returns()
func (_Medianizer *MedianizerSession) Initialize(_owner common.Address, _fiatCode [32]byte, _collateralCode [32]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.Initialize(&_Medianizer.TransactOpts, _owner, _fiatCode, _collateralCode)
}

// Initialize is a paid mutator transaction binding the contract method 0x85f812c3.
//
// Solidity: function initialize(address _owner, bytes32 _fiatCode, bytes32 _collateralCode) returns()
func (_Medianizer *MedianizerTransactorSession) Initialize(_owner common.Address, _fiatCode [32]byte, _collateralCode [32]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.Initialize(&_Medianizer.TransactOpts, _owner, _fiatCode, _collateralCode)
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns()
func (_Medianizer *MedianizerTransactor) Post(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "post")
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns()
func (_Medianizer *MedianizerSession) Post() (*types.Transaction, error) {
	return _Medianizer.Contract.Post(&_Medianizer.TransactOpts)
}

// Post is a paid mutator transaction binding the contract method 0x05c1f502.
//
// Solidity: function post() returns()
func (_Medianizer *MedianizerTransactorSession) Post() (*types.Transaction, error) {
	return _Medianizer.Contract.Post(&_Medianizer.TransactOpts)
}

// RemoveFeeder is a paid mutator transaction binding the contract method 0x8e1207dc.
//
// Solidity: function removeFeeder(address feeder) returns()
func (_Medianizer *MedianizerTransactor) RemoveFeeder(opts *bind.TransactOpts, feeder common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "removeFeeder", feeder)
}

// RemoveFeeder is a paid mutator transaction binding the contract method 0x8e1207dc.
//
// Solidity: function removeFeeder(address feeder) returns()
func (_Medianizer *MedianizerSession) RemoveFeeder(feeder common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.RemoveFeeder(&_Medianizer.TransactOpts, feeder)
}

// RemoveFeeder is a paid mutator transaction binding the contract method 0x8e1207dc.
//
// Solidity: function removeFeeder(address feeder) returns()
func (_Medianizer *MedianizerTransactorSession) RemoveFeeder(feeder common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.RemoveFeeder(&_Medianizer.TransactOpts, feeder)
}

// SetMinFedPrices is a paid mutator transaction binding the contract method 0x5d967b9e.
//
// Solidity: function setMinFedPrices(uint8 newMinFedPrices) returns()
func (_Medianizer *MedianizerTransactor) SetMinFedPrices(opts *bind.TransactOpts, newMinFedPrices uint8) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "setMinFedPrices", newMinFedPrices)
}

// SetMinFedPrices is a paid mutator transaction binding the contract method 0x5d967b9e.
//
// Solidity: function setMinFedPrices(uint8 newMinFedPrices) returns()
func (_Medianizer *MedianizerSession) SetMinFedPrices(newMinFedPrices uint8) (*types.Transaction, error) {
	return _Medianizer.Contract.SetMinFedPrices(&_Medianizer.TransactOpts, newMinFedPrices)
}

// SetMinFedPrices is a paid mutator transaction binding the contract method 0x5d967b9e.
//
// Solidity: function setMinFedPrices(uint8 newMinFedPrices) returns()
func (_Medianizer *MedianizerTransactorSession) SetMinFedPrices(newMinFedPrices uint8) (*types.Transaction, error) {
	return _Medianizer.Contract.SetMinFedPrices(&_Medianizer.TransactOpts, newMinFedPrices)
}

// SetValidityPeriod is a paid mutator transaction binding the contract method 0xd653bdc0.
//
// Solidity: function setValidityPeriod(int256 newValidityPeriod) returns()
func (_Medianizer *MedianizerTransactor) SetValidityPeriod(opts *bind.TransactOpts, newValidityPeriod *big.Int) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "setValidityPeriod", newValidityPeriod)
}

// SetValidityPeriod is a paid mutator transaction binding the contract method 0xd653bdc0.
//
// Solidity: function setValidityPeriod(int256 newValidityPeriod) returns()
func (_Medianizer *MedianizerSession) SetValidityPeriod(newValidityPeriod *big.Int) (*types.Transaction, error) {
	return _Medianizer.Contract.SetValidityPeriod(&_Medianizer.TransactOpts, newValidityPeriod)
}

// SetValidityPeriod is a paid mutator transaction binding the contract method 0xd653bdc0.
//
// Solidity: function setValidityPeriod(int256 newValidityPeriod) returns()
func (_Medianizer *MedianizerTransactorSession) SetValidityPeriod(newValidityPeriod *big.Int) (*types.Transaction, error) {
	return _Medianizer.Contract.SetValidityPeriod(&_Medianizer.TransactOpts, newValidityPeriod)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Medianizer *MedianizerTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Medianizer *MedianizerSession) Void() (*types.Transaction, error) {
	return _Medianizer.Contract.Void(&_Medianizer.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Medianizer *MedianizerTransactorSession) Void() (*types.Transaction, error) {
	return _Medianizer.Contract.Void(&_Medianizer.TransactOpts)
}

// MedianizerFeederAddIterator is returned from FilterFeederAdd and is used to iterate over the raw logs and unpacked data for FeederAdd events raised by the Medianizer contract.
type MedianizerFeederAddIterator struct {
	Event *MedianizerFeederAdd // Event containing the contract specifics and raw log

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
func (it *MedianizerFeederAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerFeederAdd)
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
		it.Event = new(MedianizerFeederAdd)
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
func (it *MedianizerFeederAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerFeederAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerFeederAdd represents a FeederAdd event raised by the Medianizer contract.
type MedianizerFeederAdd struct {
	Caller     common.Address
	Medianizer common.Address
	Feeder     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeederAdd is a free log retrieval operation binding the contract event 0xd2bb3386b022c08ee866ea92152a49a6f30f37aba23b60afe24e315a32e5ab28.
//
// Solidity: event FeederAdd(address caller, address medianizer, address feeder)
func (_Medianizer *MedianizerFilterer) FilterFeederAdd(opts *bind.FilterOpts) (*MedianizerFeederAddIterator, error) {

	logs, sub, err := _Medianizer.contract.FilterLogs(opts, "FeederAdd")
	if err != nil {
		return nil, err
	}
	return &MedianizerFeederAddIterator{contract: _Medianizer.contract, event: "FeederAdd", logs: logs, sub: sub}, nil
}

// WatchFeederAdd is a free log subscription operation binding the contract event 0xd2bb3386b022c08ee866ea92152a49a6f30f37aba23b60afe24e315a32e5ab28.
//
// Solidity: event FeederAdd(address caller, address medianizer, address feeder)
func (_Medianizer *MedianizerFilterer) WatchFeederAdd(opts *bind.WatchOpts, sink chan<- *MedianizerFeederAdd) (event.Subscription, error) {

	logs, sub, err := _Medianizer.contract.WatchLogs(opts, "FeederAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerFeederAdd)
				if err := _Medianizer.contract.UnpackLog(event, "FeederAdd", log); err != nil {
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

// ParseFeederAdd is a log parse operation binding the contract event 0xd2bb3386b022c08ee866ea92152a49a6f30f37aba23b60afe24e315a32e5ab28.
//
// Solidity: event FeederAdd(address caller, address medianizer, address feeder)
func (_Medianizer *MedianizerFilterer) ParseFeederAdd(log types.Log) (*MedianizerFeederAdd, error) {
	event := new(MedianizerFeederAdd)
	if err := _Medianizer.contract.UnpackLog(event, "FeederAdd", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MedianizerFeederRemoveIterator is returned from FilterFeederRemove and is used to iterate over the raw logs and unpacked data for FeederRemove events raised by the Medianizer contract.
type MedianizerFeederRemoveIterator struct {
	Event *MedianizerFeederRemove // Event containing the contract specifics and raw log

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
func (it *MedianizerFeederRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerFeederRemove)
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
		it.Event = new(MedianizerFeederRemove)
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
func (it *MedianizerFeederRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerFeederRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerFeederRemove represents a FeederRemove event raised by the Medianizer contract.
type MedianizerFeederRemove struct {
	Caller     common.Address
	Medianizer common.Address
	Feeder     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeederRemove is a free log retrieval operation binding the contract event 0x2dc15dd92190df6fa5ffe9febdce6505f3fc1b4f3846044441a4e37ccecfe6fb.
//
// Solidity: event FeederRemove(address caller, address medianizer, address feeder)
func (_Medianizer *MedianizerFilterer) FilterFeederRemove(opts *bind.FilterOpts) (*MedianizerFeederRemoveIterator, error) {

	logs, sub, err := _Medianizer.contract.FilterLogs(opts, "FeederRemove")
	if err != nil {
		return nil, err
	}
	return &MedianizerFeederRemoveIterator{contract: _Medianizer.contract, event: "FeederRemove", logs: logs, sub: sub}, nil
}

// WatchFeederRemove is a free log subscription operation binding the contract event 0x2dc15dd92190df6fa5ffe9febdce6505f3fc1b4f3846044441a4e37ccecfe6fb.
//
// Solidity: event FeederRemove(address caller, address medianizer, address feeder)
func (_Medianizer *MedianizerFilterer) WatchFeederRemove(opts *bind.WatchOpts, sink chan<- *MedianizerFeederRemove) (event.Subscription, error) {

	logs, sub, err := _Medianizer.contract.WatchLogs(opts, "FeederRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerFeederRemove)
				if err := _Medianizer.contract.UnpackLog(event, "FeederRemove", log); err != nil {
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

// ParseFeederRemove is a log parse operation binding the contract event 0x2dc15dd92190df6fa5ffe9febdce6505f3fc1b4f3846044441a4e37ccecfe6fb.
//
// Solidity: event FeederRemove(address caller, address medianizer, address feeder)
func (_Medianizer *MedianizerFilterer) ParseFeederRemove(log types.Log) (*MedianizerFeederRemove, error) {
	event := new(MedianizerFeederRemove)
	if err := _Medianizer.contract.UnpackLog(event, "FeederRemove", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MedianizerMedianActivateIterator is returned from FilterMedianActivate and is used to iterate over the raw logs and unpacked data for MedianActivate events raised by the Medianizer contract.
type MedianizerMedianActivateIterator struct {
	Event *MedianizerMedianActivate // Event containing the contract specifics and raw log

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
func (it *MedianizerMedianActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerMedianActivate)
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
		it.Event = new(MedianizerMedianActivate)
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
func (it *MedianizerMedianActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerMedianActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerMedianActivate represents a MedianActivate event raised by the Medianizer contract.
type MedianizerMedianActivate struct {
	Caller     common.Address
	Medianizer common.Address
	IsActive   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMedianActivate is a free log retrieval operation binding the contract event 0x9137522953a20570987f56292aae83c97008172af9c07b27079c7e87918c0544.
//
// Solidity: event MedianActivate(address caller, address medianizer, bool isActive)
func (_Medianizer *MedianizerFilterer) FilterMedianActivate(opts *bind.FilterOpts) (*MedianizerMedianActivateIterator, error) {

	logs, sub, err := _Medianizer.contract.FilterLogs(opts, "MedianActivate")
	if err != nil {
		return nil, err
	}
	return &MedianizerMedianActivateIterator{contract: _Medianizer.contract, event: "MedianActivate", logs: logs, sub: sub}, nil
}

// WatchMedianActivate is a free log subscription operation binding the contract event 0x9137522953a20570987f56292aae83c97008172af9c07b27079c7e87918c0544.
//
// Solidity: event MedianActivate(address caller, address medianizer, bool isActive)
func (_Medianizer *MedianizerFilterer) WatchMedianActivate(opts *bind.WatchOpts, sink chan<- *MedianizerMedianActivate) (event.Subscription, error) {

	logs, sub, err := _Medianizer.contract.WatchLogs(opts, "MedianActivate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerMedianActivate)
				if err := _Medianizer.contract.UnpackLog(event, "MedianActivate", log); err != nil {
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

// ParseMedianActivate is a log parse operation binding the contract event 0x9137522953a20570987f56292aae83c97008172af9c07b27079c7e87918c0544.
//
// Solidity: event MedianActivate(address caller, address medianizer, bool isActive)
func (_Medianizer *MedianizerFilterer) ParseMedianActivate(log types.Log) (*MedianizerMedianActivate, error) {
	event := new(MedianizerMedianActivate)
	if err := _Medianizer.contract.UnpackLog(event, "MedianActivate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MedianizerMedianVoidIterator is returned from FilterMedianVoid and is used to iterate over the raw logs and unpacked data for MedianVoid events raised by the Medianizer contract.
type MedianizerMedianVoidIterator struct {
	Event *MedianizerMedianVoid // Event containing the contract specifics and raw log

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
func (it *MedianizerMedianVoidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerMedianVoid)
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
		it.Event = new(MedianizerMedianVoid)
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
func (it *MedianizerMedianVoidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerMedianVoidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerMedianVoid represents a MedianVoid event raised by the Medianizer contract.
type MedianizerMedianVoid struct {
	Caller     common.Address
	Medianizer common.Address
	IsActive   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMedianVoid is a free log retrieval operation binding the contract event 0xc1c85319c7923477ac646408ce7d1255f85bf413c80258f2367de8f4fd51c950.
//
// Solidity: event MedianVoid(address caller, address medianizer, bool isActive)
func (_Medianizer *MedianizerFilterer) FilterMedianVoid(opts *bind.FilterOpts) (*MedianizerMedianVoidIterator, error) {

	logs, sub, err := _Medianizer.contract.FilterLogs(opts, "MedianVoid")
	if err != nil {
		return nil, err
	}
	return &MedianizerMedianVoidIterator{contract: _Medianizer.contract, event: "MedianVoid", logs: logs, sub: sub}, nil
}

// WatchMedianVoid is a free log subscription operation binding the contract event 0xc1c85319c7923477ac646408ce7d1255f85bf413c80258f2367de8f4fd51c950.
//
// Solidity: event MedianVoid(address caller, address medianizer, bool isActive)
func (_Medianizer *MedianizerFilterer) WatchMedianVoid(opts *bind.WatchOpts, sink chan<- *MedianizerMedianVoid) (event.Subscription, error) {

	logs, sub, err := _Medianizer.contract.WatchLogs(opts, "MedianVoid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerMedianVoid)
				if err := _Medianizer.contract.UnpackLog(event, "MedianVoid", log); err != nil {
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

// ParseMedianVoid is a log parse operation binding the contract event 0xc1c85319c7923477ac646408ce7d1255f85bf413c80258f2367de8f4fd51c950.
//
// Solidity: event MedianVoid(address caller, address medianizer, bool isActive)
func (_Medianizer *MedianizerFilterer) ParseMedianVoid(log types.Log) (*MedianizerMedianVoid, error) {
	event := new(MedianizerMedianVoid)
	if err := _Medianizer.contract.UnpackLog(event, "MedianVoid", log); err != nil {
		return nil, err
	}
	return event, nil
}
