package vclient

import (
	"context"
	"github.com/Evrynetlabs/evrynet-node"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTxHelper_ConfirmTx(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := testHelperWithMock(t)
		defer h.Done()

		h.MockConnection.EXPECT().
			TransactionReceipt(context.Background(), gomock.AssignableToTypeOf(common.Hash{})).
			Return(&types.Receipt{Status: 1}, nil)

		result, err := h.TxHelper.ConfirmTx(context.Background(), NewMockedTx(), common.Address{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
	t.Run("fail, receipt status 0, but cannot get revert message", func(t *testing.T) {
		h := testHelperWithMock(t)
		defer h.Done()
		h.MockConnection.EXPECT().
			TransactionReceipt(context.Background(), gomock.AssignableToTypeOf(common.Hash{})).
			Return(&types.Receipt{Status: 0}, nil)
		h.MockConnection.EXPECT().
			CallContract(context.Background(), gomock.AssignableToTypeOf(evrynet.CallMsg{}), nil).
			Return(nil, errors.New("some error has occurred"))

		_, err := h.TxHelper.ConfirmTx(context.Background(), NewMockedTx(), common.Address{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "fail to get revert message")
	})
	t.Run("fail, receipt status 0, with revert message", func(t *testing.T) {
		h := testHelperWithMock(t)
		defer h.Done()
		h.MockConnection.EXPECT().
			TransactionReceipt(context.Background(), gomock.AssignableToTypeOf(common.Hash{})).
			Return(&types.Receipt{Status: 0}, nil)
		h.MockConnection.EXPECT().
			CallContract(context.Background(), gomock.AssignableToTypeOf(evrynet.CallMsg{}), nil).
			Return([]byte("DRS.setup: some error has occurred"), nil)

		_, err := h.TxHelper.ConfirmTx(context.Background(), NewMockedTx(), common.Address{})

		assert.EqualError(t, err, "DRS.setup: some error has occurred")
	})
}
