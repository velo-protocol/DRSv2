package logic

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
)

func (lo *logic) ListAccount() ([]*entity.Account, error) {
	accountsBytes, err := lo.DB.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account from db")
	}

	defaultAccount := lo.AppConfig.GetDefaultAccount()

	var accounts []*entity.Account
	for _, accountBytes := range accountsBytes {
		tmpAccount := &entity.Account{}

		err := json.Unmarshal(accountBytes, tmpAccount)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal account")
		}

		tmpAccount.IsDefault = defaultAccount == tmpAccount.PublicAddress
		accounts = append(accounts, tmpAccount)
	}

	return accounts, nil
}
