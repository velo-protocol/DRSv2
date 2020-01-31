package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"regexp"
)

type MintFromCollateralAmountInput struct {
	AssetCode        string
	CollateralAmount string
}

func (i *MintFromCollateralAmountInput) Validate() error {
	if i.AssetCode == "" {
		return errors.New("assetCode must not be blank")
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,7}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	if i.CollateralAmount == "" {
		return errors.New("collateralAmount must not be blank")
	}

	amount, err := decimal.NewFromString(i.CollateralAmount)
	if err != nil {
		return errors.Wrap(err, "invalid collateralAmount format")
	}
	if !utils.IsDecimalValid(amount) {
		return errors.New("invalid collateralAmount format")
	}
	if !amount.IsPositive() {
		return errors.New("collateralAmount must be greater than 0")
	}

	return nil
}

type MintFromCollateralAmountAbiInput struct {
	NetCollateralAmount *big.Int
	AssetCode           string
}

func (i *MintFromCollateralAmountInput) ToAbiInput() *MintFromCollateralAmountAbiInput {
	netAmount, _ := utils.StringToAmount(i.CollateralAmount)
	return &MintFromCollateralAmountAbiInput{
		NetCollateralAmount: netAmount,
		AssetCode:           i.AssetCode,
	}
}

type MintFromCollateralAmountCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *vabi.DigitalReserveSystemMint
}

func (c *Client) MintFromCollateralAmount(ctx context.Context, input *MintFromCollateralAmountInput) (*MintFromCollateralAmountCreditOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	abiInput := input.ToAbiInput()
	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.drs.MintFromCollateralAmount(opt, abiInput.NetCollateralAmount, abiInput.AssetCode)
	if err != nil {
		return nil, err
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	event, err := c.txHelper.ExtractMintEvent("Mint", receipt.Logs[12])
	if err != nil {
		return nil, err
	}

	return &MintFromCollateralAmountCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   event,
	}, nil
}
