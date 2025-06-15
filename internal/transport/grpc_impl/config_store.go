package grpc_impl

import (
	"context"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"
	"google.golang.org/grpc/codes"
	"gopkg.in/yaml.v3"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) StoreConfig(ctx context.Context, req *api.StoreConfig_Request) (
	*api.StoreConfig_Response, error) {

	pref, name := ParseConfigName(req.ConfigName)

	if pref == nil {
		pref = toolbox.ToPtr(api.ConfigTypePrefix_plain)
	}

	cfgName := domain.NewConfigName(*pref, name)

	_, err := a.evonConfigService.Create(ctx, cfgName)
	if err != nil {
		return nil, rerrors.Wrap(err, "error creating config")
	}

	replaceReq := domain.ReplaceConfigReq{
		Name:    cfgName,
		Version: toolbox.Coalesce(req.GetVersion(), domain.MasterVersion),
		Config:  nil,
	}

	switch req.Format {
	case api.Format_env:
		replaceReq.Config, err = fromEvon(req.Config)
	default:
		replaceReq.Config, err = fromYaml(req.Config)
	}

	if err != nil {
		return nil, rerrors.Wrap(err, "error parsing config", codes.InvalidArgument)
	}

	err = a.evonConfigService.Replace(ctx, replaceReq)
	if err != nil {
		return nil, rerrors.Wrap(err, "error updating config")
	}

	return &api.StoreConfig_Response{}, nil
}

func fromYaml(cfg []byte) (*evon.Node, error) {
	yamlMap := map[string]any{}
	err := yaml.Unmarshal(cfg, yamlMap)
	if err != nil {
		return nil, rerrors.Wrap(err, "")
	}

	env, err := evon.MarshalEnv(yamlMap)
	if err != nil {
		return nil, rerrors.Wrap(err, "")
	}

	return env, nil
}

func fromEvon(cfg []byte) (*evon.Node, error) {
	n := &evon.Node{}
	err := evon.Unmarshal(cfg, n)
	if err != nil {
		return nil, rerrors.Wrap(err)
	}
	return n, nil
}
