package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTxHelper_ConfirmTx(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := testHelperWithMock(t)
		h.MockConnection.EXPECT().
			TransactionReceipt(context.Background(), gomock.AssignableToTypeOf(common.Hash{})).
			Return(&types.Receipt{}, nil)

		result, err := h.TxHelper.ConfirmTx(context.Background(), &types.Transaction{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestTxHelper_ExtractSetupEvent(t *testing.T) {

}
