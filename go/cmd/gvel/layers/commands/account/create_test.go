package account_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"testing"
)

func TestCommandHandler_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestPassphrase().
			Return("strong_password!")

		h.mockLogic.EXPECT().
			CreateAccount(&entity.CreateAccountInput{
				Passphrase:          "strong_password!",
				SetAsDefaultAccount: false,
			}).
			Return(&entity.CreateAccountOutput{PublicAddress: "0x01"}, nil)

		h.commandHandler.Create(h.createCmd, nil)
		assert.Equal(t, "A new account is created with address 0x01. Please remember to keep your passphrase safe. You will not be able to recover this passphrase.", h.logHook.LastEntry().Message)
	})
	t.Run("success, with default flag", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		_ = h.createCmd.Flags().Set("default", "true")

		h.mockPrompt.EXPECT().
			RequestPassphrase().
			Return("strong_password!")

		h.mockLogic.EXPECT().
			CreateAccount(&entity.CreateAccountInput{
				Passphrase:          "strong_password!",
				SetAsDefaultAccount: true,
			}).
			Return(&entity.CreateAccountOutput{PublicAddress: "0x01"}, nil)

		h.commandHandler.Create(h.createCmd, nil)
		assert.Equal(t, "A new account is created with address 0x01. Please remember to keep your passphrase safe. You will not be able to recover this passphrase.", h.logHook.LastEntry().Message)
	})

	t.Run("fail, logic.CreateAccount returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestPassphrase().
			Return("strong_password!")

		h.mockLogic.EXPECT().
			CreateAccount(&entity.CreateAccountInput{
				Passphrase:          "strong_password!",
				SetAsDefaultAccount: false,
			}).
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Create(h.createCmd, nil)
		})
		assert.Equal(t, "some error has occurred", h.logHook.LastEntry().Message)
	})
}
