package in_memory

import (
	"context"
	"strings"

	"github.com/Red-Sock/evon"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) PatchConfig(ctx context.Context, req domain.PatchConfigRequest) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	cfg, ok := d.data[req.ServiceName]
	if !ok {
		return nil
	}

	for _, b := range req.Batch {
		v := cfg.values[strings.ToUpper(b.FieldName)]
		if v == nil {
			v = &evon.Node{
				Name: b.FieldName,
			}

			cfg.values[strings.ToUpper(b.FieldName)] = v
			cfg.nodes = append(cfg.nodes, v)
		}
		v.Value = b.FieldValue
	}

	return nil
}
