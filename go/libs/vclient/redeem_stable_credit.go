package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"regexp"
)

type RedeemStableCreditInput struct {
	RedeemAmount string
	AssetCode    string
}

func (i *RedeemStableCreditInput) Validate() error {
	// validate RedeemAmount
	if i.RedeemAmount == "" {
		return errors.New("redeemAmount must not be blank")
	}

	amount, err := decimal.NewFromString(i.RedeemAmount)
	if err != nil {
		return errors.Wrap(err, "invalid redeemAmount format")
	}
	if !utils.IsDecimalValid(amount) {
		return errors.New("invalid redeemAmount format")
	}
	if !amount.IsPositive() {
		return errors.New("redeemAmount must be positive")
	}

	// validate AssetCode
	if len(i.AssetCode) == 0 {
		return errors.New("assetCode must not be blank")
	}

	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,7}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	return nil
}

type RedeemStableCreditAbiInput struct {
	RedeemAmount *big.Int
	AssetCode    string
}

func (i *RedeemStableCreditInput) ToAbiInput() *RedeemStableCreditAbiInput {
	redeemAmount, _ := utils.StringToAmount(i.RedeemAmount)
	return &RedeemStableCreditAbiInput{
		AssetCode:    i.AssetCode,
		RedeemAmount: redeemAmount,
	}
}

type RedeemStableCreditEvent struct {
	// all fields must be string formatted
	AssetCode              string
	StableCreditAmount     *big.Int
	CollateralAssetAddress common.Address
	CollateralAssetCode    [32]byte
	CollateralAmount       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

type RedeemStableCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *RedeemStableCreditEvent
	// DigitalReserveSystemRedeem represents a Redeem event raised by the DigitalReserveSystem contract.
	//type DigitalReserveSystemRedeem struct {
	//	AssetCode              string
	//	StableCreditAmount     *big.Int
	//	CollateralAssetAddress common.Address
	//	CollateralAssetCode    [32]byte
	//	CollateralAmount       *big.Int
	//	Raw                    types.Log // Blockchain specific contextual infos
	//}
}

func (c *Client) RedeemStableCredit(ctx context.Context, input *RedeemStableCreditInput) (*RedeemStableCreditOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	abiInput := input.ToAbiInput()
	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit

	tx, err := c.contract.drs.Redeem(
		opt,
		abiInput.RedeemAmount,
		abiInput.AssetCode,
	)

	receipt, err := c.txHelper.ConfirmTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	event, err := c.txHelper.ExtractRedeemEvent("Redeem", receipt.Logs[0])
	if err != nil {
		return nil, err
	}

	// TODO convert from vabi.DigitalReserveSystemRedeem to RedeemStableCreditEvent
	redeemStableCreditEvent := &RedeemStableCreditEvent{
		AssetCode:event.AssetCode,
	}
	return &RedeemStableCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   redeemStableCreditEvent,
	}, nil
}
