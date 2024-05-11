package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/transport/grpc/converters"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) UpsertConfig(ctx context.Context, req *api.PatchConfig_Request,
) (*api.PatchConfig_Response, error) {

	err := a.storage.UpsertConfig(ctx, converters.FromProtoConfig(req.Config))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.PatchConfig_Response{}, nil
}
