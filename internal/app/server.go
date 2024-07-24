package app

import (
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/transport"
	"github.com/godverv/matreshka-be/internal/transport/grpc"
)

func (a *App) InitServer() error {
	a.Server = transport.NewManager()

	grpcConfig, err := a.Cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		return errors.New("error getting grpc from config")
	}

	a.GrpcApi = grpc.NewServer(a.Cfg, grpcConfig, a.Srv, a.DataProvider)

	a.Server.AddServer(a.GrpcApi)

	return nil
}
