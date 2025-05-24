package config

import (
	"context"
	"strings"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
	"go.vervstack.ru/matreshka/internal/service/user_errors"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (c *CfgService) GetConfigWithNodes(ctx context.Context, configName domain.ConfigName, ver string) (domain.ConfigWithNodes, error) {
	nodes, err := c.configStorage.GetConfigNodes(ctx, configName.Name(), ver)
	if err != nil {
		return domain.ConfigWithNodes{}, errors.Wrap(err)
	}

	switch configName.Prefix() {
	case api.ConfigTypePrefix_pg:
		toSnake(nodes)
	}

	if nodes == nil {
		return domain.ConfigWithNodes{}, errors.Wrap(user_errors.ErrNotFound, "service with name "+configName.Name()+" not found")
	}

	versions, err := c.configStorage.GetVersions(ctx, configName.Name())
	if err != nil {
		return domain.ConfigWithNodes{}, errors.Wrap(err, "error getting config by name")
	}

	return domain.ConfigWithNodes{
		Nodes:    nodes,
		Versions: versions,
	}, nil
}

func (c *CfgService) ListConfigs(
	ctx context.Context, req domain.ListConfigsRequest) (domain.ListConfigsResponse, error) {
	resp, err := c.configStorage.ListConfigs(ctx, req)
	if err != nil {
		return domain.ListConfigsResponse{}, errors.Wrap(err)
	}

	return resp, nil
}

func toSnake(root *evon.Node) {
	if root == nil {
		return
	}

	for _, n := range root.InnerNodes {
		toSnake(n)
	}
	root.Name = strings.ReplaceAll(root.Name, "-", "_")
}
