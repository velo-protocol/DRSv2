package main

import (
	"gitlab.com/velo-lab/core/facilitator/blockchain"
	"gitlab.com/velo-lab/core/facilitator/env"
	"gitlab.com/velo-lab/core/facilitator/smart_contract/handler/enode_listener"
	"log"
)

func main() {
	env.Init()

	enode := blockchain.ConnectENode()

	err := enode_listener.NewENodeListener(enode)
	if err != nil {
		log.Fatal(err)
	}
}
