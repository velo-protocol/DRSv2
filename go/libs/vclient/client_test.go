package vclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client, err := NewClient("http://127.0.0.1:7545", privateKey1, ContractAddress{
			drsAddress:   drsAddress,
			heartAddress: heartAddress,
		})

		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("fail, bad rpc url", func(t *testing.T) {
		_, err := NewClient("badshceme://127.0.0.1:7545", privateKey1, ContractAddress{
			drsAddress:   drsAddress,
			heartAddress: heartAddress,
		})

		assert.Contains(t, err.Error(), "fail to initialize eth client")
	})

	t.Run("fail, bad private key", func(t *testing.T) {
		_, err := NewClient("http://127.0.0.1:7545", "bad_private_key", ContractAddress{
			drsAddress:   drsAddress,
			heartAddress: heartAddress,
		})

		assert.Contains(t, err.Error(), "invalid private key format")
	})

	t.Run("fail, bad drs address", func(t *testing.T) {
		_, err := NewClient("http://127.0.0.1:7545", privateKey1, ContractAddress{
			drsAddress:   "",
			heartAddress: heartAddress})

		assert.Contains(t, err.Error(), "invalid drsAddress address format")
	})

	t.Run("fail, bad heart address", func(t *testing.T) {
		_, err := NewClient("http://127.0.0.1:7545", privateKey1, ContractAddress{
			drsAddress:   drsAddress,
			heartAddress: "",
		})

		assert.Contains(t, err.Error(), "invalid heart address format")
	})

}
