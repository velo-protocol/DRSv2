package account_test

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spf13/cobra"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/commands/account"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/mocks"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/mocks"
	"testing"
)

type helper struct {
	commandHandler *account.CommandHandler
	mockLogic      *mocks.MockLogic
	mockPrompt     *mockutils.MockPrompt
	mockConfig     *mockutils.MockConfiguration
	mockController *gomock.Controller
	logHook        *test.Hook
	tableLogHook   *test.Hook
	done           func()

	cmd        *cobra.Command
	createCmd  *cobra.Command
	listCmd    *cobra.Command
	defaultCmd *cobra.Command
	importCmd  *cobra.Command
	exportCmd  *cobra.Command
}

func initTest(t *testing.T) *helper {
	mockCtrl := gomock.NewController(t)
	mockLogic := mocks.NewMockLogic(mockCtrl)
	mockPrompt := mockutils.NewMockPrompt(mockCtrl)
	mockConfig := mockutils.NewMockConfiguration(mockCtrl)

	handler := account.NewCommandHandler(mockLogic, mockPrompt, mockConfig)
	cmd := handler.Command()

	// logger
	logger, hook := test.NewNullLogger()
	console.Logger = logger

	// table logger
	tableLogger, tableLogHook := test.NewNullLogger()
	console.TableLogger = tableLogger

	// to omit what loader print
	console.DefaultLoadWriter = console.Logger.Out

	return &helper{
		commandHandler: handler,
		mockLogic:      mockLogic,
		mockPrompt:     mockPrompt,
		mockConfig:     mockConfig,
		mockController: mockCtrl,
		logHook:        hook,
		tableLogHook:   tableLogHook,
		done: func() {
			mockCtrl.Finish()
			hook.Reset()
		},

		cmd: cmd,
		// must be order by alphabet
		createCmd:  cmd.Commands()[0],
		defaultCmd: cmd.Commands()[1],
		exportCmd:  cmd.Commands()[2],
		importCmd:  cmd.Commands()[3],
		listCmd:    cmd.Commands()[4],
	}
}
