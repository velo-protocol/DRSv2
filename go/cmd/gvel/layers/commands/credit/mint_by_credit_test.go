package credit_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_MintByCredit(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of credit to be minted", nil).
			Return("vUSD")
		h.mockPrompt.EXPECT().
			RequestString("Please input amount of credit", nil).
			Return("1000")
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			MintCreditByCredit(&entity.MintCreditByCreditInput{
				AssetCode:    "vUSD",
				CreditAmount: "1000",
				Passphrase:   "password",
			}).
			Return(&entity.MintCreditByCreditOutput{
				AssetCode:           "vUSD",
				StableCreditAmount:  "1000.0000000",
				AssetAddress:        "0x1",
				CollateralAssetCode: "VELO",
				CollateralAmount:    "100.0000000",
				TxHash:              h.mockTx.Hash().String(),
			}, nil)

		h.commandHandler.MintByCredit(nil, nil)

		logs := h.loggerHook.AllEntries()
		assert.Len(t, logs, 2)
		assert.Equal(t, "1000.0000000 vUSD minted successfully using 100.0000000 VELO.", logs[0].Message)
		assert.Equal(t, "🔗 Transaction Hash: "+h.mockTx.Hash().String(), logs[1].Message)
	})

	t.Run("fail, logic.MintCreditByCredit returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of credit to be minted", nil).
			Return("vUSD")
		h.mockPrompt.EXPECT().
			RequestString("Please input amount of credit", nil).
			Return("1000")
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			MintCreditByCredit(&entity.MintCreditByCreditInput{
				AssetCode:    "vUSD",
				CreditAmount: "1000",
				Passphrase:   "password",
			}).
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.MintByCredit(nil, nil)
		})
	})
}
