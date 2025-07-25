package config

import (
	"context"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (c *CfgService) Delete(ctx context.Context,
	name domain.ConfigName, version string) error {

	nodes, err := c.configStorage.GetConfigNodes(ctx, name.Name(), version)
	if err != nil {
		return rerrors.Wrap(err, "error getting config nodes for version")
	}

	ns := evon.NodesToStorage(nodes)

	patchReq := domain.PatchConfigRequest{
		ConfigName:    name,
		ConfigVersion: version,
		Delete:        make([]string, 0, len(ns)),
	}

	for n := range ns {
		patchReq.Delete = append(patchReq.Delete, n)
	}

	err = c.configStorage.DeleteValues(ctx, patchReq)
	if err != nil {
		return rerrors.Wrap(err, "error deleting version from storage")
	}

	go func() {
		c.pubService.Publish(patchReq)
	}()

	return nil
}
