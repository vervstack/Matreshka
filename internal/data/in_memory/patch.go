package in_memory

import (
	"context"
	"strings"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) PatchConfig(ctx context.Context, req domain.PatchConfigRequest) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// TODO сделать обработку переменных
	cfg, ok := d.data[req.ServiceName]
	if !ok {
		return nil
	}

	for _, b := range req.Batch {
		cfg.values[strings.ToUpper(b.FieldPath)] = b.FieldValue
	}

	return nil
}
