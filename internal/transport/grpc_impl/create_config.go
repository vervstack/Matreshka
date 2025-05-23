package grpc_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) CreateConfig(
	ctx context.Context,
	req *api.CreateConfig_Request) (
	*api.CreateConfig_Response, error) {

	name := req.ConfigName

	pref := toolbox.ToPtr(api.ConfigTypePrefix_kv)

	if req.Type != nil {
		pref = req.Type
	}

	parsedPref, name := parseConfigName(req.ConfigName)
	if parsedPref != nil {
		pref = parsedPref
	}

	configName := domain.NewConfigName(*pref, name)

	var resp api.CreateConfig_Response
	var err error

	aboutConfig, err := a.configService.Create(ctx, configName)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp.Name = aboutConfig.Name

	return &resp, nil
}
