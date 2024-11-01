package servicev1

import (
	"context"
	"database/sql"
	"net/http"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/internal/storage"
)

func (c *ConfigService) CreateConfig(ctx context.Context, serviceName string, cfg matreshka.AppConfig) (
	*service.Response, error) {

	var resp service.Response
	err := c.txManager.Execute(func(tx *sql.Tx) error {
		configStorage := c.configStorage.WithTx(tx)

		_, err := configStorage.GetConfigNodes(ctx, serviceName)
		if err != nil {
			if !errors.Is(err, storage.ErrNoNodes) {
				return errors.Wrap(err, "error reading config from storage")
			}
		} else {
			resp.UserError = "Config already exists"
			resp.HTTPCode = http.StatusConflict
			return nil
		}

		err = configStorage.SaveConfig(ctx, serviceName, cfg)
		if err != nil {
			return errors.Wrap(err, "error saving config")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error executing transaction")
	}

	if resp.UserError != "" {
		return &resp, nil
	}

	return nil, nil
}
