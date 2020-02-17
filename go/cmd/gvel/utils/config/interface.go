package config

type Configuration interface {
	InitSharedConfig(baseDir string) error
	InitEnvBasedConfig(baseDir string, envName string) error
	Exists() bool
	GetAccountDbPath() string
	GetDefaultAccount() string
	SetDefaultAccount(account string) error
	GetCurrentEnv() string
	SetCurrentEnv(account string) error
	GetEnvList() []string
	GetRpcUrl() string
	GetDrsAddress() string
	GetHeartAddress() string
}
