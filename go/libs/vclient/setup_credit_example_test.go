// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_SetupCredit() {
	client, err := NewClient("http://52.220.128.212:22001", "3dd34af7357715c3e550153b0a88824b2c42e00bf745355cec395d40b24d2091", ContractAddress{
		DrsAddress:   "0x658280e4FD6454f63f0D6056735f8c6c2C62278c",
		HeartAddress: "0x414f380DfC7efD76eD757a901A0530aa6895d62d",
	})
	if err != nil {
		panic(err)
	}

	output, err := client.SetupCredit(context.Background(), &SetupCreditInput{
		CollateralAssetCode: "VELOe",
		PeggedCurrency:      "USD",
		AssetCode:           "vUSD4",
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
