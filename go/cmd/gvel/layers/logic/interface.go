package logic

import "github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"

type Logic interface {
	// Initialize module
	Init(configFilePath string) error

	// Account module
	CreateAccount(input *entity.CreateAccountInput) (*entity.CreateAccountOutput, error)
	ImportAccount(input *entity.ImportAccountInput) (*entity.ImportAccountOutput, error)
	ExportAccount(input *entity.ExportAccountInput) (*entity.ExportAccountOutput, error)
	ListAccount() ([]*entity.Account, error)
	SetDefaultAccount(input *entity.SetDefaultAccountInput) (*entity.SetDefaultAccountOutput, error)

	// Credit module
	SetupCredit(input *entity.SetupCreditInput) (*entity.SetupCreditOutput, error)

	// Environment module
	SetEnv(input *entity.SetEnvInput) error
}
