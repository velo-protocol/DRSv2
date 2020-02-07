package constants

const (
	EnvMainNet = "mainnet"
	EnvTestNet = "testnet"
)

const (
	DefaultCurrentEnv = EnvTestNet
)

var (
	DefaultEnvList = []string{
		EnvTestNet,
		EnvMainNet,
	}
)
