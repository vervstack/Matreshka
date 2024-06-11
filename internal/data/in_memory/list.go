package in_memory

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) ListConfigs(_ context.Context, _ domain.ListConfigsRequest) ([]matreshka.AppInfo, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	out := make([]matreshka.AppInfo, 0, len(d.data))

	for _, c := range d.data {
		out = append(out, c.appConfig.AppInfo)
	}

	return out, nil
}
