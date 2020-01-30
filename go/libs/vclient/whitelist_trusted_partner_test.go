package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
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
			SetTrustedPartner(gomock.AssignableToTypeOf(&bind.TransactOpts{}), abiInput.Address).
			Return(&types.Transaction{}, nil)

		result, err := testHelper.Client.WhitelistTrustedPartner(whitelistTrustedPartnerInput)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Tx.Hash())
	})

	t.Run("error, validation fail", func(t *testing.T) {
		whitelistTrustedPartnerInput := &WhitelistTrustedPartnerInput{Address: invalidAddress}

		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		result, err := testHelper.Client.WhitelistTrustedPartner(whitelistTrustedPartnerInput)
		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, "invalid address format", err.Error())
	})

}
