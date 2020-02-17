package logic_test

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"testing"
)

func TestLogic_ExportAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(accountsBytes(), nil)

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
			Passphrase:    "password",
		})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.PrivateKey)
		assert.Equal(t, "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4", output.PublicAddress)
		assert.Equal(t, "SBR25NMQRKQ4RLGNV5XB3MMQB4ADVYSMPGVBODQVJE7KPTDR6KGK3XMX", output.PrivateKey)
	})

	t.Run("fail, unmarshal account error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		badAccountByte, _ := json.Marshal("")

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(badAccountByte, nil)

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
			Passphrase:    "password",
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "failed to unmarshal account: json: cannot unmarshal string into Go value of type entity.Account", err.Error())
	})

	t.Run("fail, failed to get the account from DB", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(nil, errors.New("failed to get account from db"))

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
			Passphrase:    "password",
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "account not found in config file, please run gvel account create or gvel account import: failed to get account from db", err.Error())
	})

	t.Run("fail, to decrypt the seed of passphrase", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")).
			Return(accountsBytes(), nil)

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
			Passphrase:    "bad password",
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, fmt.Sprintf("failed to decrypt the seed of %s with given passphrase: failed to decipher and authenticate: cipher: message authentication failed", "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4"), err.Error())
	})
}
