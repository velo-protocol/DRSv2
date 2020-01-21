package client

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
)

// Client struct
type Client struct {
	rpcUrl          string
	privateKey      ecdsa.PrivateKey
	conn            *ethclient.Client
	contractAddress ContractAddress

	// Contracts
	drs   *vabi.DigitalReserveSystem
	heart *vabi.Heart
}

type ContractAddress struct {
	DRS   string
	Heart string
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

	return &Client{
		rpcUrl:          rpcUrl,
		privateKey:      *privKey,
		conn:            conn,
		contractAddress: contractAddress,
	}, nil
}

func validateContractAddress(contractAddress ContractAddress) error {
	if !common.IsHexAddress(contractAddress.DRS) {
		return errors.New("invalid DRS address format")
	}
	if !common.IsHexAddress(contractAddress.Heart) {
		return errors.New("invalid heart address format")
	}

	return nil
}

func (c *Client) DRS() *vabi.DigitalReserveSystem {
	if c.drs == nil {
		c.drs, _ = vabi.NewDigitalReserveSystem(common.HexToAddress(c.contractAddress.DRS), c.conn)
	}

	return c.drs
}

func (c *Client) Heart() *vabi.Heart {
	if c.heart == nil {
		c.heart, _ = vabi.NewHeart(common.HexToAddress(c.contractAddress.Heart), c.conn)
	}

	return c.heart
}
