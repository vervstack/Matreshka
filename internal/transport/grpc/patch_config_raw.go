package grpc

import (
	"context"

	"github.com/godverv/matreshka"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PatchConfigRaw(ctx context.Context, patch *matreshka_api.PatchConfigRaw_Request,
) (*matreshka_api.PatchConfigRaw_Response, error) {

	cfg := matreshka.NewEmptyConfig()

	err := cfg.Unmarshal(patch.Raw)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = a.storage.UpsertConfig(ctx, cfg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &matreshka_api.PatchConfigRaw_Response{}, nil
}
