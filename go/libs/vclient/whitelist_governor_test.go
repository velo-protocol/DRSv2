package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
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
			SetGovernor(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(&types.Transaction{}, nil)

		result, err := testHelper.Client.WhitelistGovernor(input)

		assert.NoError(t, err)
		assert.NotNil(t, result.Tx)
	})

	t.Run("fail, should throw error address must not be blank", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		input := &WhitelistGovernorInput{
			Address: "",
		}

		result, err := testHelper.Client.WhitelistGovernor(input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "address must not be blank", err.Error())
	})
}
