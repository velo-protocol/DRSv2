package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/constants"
)

type WhitelistGovernorInput struct {
	Address string
}

type WhitelistGovernorOutput struct {
	Tx *types.Transaction
}

type WhitelistGovernorAbiInput struct {
	Address common.Address
}

func (i *WhitelistGovernorInput) Validate() error {
	if len(i.Address) == 0 {
		return errors.New("address must not be blank")
	}

	if !common.IsHexAddress(i.Address) {
		return errors.New("invalid address format")
	}

	return nil
}

func (i *WhitelistGovernorInput) ToAbiInput() WhitelistGovernorAbiInput {
	return WhitelistGovernorAbiInput{
		Address: common.HexToAddress(i.Address),
	}
}

func (c *Client) WhitelistGovernor(input *WhitelistGovernorInput) (*WhitelistGovernorOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.heart.SetGovernor(opt, input.ToAbiInput().Address)
	if err != nil {
		return nil, err
	}

	return &WhitelistGovernorOutput{
		Tx: tx,
	}, nil
}
