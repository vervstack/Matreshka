package app

import (
	errors "github.com/Red-Sock/trace-errors"

	sqliteclient "github.com/godverv/matreshka-be/internal/clients/sqlite"
	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data/storage"
)

func (a *App) InitSqlite() (err error) {
	sqliteConf, err := a.Cfg.GetDataSources().Sqlite(config.ResourceSqlite)
	if err != nil {
		return errors.Wrap(err, "error getting sqlite from config")
	}

	a.DbConn, err = sqliteclient.NewConn(sqliteConf)
	if err != nil {
		return errors.Wrap(err, "error getting sqlite connection")
	}

	a.DataProvider, err = storage.New(a.DbConn)
	if err != nil {
		return errors.Wrap(err, "error initializing sqlite")
	}

	return nil
}
