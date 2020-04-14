package initialize

import (
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

// NewCommandHandler creates a initialize command handler instance.
func NewCommandHandler(logic logic.Logic, prompt console.Prompt, config config.Configuration) *CommandHandler {
	return &CommandHandler{
		Logic:     logic,
		Prompt:    prompt,
		AppConfig: config,
	}
}

// Command function generates all initialize commands
func (initCommand *CommandHandler) Command() *cobra.Command {
	return &cobra.Command{
		Use:   constants.CmdInit,
		Short: "Use init command for initializing all configurations",
		Run:   initCommand.Init,
	}
}
