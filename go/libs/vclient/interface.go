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
	MintFromCollateralAmount(opts *bind.TransactOpts, netCollateralAmount *big.Int, assetCode string) (*types.Transaction, error)
	GetExchange(opts *bind.CallOpts, assetCode string) (string, [32]byte, *big.Int, error)

	CollateralHealthCheck(opts *bind.CallOpts, assetCode string) ([32]byte, *big.Int, *big.Int, error)
}

type HeartContract interface {
	SetGovernor(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error)
	SetTrustedPartner(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error)

	GetStableCreditCount(opts *bind.CallOpts) (uint8, error)
	GetRecentStableCredit(opts *bind.CallOpts) (common.Address, error)
	GetNextStableCredit(opts *bind.CallOpts, stableCreditId [32]byte) (common.Address, error)
}

type StableCreditContract interface {
	StableCreditAssetCode(opts *bind.CallOpts) (string, error)
}

type TxHelper interface {
	ConfirmTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
	ExtractSetupEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemSetup, error)
	ExtractMintEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemMint, error)
	StableCreditAssetCode(addr common.Address) (*string, *[32]byte, error)
}
