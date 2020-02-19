// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_RedeemStableCredit() {
	client, err := NewClient(smartContractUrl, "b68a7d2260076d2e73f295cc0d946c76db7f6d96f184ea00e97376680a29749a", ContractAddress{
		DrsAddress:   "0xE5352EB3a0FA6C897da1F5FfB90d57AbA099F5bB",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "0xF9baEAdFC5e0CAbC0c1E5E88571c2d7a75099Df1", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})
	//client, err := NewClient(smartContractUrl, "<PRIVATE_KEY>", ContractAddress{
	//	DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
	//	HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	//})

	if err != nil {
		panic(err)
	}

	result, err := client.RedeemStableCredit(context.Background(), &RedeemStableCreditInput{
		RedeemAmount:"104",
		AssetCode:  "vUSD",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Redeem Stable Credit Transaction Hash: ", result.Tx.Hash().String())
	log.Println("Asset Code: ", result.Event.AssetCode)
	log.Println("Stable Credit Amount: ", result.Event.StableCreditAmount)
	log.Println("Collateral Asset Address: ", result.Event.CollateralAssetAddress)
	log.Println("Collateral Asset Code: ", result.Event.CollateralAssetCode)
	log.Println("Collateral Amount: ", result.Event.CollateralAmount)
	// Output:
	// Mint From Stable Credit Amount Transaction Hash: 0x1afaa6fc22b88f875bb16235128e245589fa460f8325f84ace2d89bb4204204b
	// Asset Address: 0x23Cf6f4656218Bd25733f27aadBEe009A0f6C3Fd
	// Asset Code: vUSD
	// Collateral Amount: 0.2736842
	// Collateral Asset Code: VELO
	// Mint Amount: 2.0000000
}
