package deposit

import (
	"context"
	"gitlab.com/velo-lab/core/gate/entities"
	"math/big"
)

type Usecase interface {
	Create(ctx context.Context, deposit entities.Deposit) (*entities.DepositInfo, error)
	Commit(ctx context.Context, signedTx string, targetAddress string, amount *big.Int) (*entities.Commit, error)
}
