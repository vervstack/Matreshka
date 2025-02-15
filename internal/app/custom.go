package app

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"go.vervstack.ru/matreshka-be/internal/service"
	"go.vervstack.ru/matreshka-be/internal/service/user_errors"
	"go.vervstack.ru/matreshka-be/internal/service/v1"
	"go.vervstack.ru/matreshka-be/internal/storage"
	"go.vervstack.ru/matreshka-be/internal/storage/sqlite"
	"go.vervstack.ru/matreshka-be/internal/storage/tx_manager"
	"go.vervstack.ru/matreshka-be/internal/transport/grpc_impl"
	"go.vervstack.ru/matreshka-be/internal/transport/web"
	docs "go.vervstack.ru/matreshka-be/pkg/docs/api"
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

	c.WebClient = web.NewServer()
	a.ServerMaster.AddHttpHandler("/", c.WebClient)

	a.ServerMaster.AddHttpHandler(docs.Swagger())
	return nil
}

func (c *Custom) Start(ctx context.Context) error {
	return nil
}

func (c *Custom) Stop() error {
	return nil
}
