package vclient

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"regexp"
)

//	function redeem(
//	uint256 amount,
//	string calldata assetCode
//) external returns (bool);

type RedeemStableCreditInput struct {
	RedeemAmount	string
	AssetCode		string
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
	RedeemAmount	*big.Int
	AssetCode		string
}

func (i *RedeemStableCreditInput) ToAbiInput() *RedeemStableCreditAbiInput {
	redeemAmount, _ := utils.StringToAmount(i.RedeemAmount)
	return &RedeemStableCreditAbiInput{
		AssetCode:			i.AssetCode,
		RedeemAmount:		redeemAmount,
	}
}

func (c *Client) RedeemStableCredit(input *RedeemStableCreditInput) (bool, error) {
	err := input.Validate()
	if err != nil {
		return false, err
	}

	assetCode, collateralAssetCode, priceInCollateralPerAssetUnit, err := c.contract.drs.GetExchange(
		nil,
		input.AssetCode,
	)
	if err != nil {
		return nil, err
	}

	return &RedeemStableCreditOutput{
		AssetCode:                     assetCode,
		CollateralAssetCode:           collateralAssetCode,
		PriceInCollateralPerAssetUnit: priceInCollateralPerAssetUnit,
	}, nil
}
