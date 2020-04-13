package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/constants"
	"strings"
)

// WhitelistTrustedPartnerInput required input fields of whitelist trusted partner
type WhitelistTrustedPartnerInput struct {
	Address string
}

// Validation function for WhitelistTrustedPartner. Validates the required struct fields. It returns an error if any of the fields are
// invalid. Otherwise, it returns nil.
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

// WhitelistTrustedPartnerOutput output fields of whitelist trusted partner
type WhitelistTrustedPartnerOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
}

func WhitelistTrustedPartnerReplaceError(prefix string, err error) error {
	if strings.Contains(err.Error(), "the message sender is not found or does not have sufficient permission") {
		return errors.Wrap(errors.New("the message sender is not found or does not have sufficient permission to perform whitelist user"), prefix)
	}
	return errors.Wrap(err, prefix)
}

func (c *Client) WhitelistTrustedPartner(ctx context.Context, input *WhitelistTrustedPartnerInput) (*WhitelistTrustedPartnerOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	senderIsGovernor, err := c.contract.heart.IsGovernor(nil, c.publicKey)
	if err != nil {
		return nil, err
	}
	if !senderIsGovernor {
		return nil, errors.New("the message sender is not found or does not have sufficient permission to perform whitelist user")
	}

	isTrustedPartner, err := c.contract.heart.IsTrustedPartner(nil, input.ToAbiInput().Address)
	if err != nil {
		return nil, err
	}
	if isTrustedPartner {
		return nil, errors.Errorf("the address %s has already been whitelisted as trusted partner", input.Address)
	}

	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.heart.SetTrustedPartner(opt, input.ToAbiInput().Address)
	if err != nil {
		return nil, WhitelistTrustedPartnerReplaceError("smart contract call error", err)
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx, opt.From)
	if err != nil {
		return nil, WhitelistTrustedPartnerReplaceError("confirm transaction error", err)
	}

	return &WhitelistTrustedPartnerOutput{
		Tx:      tx,
		Receipt: receipt,
	}, nil
}
