package env

import (
	"fmt"
	"syscall"
)

var (
	FacilitatorStellarAddress   string
	FacilitatorStellarSecretKey string
	HorizonUrl                  string
	NetworkPassphrase           string
	VeloIssuerAddress           string
	Port                        string
	DRSSCAddress                string
	GateETHPrivateKey           string
	ETHNode                     string
)

// Init Initialize env variables
func Init() {
	FacilitatorStellarAddress = requireEnv("GATE_STELLAR_ADDRESS")
	FacilitatorStellarSecretKey = requireEnv("GATE_STELLAR_SECRET_KEY")
	HorizonUrl = requireEnv("HORIZON_URL")
	NetworkPassphrase = requireEnv("NETWORK_PASSPHRASE")
	VeloIssuerAddress = requireEnv("VELO_ISSUER_ADDRESS")
	Port = requireEnv("PORT")
	DRSSCAddress = requireEnv("DRSSC_ADDRESS")
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
