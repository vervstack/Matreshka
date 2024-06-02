package v1

import (
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (c *ConfigService) Patch(cfg domain.Config, patch domain.PatchConfigRequest) (domain.Config, error) {
	keys := matreshka.GenerateKeys(cfg.Cfg)
	_ = keys

	for _, kv := range patch.Batch {
		_ = kv
	}

	return cfg, nil
}
