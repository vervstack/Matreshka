package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	//_transport_imports

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data/sqlite"
	v1 "github.com/godverv/matreshka-be/internal/service/v1"
	"github.com/godverv/matreshka-be/internal/transport"
	"github.com/godverv/matreshka-be/internal/transport/grpc"
	"github.com/godverv/matreshka-be/internal/utils/closer"
)

func main() {
	err := startServer()
	if err != nil {
		logrus.Fatal(err)
	}
}

func startServer() error {
	logrus.Println("starting app")

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return errors.Wrap(err, "error reading config")
	}

	if cfg.GetAppInfo().StartupDuration == 0 {
		return errors.New("no startup duration in config")
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.GetAppInfo().StartupDuration)
	closer.Add(func() error {
		cancel()
		return nil
	})

	transportManager := transport.NewManager()

	grpcConfig, err := cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		return errors.New("error getting grpc from config")
	}

	sqliteCfg, err := cfg.GetDataSources().Sqlite(config.ResourceSqlite)
	if err != nil {
		return errors.Wrap(err, "error getting sqlite from config")
	}
	data, err := sqlite.New(sqliteCfg)

	srv := v1.New(data)

	grpcServer, err := grpc.NewServer(cfg, grpcConfig, srv, data)
	if err != nil {
		return errors.New("error creating grpc server")
	}

	transportManager.AddServer(grpcServer)

	err = transportManager.Start(ctx)
	if err != nil {
		return errors.New("error starting server manager")
	}

	closer.Add(
		func() error {
			return transportManager.Stop(ctx)
		})
	waitingForTheEnd()

	logrus.Println("shutting down the app")

	err = closer.Close()
	if err != nil {
		return errors.Wrap(err, "error while shutting down application")
	}

	return nil
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is an LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
