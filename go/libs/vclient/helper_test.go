package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/constants"
	"log"
	"math/big"
)

const (
	smartContractUrl = "http://127.0.0.1:7545"

	drsAddress            = "0x0000000000000000000000000000000000000001"
	heartAddress          = "0x0000000000000000000000000000000000000002"
	governorAddress       = "0x0000000000000000000000000000000000000003"
	trustedPartnerAddress = "0x50637DeE3598e080B7B605B00f4FfC721046E4E0"
	invalidAddress        = "GD4K"

	privateKey1 = "da17d295e2fd005747cca4de855bbb0493f2e0669753bba1e752700dbad4c78c"
	privateKey2 = "2e8b8b7fccb7fa535857a6edf74d72fd389fb4b88d877ed57a1fdbeaac1d6862"
)

func getOpts(hexPrivateKey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatalf("hex private key :%s is not support", hexPrivateKey)
	}
	opts := bind.NewKeyedTransactor(privateKey)
	opts.GasLimit = constants.GasLimit
	return opts
}

type TestHelper struct {
	DrsAddress     common.Address
	DrsContract    *vabi.DigitalReserveSystem
	HeartAddress   common.Address
	HeartContract  *vabi.Heart
	GenesisAccount *bind.TransactOpts
	Conn           *backends.SimulatedBackend
	Client         *Client
}

func testHelper() *TestHelper {
	opts := getOpts(privateKey1)
	opts2 := getOpts(privateKey2)
	alloc := make(core.GenesisAlloc)
	blockGasLimit := uint64(constants.GasLimit)
	address1 := opts.From
	address2 := opts2.From
	alloc = map[common.Address]core.GenesisAccount{
		address1: {
			Balance: big.NewInt(1000000000000000000),
		},
		address2: {
			Balance: big.NewInt(1000000000000000000),
		},
	}
	conn := backends.NewSimulatedBackend(alloc, blockGasLimit)

	//Deploy Heart contract
	heartAddress, _, heartContract, _ := vabi.DeployHeart(
		opts,
		conn,
	)
	conn.Commit()

	//Deploy DigitalReserveSystem contract
	drsAddress, _, drsContract, _ := vabi.DeployDigitalReserveSystem(
		opts,
		conn,
		heartAddress,
	)

	conn.Commit()

	client, _ := NewClientWithEthClient(conn, privateKey1, ContractAddress{
		DRS:   drsAddress.String(),
		Heart: heartAddress.String(),
	})

	return &TestHelper{
		DrsAddress:     drsAddress,
		DrsContract:    drsContract,
		HeartAddress:   heartAddress,
		HeartContract:  heartContract,
		GenesisAccount: opts,
		Conn:           conn,
		Client:         client,
	}
}
