package config

import (
	"context"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka-be/internal/domain"
	"go.vervstack.ru/matreshka-be/internal/service/user_errors"
)

func (c *CfgService) GetNodes(ctx context.Context, serviceName string) (*evon.Node, error) {
	cfgNodes, err := c.configStorage.GetConfigNodes(ctx, serviceName)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if cfgNodes == nil {
		return nil, errors.Wrap(user_errors.ErrNotFound, "service with name "+serviceName+" not found")
	}

	return cfgNodes, nil
}

func (c *CfgService) ListConfigs(
	ctx context.Context, req domain.ListConfigsRequest) (domain.ListConfigsResponse, error) {
	resp, err := c.configStorage.ListConfigs(ctx, req)
	if err != nil {
		return domain.ListConfigsResponse{}, errors.Wrap(err)
	}

	return resp, nil
}
