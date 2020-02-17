package account

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

func (accountCommand *CommandHandler) Default(_ *cobra.Command, _ []string) {
	accounts, err := accountCommand.Logic.ListAccount()
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	if len(accounts) == 0 {
		console.ExitWithError(console.ExitError, errors.New("no account exists on gvel database"))
	}

	accountList := make([]string, len(accounts))
	for key, value := range accounts {
		accountList[key] = value.PublicAddress
	}

	choiceIndex := accountCommand.Prompt.RequestChoice(
		"Please select the account you want to make default",
		accountList,
		accountCommand.AppConfig.GetDefaultAccount(),
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
