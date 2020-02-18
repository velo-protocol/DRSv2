package logic

import (
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/repositories/database"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/repositories/vfactory"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/config"
)

type logic struct {
	DB        database.Repository
	AppConfig config.Configuration
	VFactory  vfactory.Repository
}

func NewLogic(db database.Repository, config config.Configuration, vFactory vfactory.Repository) Logic {
	return &logic{
		DB:        db,
		AppConfig: config,
		VFactory:  vFactory,
	}
}
