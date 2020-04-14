package account

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// Create a function to process creating a new account
func (accountCommand *CommandHandler) Create(cmd *cobra.Command, _ []string) {
	setAsDefault, err := cmd.Flags().GetBool("default")
	if err != nil {
		console.ExitWithError(console.ExitInvalidInput, err)
	}

	passphrase := accountCommand.Prompt.RequestPassphrase()

	output, err := accountCommand.Logic.CreateAccount(&entity.CreateAccountInput{
		Passphrase:          passphrase,
		SetAsDefaultAccount: setAsDefault,
	})
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Printf("A new account is created with address %s. Please remember to keep your passphrase safe. You will not be able to recover this passphrase.", output.PublicAddress)
}
