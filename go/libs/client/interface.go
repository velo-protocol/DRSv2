package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/velo-protocol/DRSv2/go/abi"
	"math/big"
)

type Connection interface {
	bind.ContractBackend
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type DRSContract interface {
	Setup(opts *bind.TransactOpts, collateralAssetCode [32]byte, peggedCurrency [32]byte, assetCode string, peggedValue *big.Int) (*types.Transaction, error)
}

type HeartContract interface {
}

type TxHelper interface {
	ConfirmTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
	ExtractSetupEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemSetup, error)
}
