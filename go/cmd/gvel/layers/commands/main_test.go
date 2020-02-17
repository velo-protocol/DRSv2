package commands

import (
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/constants"
	"testing"
)

func TestGvelHandler_Init(t *testing.T) {
	gvelHandler := NewGvelHandler(nil, nil)
	gvelHandler.Init()

	assert.Equal(t, constants.CmdRootGvel, gvelHandler.RootCommand.Use)

	assert.Len(t, gvelHandler.RootCommand.Commands(), 4)
	assert.True(t, gvelHandler.AccountCommand == gvelHandler.RootCommand.Commands()[0])
	assert.True(t, gvelHandler.CreditCommand == gvelHandler.RootCommand.Commands()[1])
	assert.True(t, gvelHandler.EnvCommand == gvelHandler.RootCommand.Commands()[2])
	assert.True(t, gvelHandler.InitCommand == gvelHandler.RootCommand.Commands()[3])

	assert.Contains(t, gvelHandler.AccountCommand.Use, constants.CmdAccount)
	assert.Contains(t, gvelHandler.EnvCommand.Use, constants.CmdEnv)
	assert.Contains(t, gvelHandler.CreditCommand.Use, constants.CmdCredit)
	assert.Equal(t, constants.CmdInit, gvelHandler.InitCommand.Use)

	assert.Len(t, gvelHandler.AccountCommand.Commands(), 4)
	assert.Equal(t, constants.CmdAccountCreate, gvelHandler.AccountCommand.Commands()[0].Use)
	assert.Equal(t, constants.CmdAccountDefault, gvelHandler.AccountCommand.Commands()[1].Use)
	assert.Equal(t, constants.CmdAccountImport, gvelHandler.AccountCommand.Commands()[2].Use)
	assert.Equal(t, constants.CmdAccountList, gvelHandler.AccountCommand.Commands()[3].Use)

	assert.Len(t, gvelHandler.CreditCommand.Commands(), 0)

	assert.Len(t, gvelHandler.EnvCommand.Commands(), 1)
	assert.Equal(t, constants.CmdEnvSet, gvelHandler.EnvCommand.Commands()[0].Use)

	assert.Len(t, gvelHandler.InitCommand.Commands(), 0)
}
