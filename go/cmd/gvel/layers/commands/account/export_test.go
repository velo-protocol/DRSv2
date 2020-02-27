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

func TestCommandHandler_Export(t *testing.T) {

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
			RequestChoice(gomock.AssignableToTypeOf("Please select the account you want to export"), mockAccountList, console.RequestChoiceOptions{
				ActiveChoice:      "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
				ActiveChoiceLabel: constants.CursorDefault,
			}).
			Return(0)

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input the passphrase of the account", nil).
			Return("strong_password!")

		h.mockLogic.EXPECT().
			ExportAccount(&entity.ExportAccountInput{
				PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
				Passphrase:    "strong_password!",
			}).
			Return(&entity.ExportAccountOutput{PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1", PrivateKey: "SBR25NMQRKQ4RLGNV5XB3MMQB4ADVYSMPGVBODQVJE7KPTDR6KGK3XMX"}, nil)

		assert.NotPanics(t, func() {
			h.commandHandler.Export(h.exportCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "Your public key is: 0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1", logEntries[0].Message)
		assert.Equal(t, "Your private key is: SBR25NMQRKQ4RLGNV5XB3MMQB4ADVYSMPGVBODQVJE7KPTDR6KGK3XMX", logEntries[1].Message)
	})

	t.Run("fail, account list is no once account", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return([]*entity.Account{}, nil)

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Export(h.exportCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "account not found in config file, please run `gvel account create` or `gvel account import`", logEntries[0].Message)
	})

	t.Run("fail, get account list return error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return(nil, errors.New("fail to get Logic.ExportAccount()"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Export(h.exportCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "fail to get Logic.ExportAccount()", logEntries[0].Message)
	})

	t.Run("fail, export account", func(t *testing.T) {
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
			RequestChoice(gomock.AssignableToTypeOf("Please select the account you want to export"), mockAccountList, console.RequestChoiceOptions{
				ActiveChoice:      "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
				ActiveChoiceLabel: constants.CursorDefault,
			}).
			Return(0)

		h.mockPrompt.EXPECT().
			RequestHiddenString("🔑 Please input the passphrase of the account", nil).
			Return("strong_password!")

		h.mockLogic.EXPECT().
			ExportAccount(&entity.ExportAccountInput{
				PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
				Passphrase:    "strong_password!",
			}).
			Return(nil, errors.New("fail to call Logic.ExportAccount()"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Export(h.exportCmd, nil)
		})

		logEntries := h.logHook.AllEntries()

		assert.Equal(t, "fail to call Logic.ExportAccount()", logEntries[0].Message)
	})
}
