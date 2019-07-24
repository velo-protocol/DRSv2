package entities

import (
	gate "gitlab.com/velo-lab/core/proto"
)

type Deposit struct {
	HashLock       string
	StellarAddress string
	DrsAddress     string
	Amount         string
	AssetCode      string
	AssetIssuer    string
}

func (entity Deposit) Parse(grpcModel *gate.DepositRequest) *Deposit {
	var depositEntity Deposit

	depositEntity.HashLock = grpcModel.HashLock
	depositEntity.StellarAddress = grpcModel.StellarAddress
	depositEntity.DrsAddress = grpcModel.DrsAddress
	depositEntity.Amount = grpcModel.Amount
	depositEntity.AssetCode = grpcModel.AssetCode
	depositEntity.AssetIssuer = grpcModel.AssetIssuer

	return &depositEntity
}

type DepositInfo struct {
	DepositAddress string
	RefundTxB64    string
	DeleteTxB64    string
}
