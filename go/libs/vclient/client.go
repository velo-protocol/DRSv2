package vclient

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
)

// Client struct
type Client struct {
	rpcUrl     string
	privateKey ecdsa.PrivateKey
	conn       Connection
	contract   *Contract
	txHelper   TxHelper

	// Contracts
	drs      *vabi.DigitalReserveSystem
	drsAbi   *abi.ABI
	heart    *vabi.Heart
	heartAbi *abi.ABI
}

type ContractAddress struct {
	drsAddress   string
	heartAddress string
}

func NewClient(rpcUrl string, privateKey string, contractAddress ContractAddress) (*Client, error) {
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, errors.Wrap(err, "fail to initialize eth client")
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "invalid private key format")
	}

	err = validateContractAddress(contractAddress)
	if err != nil {
		return nil, err
	}

	drsContract, err := vabi.NewDigitalReserveSystem(common.HexToAddress(contractAddress.drsAddress), conn)
	if err != nil {
		return nil, err
	}

	heartContract, err := vabi.NewHeart(common.HexToAddress(contractAddress.heartAddress), conn)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpcUrl:     rpcUrl,
		privateKey: *privKey,
		conn:       conn,
		contract:   NewContract(drsContract, heartContract),
		txHelper:   NewTxHelper(conn),
	}, nil
}

type ClientOptions struct {
	PrivateKey    ecdsa.PrivateKey
	Conn          Connection
	DRSContract   DRSContract
	HeartContract HeartContract
	TxHelper      TxHelper
}

func NewClientWithOptions(options *ClientOptions) *Client {
	return &Client{
		privateKey: options.PrivateKey,
		conn:       options.Conn,
		contract:   NewContract(options.DRSContract, options.HeartContract),
		txHelper:   options.TxHelper,
	}
}

func validateContractAddress(contractAddress ContractAddress) error {
	if !common.IsHexAddress(contractAddress.drsAddress) {
		return errors.New("invalid drsAddress address format")
	}
	if !common.IsHexAddress(contractAddress.heartAddress) {
		return errors.New("invalid heart address format")
	}

	return nil
}
