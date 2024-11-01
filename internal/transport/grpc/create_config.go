package grpc

import (
	"context"
	"net/http"
	"time"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) CreateConfig(ctx context.Context, req *api.CreateConfig_Request) (
	*api.CreateConfig_Response, error) {
	var cfg matreshka.AppConfig

	cfg.AppInfo = matreshka.AppInfo{
		Name:            req.ServiceName,
		Version:         "v0.0.1",
		StartupDuration: 5 * time.Second,
	}

	resp, err := a.service.CreateConfig(ctx, req.GetServiceName(), cfg)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	out := &api.CreateConfig_Response{}

	if resp != nil {
		out.ErrorMessage = resp.UserError
		out.HttpStatusCode = http.StatusText(resp.HTTPCode)
	}

	return out, err
}
