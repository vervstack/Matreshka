package storage

import (
	"go.redsock.ru/evon"
	"go.vervstack.ru/matreshka"
)

type Config struct {
	Cfg    *matreshka.AppConfig
	Values map[string]*evon.Node `env:"-"`
	Nodes  []*evon.Node          `env:"-"`
}
