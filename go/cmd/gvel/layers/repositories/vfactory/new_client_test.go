package vfactory_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/repositories/vfactory"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
	"testing"
)

func TestVeloFactory_NewClient(t *testing.T) {
	mockedPub, mockedPriv, _ := crypto.GenerateKeypair()

	t.Run("success", func(t *testing.T) {
		vClient, err := vfactory.NewVeloFactory(nil).NewClient(&entity.NewClientInput{
			RpcUrl:     "http://0.0.0.0:7545",
			PrivateKey: mockedPriv,
			ContractAddress: &vclient.ContractAddress{
				DrsAddress:   mockedPub,
				HeartAddress: mockedPub,
			},
		})

		assert.NoError(t, err)
		assert.NotNil(t, vClient)
	})

	t.Run("error, empty rpc url", func(t *testing.T) {
		_, err := vfactory.NewVeloFactory(nil).NewClient(&entity.NewClientInput{
			RpcUrl:     "",
			PrivateKey: mockedPriv,
			ContractAddress: &vclient.ContractAddress{
				DrsAddress:   mockedPub,
				HeartAddress: mockedPub,
			},
		})

		assert.EqualError(t, err, "rpcUrl must not be empty")
	})

	t.Run("error, empty private key", func(t *testing.T) {
		_, err := vfactory.NewVeloFactory(nil).NewClient(&entity.NewClientInput{
			RpcUrl:     "http://0.0.0.0:7545",
			PrivateKey: "",
			ContractAddress: &vclient.ContractAddress{
				DrsAddress:   mockedPub,
				HeartAddress: mockedPub,
			},
		})

		assert.EqualError(t, err, "privateKey must not be empty")
	})

	t.Run("error, empty contract address", func(t *testing.T) {
		_, err := vfactory.NewVeloFactory(nil).NewClient(&entity.NewClientInput{
			RpcUrl:          "http://0.0.0.0:7545",
			PrivateKey:      mockedPriv,
			ContractAddress: nil,
		})

		assert.EqualError(t, err, "contractAddress must not be empty")
	})
}
