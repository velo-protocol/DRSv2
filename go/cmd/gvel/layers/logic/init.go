package logic

func (lo *logic) Init(configFilePath string) error {
	err := lo.AppConfig.InitSharedConfig(configFilePath)
	if err != nil {
		return err
	}

	err = lo.AppConfig.InitEnvBasedConfig(configFilePath, "testnet")
	if err != nil {
		return err
	}

	err = lo.AppConfig.InitEnvBasedConfig(configFilePath, "mainnet")
	if err != nil {
		return err
	}

	return nil
}
