package v1

import (
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (c *ConfigService) Patch(cfg domain.Config, patch domain.PatchConfigRequest) (domain.Config, error) {
	keys, err := matreshka.GenerateKeys(cfg.Cfg)
	if err != nil {
		return cfg, errors.Wrap(err, "error generating map")
	}

	_ = keys

	for _, kv := range patch.Batch {
		_ = kv
	}

	return cfg, nil
}
