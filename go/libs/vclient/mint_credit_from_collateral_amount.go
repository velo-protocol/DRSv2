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

// MintFromCollateralAmountInput required input fields of mint from collateral amount
type MintFromCollateralAmountInput struct {
	AssetCode        string
	CollateralAmount string
}

// Validation function for MintFromCollateralAmount. Validates the required struct fields. It returns an error if any of the fields are
// invalid. Otherwise, it returns nil.
func (i *MintFromCollateralAmountInput) Validate() error {
	if i.AssetCode == "" {
		return errors.New("assetCode must not be blank")
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9]{1,12}$`, i.AssetCode); !matched {
		return errors.New("invalid assetCode format")
	}

	if i.CollateralAmount == "" {
		return errors.New("collateralAmount must not be blank")
	}

	amount, err := decimal.NewFromString(i.CollateralAmount)
	if err != nil {
		return errors.Wrap(err, "invalid collateralAmount format")
	}
	if !utils.IsDecimalValid(amount) {
		return errors.New("invalid collateralAmount format")
	}
	if !amount.IsPositive() {
		return errors.New("collateralAmount must be greater than 0")
	}

	return nil
}

type MintFromCollateralAmountAbiInput struct {
	NetCollateralAmount *big.Int
	AssetCode           string
}

func (i *MintFromCollateralAmountInput) ToAbiInput() *MintFromCollateralAmountAbiInput {
	netAmount, _ := utils.StringToAmount(i.CollateralAmount)
	return &MintFromCollateralAmountAbiInput{
		NetCollateralAmount: netAmount,
		AssetCode:           i.AssetCode,
	}
}

type MintFromCollateralAmountEvent struct {
	AssetCode           string
	StableCreditAmount  string
	AssetAddress        string
	CollateralAssetCode string
	CollateralAmount    string
	Raw                 *types.Log
}

func (i *MintFromCollateralAmountEvent) ToEventOutput(eventAbi *vabi.DigitalReserveSystemMint) {
	i.AssetCode = eventAbi.AssetCode
	i.StableCreditAmount = utils.AmountToString(eventAbi.MintAmount)
	i.AssetAddress = eventAbi.AssetAddress.String()
	i.CollateralAssetCode = utils.Byte32ToString(eventAbi.CollateralAssetCode)
	i.CollateralAmount = utils.AmountToString(eventAbi.CollateralAmount)
	i.Raw = &eventAbi.Raw
}

// MintFromCollateralAmountCreditOutput output fields of mint from collateral amount
type MintFromCollateralAmountCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *MintFromCollateralAmountEvent
}

func MintFromCollateralAmountReplaceError(prefix string, abiInput *MintFromCollateralAmountAbiInput, err error) error {
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

// MintFromCollateralAmount calls MintFromCollateralAmount on Velo smart contract.
func (c *Client) MintFromCollateralAmount(ctx context.Context, input *MintFromCollateralAmountInput) (*MintFromCollateralAmountCreditOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	abiInput := input.ToAbiInput()
	opt := bind.NewKeyedTransactor(&c.privateKey)
	opt.GasLimit = constants.GasLimit
	tx, err := c.contract.drs.MintFromCollateralAmount(opt, abiInput.NetCollateralAmount, abiInput.AssetCode)
	if err != nil {
		return nil, MintFromCollateralAmountReplaceError("smart contract call error", abiInput, err)
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx, opt.From)
	if err != nil {
		return nil, MintFromCollateralAmountReplaceError("confirm transaction error", abiInput, err)
	}

	eventLog := utils.FindLogEvent(receipt.Logs, "Mint(string,uint256,address,bytes32,uint256)")
	if eventLog == nil {
		return nil, errors.Errorf("cannot find mint event from transaction receipt %s", tx.Hash().String())
	}

	eventAbi, err := c.txHelper.ExtractMintEvent("Mint", eventLog)
	if err != nil {
		return nil, err
	}

	event := new(MintFromCollateralAmountEvent)

	event.ToEventOutput(eventAbi)

	return &MintFromCollateralAmountCreditOutput{
		Tx:      tx,
		Receipt: receipt,
		Event:   event,
	}, nil
}
