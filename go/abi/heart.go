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
const HeartABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collateralAssets\",\"outputs\":[{\"internalType\":\"contractICollateralAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collateralRatios\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collectedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditIssuanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIPrice\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"reserveFreeze\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveManager\",\"outputs\":[{\"internalType\":\"contractIRM\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stableCredits\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableCreditsLL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"llSize\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedPartners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReserveManager\",\"type\":\"address\"}],\"name\":\"setReserveManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserveManager\",\"outputs\":[{\"internalType\":\"contractIRM\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"newSeconds\",\"type\":\"uint32\"}],\"name\":\"setReserveFreeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getReserveFreeze\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDrsAddress\",\"type\":\"address\"}],\"name\":\"setDrsAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDrsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setCollateralAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollateralAsset\",\"outputs\":[{\"internalType\":\"contractICollateralAsset\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setCollateralRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setCreditIssuanceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreditIssuanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTrustedPartner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isTrustedPartner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIPrice\",\"name\":\"newPrice\",\"type\":\"address\"}],\"name\":\"addPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"}],\"name\":\"getPriceContract\",\"outputs\":[{\"internalType\":\"contractIPrice\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"}],\"name\":\"collectFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollectedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"newStableCredit\",\"type\":\"address\"}],\"name\":\"addStableCredit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stableCreditId\",\"type\":\"bytes32\"}],\"name\":\"getStableCreditById\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecentStableCredit\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stableCreditId\",\"type\":\"bytes32\"}],\"name\":\"getNextStableCredit\",\"outputs\":[{\"internalType\":\"contractIStableCredit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStableCreditCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setAllowedLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkId\",\"type\":\"bytes32\"}],\"name\":\"isLinkAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HeartBin is the compiled bytecode used for deploying new contracts.
var HeartBin = "0x608060405234801561001057600080fd5b50336000908152600b60209081526040909120805460ff19166001179055610042906006906115df610048821b17901c565b5061009b565b610050610089565b5060016000818152828201602090815260409182902080546001600160a01b03191690931790925580519182019052905460ff16815290565b60408051602081019091526000815290565b61171c806100aa6000396000f3fe608060405234801561001057600080fd5b50600436106102485760003560e01c806388bed3521161013b578063d77da9a2116100b8578063e402da4e1161007c578063e402da4e14610678578063e43581b81461069d578063f1e476a6146106c3578063f3685b7f146106e0578063f9af9ce01461070957610248565b8063d77da9a2146105c9578063db4d96d1146105ec578063dd06356c1461060f578063df7da7541461062c578063e12b39f51461065257610248565b8063bd09b85d116100ff578063bd09b85d1461052f578063c084770d1461054c578063c42cf53514610569578063d0d4212e1461058f578063d476f04b146105ac57610248565b806388bed35214610495578063aedb08d4146104c7578063bb004abc146104e4578063bb31146e146104ec578063bb33674a1461050957610248565b8063573ae02a116101c9578063716ec8ee1161018d578063716ec8ee1461042a57806378f4dce8146104325780637c0fa9a7146104685780638019155e1461048557806381c210f81461048d57610248565b8063573ae02a1461039b57806360846bc6146103b957806368698406146103d657806370ec42e8146103fc57806370fa9a6d1461042257610248565b806322042f9f1161021057806322042f9f146102ea578063295820801461030757806329fc9ec81461033357806333aec2fd1461036d578063449c3b741461037557610248565b8063016b3a111461024d5780630171af56146102675780630f09f97c1461028c5780630fb79eb2146102c55780631f6f2072146102e2575b600080fd5b610255610726565b60408051918252519081900360200190f35b61028a6004803603604081101561027d57600080fd5b508035906020013561072c565b005b6102a9600480360360208110156102a257600080fd5b50356107a7565b604080516001600160a01b039092168252519081900360200190f35b610255600480360360208110156102db57600080fd5b50356107c2565b6102556107d4565b6102a96004803603602081101561030057600080fd5b50356107da565b61028a6004803603604081101561031d57600080fd5b50803590602001356001600160a01b03166107f5565b6103596004803603602081101561034957600080fd5b50356001600160a01b031661092c565b604080519115158252519081900360200190f35b6102a9610941565b6103596004803603602081101561038b57600080fd5b50356001600160a01b031661095c565b6103a361097a565b6040805160ff9092168252519081900360200190f35b6102a9600480360360208110156103cf57600080fd5b5035610983565b61028a600480360360208110156103ec57600080fd5b50356001600160a01b031661099e565b6103596004803603602081101561041257600080fd5b50356001600160a01b0316610a06565b6102a9610a1b565b6102a9610a2a565b61044f6004803603602081101561044857600080fd5b5035610a39565b6040805163ffffffff9092168252519081900360200190f35b6102556004803603602081101561047e57600080fd5b5035610a51565b6103a3610a63565b6102a9610a6c565b61028a600480360360608110156104ab57600080fd5b508035906001600160a01b036020820135169060400135610a7b565b61028a600480360360208110156104dd57600080fd5b5035610af8565b6102a9610b41565b6102a96004803603602081101561050257600080fd5b5035610b50565b61028a6004803603602081101561051f57600080fd5b50356001600160a01b0316610b82565b61044f6004803603602081101561054557600080fd5b5035610ecb565b6103596004803603602081101561056257600080fd5b5035610ee3565b61028a6004803603602081101561057f57600080fd5b50356001600160a01b0316610ef8565b610255600480360360208110156105a557600080fd5b5035610f60565b6102a9600480360360208110156105c257600080fd5b5035610f72565b61028a600480360360408110156105df57600080fd5b5080359060200135610f8d565b61028a6004803603604081101561060257600080fd5b50803590602001356110f4565b6102a96004803603602081101561062557600080fd5b50356111b3565b61028a6004803603602081101561064257600080fd5b50356001600160a01b03166111ce565b61028a6004803603602081101561066857600080fd5b50356001600160a01b0316611234565b61028a6004803603604081101561068e57600080fd5b5080359060200135151561129a565b610359600480360360208110156106b357600080fd5b50356001600160a01b03166112fe565b6102a9600480360360208110156106d957600080fd5b503561131c565b61028a600480360360408110156106f657600080fd5b508035906020013563ffffffff16611337565b6102556004803603602081101561071f57600080fd5b503561135f565b60085490565b6000546001600160a01b031633146107755760405162461bcd60e51b81526004018080602001828103825260278152602001806116676027913960400191505060405180910390fd5b600081815260096020526040902054610794908363ffffffff61137116565b6000918252600960205260409091205550565b6003602052600090815260409020546001600160a01b031681565b60009081526004602052604090205490565b60085481565b6005602052600090815260409020546001600160a01b031681565b6107fe336112fe565b6108395760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b6001600160a01b038116610894576040805162461bcd60e51b815260206004820152601e60248201527f6e657750726963652061646472657373206d757374206e6f7420626520300000604482015290519081900360640190fd5b6000828152600160205260409020546001600160a01b0316156108fe576040805162461bcd60e51b815260206004820152601960248201527f50726963652068617320616c7265616479206578697374656400000000000000604482015290519081900360640190fd5b60009182526001602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b600a6020526000908152604090205460ff1681565b6000806109566006600163ffffffff6113d216565b91505090565b6001600160a01b03166000908152600a602052604090205460ff1690565b60065460ff1690565b6001602052600090815260409020546001600160a01b031681565b6109a7336112fe565b6109e25760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b6001600160a01b03166000908152600a60205260409020805460ff19166001179055565b600b6020526000908152604090205460ff1681565b6000546001600160a01b031690565b6000546001600160a01b031681565b600d6020526000908152604090205463ffffffff1681565b60046020526000908152604090205481565b60065460ff1681565b6002546001600160a01b031690565b610a84336112fe565b610abf5760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b600092835260036020908152604080852080546001600160a01b0319166001600160a01b03959095169490941790935560049052912055565b610b01336112fe565b610b3c5760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b600855565b6002546001600160a01b031681565b6000818152600560205260408120546001600160a01b031681610b7a60068363ffffffff6113d216565b949350505050565b6000546001600160a01b03163314610bcb5760405162461bcd60e51b81526004018080602001828103825260218152602001806116216021913960400191505060405180910390fd5b6001600160a01b038116610c105760405162461bcd60e51b81526004018080602001828103825260258152602001806116426025913960400191505060405180910390fd5b600073__Hasher________________________________63c2bebec7836001600160a01b0316634e841e3e6040518163ffffffff1660e01b815260040160006040518083038186803b158015610c6557600080fd5b505afa158015610c79573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ca257600080fd5b8101908080516040519392919084640100000000821115610cc257600080fd5b908301906020820185811115610cd757600080fd5b8251640100000000811182820188101715610cf157600080fd5b82525081516020918201929091019080838360005b83811015610d1e578181015183820152602001610d06565b50505050905090810190601f168015610d4b5780820380516001836020036101000a031916815260200191505b506040525050506040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d9e578181015183820152602001610d86565b50505050905090810190601f168015610dcb5780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610de857600080fd5b505af4158015610dfc573d6000803e3d6000fd5b505050506040513d6020811015610e1257600080fd5b50516000818152600560205260409020549091506001600160a01b031615610e81576040805162461bcd60e51b815260206004820181905260248201527f737461626c654372656469742068617320616c72656164792065786973746564604482015290519081900360640190fd5b600081815260056020526040902080546001600160a01b0319166001600160a01b038416179055610eb36006836113f5565b516006805460ff191660ff9092169190911790555050565b6000908152600d602052604090205463ffffffff1690565b6000908152600c602052604090205460ff1690565b610f01336112fe565b610f3c5760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b6001600160a01b03166000908152600b60205260409020805460ff19166001179055565b60096020526000908152604090205481565b6000908152600360205260409020546001600160a01b031690565b610f96336112fe565b610fd15760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b600082815260096020526040902054811115611034576040805162461bcd60e51b815260206004820152601e60248201527f616d6f756e74206d757374203c3d20746f20636f6c6c65637465644665650000604482015290519081900360640190fd5b600082815260036020908152604080832054815163a9059cbb60e01b81523360048201526024810186905291516001600160a01b039091169363a9059cbb93604480850194919392918390030190829087803b15801561109357600080fd5b505af11580156110a7573d6000803e3d6000fd5b505050506040513d60208110156110bd57600080fd5b50506000828152600960205260409020546110de908263ffffffff6114cf16565b6000928352600960205260409092209190915550565b6110fd336112fe565b6111385760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b6000828152600360205260409020546001600160a01b03166111a1576040805162461bcd60e51b815260206004820152601c60248201527f6173736574436f646520686173206e6f74206265656e20616464656400000000604482015290519081900360640190fd5b60009182526004602052604090912055565b6000908152600160205260409020546001600160a01b031690565b6111d7336112fe565b6112125760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b61123d336112fe565b6112785760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6112a3336112fe565b6112de5760405162461bcd60e51b815260040180806020018281038252605a81526020018061168e605a913960600191505060405180910390fd5b6000918252600c6020526040909120805460ff1916911515919091179055565b6001600160a01b03166000908152600b602052604090205460ff1690565b6000908152600560205260409020546001600160a01b031690565b6000918252600d6020526040909120805463ffffffff191663ffffffff909216919091179055565b60009081526009602052604090205490565b6000828201838110156113cb576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0380821660009081526001840160205260409020541692915050565b6113fd6115cd565b6114078383611511565b15611459576040805162461bcd60e51b815260206004820152601b60248201527f6164647220697320616c726561647920696e20746865206c6973740000000000604482015290519081900360640190fd5b50600160008181528382016020908152604080832080546001600160a01b039687168086528386208054989092166001600160a01b031998891617909155938590528054909516909217909355835460ff19811660ff918216909301811692909217938490558051928301905291909116815290565b60006113cb83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611536565b6001600160a01b03818116600090815260018401602052604090205416151592915050565b600081848411156115c55760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561158a578181015183820152602001611572565b50505050905090810190601f1680156115b75780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60408051602081019091526000815290565b6115e76115cd565b5060016000818152828201602090815260409182902080546001600160a01b03191690931790925580519182019052905460ff1681529056fe48656172742e6f6e6c794452533a2063616c6c6572206d757374206265204452536e6577537461626c654372656469742061646472657373206d757374206e6f7420626520306f6e6c792044525353432063616e207570646174652074686520636f6c6c65637465642066656548656172742e6f6e6c79476f7665726e6f723a20546865206d6573736167652073656e646572206973206e6f7420666f756e64206f7220646f6573206e6f7420686176652073756666696369656e74207065726d697373696f6ea265627a7a72315820acf1716c751c0babd18adc2dc776403bd627d68c3d5b431e90cfdc668bd332ad64736f6c63430005100032"

// DeployHeart deploys a new Ethereum contract, binding an instance of Heart to it.
func DeployHeart(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Heart, error) {
	parsed, err := abi.JSON(strings.NewReader(HeartABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HeartBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Heart{HeartCaller: HeartCaller{contract: contract}, HeartTransactor: HeartTransactor{contract: contract}, HeartFilterer: HeartFilterer{contract: contract}}, nil
}

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

// GetPriceContract is a free data retrieval call binding the contract method 0xdd06356c.
//
// Solidity: function getPriceContract(bytes32 linkId) constant returns(address)
func (_Heart *HeartCaller) GetPriceContract(opts *bind.CallOpts, linkId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "getPriceContract", linkId)
	return *ret0, err
}

// GetPriceContract is a free data retrieval call binding the contract method 0xdd06356c.
//
// Solidity: function getPriceContract(bytes32 linkId) constant returns(address)
func (_Heart *HeartSession) GetPriceContract(linkId [32]byte) (common.Address, error) {
	return _Heart.Contract.GetPriceContract(&_Heart.CallOpts, linkId)
}

// GetPriceContract is a free data retrieval call binding the contract method 0xdd06356c.
//
// Solidity: function getPriceContract(bytes32 linkId) constant returns(address)
func (_Heart *HeartCallerSession) GetPriceContract(linkId [32]byte) (common.Address, error) {
	return _Heart.Contract.GetPriceContract(&_Heart.CallOpts, linkId)
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

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(address)
func (_Heart *HeartCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Heart.contract.Call(opts, out, "prices", arg0)
	return *ret0, err
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(address)
func (_Heart *HeartSession) Prices(arg0 [32]byte) (common.Address, error) {
	return _Heart.Contract.Prices(&_Heart.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(address)
func (_Heart *HeartCallerSession) Prices(arg0 [32]byte) (common.Address, error) {
	return _Heart.Contract.Prices(&_Heart.CallOpts, arg0)
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

// AddPrice is a paid mutator transaction binding the contract method 0x29582080.
//
// Solidity: function addPrice(bytes32 linkId, address newPrice) returns()
func (_Heart *HeartTransactor) AddPrice(opts *bind.TransactOpts, linkId [32]byte, newPrice common.Address) (*types.Transaction, error) {
	return _Heart.contract.Transact(opts, "addPrice", linkId, newPrice)
}

// AddPrice is a paid mutator transaction binding the contract method 0x29582080.
//
// Solidity: function addPrice(bytes32 linkId, address newPrice) returns()
func (_Heart *HeartSession) AddPrice(linkId [32]byte, newPrice common.Address) (*types.Transaction, error) {
	return _Heart.Contract.AddPrice(&_Heart.TransactOpts, linkId, newPrice)
}

// AddPrice is a paid mutator transaction binding the contract method 0x29582080.
//
// Solidity: function addPrice(bytes32 linkId, address newPrice) returns()
func (_Heart *HeartTransactorSession) AddPrice(linkId [32]byte, newPrice common.Address) (*types.Transaction, error) {
	return _Heart.Contract.AddPrice(&_Heart.TransactOpts, linkId, newPrice)
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
