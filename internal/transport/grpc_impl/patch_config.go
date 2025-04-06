package grpc_impl

import (
	"context"
	"strings"

	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) PatchConfig(ctx context.Context, req *api.PatchConfig_Request) (*api.PatchConfig_Response, error) {
	patchReq := domain.PatchConfigRequest{
		ServiceName:   req.GetServiceName(),
		Batch:         fromNodeToPatch(&api.Node{InnerNodes: req.GetChanges()}),
		ConfigVersion: toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion),
	}

	err := a.configService.Patch(ctx, patchReq)
	if err != nil {
		return nil, errors.Wrap(err)
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
