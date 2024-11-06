package grpc_impl

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
	cfgNodes, err := a.storage.GetConfigNodes(ctx, req.GetServiceName())
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if cfgNodes == nil {
		return nil, status.Error(codes.NotFound, "config not found")
	}

	targetConfig := matreshka.NewEmptyConfig()

	nodeStorage := evon.NodesToStorage([]*evon.Node{cfgNodes})

	err = evon.UnmarshalWithNodes(nodeStorage, &targetConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling config")
	}

	resp := &api.GetConfig_Response{}

	resp.Config, err = targetConfig.Marshal()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return resp, nil
}
