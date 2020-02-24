package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/constants"
	"math/big"
	"testing"
)

func TestTxHelper_ConfirmTx(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := testHelperWithMock(t)
		h.MockConnection.EXPECT().
			TransactionReceipt(context.Background(), gomock.AssignableToTypeOf(common.Hash{})).
			Return(&types.Receipt{
				Status: 1,
			}, nil)

		result, err := h.TxHelper.ConfirmTx(context.Background(), &types.Transaction{}, common.HexToAddress("0x0000000000000000000000000000000000000002"))

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("fail, receipt status = 0 and got error from call contract", func(t *testing.T) {
		fromAddress := common.HexToAddress("0x0000000000000000000000000000000000000002")
		curContext := context.Background()
		blockNumber := big.NewInt(1023)
		h := testHelperWithMock(t)
		h.MockConnection.EXPECT().
			TransactionReceipt(curContext, gomock.AssignableToTypeOf(common.Hash{})).
			Return(&types.Receipt{
				Status:      0,
				BlockNumber: blockNumber,
			}, nil)
		h.MockConnection.EXPECT().CallContract(curContext, gomock.AssignableToTypeOf(ethereum.CallMsg{}), blockNumber).Return(nil, errors.New("some error has occurred"))

		transaction := types.NewTransaction(1, common.HexToAddress("0x0000000000000000000000000000000000000003"), big.NewInt(10000000), uint64(constants.GasLimit), big.NewInt(400000), nil)
		result, err := h.TxHelper.ConfirmTx(context.Background(), transaction, fromAddress)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "fail to get revert message")
	})
}
