package domain

import (
	"go.redsock.ru/evon"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

const MasterVersion = "master"

type ConfigWithNodes struct {
	Nodes    *evon.Node
	Versions []string
}

type ConfigName struct {
	prefix api.ConfigTypePrefix
	name   string
}

func NewConfigName(prefix api.ConfigTypePrefix, name string) ConfigName {
	return ConfigName{
		prefix: prefix,
		name:   prefix.String() + "_" + name,
	}
}

func (c ConfigName) Name() string {
	return c.name
}

func (c ConfigName) Prefix() api.ConfigTypePrefix {
	return c.prefix
}
