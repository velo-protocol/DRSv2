package environment_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_Set(t *testing.T) {
	t.Run("success, select mainnet", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfig.EXPECT().
			GetEnvList().
			Return(constants.DefaultEnvList)
		h.mockConfig.EXPECT().
			GetCurrentEnv().
			Return("")
		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf(""), constants.DefaultEnvList, console.RequestChoiceOptions{
				ActiveChoice:      "",
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(1)
		h.mockPrompt.EXPECT().
			RequestConfirmation(gomock.AssignableToTypeOf("")).
			Return(true)
		h.mockLogic.EXPECT().
			SetEnv(gomock.AssignableToTypeOf(&entity.SetEnvInput{})).
			Return(nil)

		assert.NotPanics(t, func() {
			h.commandHandler.Set(nil, nil)
		})

		logEntries := h.loggerHook.AllEntries()
		assert.Equal(t, "Switch to MAINNET. Warning: All commands will be executed on MAINNET chain. Please take care when executing transaction.\n", logEntries[0].Message)
	})

	t.Run("success, select testnet", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfig.EXPECT().
			GetEnvList().
			Return(constants.DefaultEnvList)
		h.mockConfig.EXPECT().
			GetCurrentEnv().
			Return("")
		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf(""), constants.DefaultEnvList, console.RequestChoiceOptions{
				ActiveChoice:      "",
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(0)
		h.mockLogic.EXPECT().
			SetEnv(gomock.AssignableToTypeOf(&entity.SetEnvInput{})).
			Return(nil)

		assert.NotPanics(t, func() {
			h.commandHandler.Set(nil, nil)
		})

		logEntries := h.loggerHook.AllEntries()
		assert.Equal(t, "Switch to TESTNET. This is for testing. The tokens holds no value and should not be used to do real transaction.\n", logEntries[0].Message)
	})
	t.Run("success, select mainnet but cancel confirmation", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfig.EXPECT().
			GetEnvList().
			Return(constants.DefaultEnvList)
		h.mockConfig.EXPECT().
			GetCurrentEnv().
			Return("")
		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf(""), constants.DefaultEnvList, console.RequestChoiceOptions{
				ActiveChoice:      "",
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(1)
		h.mockPrompt.EXPECT().
			RequestConfirmation(gomock.AssignableToTypeOf("")).
			Return(false)

		assert.NotPanics(t, func() {
			h.commandHandler.Set(nil, nil)
		})
	})
	t.Run("fail, empty env list", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfig.EXPECT().
			GetEnvList().
			Return(nil)

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Set(nil, nil)
		})

		logEntries := h.loggerHook.AllEntries()
		assert.Equal(t, "env list is empty", logEntries[0].Message)
	})

	t.Run("fail, Logic.SetEnv returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfig.EXPECT().
			GetEnvList().
			Return(constants.DefaultEnvList)
		h.mockConfig.EXPECT().
			GetCurrentEnv().
			Return(constants.EnvMainNet)
		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf(""), constants.DefaultEnvList, console.RequestChoiceOptions{
				ActiveChoice:      constants.EnvMainNet,
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(1)
		h.mockPrompt.EXPECT().
			RequestConfirmation(gomock.AssignableToTypeOf("")).
			Return(true)
		h.mockLogic.EXPECT().
			SetEnv(gomock.AssignableToTypeOf(&entity.SetEnvInput{})).
			Return(errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Set(nil, nil)
		})
	})
}
