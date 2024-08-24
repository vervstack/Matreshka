package app

import (
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/transport"
	"github.com/godverv/matreshka-be/internal/transport/grpc"
	"github.com/godverv/matreshka-be/internal/transport/web_client"
)

func (a *App) InitServer() error {
	var err error
	a.Server, err = transport.NewManager(":8080") // TODO
	if err != nil {
		return errors.Wrap(err, "error creating new server manager")
	}

	a.Server.AddGrpcServer(grpc.NewServer(a.Cfg, a.Srv, a.DataProvider))

	a.Server.AddRestServer("/*", web_client.NewServer())

	return nil
}
