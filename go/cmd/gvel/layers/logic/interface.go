package logic

import "github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"

type Logic interface {
	// Initialize module
	Init(configFilePath string) error

	// Account module
	CreateAccount(input *entity.CreateAccountInput) (*entity.CreateAccountOutput, error)
}
