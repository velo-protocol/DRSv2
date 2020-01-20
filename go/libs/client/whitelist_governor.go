package client

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO: implement correctly
type WhitelistGovernorInput struct {
}

func (i *WhitelistGovernorInput) Validate() error {
	return nil
}

func (i *WhitelistGovernorInput) ToAbiInput() common.Address {
	return common.Address{}
}

type WhitelistGovernorOutput struct {
	Tx *types.Transaction
}

func (c *Client) WhitelistGovernor(input *WhitelistGovernorInput) (*WhitelistGovernorOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	opt := bind.NewKeyedTransactor(&c.privateKey)
	tx, err := c.Heart().SetGovernor(opt, input.ToAbiInput())
	if err != nil {
		return nil, err
	}

	return &WhitelistGovernorOutput{
		Tx: tx,
	}, nil
}
