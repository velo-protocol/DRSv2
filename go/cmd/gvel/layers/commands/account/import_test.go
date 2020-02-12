package account_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/validation"
	"testing"
)

func TestCommandHandler_Import(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestHiddenString("Please enter private key of the address", gomock.AssignableToTypeOf(validation.ValidatePrivateKeyFormat)).
			Return("ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e")

		h.mockPrompt.EXPECT().
			RequestPassphrase().
			Return("strong_password!")

		h.mockLogic.EXPECT().
			ImportAccount(&entity.ImportAccountInput{
				PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
				Passphrase:          "strong_password!",
				SetAsDefaultAccount: false,
			}).
			Return(&entity.ImportAccountOutput{ImportedAccountAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4"}, nil)

		h.commandHandler.Import(h.importCmd, nil)
		assert.Equal(t, "Add account with address 0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4 to gvel. Please keep your passphrase safe.", h.logHook.LastEntry().Message)
	})
	t.Run("success, with default flag", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		_ = h.importCmd.Flags().Set("default", "true")

		h.mockPrompt.EXPECT().
			RequestHiddenString("Please enter private key of the address", gomock.AssignableToTypeOf(validation.ValidatePrivateKeyFormat)).
			Return("ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e")

		h.mockPrompt.EXPECT().
			RequestPassphrase().
			Return("strong_password!")

		h.mockLogic.EXPECT().
			ImportAccount(&entity.ImportAccountInput{
				PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
				Passphrase:          "strong_password!",
				SetAsDefaultAccount: true,
			}).
			Return(&entity.ImportAccountOutput{ImportedAccountAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4"}, nil)

		h.commandHandler.Import(h.importCmd, nil)
		assert.Equal(t, "Account with address 0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4 is now the default account. Please keep your passphrase safe.", h.logHook.LastEntry().Message)
	})

	t.Run("fail, logic.ImportAccount returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockPrompt.EXPECT().
			RequestHiddenString("Please enter private key of the address", gomock.AssignableToTypeOf(validation.ValidatePrivateKeyFormat)).
			Return("ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e")

		h.mockPrompt.EXPECT().
			RequestPassphrase().
			Return("strong_password!")

		h.mockLogic.EXPECT().
			ImportAccount(&entity.ImportAccountInput{
				PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
				Passphrase:          "strong_password!",
				SetAsDefaultAccount: false,
			}).
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.Import(h.importCmd, nil)
		})
		assert.Equal(t, "some error has occurred", h.logHook.LastEntry().Message)
	})
}
