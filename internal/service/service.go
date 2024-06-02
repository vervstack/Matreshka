package service

import (
	"github.com/godverv/matreshka-be/internal/domain"
)

type ConfigService interface {
	Patch(cfg domain.Config, patch domain.PatchConfigRequest) (domain.Config, error)
}
