package usecase

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/strkey"
	"github.com/stellar/go/txnbuild"
	"github.com/stellar/go/xdr"
	"gitlab.com/velo-lab/core/gate/deposit"
	"gitlab.com/velo-lab/core/gate/entities"
	"gitlab.com/velo-lab/core/gate/env"
	"gitlab.com/velo-lab/core/gate/smart_contract"
	"gitlab.com/velo-lab/core/gate/stellar"
	"gitlab.com/velo-lab/core/util"
	"math/big"
	"strconv"
	"time"
)

type depositUsecase struct {
	StellarRepository       stellar.Repository
	SmartContractRepository smart_contract.Repository
}

func NewDepositUsecase(stellarRepository stellar.Repository, smartContractRepository smart_contract.Repository) deposit.Usecase {
	return &depositUsecase{
		StellarRepository:       stellarRepository,
		SmartContractRepository: smartContractRepository,
	}
}

func (d *depositUsecase) Create(c context.Context, de entities.Deposit) (*entities.DepositInfo, error) {
	gateKP, err := util.KpFromSeedString(env.GateStellarSecretKey)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get KP from string")
	}

	depositKP, err := d.createDepositAccount(gateKP)
	if err != nil {
		return nil, err
	}

	depositAccount, err := d.StellarRepository.LoadAccount(depositKP.Address())
	if err != nil {
		return nil, err
	}

	depositSeq, err := depositAccount.GetSequenceNumber()
	if err != nil {
		return nil, err
	}
	preAuthSeq := strconv.Itoa(int(depositSeq + 3))

	preAuthRefundTxB64, err := d.createPreAuthRefund(gateKP, depositKP, de.StellarAddress, preAuthSeq, de.Amount)
	if err != nil {
		return nil, err
	}

	preAuthDeleteTxB64, err := d.createPreAuthDelete(gateKP, depositKP, preAuthSeq, de.Amount)
	if err != nil {
		return nil, err
	}

	err = d.makeDepositAccount(de, depositKP)
	if err != nil {
		return nil, err
	}

	depositInfo := &entities.DepositInfo{
		DepositAddress: depositKP.Address(),
		RefundTxB64:    *preAuthRefundTxB64,
		DeleteTxB64:    *preAuthDeleteTxB64,
	}

	return depositInfo, nil
}

func (d *depositUsecase) Commit(ctx context.Context, signedTx string, smartContractChainAddress string, amount *big.Int) (*entities.Commit, error) {
	var txe xdr.TransactionEnvelope

	err := xdr.SafeUnmarshalBase64(signedTx, &txe)
	if err != nil {
		return nil, errors.Wrap(err, "fail to unmarshal signed tx")
	}

	if len(txe.Tx.Operations) != 1 {
		return nil, errors.Wrap(err, "only 1 operation will be accepted")
	}

	if txe.Tx.Operations[0].Body.PaymentOp.Destination.Address() != env.LockUpStellarAddress {
		return nil, errors.Wrap(err, "the destination is not a lockup account")
	}

	txSuccess, err := d.StellarRepository.SubmitTransaction(signedTx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to submit signed tx")
	}

	ethHash, err := d.SmartContractRepository.Mint(smartContractChainAddress, amount)
	if err != nil {
		return nil, errors.Wrap(err, "fail to mint on smart contract chain")
	}

	return &entities.Commit{
		StellarTxHash:       txSuccess.Hash,
		SmartContractTxHash: ethHash,
	}, nil
}

func (d *depositUsecase) createDepositAccount(gateKP *keypair.Full) (*keypair.Full, error) {
	depositKP, err := keypair.Random()
	if err != nil {
		return nil, errors.Wrap(err, "fail to generate a deposit account")
	}

	createAccountOp := txnbuild.CreateAccount{
		SourceAccount: &horizon.Account{
			AccountID: env.GateStellarAddress,
		},
		Destination: depositKP.Address(),
		Amount:      "5",
	}

	validatorAccount, err := d.StellarRepository.LoadAccount(gateKP.Address())
	if err != nil {
		return nil, errors.Wrap(err, "fail to load a validator account")
	}

	createAccountTx := txnbuild.Transaction{
		SourceAccount: validatorAccount,
		Operations:    []txnbuild.Operation{&createAccountOp},
		Timebounds:    txnbuild.NewTimeout(300),
		Network:       env.NetworkPassphrase,
	}

	createAccountTxB64, err := createAccountTx.BuildSignEncode(gateKP)
	if err != nil {
		return nil, errors.Wrap(err, "fail to build a transaction")
	}

	_, err = d.StellarRepository.SubmitTransaction(createAccountTxB64)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create deposit account")
	}

	return depositKP, nil
}

func (d *depositUsecase) createPreAuthRefund(gateKP *keypair.Full, depositKP *keypair.Full, uaddress string,
	seqNum string, amount string) (*string, error) {

	var ops []txnbuild.Operation
	refundOp := txnbuild.Payment{
		Amount: amount,
		Asset: txnbuild.CreditAsset{
			Code:   "VELO",
			Issuer: env.VeloIssuerAddress,
		},
		Destination: uaddress,
	}
	ops = append(ops, &refundOp)

	mergeToValidatorOp := txnbuild.AccountMerge{
		Destination: gateKP.Address(),
	}
	ops = append(ops, &mergeToValidatorOp)

	t4 := time.Now().AddDate(0, 0, 4)
	refundTx := txnbuild.Transaction{
		SourceAccount: &horizon.Account{
			AccountID: depositKP.Address(),
			Sequence:  seqNum,
		},
		Operations: ops,
		Timebounds: txnbuild.NewTimebounds(t4.Unix(), 0),
		Network:    env.NetworkPassphrase,
	}

	preAuthRefundTx, err := d.buildPreAuth(refundTx, depositKP)
	if err != nil {
		return nil, errors.Wrap(err, "fail to build pre-auth refund tx")
	}

	return preAuthRefundTx, nil
}

func (d *depositUsecase) createPreAuthDelete(gateKP *keypair.Full, depositKP *keypair.Full,
	seqNum string, amount string) (*string, error) {
	var ops []txnbuild.Operation

	deleteTrustline := txnbuild.ChangeTrust{
		Line: txnbuild.CreditAsset{
			Code:   "VELO",
			Issuer: env.VeloIssuerAddress,
		},
		Limit: "0",
	}
	ops = append(ops, &deleteTrustline)

	mergeToValidatorOp := txnbuild.AccountMerge{
		Destination: gateKP.Address(),
	}
	ops = append(ops, &mergeToValidatorOp)

	t2 := time.Now().AddDate(0, 0, 2)
	t3 := time.Now().AddDate(0, 0, 3)
	deleteTx := txnbuild.Transaction{
		SourceAccount: &horizon.Account{
			AccountID: depositKP.Address(),
			Sequence:  seqNum,
		},
		Operations: ops,
		Timebounds: txnbuild.NewTimebounds(t2.Unix(), t3.Unix()),
		Network:    env.NetworkPassphrase,
	}

	preAuthDeleteTxB64, err := d.buildPreAuth(deleteTx, depositKP)
	if err != nil {
		return nil, errors.Wrap(err, "fail to build pre-auth delete tx")
	}

	return preAuthDeleteTxB64, nil
}

func (d *depositUsecase) buildPreAuth(tx txnbuild.Transaction, preAuthKP *keypair.Full) (*string, error) {
	err := tx.Build()
	if err != nil {
		return nil, errors.Wrap(err, "unable to build tx")
	}

	txHash, err := tx.Hash()
	if err != nil {
		return nil, errors.Wrap(err, "fail to get tx hash")
	}
	encodeTxHash, err := strkey.Encode(strkey.VersionByteHashTx, txHash[:])

	var ops []txnbuild.Operation
	preAuthOp := txnbuild.SetOptions{
		Signer: &txnbuild.Signer{
			Address: encodeTxHash,
			Weight:  1,
		},
	}
	ops = append(ops, &preAuthOp)

	preAuthAccount, err := d.StellarRepository.LoadAccount(preAuthKP.Address())
	if err != nil {
		return nil, errors.Wrap(err, "fail to get details of pre-auth account")
	}

	preAuthTx := txnbuild.Transaction{
		SourceAccount: preAuthAccount,
		Operations:    ops,
		Network:       env.NetworkPassphrase,
		Timebounds:    txnbuild.NewTimeout(300),
	}

	preAuthTxB64, err := preAuthTx.BuildSignEncode(preAuthKP)
	if err != nil {
		return nil, errors.Wrap(err, "fail to build pre-auth transaction")
	}

	_, err = d.StellarRepository.SubmitTransaction(preAuthTxB64)
	if err != nil {
		return nil, errors.Wrap(err, "fail to submit pre-auth transaction")
	}

	txB64, err := tx.Base64()
	if err != nil {
		return nil, errors.Wrap(err, "fail to build a txB64")
	}

	return pointer.ToString(txB64), nil
}

func (d *depositUsecase) makeDepositAccount(de entities.Deposit, depositKP *keypair.Full) error {
	var ops []txnbuild.Operation

	encodeHashLock, err := strkey.Encode(strkey.VersionByteHashX, []byte(de.HashLock))
	if err != nil {
		return errors.Wrap(err, "fail to encode hash lock")
	}

	addHashLockOp := txnbuild.SetOptions{
		Signer: &txnbuild.Signer{
			Address: encodeHashLock,
			Weight:  1,
		},
	}
	ops = append(ops, &addHashLockOp)

	addUserOp := txnbuild.SetOptions{
		Signer: &txnbuild.Signer{
			Address: de.StellarAddress,
			Weight:  1,
		},
	}
	ops = append(ops, &addUserOp)

	addValidatorOp := txnbuild.SetOptions{
		Signer: &txnbuild.Signer{
			Address: env.GateStellarAddress,
			Weight:  1,
		},
	}
	ops = append(ops, &addValidatorOp)

	masterWeightThreshold := txnbuild.Threshold(0)
	lowThreshold := txnbuild.Threshold(3)
	midThreshold := txnbuild.Threshold(3)
	highThreshold := txnbuild.Threshold(3)

	setThresholdOp := txnbuild.SetOptions{
		MasterWeight:    &masterWeightThreshold,
		LowThreshold:    &lowThreshold,
		MediumThreshold: &midThreshold,
		HighThreshold:   &highThreshold,
	}
	ops = append(ops, &setThresholdOp)

	trustVeloOp := txnbuild.ChangeTrust{
		Line: txnbuild.CreditAsset{
			Code:   "VELO",
			Issuer: env.VeloIssuerAddress,
		},
		Limit: txnbuild.MaxTrustlineLimit,
		SourceAccount: &horizon.Account{
			AccountID: depositKP.Address(),
		},
	}
	ops = append(ops, &trustVeloOp)

	depositAccount, err := d.StellarRepository.LoadAccount(depositKP.Address())
	if err != nil {
		return errors.Wrap(err, "fail to get deposit account's detail")
	}

	makeDepositAccountTx := txnbuild.Transaction{
		SourceAccount: depositAccount,
		Operations:    ops,
		Network:       env.NetworkPassphrase,
		Timebounds:    txnbuild.NewTimeout(300),
	}

	makeDepositAccountTxeB64, err := makeDepositAccountTx.BuildSignEncode(depositKP)
	if err != nil {
		return errors.Wrap(err, "fail to make deposit account")
	}

	_, err = d.StellarRepository.SubmitTransaction(makeDepositAccountTxeB64)
	if err != nil {
		return errors.Wrap(err, "fail to confirm makeDepositAccountTx with Stellar")
	}

	return nil
}
