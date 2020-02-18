package main

import (
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/commands"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/logic"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/repositories/database"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/repositories/vfactory"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/config"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

func main() {
	// Prepare logger
	console.InitLogger()

	// Load config
	appConfig, err := config.NewConfiguration()
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	var logicInstance logic.Logic
	{
		if appConfig.Exists() {
			// db
			accountDbRepository, err := database.NewLevelDb(appConfig.GetAccountDbPath())
			if err != nil {
				console.ExitWithError(console.ExitError, err)
			}

			// logic
			logicInstance = logic.NewLogic(accountDbRepository, appConfig, vfactory.NewVeloFactory(appConfig))
		} else {
			logicInstance = logic.NewLogic(nil, appConfig, nil)
		}
	}

	commandHandler := commands.NewGvelHandler(logicInstance, appConfig)
	commandHandler.Init()

	err = commandHandler.RootCommand.Execute()
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}
}
