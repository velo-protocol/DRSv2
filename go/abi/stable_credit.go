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

// StableCreditABI is the input ABI used to generate the binding from.
const StableCreditABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_peggedCurrency\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_creditOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_collateralAssetCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_collateralAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_peggedValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"heartAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractICollateralAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralAssetCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heart\",\"outputs\":[{\"internalType\":\"contractIHeart\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peggedCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peggedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"approveCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralDetail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetCode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCollateralToReserve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StableCreditBin is the compiled bytecode used for deploying new contracts.
var StableCreditBin = "0x60806040523480156200001157600080fd5b5060405162001baf38038062001baf833981810160405260e08110156200003757600080fd5b81516020830151604080850151606086015160808701805193519597949692959194919392820192846401000000008211156200007357600080fd5b9083019060208201858111156200008957600080fd5b8251640100000000811182820188101715620000a457600080fd5b82525081516020918201929091019080838360005b83811015620000d3578181015183820152602001620000b9565b50505050905090810190601f168015620001015780820380516001836020036101000a031916815260200191505b50604090815260208281015192909101518551929450925084918291600791620001329160039190850190620001c3565b50815162000148906004906020850190620001c3565b5060058054600a80546001600160a01b03199081166001600160a01b039d8e161790915560089790975560099b909b5560068054909616978a1697909717909455505060079490945560ff1990951660ff90951694909417610100600160a81b03191661010092909316919091029190911790555062000268565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200020657805160ff191683800117855562000236565b8280016001018555821562000236579182015b828111156200023657825182559160200191906001019062000219565b506200024492915062000248565b5090565b6200026591905b808211156200024457600081556001016200024f565b90565b61193780620002786000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80634e841e3e116100de5780639dc29fac11610097578063aefee60e11610071578063aefee60e14610496578063d8dfeb451461049e578063dd62ed3e146104a6578063f58d1c94146104d45761018e565b80639dc29fac14610412578063a457c2d71461043e578063a9059cbb1461046a5761018e565b80634e841e3e146103af5780635312424c146103b75780635d1ca631146103bf57806370a08231146103c757806383044595146103ed57806395d89b411461040a5761018e565b80632664ecf91161014b578063395093511161012557806339509351146103265780633cf252a91461035257806340c10f191461035a57806349bf0f2d146103865761018e565b80632664ecf9146102cc5780632b83cccd146102d6578063313ce567146103085761018e565b806306fdde0314610193578063095ea7b31461021057806318160ddd146102505780631e2650de1461026a57806320dc407a1461028e57806323b872dd14610296575b600080fd5b61019b6104dc565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101d55781810151838201526020016101bd565b50505050905090810190601f1680156102025780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61023c6004803603604081101561022657600080fd5b506001600160a01b038135169060200135610572565b604080519115158252519081900360200190f35b61025861058f565b60408051918252519081900360200190f35b610272610595565b604080516001600160a01b039092168252519081900360200190f35b6102726105a4565b61023c600480360360608110156102ac57600080fd5b506001600160a01b038135811691602081013590911690604001356105b3565b6102d4610640565b005b6102d4600480360360608110156102ec57600080fd5b506001600160a01b038135169060208101359060400135610803565b61031061095b565b6040805160ff9092168252519081900360200190f35b61023c6004803603604081101561033c57600080fd5b506001600160a01b038135169060200135610964565b6102586109b8565b6102d46004803603604081101561037057600080fd5b506001600160a01b0381351690602001356109be565b61038e610a8e565b604080519283526001600160a01b0390911660208301528051918290030190f35b61019b610b1f565b610258610c4c565b610258610c52565b610258600480360360208110156103dd57600080fd5b50356001600160a01b0316610e5b565b61023c6004803603602081101561040357600080fd5b5035610e76565b61019b611045565b6102d46004803603604081101561042857600080fd5b506001600160a01b0381351690602001356110a6565b61023c6004803603604081101561045457600080fd5b506001600160a01b038135169060200135611176565b61023c6004803603604081101561048057600080fd5b506001600160a01b0381351690602001356111e4565b6102586111f8565b6102726111fe565b610258600480360360408110156104bc57600080fd5b506001600160a01b038135811691602001351661120d565b610272611238565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105685780601f1061053d57610100808354040283529160200191610568565b820191906000526020600020905b81548152906001019060200180831161054b57829003601f168201915b5050505050905090565b600061058661057f61124c565b8484611250565b50600192915050565b60025490565b600b546001600160a01b031681565b600a546001600160a01b031681565b60006105c084848461133c565b610636846105cc61124c565b6106318560405180606001604052806028815260200161184c602891396001600160a01b038a1660009081526001602052604081209061060a61124c565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61149816565b611250565b5060019392505050565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b15801561068957600080fd5b505afa15801561069d573d6000803e3d6000fd5b505050506040513d60208110156106b357600080fd5b50516001600160a01b031614610706576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b600654604080516370a0823160e01b815230600482015290516001600160a01b039092169163095ea7b391339184916370a08231916024808301926020929190829003018186803b15801561075a57600080fd5b505afa15801561076e573d6000803e3d6000fd5b505050506040513d602081101561078457600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091525160448083019260209291908290030181600087803b1580156107d557600080fd5b505af11580156107e9573d6000803e3d6000fd5b505050506040513d60208110156107ff57600080fd5b5050565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b15801561084c57600080fd5b505afa158015610860573d6000803e3d6000fd5b505050506040513d602081101561087657600080fd5b50516001600160a01b0316146108c9576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b6006546040805163a9059cbb60e01b81526001600160a01b038681166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561091f57600080fd5b505af1158015610933573d6000803e3d6000fd5b505050506040513d602081101561094957600080fd5b506109569050838361152f565b505050565b60055460ff1690565b600061058661097161124c565b84610631856001600061098261124c565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61162b16565b60085481565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b158015610a0757600080fd5b505afa158015610a1b573d6000803e3d6000fd5b505050506040513d6020811015610a3157600080fd5b50516001600160a01b031614610a84576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b6107ff828261168c565b600654604080516370a0823160e01b8152306004820152905160009283926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610ade57600080fd5b505afa158015610af2573d6000803e3d6000fd5b505050506040513d6020811015610b0857600080fd5b50516006549092506001600160a01b031690509091565b6060306001600160a01b03166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b158015610b5a57600080fd5b505afa158015610b6e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610b9757600080fd5b8101908080516040519392919084640100000000821115610bb757600080fd5b908301906020820185811115610bcc57600080fd5b8251640100000000811182820188101715610be657600080fd5b82525081516020918201929091019080838360005b83811015610c13578181015183820152602001610bfb565b50505050905090810190601f168015610c405780820380516001836020036101000a031916815260200191505b50604052505050905090565b60075481565b600073__Hasher________________________________63c2bebec7306001600160a01b03166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b158015610ca757600080fd5b505afa158015610cbb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ce457600080fd5b8101908080516040519392919084640100000000821115610d0457600080fd5b908301906020820185811115610d1957600080fd5b8251640100000000811182820188101715610d3357600080fd5b82525081516020918201929091019080838360005b83811015610d60578181015183820152602001610d48565b50505050905090810190601f168015610d8d5780820380516001836020036101000a031916815260200191505b506040525050506040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610de0578181015183820152602001610dc8565b50505050905090810190601f168015610e0d5780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610e2a57600080fd5b505af4158015610e3e573d6000803e3d6000fd5b505050506040513d6020811015610e5457600080fd5b5051905090565b6001600160a01b031660009081526020819052604090205490565b6000336001600160a01b0316600560019054906101000a90046001600160a01b03166001600160a01b03166370fa9a6d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ed057600080fd5b505afa158015610ee4573d6000803e3d6000fd5b505050506040513d6020811015610efa57600080fd5b50516001600160a01b031614610f4d576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b60065460055460408051631038421f60e31b815290516001600160a01b039384169363a9059cbb93610100900416916381c210f8916004808301926020929190829003018186803b158015610fa157600080fd5b505afa158015610fb5573d6000803e3d6000fd5b505050506040513d6020811015610fcb57600080fd5b5051604080516001600160e01b031960e085901b1681526001600160a01b039092166004830152602482018690525160448083019260209291908290030181600087803b15801561101b57600080fd5b505af115801561102f573d6000803e3d6000fd5b505050506040513d602081101561063657600080fd5b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105685780601f1061053d57610100808354040283529160200191610568565b600554604080516370fa9a6d60e01b81529051339261010090046001600160a01b0316916370fa9a6d916004808301926020929190829003018186803b1580156110ef57600080fd5b505afa158015611103573d6000803e3d6000fd5b505050506040513d602081101561111957600080fd5b50516001600160a01b03161461116c576040805162461bcd60e51b815260206004820152601360248201527263616c6c6572206973206e6f7420445253534360681b604482015290519081900360640190fd5b6107ff828261152f565b600061058661118361124c565b84610631856040518060600160405280602581526020016118de60259139600160006111ad61124c565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61149816565b60006105866111f161124c565b848461133c565b60095481565b6006546001600160a01b031681565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60055461010090046001600160a01b031681565b3390565b6001600160a01b0383166112955760405162461bcd60e51b81526004018080602001828103825260248152602001806118ba6024913960400191505060405180910390fd5b6001600160a01b0382166112da5760405162461bcd60e51b81526004018080602001828103825260228152602001806118046022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166113815760405162461bcd60e51b81526004018080602001828103825260258152602001806118956025913960400191505060405180910390fd5b6001600160a01b0382166113c65760405162461bcd60e51b81526004018080602001828103825260238152602001806117bf6023913960400191505060405180910390fd5b61140981604051806060016040528060268152602001611826602691396001600160a01b038616600090815260208190526040902054919063ffffffff61149816565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461143e908263ffffffff61162b16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156115275760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156114ec5781810151838201526020016114d4565b50505050905090810190601f1680156115195780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b0382166115745760405162461bcd60e51b81526004018080602001828103825260218152602001806118746021913960400191505060405180910390fd5b6115b7816040518060600160405280602281526020016117e2602291396001600160a01b038516600090815260208190526040902054919063ffffffff61149816565b6001600160a01b0383166000908152602081905260409020556002546115e3908263ffffffff61177c16565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600082820183811015611685576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382166116e7576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6002546116fa908263ffffffff61162b16565b6002556001600160a01b038216600090815260208190526040902054611726908263ffffffff61162b16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600061168583836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061149856fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820628251aa2848973fe51ab7f95c53b02d24bdf1596f11267a829833557f473f4564736f6c63430005100032"

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

// TransferCollateralToReserve is a paid mutator transaction binding the contract method 0x83044595.
//
// Solidity: function transferCollateralToReserve(uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactor) TransferCollateralToReserve(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.contract.Transact(opts, "transferCollateralToReserve", amount)
}

// TransferCollateralToReserve is a paid mutator transaction binding the contract method 0x83044595.
//
// Solidity: function transferCollateralToReserve(uint256 amount) returns(bool)
func (_StableCredit *StableCreditSession) TransferCollateralToReserve(amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.TransferCollateralToReserve(&_StableCredit.TransactOpts, amount)
}

// TransferCollateralToReserve is a paid mutator transaction binding the contract method 0x83044595.
//
// Solidity: function transferCollateralToReserve(uint256 amount) returns(bool)
func (_StableCredit *StableCreditTransactorSession) TransferCollateralToReserve(amount *big.Int) (*types.Transaction, error) {
	return _StableCredit.Contract.TransferCollateralToReserve(&_StableCredit.TransactOpts, amount)
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
