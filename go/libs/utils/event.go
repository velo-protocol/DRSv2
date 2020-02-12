package utils

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func FindLogEvent(logs []*types.Log, event string) *types.Log {
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}
		if log.Topics[0].String() == crypto.Keccak256Hash([]byte(event)).String() {
			return log
		}
	}
	return nil
}
