package environment

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/logic"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/config"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

type CommandHandler struct {
	Logic     logic.Logic
	Prompt    console.Prompt
	AppConfig config.Configuration
}

func NewCommandHandler(logic logic.Logic, prompt console.Prompt, config config.Configuration) *CommandHandler {
	return &CommandHandler{
		Logic:     logic,
		Prompt:    prompt,
		AppConfig: config,
	}
}

func (envCommand *CommandHandler) Command() *cobra.Command {
	command := &cobra.Command{
		Use:   fmt.Sprintf("%s %s", constants.CmdEnv, "<arg>"),
		Short: "Use env command to config environment",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if !envCommand.AppConfig.Exists() {
				console.ExitWithError(console.ExitError, errors.New("config file not found, please run `gvel init`"))
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			console.Logger.Printf("The current env is %s", envCommand.AppConfig.GetCurrentEnv())
			console.Logger.Printf("You can set the env by using command `gvel env set`")
		},
	}

	command.AddCommand(
		envCommand.GetSetCommand(),
	)
	return command
}

func (envCommand *CommandHandler) GetSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   constants.CmdEnvSet,
		Short: "Select and set env",
		Run:   envCommand.Set,
	}
}
