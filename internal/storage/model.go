package storage

import (
	"github.com/godverv/matreshka"
	"go.redsock.ru/evon"
)

type Config struct {
	Cfg    *matreshka.AppConfig
	Values map[string]*evon.Node `env:"-"`
	Nodes  []*evon.Node          `env:"-"`
}
