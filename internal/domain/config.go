package domain

import (
	"go.redsock.ru/evon"
)

const MasterVersion = "master"

type ConfigWithNodes struct {
	Nodes    *evon.Node
	Versions []string
}
