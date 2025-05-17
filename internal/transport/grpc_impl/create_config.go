package grpc_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) CreateConfig(
	ctx context.Context,
	req *api.CreateConfig_Request) (
	*api.CreateConfig_Response, error) {

	var pref string
	if req.Type != nil {
		pref = req.Type.String()
	} else {
		pref = api.ConfigTypePrefix_kv.String()
	}

	req.ConfigName = pref + "_" + req.ConfigName

	var resp api.CreateConfig_Response
	var err error

	aboutConfig, err := a.configService.Create(ctx, req.ConfigName)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp.Name = aboutConfig.Name

	return &resp, nil
}
