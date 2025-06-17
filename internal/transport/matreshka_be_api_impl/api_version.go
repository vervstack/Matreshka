package matreshka_be_api_impl

import (
	"context"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) ApiVersion(_ context.Context, _ *api.ApiVersion_Request) (*api.ApiVersion_Response, error) {
	resp := &api.ApiVersion_Response{
		Version: a.version,
	}

	return resp, nil
}
