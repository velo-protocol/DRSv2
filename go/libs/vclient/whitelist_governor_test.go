package vclient

import (
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
		mockContract := testHelper()

		tx, err := mockContract.Client.WhitelistGovernor(&WhitelistGovernorInput{Address: governorAddress})

		assert.NoError(t, err)
		assert.NotNil(t, mockContract.Client)
		assert.NotNil(t, tx)
	})

	t.Run("fail, should throw error address must not be blank", func(t *testing.T) {
		mockContract := testHelper()

		tx, err := mockContract.Client.WhitelistGovernor(&WhitelistGovernorInput{Address: ""})

		assert.NotNil(t, mockContract.Client)
		assert.Nil(t, tx)
		assert.Equal(t, "address must not be blank", err.Error())
	})
}
