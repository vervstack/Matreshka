package data

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

type Data interface {
	GetConfig(ctx context.Context, serviceName domain.GetConfigReq) (*matreshka.AppConfig, error)
	UpsertConfig(ctx context.Context, config matreshka.AppConfig) error
	ListConfigs(ctx context.Context, req domain.ListConfigsRequest) ([]matreshka.AppInfo, error)
}
