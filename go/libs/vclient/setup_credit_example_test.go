// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_SetupCredit() {
	client, err := NewClient("http://127.0.0.1:7545", "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0x4Db9c67836A3735f63c0eCe4cFBc486bB80732b0
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1F1247eDEa84dC392C857A7887203a5640f3f2Fd
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
