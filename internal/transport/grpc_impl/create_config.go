package grpc_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) CreateConfig(
	ctx context.Context,
	req *api.CreateConfig_Request) (
	*api.CreateConfig_Response, error) {

	var pref api.ConfigTypePrefix
	if req.Type != nil {
		pref = *req.Type
	} else {
		pref = api.ConfigTypePrefix_kv
	}

	name := domain.NewConfigName(pref, req.ConfigName)

	var resp api.CreateConfig_Response
	var err error

	aboutConfig, err := a.configService.Create(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp.Name = aboutConfig.Name

	return &resp, nil
}
