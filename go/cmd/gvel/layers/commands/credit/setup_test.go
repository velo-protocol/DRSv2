package credit_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_Setup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code", nil).
			Return("vUSD")
		h.mockPrompt.EXPECT().
			RequestString("Please input collateral asset code", nil).
			Return("VELO")
		h.mockPrompt.EXPECT().
			RequestString("Please input pegged value", nil).
			Return("1")
		h.mockPrompt.EXPECT().
			RequestString("Please input pegged currency", nil).
			Return("USD")
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			SetupCredit(&entity.SetupCreditInput{
				Passphrase:          "password",
				AssetCode:           "vUSD",
				PeggedValue:         "1",
				PeggedCurrency:      "USD",
				CollateralAssetCode: "VELO",
			}).
			Return(&entity.SetupCreditOutput{
				TxHash:              h.mockTx.Hash().String(),
				CreditOwnerAddress:  "0x1",
				AssetCode:           "vUSD",
				AssetAddress:        "0x2",
				PeggedValue:         "1.0000000",
				PeggedCurrency:      "USD",
				CollateralAssetCode: "VELO",
			}, nil)

		h.commandHandler.Setup(nil, nil)

		logs := h.loggerHook.AllEntries()
		assert.Len(t, logs, 2)
		assert.Equal(t, "Stable credit vUSD (0x2) set up for account 0x1 successfully.", logs[0].Message)
		assert.Equal(t, "🔗 Transaction Hash: "+h.mockTx.Hash().String(), logs[1].Message)
	})

	t.Run("fail, logic.SetupCredit returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code", nil).
			Return("vUSD")
		h.mockPrompt.EXPECT().
			RequestString("Please input collateral asset code", nil).
			Return("VELO")
		h.mockPrompt.EXPECT().
			RequestString("Please input pegged value", nil).
			Return("1")
		h.mockPrompt.EXPECT().
			RequestString("Please input pegged currency", nil).
			Return("USD")
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			SetupCredit(&entity.SetupCreditInput{
				Passphrase:          "password",
				AssetCode:           "vUSD",
				PeggedValue:         "1",
				PeggedCurrency:      "USD",
				CollateralAssetCode: "VELO",
			}).
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Setup(nil, nil)
		})
	})
}
