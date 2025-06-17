package matreshka_be_api_impl

import (
	"context"
	"fmt"
	"strings"

	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"
	"google.golang.org/grpc/codes"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) PatchConfig(ctx context.Context, req *api.PatchConfig_Request) (*api.PatchConfig_Response, error) {
	patch, err := fromPatch(req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	err = a.evonConfigService.Patch(ctx, patch)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.PatchConfig_Response{}, nil
}

func fromPatch(req *api.PatchConfig_Request) (domain.PatchConfigRequest, error) {
	out := domain.PatchConfigRequest{
		ConfigVersion: toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion),
	}
	var err error
	out.ConfigName, err = fromPlainName(req.GetConfigName())
	if err != nil {
		return domain.PatchConfigRequest{}, errors.Wrap(err)
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
			out.Upsert = append(out.Upsert, domain.PatchUpdate{
				FieldName:  patch.FieldName,
				FieldValue: v.UpdateValue,
			})
		case *api.PatchConfig_Patch_Delete:
			out.Delete = append(out.Delete, patch.FieldName)
		}
	}

	return out, nil
}

func ParseConfigName(name string) (*api.ConfigTypePrefix, string) {
	nameSplited := strings.Split(name, "_")
	if len(nameSplited) < 2 {
		return nil, name
	}

	pref, ok := api.ConfigTypePrefix_value[nameSplited[0]]
	if ok {
		return toolbox.ToPtr(api.ConfigTypePrefix(pref)), strings.Join(nameSplited[1:], "_")
	}

	return nil, name

}
func fromPlainName(name string) (domain.ConfigName, error) {
	nameSplited := strings.Split(name, "_")
	if len(nameSplited) < 2 {
		return domain.ConfigName{},
			errors.NewUserError("Invalid name pattern. Name must start with type prefix", codes.InvalidArgument)
	}

	pref, ok := api.ConfigTypePrefix_value[nameSplited[0]]
	if !ok {
		return domain.ConfigName{},
			errors.NewUserError(fmt.Sprintf("Unknown type prefix: %s", nameSplited[0]), codes.InvalidArgument)
	}

	return domain.NewConfigName(api.ConfigTypePrefix(pref), strings.Join(nameSplited[1:], "_")), nil
}
