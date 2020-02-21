package logic_test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"testing"
)

func TestLogic_MintCreditByCollateral(t *testing.T) {
	var (
		rpcURL              = "http://0.0.0.0:7575"
		assetCode           = "vUSD"
		collateralAssetCode = "VELO"
		collateralAmount    = "100.0000000"
		stableCreditAmount  = "100.0000000"
	)

	mockedMintCreditByCollateralInput := func() *entity.MintCreditByCollateralInput {
		return &entity.MintCreditByCollateralInput{
			AssetCode:        assetCode,
			CollateralAmount: collateralAmount,
			Passphrase:       "password",
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
			MintFromCollateralAmount(gomock.Any(), &vclient.MintFromCollateralAmountInput{
				AssetCode:        assetCode,
				CollateralAmount: collateralAmount,
			}).Return(&vclient.MintFromCollateralAmountCreditOutput{
			Tx:      &types.Transaction{},
			Receipt: &types.Receipt{},
			Event: &vclient.MintFromCollateralAmountEvent{
				AssetCode:           "vUSD",
				StableCreditAmount:  stableCreditAmount,
				AssetAddress:        "0x03",
				CollateralAssetCode: collateralAssetCode,
				CollateralAmount:    collateralAmount,
				Raw:                 nil,
			},
		}, nil)

		output, err := h.logic.MintCreditByCollateral(mockedMintCreditByCollateralInput())

		assert.NoError(t, err)
		assert.NotEmpty(t, output)
		assert.Equal(t, assetCode, output.AssetCode)
		assert.Equal(t, "0xc5b2c658f5fa236c598a6e7fbf7f21413dc42e2a41dd982eb772b30707cba2eb", output.TxHash)
		assert.Equal(t, stableCreditAmount, output.StableCreditAmount)
	})

	t.Run("fail, failed to load accounts from db", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(nil, errors.New("error here"))

		output, err := h.logic.MintCreditByCollateral(mockedMintCreditByCollateralInput())

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

		output, err := h.logic.MintCreditByCollateral(mockedMintCreditByCollateralInput())

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("fail, to decrypt the seed of passphrase", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockConfiguration.EXPECT().
			GetDefaultAccount().
			Return(accountEntity().PublicAddress)

		h.mockDbRepo.EXPECT().
			Get([]byte(accountEntity().PublicAddress)).
			Return(accountsBytes(), nil)

		mockedBadPassphrase := mockedMintCreditByCollateralInput()
		mockedBadPassphrase.Passphrase = "bad passphrase"
		output, err := h.logic.MintCreditByCollateral(mockedBadPassphrase)

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

		output, err := h.logic.MintCreditByCollateral(mockedMintCreditByCollateralInput())

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("fail, connect to veloClient.MintFromCollateralAmount", func(t *testing.T) {
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
			MintFromCollateralAmount(gomock.Any(), gomock.AssignableToTypeOf(&vclient.MintFromCollateralAmountInput{AssetCode: assetCode, CollateralAmount: collateralAmount})).
			Return(nil, errors.New("error here"))

		output, err := h.logic.MintCreditByCollateral(mockedMintCreditByCollateralInput())

		assert.Error(t, err)
		assert.Nil(t, output)
	})
}
