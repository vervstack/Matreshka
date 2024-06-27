package data

import (
	"context"

	"github.com/Red-Sock/evon"
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

type Data interface {
	GetConfig(ctx context.Context, serviceName domain.GetConfigReq) (*evon.Node, error)
	ListConfigs(ctx context.Context, req domain.ListConfigsRequest) ([]string, error)

	SaveConfig(ctx context.Context, serviceConfig string, config matreshka.AppConfig) error
	PatchConfig(ctx context.Context, req domain.PatchConfigRequest) error
}
