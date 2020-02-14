package logic_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogic_ListAccount(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		mockedAccountsBytes := arrayOfAccountsBytes()

		h.mockDbRepo.EXPECT().
			GetAll().
			Return(mockedAccountsBytes, nil)
		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		accounts, err := h.logic.ListAccount()

		assert.NoError(t, err)
		assert.NotEmpty(t, accounts)
		assert.Equal(t, ((accounts)[0]).PublicAddress, accountEntity().PublicAddress)
		assert.True(t, ((accounts)[0]).IsDefault)
	})

	t.Run("happy, default account not found", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		mockedAccountsBytes := arrayOfAccountsBytes()

		h.mockDbRepo.EXPECT().
			GetAll().
			Return(mockedAccountsBytes, nil)
		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("")

		accounts, err := h.logic.ListAccount()

		assert.NoError(t, err)
		assert.NotEmpty(t, accounts)
		assert.Equal(t, ((accounts)[0]).PublicAddress, accountEntity().PublicAddress)
		assert.False(t, ((accounts)[0]).IsDefault)
	})

	t.Run("fail, failed to load accounts from db", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().GetAll().Return(nil, errors.New("error here"))

		accounts, err := h.logic.ListAccount()

		assert.Empty(t, accounts)
		assert.Error(t, err)
	})

	t.Run("fail, failed to unmarshal stored data to entity", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			GetAll().
			Return([][]byte{[]byte("fake")}, nil)
		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return("0x0")
		accounts, err := h.logic.ListAccount()

		assert.Error(t, err)
		assert.Empty(t, accounts)
	})
}
