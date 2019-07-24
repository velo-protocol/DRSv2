package grpc

import (
	"context"
	"github.com/stellar/go/xdr"
	"gitlab.com/velo-lab/core/gate/deposit"
	"gitlab.com/velo-lab/core/gate/entities"
	"gitlab.com/velo-lab/core/gate/env"
	gate "gitlab.com/velo-lab/core/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
)

type DepositHandler struct {
	DepositUsecase deposit.Usecase
}

func NewDepositHandler(s *grpc.Server, us deposit.Usecase) {
	handler := DepositHandler{
		DepositUsecase: us,
	}

	gate.RegisterGateServer(s, &handler)
}

func (dh *DepositHandler) SubmitDepositRequest(c context.Context, r *gate.DepositRequest) (*gate.SubmitDepositResponse, error) {
	if r.HashLock == "" {
		return nil, status.Error(codes.InvalidArgument, "hashLock is missing")
	}

	if r.Amount == "" {
		return nil, status.Error(codes.InvalidArgument, "amount is required")
	}

	if r.AssetCode != "VELO" {
		return nil, status.Error(codes.InvalidArgument, "only VELO will be accepted")
	}

	if r.AssetIssuer != env.VeloIssuerAddress {
		return nil, status.Error(codes.InvalidArgument, "not VELO issuer address")
	}

	if r.StellarAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "stellar address is required")
	}

	if r.DrsAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "drs address is required")
	}

	de := new(entities.Deposit).Parse(r)

	depositInfo, err := dh.DepositUsecase.Create(c, *de)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gate.SubmitDepositResponse{
		DepositAddress: depositInfo.DepositAddress,
		RefundTxeB64:   depositInfo.RefundTxB64,
		DeleteTxeB64:   depositInfo.DeleteTxB64,
	}, nil
}

func (dh *DepositHandler) CommitDeposit(c context.Context, r *gate.CommitDepositRequest) (*gate.CommitDepositResponse, error) {
	if r.SignedTx == "" {
		return nil, status.Error(codes.InvalidArgument, "signedTx is required")
	}

	var txe xdr.TransactionEnvelope
	err := xdr.SafeUnmarshalBase64(r.SignedTx, &txe)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if len(txe.Tx.Operations) != 1 {
		return nil, status.Error(codes.InvalidArgument, "only 1 operation is accepted")
	}

	if txe.Tx.Operations[0].Body.PaymentOp.Destination.Address() != env.LockUpStellarAddress {
		return nil, status.Error(codes.InvalidArgument, "the destination account is incorrect")
	}

	amountParam := big.NewInt(int64(txe.Tx.Operations[0].Body.PaymentOp.Amount))
	committedTx, err := dh.DepositUsecase.Commit(c, r.SignedTx, r.SmartContractChainAddress, amountParam)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gate.CommitDepositResponse{
		StellarTxHash:       committedTx.StellarTxHash,
		SmartContractTxHash: committedTx.SmartContractTxHash,
	}, nil
}
