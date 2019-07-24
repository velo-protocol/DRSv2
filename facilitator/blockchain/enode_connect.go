package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/velo-lab/core/facilitator/env"
)

func ConnectENode() bind.ContractBackend {
	ethContractBackend, err := ethclient.Dial(env.ETHNode)
	if err != nil {
		panic(err)
	}

	return ethContractBackend
}
