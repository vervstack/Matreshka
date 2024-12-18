package servicev1

import (
	"context"
	"database/sql"
	"time"

	"github.com/godverv/matreshka"
	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"

	"github.com/godverv/matreshka-be/internal/domain"
	"github.com/godverv/matreshka-be/internal/service/user_errors"
)

func (c *ConfigService) CreateConfig(ctx context.Context, serviceName string) (int64, error) {
	err := c.validator.validateServiceName(serviceName)
	if err != nil {
		return 0, errors.Wrap(err)
	}

	newCfg := c.initNewConfig(serviceName)

	var createdCfgId int64
	newCfgPatch, err := c.convertConfigToPatch(newCfg)
	if err != nil {
		return 0, errors.Wrap(err, "error converting config to patch")
	}

	err = c.txManager.Execute(func(tx *sql.Tx) error {
		configStorage := c.configStorage.WithTx(tx)

		var nodes *evon.Node
		nodes, err = configStorage.GetConfigNodes(ctx, serviceName)
		if err != nil {
			return errors.Wrap(err, "error reading config from storage")
		}

		if nodes != nil {
			return errors.Wrap(user_errors.ErrAlreadyExists, "Name \""+serviceName+"\" is already taken")
		}

		createdCfgId, err = configStorage.Create(ctx, serviceName)
		if err != nil {
			return errors.Wrap(err, "error saving config")
		}

		err = configStorage.UpsertValues(ctx, serviceName, newCfgPatch)
		if err != nil {
			return errors.Wrap(err, "error upserting new config")
		}
		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	return createdCfgId, nil
}

func (c *ConfigService) convertConfigToPatch(cfg matreshka.AppConfig) ([]domain.PatchConfig, error) {
	newCfgNodes, err := evon.MarshalEnv(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling config")
	}
	newCfgNodesStore := evon.NodesToStorage(newCfgNodes.InnerNodes)

	cfgPatch := make([]domain.PatchConfig, 0, len(newCfgNodesStore))
	for _, node := range newCfgNodesStore {
		if node.Value != nil {
			cfgPatch = append(cfgPatch,
				domain.PatchConfig{
					FieldName:  node.Name,
					FieldValue: node.Value,
				})
		}
	}
	return cfgPatch, nil
}

func (c *ConfigService) initNewConfig(serviceName string) matreshka.AppConfig {
	newCfg := matreshka.NewEmptyConfig()

	newCfg.AppInfo = matreshka.AppInfo{
		Name:            serviceName,
		Version:         "v0.0.1",
		StartupDuration: time.Second * 5,
	}

	return newCfg
}
