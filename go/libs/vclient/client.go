/*
Package vclient provides client access to a Velo smart contract.
For more information and further examples, see https://docs.velo.org/.
*/
package vclient

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
)

// Client struct contains data for connecting to the Velo Smart Contract
type Client struct {
	// The url which is used to connect to Velo smart contract
	rpcUrl string
	// A private key that will be used to sign every transaction when submitting to smart contract
	privateKey ecdsa.PrivateKey
	// A public key of private key
	publicKey common.Address
	// A connection that is used by Velo.
	conn Connection
	// The smart contracts that are used by Velo
	contract *Contract
	// The transaction helper that are used by Velo
	txHelper TxHelper
}

type ContractAddress struct {
	DrsAddress   string
	HeartAddress string
}

// NewClient creates a default client instance.
func NewClient(rpcUrl string, privateKey string, contractAddress ContractAddress) (*Client, error) {
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, errors.Wrap(err, "fail to initialize eth client")
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "invalid private key format")
	}

	pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	err = validateContractAddress(contractAddress)
	if err != nil {
		return nil, err
	}

	drsContract, err := vabi.NewDigitalReserveSystem(common.HexToAddress(contractAddress.DrsAddress), conn)
	if err != nil {
		return nil, err
	}

	heartContract, err := vabi.NewHeart(common.HexToAddress(contractAddress.HeartAddress), conn)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpcUrl:     rpcUrl,
		privateKey: *privKey,
		publicKey:  crypto.PubkeyToAddress(*pubKeyECDSA),
		conn:       conn,
		contract:   NewContract(drsContract, heartContract),
		txHelper:   NewTxHelper(conn),
	}, nil
}

// ClientOptions struct contains data for connecting to the Velo Smart Contract
type ClientOptions struct {
	// A private key that will be used to sign every transaction when submitting to smart contract
	PrivateKey ecdsa.PrivateKey
	// A connection that is used by Velo.
	Conn Connection
	// The DRS contract interface
	DRSContract DRSContract
	// The heart contract interface
	HeartContract HeartContract
	// The transaction helper that are used by Velo
	TxHelper TxHelper
}

// NewClientWithOptions creates a client with config option instance.
func NewClientWithOptions(options *ClientOptions) *Client {
	return &Client{
		privateKey: options.PrivateKey,
		conn:       options.Conn,
		contract:   NewContract(options.DRSContract, options.HeartContract),
		txHelper:   options.TxHelper,
	}
}

func validateContractAddress(contractAddress ContractAddress) error {
	if !common.IsHexAddress(contractAddress.DrsAddress) {
		return errors.New("invalid drsAddress address format")
	}
	if !common.IsHexAddress(contractAddress.HeartAddress) {
		return errors.New("invalid heart address format")
	}

	return nil
}

// Contract return contracts of client
func (c *Client) Contract() *Contract {
	return c.contract
}

// Conn return connection of client
func (c *Client) Conn() Connection {
	return c.conn
}
