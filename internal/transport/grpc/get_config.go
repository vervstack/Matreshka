package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	"github.com/godverv/matreshka-be/internal/transport/grpc/converters"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) GetConfig(ctx context.Context, req *matreshka_api.GetConfig_Request,
) (*matreshka_api.GetConfig_Response, error) {
	config, err := a.storage.GetConfig(ctx, getConfigRequest(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if config == nil {
		return nil, status.Error(codes.NotFound, "config not found")
	}
	resp := &matreshka_api.GetConfig_Response{}

	resp.Config = converters.ToProtoConfig(*config)

	return resp, nil
}

func getConfigRequest(in *matreshka_api.GetConfig_Request) domain.GetConfigReq {
	out := domain.GetConfigReq{}

	out.ServiceName = in.GetServiceName()

	return out
}
