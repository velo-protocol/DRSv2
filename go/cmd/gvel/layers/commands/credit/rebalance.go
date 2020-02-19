package credit

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
)

func (creditCommand *CommandHandler) Rebalance(_ *cobra.Command, _ []string) {
	rebalanceCreditInput := &entity.RebalanceCreditInput{
		Passphrase: creditCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Rebalancing Collateral and Reserve pool")
	output, err := creditCommand.Logic.RebalanceCredit(rebalanceCreditInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	console.Logger.Infof("Rebalancing completed.")
	resultMessages := make([]string, len(output.Results))
	for i, r := range output.Results {
		resultMessages[i] = fmt.Sprintf("%s (%s): %s", r.AssetCode, r.CollateralAssetCode, r.TxHash)
	}
	console.Logger.Infof("🔗 Transaction Hash:\n%s", strings.Join(resultMessages, "\n"))
}
