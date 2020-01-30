package vclient

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
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
