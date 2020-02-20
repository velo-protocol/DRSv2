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

func (lo *logic) RebalanceCollateral(input *entity.RebalanceCollateralInput) ([]*entity.RebalanceCollateralOutput, error) {
	// 1. Get default account
	defaultAccount := lo.AppConfig.GetDefaultAccount()

	// 2. Prepare private key
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
		return nil, errors.Wrapf(err, "failed to decrypt the private key of %s with given passphrase", defaultAccount)
	}

	// 3. Init Velo client
	veloClient, err := lo.VFactory.NewClientFromConfig(string(privateKeyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "fail to initiate velo client")
	}

	// 4. Rebalance
	ctx, cancel := context.WithTimeout(context.Background(), constants.Timeout)
	defer cancel()

	output, err := veloClient.Rebalance(ctx, &vclient.RebalanceInput{})
	if err != nil {
		return nil, err
	}

	// 5. Map result
	outputs := make([]*entity.RebalanceCollateralOutput, len(output.Events))
	for i := range output.Events {
		outputs[i] = &entity.RebalanceCollateralOutput{
			TxHash:         output.Txs[i].Hash().String(),
			RebalanceEvent: output.Events[i],
		}
	}

	return outputs, nil
}
