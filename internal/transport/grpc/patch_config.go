package grpc

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PatchConfig(ctx context.Context, req *api.PatchConfig_Request) (*api.PatchConfig_Response, error) {
	patchReq := domain.PatchConfigRequest{
		ServiceName: req.GetServiceName(),
		Batch:       fromNodeToPatch(&api.Node{InnerNodes: req.GetChanges()}),
	}

	err := a.service.PatchConfig(ctx, patchReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.PatchConfig_Response{}, nil
}

func fromNodeToPatch(root *api.Node) []domain.PatchConfig {
	batch := make([]domain.PatchConfig, 0, len(root.InnerNodes))

	for _, node := range root.InnerNodes {
		patch := domain.PatchConfig{
			FieldName:  node.Name,
			FieldValue: node.Value,
		}
		if !strings.HasPrefix(patch.FieldName, root.Name) {
			patch.FieldName = root.Name + "_" + patch.FieldName
		}
		batch = append(batch, patch)
		batch = append(batch, fromNodeToPatch(node)...)
	}

	return batch
}
