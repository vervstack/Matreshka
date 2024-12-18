package grpc_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) CreateConfig(
	ctx context.Context,
	req *api.CreateConfig_Request) (
	*api.CreateConfig_Response, error) {

	var resp api.CreateConfig_Response
	var err error

	resp.Id, err = a.service.CreateConfig(ctx, req.GetServiceName())
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &resp, nil
}
