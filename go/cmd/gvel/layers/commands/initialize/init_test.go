package initialize_test

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
	"testing"
)

func TestCommandHandler_Init(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		helper := initTest(t)
		defer helper.done()

		helper.mockConfig.EXPECT().
			Exists().
			Return(false)

		helper.mockLogic.EXPECT().
			Init(constants.FsBaseDir).
			Return(nil)

		helper.commandHandler.Init(nil, nil)

		logEntries := helper.loggerHook.AllEntries()
		assert.Equal(t, "gvel has been initialized\n", logEntries[0].Message)
		assert.Equal(t, fmt.Sprintf("using config file at: %s\n", constants.FsBaseDir), logEntries[1].Message)
		assert.Equal(t, fmt.Sprintf("gvel is currently connected to %s. Please use `gvel env set` to change the env.\n", strings.ToUpper(constants.EnvTestNet)), logEntries[2].Message)
	})

	t.Run("fail, config already exist", func(t *testing.T) {
		helper := initTest(t)
		defer helper.done()

		helper.mockConfig.EXPECT().
			Exists().
			Return(true)

		assert.PanicsWithValue(t, console.ExitError, func() {
			helper.commandHandler.Init(nil, nil)
		})
		assert.Equal(t, fmt.Sprintf("gvel has already been initialized, configuration can be found at %s", constants.FsBaseDir), helper.loggerHook.LastEntry().Message)
	})

	t.Run("fail, logic returns error", func(t *testing.T) {
		helper := initTest(t)
		defer helper.done()

		helper.mockConfig.EXPECT().
			Exists().
			Return(false)

		helper.mockLogic.EXPECT().
			Init(constants.FsBaseDir).
			Return(errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			helper.commandHandler.Init(nil, nil)
		})
	})
}
