package vclient

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"strings"
)

// CollateralHealthCheckInput required input fields of collateral health check
type CollateralHealthCheckInput struct {
	// to be implement more in the future
}

type CollateralHealthCheckAbiInput struct {
	AssetCode string
}

type CollateralHealthCheckAbiOutput struct {
	CollateralAssetAddress common.Address
	CollateralAssetCode    [32]byte
	RequiredAmount         *big.Int
	PresentAmount          *big.Int
	IndexKey               int
}

// CollateralHealthCheckOutput output fields of collateral health check
type CollateralHealthCheckOutput struct {
	CollateralAssetAddress string
	CollateralAssetCode    string
	RequiredAmount         string
	PresentAmount          string
}

// CollateralHealthCheck calls CollateralHealthCheck on Velo smart contract.
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
	var prevStableCreditId [32]byte
	// 2. loop StableCredit from stableCreditSize
	for index := 0; index < int(stableCreditSize); index++ {
		if index == 0 {
			// 2.1 get recent StableCredit from linklist in Heart Contract
			curStableCreditAddress, err = c.contract.heart.GetRecentStableCredit(nil)
			if err != nil {
				return nil, err
			}
		} else {
			// 2.2 get next StableCredit for each stableCreditId
			curStableCreditAddress, err = c.contract.heart.GetNextStableCredit(nil, prevStableCreditId)
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
		collateralAddress, collateralAssetCode, requiredAmount, presentAmount, err := c.contract.drs.CollateralHealthCheck(
			nil,
			toAbiInput.AssetCode,
		)
		if err != nil {
			msg := err.Error()
			if strings.Contains(msg, "valid price not found") ||
				strings.Contains(msg, "invalid price") ||
				strings.Contains(msg, "price is not active") ||
				strings.Contains(msg, "price has an error") {
				return nil, errors.Wrap(errors.New("valid price not found"), "smart contract call error")
			}
			return nil, err
		}

		// Append to collateralOutput
		collateralOutputs = append(collateralOutputs, CollateralHealthCheckAbiOutput{
			CollateralAssetAddress: collateralAddress,
			CollateralAssetCode:    collateralAssetCode,
			RequiredAmount:         requiredAmount,
			PresentAmount:          presentAmount,
		})

		// keep the recent of prevStableCreditAddress from loop
		prevStableCreditId = *curStableCreditId
	}

	var collateralHealthCheckOutput []CollateralHealthCheckOutput
	outputMap := map[[32]byte]CollateralHealthCheckAbiOutput{}
	for _, collateralOutput := range collateralOutputs {
		collateralAssetAddress := collateralOutput.CollateralAssetAddress
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
			*outputMap[output.CollateralAssetCode].RequiredAmount = *requiredAmount
			*outputMap[output.CollateralAssetCode].PresentAmount = *presentAmount

			// put back to repeat of collateralHealthCheckOutput record after calculate requiredAmount and presentAmount
			if len(collateralHealthCheckOutput) == 0 {
				return nil, errors.New("internal error")
			}
			collateralHealthCheckOutput[outputMap[output.CollateralAssetCode].IndexKey].RequiredAmount = utils.AmountToString(requiredAmount)
			collateralHealthCheckOutput[outputMap[output.CollateralAssetCode].IndexKey].PresentAmount = utils.AmountToString(presentAmount)
		} else {
			// init, add to map
			// information from each loop of new CollateralAssetCode
			output.CollateralAssetAddress = collateralAssetAddress
			output.CollateralAssetCode = collateralAssetCode
			output.RequiredAmount = requiredAmount
			output.PresentAmount = presentAmount
			output.IndexKey = len(collateralHealthCheckOutput)

			// store for checked map key of CollateralAssetCode in next loop
			outputMap[output.CollateralAssetCode] = output

			// put new record of CollateralHealthCheckOutput
			collateralHealthCheckOutput = append(collateralHealthCheckOutput, CollateralHealthCheckOutput{
				CollateralAssetAddress: output.CollateralAssetAddress.String(),
				CollateralAssetCode:    utils.Byte32ToString(output.CollateralAssetCode),
				RequiredAmount:         utils.AmountToString(output.RequiredAmount),
				PresentAmount:          utils.AmountToString(output.PresentAmount),
			})
		}
	}

	return collateralHealthCheckOutput, nil
}
