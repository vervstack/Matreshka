package grpc

import (
	"context"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) GetConfig(ctx context.Context, req *api.GetConfig_Request) (*api.GetConfig_Response, error) {
	cfgNode, err := a.storage.GetConfigNodes(ctx, req.GetServiceName())
	if err != nil {
		return nil, errors.Wrap(err)
	}

	cfg := matreshka.NewEmptyConfig()

	nodeStorage := evon.NodesToStorage([]*evon.Node{cfgNode})
	err = evon.UnmarshalWithNodes(nodeStorage, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling config")
	}

	resp := &api.GetConfig_Response{}

	resp.Config, err = cfg.Marshal()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return resp, nil
}
