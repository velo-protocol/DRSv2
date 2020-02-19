package logic_test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"testing"
)

func TestLogic_CollateralHealthCheck(t *testing.T) {
	var (
		rpcURL                 = "http://0.0.0.0:7575"
		collateralAssetCode    = "VELO"
		collateralAssetAddress = "0x04"
		presentAmount          = "1.0000000"
		requiredAmount         = "2.0000000"
	)

	mockedCollateralHealthCheckInput := func() *entity.CollateralHealthCheckInput {
		return &entity.CollateralHealthCheckInput{
			Passphrase: "password",
		}
	}

	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(accountsBytes(), nil)

		h.mockConfiguration.EXPECT().
			GetRpcUrl().
			Return(rpcURL)

		h.mockConfiguration.EXPECT().
			GetDrsAddress().
			Return("0x01")

		h.mockConfiguration.EXPECT().
			GetHeartAddress().
			Return("0x02")

		h.mockVFactory.EXPECT().
			NewClient(&entity.NewClientInput{
				RpcUrl:          rpcURL,
				PrivateKey:      "6d71af6c908ff8b618825926f1004431915faf9b66238c30af9f86438d2bcd89",
				ContractAddress: &vclient.ContractAddress{"0x01", "0x02"},
			}).Return(h.mockVClient, nil)

		h.mockVClient.EXPECT().
			CollateralHealthCheck(&vclient.CollateralHealthCheckInput{}).
			Return([]vclient.CollateralHealthCheckOutput{
				{
					CollateralAssetAddress: collateralAssetAddress,
					CollateralAssetCode:    collateralAssetCode,
					PresentAmount:          presentAmount,
					RequiredAmount:         requiredAmount,
				},
			}, nil)

		output, err := h.logic.CollateralHealthCheck(mockedCollateralHealthCheckInput())

		assert.NoError(t, err)
		assert.NotEmpty(t, output)
		assert.Equal(t, collateralAssetCode, output[0].CollateralAssetCode)
		assert.Equal(t, collateralAssetAddress, output[0].CollateralAssetAddress)
		assert.Equal(t, presentAmount, output[0].CollateralPoolAmount)
		assert.Equal(t, requiredAmount, output[0].RequiredCollateralAmount)
	})

	t.Run("fail, default account is not found", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(nil, errors.New("error here"))

		output, err := h.logic.CollateralHealthCheck(mockedCollateralHealthCheckInput())

		assert.Nil(t, output)
		assert.Error(t, err)
	})

	t.Run("fail, failed to unmarshal stored data to entity", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		badAccountByte, _ := json.Marshal("")

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(badAccountByte, nil)

		output, err := h.logic.CollateralHealthCheck(mockedCollateralHealthCheckInput())

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("fail, to decrypt the private of passphrase", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(accountsBytes(), nil)

		mockedBadPassphrase := mockedCollateralHealthCheckInput()
		mockedBadPassphrase.Passphrase = "bad passphrase"
		output, err := h.logic.CollateralHealthCheck(mockedBadPassphrase)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, fmt.Sprintf("failed to decrypt the private key of %s with given passphrase: failed to decipher and authenticate: cipher: message authentication failed", "0xf41E18a9573832265F74a671d3E275ec76790b5C"), err.Error())
	})

	t.Run("fail, connect to veloFactory.NewClient", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(accountsBytes(), nil)

		h.mockConfiguration.EXPECT().
			GetRpcUrl().
			Return(rpcURL)

		h.mockConfiguration.EXPECT().
			GetDrsAddress().
			Return("0x01")

		h.mockConfiguration.EXPECT().
			GetHeartAddress().
			Return("0x02")

		h.mockVFactory.EXPECT().
			NewClient(&entity.NewClientInput{
				RpcUrl:     rpcURL,
				PrivateKey: "6d71af6c908ff8b618825926f1004431915faf9b66238c30af9f86438d2bcd89",
				ContractAddress: &vclient.ContractAddress{
					"0x01",
					"0x02",
				},
			}).
			Return(nil, errors.New("error here"))

		output, err := h.logic.CollateralHealthCheck(mockedCollateralHealthCheckInput())

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("fail, connect to veloClient.GetExchangeRate", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(accountsBytes(), nil)

		h.mockConfiguration.EXPECT().
			GetRpcUrl().
			Return(rpcURL)

		h.mockConfiguration.EXPECT().
			GetDrsAddress().
			Return("0x01")

		h.mockConfiguration.EXPECT().
			GetHeartAddress().
			Return("0x02")

		h.mockVFactory.EXPECT().
			NewClient(&entity.NewClientInput{
				RpcUrl:          rpcURL,
				PrivateKey:      "6d71af6c908ff8b618825926f1004431915faf9b66238c30af9f86438d2bcd89",
				ContractAddress: &vclient.ContractAddress{"0x01", "0x02"},
			}).Return(h.mockVClient, nil)

		h.mockVClient.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&vclient.CollateralHealthCheckInput{})).
			Return(nil, errors.New("error here"))

		output, err := h.logic.CollateralHealthCheck(mockedCollateralHealthCheckInput())

		assert.Error(t, err)
		assert.Nil(t, output)
	})
}
