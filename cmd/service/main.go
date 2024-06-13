package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data/in_memory"
	v1 "github.com/godverv/matreshka-be/internal/service/v1"
	"github.com/godverv/matreshka-be/internal/transport"
	"github.com/godverv/matreshka-be/internal/transport/grpc"
	"github.com/godverv/matreshka-be/internal/utils/closer"
	//_transport_imports
)

func main() {
	logrus.Println("starting app")

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("error reading config %s", err.Error())
	}

	if cfg.GetAppInfo().StartupDuration == 0 {
		logrus.Fatalf("no startup duration in config")
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.GetAppInfo().StartupDuration)
	closer.Add(func() error {
		cancel()
		return nil
	})

	mngr := transport.NewManager()
	grpcConfig, err := cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		logrus.Fatalf("error getting grpc from config")
	}

	data := in_memory.New()
	srv := v1.New(data)

	grpcServer, err := grpc.NewServer(cfg, grpcConfig, srv, data)
	if err != nil {
		logrus.Fatalf("error creating grpc server")
	}

	mngr.AddServer(grpcServer)

	err = mngr.Start(ctx)
	if err != nil {
		logrus.Fatalf("error starting server manager")
	}

	closer.Add(
		func() error {
			return mngr.Stop(ctx)
		})
	waitingForTheEnd()

	logrus.Println("shutting down the app")

	if err = closer.Close(); err != nil {
		logrus.Fatalf("errors while shutting down application %s", err.Error())
	}
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is a LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
