package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) GetConfig(ctx context.Context, req *api.GetConfig_Request) (*api.GetConfig_Response, error) {
	r := domain.GetConfigReq{
		ServiceName: req.GetServiceName(),
	}

	cfg, err := a.storage.GetConfig(ctx, r)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if cfg == nil {
		return nil, status.Error(codes.NotFound, "config not found")
	}

	resp := &api.GetConfig_Response{}

	resp.Config, err = cfg.Marshal()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
