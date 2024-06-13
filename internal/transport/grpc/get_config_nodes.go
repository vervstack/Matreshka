package grpc

import (
	"context"
	"fmt"

	"github.com/Red-Sock/evon"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) GetConfigNodes(ctx context.Context, req *api.GetConfigNode_Request) (*api.GetConfigNode_Response, error) {
	getConfigReq := domain.GetConfigReq{ServiceName: req.GetServiceName()}
	cfgNode, err := a.storage.GetConfig(ctx, getConfigReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if cfgNode == nil {
		return nil, status.Error(codes.NotFound, "config not found")
	}

	resp := &api.GetConfigNode_Response{
		Root: toApiNode(*cfgNode),
	}

	return resp, nil
}

func toApiNode(node evon.Node) *api.Node {
	resp := &api.Node{
		Name:  node.Name,
		Value: fmt.Sprint(node.Value),
	}

	for _, innerNode := range node.InnerNodes {
		resp.InnerNodes = append(resp.InnerNodes, toApiNode(*innerNode))
	}

	return resp
}
