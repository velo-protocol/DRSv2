package logic_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"testing"
)

func TestLogic_SetupCredit(t *testing.T) {
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
		h.mockConfiguration.EXPECT().
			GetRpcUrl().
			Return("http://0.0.0.0:7475")
		h.mockConfiguration.EXPECT().
			GetDrsAddress().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")
		h.mockConfiguration.EXPECT().
			GetHeartAddress().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")
		h.mockVFactory.EXPECT().
			NewClient(&entity.NewClientInput{
				RpcUrl:     "http://0.0.0.0:7475",
				PrivateKey: priv,
				ContractAddress: &vclient.ContractAddress{
					DrsAddress:   "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
					HeartAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
				},
			}).
			Return(h.mockVClient, nil)
		h.mockVClient.EXPECT().
			SetupCredit(gomock.Any(), gomock.AssignableToTypeOf(&vclient.SetupCreditInput{})).
			Return(&vclient.SetupCreditOutput{
				Tx:    h.mockTx,
				Event: &vclient.SetupCreditEvent{},
			}, nil)

		output, err := h.logic.SetupCredit(&entity.SetupCreditInput{
			Passphrase: "password",
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

		_, err := h.logic.SetupCredit(&entity.SetupCreditInput{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get account from db")
	})

	t.Run("fail, VFactory.NewClient returns error", func(t *testing.T) {
		acc := accountEntity()
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(acc.PublicAddress)
		h.mockDbRepo.EXPECT().
			Get([]byte(acc.PublicAddress)).
			Return(accountsBytes(), nil)
		h.mockConfiguration.EXPECT().
			GetRpcUrl().
			Return("http://0.0.0.0:7475")
		h.mockConfiguration.EXPECT().
			GetDrsAddress().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")
		h.mockConfiguration.EXPECT().
			GetHeartAddress().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")
		h.mockVFactory.EXPECT().
			NewClient(&entity.NewClientInput{
				RpcUrl:     "http://0.0.0.0:7475",
				PrivateKey: priv,
				ContractAddress: &vclient.ContractAddress{
					DrsAddress:   "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
					HeartAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
				},
			}).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.SetupCredit(&entity.SetupCreditInput{
			Passphrase: "password",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "fail to initiate velo client")
	})

	t.Run("fail, vclient.SetupCredit returns error", func(t *testing.T) {
		acc := accountEntity()
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(acc.PublicAddress)
		h.mockDbRepo.EXPECT().
			Get([]byte(acc.PublicAddress)).
			Return(accountsBytes(), nil)
		h.mockConfiguration.EXPECT().
			GetRpcUrl().
			Return("http://0.0.0.0:7475")
		h.mockConfiguration.EXPECT().
			GetDrsAddress().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")
		h.mockConfiguration.EXPECT().
			GetHeartAddress().
			Return("0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4")
		h.mockVFactory.EXPECT().
			NewClient(&entity.NewClientInput{
				RpcUrl:     "http://0.0.0.0:7475",
				PrivateKey: priv,
				ContractAddress: &vclient.ContractAddress{
					DrsAddress:   "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
					HeartAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
				},
			}).
			Return(h.mockVClient, nil)
		h.mockVClient.EXPECT().
			SetupCredit(gomock.Any(), gomock.AssignableToTypeOf(&vclient.SetupCreditInput{})).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.SetupCredit(&entity.SetupCreditInput{
			Passphrase: "password",
		})

		assert.Error(t, err)
	})
}
