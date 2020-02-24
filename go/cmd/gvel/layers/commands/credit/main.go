package credit

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

func (creditCommand *CommandHandler) Command() *cobra.Command {
	command := &cobra.Command{
		Use:   fmt.Sprintf("%s %s", constants.CmdCredit, "<arg>"),
		Short: "Use credit command for managing the stable credit on Velo",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if !creditCommand.AppConfig.Exists() {
				console.ExitWithError(console.ExitError, errors.New("config file not found, please run `gvel init`"))
			}

			if creditCommand.AppConfig.GetDefaultAccount() == "" {
				console.ExitWithError(console.ExitError, errors.New("default account not found in config file, please run `gvel account create` or `gvel account import`"))
			}
		},
	}

	command.AddCommand(
		creditCommand.GetSetupCommand(),
		creditCommand.GetMintCreditByCollateralCommand(),
		creditCommand.GetMintByCreditCommand(),
		creditCommand.GetRedeemStableCreditCommand(),
		creditCommand.GetExchangeCommand(),
	)

	return command
}

func (creditCommand *CommandHandler) GetSetupCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCreditSetup,
		Short: "Setup a stable credit on Velo",
		Run:   creditCommand.Setup,
	}

	return command
}

func (creditCommand *CommandHandler) GetMintCreditByCollateralCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCreditMintByCollateral,
		Short: "Mint a stable credit by collateral on Velo",
		Run:   creditCommand.MintByCollateral,
	}

	return command
}

func (creditCommand *CommandHandler) GetMintByCreditCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCreditMintByCredit,
		Short: "Mint stable credit on Velo by providing amount of desired stable credit",
		Run:   creditCommand.MintByCredit,
	}

	return command
}

func (creditCommand *CommandHandler) GetExchangeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCreditGetExchange,
		Short: "Get exchange rate of a stable credit on Velo",
		Run:   creditCommand.GetExchange,
	}

	return command
}

func (creditCommand *CommandHandler) GetRedeemStableCreditCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   constants.CmdCreditRedeem,
		Short: "Redeem stable credit to collateral Asset on Velo",
		Run:   creditCommand.RedeemStableCredit,
	}

	return command
}
