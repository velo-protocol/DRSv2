package constants

import "strings"

const (
	EnvMainNet = "mainnet"
	EnvTestNet = "testnet"
)

const (
	DefaultCurrentEnv = EnvTestNet
)

var (
	DefaultEnvList = []string{
		strings.ToUpper(EnvTestNet),
		strings.ToUpper(EnvMainNet),
	}
)
