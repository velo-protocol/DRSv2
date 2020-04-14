package account

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// Export a function to process exporting an account
func (accountCommand *CommandHandler) Export(_ *cobra.Command, _ []string) {
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
		"Please select the account you want to export",
		accountList,
		console.RequestChoiceOptions{
			ActiveChoice:      accountCommand.AppConfig.GetDefaultAccount(),
			ActiveChoiceLabel: constants.CursorDefault,
		},
	)
	selectedAccountExport := accountList[choiceIndex]

	passphrase := accountCommand.Prompt.RequestHiddenString("🔑 Please input the passphrase of the account", nil)

	output, err := accountCommand.Logic.ExportAccount(&entity.ExportAccountInput{
		PublicAddress: selectedAccountExport,
		Passphrase:    passphrase,
	})
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Printf("Your public key is: %s", output.PublicAddress)
	console.Logger.Printf("Your private key is: %s", output.PrivateKey)
}
