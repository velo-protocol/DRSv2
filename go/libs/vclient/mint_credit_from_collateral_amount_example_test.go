// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_MintFromCollateralAmount() {
	client, err := NewClient(smartContractUrl, "6bc0e50d09b273bf7cc5193eeee57fe807011b33ab0ea9a05c922f5eac099642", ContractAddress{
		DrsAddress:   "0xCa2ab1d86B115771c7C2fAF17d58E5D258BE0F6F", // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "0x7B047A389b2CDCAf7FE0e1eE04E6a04E77D00Ba8", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})

	if err != nil {
		panic(err)
	}

	result, err := client.MintFromCollateralAmount(context.Background(), &MintFromCollateralAmountInput{
		AssetCode:        "vUSD",
		CollateralAmount: "2",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Mint From Collateral Amount Transaction Hash: ", result.Tx.Hash().String())
	log.Println("Asset Address: ", result.Event.AssetAddress)
	log.Println("Asset Code: ", result.Event.AssetCode)
	log.Println("Collateral Amount: ", result.Event.CollateralAmount)
	log.Println("Collateral Asset Code: ", result.Event.CollateralAssetCode)
	log.Println("Mint Amount: ", result.Event.MintAmount)
	// Output:
	// Mint From Collateral Amount Transaction Hash: 0x0d6658bf79d9e4541de0617a14c17e7d642169d3f00d7a035595136b59f6f8ac
}
