// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vabi

import (
	"math/big"
	"strings"

	ethereum "github.com/Evrynetlabs/evrynet-node"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi/bind"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/Evrynetlabs/evrynet-node/event"
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

// StableCreditABI is the input ABI used to generate the binding from.
const StableCreditABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_peggedCurrency\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_creditOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_collateralAssetCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_collateralAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_peggedValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"heartAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractICollateralAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralAssetCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heart\",\"outputs\":[{\"internalType\":\"contractIHeart\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peggedCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peggedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"approveCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralDetail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetCode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StableCreditBin is the compiled bytecode used for deploying new contracts.
var StableCreditBin = "0x60806040523480156200001157600080fd5b50604051620019a8380380620019a8833981810160405260e08110156200003757600080fd5b81516020830151604080850151606086015160808701805193519597949692959194919392820192846401000000008211156200007357600080fd5b9083019060208201858111156200008957600080fd5b8251640100000000811182820188101715620000a457600080fd5b82525081516020918201929091019080838360005b83811015620000d3578181015183820152602001620000b9565b50505050905090810190601f168015620001015780820380516001836020036101000a031916815260200191505b50604090815260208281015192909101518551929450925084918291600791620001329160039190850190620001c3565b50815162000148906004906020850190620001c3565b5060058054600a80546001600160a01b03199081166001600160a01b039d8e161790915560089790975560099b909b5560068054909616978a1697909717909455505060079490945560ff1990951660ff90951694909417610100600160a81b03191661010092909316919091029190911790555062000268565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200020657805160ff191683800117855562000236565b8280016001018555821562000236579182015b828111156200023657825182559160200191906001019062000219565b506200024492915062000248565b5090565b6200026591905b808211156200024457600081556001016200024f565b90565b61173080620002786000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c806349bf0f2d116100de5780639dc29fac11610097578063aefee60e11610071578063aefee60e1461045e578063d8dfeb4514610466578063dd62ed3e1461046e578063f58d1c941461049c57610173565b80639dc29fac146103da578063a457c2d714610406578063a9059cbb1461043257610173565b806349bf0f2d1461036b5780634e841e3e146103945780635312424c1461039c5780635d1ca631146103a457806370a08231146103ac57806395d89b41146103d257610173565b80632664ecf9116101305780632664ecf9146102b15780632b83cccd146102bb578063313ce567146102ed578063395093511461030b5780633cf252a91461033757806340c10f191461033f57610173565b806306fdde0314610178578063095ea7b3146101f557806318160ddd146102355780631e2650de1461024f57806320dc407a1461027357806323b872dd1461027b575b600080fd5b6101806104a4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ba5781810151838201526020016101a2565b50505050905090810190601f1680156101e75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102216004803603604081101561020b57600080fd5b506001600160a01b03813516906020013561053a565b604080519115158252519081900360200190f35b61023d610557565b60408051918252519081900360200190f35b61025761055d565b604080516001600160a01b039092168252519081900360200190f35b61025761056c565b6102216004803603606081101561029157600080fd5b506001600160a01b0381358116916020810135909116906040013561057b565b6102b9610608565b005b6102b9600480360360608110156102d157600080fd5b506001600160a01b0381351690602081013590604001356107cb565b6102f5610923565b6040805160ff9092168252519081900360200190f35b6102216004803603604081101561032157600080fd5b506001600160a01b03813516906020013561092c565b61023d610980565b6102b96004803603604081101561035557600080fd5b506001600160a01b038135169060200135610986565b610373610a56565b604080519283526001600160a01b0390911660208301528051918290030190f35b610180610ae7565b61023d610c14565b61023d610c1a565b61023d600480360360208110156103c257600080fd5b50356001600160a01b0316610e23565b610180610e3e565b6102b9600480360360408110156103f057600080fd5b506001600160a01b038135169060200135610e9f565b6102216004803603604081101561041c57600080fd5b506001600160a01b038135169060200135610f6f565b6102216004803603604081101561044857600080fd5b506001600160a01b038135169060200135610fdd565b61023d610ff1565b610257610ff7565b61023d6004803603604081101561048457600080fd5b506001600160a01b0381358116916020013516611006565b610257611031565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105305780601f1061050557610100808354040283529160200191610530565b820191906000526020600020905b81548152906001019060200180831161051357829003601f168201915b5050505050905090565b600061054e610547611045565b8484611049565b50600192915050565b60025490565b600b546001600160a01b031681565b600a546001600160a01b031681565b6000610588848484611135565b6105fe84610594611045565b6105f985604051806060016040528060288152602001611645602891396001600160a01b038a166000908152600160205260408120906105d2611045565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61129116565b611049565b5060019392505050565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b15801561065157600080fd5b505afa158015610665573d6000803e3d6000fd5b505050506040513d602081101561067b57600080fd5b50516001600160a01b0316146106ce576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b600654604080516370a0823160e01b815230600482015290516001600160a01b039092169163095ea7b391339184916370a08231916024808301926020929190829003018186803b15801561072257600080fd5b505afa158015610736573d6000803e3d6000fd5b505050506040513d602081101561074c57600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091525160448083019260209291908290030181600087803b15801561079d57600080fd5b505af11580156107b1573d6000803e3d6000fd5b505050506040513d60208110156107c757600080fd5b5050565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b15801561081457600080fd5b505afa158015610828573d6000803e3d6000fd5b505050506040513d602081101561083e57600080fd5b50516001600160a01b031614610891576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b6006546040805163a9059cbb60e01b81526001600160a01b038681166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156108e757600080fd5b505af11580156108fb573d6000803e3d6000fd5b505050506040513d602081101561091157600080fd5b5061091e90508383611328565b505050565b60055460ff1690565b600061054e610939611045565b846105f9856001600061094a611045565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61142416565b60085481565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b1580156109cf57600080fd5b505afa1580156109e3573d6000803e3d6000fd5b505050506040513d60208110156109f957600080fd5b50516001600160a01b031614610a4c576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b6107c78282611485565b600654604080516370a0823160e01b8152306004820152905160009283926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610aa657600080fd5b505afa158015610aba573d6000803e3d6000fd5b505050506040513d6020811015610ad057600080fd5b50516006549092506001600160a01b031690509091565b6060306001600160a01b03166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b158015610b2257600080fd5b505afa158015610b36573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610b5f57600080fd5b8101908080516040519392919084640100000000821115610b7f57600080fd5b908301906020820185811115610b9457600080fd5b8251640100000000811182820188101715610bae57600080fd5b82525081516020918201929091019080838360005b83811015610bdb578181015183820152602001610bc3565b50505050905090810190601f168015610c085780820380516001836020036101000a031916815260200191505b50604052505050905090565b60075481565b600073__Hasher________________________________63c2bebec7306001600160a01b03166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b158015610c6f57600080fd5b505afa158015610c83573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610cac57600080fd5b8101908080516040519392919084640100000000821115610ccc57600080fd5b908301906020820185811115610ce157600080fd5b8251640100000000811182820188101715610cfb57600080fd5b82525081516020918201929091019080838360005b83811015610d28578181015183820152602001610d10565b50505050905090810190601f168015610d555780820380516001836020036101000a031916815260200191505b506040525050506040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610da8578181015183820152602001610d90565b50505050905090810190601f168015610dd55780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610df257600080fd5b505af4158015610e06573d6000803e3d6000fd5b505050506040513d6020811015610e1c57600080fd5b5051905090565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105305780601f1061050557610100808354040283529160200191610530565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b158015610ee857600080fd5b505afa158015610efc573d6000803e3d6000fd5b505050506040513d6020811015610f1257600080fd5b50516001600160a01b031614610f65576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b6107c78282611328565b600061054e610f7c611045565b846105f9856040518060600160405280602581526020016116d76025913960016000610fa6611045565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61129116565b600061054e610fea611045565b8484611135565b60095481565b6006546001600160a01b031681565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60055461010090046001600160a01b031681565b3390565b6001600160a01b03831661108e5760405162461bcd60e51b81526004018080602001828103825260248152602001806116b36024913960400191505060405180910390fd5b6001600160a01b0382166110d35760405162461bcd60e51b81526004018080602001828103825260228152602001806115fd6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661117a5760405162461bcd60e51b815260040180806020018281038252602581526020018061168e6025913960400191505060405180910390fd5b6001600160a01b0382166111bf5760405162461bcd60e51b81526004018080602001828103825260238152602001806115b86023913960400191505060405180910390fd5b6112028160405180606001604052806026815260200161161f602691396001600160a01b038616600090815260208190526040902054919063ffffffff61129116565b6001600160a01b038085166000908152602081905260408082209390935590841681522054611237908263ffffffff61142416565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156113205760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156112e55781810151838201526020016112cd565b50505050905090810190601f1680156113125780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b03821661136d5760405162461bcd60e51b815260040180806020018281038252602181526020018061166d6021913960400191505060405180910390fd5b6113b0816040518060600160405280602281526020016115db602291396001600160a01b038516600090815260208190526040902054919063ffffffff61129116565b6001600160a01b0383166000908152602081905260409020556002546113dc908263ffffffff61157516565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60008282018381101561147e576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382166114e0576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6002546114f3908263ffffffff61142416565b6002556001600160a01b03821660009081526020819052604090205461151f908263ffffffff61142416565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600061147e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061129156fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820245d677aba950e387a73b81e971f5b0caab9cb4b24702ce79d742edbc70fc8b164736f6c634300050c0032"

// DeployStableCredit deploys a new Ethereum contract, binding an instance of StableCredit to it.
func DeployStableCredit(auth *bind.TransactOpts, backend bind.ContractBackend, _peggedCurrency [32]byte, _creditOwner common.Address, _collateralAssetCode [32]byte, _collateralAddress common.Address, _code string, _peggedValue *big.Int, heartAddr common.Address) (common.Address, *types.Transaction, *StableCredit, error) {
	parsed, err := abi.JSON(strings.NewReader(StableCreditABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StableCreditBin), backend, _peggedCurrency, _creditOwner, _collateralAssetCode, _collateralAddress, _code, _peggedValue, heartAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StableCredit{StableCreditCaller: StableCreditCaller{contract: contract}, StableCreditTransactor: StableCreditTransactor{contract: contract}, StableCreditFilterer: StableCreditFilterer{contract: contract}}, nil
}

// StableCredit is an auto generated Go binding around an Ethereum contract.
type StableCredit struct {
	StableCreditCaller     // Read-only binding to the contract
	StableCreditTransactor // Write-only binding to the contract
	StableCreditFilterer   // Log filterer for contract events
}

// StableCreditCaller is an auto generated read-only Go binding around an Ethereum contract.
type StableCreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableCreditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StableCreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableCreditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StableCreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableCreditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StableCreditSession struct {
	Contract     *StableCredit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StableCreditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StableCreditCallerSession struct {
	Contract *StableCreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StableCreditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StableCreditTransactorSession struct {
	Contract     *StableCreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StableCreditRaw is an auto generated low-level Go binding around an Ethereum contract.
type StableCreditRaw struct {
	Contract *StableCredit // Generic contract binding to access the raw methods on
}

// StableCreditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StableCreditCallerRaw struct {
	Contract *StableCreditCaller // Generic read-only contract binding to access the raw methods on
}

// StableCreditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StableCreditTransactorRaw struct {
	Contract *StableCreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStableCredit creates a new instance of StableCredit, bound to a specific deployed contract.
func NewStableCredit(address common.Address, backend bind.ContractBackend) (*StableCredit, error) {
	contract, err := bindStableCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StableCredit{StableCreditCaller: StableCreditCaller{contract: contract}, StableCreditTransactor: StableCreditTransactor{contract: contract}, StableCreditFilterer: StableCreditFilterer{contract: contract}}, nil
}

// NewStableCreditCaller creates a new read-only instance of StableCredit, bound to a specific deployed contract.
func NewStableCreditCaller(address common.Address, caller bind.ContractCaller) (*StableCreditCaller, error) {
	contract, err := bindStableCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StableCreditCaller{contract: contract}, nil
}

// NewStableCreditTransactor creates a new write-only instance of StableCredit, bound to a specific deployed contract.
func NewStableCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*StableCreditTransactor, error) {
	contract, err := bindStableCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StableCreditTransactor{contract: contract}, nil
}

// NewStableCreditFilterer creates a new log filterer instance of StableCredit, bound to a specific deployed contract.
func NewStableCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*StableCreditFilterer, error) {
	contract, err := bindStableCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StableCreditFilterer{contract: contract}, nil
}

// bindStableCredit binds a generic wrapper to an already deployed contract.
func bindStableCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StableCreditABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableCredit *StableCreditRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StableCredit.Contract.StableCreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableCredit *StableCreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableCredit.Contract.StableCreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableCredit *StableCreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableCredit.Contract.StableCreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableCredit *StableCreditCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StableCredit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableCredit *StableCreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableCredit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableCredit *StableCreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableCredit.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_StableCredit *StableCreditCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_StableCredit *StableCreditSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StableCredit.Contract.Allowance(&_StableCredit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_StableCredit *StableCreditCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StableCredit.Contract.Allowance(&_StableCredit.CallOpts, owner, spender)
}

// AssetCode is a free data retrieval call binding the contract method 0x4e841e3e.
//
// Solidity: function assetCode() constant returns(string)
func (_StableCredit *StableCreditCaller) AssetCode(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "assetCode")
	return *ret0, err
}

// AssetCode is a free data retrieval call binding the contract method 0x4e841e3e.
//
// Solidity: function assetCode() constant returns(string)
func (_StableCredit *StableCreditSession) AssetCode() (string, error) {
	return _StableCredit.Contract.AssetCode(&_StableCredit.CallOpts)
}

// AssetCode is a free data retrieval call binding the contract method 0x4e841e3e.
//
// Solidity: function assetCode() constant returns(string)
func (_StableCredit *StableCreditCallerSession) AssetCode() (string, error) {
	return _StableCredit.Contract.AssetCode(&_StableCredit.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_StableCredit *StableCreditCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_StableCredit *StableCreditSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StableCredit.Contract.BalanceOf(&_StableCredit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_StableCredit *StableCreditCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StableCredit.Contract.BalanceOf(&_StableCredit.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() constant returns(address)
func (_StableCredit *StableCreditCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "collateral")
	return *ret0, err
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() constant returns(address)
func (_StableCredit *StableCreditSession) Collateral() (common.Address, error) {
	return _StableCredit.Contract.Collateral(&_StableCredit.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() constant returns(address)
func (_StableCredit *StableCreditCallerSession) Collateral() (common.Address, error) {
	return _StableCredit.Contract.Collateral(&_StableCredit.CallOpts)
}

// CollateralAssetCode is a free data retrieval call binding the contract method 0x5312424c.
//
// Solidity: function collateralAssetCode() constant returns(bytes32)
func (_StableCredit *StableCreditCaller) CollateralAssetCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "collateralAssetCode")
	return *ret0, err
}

// CollateralAssetCode is a free data retrieval call binding the contract method 0x5312424c.
//
// Solidity: function collateralAssetCode() constant returns(bytes32)
func (_StableCredit *StableCreditSession) CollateralAssetCode() ([32]byte, error) {
	return _StableCredit.Contract.CollateralAssetCode(&_StableCredit.CallOpts)
}

// CollateralAssetCode is a free data retrieval call binding the contract method 0x5312424c.
//
// Solidity: function collateralAssetCode() constant returns(bytes32)
func (_StableCredit *StableCreditCallerSession) CollateralAssetCode() ([32]byte, error) {
	return _StableCredit.Contract.CollateralAssetCode(&_StableCredit.CallOpts)
}

// CreditOwner is a free data retrieval call binding the contract method 0x20dc407a.
//
// Solidity: function creditOwner() constant returns(address)
func (_StableCredit *StableCreditCaller) CreditOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "creditOwner")
	return *ret0, err
}

// CreditOwner is a free data retrieval call binding the contract method 0x20dc407a.
//
// Solidity: function creditOwner() constant returns(address)
func (_StableCredit *StableCreditSession) CreditOwner() (common.Address, error) {
	return _StableCredit.Contract.CreditOwner(&_StableCredit.CallOpts)
}

// CreditOwner is a free data retrieval call binding the contract method 0x20dc407a.
//
// Solidity: function creditOwner() constant returns(address)
func (_StableCredit *StableCreditCallerSession) CreditOwner() (common.Address, error) {
	return _StableCredit.Contract.CreditOwner(&_StableCredit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StableCredit *StableCreditCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StableCredit *StableCreditSession) Decimals() (uint8, error) {
	return _StableCredit.Contract.Decimals(&_StableCredit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StableCredit *StableCreditCallerSession) Decimals() (uint8, error) {
	return _StableCredit.Contract.Decimals(&_StableCredit.CallOpts)
}

// DrsAddress is a free data retrieval call binding the contract method 0x1e2650de.
//
// Solidity: function drsAddress() constant returns(address)
func (_StableCredit *StableCreditCaller) DrsAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "drsAddress")
	return *ret0, err
}

// DrsAddress is a free data retrieval call binding the contract method 0x1e2650de.
//
// Solidity: function drsAddress() constant returns(address)
func (_StableCredit *StableCreditSession) DrsAddress() (common.Address, error) {
	return _StableCredit.Contract.DrsAddress(&_StableCredit.CallOpts)
}

// DrsAddress is a free data retrieval call binding the contract method 0x1e2650de.
//
// Solidity: function drsAddress() constant returns(address)
func (_StableCredit *StableCreditCallerSession) DrsAddress() (common.Address, error) {
	return _StableCredit.Contract.DrsAddress(&_StableCredit.CallOpts)
}

// GetCollateralDetail is a free data retrieval call binding the contract method 0x49bf0f2d.
//
// Solidity: function getCollateralDetail() constant returns(uint256, address)
func (_StableCredit *StableCreditCaller) GetCollateralDetail(opts *bind.CallOpts) (*big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _StableCredit.contract.Call(opts, out, "getCollateralDetail")
	return *ret0, *ret1, err
}

// GetCollateralDetail is a free data retrieval call binding the contract method 0x49bf0f2d.
//
// Solidity: function getCollateralDetail() constant returns(uint256, address)
func (_StableCredit *StableCreditSession) GetCollateralDetail() (*big.Int, common.Address, error) {
	return _StableCredit.Contract.GetCollateralDetail(&_StableCredit.CallOpts)
}

// GetCollateralDetail is a free data retrieval call binding the contract method 0x49bf0f2d.
//
// Solidity: function getCollateralDetail() constant returns(uint256, address)
func (_StableCredit *StableCreditCallerSession) GetCollateralDetail() (*big.Int, common.Address, error) {
	return _StableCredit.Contract.GetCollateralDetail(&_StableCredit.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() constant returns(bytes32)
func (_StableCredit *StableCreditCaller) GetId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "getId")
	return *ret0, err
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() constant returns(bytes32)
func (_StableCredit *StableCreditSession) GetId() ([32]byte, error) {
	return _StableCredit.Contract.GetId(&_StableCredit.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() constant returns(bytes32)
func (_StableCredit *StableCreditCallerSession) GetId() ([32]byte, error) {
	return _StableCredit.Contract.GetId(&_StableCredit.CallOpts)
}

// Heart is a free data retrieval call binding the contract method 0xf58d1c94.
//
// Solidity: function heart() constant returns(address)
func (_StableCredit *StableCreditCaller) Heart(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "heart")
	return *ret0, err
}

// Heart is a free data retrieval call binding the contract method 0xf58d1c94.
//
// Solidity: function heart() constant returns(address)
func (_StableCredit *StableCreditSession) Heart() (common.Address, error) {
	return _StableCredit.Contract.Heart(&_StableCredit.CallOpts)
}

// Heart is a free data retrieval call binding the contract method 0xf58d1c94.
//
// Solidity: function heart() constant returns(address)
func (_StableCredit *StableCreditCallerSession) Heart() (common.Address, error) {
	return _StableCredit.Contract.Heart(&_StableCredit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StableCredit *StableCreditCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StableCredit *StableCreditSession) Name() (string, error) {
	return _StableCredit.Contract.Name(&_StableCredit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StableCredit *StableCreditCallerSession) Name() (string, error) {
	return _StableCredit.Contract.Name(&_StableCredit.CallOpts)
}

// PeggedCurrency is a free data retrieval call binding the contract method 0xaefee60e.
//
// Solidity: function peggedCurrency() constant returns(bytes32)
func (_StableCredit *StableCreditCaller) PeggedCurrency(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "peggedCurrency")
	return *ret0, err
}

// PeggedCurrency is a free data retrieval call binding the contract method 0xaefee60e.
//
// Solidity: function peggedCurrency() constant returns(bytes32)
func (_StableCredit *StableCreditSession) PeggedCurrency() ([32]byte, error) {
	return _StableCredit.Contract.PeggedCurrency(&_StableCredit.CallOpts)
}

// PeggedCurrency is a free data retrieval call binding the contract method 0xaefee60e.
//
// Solidity: function peggedCurrency() constant returns(bytes32)
func (_StableCredit *StableCreditCallerSession) PeggedCurrency() ([32]byte, error) {
	return _StableCredit.Contract.PeggedCurrency(&_StableCredit.CallOpts)
}

// PeggedValue is a free data retrieval call binding the contract method 0x3cf252a9.
//
// Solidity: function peggedValue() constant returns(uint256)
func (_StableCredit *StableCreditCaller) PeggedValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "peggedValue")
	return *ret0, err
}

// PeggedValue is a free data retrieval call binding the contract method 0x3cf252a9.
//
// Solidity: function peggedValue() constant returns(uint256)
func (_StableCredit *StableCreditSession) PeggedValue() (*big.Int, error) {
	return _StableCredit.Contract.PeggedValue(&_StableCredit.CallOpts)
}

// PeggedValue is a free data retrieval call binding the contract method 0x3cf252a9.
//
// Solidity: function peggedValue() constant returns(uint256)
func (_StableCredit *StableCreditCallerSession) PeggedValue() (*big.Int, error) {
	return _StableCredit.Contract.PeggedValue(&_StableCredit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StableCredit *StableCreditCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StableCredit *StableCreditSession) Symbol() (string, error) {
	return _StableCredit.Contract.Symbol(&_StableCredit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StableCredit *StableCreditCallerSession) Symbol() (string, error) {
	return _StableCredit.Contract.Symbol(&_StableCredit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StableCredit *StableCreditCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableCredit.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StableCredit *StableCreditSession) TotalSupply() (*big.Int, error) {
	return _StableCredit.Contract.TotalSupply(&_StableCredit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StableCredit *StableCreditCallerSession) TotalSupply() (*big.Int, error) {
	return _StableCredit.Contract.TotalSupply(&_StableCredit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StableCredit *StableCreditSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Approve(&_StableCredit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Approve(&_StableCredit.TransactOpts, spender, amount)
}

// ApproveCollateral is a paid mutator transaction binding the contract method 0x2664ecf9.
//
// Solidity: function approveCollateral() returns()
func (_StableCredit *StableCreditTransactor) ApproveCollateral(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "approveCollateral")
}

// ApproveCollateral is a paid mutator transaction binding the contract method 0x2664ecf9.
//
// Solidity: function approveCollateral() returns()
func (_StableCredit *StableCreditSession) ApproveCollateral() (*types.Transaction, error) {
	return _StableCredit.Contract.ApproveCollateral(&_StableCredit.TransactOpts)
}

// ApproveCollateral is a paid mutator transaction binding the contract method 0x2664ecf9.
//
// Solidity: function approveCollateral() returns()
func (_StableCredit *StableCreditTransactorSession) ApproveCollateral() (*types.Transaction, error) {
	return _StableCredit.Contract.ApproveCollateral(&_StableCredit.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tokenOwner, uint256 amount) returns()
func (_StableCredit *StableCreditTransactor) Burn(opts *bind.TransactOpts, tokenOwner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "burn", tokenOwner, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tokenOwner, uint256 amount) returns()
func (_StableCredit *StableCreditSession) Burn(tokenOwner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Burn(&_StableCredit.TransactOpts, tokenOwner, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tokenOwner, uint256 amount) returns()
func (_StableCredit *StableCreditTransactorSession) Burn(tokenOwner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Burn(&_StableCredit.TransactOpts, tokenOwner, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StableCredit *StableCreditTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StableCredit *StableCreditSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.DecreaseAllowance(&_StableCredit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StableCredit *StableCreditTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.DecreaseAllowance(&_StableCredit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StableCredit *StableCreditTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StableCredit *StableCreditSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.IncreaseAllowance(&_StableCredit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StableCredit *StableCreditTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.IncreaseAllowance(&_StableCredit.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_StableCredit *StableCreditTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "mint", recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_StableCredit *StableCreditSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Mint(&_StableCredit.TransactOpts, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_StableCredit *StableCreditTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Mint(&_StableCredit.TransactOpts, recipient, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address redeemer, uint256 burnAmount, uint256 returnAmount) returns()
func (_StableCredit *StableCreditTransactor) Redeem(opts *bind.TransactOpts, redeemer common.Address, burnAmount *big.Int, returnAmount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "redeem", redeemer, burnAmount, returnAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address redeemer, uint256 burnAmount, uint256 returnAmount) returns()
func (_StableCredit *StableCreditSession) Redeem(redeemer common.Address, burnAmount *big.Int, returnAmount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Redeem(&_StableCredit.TransactOpts, redeemer, burnAmount, returnAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address redeemer, uint256 burnAmount, uint256 returnAmount) returns()
func (_StableCredit *StableCreditTransactorSession) Redeem(redeemer common.Address, burnAmount *big.Int, returnAmount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Redeem(&_StableCredit.TransactOpts, redeemer, burnAmount, returnAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StableCredit *StableCreditSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Transfer(&_StableCredit.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.Transfer(&_StableCredit.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StableCredit *StableCreditSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.TransferFrom(&_StableCredit.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.TransferFrom(&_StableCredit.TransactOpts, sender, recipient, amount)
}

// StableCreditApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StableCredit contract.
type StableCreditApprovalIterator struct {
	Event *StableCreditApproval // Event containing the contract specifics and raw log

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
func (it *StableCreditApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableCreditApproval)
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
		it.Event = new(StableCreditApproval)
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
func (it *StableCreditApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableCreditApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableCreditApproval represents a Approval event raised by the StableCredit contract.
type StableCreditApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableCredit *StableCreditFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StableCreditApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableCredit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StableCreditApprovalIterator{contract: _StableCredit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableCredit *StableCreditFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StableCreditApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableCredit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableCreditApproval)
				if err := _StableCredit.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableCredit *StableCreditFilterer) ParseApproval(log types.Log) (*StableCreditApproval, error) {
	event := new(StableCreditApproval)
	if err := _StableCredit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableCreditTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StableCredit contract.
type StableCreditTransferIterator struct {
	Event *StableCreditTransfer // Event containing the contract specifics and raw log

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
func (it *StableCreditTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableCreditTransfer)
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
		it.Event = new(StableCreditTransfer)
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
func (it *StableCreditTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableCreditTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableCreditTransfer represents a Transfer event raised by the StableCredit contract.
type StableCreditTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableCredit *StableCreditFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StableCreditTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableCredit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StableCreditTransferIterator{contract: _StableCredit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableCredit *StableCreditFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StableCreditTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableCredit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableCreditTransfer)
				if err := _StableCredit.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableCredit *StableCreditFilterer) ParseTransfer(log types.Log) (*StableCreditTransfer, error) {
	event := new(StableCreditTransfer)
	if err := _StableCredit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
