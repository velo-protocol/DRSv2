// +build !unit

package vclient

import (
	"log"
)

func ExampleClient_GetExchangeRate() {
	client, err := NewClient("http://127.0.0.1:7545", privateKey, ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})

	if err != nil {
		panic(err)
	}

	result, err := client.GetExchangeRate(&GetExchangeRateInput{AssetCode: "vUSD"})
	if err != nil {
		panic(err)
	}

	log.Println("AssetCode: ", result.AssetCode)
	log.Println("CollateralAssetCode: ", result.CollateralAssetCode)
	log.Println("priceInCollateralPerAssetUnit: ", result.PriceInCollateralPerAssetUnit)
	// Output:
	// AssetCode: vUSD
	// CollateralAssetCode: VELO
	// priceInCollateralPerAssetUnit: 1
}
