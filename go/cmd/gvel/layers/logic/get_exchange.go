package logic

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)

func (lo *logic) GetCreditExchange(input *entity.GetCreditExchangeInput) (*entity.GetCreditExchangeOutput, error) {
	defaultAccount := lo.AppConfig.GetDefaultAccount()
	accountBytes, err := lo.DB.Get([]byte(defaultAccount))
	if err != nil {
		return nil, errors.Wrap(err, "default account is not found")
	}

	var account entity.Account
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal account")
	}

	privateKeyBytes, err := crypto.Decrypt(account.EncryptedPrivateKey, input.Passphrase)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt the private key of %s with given passphrase", account.PublicAddress)
	}

	veloClient, err := lo.VFactory.NewClient(&entity.NewClientInput{
		RpcUrl:     lo.AppConfig.GetRpcUrl(),
		PrivateKey: string(privateKeyBytes),
		ContractAddress: &vclient.ContractAddress{
			DrsAddress:   lo.AppConfig.GetDrsAddress(),
			HeartAddress: lo.AppConfig.GetHeartAddress(),
		},
	})
	if err != nil {
		return nil, err
	}

	output, err := veloClient.GetExchangeRate(&vclient.GetExchangeRateInput{
		AssetCode: input.AssetCode,
	})
	if err != nil {
		return nil, err
	}

	return &entity.GetCreditExchangeOutput{
		AssetCode:                     output.AssetCode,
		CollateralAssetCode:           output.CollateralAssetCode,
		PriceInCollateralPerAssetUnit: output.PriceInCollateralPerAssetUnit,
	}, nil
}
