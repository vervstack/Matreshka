package grpc_impl

import (
	"context"
	"fmt"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka-be/internal/domain"
	api "go.vervstack.ru/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) GetConfigNodes(ctx context.Context, req *api.GetConfigNode_Request) (*api.GetConfigNode_Response, error) {
	name := req.GetServiceName()
	ver := toolbox.Coalesce(toolbox.FromPtr(req.Version), domain.MasterVersion)

	cfgNodes, err := a.configService.GetNodes(ctx, name, ver)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp := &api.GetConfigNode_Response{
		Root: toApiNode(cfgNodes),
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
