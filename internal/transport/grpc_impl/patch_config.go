package grpc_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) PatchConfig(ctx context.Context, req *api.PatchConfig_Request) (*api.PatchConfig_Response, error) {
	err := a.configService.Patch(ctx, fromPatch(req))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.PatchConfig_Response{}, nil
}

func fromPatch(req *api.PatchConfig_Request) domain.PatchConfigRequest {
	out := domain.PatchConfigRequest{
		ConfigName:    req.GetConfigName(),
		ConfigVersion: toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion),
	}

	for _, patch := range req.Patches {

		switch v := patch.GetPatch().(type) {
		case *api.PatchConfig_Patch_Rename:
			out.RenameTo = append(out.RenameTo,
				domain.PatchRename{
					OldName: patch.FieldName,
					NewName: v.Rename,
				})
		case *api.PatchConfig_Patch_UpdateValue:
			out.Update = append(out.Update, domain.PatchUpdate{
				FieldName:  patch.FieldName,
				FieldValue: v.UpdateValue,
			})
		case *api.PatchConfig_Patch_Delete:
			out.Delete = append(out.Delete, patch.FieldName)
		}
	}

	return out
}
