package credit

import (
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

func (creditCommand *CommandHandler) MintByCollateral(_ *cobra.Command, _ []string) {
	mintCreditByCollateralInput := &entity.MintCreditByCollateralInput{
		AssetCode:        creditCommand.Prompt.RequestString("Please enter asset code of credit to be minted", nil),
		CollateralAmount: creditCommand.Prompt.RequestString("Please input amount of collateral", nil),
		Passphrase:       creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Minting stable credit")
	output, err := creditCommand.Logic.MintCreditByCollateral(mintCreditByCollateralInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("%s %s minted successfully.", output.MintAmount, output.AssetCode)
	console.Logger.Infof("🔗 Transaction Hash: %s", output.TxHash)
}
