package matreshka_be_api_impl

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Impl struct {
	matreshka_be_api.UnimplementedMatreshkaBeAPIServer
}

func New() *Impl {
	return &Impl{}
}

func (impl *Impl) Register(server grpc.ServiceRegistrar) {
	matreshka_be_api.RegisterMatreshkaBeAPIServer(server, impl)
}

func (impl *Impl) Gateway(ctx context.Context, endpoint string, opts ...grpc.DialOption) (route string, handler http.Handler) {
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
