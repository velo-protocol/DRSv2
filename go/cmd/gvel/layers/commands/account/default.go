package account

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// Default a function to process getting a default account
func (accountCommand *CommandHandler) Default(_ *cobra.Command, _ []string) {
	accounts, err := accountCommand.Logic.ListAccount()
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	if len(accounts) == 0 {
		console.ExitWithError(console.ExitError, errors.New("account not found in config file, please run `gvel account create` or `gvel account import`"))
	}

	accountList := make([]string, len(accounts))
	for key, value := range accounts {
		accountList[key] = value.PublicAddress
	}

	choiceIndex := accountCommand.Prompt.RequestChoice(
		"Please select the account you want to make default",
		accountList,
		console.RequestChoiceOptions{
			ActiveChoice:      accountCommand.AppConfig.GetDefaultAccount(),
			ActiveChoiceLabel: constants.CursorCurrent,
		},
	)
	selectedAccountDefault := accountList[choiceIndex]

	output, err := accountCommand.Logic.SetDefaultAccount(&entity.SetDefaultAccountInput{
		Account: selectedAccountDefault,
	})
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("%s is now set as the default account for signing transaction.", output.Account)
}
