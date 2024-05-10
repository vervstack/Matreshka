package grpc

import (
	"context"

	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) ListConfigs(ctx context.Context, req *api.ListConfigs_Request,
) (*api.ListConfigs_Response, error) {
	return &api.ListConfigs_Response{
		Services: []*api.ListConfigs_Response_ServiceInfo{
			{
				Name: "velez",
			},
			{
				Name: "matreshka",
			},
		},
	}, nil
}
