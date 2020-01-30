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

type SetupCreditInput struct {
	CollateralAssetCode string
	PeggedCurrency      string
	AssetCode           string
	PeggedValue         string
}

func (i *SetupCreditInput) Validate() error {
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,7}$`, i.CollateralAssetCode); !matched {
		return errors.New("invalid collateralAssetCode format")
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{3}$`, i.PeggedCurrency); !matched {
		return errors.New("invalid peggedCurrency format")
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,7}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	peggedValue, err := decimal.NewFromString(i.PeggedValue)
	if err != nil {
		return errors.Wrap(err, "invalid peggedValue format")
	}
	if !peggedValue.IsPositive() {
		return errors.New("peggedValue must be positive")
	}

	return nil
}

type SetupCreditAbiInput struct {
	CollateralAssetCode [32]byte
	PeggedCurrency      [32]byte
	AssetCode           string
	PeggedValue         *big.Int
}

func (i *SetupCreditInput) ToAbiInput() *SetupCreditAbiInput {
	peggedValue, _ := utils.StringToAmount(i.PeggedValue)
	return &SetupCreditAbiInput{
		CollateralAssetCode: utils.StringToByte32(i.CollateralAssetCode),
		PeggedCurrency:      utils.StringToByte32(i.PeggedCurrency),
		AssetCode:           i.AssetCode,
		PeggedValue:         peggedValue,
	}
}

type SetupCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *vabi.DigitalReserveSystemSetup
}

func (c *Client) SetupCredit(ctx context.Context, input *SetupCreditInput) (*SetupCreditOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	abiInput := input.ToAbiInput()
	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.drs.Setup(
		opt,
		abiInput.CollateralAssetCode,
		abiInput.PeggedCurrency,
		abiInput.AssetCode,
		abiInput.PeggedValue,
	)
	if err != nil {
		return nil, err
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	event, err := c.txHelper.ExtractSetupEvent("Setup", receipt.Logs[0])
	if err != nil {
		return nil, err
	}

	return &SetupCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   event,
	}, nil
}
