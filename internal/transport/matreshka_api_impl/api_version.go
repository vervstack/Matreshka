package matreshka_api_impl

import (
	"context"

	api "go.vervstack.ru/matreshka/pkg/matreshka_api"
)

func (s *Impl) ApiVersion(_ context.Context, _ *api.ApiVersion_Request) (*api.ApiVersion_Response, error) {
	resp := &api.ApiVersion_Response{
		Version: s.version,
	}

	return resp, nil
}
