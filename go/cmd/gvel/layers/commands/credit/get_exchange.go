package credit

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// GetExchange a function to process getting credit exchange
func (creditCommand *CommandHandler) GetExchange(_ *cobra.Command, _ []string) {
	getExchangeInput := &entity.GetCreditExchangeInput{
		AssetCode:  creditCommand.Prompt.RequestString("Please input asset code of the stable credit", nil),
		Passphrase: creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Getting exchange rate")
	output, err := creditCommand.Logic.GetCreditExchange(getExchangeInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("You will get %s %s for 1 %s.", output.PriceInCollateralPerAssetUnit, output.CollateralAssetCode, output.AssetCode)

}
