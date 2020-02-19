package constants

import "strings"

const (
	EnvMainNet = "mainnet"
	EnvTestNet = "testnet"

	TestNetRpcUrl       = "http://127.0.0.1:7545"
	TestNetDrsAddress   = "0x08d6b93451d00bE089e1c4f2A74D1485F559B6F1"
	TestNetHeartAddress = "0x833Aa557EDd8520055624E3Adc2315771f98248d"

	MainNetRpcUrl       = "http://10.14.129.254:8545"
	MainNetDrsAddress   = "0x3B39BaBa62bafF6367AE6BD3352FAD7C34C5938E"
	MainNetHeartAddress = "0xa910805697c4218a6b42ff3c11Be01990209f7Cf"
)

const (
	DefaultCurrentEnv = EnvTestNet
)

var (
	DefaultEnvList = []string{
		strings.ToUpper(EnvTestNet),
		strings.ToUpper(EnvMainNet),
	}

	DefaultConfigMap = map[string]map[string]string{
		EnvMainNet: {
			CfgKeyRpcUrl:       MainNetRpcUrl,
			CfgKeyDrsAddress:   MainNetDrsAddress,
			CfgKeyHeartAddress: MainNetHeartAddress,
		},
		EnvTestNet: {
			CfgKeyRpcUrl:       TestNetRpcUrl,
			CfgKeyDrsAddress:   TestNetDrsAddress,
			CfgKeyHeartAddress: TestNetHeartAddress,
		},
	}
)
