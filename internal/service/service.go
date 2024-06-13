package service

import (
	"context"

	"github.com/godverv/matreshka-be/internal/domain"
)

type ConfigService interface {
	PatchConfig(ctx context.Context, configPatch domain.PatchConfigRequest) error
}
