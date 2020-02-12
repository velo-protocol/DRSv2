package logic_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"testing"
)

func TestLogic_CreateAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("0x00")

		output, err := h.logic.CreateAccount(&entity.CreateAccountInput{
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.NoError(t, err)
		assert.False(t, output.IsDefault)
	})

	t.Run("success, no default account, with SetAsDefaultAccount=false", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("")

		h.mockConfiguration.EXPECT().
			SetDefaultAccount(gomock.AssignableToTypeOf("")).
			Return(nil)

		output, err := h.logic.CreateAccount(&entity.CreateAccountInput{
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.NoError(t, err)
		assert.True(t, output.IsDefault)
	})

	t.Run("success, default account is already set, with SetAsDefaultAccount=true", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("0x00")

		h.mockConfiguration.EXPECT().
			SetDefaultAccount(gomock.AssignableToTypeOf("")).
			Return(nil)

		output, err := h.logic.CreateAccount(&entity.CreateAccountInput{
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: true,
		})

		assert.NoError(t, err)
		assert.True(t, output.IsDefault)
	})

	t.Run("fail, DB.Save returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(errors.New("some error has occurred"))

		_, err := h.logic.CreateAccount(&entity.CreateAccountInput{
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to save an account")
	})

	t.Run("fail, AppConfig.SetDefaultAccount returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("")

		h.mockConfiguration.EXPECT().
			SetDefaultAccount(gomock.AssignableToTypeOf("")).
			Return(errors.New("some error has occurred"))

		_, err := h.logic.CreateAccount(&entity.CreateAccountInput{
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to write config file")
	})
}
