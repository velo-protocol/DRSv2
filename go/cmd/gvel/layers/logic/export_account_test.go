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
			Get([]byte(pub)).
			Return(accountsBytes(), nil)

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: pub,
			Passphrase:    "password",
		})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.PrivateKey)
		assert.Equal(t, pub, output.PublicAddress)
		assert.Equal(t, priv, output.PrivateKey)
	})

	t.Run("fail, unmarshal account error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		badAccountByte, _ := json.Marshal("")

		h.mockDbRepo.EXPECT().
			Get([]byte(pub)).
			Return(badAccountByte, nil)

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: pub,
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
			Get([]byte(pub)).
			Return(nil, errors.New("failed to get account from db"))

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: pub,
			Passphrase:    "password",
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "default account not found in config file, please run `gvel account create` or `gvel account import`")
	})

	t.Run("fail, to decrypt the seed of passphrase", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockDbRepo.EXPECT().
			Get([]byte(pub)).
			Return(accountsBytes(), nil)

		output, err := h.logic.ExportAccount(&entity.ExportAccountInput{
			PublicAddress: pub,
			Passphrase:    "bad password",
		})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, fmt.Sprintf("failed to decrypt the seed of %s with given passphrase: failed to decipher and authenticate: cipher: message authentication failed", "0xf41E18a9573832265F74a671d3E275ec76790b5C"), err.Error())
	})
}
