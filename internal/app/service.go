package app

import (
	v1 "github.com/godverv/matreshka-be/internal/service/v1"
)

func (a *App) InitService() {
	a.Srv = v1.New(a.DataProvider)
}
