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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FeederABI is the input ABI used to generate the binding from.
const FeederABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pf1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pf2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pf3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fiatCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_collateralCode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AcceptPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"CommitPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPriceFeed\",\"type\":\"address\"}],\"name\":\"UpdatePriceFeed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastOperationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operationCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeedTimeTol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeedTolInBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceTolInBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceUpdateCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valueTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"}],\"name\":\"startOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"}],\"name\":\"commitPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeder\",\"type\":\"address\"}],\"name\":\"updatePriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeederBin is the compiled bytecode used for deploying new contracts.
var FeederBin = "0x60806040526101f46009556064600a55603c600b819055600c556000600d819055600f556010805460ff1916905534801561003957600080fd5b50604051610e18380380610e18833981810160405260c081101561005c57600080fd5b508051602082015160408301516060840151608085015160a09095015160108054610100600160a81b0319166101006001600160a01b038089169190910291909117909155601180546001600160a01b031990811683881617909155601280548216838716179055601380549091169184169190911790556016869055601781905593949293919290916100ee610114565b600755505060006006555050600880546001600160a01b03191633179055506101189050565b4290565b610cf1806101276000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806364bc5204116100c3578063acfc827e1161007c578063acfc827e1461029b578063d1e61e24146102c7578063d8cf24fd146102cf578063dc9eaf38146102f0578063dcfe4c53146102f8578063e5693f411461030057610158565b806364bc52041461023b57806377680bb5146102435780637b8d56e31461024b5780638da5cb5b1461026e578063a99ae65114610276578063ab0ca0e11461029357610158565b8063427cb6fe11610115578063427cb6fe146101ef578063537f5373146102135780635587c4361461021b5780635623732e1461022357806357ef7f571461022b57806361efd4901461023357610158565b8063028a82f41461015d57806302fb0c5e1461018e578063053f14da1461019657806316d43a97146101c55780631f2698ab146101df57806339a3da4a146101e7575b600080fd5b61017a6004803603602081101561017357600080fd5b5035610308565b604080519115158252519081900360200190f35b61017a61075b565b61019e610764565b6040805193845260208401929092526001600160a01b031682820152519081900360600190f35b6101cd610779565b60408051918252519081900360200190f35b61017a61077f565b61019e610788565b6101f761079d565b604080516001600160a01b039092168252519081900360200190f35b6101cd6107ac565b6101cd6107b2565b6101cd6107b8565b6101cd6107be565b6101cd6107c4565b6101cd6107ca565b6101cd6107d0565b61017a6004803603604081101561026157600080fd5b50803590602001356107d6565b6101f76108c2565b61017a6004803603602081101561028c57600080fd5b50356108d1565b6101f76109a1565b61017a600480360360408110156102b157600080fd5b50803590602001356001600160a01b03166109b5565b61019e610acf565b6102d7610ae4565b6040805192835260208301919091528051918290030190f35b6101cd610aee565b6101cd610af4565b6101f7610afa565b60105460009061010090046001600160a01b031633148061033357506011546001600160a01b031633145b8061034857506012546001600160a01b031633145b61035157600080fd5b600061035b610b09565b60105490915060ff1680156103845750600c546007546103809163ffffffff610b0d16565b8110155b61038d57600080fd5b6000600d5460001415610470576006546103ae90859063ffffffff610b2316565b600954600654919250906103da906103ce8461271063ffffffff610b4516565b9063ffffffff610b6516565b116103ef576103ea848333610b7a565b61046b565b6040805160608101825285815260208082018590523391830182905260008781556001869055600280546001600160a01b031916841790558351928352908201528151849287927fe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9929081900390910190a3600d805460010190555b61074f565b600d546001141561062057600c546001546104909163ffffffff610b0d16565b8211156104d5576002546001600160a01b03163314156104ba576104b5848333610b7a565b6103ea565b6000546002546103ea919084906001600160a01b0316610b7a565b6002546001600160a01b03163314156104ed57600080fd5b600b546001548391610505919063ffffffff610b0d16565b10806105265750600b546001548391610524919063ffffffff610be316565b115b15610548576000546001546002546103ea9291906001600160a01b0316610b7a565b60005461055c90859063ffffffff610b2316565b6009546000549192509061057c906103ce8461271063ffffffff610b4516565b1161059e576000546001546002546103ea9291906001600160a01b0316610b7a565b6040805160608101825285815260208082018590523391830182905260038790556004859055600580546001600160a01b0319168317905582519182526001908201528151849287927fe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9929081900390910190a3600d8054600101905561074f565b600d546002141561074457600c5460015401821115610687576002546001600160a01b031633148061065c57506005546001600160a01b031633145b1561066c576104b5848333610b7a565b6003546005546103ea919084906001600160a01b0316610b7a565b6002546001600160a01b031633148015906106ad57506005546001600160a01b03163314155b6106b657600080fd5b600b5460015460009184916106d09163ffffffff610b0d16565b10806106f15750600b5460015484916106ef919063ffffffff610be316565b115b156106ff5750600054610724565b600354851415610710575083610724565b600054600354610721919087610bf5565b90505b60015460025461073e9183916001600160a01b0316610b7a565b5061074f565b600092505050610756565b6001925050505b919050565b60155460ff1681565b6006546007546008546001600160a01b031683565b600e5481565b60105460ff1681565b6000546001546002546001600160a01b031683565b6012546001600160a01b031681565b600a5481565b60175481565b600f5481565b600b5481565b60095481565b60165481565b600d5481565b6013546000906001600160a01b031633146108225760405162461bcd60e51b815260040180806020018281038252603a815260200180610c83603a913960400191505060405180910390fd5b60008361083757506009805490839055610878565b836001141561084e5750600a805490839055610878565b83600214156108655750600b805490839055610878565b83600314156101585750600c8054908390555b604080518581526020810183905280820185905290517fabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba9181900360600190a15060019392505050565b6013546001600160a01b031681565b60105460009061010090046001600160a01b03163314806108fc57506011546001600160a01b031633145b8061091157506012546001600160a01b031633145b61091a57600080fd5b6000610924610b09565b60105490915060ff161561093757600080fd5b426007556006839055600880546001600160a01b031916339081179091556010805460ff191660011790556040805191825251829185917f60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e9181900360200190a350600192915050565b60105461010090046001600160a01b031681565b6013546000906001600160a01b03163314610a015760405162461bcd60e51b815260040180806020018281038252603a815260200180610c83603a913960400191505060405180910390fd5b60038310610a0e57600080fd5b338284610a3a5760108054610100600160a81b0319166101006001600160a01b03841602179055610a7f565b8460011415610a6357601180546001600160a01b0319166001600160a01b038316179055610a7f565b601280546001600160a01b0319166001600160a01b0383161790555b604080516001600160a01b0380851682528316602082015281517f0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316929181900390910190a1506001949350505050565b6003546004546005546001600160a01b031683565b6006546007549091565b60145481565b600c5481565b6011546001600160a01b031681565b4290565b600082820183811015610b1c57fe5b9392505050565b6000818311610b3b57610b368284610be3565b610b1c565b610b1c8383610be3565b6000828202831580610b5f575082848281610b5c57fe5b04145b610b1c57fe5b600080828481610b7157fe5b04949350505050565b60068390556007829055600880546001600160a01b0383166001600160a01b031990911681179091556000600d556040805191825251839185917f60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e9181900360200190a3505050565b600082821115610bef57fe5b50900390565b6000610c07828563ffffffff610c6a16565b610c17858563ffffffff610c6a16565b186001600160f81b031916610c2d575082610b1c565b610c3d828463ffffffff610c6a16565b610c4d848663ffffffff610c6a16565b186001600160f81b031916610c63575081610b1c565b5092915050565b60008082841115610b1c5750600160f81b939250505056fe4665656465722e6f6e6c794f776e65723a2063616c6c6572206d75737420626520616e206f776e6572206f66207468697320636f6e7472616374a265627a7a72315820435ead9b92eb3314d4c19b11fb3bdc9eb8758f782ddf215904de325a3511913364736f6c634300050f0032"

// DeployFeeder deploys a new Ethereum contract, binding an instance of Feeder to it.
func DeployFeeder(auth *bind.TransactOpts, backend bind.ContractBackend, pf1 common.Address, pf2 common.Address, pf3 common.Address, _owner common.Address, _fiatCode [32]byte, _collateralCode [32]byte) (common.Address, *types.Transaction, *Feeder, error) {
	parsed, err := abi.JSON(strings.NewReader(FeederABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeederBin), backend, pf1, pf2, pf3, _owner, _fiatCode, _collateralCode)
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
// Solidity: function active() view returns(bool)
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
// Solidity: function active() view returns(bool)
func (_Feeder *FeederSession) Active() (bool, error) {
	return _Feeder.Contract.Active(&_Feeder.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Feeder *FeederCallerSession) Active() (bool, error) {
	return _Feeder.Contract.Active(&_Feeder.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() view returns(bytes32)
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
// Solidity: function collateralCode() view returns(bytes32)
func (_Feeder *FeederSession) CollateralCode() ([32]byte, error) {
	return _Feeder.Contract.CollateralCode(&_Feeder.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() view returns(bytes32)
func (_Feeder *FeederCallerSession) CollateralCode() ([32]byte, error) {
	return _Feeder.Contract.CollateralCode(&_Feeder.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() view returns(bytes32)
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
// Solidity: function fiatCode() view returns(bytes32)
func (_Feeder *FeederSession) FiatCode() ([32]byte, error) {
	return _Feeder.Contract.FiatCode(&_Feeder.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() view returns(bytes32)
func (_Feeder *FeederCallerSession) FiatCode() ([32]byte, error) {
	return _Feeder.Contract.FiatCode(&_Feeder.CallOpts)
}

// FirstPrice is a free data retrieval call binding the contract method 0x39a3da4a.
//
// Solidity: function firstPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCaller) FirstPrice(opts *bind.CallOpts) (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	ret := new(struct {
		PriceInWei   *big.Int
		TimeInSecond *big.Int
		Source       common.Address
	})
	out := ret
	err := _Feeder.contract.Call(opts, out, "firstPrice")
	return *ret, err
}

// FirstPrice is a free data retrieval call binding the contract method 0x39a3da4a.
//
// Solidity: function firstPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederSession) FirstPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.FirstPrice(&_Feeder.CallOpts)
}

// FirstPrice is a free data retrieval call binding the contract method 0x39a3da4a.
//
// Solidity: function firstPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCallerSession) FirstPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.FirstPrice(&_Feeder.CallOpts)
}

// GetLastPrice is a free data retrieval call binding the contract method 0xd8cf24fd.
//
// Solidity: function getLastPrice() view returns(uint256, uint256)
func (_Feeder *FeederCaller) GetLastPrice(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Feeder.contract.Call(opts, out, "getLastPrice")
	return *ret0, *ret1, err
}

// GetLastPrice is a free data retrieval call binding the contract method 0xd8cf24fd.
//
// Solidity: function getLastPrice() view returns(uint256, uint256)
func (_Feeder *FeederSession) GetLastPrice() (*big.Int, *big.Int, error) {
	return _Feeder.Contract.GetLastPrice(&_Feeder.CallOpts)
}

// GetLastPrice is a free data retrieval call binding the contract method 0xd8cf24fd.
//
// Solidity: function getLastPrice() view returns(uint256, uint256)
func (_Feeder *FeederCallerSession) GetLastPrice() (*big.Int, *big.Int, error) {
	return _Feeder.Contract.GetLastPrice(&_Feeder.CallOpts)
}

// LastOperationTime is a free data retrieval call binding the contract method 0x16d43a97.
//
// Solidity: function lastOperationTime() view returns(uint256)
func (_Feeder *FeederCaller) LastOperationTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "lastOperationTime")
	return *ret0, err
}

// LastOperationTime is a free data retrieval call binding the contract method 0x16d43a97.
//
// Solidity: function lastOperationTime() view returns(uint256)
func (_Feeder *FeederSession) LastOperationTime() (*big.Int, error) {
	return _Feeder.Contract.LastOperationTime(&_Feeder.CallOpts)
}

// LastOperationTime is a free data retrieval call binding the contract method 0x16d43a97.
//
// Solidity: function lastOperationTime() view returns(uint256)
func (_Feeder *FeederCallerSession) LastOperationTime() (*big.Int, error) {
	return _Feeder.Contract.LastOperationTime(&_Feeder.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCaller) LastPrice(opts *bind.CallOpts) (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	ret := new(struct {
		PriceInWei   *big.Int
		TimeInSecond *big.Int
		Source       common.Address
	})
	out := ret
	err := _Feeder.contract.Call(opts, out, "lastPrice")
	return *ret, err
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederSession) LastPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.LastPrice(&_Feeder.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCallerSession) LastPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.LastPrice(&_Feeder.CallOpts)
}

// NumOfPrices is a free data retrieval call binding the contract method 0x77680bb5.
//
// Solidity: function numOfPrices() view returns(uint256)
func (_Feeder *FeederCaller) NumOfPrices(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "numOfPrices")
	return *ret0, err
}

// NumOfPrices is a free data retrieval call binding the contract method 0x77680bb5.
//
// Solidity: function numOfPrices() view returns(uint256)
func (_Feeder *FeederSession) NumOfPrices() (*big.Int, error) {
	return _Feeder.Contract.NumOfPrices(&_Feeder.CallOpts)
}

// NumOfPrices is a free data retrieval call binding the contract method 0x77680bb5.
//
// Solidity: function numOfPrices() view returns(uint256)
func (_Feeder *FeederCallerSession) NumOfPrices() (*big.Int, error) {
	return _Feeder.Contract.NumOfPrices(&_Feeder.CallOpts)
}

// OperationCoolDown is a free data retrieval call binding the contract method 0x5623732e.
//
// Solidity: function operationCoolDown() view returns(uint256)
func (_Feeder *FeederCaller) OperationCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "operationCoolDown")
	return *ret0, err
}

// OperationCoolDown is a free data retrieval call binding the contract method 0x5623732e.
//
// Solidity: function operationCoolDown() view returns(uint256)
func (_Feeder *FeederSession) OperationCoolDown() (*big.Int, error) {
	return _Feeder.Contract.OperationCoolDown(&_Feeder.CallOpts)
}

// OperationCoolDown is a free data retrieval call binding the contract method 0x5623732e.
//
// Solidity: function operationCoolDown() view returns(uint256)
func (_Feeder *FeederCallerSession) OperationCoolDown() (*big.Int, error) {
	return _Feeder.Contract.OperationCoolDown(&_Feeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
// Solidity: function owner() view returns(address)
func (_Feeder *FeederSession) Owner() (common.Address, error) {
	return _Feeder.Contract.Owner(&_Feeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Feeder *FeederCallerSession) Owner() (common.Address, error) {
	return _Feeder.Contract.Owner(&_Feeder.CallOpts)
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() view returns(address)
func (_Feeder *FeederCaller) PriceFeed1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeed1")
	return *ret0, err
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() view returns(address)
func (_Feeder *FeederSession) PriceFeed1() (common.Address, error) {
	return _Feeder.Contract.PriceFeed1(&_Feeder.CallOpts)
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() view returns(address)
func (_Feeder *FeederCallerSession) PriceFeed1() (common.Address, error) {
	return _Feeder.Contract.PriceFeed1(&_Feeder.CallOpts)
}

// PriceFeed2 is a free data retrieval call binding the contract method 0xe5693f41.
//
// Solidity: function priceFeed2() view returns(address)
func (_Feeder *FeederCaller) PriceFeed2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeed2")
	return *ret0, err
}

// PriceFeed2 is a free data retrieval call binding the contract method 0xe5693f41.
//
// Solidity: function priceFeed2() view returns(address)
func (_Feeder *FeederSession) PriceFeed2() (common.Address, error) {
	return _Feeder.Contract.PriceFeed2(&_Feeder.CallOpts)
}

// PriceFeed2 is a free data retrieval call binding the contract method 0xe5693f41.
//
// Solidity: function priceFeed2() view returns(address)
func (_Feeder *FeederCallerSession) PriceFeed2() (common.Address, error) {
	return _Feeder.Contract.PriceFeed2(&_Feeder.CallOpts)
}

// PriceFeed3 is a free data retrieval call binding the contract method 0x427cb6fe.
//
// Solidity: function priceFeed3() view returns(address)
func (_Feeder *FeederCaller) PriceFeed3(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeed3")
	return *ret0, err
}

// PriceFeed3 is a free data retrieval call binding the contract method 0x427cb6fe.
//
// Solidity: function priceFeed3() view returns(address)
func (_Feeder *FeederSession) PriceFeed3() (common.Address, error) {
	return _Feeder.Contract.PriceFeed3(&_Feeder.CallOpts)
}

// PriceFeed3 is a free data retrieval call binding the contract method 0x427cb6fe.
//
// Solidity: function priceFeed3() view returns(address)
func (_Feeder *FeederCallerSession) PriceFeed3() (common.Address, error) {
	return _Feeder.Contract.PriceFeed3(&_Feeder.CallOpts)
}

// PriceFeedTimeTol is a free data retrieval call binding the contract method 0x57ef7f57.
//
// Solidity: function priceFeedTimeTol() view returns(uint256)
func (_Feeder *FeederCaller) PriceFeedTimeTol(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeedTimeTol")
	return *ret0, err
}

// PriceFeedTimeTol is a free data retrieval call binding the contract method 0x57ef7f57.
//
// Solidity: function priceFeedTimeTol() view returns(uint256)
func (_Feeder *FeederSession) PriceFeedTimeTol() (*big.Int, error) {
	return _Feeder.Contract.PriceFeedTimeTol(&_Feeder.CallOpts)
}

// PriceFeedTimeTol is a free data retrieval call binding the contract method 0x57ef7f57.
//
// Solidity: function priceFeedTimeTol() view returns(uint256)
func (_Feeder *FeederCallerSession) PriceFeedTimeTol() (*big.Int, error) {
	return _Feeder.Contract.PriceFeedTimeTol(&_Feeder.CallOpts)
}

// PriceFeedTolInBP is a free data retrieval call binding the contract method 0x537f5373.
//
// Solidity: function priceFeedTolInBP() view returns(uint256)
func (_Feeder *FeederCaller) PriceFeedTolInBP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeedTolInBP")
	return *ret0, err
}

// PriceFeedTolInBP is a free data retrieval call binding the contract method 0x537f5373.
//
// Solidity: function priceFeedTolInBP() view returns(uint256)
func (_Feeder *FeederSession) PriceFeedTolInBP() (*big.Int, error) {
	return _Feeder.Contract.PriceFeedTolInBP(&_Feeder.CallOpts)
}

// PriceFeedTolInBP is a free data retrieval call binding the contract method 0x537f5373.
//
// Solidity: function priceFeedTolInBP() view returns(uint256)
func (_Feeder *FeederCallerSession) PriceFeedTolInBP() (*big.Int, error) {
	return _Feeder.Contract.PriceFeedTolInBP(&_Feeder.CallOpts)
}

// PriceTolInBP is a free data retrieval call binding the contract method 0x61efd490.
//
// Solidity: function priceTolInBP() view returns(uint256)
func (_Feeder *FeederCaller) PriceTolInBP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceTolInBP")
	return *ret0, err
}

// PriceTolInBP is a free data retrieval call binding the contract method 0x61efd490.
//
// Solidity: function priceTolInBP() view returns(uint256)
func (_Feeder *FeederSession) PriceTolInBP() (*big.Int, error) {
	return _Feeder.Contract.PriceTolInBP(&_Feeder.CallOpts)
}

// PriceTolInBP is a free data retrieval call binding the contract method 0x61efd490.
//
// Solidity: function priceTolInBP() view returns(uint256)
func (_Feeder *FeederCallerSession) PriceTolInBP() (*big.Int, error) {
	return _Feeder.Contract.PriceTolInBP(&_Feeder.CallOpts)
}

// PriceUpdateCoolDown is a free data retrieval call binding the contract method 0xdcfe4c53.
//
// Solidity: function priceUpdateCoolDown() view returns(uint256)
func (_Feeder *FeederCaller) PriceUpdateCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceUpdateCoolDown")
	return *ret0, err
}

// PriceUpdateCoolDown is a free data retrieval call binding the contract method 0xdcfe4c53.
//
// Solidity: function priceUpdateCoolDown() view returns(uint256)
func (_Feeder *FeederSession) PriceUpdateCoolDown() (*big.Int, error) {
	return _Feeder.Contract.PriceUpdateCoolDown(&_Feeder.CallOpts)
}

// PriceUpdateCoolDown is a free data retrieval call binding the contract method 0xdcfe4c53.
//
// Solidity: function priceUpdateCoolDown() view returns(uint256)
func (_Feeder *FeederCallerSession) PriceUpdateCoolDown() (*big.Int, error) {
	return _Feeder.Contract.PriceUpdateCoolDown(&_Feeder.CallOpts)
}

// SecondPrice is a free data retrieval call binding the contract method 0xd1e61e24.
//
// Solidity: function secondPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCaller) SecondPrice(opts *bind.CallOpts) (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	ret := new(struct {
		PriceInWei   *big.Int
		TimeInSecond *big.Int
		Source       common.Address
	})
	out := ret
	err := _Feeder.contract.Call(opts, out, "secondPrice")
	return *ret, err
}

// SecondPrice is a free data retrieval call binding the contract method 0xd1e61e24.
//
// Solidity: function secondPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederSession) SecondPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.SecondPrice(&_Feeder.CallOpts)
}

// SecondPrice is a free data retrieval call binding the contract method 0xd1e61e24.
//
// Solidity: function secondPrice() view returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCallerSession) SecondPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.SecondPrice(&_Feeder.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Feeder *FeederCaller) Started(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "started")
	return *ret0, err
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Feeder *FeederSession) Started() (bool, error) {
	return _Feeder.Contract.Started(&_Feeder.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Feeder *FeederCallerSession) Started() (bool, error) {
	return _Feeder.Contract.Started(&_Feeder.CallOpts)
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() view returns(uint256)
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
// Solidity: function valueTimestamp() view returns(uint256)
func (_Feeder *FeederSession) ValueTimestamp() (*big.Int, error) {
	return _Feeder.Contract.ValueTimestamp(&_Feeder.CallOpts)
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() view returns(uint256)
func (_Feeder *FeederCallerSession) ValueTimestamp() (*big.Int, error) {
	return _Feeder.Contract.ValueTimestamp(&_Feeder.CallOpts)
}

// CommitPrice is a paid mutator transaction binding the contract method 0x028a82f4.
//
// Solidity: function commitPrice(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactor) CommitPrice(opts *bind.TransactOpts, priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "commitPrice", priceInWei)
}

// CommitPrice is a paid mutator transaction binding the contract method 0x028a82f4.
//
// Solidity: function commitPrice(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederSession) CommitPrice(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.CommitPrice(&_Feeder.TransactOpts, priceInWei)
}

// CommitPrice is a paid mutator transaction binding the contract method 0x028a82f4.
//
// Solidity: function commitPrice(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactorSession) CommitPrice(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.CommitPrice(&_Feeder.TransactOpts, priceInWei)
}

// SetValue is a paid mutator transaction binding the contract method 0x7b8d56e3.
//
// Solidity: function setValue(uint256 idx, uint256 newValue) returns(bool success)
func (_Feeder *FeederTransactor) SetValue(opts *bind.TransactOpts, idx *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "setValue", idx, newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x7b8d56e3.
//
// Solidity: function setValue(uint256 idx, uint256 newValue) returns(bool success)
func (_Feeder *FeederSession) SetValue(idx *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.SetValue(&_Feeder.TransactOpts, idx, newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x7b8d56e3.
//
// Solidity: function setValue(uint256 idx, uint256 newValue) returns(bool success)
func (_Feeder *FeederTransactorSession) SetValue(idx *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.SetValue(&_Feeder.TransactOpts, idx, newValue)
}

// StartOracle is a paid mutator transaction binding the contract method 0xa99ae651.
//
// Solidity: function startOracle(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactor) StartOracle(opts *bind.TransactOpts, priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "startOracle", priceInWei)
}

// StartOracle is a paid mutator transaction binding the contract method 0xa99ae651.
//
// Solidity: function startOracle(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederSession) StartOracle(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.StartOracle(&_Feeder.TransactOpts, priceInWei)
}

// StartOracle is a paid mutator transaction binding the contract method 0xa99ae651.
//
// Solidity: function startOracle(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactorSession) StartOracle(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.StartOracle(&_Feeder.TransactOpts, priceInWei)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xacfc827e.
//
// Solidity: function updatePriceFeed(uint256 index, address feeder) returns(bool)
func (_Feeder *FeederTransactor) UpdatePriceFeed(opts *bind.TransactOpts, index *big.Int, feeder common.Address) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "updatePriceFeed", index, feeder)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xacfc827e.
//
// Solidity: function updatePriceFeed(uint256 index, address feeder) returns(bool)
func (_Feeder *FeederSession) UpdatePriceFeed(index *big.Int, feeder common.Address) (*types.Transaction, error) {
	return _Feeder.Contract.UpdatePriceFeed(&_Feeder.TransactOpts, index, feeder)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xacfc827e.
//
// Solidity: function updatePriceFeed(uint256 index, address feeder) returns(bool)
func (_Feeder *FeederTransactorSession) UpdatePriceFeed(index *big.Int, feeder common.Address) (*types.Transaction, error) {
	return _Feeder.Contract.UpdatePriceFeed(&_Feeder.TransactOpts, index, feeder)
}

// FeederAcceptPriceIterator is returned from FilterAcceptPrice and is used to iterate over the raw logs and unpacked data for AcceptPrice events raised by the Feeder contract.
type FeederAcceptPriceIterator struct {
	Event *FeederAcceptPrice // Event containing the contract specifics and raw log

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
func (it *FeederAcceptPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederAcceptPrice)
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
		it.Event = new(FeederAcceptPrice)
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
func (it *FeederAcceptPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederAcceptPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederAcceptPrice represents a AcceptPrice event raised by the Feeder contract.
type FeederAcceptPrice struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Sender       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAcceptPrice is a free log retrieval operation binding the contract event 0x60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e.
//
// Solidity: event AcceptPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender)
func (_Feeder *FeederFilterer) FilterAcceptPrice(opts *bind.FilterOpts, priceInWei []*big.Int, timeInSecond []*big.Int) (*FeederAcceptPriceIterator, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "AcceptPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return &FeederAcceptPriceIterator{contract: _Feeder.contract, event: "AcceptPrice", logs: logs, sub: sub}, nil
}

// WatchAcceptPrice is a free log subscription operation binding the contract event 0x60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e.
//
// Solidity: event AcceptPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender)
func (_Feeder *FeederFilterer) WatchAcceptPrice(opts *bind.WatchOpts, sink chan<- *FeederAcceptPrice, priceInWei []*big.Int, timeInSecond []*big.Int) (event.Subscription, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "AcceptPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederAcceptPrice)
				if err := _Feeder.contract.UnpackLog(event, "AcceptPrice", log); err != nil {
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

// ParseAcceptPrice is a log parse operation binding the contract event 0x60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e.
//
// Solidity: event AcceptPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender)
func (_Feeder *FeederFilterer) ParseAcceptPrice(log types.Log) (*FeederAcceptPrice, error) {
	event := new(FeederAcceptPrice)
	if err := _Feeder.contract.UnpackLog(event, "AcceptPrice", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeederCommitPriceIterator is returned from FilterCommitPrice and is used to iterate over the raw logs and unpacked data for CommitPrice events raised by the Feeder contract.
type FeederCommitPriceIterator struct {
	Event *FeederCommitPrice // Event containing the contract specifics and raw log

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
func (it *FeederCommitPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederCommitPrice)
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
		it.Event = new(FeederCommitPrice)
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
func (it *FeederCommitPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederCommitPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederCommitPrice represents a CommitPrice event raised by the Feeder contract.
type FeederCommitPrice struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Sender       common.Address
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCommitPrice is a free log retrieval operation binding the contract event 0xe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9.
//
// Solidity: event CommitPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender, uint256 index)
func (_Feeder *FeederFilterer) FilterCommitPrice(opts *bind.FilterOpts, priceInWei []*big.Int, timeInSecond []*big.Int) (*FeederCommitPriceIterator, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "CommitPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return &FeederCommitPriceIterator{contract: _Feeder.contract, event: "CommitPrice", logs: logs, sub: sub}, nil
}

// WatchCommitPrice is a free log subscription operation binding the contract event 0xe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9.
//
// Solidity: event CommitPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender, uint256 index)
func (_Feeder *FeederFilterer) WatchCommitPrice(opts *bind.WatchOpts, sink chan<- *FeederCommitPrice, priceInWei []*big.Int, timeInSecond []*big.Int) (event.Subscription, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "CommitPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederCommitPrice)
				if err := _Feeder.contract.UnpackLog(event, "CommitPrice", log); err != nil {
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

// ParseCommitPrice is a log parse operation binding the contract event 0xe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9.
//
// Solidity: event CommitPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender, uint256 index)
func (_Feeder *FeederFilterer) ParseCommitPrice(log types.Log) (*FeederCommitPrice, error) {
	event := new(FeederCommitPrice)
	if err := _Feeder.contract.UnpackLog(event, "CommitPrice", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeederSetValueIterator is returned from FilterSetValue and is used to iterate over the raw logs and unpacked data for SetValue events raised by the Feeder contract.
type FeederSetValueIterator struct {
	Event *FeederSetValue // Event containing the contract specifics and raw log

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
func (it *FeederSetValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederSetValue)
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
		it.Event = new(FeederSetValue)
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
func (it *FeederSetValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederSetValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederSetValue represents a SetValue event raised by the Feeder contract.
type FeederSetValue struct {
	Index    *big.Int
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetValue is a free log retrieval operation binding the contract event 0xabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba.
//
// Solidity: event SetValue(uint256 index, uint256 oldValue, uint256 newValue)
func (_Feeder *FeederFilterer) FilterSetValue(opts *bind.FilterOpts) (*FeederSetValueIterator, error) {

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "SetValue")
	if err != nil {
		return nil, err
	}
	return &FeederSetValueIterator{contract: _Feeder.contract, event: "SetValue", logs: logs, sub: sub}, nil
}

// WatchSetValue is a free log subscription operation binding the contract event 0xabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba.
//
// Solidity: event SetValue(uint256 index, uint256 oldValue, uint256 newValue)
func (_Feeder *FeederFilterer) WatchSetValue(opts *bind.WatchOpts, sink chan<- *FeederSetValue) (event.Subscription, error) {

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "SetValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederSetValue)
				if err := _Feeder.contract.UnpackLog(event, "SetValue", log); err != nil {
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

// ParseSetValue is a log parse operation binding the contract event 0xabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba.
//
// Solidity: event SetValue(uint256 index, uint256 oldValue, uint256 newValue)
func (_Feeder *FeederFilterer) ParseSetValue(log types.Log) (*FeederSetValue, error) {
	event := new(FeederSetValue)
	if err := _Feeder.contract.UnpackLog(event, "SetValue", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeederUpdatePriceFeedIterator is returned from FilterUpdatePriceFeed and is used to iterate over the raw logs and unpacked data for UpdatePriceFeed events raised by the Feeder contract.
type FeederUpdatePriceFeedIterator struct {
	Event *FeederUpdatePriceFeed // Event containing the contract specifics and raw log

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
func (it *FeederUpdatePriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederUpdatePriceFeed)
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
		it.Event = new(FeederUpdatePriceFeed)
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
func (it *FeederUpdatePriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederUpdatePriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederUpdatePriceFeed represents a UpdatePriceFeed event raised by the Feeder contract.
type FeederUpdatePriceFeed struct {
	Updater      common.Address
	NewPriceFeed common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatePriceFeed is a free log retrieval operation binding the contract event 0x0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316.
//
// Solidity: event UpdatePriceFeed(address updater, address newPriceFeed)
func (_Feeder *FeederFilterer) FilterUpdatePriceFeed(opts *bind.FilterOpts) (*FeederUpdatePriceFeedIterator, error) {

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "UpdatePriceFeed")
	if err != nil {
		return nil, err
	}
	return &FeederUpdatePriceFeedIterator{contract: _Feeder.contract, event: "UpdatePriceFeed", logs: logs, sub: sub}, nil
}

// WatchUpdatePriceFeed is a free log subscription operation binding the contract event 0x0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316.
//
// Solidity: event UpdatePriceFeed(address updater, address newPriceFeed)
func (_Feeder *FeederFilterer) WatchUpdatePriceFeed(opts *bind.WatchOpts, sink chan<- *FeederUpdatePriceFeed) (event.Subscription, error) {

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "UpdatePriceFeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederUpdatePriceFeed)
				if err := _Feeder.contract.UnpackLog(event, "UpdatePriceFeed", log); err != nil {
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

// ParseUpdatePriceFeed is a log parse operation binding the contract event 0x0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316.
//
// Solidity: event UpdatePriceFeed(address updater, address newPriceFeed)
func (_Feeder *FeederFilterer) ParseUpdatePriceFeed(log types.Log) (*FeederUpdatePriceFeed, error) {
	event := new(FeederUpdatePriceFeed)
	if err := _Feeder.contract.UnpackLog(event, "UpdatePriceFeed", log); err != nil {
		return nil, err
	}
	return event, nil
}
