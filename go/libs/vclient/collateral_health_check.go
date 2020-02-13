package vclient

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
)

type CollateralHealthCheckInput struct {
	// to be implement more in the future
}

type CollateralHealthCheckAbiInput struct {
	AssetCode string
}

type CollateralHealthCheckAbiOutput struct {
	CollateralAssetCode [32]byte
	RequiredAmount      *big.Int
	PresentAmount       *big.Int
	IndexKey            int
}

type CollateralHealthCheckOutput struct {
	CollateralAssetCode string
	RequiredAmount      string
	PresentAmount       string
}

func (c *Client) CollateralHealthCheck(input *CollateralHealthCheckInput) ([]CollateralHealthCheckOutput, error) {
	var collateralOutputs []CollateralHealthCheckAbiOutput

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

		// 2.4 add abiInput with assetCode get from stableCredit Contract
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: *curAssetCode,
		}
		collateralAssetCode, requiredAmount, presentAmount, err := c.contract.drs.CollateralHealthCheck(
			nil,
			toAbiInput.AssetCode,
		)
		if err != nil {
			return nil, err
		}

		// Append to collateralOutput
		collateralOutputs = append(collateralOutputs, CollateralHealthCheckAbiOutput{
			CollateralAssetCode: collateralAssetCode,
			RequiredAmount:      requiredAmount,
			PresentAmount:       presentAmount,
		})

		// keep the recent of prevStableCreditAddress from loop
		prevStableCreditId = curStableCreditId
	}

	var collateralHealthCheckOutput []CollateralHealthCheckOutput
	outputMap := map[[32]byte]CollateralHealthCheckAbiOutput{}
	for index, collateralOutput := range collateralOutputs {
		collateralAssetCode := collateralOutput.CollateralAssetCode
		requiredAmount := collateralOutput.RequiredAmount
		presentAmount := collateralOutput.PresentAmount

		output, ok := outputMap[collateralOutput.CollateralAssetCode]
		if ok {
			// add
			// get information from outputMap of repeat the CollateralAssetCode
			// calculate requiredAmount to repeat outputMap[output.CollateralAssetCode].RequiredAmount
			// calculate presentAmount to repeat outputMap[output.CollateralAssetCode].PresentAmount
			requiredAmount = new(big.Int).Add(outputMap[output.CollateralAssetCode].RequiredAmount, requiredAmount)
			presentAmount = new(big.Int).Add(outputMap[output.CollateralAssetCode].PresentAmount, presentAmount)

			// put back to repeat of collateralHealthCheckOutput record after calculate requiredAmount and presentAmount
			collateralHealthCheckOutput[outputMap[output.CollateralAssetCode].IndexKey].RequiredAmount = utils.AmountToString(requiredAmount)
			collateralHealthCheckOutput[outputMap[output.CollateralAssetCode].IndexKey].PresentAmount = utils.AmountToString(presentAmount)
		} else {
			// init, add to map
			// information from each loop of new CollateralAssetCode
			output.CollateralAssetCode = collateralAssetCode
			output.RequiredAmount = requiredAmount
			output.PresentAmount = presentAmount
			output.IndexKey = index

			// store for checked map key of CollateralAssetCode in next loop
			outputMap[output.CollateralAssetCode] = output

			// put new record of CollateralHealthCheckOutput
			collateralHealthCheckOutput = append(collateralHealthCheckOutput, CollateralHealthCheckOutput{
				CollateralAssetCode: utils.Byte32ToString(output.CollateralAssetCode),
				RequiredAmount:      utils.AmountToString(output.RequiredAmount),
				PresentAmount:       utils.AmountToString(output.PresentAmount),
			})
		}
	}

	return collateralHealthCheckOutput, nil
}
