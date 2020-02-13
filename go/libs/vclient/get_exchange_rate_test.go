package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"testing"
)

func TestValidateGetExchangeRate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		getExchangeRateInput := &GetExchangeRateInput{AssetCode: "vUSD"}
		err := getExchangeRateInput.Validate()

		assert.Nil(t, err)
	})

	t.Run("fail, should throw error assetCode must not be blank", func(t *testing.T) {
		getExchangeRateInput := &GetExchangeRateInput{AssetCode: ""}
		err := getExchangeRateInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "assetCode must not be blank", err.Error())
	})

	t.Run("fail, should throw error invalid assetCode format", func(t *testing.T) {
		getExchangeRateInput := &GetExchangeRateInput{AssetCode: "BAD_ASSET_CODE"}
		err := getExchangeRateInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "invalid assetCode format", err.Error())
	})
}

func TestValidateGetExchangeRateToAbiInput(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		getExchangeRateInput := &GetExchangeRateInput{AssetCode: "vUSD"}
		abiInput := getExchangeRateInput.ToAbiInput()

		assert.NotNil(t, abiInput)
	})
}

func TestGetExchangeRate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		pricePerAssetUnit, _ := utils.StringToAmount("1")

		input := &GetExchangeRateInput{
			AssetCode: "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			GetExchange(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.AssetCode).
			Return("vUSD", utils.StringToByte32("VELO"), pricePerAssetUnit, nil)

		result, err := testHelper.Client.GetExchangeRate(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "vUSD", result.AssetCode)
		assert.Equal(t, "VELO", result.CollateralAssetCode)
		assert.Equal(t, "1.0000000", result.PriceInCollateralPerAssetUnit)
	})

	t.Run("fail, should throw error assetCode must not be blank", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &GetExchangeRateInput{
			AssetCode: "",
		}

		result, err := testHelper.Client.GetExchangeRate(input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "assetCode must not be blank", err.Error())
	})
}
