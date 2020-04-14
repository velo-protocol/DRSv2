package environment

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
)

// Set a function to process setting environment
func (envCommand *CommandHandler) Set(_ *cobra.Command, _ []string) {
	envList := envCommand.AppConfig.GetEnvList()
	if envList == nil || len(envList) == 0 {
		console.ExitWithError(console.ExitError, errors.New("env list is empty"))
		return
	}

	choiceIndex := envCommand.Prompt.RequestChoice(
		"Pick the env you want to use",
		envList,
		console.RequestChoiceOptions{
			ActiveChoice:      envCommand.AppConfig.GetCurrentEnv(),
			ActiveChoiceLabel: constants.CursorCurrent,
		},
	)

	chosenEnv := strings.ToLower(envList[choiceIndex])

	if chosenEnv == constants.EnvMainNet {
		confirm := envCommand.Prompt.RequestConfirmation("Are you sure you want to change to MAINNET? All command will be executed on MAINNET chain")
		if !confirm {
			return
		}
	}

	err := envCommand.Logic.SetEnv(&entity.SetEnvInput{Env: chosenEnv})
	if err != nil {
		console.ExitWithError(console.ExitError, err)
		return
	}

	switch chosenEnv {
	case constants.EnvTestNet:
		console.Logger.Printf("Switch to TESTNET. This is for testing. The tokens holds no value and should not be used to do real transaction.\n")
	case constants.EnvMainNet:
		console.Logger.Printf("Switch to MAINNET. Warning: All commands will be executed on MAINNET chain. Please take care when executing transaction.\n")
	default:
		console.Logger.Printf("Using env %s\n", strings.ToUpper(chosenEnv))
	}
}
