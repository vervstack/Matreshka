package matreshka_api_impl

import (
	"context"

	errors "go.redsock.ru/rerrors"

	api "go.vervstack.ru/matreshka/pkg/matreshka_api"
)

func (s *Impl) CreateConfig(
	ctx context.Context,
	req *api.CreateConfig_Request) (
	*api.CreateConfig_Response, error) {

	name := fromName(req.ConfigName)

	var resp api.CreateConfig_Response
	var err error

	aboutConfig, err := s.evonConfigService.Create(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp.Name = aboutConfig.Name

	return &resp, nil
}
