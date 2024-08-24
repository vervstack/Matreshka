package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func NewGrpcGateway(ctx context.Context, grpcClient matreshka_api.MatreshkaBeAPIServer) http.Handler {
	gwHttpMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{},
		),
	)

	err := matreshka_api.RegisterMatreshkaBeAPIHandlerServer(ctx, gwHttpMux, grpcClient)
	if err != nil {
		logrus.Errorf("error registering grpc2http handler: %s", err)
	}

	return gwHttpMux
}
