// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_MintFromCollateralAmount() {
	client, err := NewClient(smartContractUrl, "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
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
	log.Println("Stable Credit Amount: ", result.Event.StableCreditAmount)
	// Output:
	// Mint From Collateral Amount Transaction Hash: 0x0d6658bf79d9e4541de0617a14c17e7d642169d3f00d7a035595136b59f6f8ac
	// Asset Address: 0x23Cf6f4656218Bd25733f27aadBEe009A0f6C3Fd
	// Asset Code: vUSD
	// Collateral Amount:  2.0000000
	// Collateral Asset Code: VELO
	// Stable Credit Amount: 4.0000000
}
