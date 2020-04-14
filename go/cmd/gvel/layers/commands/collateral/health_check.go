package collateral

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

// HealthCheck a function to process getting health check collateral
func (collateralCommand *CommandHandler) HealthCheck(_ *cobra.Command, _ []string) {
	collateralHealthCheckInput := &entity.CollateralHealthCheckInput{
		Passphrase: collateralCommand.Prompt.RequestHiddenString("🔑 Please input passphrase", nil),
	}

	console.StartLoading("Getting exchange rate")
	collateralHealthCheckOutput, err := collateralCommand.Logic.CollateralHealthCheck(collateralHealthCheckInput)
	console.StopLoading()

	if err != nil {
		console.ExitWithError(console.ExitError, err)
	}

	var data [][]string
	headers := []string{"ASSET", "COLLATERAL POOL", "REQUIRED COLLATERAL"}
	for _, collateralHealthCheck := range collateralHealthCheckOutput {
		data = append(data, []string{
			fmt.Sprintf("%s (%s)", collateralHealthCheck.CollateralAssetCode, collateralHealthCheck.CollateralAssetAddress),
			fmt.Sprintf("%s", collateralHealthCheck.CollateralPoolAmount),
			fmt.Sprintf("%s", collateralHealthCheck.RequiredCollateralAmount),
		})
	}

	console.WriteTable(headers, data)
}
