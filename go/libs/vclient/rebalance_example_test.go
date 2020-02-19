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

	fmt.Println("Message:", output.Message)
	fmt.Println("Transaction Hashes:", output.Txs)
	fmt.Println("Transaction Receipts:", output.Receipts)
	fmt.Println("Results 0 AssetCode:", output.Events[0].AssetCode)
	fmt.Println("Results 0 CollateralAssetCode:", output.Events[0].CollateralAssetCode)
	fmt.Println("Results 0 PresentAmount:", output.Events[0].PresentAmount)
	fmt.Println("Results 0 RequiredAmount:", output.Events[0].RequiredAmount)
	// Output:
	// Message: rebalance process completed.
	// Transaction Hashes: [0xc000476240]
	// Transaction Receipts: [0xc000460a80]
	// Results 0 AssetCode: vUSD
	// Results 0 CollateralAssetCode: VELO
	// Results 0 PresentAmount: 1.1253847
	// Results 0 RequiredAmount: 1.6076923
}
