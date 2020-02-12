package logic

import "github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"

func (lo *logic) Init(configFilePath string) error {
	err := lo.AppConfig.InitSharedConfig(configFilePath)
	if err != nil {
		return err
	}

	err = lo.AppConfig.InitEnvBasedConfig(configFilePath, constants.EnvTestNet)
	if err != nil {
		return err
	}

	err = lo.AppConfig.InitEnvBasedConfig(configFilePath, constants.EnvMainNet)
	if err != nil {
		return err
	}

	return nil
}
