package in_memory

import (
	"github.com/Red-Sock/evon"
	"github.com/godverv/matreshka"
)

type Config struct {
	Cfg    *matreshka.AppConfig
	values map[string]*evon.Node `env:"-"`
	nodes  []*evon.Node          `env:"-"`
}
