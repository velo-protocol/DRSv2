package credit_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_GetExchange(t *testing.T) {

	var (
		assetCode                     = "vUSD"
		collateralAssetCode           = "VELO"
		priceInCollateralPerAssetUnit = "1.0000000"
	)

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of the stable credit", nil).
			Return(assetCode)

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			GetCreditExchange(&entity.GetCreditExchangeInput{
				AssetCode:  assetCode,
				Passphrase: "password",
			}).
			Return(&entity.GetCreditExchangeOutput{
				AssetCode:                     assetCode,
				CollateralAssetCode:           collateralAssetCode,
				PriceInCollateralPerAssetUnit: priceInCollateralPerAssetUnit,
			}, nil)

		h.commandHandler.GetExchange(nil, nil)

		logs := h.loggerHook.AllEntries()
		assert.Len(t, logs, 1)
		assert.Equal(t, "You will get 1.0000000 VELO for 1 vUSD.", logs[0].Message)
	})

	t.Run("fail, logic.GetExchange returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of the stable credit", nil).
			Return(assetCode)

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			GetCreditExchange(&entity.GetCreditExchangeInput{
				AssetCode:  assetCode,
				Passphrase: "password",
			}).
			Return(nil, errors.New("error here"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.GetExchange(nil, nil)
		})
	})
}
