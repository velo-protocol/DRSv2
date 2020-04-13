package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"regexp"
	"strings"
)

// MintFromStableCreditAmountInput required input fields of mint from stable credit amount
type MintFromStableCreditAmountInput struct {
	AssetCode          string
	StableCreditAmount string
}

// Validation function for MintFromStableCreditAmount. Validates the required struct fields. It returns an error if any of the fields are
// invalid. Otherwise, it returns nil.
func (i *MintFromStableCreditAmountInput) Validate() error {
	if i.AssetCode == "" {
		return errors.New("assetCode must not be blank")
	}

	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,12}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	if i.StableCreditAmount == "" {
		return errors.New("stableCreditAmount must not be blank")
	}

	amount, err := decimal.NewFromString(i.StableCreditAmount)
	if err != nil {
		return errors.Wrap(err, "invalid stableCreditAmount format")
	}
	if !utils.IsDecimalValid(amount) {
		return errors.New("invalid stableCreditAmount format")
	}
	if !amount.IsPositive() {
		return errors.New("stableCreditAmount must be greater than 0")
	}

	return nil
}

type MintFromStableCreditAmountAbiInput struct {
	MintAmount *big.Int
	AssetCode  string
}

func (i *MintFromStableCreditAmountInput) ToAbiInput() *MintFromStableCreditAmountAbiInput {

	MintAmount, _ := utils.StringToAmount(i.StableCreditAmount)
	return &MintFromStableCreditAmountAbiInput{
		MintAmount: MintAmount,
		AssetCode:  i.AssetCode,
	}
}

type MintFromStableCreditAmountEvent struct {
	AssetCode           string
	StableCreditAmount  string
	AssetAddress        string
	CollateralAssetCode string
	CollateralAmount    string
	Raw                 *types.Log
}

func (i *MintFromStableCreditAmountEvent) ToEventOutput(eventAbi *vabi.DigitalReserveSystemMint) {
	i.AssetCode = eventAbi.AssetCode
	i.StableCreditAmount = utils.AmountToString(eventAbi.MintAmount)
	i.AssetAddress = eventAbi.AssetAddress.String()
	i.CollateralAssetCode = utils.Byte32ToString(eventAbi.CollateralAssetCode)
	i.CollateralAmount = utils.AmountToString(eventAbi.CollateralAmount)
	i.Raw = &eventAbi.Raw
}

// MintFromStableCreditAmountCreditOutput output fields of mint from stable credit amount
type MintFromStableCreditAmountCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *MintFromStableCreditAmountEvent
}

func MintFromStableCreditAmountReplaceError(prefix string, abiInput *MintFromStableCreditAmountAbiInput, err error) error {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "caller must be a trusted partner"):
		return errors.Wrap(errors.New("the message sender is not found or does not have sufficient permission to perform mint stable credit"), prefix)
	case strings.Contains(msg, "stableCredit not exist"):
		return errors.Wrap(errors.Errorf("the stable credit %s is not found", abiInput.AssetCode), prefix)
	case strings.Contains(msg, "transfer amount exceeds balance"):
		return errors.Wrap(errors.New("the collateral in your address is insufficient"), prefix)
	case strings.Contains(msg, "the stable credit does not belong to you"):
		return errors.Wrap(errors.Errorf("the stable credit %s does not belong to you", abiInput.AssetCode), prefix)
	case strings.Contains(msg, "valid price not found"):
		return errors.Wrap(errors.New("valid price not found"), prefix)
	default:
		return errors.Wrap(err, prefix)
	}
}

// MintFromStableCreditAmount calls MintFromStableCreditAmount on Velo smart contract.
func (c *Client) MintFromStableCreditAmount(ctx context.Context, input *MintFromStableCreditAmountInput) (*MintFromStableCreditAmountCreditOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	abiInput := input.ToAbiInput()
	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.drs.MintFromStableCreditAmount(opt, abiInput.MintAmount, abiInput.AssetCode)
	if err != nil {
		return nil, MintFromStableCreditAmountReplaceError("smart contract call error", abiInput, err)
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx, opt.From)
	if err != nil {
		return nil, MintFromStableCreditAmountReplaceError("confirm transaction error", abiInput, err)
	}

	eventLog := utils.FindLogEvent(receipt.Logs, "Mint(string,uint256,address,bytes32,uint256)")
	if eventLog == nil {
		return nil, errors.Errorf("cannot find mint event from transaction receipt %s", tx.Hash().String())
	}

	eventAbi, err := c.txHelper.ExtractMintEvent("Mint", eventLog)
	if err != nil {
		return nil, err
	}

	event := new(MintFromStableCreditAmountEvent)

	event.ToEventOutput(eventAbi)

	return &MintFromStableCreditAmountCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   event,
	}, nil

}
