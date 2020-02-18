package vfactory

import (
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient/ivclient"
)

type Repository interface {
	NewClient(input *entity.NewClientInput) (ivclient.VClient, error)
}
