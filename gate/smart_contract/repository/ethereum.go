package repository

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"gitlab.com/velo-lab/core/gate/env"
	"gitlab.com/velo-lab/core/gate/smart_contract"
	abi "gitlab.com/velo-lab/core/smart_contract/abi/warp"
	"math/big"
)

type ethereumRepository struct {
	EthContractBackend bind.ContractBackend
}

func NewEthereumRepository(EthContractBackend bind.ContractBackend) smart_contract.Repository {
	return &ethereumRepository{
		EthContractBackend: EthContractBackend,
	}
}

func (e ethereumRepository) Mint(targetAccount string, amount *big.Int) (string, error) {
	contractInstance, err := abi.NewWarp(common.HexToAddress(env.WarpContractAddress), e.EthContractBackend)
	if err != nil {
		return "", errors.Wrap(err, "fail to get velo token smart contract on smart contract chain")
	}

	priv, err := crypto.HexToECDSA(env.GateETHPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "fail to convert owner pk to ecdsa")
	}

	txOpts := bind.NewKeyedTransactor(priv)
	txOpts.GasLimit = 4000000

	veloB32 := [32]byte{}
	copy(veloB32[:], []byte("VELO"))

	tx, err := contractInstance.Unlock(txOpts, common.HexToAddress(targetAccount), veloB32, amount)
	if err != nil {
		return "", errors.Wrap(err, "fail to execute unlock operation")
	}

	return tx.Hash().String(), nil
}
