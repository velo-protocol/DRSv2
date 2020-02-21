// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_MintFromStableCreditAmount() {
	client, err := NewClient(smartContractUrl, "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})

	if err != nil {
		panic(err)
	}

	result, err := client.MintFromStableCreditAmount(context.Background(), &MintFromStableCreditAmountInput{
		AssetCode:          "vUSD",
		StableCreditAmount: "2",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Mint From Stable Credit Amount Transaction Hash: ", result.Tx.Hash().String())
	log.Println("Asset Address: ", result.Event.AssetAddress)
	log.Println("Asset Code: ", result.Event.AssetCode)
	log.Println("Collateral Amount: ", result.Event.CollateralAmount)
	log.Println("Collateral Asset Code: ", result.Event.CollateralAssetCode)
	log.Println("Stable Credit Amount: ", result.Event.StableCreditAmount)
	// Output:
	// Mint From Stable Credit Amount Transaction Hash: 0x1afaa6fc22b88f875bb16235128e245589fa460f8325f84ace2d89bb4204204b
	// Asset Address: 0x23Cf6f4656218Bd25733f27aadBEe009A0f6C3Fd
	// Asset Code: vUSD
	// Collateral Amount: 0.2736842
	// Collateral Asset Code: VELO
	// Stable Credit Amount: 2.0000000
}
