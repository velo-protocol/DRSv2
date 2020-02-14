package logic

import "github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"

type Logic interface {
	// Initialize module
	Init(configFilePath string) error

	// Account module
	CreateAccount(input *entity.CreateAccountInput) (*entity.CreateAccountOutput, error)
	ImportAccount(input *entity.ImportAccountInput) (*entity.ImportAccountOutput, error)
	ListAccount() ([]*entity.Account, error)
	SetDefaultAccount(input *entity.SetDefaultAccountInput) (*entity.SetDefaultAccountOutput, error)

	// Environment module
	SetEnv(input *entity.SetEnvInput) error
}
