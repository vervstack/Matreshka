package domain

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
