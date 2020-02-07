package environment

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
)

func (envCommand *CommandHandler) Set(_ *cobra.Command, _ []string) {
	envList := envCommand.AppConfig.GetEnvList()
	if len(envList) == 0 {
		console.ExitWithError(console.ExitError, errors.New("env list is empty"))
	}

	choiceIndex := envCommand.Prompt.RequestChoice(
		"Pick the env you want to use",
		envList,
		envCommand.AppConfig.GetCurrentEnv(),
	)

	switch envList[choiceIndex] {
	case constants.EnvTestNet:
		err := envCommand.AppConfig.SetCurrentEnv(envList[choiceIndex])
		if err != nil {
			console.ExitWithError(console.ExitError, err)
		}

		console.Logger.Printf("Switched to TESTNET. This is for testing. The tokens holds no value and should not be used to do real transaction.\n")
	case constants.EnvMainNet:
		confirm := envCommand.Prompt.RequestConfirmation("Are you sure you want to change to MAINNET? All command will be executed on MAINNET chain")
		if !confirm {
			return
		}

		err := envCommand.AppConfig.SetCurrentEnv(envList[choiceIndex])
		if err != nil {
			console.ExitWithError(console.ExitError, err)
		}

		console.Logger.Printf("Switched to MAINNET. Warning: All commands will be executed on MAINNET chain. Please take care when executing transaction.\n")
	default:
		err := envCommand.AppConfig.SetCurrentEnv(envList[choiceIndex])
		if err != nil {
			console.ExitWithError(console.ExitError, err)
		}

		console.Logger.Printf("Using env %s\n", envList[choiceIndex])
	}
}
