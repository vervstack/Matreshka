package in_memory

import (
	"context"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) ListConfigs(_ context.Context, _ domain.ListConfigsRequest) ([]matreshka.AppInfo, error) {
	d.m.RLock()
	defer d.m.RUnlock()

	out := make([]matreshka.AppInfo, 0, len(d.mp))

	for _, c := range d.mp {
		out = append(out, c.AppInfo)
	}

	return out, nil
}
