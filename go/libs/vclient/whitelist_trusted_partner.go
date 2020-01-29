package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type WhitelistTrustedPartnerInput struct {
	Address string
}

func (i *WhitelistTrustedPartnerInput) Validate() error {
	if !common.IsHexAddress(i.Address) {
		return errors.New("invalid address format")
	}
	return nil
}

type WhitelistTrustedPartnerAbiInput struct {
	Address common.Address
}

func (i *WhitelistTrustedPartnerInput) ToAbiInput() WhitelistTrustedPartnerAbiInput {

	return WhitelistTrustedPartnerAbiInput{
		Address: common.HexToAddress(i.Address),
	}
}

type WhitelistTrustedPartnerOutput struct {
	Tx *types.Transaction
}

func (c *Client) WhitelistTrustedPartner(input *WhitelistTrustedPartnerInput) (*WhitelistTrustedPartnerOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = 5000000
	tx, err := c.Heart().SetTrustedPartner(opt, input.ToAbiInput().Address)
	if err != nil {
		return nil, err
	}

	return &WhitelistTrustedPartnerOutput{
		Tx: tx,
	}, nil
}
