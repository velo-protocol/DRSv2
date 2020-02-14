package constants

import "strings"

const (
	EnvMainNet = "mainnet"
	EnvTestNet = "testnet"

	TestNetRpcUrl       = "http://127.0.0.1:7545"
	TestNetDrsAddress   = "0x3931B8d4C37b2E15555A49cAfAF807ef20CA77D2"
	TestNetHeartAddress = "0x68C3497D9066C9e237b7140A4fC723658ceA2E47"

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
