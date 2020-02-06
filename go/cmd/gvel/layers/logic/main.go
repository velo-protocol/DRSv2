package logic

import (
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/repositories/database"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/config"
)

type logic struct {
	DB        database.Repository
	AppConfig config.Configuration
}

func NewLogic(db database.Repository, config config.Configuration) Logic {
	return &logic{
		DB:        db,
		AppConfig: config,
	}
}
