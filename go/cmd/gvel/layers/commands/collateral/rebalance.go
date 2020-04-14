package collateral

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
)

// Rebalance a function to process rebalancing  collateral
func (collateralCommand *CommandHandler) Rebalance(_ *cobra.Command, _ []string) {
	rebalanceCreditInput := &entity.RebalanceCollateralInput{
		Passphrase: collateralCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Rebalancing Collateral and Reserve pool")
	outputs, err := collateralCommand.Logic.RebalanceCollateral(rebalanceCreditInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("Rebalancing completed.")
	resultMessages := make([]string, len(outputs))
	for i, output := range outputs {
		resultMessages[i] = fmt.Sprintf("%s (%s): %s", output.AssetCode, output.CollateralAssetCode, output.TxHash)
	}
	console.Logger.Infof("🔗 Transaction Hash:\n%s", strings.Join(resultMessages, "\n"))
}
