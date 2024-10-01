package grpc

import (
	"context"

	"github.com/godverv/matreshka"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) PostConfig(ctx context.Context, req *api.PostConfig_Request) (*api.PostConfig_Response, error) {
	var cfg matreshka.AppConfig
	err := cfg.Unmarshal(req.Content)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = a.service.SaveConfig(ctx, req.GetServiceName(), cfg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.PostConfig_Response{}, nil
}
