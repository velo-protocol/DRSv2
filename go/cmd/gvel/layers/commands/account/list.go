package account

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// List a function to process listing all accounts
func (accountCommand *CommandHandler) List(cmd *cobra.Command, _ []string) {
	accounts, err := accountCommand.Logic.ListAccount()
	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	var data [][]string
	headers := []string{"Index", "Address", "Default"}
	for index, account := range accounts {
		data = append(data, []string{
			fmt.Sprintf("%d", index+1),
			fmt.Sprintf("%s", account.PublicAddress),
			fmt.Sprintf("%v", account.IsDefault),
		})
	}

	console.WriteTable(headers, data)
}
