package app

import (
	"context"
	"database/sql"

	"github.com/Red-Sock/toolbox"
	"github.com/Red-Sock/toolbox/closer"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/internal/transport"
)

type App struct {
	Ctx  context.Context
	Stop func()
	Cfg  config.Config

	DbConn *sql.DB

	DataProvider data.Data
	Srv          service.ConfigService

	Server *transport.ServersManager
}

func New() (app App, err error) {
	logrus.Println("starting app")

	app.Ctx, app.Stop = context.WithCancel(context.Background())

	err = app.InitConfig()
	if err != nil {
		return App{}, errors.Wrap(err, "error initializing config")
	}

	err = app.InitSqlite()
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
	err := a.Server.Start(a.Ctx)
	if err != nil {
		return errors.Wrap(err, "error starting Server manager")
	}
	closer.Add(func() error { return a.Server.Stop() })

	toolbox.WaitForInterrupt()

	logrus.Println("shutting down the app")

	err = closer.Close()
	if err != nil {
		return errors.Wrap(err, "error while shutting down application")
	}

	return nil
}
