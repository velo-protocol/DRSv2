package vclient

import (
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"regexp"
	"strings"
)

// GetExchangeRateInput required input fields of get exchange rate
type GetExchangeRateInput struct {
	AssetCode string
}

// GetExchangeRateOutput output fields of get exchange rate
type GetExchangeRateOutput struct {
	AssetCode                     string
	CollateralAssetCode           string
	PriceInCollateralPerAssetUnit string
}

type GetExchangeRateAbiInput struct {
	AssetCode string
}

// Validation function for GetExchangeRate. Validates the required struct fields. It returns an error if any of the fields are
// invalid. Otherwise, it returns nil.
func (i *GetExchangeRateInput) Validate() error {
	if len(i.AssetCode) == 0 {
		return errors.New("assetCode must not be blank")
	}

	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,12}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	return nil
}

func (i *GetExchangeRateInput) ToAbiInput() GetExchangeRateAbiInput {
	return GetExchangeRateAbiInput{
		AssetCode: i.AssetCode,
	}
}

// GetExchangeRate calls GetExchange on Velo smart contract.
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
		msg := err.Error()
		switch {
		case strings.Contains(msg, "stableCredit not exist"):
			return nil, errors.Wrap(errors.Errorf("the stable credit %s does not exist", input.AssetCode), "smart contract call error")
		case strings.Contains(msg, "valid price not found") ||
			strings.Contains(msg, "invalid price") ||
			strings.Contains(msg, "price is not active") ||
			strings.Contains(msg, "price has an error"):
			return nil, errors.Wrap(errors.New("valid price not found"), "smart contract call error")
		default:
			return nil, err
		}
	}

	return &GetExchangeRateOutput{
		AssetCode:                     assetCode,
		CollateralAssetCode:           utils.Byte32ToString(collateralAssetCode),
		PriceInCollateralPerAssetUnit: utils.AmountToString(priceInCollateralPerAssetUnit),
	}, nil
}
