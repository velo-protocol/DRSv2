package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	vabi "github.com/velo-protocol/DRSv2/go/abi"
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
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Setup(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.CollateralAssetCode, abiInput.PeggedCurrency, abiInput.AssetCode, abiInput.PeggedValue).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractSetupEvent("Setup", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemSetup{}, nil)

		result, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.NoError(t, err)
		assert.NotNil(t, result.Tx)
		assert.NotNil(t, result.Receipt)
		assert.NotNil(t, result.Event)
	})

	t.Run("error, validation fail", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		_, err := testHelper.Client.SetupCredit(context.Background(), &SetupCreditInput{})

		assert.Error(t, err)
	})
}
