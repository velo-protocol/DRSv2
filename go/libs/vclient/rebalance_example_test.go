// +build !unit

package vclient

import (
	"context"
	"fmt"
)

func ExampleClient_Rebalance() {
	client, err := NewClient("http://127.0.0.1:7545", "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0x4Db9c67836A3735f63c0eCe4cFBc486bB80732b0
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1F1247eDEa84dC392C857A7887203a5640f3f2Fd
	})
	if err != nil {
		panic(err)
	}

	output, err := client.Rebalance(context.Background(), &RebalanceInput{})
	if err != nil {
		panic(err)
	}

	for index, rebalanceTransaction := range output.RebalanceTransactions {
		fmt.Printf("Event %d Transaction Hashes: %s", index, rebalanceTransaction.Tx.Hash().String())
		fmt.Printf("Event %d Transaction Receipt: %d", index, rebalanceTransaction.Receipt.Status)
		fmt.Printf("Events %d AssetCode: %s", index, rebalanceTransaction.AssetCode)
		fmt.Printf("Events %d CollateralAssetCode: %s", index, rebalanceTransaction.CollateralAssetCode)
		fmt.Printf("Events %d PresentAmount: %s", index, rebalanceTransaction.PresentAmount)
		fmt.Printf("Events %d RequiredAmount: %s", index, rebalanceTransaction.RequiredAmount)
	}

	// Output:
	// Event 0 Transaction Hashes: 0xf9647f3e917d75f124876928fe60d015763fa807c621bf956167a37653197d4a
	// Event 0 Transaction Receipt: 1
	// Events 0 AssetCode: vUSD
	// Events 0 CollateralAssetCode: VELO
	// Events 0 PresentAmount: 1.1253847
	// Events 0 RequiredAmount: 1.6076923
}
