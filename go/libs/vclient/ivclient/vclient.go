package ivclient

import (
	"context"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)

type VClient interface {
	SetupCredit(ctx context.Context, input *vclient.SetupCreditInput) (*vclient.SetupCreditOutput, error)
	MintFromCollateralAmount(ctx context.Context, input *vclient.MintFromCollateralAmountInput) (*vclient.MintFromCollateralAmountCreditOutput, error)
	MintFromStableCreditAmount(ctx context.Context, input *vclient.MintFromStableCreditAmountInput) (*vclient.MintFromStableCreditAmountCreditOutput, error)
	RedeemCredit(ctx context.Context, input *vclient.RedeemCreditInput) (*vclient.RedeemCreditOutput, error)
	GetExchangeRate(input *vclient.GetExchangeRateInput) (*vclient.GetExchangeRateOutput, error)
	CollateralHealthCheck(input *vclient.CollateralHealthCheckInput) ([]vclient.CollateralHealthCheckOutput, error)
	Rebalance(ctx context.Context, input *vclient.RebalanceInput) (*vclient.RebalanceOutput, error)
}
