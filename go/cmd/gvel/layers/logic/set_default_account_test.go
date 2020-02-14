package logic_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"testing"
)

func TestLogic_SetDefaultAccount(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(accountsBytes(), nil)
		h.mockConfiguration.EXPECT().
			SetDefaultAccount("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4").
			Return(nil)

		output, err := h.logic.SetDefaultAccount(&entity.SetDefaultAccountInput{
			Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
		})

		assert.NoError(t, err)
		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4", output.Account)
	})
	t.Run("error, database returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.SetDefaultAccount(&entity.SetDefaultAccountInput{
			Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address 0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4 is not found in gvel")
	})
	t.Run("error, config returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(accountsBytes(), nil)
		h.mockConfiguration.EXPECT().
			SetDefaultAccount("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4").
			Return(errors.New("some error has occurred"))

		_, err := h.logic.SetDefaultAccount(&entity.SetDefaultAccountInput{
			Account: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to write config file")
	})
}
