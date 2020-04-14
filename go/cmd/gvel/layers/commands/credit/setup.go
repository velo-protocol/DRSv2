package credit

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// Setup a function to process setup new credit
func (creditCommand *CommandHandler) Setup(_ *cobra.Command, _ []string) {
	setupCreditInput := &entity.SetupCreditInput{
		AssetCode:           creditCommand.Prompt.RequestString("Please input asset code", nil),
		CollateralAssetCode: creditCommand.Prompt.RequestString("Please input collateral asset code", nil),
		PeggedValue:         creditCommand.Prompt.RequestString("Please input pegged value", nil),
		PeggedCurrency:      creditCommand.Prompt.RequestString("Please input pegged currency", nil),
		Passphrase:          creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Setting up stable credit")
	output, err := creditCommand.Logic.SetupCredit(setupCreditInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("Stable credit %s (%s) set up for account %s successfully.", output.AssetCode, output.AssetAddress, output.CreditOwnerAddress)
	console.Logger.Infof("🔗 Transaction Hash: %s", output.TxHash)
}
