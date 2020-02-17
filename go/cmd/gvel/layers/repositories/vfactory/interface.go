package vfactory

import (
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)

type Repository interface {
	NewClient(input *entity.NewClientInput) (vclient.VClient, error)
}
