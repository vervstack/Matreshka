package app

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"go.vervstack.ru/matreshka/internal/service"
	"go.vervstack.ru/matreshka/internal/service/user_errors"
	"go.vervstack.ru/matreshka/internal/service/v1"
	"go.vervstack.ru/matreshka/internal/storage"
	"go.vervstack.ru/matreshka/internal/storage/sqlite"
	"go.vervstack.ru/matreshka/internal/storage/tx_manager"
	"go.vervstack.ru/matreshka/internal/transport/grpc_impl"
	"go.vervstack.ru/matreshka/internal/transport/web"
	"go.vervstack.ru/matreshka/internal/transport/web_api"
	"go.vervstack.ru/matreshka/internal/web/auth"
	"go.vervstack.ru/matreshka/pkg/docs"
)

type Custom struct {
	DataProvider storage.Data

	Service service.Services

	GrpcImpl  *grpc_impl.Impl
	WebClient http.Handler
}

func (c *Custom) Init(a *App) (err error) {
	// Repository, Service logic, transport registration happens here
	c.DataProvider = sqlite.New(a.Sqlite)

	txManager := tx_manager.New(a.Sqlite)

	c.Service = v1.New(c.DataProvider, txManager)

	c.GrpcImpl = grpc_impl.NewServer(a.Cfg, c.Service)

	a.ServerMaster.AddImplementation(c.GrpcImpl)
	a.ServerMaster.AddServerOption(grpc.UnaryInterceptor(user_errors.ErrorInterceptor()))

	if a.Cfg.Environment.Pass != "" {
		a.ServerMaster.AddServerOption(auth.Interceptor(a.Cfg.Environment.Pass))
	}

	c.WebClient = web.NewServer()
	a.ServerMaster.AddHttpHandler("/", c.WebClient)

	a.ServerMaster.AddHttpHandler(docs.Swagger())

	a.ServerMaster.AddHttpHandler("/download/", web_api.New(c.GrpcImpl))
	return nil
}

func (c *Custom) Start(ctx context.Context) error {
	return nil
}

func (c *Custom) Stop() error {
	return nil
}
