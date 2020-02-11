package logic_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"testing"
)

func TestLogic_SetEnv(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			SetCurrentEnv(constants.EnvMainNet).
			Return(nil)

		err := h.logic.SetEnv(&entity.SetEnvInput{
			Env: constants.EnvMainNet,
		})

		assert.NoError(t, err)
	})
	t.Run("fail, env is empty", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		err := h.logic.SetEnv(&entity.SetEnvInput{})

		assert.EqualError(t, err, "env must not be empty")
	})
}
