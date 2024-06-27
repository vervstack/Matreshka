package v1

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
)

func (c *ConfigService) SaveConfig(ctx context.Context, serviceName string, cfg matreshka.AppConfig) error {
	err := c.data.SaveConfig(ctx, serviceName, cfg)
	if err != nil {
		return errors.Wrap(err, "error saving config")
	}

	return nil
}
