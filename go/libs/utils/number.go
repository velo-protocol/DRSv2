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

func AmountToString(amount *big.Int) string {
	d := decimal.NewFromInt(amount.Int64())
	d = d.Shift(-7).Truncate(7)

	return d.StringFixed(7)
}

func StringToByte32(s string) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], s)
	return byteArr
}

func Byte32ToString(b [32]byte) string {
	return string(b[:])
}
