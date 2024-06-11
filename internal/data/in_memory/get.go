package in_memory

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) GetConfig(_ context.Context, req domain.GetConfigReq) (*matreshka.AppConfig, error) {
	cfg := d.getConfig(req.ServiceName)
	return cfg, nil
}

func (d *inMemory) getConfig(serviceName string) *matreshka.AppConfig {
	d.mu.RLock()
	a, _ := d.data[serviceName]
	d.mu.RUnlock()

	if a == nil {
		return nil
	}

	cfg := matreshka.MergeConfigs(matreshka.NewEmptyConfig(), a.appConfig)
	return &cfg
}
