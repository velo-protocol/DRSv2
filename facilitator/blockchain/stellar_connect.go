package blockchain

import (
	"github.com/stellar/go/clients/horizon"
	"gitlab.com/velo-lab/core/facilitator/env"
	"net/http"
)

func ConnectHorizon() *horizon.Client {
	horizonUrl := env.HorizonUrl

	horizonClient := horizon.Client{
		URL:  horizonUrl,
		HTTP: http.DefaultClient,
	}

	return &horizonClient
}
