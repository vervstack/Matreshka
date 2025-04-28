package config

import (
	"context"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
	"go.vervstack.ru/matreshka/internal/service/user_errors"
)

func (c *CfgService) GetConfigWithNodes(ctx context.Context, serviceName string, ver string) (domain.ConfigWithNodes, error) {
	nodes, err := c.configStorage.GetConfigNodes(ctx, serviceName, ver)
	if err != nil {
		return domain.ConfigWithNodes{}, errors.Wrap(err)
	}

	if nodes == nil {
		return domain.ConfigWithNodes{}, errors.Wrap(user_errors.ErrNotFound, "service with name "+serviceName+" not found")
	}

	versions, err := c.configStorage.GetVersions(ctx, serviceName)
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
