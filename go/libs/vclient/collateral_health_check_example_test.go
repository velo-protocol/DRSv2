// +build !unit

package vclient

import (
	"fmt"
)

func ExampleClient_CollateralHealthCheck() {
	client, err := NewClient("<SMART_CONTRACT_NODE_URL>", "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0x4Db9c67836A3735f63c0eCe4cFBc486bB80732b0
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1F1247eDEa84dC392C857A7887203a5640f3f2Fd
	})
	if err != nil {
		panic(err)
	}

	output, err := client.CollateralHealthCheck(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("CollateralAssetCode:", output[0].CollateralAssetCode)
	fmt.Println("CollateralAssetAddress:", output[0].CollateralAssetAddress)
	fmt.Println("RequiredAmount:", output[0].RequiredAmount)
	fmt.Println("PresentAmount:", output[0].PresentAmount)
	// Output:
	// CollateralAssetCode: VELO
	// CollateralAssetAddress: 0xbBa7C1C0222Eb4e323382FA3914fBC9E2dD58A9F
	// RequiredAmount: 0.0000000
	// PresentAmount: 0.0000000
}
