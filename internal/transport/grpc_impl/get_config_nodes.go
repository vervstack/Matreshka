package grpc_impl

import (
	"context"
	"fmt"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"

	"go.verv.tech/matreshka-be/internal/service/user_errors"
	api "go.verv.tech/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) GetConfigNodes(ctx context.Context, req *api.GetConfigNode_Request) (*api.GetConfigNode_Response, error) {
	cfgNodes, err := a.storage.GetConfigNodes(ctx, req.GetServiceName())
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if cfgNodes == nil {
		return nil, errors.Wrap(user_errors.ErrNotFound, "service with name "+req.ServiceName+" not found")
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
