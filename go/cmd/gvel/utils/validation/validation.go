package validation

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func ValidatePrivateKeyFormat(input string) error {
	_, err := crypto.HexToECDSA(input)
	if err != nil {
		return errors.New("invalid private key format")
	}

	return nil
}
