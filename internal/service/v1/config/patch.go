package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox"

	"go.vervstack.ru/matreshka/internal/domain"
	"go.vervstack.ru/matreshka/internal/domain/patch"
	"go.vervstack.ru/matreshka/internal/service/user_errors"
)

func (c *CfgService) Patch(ctx context.Context, req domain.PatchConfigRequest) error {
	cfgNodes, err := c.getConfig(ctx, req.ServiceName, req.ConfigVersion)
	if err != nil {
		return rerrors.Wrap(err, "")
	}

	cfgPatch := patch.NewPatch(req.Batch, cfgNodes)

	delReq := domain.PatchConfigRequest{
		ServiceName:   req.ServiceName,
		Batch:         cfgPatch.Delete,
		ConfigVersion: req.ConfigVersion,
	}

	upserReq := domain.PatchConfigRequest{
		ServiceName:   req.ServiceName,
		Batch:         append(cfgPatch.Upsert, cfgPatch.EnvUpsert...),
		ConfigVersion: req.ConfigVersion,
	}

	newName, isChanged := cfgPatch.NameChanged()

	err = c.txManager.Execute(func(tx *sql.Tx) error {
		configStorage := c.configStorage.WithTx(tx)

		err = configStorage.DeleteValues(ctx, delReq)
		if err != nil {
			return rerrors.Wrap(err, "error deleting values")
		}

		err = configStorage.UpsertValues(ctx, upserReq)
		if err != nil {
			return rerrors.Wrap(err, "error patching config in data storage")
		}

		err = configStorage.SetUpdatedAt(ctx, req.ServiceName, time.Now())
		if err != nil {
			return rerrors.Wrap(err, "error updating time")
		}

		if isChanged && req.ConfigVersion == domain.MasterVersion {
			err = configStorage.Rename(ctx, req.ServiceName, newName)
			if err != nil {
				return rerrors.Wrap(err, "error renaming config")
			}
		}

		return nil
	})

	go func() {
		event := domain.PatchConfigRequest{
			ServiceName: req.ServiceName,
			Batch:       append([]domain.PatchConfig{}, cfgPatch.Upsert...),
		}

		event.Batch = append(event.Batch, cfgPatch.EnvUpsert...)
		event.Batch = append(event.Batch, cfgPatch.Delete...)

		c.pubService.Publish(event)
	}()

	if len(cfgPatch.Invalid) != 0 {
		return rerrors.Wrap(user_errors.ErrValidation, "Invalid patched env var name: "+fmt.Sprint(cfgPatch.Invalid))
	}

	return nil
}

func (c *CfgService) getConfig(ctx context.Context, serviceName, version string) (*evon.Node, error) {
	ver := toolbox.Coalesce(version, domain.MasterVersion)

	cfgNodes, err := c.configStorage.GetConfigNodes(ctx, serviceName, ver)
	if err != nil {
		return nil, rerrors.Wrap(err, "error getting nodes")
	}

	if cfgNodes != nil {
		return cfgNodes, nil
	}
	_, err = c.Create(ctx, serviceName)
	if err != nil {
		return nil, rerrors.Wrap(err, "error creating config to patch to")
	}

	return &evon.Node{}, nil
}
