package vclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMintFromCollateralAmountInput_Validate(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode:        "vUSD",
			CollateralAmount: "100",
		}).Validate()

		assert.NoError(t, err)
	})

	t.Run("error, validation fail assetCode must not be blank", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			CollateralAmount: "100",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "assetCode must not be blank", err.Error())
	})

	t.Run("error, validation fail invalid assetCode format", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode:        "AssetCodee!@",
			CollateralAmount: "100",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "invalid assetCode format", err.Error())
	})

	t.Run("error, validation fail collateralAmount must not be blank", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode: "vUSD",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "collateralAmount must not be blank", err.Error())
	})

	t.Run("error, validation fail invalid collateralAmount format", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode:        "vUSD",
			CollateralAmount: "amount",
		}).Validate()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid collateralAmount format")
	})

	t.Run("error, validation fail collateralAmount must be greater than 0", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode:        "vUSD",
			CollateralAmount: "0",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "collateralAmount must be greater than 0", err.Error())
	})

	t.Run("error, validation fail invalid collateralAmount is a number with greater than 7 decimal places", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode:        "vUSD",
			CollateralAmount: "10.12345678",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "invalid collateralAmount format", err.Error())
	})
}
