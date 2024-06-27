package in_memory

import (
	"context"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) ListConfigs(_ context.Context, _ domain.ListConfigsRequest) ([]string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	out := make([]string, 0, len(d.data))

	for _, c := range d.data {
		out = append(out, c.Cfg.AppInfo.Name)
	}

	return out, nil
}
