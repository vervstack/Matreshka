package domain

import (
	"github.com/Red-Sock/evon"
)

type ListRequest struct {
	Limit  uint32
	Offset uint32
}

type ListConfigsRequest struct {
	ListRequest
	SearchPattern string
}

type ConfigListItem struct {
	Name    string
	Version string
}

type ConfigDescription struct {
	Id          int64
	ServiceName string
}

type ConfigEnvVals struct {
	ConfigDescription
	Nodes *evon.Node
}
