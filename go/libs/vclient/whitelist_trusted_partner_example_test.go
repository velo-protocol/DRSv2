// +build !unit

package vclient

import (
	"context"
	"log"
)

func ExampleClient_WhitelistTrustedPartner() {
	client, err := NewClient("http://127.0.0.1:7545", "da17d295e2fd005747cca4de855bbb0493f2e0669753bba1e752700dbad4c78c", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // Ex: 0x604Ee3d8d9A734d4607E8aF2E4eb44D8e6c2Bf46
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // Ex: 0x1fCc1CEf04A0B0121faD5AF38C06F10c44240787
	})
	if err != nil {
		panic(err)
		return
	}

	result, err := client.WhitelistTrustedPartner(context.Background(), &WhitelistTrustedPartnerInput{
		Address: "<TRUSTED_PARTNER_ADDRESS>", // Ex: 0xf9955C6A38f74f0bfFD63141E840BB77FF3F3d38
	})
	if err != nil {
		panic(err)
		return
	}

	log.Println("Transaction Hash:", result.Tx.Hash().String())
	// Output:
	// Transaction Hash: 0xc49e5ce3e74e03c5d6ed0ba1e4993b03d8ef6f3b7f9690c523e62e22b0a547ed

}
