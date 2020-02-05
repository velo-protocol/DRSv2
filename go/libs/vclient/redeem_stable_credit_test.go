package vclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRedeemStableCredit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.Nil(t, err)
	})

	t.Run("error, validation fail invalid collateralAmount is a number with greater than 7 decimal places", func(t *testing.T) {
		err := (&MintFromCollateralAmountInput{
			AssetCode:        "vUSD",
			CollateralAmount: "10.12345678",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "invalid collateralAmount format", err.Error())
	})

	t.Run("fail, should throw error invalid RedeemAmount format", func(t *testing.T) {
		err := (&RedeemStableCreditInput{
			RedeemAmount: "10.12345678",
			AssetCode: "vUSD",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "invalid RedeemAmount format", err.Error())
	})

	t.Run("fail, should throw error RedeemAmount must be positive", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "-2", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "RedeemAmount must be positive", err.Error())
	})

	t.Run("fail, should throw error assetCode must not be blank", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: ""}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "assetCode must not be blank", err.Error())
	})

	t.Run("fail, should throw error invalid assetCode format", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "BAD_ASSET_CODE"}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "invalid assetCode format", err.Error())
	})
}

func TestValidateRedeemStableCreditToAbiInput(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "vUSD"}
		abiInput := redeemStableCreditInput.ToAbiInput()

		assert.NotNil(t, abiInput)
	})
}

func TestRedeemStableCredit(t *testing.T) {
	// TODO VELO-657
}
