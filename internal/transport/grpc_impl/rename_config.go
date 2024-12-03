package grpc_impl

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) RenameConfig(ctx context.Context, req *api.RenameConfig_Request) (*api.RenameConfig_Response, error) {
	err := a.service.Rename(ctx, req.ServiceName, req.NewName)
	if err != nil {
		return nil, errors.Wrap(err, "error renaming config")
	}

	return &api.RenameConfig_Response{NewName: req.NewName}, nil
}
