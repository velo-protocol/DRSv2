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

func TestValidateWhitelistGovernor(t *testing.T) {
	t.Run("fail, should throw error address must not be blank", func(t *testing.T) {
		whitelistGovernorInput := &WhitelistGovernorInput{Address: ""}
		err := whitelistGovernorInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "address must not be blank", err.Error())
	})

	t.Run("fail, should throw error invalid address format", func(t *testing.T) {
		whitelistGovernorInput := &WhitelistGovernorInput{Address: invalidAddress}
		err := whitelistGovernorInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "invalid address format", err.Error())
	})
}

func TestValidateWhitelistGovernorToAbiInput(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		whitelistGovernorInput := &WhitelistGovernorInput{Address: governorAddress}
		abiInput := whitelistGovernorInput.ToAbiInput()

		assert.NotNil(t, abiInput)
	})
}

func TestWhitelistGovernor(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetGovernor(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(&types.Transaction{}, nil)

		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NoError(t, err)
		assert.NotNil(t, result.Tx)
		assert.NotNil(t, result.Receipt)
	})

	t.Run("fail, should throw error address must not be blank", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: "",
		}

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "address must not be blank", err.Error())
	})

	t.Run("fail, should throw error the message sender is not found or does not have sufficient permission to perform whitelist user", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(false, nil)

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "the message sender is not found or does not have sufficient permission to perform whitelist user", err.Error())
	})

	t.Run("fail, should throw error call heart.contract.IsGovernor validate sender", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(false, errors.New("error here"))

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, should throw error The address ${address} has already been whitelisted as governor", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(true, nil)

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "the address 0x0000000000000000000000000000000000000003 has already been whitelisted as governor", err.Error())
	})

	t.Run("fail, should throw error call heart.contract.IsGovernor", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(true, errors.New("error here"))

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, should throw error call heart.contract.SetGovernor", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetGovernor(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, should throw error The message sender is not found or does not have sufficient permission to perform whitelist user from heart.contract.SetGovernor", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetGovernor(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(nil, errors.New("the message sender is not found or does not have sufficient permission"))

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "smart contract call error: the message sender is not found or does not have sufficient permission to perform whitelist user", err.Error())
	})

	t.Run("fail, should throw error txHelper.ConfirmTx", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: governorAddress,
		}
		abiInput := input.ToAbiInput()

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), testHelper.Client.publicKey).
			Return(true, nil)

		testHelper.MockHeartContract.EXPECT().
			IsGovernor(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.Address).
			Return(false, nil)

		testHelper.MockHeartContract.EXPECT().
			SetGovernor(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(&types.Transaction{}, nil)

		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.WhitelistGovernor(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
