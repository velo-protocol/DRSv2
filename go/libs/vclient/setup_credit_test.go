package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/abi"
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
	t.Run("error, input parameters must not be blank", func(t *testing.T) {
		errCollateralAssetCode := (&SetupCreditInput{
			CollateralAssetCode: "",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}).Validate()
		assert.EqualError(t, errCollateralAssetCode, "collateralAssetCode must not be blank")

		errPeggedCurrency := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "",
			AssetCode:           "vUSD",
			PeggedValue:         "1.50",
		}).Validate()
		assert.EqualError(t, errPeggedCurrency, "peggedCurrency must not be blank")

		errAssetCode := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "",
			PeggedValue:         "1.50",
		}).Validate()
		assert.EqualError(t, errAssetCode, "assetCode must not be blank")

		errPeggedValue := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "",
		}).Validate()
		assert.EqualError(t, errPeggedValue, "peggedValue must not be blank")
	})
	t.Run("error, invalid peggedValue with more than 7 decimal places is not allowed", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "1.12345678",
		}).Validate()
		assert.EqualError(t, err, "peggedValue with more than 7 decimal places is not allowed")
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

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid peggedValue format")
	})
	t.Run("error, peggedValue must be positive", func(t *testing.T) {
		err := (&SetupCreditInput{
			CollateralAssetCode: "VELO",
			PeggedCurrency:      "USD",
			AssetCode:           "vUSD",
			PeggedValue:         "-1.50",
		}).Validate()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "peggedValue must be greater than 0")
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
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Setup(string,bytes32,uint256,bytes32,address)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractSetupEvent("Setup", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemSetup{
				AssetCode:           "vUSD",
				PeggedCurrency:      [32]byte{},
				PeggedValue:         big.NewInt(10000000),
				CollateralAssetCode: [32]byte{},
				AssetAddress:        common.Address{},
			}, nil)

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

	t.Run("error, drs.Setup returns an error", func(t *testing.T) {
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
			Return(nil, errors.New("some error has occurred"))

		_, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.Error(t, err)
	})

	t.Run("error, drs.ConfirmTx returns an error revert DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", func(t *testing.T) {
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
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New("VM Exception while processing transaction: revert DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner"))
		_, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.Error(t, err)
		assert.Equal(t, "confirm transaction error: the message sender is not found or does not have sufficient permission to perform setup stable credit", err.Error())
	})

	t.Run("error, drs.ConfirmTx returns an error revert DigitalReserveSystem.setup: assetCode has already been used", func(t *testing.T) {
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
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New("VM Exception while processing transaction: revert DigitalReserveSystem.setup: assetCode has already been used"))

		_, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.Error(t, err)
		assert.Equal(t, "confirm transaction error: asset code vUSD has already been used", err.Error())
	})

	t.Run("error, txHelper.ConfirmTx returns an error", func(t *testing.T) {
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
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New("some error has occurred"))
		_, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.Error(t, err)
	})

	t.Run("error, cannot find event log", func(t *testing.T) {
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
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)

		_, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot find setup event from transaction receipt")
	})

	t.Run("error, txHelper.ExtractSetupEvent returns an error", func(t *testing.T) {
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
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Setup(string,bytes32,uint256,bytes32,address)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractSetupEvent("Setup", gomock.AssignableToTypeOf(&types.Log{})).
			Return(nil, errors.New("some error has occurred"))

		_, err := testHelper.Client.SetupCredit(context.Background(), input)

		assert.Error(t, err)
	})
}
