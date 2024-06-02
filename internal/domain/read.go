package domain

type GetConfigReq struct {
	ServiceName string
}

type ListRequest struct {
	Limit  uint32
	Offset uint32
}

type ListConfigsRequest struct {
	ListRequest
	SearchPattern string
}
