package utils_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"testing"
)

func TestParseRevertMessage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := utils.ParseRevertMessage([]byte("0x00000000000000000000000000000000000000000000000000000000000000000DigitalReserveSystem._validateAssetCode: stableCredit not exist"))
		assert.NotEmpty(t, result)
		assert.Equal(t, "DigitalReserveSystem._validateAssetCode: stableCredit not exist", result)
	})
}
