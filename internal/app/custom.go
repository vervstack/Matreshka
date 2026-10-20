package app

import (
	"context"

	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/middleware"
	"go.vervstack.ru/matreshka/internal/service"
	"go.vervstack.ru/matreshka/internal/service/v1"
	"go.vervstack.ru/matreshka/internal/storage"
	"go.vervstack.ru/matreshka/internal/storage/sqlite"
	"go.vervstack.ru/matreshka/internal/transport"

	//"go.vervstack.ru/matreshka/internal/storage/pg"
	"go.vervstack.ru/matreshka/internal/storage/tx_manager"
	"go.vervstack.ru/matreshka/internal/transport/matreshka_api_impl"
	"go.vervstack.ru/matreshka/internal/version"
	"go.vervstack.ru/matreshka/internal/web/auth"
	"go.vervstack.ru/matreshka/pkg/docs"
)

type Custom struct {
	SqliteData storage.Data

	Service service.Services

	ServerMaster *transport.ServersManager

	GrpcImpl *matreshka_api_impl.Impl
}

func (c *Custom) Init(a *App) (err error) {
	a.Cfg.AppInfo.Version = version.Get()

	txManager := tx_manager.New(a.Sqlite)
	c.SqliteData = sqlite.New(a.Sqlite)

	c.Service = v1.New(c.SqliteData, txManager)

	c.GrpcImpl = matreshka_api_impl.NewServer(a.Cfg, c.Service)

	c.ServerMaster, err = transport.NewServerManager(a.Ctx, a.MASTER)
	if err != nil {
		return rerrors.Wrap(err)
	}

	c.ServerMaster.AddImplementation(c.GrpcImpl)

	if a.Cfg.Environment.Pass != "" {
		c.ServerMaster.AddServerOption(auth.Interceptor(a.Cfg.Environment.Pass))
	}

	c.ServerMaster.AddServerOption(
		middleware.PanicInterceptor(),
		middleware.LogInterceptor(),
	)

	c.ServerMaster.AddHttpHandler(docs.Swagger())

	return nil
}

func (c *Custom) Start(_ context.Context) error {
	err := c.ServerMaster.Start()
	if err != nil {
		return rerrors.Wrap(err)
	}

	return nil
}

func (c *Custom) Stop() error {
	return nil
}
