// +build !unit

package vclient

import (
	"context"
	"log"
)

const (
	privateKey = "<PRIVATE_KEY>" // b673aace6739646820330920307288260703487da63525f944c96039931d8ed2
)

func ExampleClient_WhitelistGovernor() {
	client, err := NewClient(smartContractUrl, privateKey, ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0xBdA518a6245480652d1A217192EBB299C94F623f
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1623C9c8600319E7CfAff0Ca1c4a05e1a61D954D
	})

	if err != nil {
		panic(err)
	}

	result, err := client.WhitelistGovernor(context.Background(), &WhitelistGovernorInput{Address: "0xcB74De0a5c8E24E81bEd76D57C4D89fcAfa7B6f"})
	if err != nil {
		panic(err)
	}

	log.Println("Whitelist Governor Transaction Hash: ", result.Tx.Hash().String())
	// Output:
	// Whitelist Governor Transaction Hash: 0x0d6658bf79d9e4541de0617a14c17e7d642169d3f00d7a035595136b59f6f8ac
}
