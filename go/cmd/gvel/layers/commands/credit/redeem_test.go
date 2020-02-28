package credit_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_Redeem(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input amount of stable credit", nil).
			Return("104")
		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of credit to be redeemed", nil).
			Return("vUSD")
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			RedeemCredit(&entity.RedeemCreditInput{
				RedeemAmount: "104",
				AssetCode:    "vUSD",
				Passphrase:   "password",
			}).
			Return(&entity.RedeemCreditOutput{
				CollateralAmount:    "10.4000000",
				CollateralAssetCode: "VELO",
				TxHash:              h.mockTx.Hash().String(),
			}, nil)

		h.commandHandler.Redeem(nil, nil)

		logs := h.loggerHook.AllEntries()
		assert.Len(t, logs, 2)
		assert.Equal(t, "Redeemed successfully. You got 10.4000000 VELO.", logs[0].Message)
		assert.Equal(t, "🔗 Transaction Hash: "+h.mockTx.Hash().String(), logs[1].Message)
	})

	t.Run("fail, logic.RedeemCredit returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input amount of stable credit", nil).
			Return("104")
		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of credit to be redeemed", nil).
			Return("vUSD")
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			RedeemCredit(&entity.RedeemCreditInput{
				RedeemAmount: "104",
				AssetCode:    "vUSD",
				Passphrase:   "password",
			}).
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Redeem(nil, nil)
		})
	})
}
