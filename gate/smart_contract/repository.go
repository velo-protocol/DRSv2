package smart_contract

import (
	"math/big"
)

type Repository interface {
	Mint(targetAccount string, amount *big.Int) (string, error)
}
