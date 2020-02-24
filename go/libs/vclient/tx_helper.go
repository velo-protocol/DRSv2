package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"strings"
)

type txHelper struct {
	conn            Connection
	drsAbi          *abi.ABI
	heartAbi        *abi.ABI
	stableCreditAbi *abi.ABI
}

func NewTxHelper(conn Connection) *txHelper {
	drsAbi, _ := abi.JSON(strings.NewReader(vabi.DigitalReserveSystemABI))
	heartAbi, _ := abi.JSON(strings.NewReader(vabi.HeartABI))
	stableCreditAbi, _ := abi.JSON(strings.NewReader(vabi.StableCreditABI))
	return &txHelper{
		conn:            conn,
		drsAbi:          &drsAbi,
		heartAbi:        &heartAbi,
		stableCreditAbi: &stableCreditAbi,
	}
}

func (h *txHelper) ConfirmTx(ctx context.Context, tx *types.Transaction, from common.Address) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, h.conn, tx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to confirm transaction")
	}
	if receipt.Status == 0 {
		untrimmedBytes, err := h.conn.CallContract(ctx, ethereum.CallMsg{
			From:     from,
			To:       tx.To(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		}, receipt.BlockNumber)
		if err != nil {
			return nil, errors.Wrap(err, "fail to get revert message")
		}
		return nil, errors.New(utils.ParseRevertMessage(untrimmedBytes))
	}
	return receipt, nil
}

func (h *txHelper) ExtractSetupEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemSetup, error) {
	event := new(vabi.DigitalReserveSystemSetup)
	err := h.drsAbi.Unpack(event, eventName, log.Data)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read event log")
	}

	// extract indexed field
	if len(log.Topics) <= 1 {
		return nil, errors.New("fail to parse indexed param of an event")
	}
	event.CollateralAssetCode = utils.BytesToBytes32(log.Topics[1].Bytes())
	return event, nil
}

func (h *txHelper) StableCreditAssetCode(addr common.Address) (*string, *[32]byte, error) {
	stableCreditContract, err := vabi.NewStableCredit(addr, h.conn)
	if err != nil {
		return nil, nil, errors.Wrap(err, "fail to call stableCreditContract")
	}

	assetCode, err := stableCreditContract.AssetCode(nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "fail to call stableCreditContract.AssetCode")
	}

	stableCreditId, err := stableCreditContract.GetId(nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "fail to call stableCreditContract.GetId")
	}

	return &assetCode, &stableCreditId, nil
}

func (h *txHelper) ExtractMintEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemMint, error) {
	event := new(vabi.DigitalReserveSystemMint)
	err := h.drsAbi.Unpack(event, eventName, log.Data)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read event log")
	}

	// extract indexed field
	if len(log.Topics) < 3 {
		return nil, errors.New("fail to parse indexed param of an event")
	}

	event.AssetAddress = common.BytesToAddress(log.Topics[1].Bytes())
	event.CollateralAssetCode = utils.BytesToBytes32(log.Topics[2].Bytes())
	return event, nil
}

func (h *txHelper) ExtractRedeemEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemRedeem, error) {
	event := new(vabi.DigitalReserveSystemRedeem)
	err := h.drsAbi.Unpack(event, eventName, log.Data)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read event log")
	}

	// extract indexed field
	if len(log.Topics) < 3 {
		return nil, errors.New("fail to parse indexed param of an event")
	}
	event.CollateralAssetAddress = common.BytesToAddress(log.Topics[1].Bytes())
	event.CollateralAssetCode = utils.BytesToBytes32(log.Topics[2].Bytes())
	return event, nil
}

func (h *txHelper) ExtractRebalanceEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemRebalance, error) {
	event := new(vabi.DigitalReserveSystemRebalance)
	err := h.drsAbi.Unpack(event, eventName, log.Data)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read event log")
	}

	// extract indexed field
	if len(log.Topics) <= 1 {
		return nil, errors.New("fail to parse indexed param of an event")
	}
	event.CollateralAssetCode = utils.BytesToBytes32(log.Topics[1].Bytes())
	return event, nil
}
