package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// Client struct
type Client struct {
	// An rpc url
	RpcUrl string
	conn   *ethclient.Client
}

func NewClient(rpcUrl string, privateKey string) Client {
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	return Client{
		RpcUrl: rpcUrl,
		conn:   conn,
	}
}
