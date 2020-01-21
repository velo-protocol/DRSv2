package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func StringToAmount(number string) (*big.Int, error) {
	d, err := decimal.NewFromString(number)
	if err != nil {
		return nil, err
	}

	d = d.Shift(7)
	b := new(big.Int)
	b, _ = b.SetString(d.String(), 10)

	return b, nil
}

func StringToByte32(s string) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], s)
	return byteArr
}
