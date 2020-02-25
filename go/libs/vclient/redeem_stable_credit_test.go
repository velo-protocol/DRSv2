package vclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	vabi "github.com/velo-protocol/DRSv2/go/abi"
	"math/big"
	"testing"
)

func TestValidateRedeemStableCredit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.Nil(t, err)
	})

	t.Run("fail, should throw error invalid redeemAmount format", func(t *testing.T) {
		err := (&RedeemStableCreditInput{
			RedeemAmount: "10.12345678",
			AssetCode:    "vUSD",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "invalid redeemAmount format", err.Error())
	})

	t.Run("fail, should throw error redeemAmount must be positive", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "-2", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "redeemAmount must be greater than 0", err.Error())
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
			AssetCode:    "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Redeem(string,uint256,address,bytes32,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRedeemEvent("Redeem", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRedeem{
				AssetCode:              "vUSD",
				StableCreditAmount:     big.NewInt(1040000000),
				CollateralAssetAddress: common.Address{},
				CollateralAssetCode:    [32]byte{},
				CollateralAmount:       big.NewInt(1040000000),
			}, nil)

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

	t.Run("error, Redeem() function returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &RedeemStableCreditInput{
			RedeemAmount: "104",
			AssetCode:    "vUSD",
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
			AssetCode:    "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
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
			AssetCode:    "vUSD",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Redeem(string,uint256,address,bytes32,uint256)")),
						},
					},
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

	t.Run("error, no redeem log emitting", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &RedeemStableCreditInput{
			RedeemAmount: "104",
			AssetCode:    "vUSD",
		}
		abiInput := input.ToAbiInput()

		tx := new(types.Transaction)

		testHelper.MockDRSContract.EXPECT().
			Redeem(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.RedeemAmount, abiInput.AssetCode).
			Return(tx, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(tx), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)

		result, err := testHelper.Client.RedeemStableCredit(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, fmt.Sprintf("cannot find redeem event from transaction receipt %s", tx.Hash().String()), err.Error())
	})
}

func TestRedeemStableCreditReplaceError(t *testing.T) {
	t.Run("parsing an error, stableCredit not exist", func(t *testing.T) {
		err := RedeemStableCreditReplaceError("smart contract call error", &RedeemStableCreditInput{AssetCode: "vUSD"}, errors.New("stableCredit not exist"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "the stable credit vUSD is not found")
		assert.Contains(t, err.Error(), "smart contract call error")
	})
	t.Run("parsing an error, valid price not found", func(t *testing.T) {
		err := RedeemStableCreditReplaceError("smart contract call error", nil, errors.New("valid price not found"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "valid price not found")
		assert.Contains(t, err.Error(), "smart contract call error")
	})
	t.Run("parsing an error, ERC20: burn amount exceeds balance", func(t *testing.T) {
		err := RedeemStableCreditReplaceError("smart contract call error", &RedeemStableCreditInput{AssetCode: "vUSD"}, errors.New("ERC20: burn amount exceeds balance"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "the stable credit vUSD in your address is insufficient")
		assert.Contains(t, err.Error(), "smart contract call error")
	})
	t.Run("parsing an error, any error", func(t *testing.T) {
		err := RedeemStableCreditReplaceError("smart contract call error", nil, errors.New("some error has occured"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "some error has occured")
		assert.Contains(t, err.Error(), "smart contract call error")
	})
}
