package enode_listener

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/velo-lab/core/facilitator/env"
	"gitlab.com/velo-lab/core/facilitator/smart_contract/handler/enode_listener/model"
	drsABI "gitlab.com/velo-lab/core/smart_contract/abi/digital_reserve_system"
	"strings"
)

type ENodeListener struct {
}

func NewENodeListener(enode bind.ContractBackend) error {
	handler := &ENodeListener{}

	drsscAddress := common.HexToAddress(env.DRSSCAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			drsscAddress,
		},
	}

	logMintSig := []byte("Mint(bytes32,uint256,uint256,address,string,bytes32)")
	logMintSignHash := crypto.Keccak256Hash(logMintSig)

	drsscABI, err := abi.JSON(strings.NewReader(string(drsABI.DigitalReserveSystemABI)))
	if err != nil {
		return err
	}

	logs := make(chan types.Log)
	sub, err := enode.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logMintSignHash.Hex():
				var mintEvent model.LogMint

				err := drsscABI.Unpack(&mintEvent, "Mint", vLog.Data)
				if err != nil {
					return err
				}

				handler.mintEvent(mintEvent)
			}
		}
	}
}
