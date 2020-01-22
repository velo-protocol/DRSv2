package client

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSetupCreditInput_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}).Validate()

		assert.NoError(t, err)
	})
	t.Run("error, invalid collateralAssetCode format", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "BAD_VELO!",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}).Validate()

		assert.EqualError(t, err, "invalid collateralAssetCode format")
	})
	t.Run("error, invalid peggedCurrency format", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "BAD_CURRENCY",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}).Validate()

		assert.EqualError(t, err, "invalid peggedCurrency format")
	})
	t.Run("error, invalid assetCode format", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "BAD_ASSET!",
			PeggedValue:         "1.50",
		}).Validate()

		assert.EqualError(t, err, "invalid assetCode format")
	})
	t.Run("error, invalid peggedValue format", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "I_am_a_number!",
		}).Validate()

		assert.Contains(t, err.Error(), "invalid peggedValue format")
	})
	t.Run("error, peggedValue must be positive", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "-1.50",
		}).Validate()

		assert.Contains(t, err.Error(), "peggedValue must be positive")
	})
}

func TestSetupCreditInput_ToAbiInput(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		abiInput := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}).ToAbiInput()

		assert.NotEmpty(t, abiInput.CollateralAssetCode)
		assert.NotEmpty(t, abiInput.PeggedCurrency)
		assert.Equal(t, "vUSD", abiInput.AssetCode)
		assert.Equal(t, big.NewInt(15000000), abiInput.PeggedValue)
	})
}

func TestClient_SetupCredit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testHelper := testHelper()

		result, err := testHelper.Client.SetupCredit(&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, result.Tx.Hash())
	})

	t.Run("error, validation fail", func(t *testing.T) {
		testHelper := testHelper()
		_, err := testHelper.Client.SetupCredit(&SetupCreditInput{})

		assert.Error(t, err)
	})
}
