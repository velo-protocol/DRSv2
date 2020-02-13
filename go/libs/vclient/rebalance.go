package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/constants"
)

type RebalanceInput struct {
	// add more if need to add more features.
}

type RebalanceOutput struct {
	Message  string
	Txs      []*types.Transaction
	Receipts []*types.Receipt
}

func (c *Client) Rebalance(ctx context.Context, input *RebalanceInput) (*RebalanceOutput, error) {
	rebalanceOutput := &RebalanceOutput{
		Message:  "rebalance process completed.",
		Txs:      []*types.Transaction{},
		Receipts: []*types.Receipt{},
	}

	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit

	// 1. get count all StableCredit in Heart Contract
	stableCreditSize, err := c.contract.heart.GetStableCreditCount(nil)
	if err != nil {
		return nil, err
	}

	if stableCreditSize == 0 {
		return nil, errors.New("stable credit not found")
	}

	var curStableCreditAddress common.Address
	var curAssetCode *string
	var curStableCreditId *[32]byte
	var prevStableCreditId *[32]byte
	// 2. loop StableCredit from stableCreditSize
	for i := 0; i < int(stableCreditSize); i++ {
		if i == 0 {
			// 2.1 get recent StableCredit from linklist in Heart Contract
			curStableCreditAddress, err = c.contract.heart.GetRecentStableCredit(nil)
			if err != nil {
				return nil, err
			}
		} else {
			// 2.2 get next StableCredit for each stableCreditId
			curStableCreditAddress, err = c.contract.heart.GetNextStableCredit(nil, *prevStableCreditId)
			if err != nil {
				return nil, err
			}
		}

		// 2.3 get stable credit contract for get asset code each stable credit address
		curAssetCode, curStableCreditId, err = c.txHelper.StableCreditAssetCode(curStableCreditAddress)
		if err != nil {
			return nil, err
		}

		tx, err := c.contract.drs.Rebalance(
			opt,
			*curAssetCode,
		)
		if err != nil {
			return nil, err
		}

		receipt, err := c.txHelper.ConfirmTx(ctx, tx)
		if err != nil {
			return nil, err
		}

		// Append to rebalanceOutput Tx, Receipt
		rebalanceOutput.Txs = append(rebalanceOutput.Txs, tx)
		rebalanceOutput.Receipts = append(rebalanceOutput.Receipts, receipt)

		// keep the recent of prevStableCreditAddress from loop
		prevStableCreditId = curStableCreditId
	}

	return rebalanceOutput, nil
}
