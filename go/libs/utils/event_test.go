package utils_test

import (
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/Evrynetlabs/evrynet-node/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"testing"
)

func TestFindLogEvent(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		event := "Mint(string,uint256,address,bytes32,uint256)"
		logs := []*types.Log{
			{
				Topics: []common.Hash{
					crypto.Keccak256Hash([]byte(event)),
				},
			},
		}
		result := utils.FindLogEvent(logs, event)
		assert.NotEmpty(t, result)
	})

	t.Run("fail, event not match", func(t *testing.T) {
		event := "Mint(string,uint256,address,bytes32,uint256)"
		logs := []*types.Log{
			{
				Topics: []common.Hash{
					crypto.Keccak256Hash([]byte("Setup(string,bytes32)")),
				},
			},
		}
		result := utils.FindLogEvent(logs, event)
		assert.Empty(t, result)
	})

	t.Run("fail, no topics in log", func(t *testing.T) {
		event := "Mint(string,uint256,address,bytes32,uint256)"
		logs := []*types.Log{
			{},
		}
		result := utils.FindLogEvent(logs, event)
		assert.Empty(t, result)
	})
}
