package account_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_Default(t *testing.T) {

	mockAccountList := []string{
		"0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
		"0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
		"0x0f1D6Ad59AE485A9ec31b36154093820337bdEA3",
	}

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return([]*entity.Account{
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
					IsDefault:     false,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
					IsDefault:     true,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA3",
					IsDefault:     false,
				},
			}, nil)

		h.mockConfig.EXPECT().
			GetDefaultAccount().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2")

		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf("Please select the account you want to make default"), mockAccountList, console.RequestChoiceOptions{
				ActiveChoice:      "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(0)

		h.mockLogic.EXPECT().
			SetDefaultAccount(&entity.SetDefaultAccountInput{
				Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
			}).
			Return(&entity.SetDefaultAccountOutput{Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1"}, nil)

		assert.NotPanics(t, func() {
			h.commandHandler.Default(h.defaultCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1 is now set as the default account for signing transaction.", logEntries[0].Message)
	})

	t.Run("success, with old default account", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return([]*entity.Account{
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
					IsDefault:     false,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
					IsDefault:     true,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA3",
					IsDefault:     false,
				},
			}, nil)

		h.mockConfig.EXPECT().
			GetDefaultAccount().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2")

		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf("Please select the account you want to make default"), mockAccountList, console.RequestChoiceOptions{
				ActiveChoice:      "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(1)

		h.mockLogic.EXPECT().
			SetDefaultAccount(&entity.SetDefaultAccountInput{
				Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
			}).
			Return(&entity.SetDefaultAccountOutput{Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2"}, nil)

		assert.NotPanics(t, func() {
			h.commandHandler.Default(h.defaultCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2 is now set as the default account for signing transaction.", logEntries[0].Message)
	})

	t.Run("fail, get account list", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return([]*entity.Account{}, nil)

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Default(h.defaultCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "account not found in config file, please run `gvel account create` or `gvel account import`", logEntries[0].Message)
	})

	t.Run("fail, account list is no once account", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return(nil, errors.New("fail to get Logic.ListAccount()"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Default(h.defaultCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "fail to get Logic.ListAccount()", logEntries[0].Message)
	})

	t.Run("fail, set default account", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return([]*entity.Account{
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
					IsDefault:     false,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
					IsDefault:     true,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA3",
					IsDefault:     false,
				},
			}, nil)

		h.mockConfig.EXPECT().
			GetDefaultAccount().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2")

		h.mockPrompt.EXPECT().
			RequestChoice(gomock.AssignableToTypeOf("Please select the account you want to make default"), mockAccountList, console.RequestChoiceOptions{
				ActiveChoice:      "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
				ActiveChoiceLabel: constants.CursorCurrent,
			}).
			Return(0)

		h.mockLogic.EXPECT().
			SetDefaultAccount(&entity.SetDefaultAccountInput{
				Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
			}).
			Return(nil, errors.New("fail to call Logic.SetDefaultAccount()"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Default(h.defaultCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "fail to call Logic.SetDefaultAccount()", logEntries[0].Message)
	})
}
