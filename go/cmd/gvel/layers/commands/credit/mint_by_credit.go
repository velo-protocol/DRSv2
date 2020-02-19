package credit

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

func (creditCommand *CommandHandler) MintByCredit(_ *cobra.Command, _ []string) {
	mintByCreditInput := &entity.MintByCreditInput{
		AssetCode:    creditCommand.Prompt.RequestString("Please input asset code of credit to be minted", nil),
		CreditAmount: creditCommand.Prompt.RequestString("Please input amount of credit", nil),
		Passphrase:   creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Minting stable credit")
	output, err := creditCommand.Logic.MintByCredit(mintByCreditInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("%s %s minted successfully using %s %s.", output.MintAmount, output.AssetCode, output.CollateralAmount, output.CollateralAssetCode)
	console.Logger.Infof("🔗 Transaction Hash: %s", output.TxHash)
}
