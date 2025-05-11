package grpc_impl

import (
	"context"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"
	"gopkg.in/yaml.v3"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) GetConfig(ctx context.Context, req *api.GetConfig_Request) (*api.GetConfig_Response, error) {
	name := req.GetConfigName()
	ver := toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion)

	cfg, err := a.configService.GetConfigWithNodes(ctx, name, ver)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp := &api.GetConfig_Response{}

	switch req.Format {
	case api.Format_env:
		resp.Config = evon.Marshal(cfg.Nodes.InnerNodes)
	default:
		nodeStorage := evon.NodesToStorage([]*evon.Node{cfg.Nodes})

		m := make(map[string]any)
		err = evon.UnmarshalWithNodes(nodeStorage, m)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling config")
		}

		resp.Config, err = yaml.Marshal(m)
		if err != nil {
			return nil, errors.Wrap(err, "error marshalling to yaml")
		}
	}

	return resp, nil
}
