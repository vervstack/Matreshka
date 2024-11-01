package grpc

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/internal/storage"
	"github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func NewServer(
	cfg config.Config,
	service service.ConfigService,
	storage storage.Data,
) *Impl {
	return &Impl{
		version: cfg.AppInfo.Version,
		service: service,
		storage: storage,
	}
}

func (a *Impl) Register(srv *grpc.Server) {
	matreshka_be_api.RegisterMatreshkaBeAPIServer(srv, a)
}

func (a *Impl) Gateway(ctx context.Context) (route string, handler http.Handler) {
	gwHttpMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{},
		),
	)

	err := matreshka_be_api.RegisterMatreshkaBeAPIHandlerServer(ctx, gwHttpMux, a)
	if err != nil {
		logrus.Errorf("error registering grpc2http handler: %s", err)
	}

	return "/api/*", gwHttpMux
}
