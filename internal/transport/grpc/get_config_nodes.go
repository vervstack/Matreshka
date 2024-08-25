package grpc

import (
	"context"
	"fmt"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) GetConfigNodes(ctx context.Context, req *api.GetConfigNode_Request) (*api.GetConfigNode_Response, error) {
	cfgNode, err := a.storage.GetConfig(ctx, req.GetServiceName())
	if err != nil {
		return nil, errors.Wrap(err)
	}

	resp := &api.GetConfigNode_Response{
		Root: toApiNode(cfgNode),
	}

	return resp, nil
}

func toApiNode(node evon.Node) *api.Node {
	resp := &api.Node{
		Name: node.Name,
	}
	if node.Value != nil {
		v := fmt.Sprint(node.Value)
		resp.Value = &v
	}

	for _, innerNode := range node.InnerNodes {
		resp.InnerNodes = append(resp.InnerNodes, toApiNode(*innerNode))
	}

	return resp
}
