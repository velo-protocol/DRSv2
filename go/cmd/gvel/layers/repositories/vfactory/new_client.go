package vfactory

import (
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"github.com/velo-protocol/DRSv2/go/libs/vclient/ivclient"
)

func (v *veloFactory) NewClient(input *entity.NewClientInput) (ivclient.VClient, error) {
	if len(input.RpcUrl) < 1 {
		return nil, errors.New("rpcUrl must not be empty")
	}
	if len(input.PrivateKey) < 1 {
		return nil, errors.New("privateKey must not be empty")
	}
	if input.ContractAddress == nil {
		return nil, errors.New("contractAddress must not be empty")
	}

	client, err := vclient.NewClient(input.RpcUrl, input.PrivateKey, *input.ContractAddress)
	if err != nil {
		return nil, err
	}

	return ivclient.VClient(client), nil
}
