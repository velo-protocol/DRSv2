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

func TestMintFromStableCreditAmountInput_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := (&MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}).Validate()

		assert.NoError(t, err)
	})

	t.Run("error, validation fail assetCode must not be blank", func(t *testing.T) {
		err := (&MintFromStableCreditAmountInput{
			StableCreditAmount: "100",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "assetCode must not be blank", err.Error())
	})

	t.Run("error, validation fail invalid assetCode format", func(t *testing.T) {
		err := (&MintFromStableCreditAmountInput{
			AssetCode:          "AssetCodee!@",
			StableCreditAmount: "100",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "invalid assetCode format", err.Error())
	})

	t.Run("error, validation fail stableCreditAmount must not be blank", func(t *testing.T) {
		err := (&MintFromStableCreditAmountInput{
			AssetCode: "vUSD",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "stableCreditAmount must not be blank", err.Error())
	})

	t.Run("error, validation fail invalid stableCreditAmount format", func(t *testing.T) {
		err := (&MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "amount",
		}).Validate()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid stableCreditAmount format")

		err = (&MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100.00000001",
		}).Validate()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid stableCreditAmount format")
	})

	t.Run("error, validation fail stableCreditAmount must be greater than 0", func(t *testing.T) {
		err := (&MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "0",
		}).Validate()

		assert.Error(t, err)
		assert.Equal(t, "stableCreditAmount must be greater than 0", err.Error())
	})
}

func TestMintFromStableCreditAmountInput_ToAbiInput(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		abiInput := (&MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}).ToAbiInput()

		assert.Equal(t, "vUSD", abiInput.AssetCode)
		assert.Equal(t, big.NewInt(1000000000), abiInput.MintAmount)
	})
}

func TestClient_MintFromStableCreditAmount(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()
		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Mint(string,uint256,address,bytes32,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractMintEvent("Mint", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemMint{
				AssetCode:           "vUSD",
				MintAmount:          big.NewInt(1000000000),
				AssetAddress:        common.Address{},
				CollateralAssetCode: [32]byte{},
				CollateralAmount:    big.NewInt(1000000000),
			}, nil)

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, result.Tx)
		assert.NotNil(t, result.Receipt)
		assert.NotNil(t, result.Event)

	})

	t.Run("error, validation fail", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), &MintFromStableCreditAmountInput{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("error, drs.MintFromStableCreditAmount returns an error caller must be a trusted partner", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(nil, errors.New("revert DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner"))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "the message sender is not found or does not have sufficient permission to perform mint stable credit")
	})

	t.Run("error, drs.MintFromStableCreditAmount returns an error stableCredit not exist", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(nil, errors.New("revert DigitalReserveSystem._validateAssetCode: stableCredit not exist"))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "the stable credit vUSD is not found")
	})

	t.Run("error, drs.MintFromStableCreditAmount returns an error transfer amount exceeds balance", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(nil, errors.New("revert ERC20: transfer amount exceeds balance"))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "the collateral in your address is insufficient")
	})

	t.Run("error, drs.MintFromStableCreditAmount returns an error stable credit does not belong to you", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(nil, errors.New("revert DigitalReserveSystem.MintFromStableCreditAmount: the stable credit does not belong to you"))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "the stable credit vUSD does not belong to you")
	})

	t.Run("error, drs.MintFromStableCreditAmount returns an error valid price not found", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(nil, errors.New("DigitalReserveSystem._validateAssetCode: valid price not found"))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "valid price not found")
	})

	t.Run("error, drs.MintFromStableCreditAmount returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(nil, errors.New(expectedMsg))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("error, txHelper.ConfirmTx returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New(expectedMsg))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("error, txHelper.ExtractMintEvent returns an error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		expectedMsg := "some error has occurred"

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Mint(string,uint256,address,bytes32,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractMintEvent("Mint", gomock.AssignableToTypeOf(&types.Log{})).
			Return(nil, errors.New(expectedMsg))

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("error, no mint log emitting", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &MintFromStableCreditAmountInput{
			AssetCode:          "vUSD",
			StableCreditAmount: "100",
		}
		abiInput := input.ToAbiInput()

		tx := new(types.Transaction)

		testHelper.MockDRSContract.EXPECT().
			MintFromStableCreditAmount(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.MintAmount, abiInput.AssetCode).
			Return(tx, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(tx), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)

		result, err := testHelper.Client.MintFromStableCreditAmount(context.Background(), input)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, fmt.Sprintf("cannot find mint event from transaction receipt %s", tx.Hash().String()), err.Error())
	})
}
