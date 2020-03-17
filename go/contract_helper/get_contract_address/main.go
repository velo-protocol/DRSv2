package get_contract_address

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Addresses struct {
	DrsAddress            string `json:"drsAddress"`
	HeartAddress          string `json:"heartAddress"`
	ReserveManagerAddress string `json:"reserveManagerAddress"`
	CollateralAddress     string `json:"collateralAddress"`
	PriceFeederAddress    string `json:"priceFeederAddress"`
}

func GetContractAddresses() *Addresses {

	byteValue, err := ioutil.ReadFile("../contract_addresses.json")
	if err != nil {
		log.Fatal(err)
	}
	addresses := new(Addresses)
	err = json.Unmarshal(byteValue, addresses)
	if err != nil {
		log.Fatal(err)
	}

	return addresses
}
