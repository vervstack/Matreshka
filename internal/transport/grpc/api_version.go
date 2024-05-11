package grpc

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) ApiVersion(_ context.Context, _ *matreshka_api.ApiVersion_Request,
) (*matreshka_api.ApiVersion_Response, error) {

	resp := &matreshka_api.ApiVersion_Response{
		Version: a.version,
	}

	return resp, nil
}
