package domain

import (
	"time"

	"go.redsock.ru/evon"

	api "go.vervstack.ru/matreshka-be/pkg/matreshka_be_api"
)

type ListConfigsRequest struct {
	Paging Paging
	Sort   Sort

	SearchPattern string
}

type ListConfigsResponse struct {
	List         []ConfigListItem
	TotalRecords uint32
}

type ConfigListItem struct {
	Name      string
	Version   string
	UpdatedAt time.Time
}

type ConfigDescription struct {
	Id          int64
	ServiceName string
}

type ConfigEnvVals struct {
	ConfigDescription
	Nodes *evon.Node
}

type Paging struct {
	Limit  uint32
	Offset uint32
}

type Sort struct {
	SortType api.Sort_Type
	Desc     bool
}
