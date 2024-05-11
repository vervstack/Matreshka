package in_memory

import (
	"context"

	"github.com/godverv/matreshka"
)

func (d *inMemory) UpsertConfig(_ context.Context, cfg matreshka.AppConfig) error {
	d.m.Lock()
	d.mp[cfg.Name] = &cfg
	d.m.Unlock()

	return nil
}
