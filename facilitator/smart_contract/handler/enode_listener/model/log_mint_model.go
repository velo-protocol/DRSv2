package model

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//event Mint(
//	bytes32 assetName,
//	uint256 amount,
//	uint256 collateral,
//	address indexed issuer,
//	string indexed stellarAddr,
//	bytes32 hashSecret
//);

type LogMint struct {
	AssetName      [32]byte
	Amount         *big.Int
	Collateral     *big.Int
	Issuer         common.Address
	StellarAddress string
	HashSecret     [32]byte
}
