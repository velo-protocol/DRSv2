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

type MintFromCollateralAmountInput struct {
	AssetCode        string
	CollateralAmount string
}

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
	MintAmount          string
	AssetAddress        string
	CollateralAssetCode string
	CollateralAmount    string
	Raw                 *types.Log
}

func (i *MintFromCollateralAmountEvent) ToEventOutput(eventAbi *vabi.DigitalReserveSystemMint) {
	i.AssetCode = eventAbi.AssetCode
	i.MintAmount = utils.AmountToString(eventAbi.MintAmount)
	i.AssetAddress = eventAbi.AssetAddress.String()
	i.CollateralAssetCode = utils.Byte32ToString(eventAbi.CollateralAssetCode)
	i.CollateralAmount = utils.AmountToString(eventAbi.CollateralAmount)
	i.Raw = &eventAbi.Raw
}

type MintFromCollateralAmountCreditOutput struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Event   *MintFromCollateralAmountEvent
}

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
		msg := err.Error()
		switch {
		case strings.Contains(msg, "caller must be a trusted partner"):
			return nil, errors.New("the message sender is not found or does not have sufficient permission to perform mint stable credit")
		case strings.Contains(msg, "stableCredit not exist"):
			return nil, errors.Errorf("the stable credit %s is not found", input.AssetCode)
		case strings.Contains(msg, "transfer amount exceeds balance"):
			return nil, errors.New("the collateral in your address is insufficient")
		case strings.Contains(msg, "the stable credit does not belong to you"):
			return nil, errors.Errorf("the stable credit %s does not belong to you", input.AssetCode)
		default:
			return nil, err
		}
	}

	receipt, err := c.txHelper.ConfirmTx(ctx, tx)
	if err != nil {
		return nil, err
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
