package env

import (
	"fmt"
	"syscall"
)

var (
	GateStellarAddress   string
	GateStellarSecretKey string
	HorizonUrl           string
	NetworkPassphrase    string
	VeloIssuerAddress    string
	Port                 string
	LockUpStellarAddress string
	WarpContractAddress  string
	GateETHPrivateKey    string
	ETHNode              string
)

// Init Initialize env variables
func Init() {
	GateStellarAddress = requireEnv("GATE_STELLAR_ADDRESS")
	GateStellarSecretKey = requireEnv("GATE_STELLAR_SECRET_KEY")
	HorizonUrl = requireEnv("HORIZON_URL")
	NetworkPassphrase = requireEnv("NETWORK_PASSPHRASE")
	VeloIssuerAddress = requireEnv("VELO_ISSUER_ADDRESS")
	Port = requireEnv("PORT")
	LockUpStellarAddress = requireEnv("LOCK_UP_STELLAR_ADDRESS")
	WarpContractAddress = requireEnv("WARP_CONTRACT_ADDRESS")
	GateETHPrivateKey = requireEnv("GATE_ETH_PRIVATE_KEY")
	ETHNode = requireEnv("ETH_NODE")
}

func requireEnv(envName string) string {
	value, found := syscall.Getenv(envName)

	if !found {
		panic(fmt.Sprintf("%s env is required", envName))
	}

	return value
}
