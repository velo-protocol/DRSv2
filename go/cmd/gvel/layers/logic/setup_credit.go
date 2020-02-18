package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)

func (lo *logic) SetupCredit(input *entity.SetupCreditInput) (*entity.SetupCreditOutput, error) {
	// 1. Get default account
	defaultAccount := lo.AppConfig.GetDefaultAccount()

	// 2. Prepare private key
	accountBytes, err := lo.DB.Get([]byte(defaultAccount))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account from db")
	}

	var account entity.Account
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal account")
	}

	privateKeyBytes, err := crypto.Decrypt(account.EncryptedPrivateKey, input.Passphrase)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt the private key of %s with given passphrase", defaultAccount)
	}

	// 3. Init Velo client
	veloClient, err := lo.VFactory.NewClient(&entity.NewClientInput{
		RpcUrl:     lo.AppConfig.GetRpcUrl(),
		PrivateKey: string(privateKeyBytes),
		ContractAddress: &vclient.ContractAddress{
			DrsAddress:   lo.AppConfig.GetDrsAddress(),
			HeartAddress: lo.AppConfig.GetHeartAddress(),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "fail to initiate velo client")
	}

	// 4. Setup credit
	ctx, cancel := context.WithTimeout(context.Background(), constants.Timeout)
	defer cancel()

	output, err := veloClient.SetupCredit(ctx, &vclient.SetupCreditInput{
		AssetCode:           input.AssetCode,
		PeggedValue:         input.PeggedValue,
		PeggedCurrency:      input.PeggedCurrency,
		CollateralAssetCode: input.CollateralAssetCode,
	})
	if err != nil {
		return nil, err
	}

	return &entity.SetupCreditOutput{
		TxHash:              output.Tx.Hash().String(),
		CreditOwnerAddress:  defaultAccount,
		AssetCode:           output.Event.AssetCode,
		AssetAddress:        output.Event.AssetAddress,
		PeggedValue:         output.Event.PeggedValue,
		PeggedCurrency:      output.Event.PeggedCurrency,
		CollateralAssetCode: output.Event.CollateralAssetCode,
	}, nil
}
