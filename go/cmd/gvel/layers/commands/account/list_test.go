package account_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"strings"
	"testing"
)

func TestCommandHandler_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return([]*entity.Account{
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1",
					IsDefault:     false,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2",
					IsDefault:     true,
				},
				{
					PublicAddress: "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA3",
					IsDefault:     false,
				},
			}, nil)

		h.commandHandler.List(h.cmd, nil)
		lines := strings.Split(h.tableLogHook.LastEntry().Message, "\n")

		assert.Contains(t, lines[1], "INDEX")
		assert.Contains(t, lines[1], "ADDRESS")
		assert.Contains(t, lines[1], "DEFAULT")

		assert.Contains(t, lines[3], "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA1")
		assert.Contains(t, lines[4], "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA2")
		assert.Contains(t, lines[5], "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA3")

		assert.Contains(t, lines[3], "false")
		assert.Contains(t, lines[4], "true")
		assert.Contains(t, lines[5], "false")
	})
	t.Run("fail, list account returns error", func(t *testing.T) {
		h := initTest(t)
		defer h.done()

		h.mockLogic.EXPECT().
			ListAccount().
			Return(nil, errors.New("some error has occurred"))

		assert.PanicsWithValue(t, console.ExitError, func() {
			h.commandHandler.List(h.listCmd, nil)
		})
	})
}
