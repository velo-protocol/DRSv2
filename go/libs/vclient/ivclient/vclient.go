package ivclient

import (
	"context"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)

type VClient interface {
	SetupCredit(ctx context.Context, input *vclient.SetupCreditInput) (*vclient.SetupCreditOutput, error)
	MintFromCollateralAmount(ctx context.Context, input *vclient.MintByCollateralAmountInput) (*vclient.MintByCollateralAmountCreditOutput, error)
	GetExchangeRate(input *vclient.GetExchangeRateInput) (*vclient.GetExchangeRateOutput, error)
	CollateralHealthCheck(input *vclient.CollateralHealthCheckInput) ([]vclient.CollateralHealthCheckOutput, error)
}
