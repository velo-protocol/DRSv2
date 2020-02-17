package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"os"
	"path"
	"strings"
)

type configuration struct {
	sharedConfig   *viper.Viper
	envBasedConfig *viper.Viper
}

func NewConfiguration() (*configuration, error) {
	c := configuration{}
	err := c.LoadDefault()
	return &c, err
}

func getConfig(file string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func (cfg *configuration) LoadDefault() error {
	var err error
	cfg.sharedConfig, err = getConfig(constants.FsSharedConfigFile)
	if err != nil {
		// just return, if no config init yet
		return nil
	}

	envName := cfg.GetCurrentEnv()
	if envName == "" {
		return errors.New("could not read currentEnv from shared-config.json")
	}

	cfg.envBasedConfig, err = getConfig(fmt.Sprintf(constants.FsConfigFileNameFormat, envName))
	if err != nil {
		return errors.Wrapf(err, "could not read %s-config.json", envName)
	}

	return nil
}

func (cfg *configuration) InitSharedConfig(baseDir string) error {
	if cfg.Exists() {
		console.Logger.Error("config file found")
		return nil
	}

	err := os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create a config folder")
	}

	// Set default config
	v := viper.New()
	v.SetConfigType("json")
	v.SetDefault(constants.CfgKeyInitialized, true) // a flag to check for config file existence
	v.SetDefault(constants.CfgKeyCurrentEnv, constants.DefaultCurrentEnv)
	v.SetDefault(constants.CfgKeyEnvList, constants.DefaultEnvList)

	err = v.WriteConfigAs(constants.FsSharedConfigFile)
	if err != nil {
		return errors.Wrap(err, "failed to write a config to the disk")
	}

	return nil
}

func (cfg *configuration) InitEnvBasedConfig(baseDir string, envName string) error {
	envName = strings.ToLower(envName)

	err := os.MkdirAll(path.Join(baseDir, envName, "/db/account"), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create a db folder")
	}

	configKeys, ok := constants.DefaultConfigMap[envName]
	if !ok {
		return errors.New("default env base config not found")
	}

	// Set default config
	v := viper.New()
	v.SetConfigType("json")
	v.SetDefault(constants.CfgKeyAccountDbPath, path.Join(baseDir, envName, "/db/account"))
	for key, value := range configKeys {
		v.SetDefault(key, value)
	}

	err = v.WriteConfigAs(fmt.Sprintf(constants.FsConfigFileNameFormat, envName))
	if err != nil {
		return errors.Wrap(err, "failed to write a config to the disk")
	}

	return nil
}

func (cfg *configuration) Exists() bool {
	if cfg.sharedConfig == nil {
		return false
	}
	return cfg.sharedConfig.GetBool(constants.CfgKeyInitialized)
}

func (cfg *configuration) GetAccountDbPath() string {
	if cfg.envBasedConfig == nil {
		return ""
	}
	return cfg.envBasedConfig.GetString(constants.CfgKeyAccountDbPath)
}

func (cfg *configuration) GetDefaultAccount() string {
	if cfg.envBasedConfig == nil {
		return ""
	}
	return cfg.envBasedConfig.GetString(constants.CfgKeyDefaultAccount)
}

func (cfg *configuration) SetDefaultAccount(account string) error {
	if cfg.envBasedConfig == nil {
		return errors.New("envBasedConfig is empty")
	}
	cfg.envBasedConfig.Set(constants.CfgKeyDefaultAccount, account)
	return cfg.envBasedConfig.WriteConfig()
}

func (cfg *configuration) GetCurrentEnv() string {
	if cfg.sharedConfig == nil {
		return ""
	}
	return cfg.sharedConfig.GetString(constants.CfgKeyCurrentEnv)
}

func (cfg *configuration) SetCurrentEnv(envName string) error {
	if cfg.sharedConfig == nil {
		return errors.New("sharedConfig is empty")
	}
	cfg.sharedConfig.Set(constants.CfgKeyCurrentEnv, envName)
	return cfg.sharedConfig.WriteConfig()
}

func (cfg *configuration) GetEnvList() []string {
	if cfg.sharedConfig == nil {
		return []string{}
	}
	return cfg.sharedConfig.GetStringSlice(constants.CfgKeyEnvList)
}

func (cfg *configuration) GetRpcUrl() string {
	if cfg.envBasedConfig == nil {
		return ""
	}
	return cfg.envBasedConfig.GetString(constants.CfgKeyRpcUrl)
}

func (cfg *configuration) GetDrsAddress() string {
	if cfg.envBasedConfig == nil {
		return ""
	}
	return cfg.envBasedConfig.GetString(constants.CfgKeyDrsAddress)
}

func (cfg *configuration) GetHeartAddress() string {
	if cfg.envBasedConfig == nil {
		return ""
	}
	return cfg.envBasedConfig.GetString(constants.CfgKeyHeartAddress)
}
