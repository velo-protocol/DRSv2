package collateral

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

// NewCommandHandler creates a collateral command handler instance.
func NewCommandHandler(logic logic.Logic, prompt console.Prompt, config config.Configuration) *CommandHandler {
	return &CommandHandler{
		Logic:     logic,
		Prompt:    prompt,
		AppConfig: config,
	}
}

// Command function generates all collateral commands
func (collateralCommand *CommandHandler) Command() *cobra.Command {
	command := &cobra.Command{
		Use:   fmt.Sprintf("%s %s", constants.CmdCollateral, "<arg>"),
		Short: "Use collateral command for managing the collateral asset on Velo",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if !collateralCommand.AppConfig.Exists() {
				console.ExitWithError(console.ExitError, errors.New("config file not found, please run `gvel init`"))
			}

			if collateralCommand.AppConfig.GetDefaultAccount() == "" {
				console.ExitWithError(console.ExitError, errors.New("default account not found in config file, please run `gvel account create` or `gvel account import`"))
			}
		},
	}

	command.AddCommand(
		collateralCommand.GetHealthCheckCommand(),
		collateralCommand.GetRebalanceCommand(),
	)

	return command
}

// GetHealthCheckCommand function return get health check command
func (collateralCommand *CommandHandler) GetHealthCheckCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCollateralHealthCheck,
		Short: "Get collateral health check",
		Run:   collateralCommand.HealthCheck,
	}
	return command
}

// GetRebalanceCommand function return get rebalance command
func (collateralCommand *CommandHandler) GetRebalanceCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCollateralRebalance,
		Short: "Rebalance collateral",
		Run:   collateralCommand.Rebalance,
	}
	return command
}
