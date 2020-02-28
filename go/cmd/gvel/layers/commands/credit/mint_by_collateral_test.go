package credit_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_MintByCollateral(t *testing.T) {

	var (
		assetCode        = "vUSD"
		collateralAmount = "100.0000000"
		mintAmount       = "100.0000000"
	)

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of credit to be minted", nil).
			Return(assetCode)
		h.mockPrompt.EXPECT().
			RequestString("Please input amount of collateral", nil).
			Return(collateralAmount)
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			MintCreditByCollateral(&entity.MintCreditByCollateralInput{
				AssetCode:        assetCode,
				CollateralAmount: collateralAmount,
				Passphrase:       "password",
			}).
			Return(&entity.MintCreditByCollateralOutput{
				AssetCode:          assetCode,
				StableCreditAmount: mintAmount,
				TxHash:             h.mockTx.Hash().String(),
			}, nil)

		h.commandHandler.MintByCollateral(nil, nil)

		logs := h.loggerHook.AllEntries()
		assert.Len(t, logs, 2)
		assert.Equal(t, "100.0000000 vUSD minted successfully.", logs[0].Message)
		assert.Equal(t, "🔗 Transaction Hash: "+h.mockTx.Hash().String(), logs[1].Message)
	})

	t.Run("fail, logic.MintByCollateral returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestString("Please input asset code of credit to be minted", nil).
			Return(assetCode)
		h.mockPrompt.EXPECT().
			RequestString("Please input amount of collateral", nil).
			Return(collateralAmount)
		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			MintCreditByCollateral(&entity.MintCreditByCollateralInput{
				AssetCode:        assetCode,
				CollateralAmount: collateralAmount,
				Passphrase:       "password",
			}).
			Return(nil, errors.New("error here"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.MintByCollateral(nil, nil)
		})
	})
}
