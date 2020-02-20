package collateral_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
	"testing"
)

func TestCommandHandler_CollateralHealthCheck(t *testing.T) {

	var (
		collateralAssetCode      = "VELO"
		collateralAssetAddress   = "0x04"
		collateralPoolAmount     = "1.0000000"
		requiredCollateralAmount = "2.0000000"
	)

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			CollateralHealthCheck(&entity.CollateralHealthCheckInput{
				Passphrase: "password",
			}).
			Return([]*entity.CollateralHealthCheckOutput{
				{
					CollateralAssetAddress:   collateralAssetAddress,
					CollateralAssetCode:      collateralAssetCode,
					CollateralPoolAmount:     collateralPoolAmount,
					RequiredCollateralAmount: requiredCollateralAmount,
				},
			}, nil)

		h.commandHandler.HealthCheck(nil, nil)

		lines := strings.Split(h.tableLogHook.LastEntry().Message, "\n")

		assert.Contains(t, lines[1], "ASSET")
		assert.Contains(t, lines[1], "COLLATERAL POOL")
		assert.Contains(t, lines[1], "REQUIRED COLLATERAL")

		assert.Contains(t, lines[3], "VELO (0x04)")
		assert.Contains(t, lines[3], requiredCollateralAmount)
		assert.Contains(t, lines[3], collateralPoolAmount)
	})

	t.Run("fail, logic.CollateralHealthCheck returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input passphrase", nil).
			Return("password")

		h.mockLogic.EXPECT().
			CollateralHealthCheck(&entity.CollateralHealthCheckInput{
				Passphrase: "password",
			}).
			Return(nil, errors.New("error here"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.HealthCheck(nil, nil)
		})
	})
}
