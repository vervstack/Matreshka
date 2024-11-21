package servicev1

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
)

func (c *ConfigService) Rename(ctx context.Context, oldName, newName string) error {
	err := c.validator.validateServiceName(newName)
	if err != nil {
		return errors.Wrap(err)
	}

	err = c.configStorage.Rename(ctx, oldName, newName)
	if err != nil {
		return errors.Wrap(err, "error during rename operation")
	}

	return nil
}
