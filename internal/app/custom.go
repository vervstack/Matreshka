package app

import (
	"net/http"

	"google.golang.org/grpc"

	"go.verv.tech/matreshka-be/internal/service"
	"go.verv.tech/matreshka-be/internal/service/user_errors"
	"go.verv.tech/matreshka-be/internal/service/v1"
	"go.verv.tech/matreshka-be/internal/storage"
	"go.verv.tech/matreshka-be/internal/storage/sqlite"
	"go.verv.tech/matreshka-be/internal/storage/tx_manager"
	"go.verv.tech/matreshka-be/internal/transport/grpc_impl"
	"go.verv.tech/matreshka-be/internal/transport/web"
	docs "go.verv.tech/matreshka-be/pkg/docs/api"
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

	a.ServerMaster.AddImplementation(c.GrpcImpl,
		grpc.UnaryInterceptor(user_errors.ErrorInterceptor()))

	c.WebClient = web.NewServer()
	a.ServerMaster.AddHttpHandler("/", c.WebClient)

	a.ServerMaster.AddHttpHandler(docs.Swagger())
	return nil
}
