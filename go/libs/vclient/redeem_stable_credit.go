package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	vabi "github.com/velo-protocol/DRSv2/go/abi"
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
		return errors.New("RedeemAmount must not be blank")
	}

	amount, err := decimal.NewFromString(i.RedeemAmount)
	if err != nil {
		return errors.Wrap(err, "invalid RedeemAmount format")
	}
	if !utils.IsDecimalValid(amount) {
		return errors.New("invalid RedeemAmount format")
	}
	if !amount.IsPositive() {
		return errors.New("RedeemAmount must be positive")
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

type RedeemStableCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *vabi.DigitalReserveSystemRedeem
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

	return &RedeemStableCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   event,
	}, nil
}
