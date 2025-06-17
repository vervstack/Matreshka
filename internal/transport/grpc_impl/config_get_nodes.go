package grpc_impl

import (
	"context"
	"fmt"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka/internal/domain"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) GetConfigNodes(ctx context.Context, req *api.GetConfigNode_Request) (*api.GetConfigNode_Response, error) {
	name := req.GetConfigName()
	ver := toolbox.Coalesce(req.Version, domain.MasterVersion)

	pref, name := ParseConfigName(name)
	if pref == nil {
		pref = toolbox.ToPtr(api.ConfigTypePrefix_plain)
	}

	confName := domain.NewConfigName(*pref, name)

	cfg, err := a.evonConfigService.GetConfigWithNodes(ctx, confName, ver)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp := &api.GetConfigNode_Response{
		Root:     toApiNode(cfg.Nodes),
		Versions: cfg.Versions,
	}

	return resp, nil
}

func toApiNode(node *evon.Node) *api.Node {
	resp := &api.Node{
		Name: node.Name,
	}
	if node.Value != nil {
		v := fmt.Sprint(node.Value)
		resp.Value = &v
	}

	for _, innerNode := range node.InnerNodes {
		resp.InnerNodes = append(resp.InnerNodes, toApiNode(innerNode))
	}

	return resp
}
