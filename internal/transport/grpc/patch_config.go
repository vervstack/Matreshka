package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PatchConfig(ctx context.Context, req *api.PatchConfig_Request) (*api.PatchConfig_Response, error) {
	var patchReq domain.PatchConfigRequest

	patchReq.ServiceName = req.GetServiceName()
	patchReq.Batch = make([]domain.PatchConfig, 0, len(req.GetChanges()))

	for _, change := range req.GetChanges() {
		patchReq.Batch = append(patchReq.Batch,
			domain.PatchConfig{
				FieldName:  change.Name,
				FieldValue: change.Value,
			})
	}

	err := a.service.PatchConfig(ctx, patchReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.PatchConfig_Response{}, nil
}
