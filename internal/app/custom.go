package app

import (
	"net/http"

	"google.golang.org/grpc"

	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/internal/service/servicev1"
	"github.com/godverv/matreshka-be/internal/service/user_errors"
	"github.com/godverv/matreshka-be/internal/storage"
	"github.com/godverv/matreshka-be/internal/storage/sqlite"
	"github.com/godverv/matreshka-be/internal/storage/tx_manager"
	"github.com/godverv/matreshka-be/internal/transport/grpc_impl"
	"github.com/godverv/matreshka-be/internal/transport/web"
	docs "github.com/godverv/matreshka-be/pkg/docs/api"
)

type Custom struct {
	DataProvider storage.Data

	Service service.ConfigService

	GrpcImpl  *grpc_impl.Impl
	WebClient http.Handler
}

func (c *Custom) Init(a *App) (err error) {
	// Repository, Service logic, transport registration happens here
	c.DataProvider = sqlite.New(a.Sqlite)

	txManager := tx_manager.New(a.Sqlite)

	c.Service = servicev1.New(c.DataProvider, txManager)

	c.GrpcImpl = grpc_impl.NewServer(a.Cfg, c.Service, c.DataProvider)

	a.ServerMaster.AddImplementation(c.GrpcImpl,
		grpc.UnaryInterceptor(user_errors.ErrorInterceptor()))

	c.WebClient = web.NewServer()
	a.ServerMaster.AddHttpHandler("/", c.WebClient)

	a.ServerMaster.AddHttpHandler(docs.Swagger())
	return nil
}
