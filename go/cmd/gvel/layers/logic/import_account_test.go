package logic_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"testing"
)

func TestLogic_ImportAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(nil, errors.New("leveldb: not found"))

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")

		output, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4", output.ImportedAccountAddress)
	})

	t.Run("success, no default account with SetAsDefaultAccount=false", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(nil, errors.New("leveldb: not found"))

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("")

		h.mockConfiguration.EXPECT().
			SetDefaultAccount(gomock.AssignableToTypeOf("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(nil)

		output, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4", output.ImportedAccountAddress)
	})

	t.Run("success, default account is already set with SetAsDefaultAccount=true", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(nil, errors.New("leveldb: not found"))

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")

		h.mockConfiguration.EXPECT().
			SetDefaultAccount(gomock.AssignableToTypeOf("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(nil)

		output, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: true,
		})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4", output.ImportedAccountAddress)
	})

	t.Run("fail, getting account from db", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(nil, errors.New("some error has occurred"))

		output, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "failed to get account from db")
	})

	t.Run("fail, account already exist", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(accountsBytes(), nil)

		output, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "the account already existed in gvel", err.Error())
	})

	t.Run("fail, ValidatePrivateKeyFormat returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		_, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid private key format")
	})

	t.Run("fail, DB.Save returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(nil, errors.New("leveldb: not found"))

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(errors.New("some error has occurred"))

		_, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: true,
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to save an account")
	})

	t.Run("fail, AppConfig.SetDefaultAccount returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get(gomock.AssignableToTypeOf([]byte{})).
			Return(nil, errors.New("leveldb: not found"))

		h.mockDbRepo.EXPECT().
			Save(gomock.AssignableToTypeOf([]byte{}), gomock.AssignableToTypeOf([]byte{})).
			Return(nil)

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("")

		h.mockConfiguration.EXPECT().
			SetDefaultAccount(gomock.AssignableToTypeOf("")).
			Return(errors.New("some error has occurred"))

		_, err := h.logic.ImportAccount(&entity.ImportAccountInput{
			PrivateKey:          "ede4697e04e1e05d58bcec13196a065de82a01f3ce8bf074a4ff36b7cc62d54e",
			Passphrase:          "strong_password!",
			SetAsDefaultAccount: false,
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to write config file")
	})
}
