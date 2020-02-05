package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	vabi "github.com/velo-protocol/DRSv2/go/abi"
	"testing"
)

func TestValidateRedeemStableCredit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.Nil(t, err)
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


func TestClient_RedeemStableCredit(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &RedeemStableCreditInput{
			RedeemAmount: "104",
			AssetCode: "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRedeemEvent("Redeem", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRedeem{}, nil)

		result, err := testHelper.Client.RedeemStableCredit(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, result.Tx)
		assert.NotNil(t, result.Receipt)
		assert.NotNil(t, result.Event)
	})

	t.Run("error, validation fail", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		result, err := testHelper.Client.RedeemStableCredit(context.Background(), &RedeemStableCreditInput{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("error, drs.RedeemStableCredit returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &RedeemStableCreditInput{
			RedeemAmount: "104",
			AssetCode: "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(nil, errors.New(expectedMsg))

		result, err := testHelper.Client.RedeemStableCredit(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("error, txHelper.ConfirmTx returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &RedeemStableCreditInput{
			RedeemAmount: "104",
			AssetCode: "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{})).
			Return(nil, errors.New(expectedMsg))

		result, err := testHelper.Client.RedeemStableCredit(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("error, txHelper.ExtractRedeemEvent returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &RedeemStableCreditInput{
			RedeemAmount: "104",
			AssetCode: "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRedeemEvent("Redeem", gomock.AssignableToTypeOf(&types.Log{})).
			Return(nil, errors.New(expectedMsg))

		result, err := testHelper.Client.RedeemStableCredit(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedMsg)
	})
}
