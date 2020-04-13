package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
)

// RebalanceInput required input fields of rebalance
type RebalanceInput struct {
	// add more if need to add more features.
}

type RebalanceTransaction struct {
	AssetCode           string
	CollateralAssetCode string
	RequiredAmount      string
	PresentAmount       string
	Raw                 *types.Log
	Tx                  *types.Transaction
	Receipt             *types.Receipt
}

// RebalanceOutput output fields of rebalance
type RebalanceOutput struct {
	RebalanceTransactions []*RebalanceTransaction
}

func (i *RebalanceTransaction) ToRebalanceOutput(eventAbi *vabi.DigitalReserveSystemRebalance, tx *types.Transaction, receipt *types.Receipt) {
	if eventAbi != nil {
		i.AssetCode = eventAbi.AssetCode
		i.CollateralAssetCode = utils.Byte32ToString(eventAbi.CollateralAssetCode)
		i.RequiredAmount = utils.AmountToString(eventAbi.RequiredAmount)
		i.PresentAmount = utils.AmountToString(eventAbi.PresentAmount)
		i.Raw = &eventAbi.Raw
	}
	i.Tx = tx
	i.Receipt = receipt
}

// Rebalance calls Rebalance on Velo smart contract.
func (c *Client) Rebalance(ctx context.Context, input *RebalanceInput) (*RebalanceOutput, error) {
	rebalanceOutput := &RebalanceOutput{
		RebalanceTransactions: []*RebalanceTransaction{},
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

		receipt, err := c.txHelper.ConfirmTx(ctx, tx, opt.From)
		if err != nil {
			return nil, err
		}
		rebalanceTransaction := new(RebalanceTransaction)
		eventLog := utils.FindLogEvent(receipt.Logs, "Rebalance(string,bytes32,uint256,uint256)")
		if eventLog != nil {

			eventAbi, err := c.txHelper.ExtractRebalanceEvent("Rebalance", eventLog)
			if err != nil {
				return nil, err
			}

			rebalanceTransaction.ToRebalanceOutput(eventAbi, tx, receipt)

			// Append to rebalanceOutput
			rebalanceOutput.RebalanceTransactions = append(rebalanceOutput.RebalanceTransactions, rebalanceTransaction)

		}
		// keep the recent of prevStableCreditAddress from loop
		prevStableCreditId = curStableCreditId
	}

	return rebalanceOutput, nil
}
