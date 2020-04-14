package credit

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// Redeem a function to process redeeming credit
func (creditCommand *CommandHandler) Redeem(_ *cobra.Command, _ []string) {
	redeemStableCreditInput := &entity.RedeemCreditInput{
		RedeemAmount: creditCommand.Prompt.RequestString("Please input amount of stable credit", nil),
		AssetCode:    creditCommand.Prompt.RequestString("Please input asset code of credit to be redeemed", nil),
		Passphrase:   creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Redeeming stable credit")
	output, err := creditCommand.Logic.RedeemCredit(redeemStableCreditInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("Redeemed successfully. You got %s %s.", output.CollateralAmount, output.CollateralAssetCode)
	console.Logger.Infof("🔗 Transaction Hash: %s", output.TxHash)
}
