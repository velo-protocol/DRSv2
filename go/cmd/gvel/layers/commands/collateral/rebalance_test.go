package collateral_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
	"testing"
)

func TestCommandHandler_Rebalance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			RebalanceCollateral(&entity.RebalanceCollateralInput{
				Passphrase: "password",
			}).
			Return([]*entity.RebalanceCollateralOutput{
				{
					TxHash:              "0xAAAA",
					AssetCode:           "vUSD",
					CollateralAssetCode: "VELO",
					RequiredAmount:      "100.0000000",
					PresentAmount:       "120.0000000",
				},
				{
					TxHash:              "0xBBBB",
					AssetCode:           "vSGD",
					CollateralAssetCode: "VELO",
					RequiredAmount:      "230.0000000",
					PresentAmount:       "200.0000000",
				},
			}, nil)

		h.commandHandler.Rebalance(nil, nil)

		logs := h.loggerHook.AllEntries()
		assert.Equal(t, "Rebalancing completed.", logs[0].Message)

		resultMessages := strings.Split(logs[1].Message, "\n")
		assert.Equal(t, "🔗 Transaction Hash:", resultMessages[0])
		assert.Equal(t, "vUSD (VELO): 0xAAAA", resultMessages[1])
		assert.Equal(t, "vSGD (VELO): 0xBBBB", resultMessages[2])
	})
	t.Run("fail, logic.Rebalance returns an error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			RebalanceCollateral(&entity.RebalanceCollateralInput{
				Passphrase: "password",
			}).
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Rebalance(nil, nil)
		})
	})
}
