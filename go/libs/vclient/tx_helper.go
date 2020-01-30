package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
	"strings"
)

type txHelper struct {
	conn     Connection
	drsAbi   *abi.ABI
	heartAbi *abi.ABI
}

func NewTxHelper(conn Connection) *txHelper {
	drsAbi, _ := abi.JSON(strings.NewReader(vabi.DigitalReserveSystemABI))
	heartAbi, _ := abi.JSON(strings.NewReader(vabi.HeartABI))
	return &txHelper{
		conn:     conn,
		drsAbi:   &drsAbi,
		heartAbi: &heartAbi,
	}
}

func (h *txHelper) ConfirmTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, h.conn, tx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to confirm transaction")
	}
	return receipt, nil
}

func (h *txHelper) ExtractSetupEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemSetup, error) {
	event := new(vabi.DigitalReserveSystemSetup)
	err := h.drsAbi.Unpack(event, eventName, log.Data)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read event log")
	}
	return event, nil
}
