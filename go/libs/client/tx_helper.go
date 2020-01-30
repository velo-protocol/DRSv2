package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type txHelper struct {
	conn Connection
}

func NewTxHelper(conn Connection) *txHelper {
	return &txHelper{
		conn: conn,
	}
}

func (h *txHelper) ConfirmTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, h.conn, tx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to confirm transaction")
	}
	return receipt, nil
}

func (h *txHelper) ExtractEventFromTx(contractAbi *abi.ABI, eventName string, log *types.Log) (interface{}, error) {
	event := new(interface{})
	err := contractAbi.Unpack(event, "Setup", log.Data)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read event log")
	}
	return event, nil
}
