package credit

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// MintByCredit a function to process minting credit by credit
func (creditCommand *CommandHandler) MintByCredit(_ *cobra.Command, _ []string) {
	mintByCreditInput := &entity.MintCreditByCreditInput{
		AssetCode:    creditCommand.Prompt.RequestString("Please input asset code of credit to be minted", nil),
		CreditAmount: creditCommand.Prompt.RequestString("Please input amount of credit", nil),
		Passphrase:   creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Minting stable credit")
	output, err := creditCommand.Logic.MintCreditByCredit(mintByCreditInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("%s %s minted successfully using %s %s.", output.StableCreditAmount, output.AssetCode, output.CollateralAmount, output.CollateralAssetCode)
	console.Logger.Infof("🔗 Transaction Hash: %s", output.TxHash)
}
