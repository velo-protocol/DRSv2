// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_RedeemStableCredit() {
	client, err := NewClient(smartContractUrl, "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})

	if err != nil {
		panic(err)
	}

	result, err := client.RedeemStableCredit(context.Background(), &RedeemStableCreditInput{
		RedeemAmount: "104",
		AssetCode:    "vUSD",
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
	//Output:
	//Redeem Stable Credit Transaction Hash:  0xf9647f3e917d75f124876928fe60d015763fa807c621bf956167a37653197d4a
	//Asset Code:  vUSD
	//Stable Credit Amount:  104.0000000
	//Collateral Asset Address:  0x11590aB398DD2fedC2C40b6E9F02fa13cC4d2dEe
	//Collateral Asset Code:  VELO
	//Collateral Amount:  0.0000000
}
