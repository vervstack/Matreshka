package grpc_impl

import (
	"context"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.vervstack.ru/matreshka/internal/domain"
	"go.vervstack.ru/matreshka/pkg/matreshka"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) GetConfig(ctx context.Context, req *api.GetConfig_Request) (*api.GetConfig_Response, error) {
	name := req.GetConfigName()
	ver := toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion)

	cfgNodes, err := a.configService.GetNodes(ctx, name, ver)
	if err != nil {
		return nil, errors.Wrap(err)
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
