package logic

import (
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"strings"
)

func (lo *logic) SetEnv(input *entity.SetEnvInput) error {
	if input.Env == "" {
		return errors.New("env must not be empty")
	}

	return lo.AppConfig.SetCurrentEnv(strings.ToLower(input.Env))
}
