package domain

import (
	"go.redsock.ru/evon"

	"go.vervstack.ru/matreshka/pkg/matreshka"
)

const MasterVersion = "master"

type Config struct {
	Cfg matreshka.AppConfig
}

type ConfigWithNodes struct {
	Nodes    *evon.Node
	Versions []string
}
