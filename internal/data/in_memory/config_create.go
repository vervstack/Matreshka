package in_memory

import (
	"context"
	"time"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (d *data) CreateConfig(_ context.Context, req *matreshka_api.CreateServiceConfig_Request) error {
	c := &config{
		cfg: matreshka.NewEmptyConfig(),
	}

	c.cfg.AppInfo = matreshka.AppInfo{
		Name:    req.AppConfig.Name,
		Version: req.AppConfig.Version,
	}

	c.cfg.StartupDuration, _ = time.ParseDuration(req.AppConfig.StartupDuration)

	d.m.Lock()
	d.mp[req.AppConfig.Name] = c
	d.m.Unlock()

	return nil
}
