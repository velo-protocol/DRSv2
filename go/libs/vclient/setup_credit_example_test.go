// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_SetupCredit() {
	client, err := NewClient(smartContractUrl, privateKey, ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})
	if err != nil {
		panic(err)
	}

	output, err := client.SetupCredit(context.Background(), &SetupCreditInput{
		CollateralAssetCode: "VELO",
		PeggedCurrency:      "USD",
		AssetCode:           "vUSD",
		PeggedValue:         "1.0",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Transaction Hash: ", output.Tx.Hash().String())
	log.Println("Block Number: ", output.Receipt.BlockNumber.String())
	log.Println("Event.PeggedCurrency: ", output.Event.PeggedCurrency)
	log.Println("Event.PeggedValue: ", output.Event.PeggedValue)
	log.Println("Event.AssetCode: ", output.Event.AssetCode)
	log.Println("Event.CollateralAssetCode: ", output.Event.CollateralAssetCode)
	log.Println("Event.AssetAddress: ", output.Event.AssetAddress)
	// Output:
	// Transaction Hash: 0x0d6658bf79d9e4541de0617a14c17e7d642169d3f00d7a035595136b59f6f8ac
	// Block Number: 10
	// Event.PeggedCurrency: USD
	// Event.PeggedValue: 1
	// Event.AssetCode: vUSD
	// Event.CollateralAssetCode: VELO
	// Event.AssetAddress: 0x50637DeE3598e080B7B605B00f4FfC721046E4E0
}
