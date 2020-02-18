package ivclient

import (
	"context"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)

type VClient interface {
	SetupCredit(ctx context.Context, input *vclient.SetupCreditInput) (*vclient.SetupCreditOutput, error)
	MintFromCollateralAmount(ctx context.Context, input *vclient.MintFromCollateralAmountInput) (*vclient.MintFromCollateralAmountCreditOutput, error)
	MintFromStableCreditAmount(ctx context.Context, input *vclient.MintFromStableCreditAmountInput) (*vclient.MintFromStableCreditAmountCreditOutput, error)
}
