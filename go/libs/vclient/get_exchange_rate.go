package vclient

import (
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"regexp"
)

type GetExchangeRateInput struct {
	AssetCode string
}

type GetExchangeRateOutput struct {
	AssetCode                     string
	CollateralAssetCode           string
	PriceInCollateralPerAssetUnit string
}

type GetExchangeRateAbiInput struct {
	AssetCode string
}

func (i *GetExchangeRateInput) Validate() error {
	if len(i.AssetCode) == 0 {
		return errors.New("assetCode must not be blank")
	}

	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,7}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	return nil
}

func (i *GetExchangeRateInput) ToAbiInput() GetExchangeRateAbiInput {
	return GetExchangeRateAbiInput{
		AssetCode: i.AssetCode,
	}
}

func (c *Client) GetExchangeRate(input *GetExchangeRateInput) (*GetExchangeRateOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	assetCode, collateralAssetCode, priceInCollateralPerAssetUnit, err := c.contract.drs.GetExchange(
		nil,
		input.AssetCode,
	)
	if err != nil {
		return nil, err
	}

	return &GetExchangeRateOutput{
		AssetCode:                     assetCode,
		CollateralAssetCode:           utils.Byte32ToString(collateralAssetCode),
		PriceInCollateralPerAssetUnit: utils.AmountToString(priceInCollateralPerAssetUnit),
	}, nil
}
