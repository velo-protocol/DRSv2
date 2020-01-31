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
const StableCreditABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drsAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peggedValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralAssetCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peggedCurrency\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heart\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_peggedCurrency\",\"type\":\"bytes32\"},{\"name\":\"_creditOwner\",\"type\":\"address\"},{\"name\":\"_collateralAssetCode\",\"type\":\"bytes32\"},{\"name\":\"_collateralAddress\",\"type\":\"address\"},{\"name\":\"_code\",\"type\":\"string\"},{\"name\":\"_peggedValue\",\"type\":\"uint256\"},{\"name\":\"heartAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"approveCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeemer\",\"type\":\"address\"},{\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralDetail\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetCode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StableCreditBin is the compiled bytecode used for deploying new contracts.
var StableCreditBin = "0x60806040523480156200001157600080fd5b5060405162001b0238038062001b02833981018060405260e08110156200003757600080fd5b8151602083015160408401516060850151608086018051949693959294919392830192916401000000008111156200006e57600080fd5b820160208101848111156200008257600080fd5b81516401000000008111828201871017156200009d57600080fd5b5050602080830151604090930151825192955092935084918291600791620000cc91600391908501906200015d565b508151620000e29060049060208501906200015d565b5060058054600a8054600160a060020a0319908116600160a060020a039d8e161790915560089790975560099b909b5560068054909616978a1697909717909455505060079490945560ff1990951660ff9095169490941761010060a860020a03191661010092909316919091029190911790555062000202565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a057805160ff1916838001178555620001d0565b82800160010185558215620001d0579182015b82811115620001d0578251825591602001919060010190620001b3565b50620001de929150620001e2565b5090565b620001ff91905b80821115620001de5760008155600101620001e9565b90565b6118f080620002126000396000f3fe60806040526004361061012f5763ffffffff60e060020a60003504166306fdde038114610134578063095ea7b3146101be57806318160ddd1461020b5780631e2650de1461023257806320dc407a1461026357806323b872dd146102785780632664ecf9146102bb5780632b83cccd146102d2578063313ce56714610311578063395093511461033c5780633cf252a91461037557806340c10f191461038a57806349bf0f2d146103c35780634e841e3e146103f95780635312424c1461040e5780635d1ca6311461042357806370a082311461043857806395d89b411461046b5780639dc29fac14610480578063a457c2d7146104b9578063a9059cbb146104f2578063aefee60e1461052b578063d8dfeb4514610540578063dd62ed3e14610555578063f58d1c9414610590575b600080fd5b34801561014057600080fd5b506101496105a5565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018357818101518382015260200161016b565b50505050905090810190601f1680156101b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ca57600080fd5b506101f7600480360360408110156101e157600080fd5b50600160a060020a03813516906020013561063b565b604080519115158252519081900360200190f35b34801561021757600080fd5b50610220610658565b60408051918252519081900360200190f35b34801561023e57600080fd5b5061024761065e565b60408051600160a060020a039092168252519081900360200190f35b34801561026f57600080fd5b5061024761066d565b34801561028457600080fd5b506101f76004803603606081101561029b57600080fd5b50600160a060020a0381358116916020810135909116906040013561067c565b3480156102c757600080fd5b506102d061075b565b005b3480156102de57600080fd5b506102d0600480360360608110156102f557600080fd5b50600160a060020a038135169060208101359060400135610933565b34801561031d57600080fd5b50610326610aa2565b6040805160ff9092168252519081900360200190f35b34801561034857600080fd5b506101f76004803603604081101561035f57600080fd5b50600160a060020a038135169060200135610aab565b34801561038157600080fd5b50610220610aff565b34801561039657600080fd5b506102d0600480360360408110156103ad57600080fd5b50600160a060020a038135169060200135610b05565b3480156103cf57600080fd5b506103d8610bd3565b60408051928352600160a060020a0390911660208301528051918290030190f35b34801561040557600080fd5b50610149610c7d565b34801561041a57600080fd5b50610220610d48565b34801561042f57600080fd5b50610220610d4e565b34801561044457600080fd5b506102206004803603602081101561045b57600080fd5b5035600160a060020a0316610ef9565b34801561047757600080fd5b50610149610f14565b34801561048c57600080fd5b506102d0600480360360408110156104a357600080fd5b50600160a060020a038135169060200135610f75565b3480156104c557600080fd5b506101f7600480360360408110156104dc57600080fd5b50600160a060020a038135169060200135611043565b3480156104fe57600080fd5b506101f76004803603604081101561051557600080fd5b50600160a060020a0381351690602001356110f5565b34801561053757600080fd5b50610220611109565b34801561054c57600080fd5b5061024761110f565b34801561056157600080fd5b506102206004803603604081101561057857600080fd5b50600160a060020a038135811691602001351661111e565b34801561059c57600080fd5b50610247611149565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106315780601f1061060657610100808354040283529160200191610631565b820191906000526020600020905b81548152906001019060200180831161061457829003601f168201915b5050505050905090565b600061064f61064861115d565b8484611161565b50600192915050565b60025490565b600b54600160a060020a031681565b600a54600160a060020a031681565b60006106898484846112ce565b6107518461069561115d565b61074c85606060405190810160405280602881526020017f45524332303a207472616e7366657220616d6f756e742065786365656473206181526020017f6c6c6f77616e6365000000000000000000000000000000000000000000000000815250600160008b600160a060020a0316600160a060020a03168152602001908152602001600020600061072561115d565b600160a060020a03168152602081019190915260400160002054919063ffffffff6114ef16565b611161565b5060019392505050565b6005546040805160e060020a6370fa9a6d028152905133926101009004600160a060020a0316916370fa9a6d916004808301926020929190829003018186803b1580156107a757600080fd5b505afa1580156107bb573d6000803e3d6000fd5b505050506040513d60208110156107d157600080fd5b5051600160a060020a03161461081f576040805160e560020a62461bcd02815260206004820152601360248201526000805160206118a5833981519152604482015290519081900360640190fd5b600654604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a039092169163095ea7b391339184916370a08231916024808301926020929190829003018186803b15801561088c57600080fd5b505afa1580156108a0573d6000803e3d6000fd5b505050506040513d60208110156108b657600080fd5b50516040805160e060020a63ffffffff8616028152600160a060020a03909316600484015260248301919091525160448083019260209291908290030181600087803b15801561090557600080fd5b505af1158015610919573d6000803e3d6000fd5b505050506040513d602081101561092f57600080fd5b5050565b6005546040805160e060020a6370fa9a6d028152905133926101009004600160a060020a0316916370fa9a6d916004808301926020929190829003018186803b15801561097f57600080fd5b505afa158015610993573d6000803e3d6000fd5b505050506040513d60208110156109a957600080fd5b5051600160a060020a0316146109f7576040805160e560020a62461bcd02815260206004820152601360248201526000805160206118a5833981519152604482015290519081900360640190fd5b600654604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015610a6657600080fd5b505af1158015610a7a573d6000803e3d6000fd5b505050506040513d6020811015610a9057600080fd5b50610a9d90508383611589565b505050565b60055460ff1690565b600061064f610ab861115d565b8461074c8560016000610ac961115d565b600160a060020a03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61170916565b60085481565b6005546040805160e060020a6370fa9a6d028152905133926101009004600160a060020a0316916370fa9a6d916004808301926020929190829003018186803b158015610b5157600080fd5b505afa158015610b65573d6000803e3d6000fd5b505050506040513d6020811015610b7b57600080fd5b5051600160a060020a031614610bc9576040805160e560020a62461bcd02815260206004820152601360248201526000805160206118a5833981519152604482015290519081900360640190fd5b61092f828261176d565b600654604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290516000928392600160a060020a03909116916370a0823191602480820192602092909190829003018186803b158015610c3c57600080fd5b505afa158015610c50573d6000803e3d6000fd5b505050506040513d6020811015610c6657600080fd5b5051600654909250600160a060020a031690509091565b606030600160a060020a03166306fdde036040518163ffffffff1660e060020a02815260040160006040518083038186803b158015610cbb57600080fd5b505afa158015610ccf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610cf857600080fd5b810190808051640100000000811115610d1057600080fd5b82016020810184811115610d2357600080fd5b8151640100000000811182820187101715610d3d57600080fd5b509094505050505090565b60075481565b600073__Hasher________________________________63c2bebec730600160a060020a03166306fdde036040518163ffffffff1660e060020a02815260040160006040518083038186803b158015610da657600080fd5b505afa158015610dba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610de357600080fd5b810190808051640100000000811115610dfb57600080fd5b82016020810184811115610e0e57600080fd5b8151640100000000811182820187101715610e2857600080fd5b505060405160e060020a63ffffffff871602815260206004820181815283516024840152835193965094508493506044909101919085019080838360005b83811015610e7e578181015183820152602001610e66565b50505050905090810190601f168015610eab5780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610ec857600080fd5b505af4158015610edc573d6000803e3d6000fd5b505050506040513d6020811015610ef257600080fd5b5051905090565b600160a060020a031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106315780601f1061060657610100808354040283529160200191610631565b6005546040805160e060020a6370fa9a6d028152905133926101009004600160a060020a0316916370fa9a6d916004808301926020929190829003018186803b158015610fc157600080fd5b505afa158015610fd5573d6000803e3d6000fd5b505050506040513d6020811015610feb57600080fd5b5051600160a060020a031614611039576040805160e560020a62461bcd02815260206004820152601360248201526000805160206118a5833981519152604482015290519081900360640190fd5b61092f8282611589565b600061064f61105061115d565b8461074c85606060405190810160405280602581526020017f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7781526020017f207a65726f000000000000000000000000000000000000000000000000000000815250600160006110be61115d565b600160a060020a03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6114ef16565b600061064f61110261115d565b84846112ce565b60095481565b600654600160a060020a031681565b600160a060020a03918216600090815260016020908152604080832093909416825291909152205490565b6005546101009004600160a060020a031681565b3390565b600160a060020a03831615156111e6576040805160e560020a62461bcd028152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a038216151561126c576040805160e560020a62461bcd02815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600160a060020a0383161515611354576040805160e560020a62461bcd02815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03821615156113da576040805160e560020a62461bcd02815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60408051606081018252602681527f45524332303a207472616e7366657220616d6f756e74206578636565647320626020808301919091527f616c616e6365000000000000000000000000000000000000000000000000000082840152600160a060020a038616600090815290819052919091205461146091839063ffffffff6114ef16565b600160a060020a038085166000908152602081905260408082209390935590841681522054611495908263ffffffff61170916565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156115815760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561154657818101518382015260200161152e565b50505050905090810190601f1680156115735780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600160a060020a038216151561160f576040805160e560020a62461bcd02815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60408051606081018252602281527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e6020808301919091527f636500000000000000000000000000000000000000000000000000000000000082840152600160a060020a038516600090815290819052919091205461169591839063ffffffff6114ef16565b600160a060020a0383166000908152602081905260409020556002546116c1908263ffffffff61186216565b600255604080518281529051600091600160a060020a038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600082820183811015611766576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600160a060020a03821615156117cd576040805160e560020a62461bcd02815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6002546117e0908263ffffffff61170916565b600255600160a060020a03821660009081526020819052604090205461180c908263ffffffff61170916565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600061176683836040805190810160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506114ef56fe63616c6c6572206973206e6f7420445253534300000000000000000000000000a165627a7a723058208369dd11091029ed39ac46994fa156c3903ab977d134d270efe55dff09cf53d70029"

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
