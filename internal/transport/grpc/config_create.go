package grpc

import (
	"context"

	api "github.com/godverv/matreshka-be/pkg/api/matreshka_api"
)

func (a *App) CreateServiceConfig(ctx context.Context, req *api.CreateServiceConfig_Request) (*api.CreateServiceConfig_Response, error) {
	return &api.CreateServiceConfig_Response{}, a.storage.CreateConfig(ctx, req)
}
