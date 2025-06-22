package matreshka_api_impl

import (
	"context"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"
	"gopkg.in/yaml.v3"

	"go.vervstack.ru/matreshka/internal/domain"
	"go.vervstack.ru/matreshka/pkg/matreshka"
	api "go.vervstack.ru/matreshka/pkg/matreshka_api"
)

func (a *Impl) GetConfig(ctx context.Context, req *api.GetConfig_Request) (*api.GetConfig_Response, error) {
	name := req.GetConfigName()
	ver := toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion)

	pref, name := ParseConfigName(name)
	if pref == nil {
		pref = toolbox.ToPtr(api.ConfigTypePrefix_plain)
	}

	configName := domain.NewConfigName(*pref, name)

	cfg, err := a.evonConfigService.GetConfigWithNodes(ctx, configName, ver)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp := &api.GetConfig_Response{}

	switch req.Format {
	case api.Format_env:
		resp.Config = evon.Marshal(cfg.Nodes.InnerNodes)
	default:
		switch *pref {
		case api.ConfigTypePrefix_verv:
			resp.Config, err = vervToYaml(cfg.Nodes)
		default:
			resp.Config, err = kvToYaml(cfg.Nodes)
		}
		if err != nil {
			return nil, errors.Wrap(err)
		}
	}

	return resp, nil
}

func vervToYaml(node *evon.Node) ([]byte, error) {
	nodeStorage := evon.NodesToStorage(node)

	matreshkaConf := matreshka.NewEmptyConfig()

	err := evon.UnmarshalWithNodes(nodeStorage, &matreshkaConf)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling config")
	}

	config, err := matreshkaConf.Marshal()
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling to yaml")
	}

	return config, nil
}

func kvToYaml(node *evon.Node) ([]byte, error) {
	nodeStorage := evon.NodesToStorage(node)

	m := make(map[string]any)
	err := evon.UnmarshalWithNodes(nodeStorage, m, evon.WithSnakeUnmarshal())
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling config")
	}

	config, err := yaml.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling to yaml")
	}

	return config, nil
}
