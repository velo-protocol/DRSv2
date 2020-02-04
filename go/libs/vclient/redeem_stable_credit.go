package vclient

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"regexp"
)

type RedeemStableCreditInput struct {
	RedeemAmount    string
	AssetCode       string
}

func (i *RedeemStableCreditInput) Validate() error {
	redeemAmount, err := decimal.NewFromString(i.RedeemAmount)
	if err != nil {
		return errors.Wrap(err, "invalid RedeemAmount format")
	}
	if !redeemAmount.IsPositive() {
		return errors.New("RedeemAmount must be positive")
	}

	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,7}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	return nil
}

type RedeemStableCreditAbiInput struct {
	RedeemAmount    *big.Int
	AssetCode       string
}

func (i *RedeemStableCreditInput) ToAbiInput() *RedeemStableCreditAbiInput {
	redeemAmount, _ := utils.StringToAmount(i.RedeemAmount)
	return &RedeemStableCreditAbiInput{
		AssetCode:       i.AssetCode,
		RedeemAmount:    redeemAmount,
	}
}

func (c *Client) RedeemStableCredit(input *RedeemStableCreditInput) (bool, error) {
	err := input.Validate()
	if err != nil {
		return false, err
	}

	abiInput := input.ToAbiInput()

	redeemSuccess, err := c.contract.drs.Redeem(abiInput.RedeemAmount, abiInput.AssetCode)

	if err != nil {
		return false, err
	}

	return redeemSuccess, nil
}
