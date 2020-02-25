package logic_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"testing"
)

func TestLogic_RebalanceCollateral(t *testing.T) {
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
			Rebalance(gomock.Any(), &vclient.RebalanceInput{}).
			Return(&vclient.RebalanceOutput{
				RebalanceTransactions: []*vclient.RebalanceTransaction{
					{
						Tx:                  h.mockTx,
						AssetCode:           "vUSD",
						CollateralAssetCode: "VELO",
						RequiredAmount:      "100.0000000",
						PresentAmount:       "120.0000000",
					},
					{
						Tx:                  h.mockTx2,
						AssetCode:           "vSGD",
						CollateralAssetCode: "VELO",
						RequiredAmount:      "230.0000000",
						PresentAmount:       "200.0000000",
					},
				},
			}, nil)

		outputs, err := h.logic.RebalanceCollateral(&entity.RebalanceCollateralInput{
			Passphrase: "password",
		})

		assert.NoError(t, err)
		assert.NotNil(t, outputs)

		assert.Equal(t, h.mockTx.Hash().String(), outputs[0].TxHash)
		assert.Equal(t, "vUSD", outputs[0].AssetCode)
		assert.Equal(t, "VELO", outputs[0].CollateralAssetCode)
		assert.Equal(t, "100.0000000", outputs[0].RequiredAmount)
		assert.Equal(t, "120.0000000", outputs[0].PresentAmount)

		assert.Equal(t, h.mockTx2.Hash().String(), outputs[1].TxHash)
		assert.Equal(t, "vSGD", outputs[1].AssetCode)
		assert.Equal(t, "VELO", outputs[1].CollateralAssetCode)
		assert.Equal(t, "230.0000000", outputs[1].RequiredAmount)
		assert.Equal(t, "200.0000000", outputs[1].PresentAmount)
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

		_, err := h.logic.RebalanceCollateral(&entity.RebalanceCollateralInput{
			Passphrase: "password",
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

		_, err := h.logic.RebalanceCollateral(&entity.RebalanceCollateralInput{
			Passphrase: "password",
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
			Rebalance(gomock.Any(), &vclient.RebalanceInput{}).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.logic.RebalanceCollateral(&entity.RebalanceCollateralInput{
			Passphrase: "password",
		})

		assert.Error(t, err)
	})
}
