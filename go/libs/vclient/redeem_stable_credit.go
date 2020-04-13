package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"regexp"
	"strings"
)

// RedeemStableCreditInput required input fields of redeem stable credit
type RedeemStableCreditInput struct {
	RedeemAmount string
	AssetCode    string
}

// Validation function for RedeemStableCredit. Validates the required struct fields. It returns an error if any of the fields are
// invalid. Otherwise, it returns nil.
func (i *RedeemStableCreditInput) Validate() error {
	// validate RedeemAmount
	if i.RedeemAmount == "" {
		return errors.New("redeemAmount must not be blank")
	}

	amount, err := decimal.NewFromString(i.RedeemAmount)
	if err != nil {
		return errors.Wrap(err, "invalid redeemAmount format")
	}
	if !utils.IsDecimalValid(amount) {
		return errors.New("invalid redeemAmount format")
	}
	if !amount.IsPositive() {
		return errors.New("redeemAmount must be greater than 0")
	}

	// validate AssetCode
	if len(i.AssetCode) == 0 {
		return errors.New("assetCode must not be blank")
	}

	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,12}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	return nil
}

type RedeemStableCreditAbiInput struct {
	RedeemAmount *big.Int
	AssetCode    string
}

func (i *RedeemStableCreditInput) ToAbiInput() *RedeemStableCreditAbiInput {
	redeemAmount, _ := utils.StringToAmount(i.RedeemAmount)
	return &RedeemStableCreditAbiInput{
		AssetCode:    i.AssetCode,
		RedeemAmount: redeemAmount,
	}
}

type RedeemStableCreditEvent struct {
	// all fields must be string formatted
	AssetCode              string
	StableCreditAmount     string
	CollateralAssetAddress string
	CollateralAssetCode    string
	CollateralAmount       string
	Raw                    *types.Log
}

// RedeemStableCreditOutput output fields of redeem stable credit
type RedeemStableCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *RedeemStableCreditEvent
}

// RedeemStableCredit calls Redeem on Velo smart contract.
func (c *Client) RedeemStableCredit(ctx context.Context, input *RedeemStableCreditInput) (*RedeemStableCreditOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	abiInput := input.ToAbiInput()
	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit

	tx, err := c.contract.drs.Redeem(
		opt,
		abiInput.RedeemAmount,
		abiInput.AssetCode,
	)
	if err != nil {
		return nil, RedeemStableCreditReplaceError("smart contract call error", input, err)
	}
	receipt, err := c.txHelper.ConfirmTx(ctx, tx, opt.From)
	if err != nil {
		return nil, RedeemStableCreditReplaceError("confirm transaction error", input, err)
	}

	eventLog := utils.FindLogEvent(receipt.Logs, "Redeem(string,uint256,address,bytes32,uint256)")
	if eventLog == nil {
		return nil, errors.Errorf("cannot find redeem event from transaction receipt %s", tx.Hash().String())
	}

	event, err := c.txHelper.ExtractRedeemEvent("Redeem", eventLog)
	if err != nil {
		return nil, err
	}

	redeemStableCreditEvent := &RedeemStableCreditEvent{
		AssetCode:              event.AssetCode,
		StableCreditAmount:     utils.AmountToString(event.StableCreditAmount),
		CollateralAssetAddress: event.CollateralAssetAddress.String(),
		CollateralAssetCode:    utils.Byte32ToString(event.CollateralAssetCode),
		CollateralAmount:       utils.AmountToString(event.CollateralAmount),
		Raw:                    &event.Raw,
	}
	return &RedeemStableCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   redeemStableCreditEvent,
	}, nil
}

func RedeemStableCreditReplaceError(prefix string, abiInput *RedeemStableCreditInput, err error) error {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "stableCredit not exist"):
		return errors.Wrap(errors.Errorf("the stable credit %s is not found", abiInput.AssetCode), prefix)
	case strings.Contains(msg, "valid price not found"):
		return errors.Wrap(errors.New("valid price not found"), prefix)
	case strings.Contains(msg, "ERC20: burn amount exceeds balance"):
		return errors.Wrap(errors.Errorf("the stable credit %s in your address is insufficient", abiInput.AssetCode), prefix)
	default:
		return errors.Wrap(err, prefix)
	}
}
