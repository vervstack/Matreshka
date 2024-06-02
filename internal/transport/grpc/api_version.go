package grpc

import (
	"context"

	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) ApiVersion(_ context.Context, _ *api.ApiVersion_Request) (*api.ApiVersion_Response, error) {
	resp := &api.ApiVersion_Response{
		Version: a.version,
	}

	return resp, nil
}
