package service

import (
	"context"

	"go.verv.tech/matreshka-be/internal/domain"
)

type ConfigService interface {
	PatchConfig(ctx context.Context, configPatch domain.PatchConfigRequest) error
	CreateConfig(ctx context.Context, serviceName string) (int64, error)
	Rename(ctx context.Context, oldName, newName string) error
}
