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
const HeartABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collateralAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x0f09f97c\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditIssuanceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x1f6f2072\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stableCredits\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x22042f9f\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedPartners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x29fc9ec8\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"governor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70ec42e8\"},{\"constant\":true,\"inputs\":[],\"name\":\"drsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x716ec8ee\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"reserveFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x78f4dce8\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collateralRatios\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x7c0fa9a7\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableCreditsLL\",\"outputs\":[{\"name\":\"llSize\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8019155e\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeeders\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xadd36f68\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xbb004abc\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xd0d4212e\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"newReserveManager\",\"type\":\"address\"}],\"name\":\"setReserveManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xdf7da754\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserveManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x81c210f8\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"name\":\"newSeconds\",\"type\":\"uint32\"}],\"name\":\"setReserveFreeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf3685b7f\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getReserveFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xbd09b85d\"},{\"constant\":false,\"inputs\":[{\"name\":\"newDrsAddress\",\"type\":\"address\"}],\"name\":\"setDrsAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xe12b39f5\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDrsAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70fa9a6d\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setCollateralAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x88bed352\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollateralAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xd476f04b\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetCode\",\"type\":\"bytes32\"},{\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setCollateralRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xdb4d96d1\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollateralRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x0fb79eb2\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setCreditIssuanceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xaedb08d4\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreditIssuanceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x016b3a11\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTrustedPartner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x68698406\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xc42cf535\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isTrustedPartner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x449c3b74\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xe43581b8\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPriceFeeders\",\"type\":\"address\"}],\"name\":\"setPriceFeeders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x1070ff5d\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceFeeders\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x794e3db0\"},{\"constant\":false,\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"}],\"name\":\"collectFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x0171af56\"},{\"constant\":true,\"inputs\":[{\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"}],\"name\":\"getCollectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xf9af9ce0\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralAssetCode\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xd77da9a2\"},{\"constant\":false,\"inputs\":[{\"name\":\"newStableCredit\",\"type\":\"address\"}],\"name\":\"addStableCredit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xbb33674a\"},{\"constant\":true,\"inputs\":[{\"name\":\"stableCreditId\",\"type\":\"bytes32\"}],\"name\":\"getStableCreditById\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xf1e476a6\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecentStableCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x33aec2fd\"},{\"constant\":true,\"inputs\":[{\"name\":\"stableCreditId\",\"type\":\"bytes32\"}],\"name\":\"getNextStableCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xbb31146e\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStableCreditCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x573ae02a\"},{\"constant\":false,\"inputs\":[{\"name\":\"linkId\",\"type\":\"bytes32\"},{\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setAllowedLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xe402da4e\"},{\"constant\":true,\"inputs\":[{\"name\":\"linkId\",\"type\":\"bytes32\"}],\"name\":\"isLinkAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xc084770d\"}]"

// HeartBin is the compiled bytecode used for deploying new contracts.
var HeartBin = "0x608060405234801561001057600080fd5b50336000908152600b60205260409020805460ff1916600117905561004260066401000000006118d261004882021704565b5061009b565b610050610089565b506001600081815282820160209081526040918290208054600160a060020a03191690931790925580519182019052905460ff16815290565b60408051602081019091526000815290565b61199f806100aa6000396000f3fe6080604052600436106101d45763ffffffff60e060020a600035041663016b3a1181146101d95780630171af56146102005780630f09f97c146102325780630fb79eb2146102785780631070ff5d146102a25780631f6f2072146102d557806322042f9f146102ea57806329fc9ec81461031457806333aec2fd1461035b578063449c3b7414610370578063573ae02a146103a357806368698406146103ce57806370ec42e81461040157806370fa9a6d14610434578063716ec8ee1461044957806378f4dce81461045e578063794e3db0146104a15780637c0fa9a7146104b65780638019155e146104e057806381c210f8146104f557806388bed3521461050a578063add36f6814610549578063aedb08d41461055e578063bb004abc14610588578063bb31146e1461059d578063bb33674a146105c7578063bd09b85d146105fa578063c084770d14610624578063c42cf5351461064e578063d0d4212e14610681578063d476f04b146106ab578063d77da9a2146106d5578063db4d96d114610705578063df7da75414610735578063e12b39f514610768578063e402da4e1461079b578063e43581b8146107cd578063f1e476a614610800578063f3685b7f1461082a578063f9af9ce014610860575b600080fd5b3480156101e557600080fd5b506101ee61088a565b60408051918252519081900360200190f35b34801561020c57600080fd5b506102306004803603604081101561022357600080fd5b5080359060200135610890565b005b34801561023e57600080fd5b5061025c6004803603602081101561025557600080fd5b503561094a565b60408051600160a060020a039092168252519081900360200190f35b34801561028457600080fd5b506101ee6004803603602081101561029b57600080fd5b5035610965565b3480156102ae57600080fd5b50610230600480360360208110156102c557600080fd5b5035600160a060020a0316610977565b3480156102e157600080fd5b506101ee610a0e565b3480156102f657600080fd5b5061025c6004803603602081101561030d57600080fd5b5035610a14565b34801561032057600080fd5b506103476004803603602081101561033757600080fd5b5035600160a060020a0316610a2f565b604080519115158252519081900360200190f35b34801561036757600080fd5b5061025c610a44565b34801561037c57600080fd5b506103476004803603602081101561039357600080fd5b5035600160a060020a0316610a5f565b3480156103af57600080fd5b506103b8610a7d565b6040805160ff9092168252519081900360200190f35b3480156103da57600080fd5b50610230600480360360208110156103f157600080fd5b5035600160a060020a0316610a86565b34801561040d57600080fd5b506103476004803603602081101561042457600080fd5b5035600160a060020a0316610b1f565b34801561044057600080fd5b5061025c610b34565b34801561045557600080fd5b5061025c610b43565b34801561046a57600080fd5b506104886004803603602081101561048157600080fd5b5035610b52565b6040805163ffffffff9092168252519081900360200190f35b3480156104ad57600080fd5b5061025c610b6a565b3480156104c257600080fd5b506101ee600480360360208110156104d957600080fd5b5035610b79565b3480156104ec57600080fd5b506103b8610b8b565b34801561050157600080fd5b5061025c610b94565b34801561051657600080fd5b506102306004803603606081101561052d57600080fd5b50803590600160a060020a036020820135169060400135610ba3565b34801561055557600080fd5b5061025c610c51565b34801561056a57600080fd5b506102306004803603602081101561058157600080fd5b5035610c60565b34801561059457600080fd5b5061025c610cda565b3480156105a957600080fd5b5061025c600480360360208110156105c057600080fd5b5035610ce9565b3480156105d357600080fd5b50610230600480360360208110156105ea57600080fd5b5035600160a060020a0316610d1b565b34801561060657600080fd5b506104886004803603602081101561061d57600080fd5b5035611089565b34801561063057600080fd5b506103476004803603602081101561064757600080fd5b50356110a1565b34801561065a57600080fd5b506102306004803603602081101561067157600080fd5b5035600160a060020a03166110b6565b34801561068d57600080fd5b506101ee600480360360208110156106a457600080fd5b503561114f565b3480156106b757600080fd5b5061025c600480360360208110156106ce57600080fd5b5035611161565b3480156106e157600080fd5b50610230600480360360408110156106f857600080fd5b508035906020013561117c565b34801561071157600080fd5b506102306004803603604081101561072857600080fd5b5080359060200135611330565b34801561074157600080fd5b506102306004803603602081101561075857600080fd5b5035600160a060020a0316611425565b34801561077457600080fd5b506102306004803603602081101561078b57600080fd5b5035600160a060020a03166114bc565b3480156107a757600080fd5b50610230600480360360408110156107be57600080fd5b50803590602001351515611553565b3480156107d957600080fd5b50610347600480360360208110156107f057600080fd5b5035600160a060020a03166115e8565b34801561080c57600080fd5b5061025c6004803603602081101561082357600080fd5b5035611606565b34801561083657600080fd5b506102306004803603604081101561084d57600080fd5b508035906020013563ffffffff16611621565b34801561086c57600080fd5b506101ee6004803603602081101561088357600080fd5b5035611649565b60085490565b600054600160a060020a03163314610918576040805160e560020a62461bcd02815260206004820152602760248201527f6f6e6c792044525353432063616e207570646174652074686520636f6c6c656360448201527f7465642066656500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600081815260096020526040902054610937908363ffffffff61165b16565b6000918252600960205260409091205550565b600360205260009081526040902054600160a060020a031681565b60009081526004602052604090205490565b610980336115e8565b15156109ec576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60085481565b600560205260009081526040902054600160a060020a031681565b600a6020526000908152604090205460ff1681565b600080610a596006600163ffffffff6116bf16565b91505090565b600160a060020a03166000908152600a602052604090205460ff1690565b60065460ff1690565b610a8f336115e8565b1515610afb576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b600160a060020a03166000908152600a60205260409020805460ff19166001179055565b600b6020526000908152604090205460ff1681565b600054600160a060020a031690565b600054600160a060020a031681565b600d6020526000908152604090205463ffffffff1681565b600154600160a060020a031690565b60046020526000908152604090205481565b60065460ff1681565b600254600160a060020a031690565b610bac336115e8565b1515610c18576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b60009283526003602090815260408085208054600160a060020a031916600160a060020a03959095169490941790935560049052912055565b600154600160a060020a031681565b610c69336115e8565b1515610cd5576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b600855565b600254600160a060020a031681565b600081815260056020526040812054600160a060020a031681610d1360068363ffffffff6116bf16565b949350505050565b600054600160a060020a03163314610da3576040805160e560020a62461bcd02815260206004820152602160248201527f48656172742e6f6e6c794452533a2063616c6c6572206d75737420626520445260448201527f5300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0381161515610e29576040805160e560020a62461bcd02815260206004820152602560248201527f6e6577537461626c654372656469742061646472657373206d757374206e6f7460448201527f2062652030000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600073__Hasher________________________________63c2bebec783600160a060020a0316634e841e3e6040518163ffffffff1660e060020a02815260040160006040518083038186803b158015610e8157600080fd5b505afa158015610e95573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ebe57600080fd5b810190808051640100000000811115610ed657600080fd5b82016020810184811115610ee957600080fd5b8151640100000000811182820187101715610f0357600080fd5b505060405160e060020a63ffffffff871602815260206004820181815283516024840152835193965094508493506044909101919085019080838360005b83811015610f59578181015183820152602001610f41565b50505050905090810190601f168015610f865780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610fa357600080fd5b505af4158015610fb7573d6000803e3d6000fd5b505050506040513d6020811015610fcd57600080fd5b5051600081815260056020526040902054909150600160a060020a03161561103f576040805160e560020a62461bcd02815260206004820181905260248201527f737461626c654372656469742068617320616c72656164792065786973746564604482015290519081900360640190fd5b60008181526005602052604090208054600160a060020a031916600160a060020a0384161790556110716006836116e2565b516006805460ff191660ff9092169190911790555050565b6000908152600d602052604090205463ffffffff1690565b6000908152600c602052604090205460ff1690565b6110bf336115e8565b151561112b576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b60096020526000908152604090205481565b600090815260036020526040902054600160a060020a031690565b611185336115e8565b15156111f1576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b600082815260096020526040902054811115611257576040805160e560020a62461bcd02815260206004820152601e60248201527f616d6f756e74206d757374203c3d20746f20636f6c6c65637465644665650000604482015290519081900360640190fd5b60008281526003602090815260408083205481517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018690529151600160a060020a039091169363a9059cbb93604480850194919392918390030190829087803b1580156112cf57600080fd5b505af11580156112e3573d6000803e3d6000fd5b505050506040513d60208110156112f957600080fd5b505060008281526009602052604090205461131a908263ffffffff6117bf16565b6000928352600960205260409092209190915550565b611339336115e8565b15156113a5576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b600082815260036020526040902054600160a060020a03161515611413576040805160e560020a62461bcd02815260206004820152601c60248201527f6173736574436f646520686173206e6f74206265656e20616464656400000000604482015290519081900360640190fd5b60009182526004602052604090912055565b61142e336115e8565b151561149a576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b6114c5336115e8565b1515611531576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b61155c336115e8565b15156115c8576040805160e560020a62461bcd02815260206004820152605a602482015260008051602061191483398151915260448201526000805160206119348339815191526064820152600080516020611954833981519152608482015290519081900360a40190fd5b6000918252600c6020526040909120805460ff1916911515919091179055565b600160a060020a03166000908152600b602052604090205460ff1690565b600090815260056020526040902054600160a060020a031690565b6000918252600d6020526040909120805463ffffffff191663ffffffff909216919091179055565b60009081526009602052604090205490565b6000828201838110156116b8576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600160a060020a0380821660009081526001840160205260409020541692915050565b6116ea6118c0565b6116f48383611801565b15611749576040805160e560020a62461bcd02815260206004820152601b60248201527f6164647220697320616c726561647920696e20746865206c6973740000000000604482015290519081900360640190fd5b5060016000818152838201602090815260408083208054600160a060020a03968716808652838620805498909216600160a060020a031998891617909155938590528054909516909217909355835460ff19811660ff918216909301811692909217938490558051928301905291909116815290565b60006116b883836040805190810160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611826565b600160a060020a03818116600090815260018401602052604090205416151592915050565b600081848411156118b85760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561187d578181015183820152602001611865565b50505050905090810190601f1680156118aa5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60408051602081019091526000815290565b6118da6118c0565b506001600081815282820160209081526040918290208054600160a060020a03191690931790925580519182019052905460ff1681529056fe48656172742e6f6e6c79476f7665726e6f723a20546865206d6573736167652073656e646572206973206e6f7420666f756e64206f7220646f6573206e6f7420686176652073756666696369656e74207065726d697373696f6e000000000000a165627a7a723058203764a88e7023bdbf66a63898636313dce8ff0cff21fc74653b55572ef160fa6b0029"

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
