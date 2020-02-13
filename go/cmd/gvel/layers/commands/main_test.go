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

	assert.True(t, gvelHandler.AccountCommand == gvelHandler.RootCommand.Commands()[0])
	assert.True(t, gvelHandler.CreditCommand == gvelHandler.RootCommand.Commands()[1])
	assert.True(t, gvelHandler.EnvCommand == gvelHandler.RootCommand.Commands()[2])
	assert.True(t, gvelHandler.InitCommand == gvelHandler.RootCommand.Commands()[3])

	assert.Len(t, gvelHandler.InitCommand.Commands(), 0)
}
