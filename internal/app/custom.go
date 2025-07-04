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
	"go.vervstack.ru/matreshka/internal/transport/matreshka_api_impl"
	"go.vervstack.ru/matreshka/internal/transport/web"
	"go.vervstack.ru/matreshka/internal/transport/web_api"
	"go.vervstack.ru/matreshka/internal/web/auth"
	"go.vervstack.ru/matreshka/pkg/docs"
)

type Custom struct {
	DataProvider storage.Data

	Service service.Services

	GrpcImpl   *matreshka_api_impl.Impl
	WebApiImpl http.Handler
}

func (c *Custom) Init(a *App) (err error) {
	// Repository, Service logic, transport registration happens here
	c.DataProvider = sqlite.New(a.Sqlite)

	txManager := tx_manager.New(a.Sqlite)

	c.Service = v1.New(c.DataProvider, txManager)

	c.GrpcImpl = matreshka_api_impl.NewServer(a.Cfg, c.Service)
	c.WebApiImpl = web_api.New(c.GrpcImpl)

	a.ServerMaster.AddImplementation(c.GrpcImpl)

	a.ServerMaster.AddServerOption(grpc.UnaryInterceptor(user_errors.ErrorInterceptor()))

	if a.Cfg.Environment.Pass != "" {
		a.ServerMaster.AddServerOption(auth.Interceptor(a.Cfg.Environment.Pass))
	}

	a.ServerMaster.AddHttpHandler("/web_api/", c.WebApiImpl)
	a.ServerMaster.AddHttpHandler("/", web.NewServer())
	a.ServerMaster.AddHttpHandler(docs.Swagger())

	return nil
}

func (c *Custom) Start(ctx context.Context) error {
	return nil
}

func (c *Custom) Stop() error {
	return nil
}
