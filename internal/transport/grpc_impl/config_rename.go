package grpc_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) RenameConfig(ctx context.Context, req *api.RenameConfig_Request) (*api.RenameConfig_Response, error) {
	oldName, err := fromPlainName(req.ConfigName)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	newName, err := fromPlainName(req.NewName)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	err = a.evonConfigService.Rename(ctx, oldName, newName)
	if err != nil {
		return nil, errors.Wrap(err, "error renaming config")
	}

	return &api.RenameConfig_Response{NewName: req.NewName}, nil
}
