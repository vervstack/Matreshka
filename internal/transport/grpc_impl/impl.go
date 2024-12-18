package grpc_impl

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"go.verv.tech/matreshka-be/internal/config"
	"go.verv.tech/matreshka-be/internal/service"
	"go.verv.tech/matreshka-be/internal/storage"
	"go.verv.tech/matreshka-be/pkg/matreshka_be_api"
)

type Impl struct {
	version string
	service service.ConfigService
	storage storage.Data

	matreshka_be_api.UnimplementedMatreshkaBeAPIServer
}

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

func (a *Impl) Gateway(ctx context.Context, endpoint string, opts ...grpc.DialOption) (route string, handler http.Handler) {
	gwHttpMux := runtime.NewServeMux()

	err := matreshka_be_api.RegisterMatreshkaBeAPIHandlerFromEndpoint(
		ctx,
		gwHttpMux,
		endpoint,
		opts,
	)
	if err != nil {
		logrus.Errorf("error registering grpc2http handler: %s", err)
	}

	return "/api/", gwHttpMux
}
