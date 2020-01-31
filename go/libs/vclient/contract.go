package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/velo-protocol/DRSv2/go/abi"
	"strings"
)

type Contract struct {
	drs      DRSContract
	heart    HeartContract
	drsAbi   *abi.ABI
	heartAbi *abi.ABI
}

func NewContract(drs DRSContract, heart HeartContract) *Contract {
	drsAbi, _ := abi.JSON(strings.NewReader(vabi.DigitalReserveSystemABI))
	heartAbi, _ := abi.JSON(strings.NewReader(vabi.HeartABI))

	return &Contract{
		drs:      drs,
		heart:    heart,
		drsAbi:   &drsAbi,
		heartAbi: &heartAbi,
	}
}

func (c *Contract) DRS() DRSContract {
	return c.drs
}

func (c *Contract) Heart() HeartContract {
	return c.heart
}
