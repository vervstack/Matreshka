package data

import (
	"github.com/Red-Sock/evon"
	"github.com/godverv/matreshka"
)

type Config struct {
	Cfg    *matreshka.AppConfig
	Values map[string]*evon.Node `env:"-"`
	Nodes  []*evon.Node          `env:"-"`
}
