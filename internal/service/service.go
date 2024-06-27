package service

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

type ConfigService interface {
	PatchConfig(ctx context.Context, configPatch domain.PatchConfigRequest) error
	SaveConfig(ctx context.Context, serviceName string, cfg matreshka.AppConfig) error
}
