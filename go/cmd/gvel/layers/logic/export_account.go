package logic

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
	"strings"
)

func (lo *logic) ExportAccount(input *entity.ExportAccountInput) (*entity.ExportAccountOutput, error) {
	accountBytes, err := lo.DB.Get([]byte(input.PublicAddress))
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return nil, errors.Wrapf(err, "default account not found in config file, please run `gvel account create` or `gvel account import`")
	}

	var account entity.Account
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal account")
	}

	privKeyBytes, err := crypto.Decrypt(account.EncryptedPrivateKey, input.Passphrase)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt the seed of %s with given passphrase", input.PublicAddress)
	}

	return &entity.ExportAccountOutput{
		PublicAddress: account.PublicAddress,
		PrivateKey:    string(privKeyBytes),
	}, nil
}
