package logic_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogic_Init(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			InitSharedConfig("~/.gvel").
			Return(nil)
		h.mockConfiguration.EXPECT().
			InitEnvBasedConfig("~/.gvel", "testnet").
			Return(nil)
		h.mockConfiguration.EXPECT().
			InitEnvBasedConfig("~/.gvel", "mainnet").
			Return(nil)

		err := h.logic.Init("~/.gvel")
		assert.NoError(t, err)
	})

	t.Run("fail, AppConfig.InitSharedConfig returns an error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			InitSharedConfig("~/.gvel").
			Return(errors.New("some error has occurred"))

		err := h.logic.Init("~/.gvel")
		assert.Error(t, err)
	})

	t.Run("fail, AppConfig.InitEnvBasedConfig for testnet returns an error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			InitSharedConfig("~/.gvel").
			Return(nil)
		h.mockConfiguration.EXPECT().
			InitEnvBasedConfig("~/.gvel", "testnet").
			Return(errors.New("some error has occurred"))

		err := h.logic.Init("~/.gvel")
		assert.Error(t, err)
	})

	t.Run("fail, AppConfig.InitEnvBasedConfig for mainnet returns an error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			InitSharedConfig("~/.gvel").
			Return(nil)
		h.mockConfiguration.EXPECT().
			InitEnvBasedConfig("~/.gvel", "testnet").
			Return(nil)
		h.mockConfiguration.EXPECT().
			InitEnvBasedConfig("~/.gvel", "mainnet").
			Return(errors.New("some error has occurred"))

		err := h.logic.Init("~/.gvel")
		assert.Error(t, err)
	})
}
