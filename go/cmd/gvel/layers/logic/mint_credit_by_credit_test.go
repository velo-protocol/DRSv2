package logic_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"testing"
)

func TestLogic_MintByCredit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		acc := accountEntity()
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(acc.PublicAddress)
		h.mockDbRepo.EXPECT().
			Get([]byte(acc.PublicAddress)).
			Return(accountsBytes(), nil)
		h.mockVFactory.EXPECT().
			NewClientFromConfig(priv).
			Return(h.mockVClient, nil)
		h.mockVClient.EXPECT().
			MintFromStableCreditAmount(gomock.Any(), &vclient.MintFromStableCreditAmountInput{
				AssetCode:          "vUSD",
				StableCreditAmount: "1000",
			}).
			Return(&vclient.MintFromStableCreditAmountCreditOutput{
				Tx:    h.mockTx,
				Event: &vclient.MintFromStableCreditAmountEvent{},
			}, nil)

		output, err := h.logic.MintCreditByCredit(&entity.MintCreditByCreditInput{
			Passphrase:   "password",
			AssetCode:    "vUSD",
			CreditAmount: "1000",
		})

		assert.NoError(t, err)
		assert.NotNil(t, output)
	})

	t.Run("fail, database repo returns error", func(t *testing.T) {
		acc := accountEntity()
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(acc.PublicAddress)
		h.mockDbRepo.EXPECT().
			Get([]byte(acc.PublicAddress)).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.MintCreditByCredit(&entity.MintCreditByCreditInput{
			Passphrase:   "password",
			AssetCode:    "vUSD",
			CreditAmount: "1000",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "default account is not found")
	})

	t.Run("fail, VFactory.NewClientFromConfig returns error", func(t *testing.T) {
		acc := accountEntity()
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(acc.PublicAddress)
		h.mockDbRepo.EXPECT().
			Get([]byte(acc.PublicAddress)).
			Return(accountsBytes(), nil)
		h.mockVFactory.EXPECT().
			NewClientFromConfig(priv).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.MintCreditByCredit(&entity.MintCreditByCreditInput{
			Passphrase:   "password",
			AssetCode:    "vUSD",
			CreditAmount: "1000",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "fail to initiate velo client")
	})

	t.Run("fail, veloClient.MintFromStableCreditAmount returns error", func(t *testing.T) {
		acc := accountEntity()
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(acc.PublicAddress)
		h.mockDbRepo.EXPECT().
			Get([]byte(acc.PublicAddress)).
			Return(accountsBytes(), nil)
		h.mockVFactory.EXPECT().
			NewClientFromConfig(priv).
			Return(h.mockVClient, nil)
		h.mockVClient.EXPECT().
			MintFromStableCreditAmount(gomock.Any(), &vclient.MintFromStableCreditAmountInput{
				AssetCode:          "vUSD",
				StableCreditAmount: "1000",
			}).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.MintCreditByCredit(&entity.MintCreditByCreditInput{
			Passphrase:   "password",
			AssetCode:    "vUSD",
			CreditAmount: "1000",
		})

		assert.Error(t, err)
	})
}
