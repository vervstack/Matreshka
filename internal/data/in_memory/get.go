package in_memory

import (
	"context"

	"github.com/Red-Sock/evon"

	"github.com/godverv/matreshka-be/internal/domain"
)

func (d *inMemory) GetConfig(_ context.Context, req domain.GetConfigReq) (*evon.Node, error) {
	cfg := d.getConfig(req.ServiceName)

	return cfg, nil
}

func (d *inMemory) getConfig(serviceName string) *evon.Node {
	d.mu.RLock()
	a, _ := d.data[serviceName]
	d.mu.RUnlock()

	if a == nil {
		return nil
	}

	return &evon.Node{
		Name:       a.Cfg.Name,
		InnerNodes: a.nodes,
	}
}
