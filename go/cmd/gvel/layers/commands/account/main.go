package account

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

// NewCommandHandler creates an account command handler instance.
func NewCommandHandler(logic logic.Logic, prompt console.Prompt, config config.Configuration) *CommandHandler {
	return &CommandHandler{
		Logic:     logic,
		Prompt:    prompt,
		AppConfig: config,
	}
}

// Command function generates all account commands
func (accountCommand *CommandHandler) Command() *cobra.Command {
	command := &cobra.Command{
		Use:   fmt.Sprintf("%s %s", constants.CmdAccount, "<arg>"),
		Short: "Use account command for managing the account interacting with Velo",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if !accountCommand.AppConfig.Exists() {
				console.ExitWithError(console.ExitError, errors.New("config file not found, please run `gvel init`"))
			}
		},
	}

	command.AddCommand(
		accountCommand.GetCreateCommand(),
		accountCommand.GetImportCommand(),
		accountCommand.GetListCommand(),
		accountCommand.GetDefaultCommand(),
		accountCommand.GetExportCommand(),
	)
	return command
}

// GetCreateCommand function return account create command
func (accountCommand *CommandHandler) GetCreateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdAccountCreate,
		Short: "Create an account and store on your disk",
		Run:   accountCommand.Create,
	}

	command.Flags().BoolP(constants.FlagDefault, "d", false, "set as default account")
	return command
}

// GetImportCommand function return account import command
func (accountCommand *CommandHandler) GetImportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdAccountImport,
		Short: "Import your account",
		Run:   accountCommand.Import,
	}

	command.Flags().BoolP(constants.FlagDefault, "d", false, "set as default account")
	return command
}

// GetListCommand function return account list command
func (accountCommand *CommandHandler) GetListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   constants.CmdAccountList,
		Short: "Print all accounts that were created",
		Run:   accountCommand.List,
	}
}

// GetDefaultCommand function return account default command
func (accountCommand *CommandHandler) GetDefaultCommand() *cobra.Command {
	return &cobra.Command{
		Use:   constants.CmdAccountDefault,
		Short: "Set default account to be used as signer",
		Run:   accountCommand.Default,
	}
}

// GetExportCommand function return account export command
func (accountCommand *CommandHandler) GetExportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdAccountExport,
		Short: "Export your account",
		Run:   accountCommand.Export,
	}

	return command
}
