package app

import (
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/config"
)

func (a *App) InitConfig() (err error) {
	a.Cfg, err = config.Load()
	if err != nil {
		return errors.Wrap(err, "error reading config")
	}

	return nil
}
