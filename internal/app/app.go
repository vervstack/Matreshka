package app

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/internal/transport"
	"github.com/godverv/matreshka-be/internal/transport/grpc"
	"github.com/godverv/matreshka-be/internal/utils/closer"
)

type App struct {
	Cfg config.Config

	DbConn *sql.DB

	DataProvider data.Data
	Srv          service.ConfigService

	GrpcApi *grpc.Server
	Server  *transport.ServersManager
}

func New() (app App, err error) {
	logrus.Println("starting app")

	err = app.InitConfig()
	if err != nil {
		return App{}, errors.Wrap(err, "error initializing config")
	}

	err = app.InitSqlite(app.Cfg)
	if err != nil {
		return App{}, errors.Wrap(err, "error initializing sqlite storage")
	}

	app.InitService()

	err = app.InitServer()
	if err != nil {
		return app, errors.Wrap(err, "errors initializing servers")
	}

	return app, nil
}

func (a *App) Start() error {
	ctx := context.Background()

	err := a.Server.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "error starting Server manager")
	}

	closer.Add(func() error { return a.Server.Stop(ctx) })

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
