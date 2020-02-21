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

func (lo *logic) MintCreditByCollateral(input *entity.MintCreditByCollateralInput) (*entity.MintCreditByCollateralOutput, error) {
	defaultAccount := lo.AppConfig.GetDefaultAccount()
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

	ctx, cancel := context.WithTimeout(context.Background(), constants.Timeout)
	defer cancel()

	output, err := veloClient.MintFromCollateralAmount(ctx, &vclient.MintFromCollateralAmountInput{
		AssetCode:        input.AssetCode,
		CollateralAmount: input.CollateralAmount,
	})
	if err != nil {
		return nil, err
	}

	return &entity.MintCreditByCollateralOutput{
		AssetCode:          output.Event.AssetCode,
		StableCreditAmount: output.Event.StableCreditAmount,
		TxHash:             output.Tx.Hash().String(),
	}, nil
}
