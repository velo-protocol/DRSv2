package config

type Configuration interface {
	InitSharedConfig(baseDir string) error
	InitEnvBasedConfig(baseDir string, envName string) error
	Exists() bool
	GetAccountDbPath() string
}
