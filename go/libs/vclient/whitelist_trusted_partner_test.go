package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhitelistTrustedPartnerInput_Validate(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		whitelistTrustedPartnerInput := WhitelistTrustedPartnerInput{Address: trustedPartnerAddress}
		err := whitelistTrustedPartnerInput.Validate()
		assert.NoError(t, err)
	})

	t.Run("error, validation fail invalid address format", func(t *testing.T) {
		whitelistTrustedPartnerInput := WhitelistTrustedPartnerInput{Address: invalidAddress}
		err := whitelistTrustedPartnerInput.Validate()
		assert.Error(t, err)
		assert.Equal(t, "invalid address format", err.Error())
	})
}

func TestWhitelistTrustedPartnerInput_ToAbiInput(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		whitelistTrustedPartnerInput := WhitelistTrustedPartnerInput{Address: trustedPartnerAddress}
		abiInput := whitelistTrustedPartnerInput.ToAbiInput()
		assert.Equal(t, common.HexToAddress(trustedPartnerAddress), abiInput.Address)
	})

}

func TestClient_WhitelistTrustedPartner(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		whitelistTrustedPartnerInput := &WhitelistTrustedPartnerInput{Address: trustedPartnerAddress}

		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		abiInput := whitelistTrustedPartnerInput.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsTrustedPartner(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetTrustedPartner(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(&types.Transaction{}, nil)

		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), whitelistTrustedPartnerInput)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Tx.Hash())
		assert.NotNil(t, result.Receipt)
	})

	t.Run("error, validation fail", func(t *testing.T) {
		whitelistTrustedPartnerInput := &WhitelistTrustedPartnerInput{Address: invalidAddress}

		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), whitelistTrustedPartnerInput)
		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, "invalid address format", err.Error())
	})

	t.Run("fail, should throw error The address ${address} has already been whitelisted as trusted partner", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(false, nil)

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "the message sender is not found or does not have sufficient permission to perform whitelist user", err.Error())
	})

	t.Run("fail, should throw error call heart.contract.IsGovernor validate sender", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(false, errors.New("error here"))

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, should throw error The address ${address} has already been whitelisted as trusted partner", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsTrustedPartner(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(true, nil)

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "the address 0x0000000000000000000000000000000000000003 has already been whitelisted as trusted partner", err.Error())
	})

	t.Run("fail, should throw error call heart.contract.IsTrustedPartner", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsTrustedPartner(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(true, errors.New("error here"))

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, should throw error call heart.contract.SetTrustedPartner", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsTrustedPartner(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetTrustedPartner(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, should throw error The message sender is not found or does not have sufficient permission to perform whitelist user from heart.contract.SetTrustedPartner", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsTrustedPartner(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetTrustedPartner(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(nil, errors.New("the message sender is not found or does not have sufficient permission"))

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "smart contract call error: the message sender is not found or does not have sufficient permission to perform whitelist user", err.Error())
	})

	t.Run("fail, should throw error txHelper.ConfirmTx", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistTrustedPartnerInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsTrustedPartner(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetTrustedPartner(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(&types.Transaction{}, nil)

		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.WhitelistTrustedPartner(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
