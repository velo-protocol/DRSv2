package vclient

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
	"strings"
)

// Client struct
type Client struct {
	rpcUrl          string
	privateKey      ecdsa.PrivateKey
	conn            Connection
	contractAddress ContractAddress

	// Contracts
	drs      *vabi.DigitalReserveSystem
	drsAbi   *abi.ABI
	heart    *vabi.Heart
	heartAbi *abi.ABI
}

type ContractAddress struct {
	DRS   string
	Heart string
}

type Connection interface {
	bind.ContractBackend
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
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

func NewClientWithEthClient(conn Connection, privateKey string, contractAddress ContractAddress) (*Client, error) {
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "invalid private key format")
	}

	err = validateContractAddress(contractAddress)
	if err != nil {
		return nil, err
	}

	return &Client{
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

func (c *Client) DRSAbi() *abi.ABI {
	if c.drsAbi == nil {
		contractAbi, _ := abi.JSON(strings.NewReader(vabi.DigitalReserveSystemABI))
		c.drsAbi = &contractAbi
	}

	return c.drsAbi
}

func (c *Client) Heart() *vabi.Heart {
	if c.heart == nil {
		c.heart, _ = vabi.NewHeart(common.HexToAddress(c.contractAddress.Heart), c.conn)
	}

	return c.heart
}

func (c *Client) HeartAbi() *abi.ABI {
	if c.heartAbi == nil {
		contractAbi, _ := abi.JSON(strings.NewReader(vabi.HeartABI))
		c.heartAbi = &contractAbi
	}

	return c.heartAbi
}

func (c *Client) ConfirmTx(tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), c.conn, tx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to confirm transaction")
	}
	return receipt, nil
}
