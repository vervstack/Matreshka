package grpc

import (
	"context"

	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PostConfig(ctx context.Context, req *api.PostConfig_Request) (*api.PostConfig_Response, error) {

	return &api.PostConfig_Response{}, nil
}
