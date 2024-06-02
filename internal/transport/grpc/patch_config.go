package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PatchConfig(ctx context.Context, req *api.PatchConfig_Request) (*api.PatchConfig_Response, error) {
	patchReq := domain.PatchConfigRequest{
		ServiceName: req.GetServiceName(),
		Batch:       make([]domain.PatchConfig, len(req.GetPathToValue())),
	}

	for k, v := range req.GetPathToValue() {
		patchReq.Batch = append(patchReq.Batch, domain.PatchConfig{
			FieldPath:  k,
			FieldValue: v,
		})
	}

	cfg, _ := a.storage.GetConfig(ctx, domain.GetConfigReq{ServiceName: req.ServiceName})

	cfg2, err := a.cfgService.Patch(domain.Config{Cfg: *cfg}, patchReq)

	_ = cfg2

	err = a.storage.PatchConfig(ctx, patchReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.PatchConfig_Response{}, nil
}
