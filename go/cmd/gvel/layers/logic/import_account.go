package logic

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
)

func (lo *logic) ImportAccount(input *entity.ImportAccountInput) (*entity.ImportAccountOutput, error) {
	// 1. Check private key format
	pubAddress, privKey, err := crypto.PrivateKeyToKeyPair(input.PrivateKey)
	if err != nil {
		return nil, err
	}

	// 2. Encrypt private key
	encryptedPrivKey, nonce, err := crypto.Encrypt([]byte(*privKey), input.Passphrase)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt private key")
	}

	// 3. Prepare account entity to be saved as bytes
	account := &entity.Account{
		PublicAddress:       *pubAddress,
		EncryptedPrivateKey: encryptedPrivKey,
		Nonce:               nonce,
	}
	accountBytes, err := json.Marshal(account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal entity")
	}

	// 4. Save account to db
	err = lo.DB.Save([]byte(*pubAddress), accountBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save an account")
	}

	// 5. Set default account
	mustSetDefault := lo.AppConfig.GetDefaultAccount() == "" || input.SetAsDefaultAccount
	if mustSetDefault {
		err = lo.AppConfig.SetDefaultAccount(*pubAddress)
		if err != nil {
			return nil, errors.Wrap(err, "failed to write config file")
		}
	}

	return &entity.ImportAccountOutput{
		ImportedAccountAddress: *pubAddress,
	}, nil
}
