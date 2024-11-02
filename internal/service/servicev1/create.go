package servicev1

import (
	"context"
	"database/sql"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/service/user_errors"
	"github.com/godverv/matreshka-be/internal/storage"
)

func (c *ConfigService) CreateConfig(ctx context.Context,
	serviceName string, cfg matreshka.AppConfig) error {

	err := c.validator.validateServiceName(serviceName)
	if err != nil {
		return errors.Wrap(err)
	}

	err = c.txManager.Execute(
		func(tx *sql.Tx) error {
			configStorage := c.configStorage.WithTx(tx)

			_, err := configStorage.GetConfigNodes(ctx, serviceName)
			if err != nil {
				if !errors.Is(err, storage.ErrNoNodes) {
					return errors.Wrap(err, "error reading config from storage")
				}
			} else {
				return errors.Wrap(user_errors.ErrAlreadyExists, "Name \""+serviceName+"\" is already taken")
			}

			err = configStorage.SaveConfig(ctx, serviceName, cfg)
			if err != nil {
				return errors.Wrap(err, "error saving config")
			}

			return nil
		})
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
