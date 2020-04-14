package account

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/validation"
)

// Import a function to process importing an existing account
func (accountCommand *CommandHandler) Import(cmd *cobra.Command, _ []string) {
	setAsDefault, err := cmd.Flags().GetBool("default")
	if err != nil {
		console.ExitWithError(console.ExitInvalidInput, err)
	}

	privateKey := accountCommand.Prompt.RequestHiddenString("Please enter private key of the address", validation.ValidatePrivateKeyFormat)
	passphrase := accountCommand.Prompt.RequestPassphrase()

	console.StartLoading("Importing account")
	output, err := accountCommand.Logic.ImportAccount(&entity.ImportAccountInput{
		Passphrase:          passphrase,
		PrivateKey:          privateKey,
		SetAsDefaultAccount: setAsDefault,
	})
	console.StopLoading()
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	if setAsDefault {
		console.Logger.Printf("Account with address %s is now the default account. Please keep your passphrase safe.", output.ImportedAccountAddress)
	} else {
		console.Logger.Printf("Add account with address %s to gvel. Please keep your passphrase safe.", output.ImportedAccountAddress)
	}
}
