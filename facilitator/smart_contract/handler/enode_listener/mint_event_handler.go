package enode_listener

import (
	"fmt"
	"gitlab.com/velo-lab/core/facilitator/smart_contract/handler/enode_listener/model"
)

func (handler *ENodeListener) mintEvent(logMint model.LogMint) {
	fmt.Println(logMint)
}
