package entity

import "github.com/velo-protocol/DRSv2/go/libs/vclient"

type NewClientInput struct {
	RpcUrl          string
	PrivateKey      string
	ContractAddress *vclient.ContractAddress
}
