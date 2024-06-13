package in_memory

import (
	"context"

	"github.com/Red-Sock/evon"
	"github.com/godverv/matreshka"
)

func (d *inMemory) UpsertConfig(_ context.Context, cfg matreshka.AppConfig) error {
	c := Config{}
	c.values = map[string]*evon.Node{}
	c.Cfg = &cfg

	nodes := evon.MarshalEnv(&cfg)

	for idx := range nodes {
		c.values[nodes[idx].Name] = &nodes[idx]
		c.nodes = append(c.nodes, &nodes[idx])
	}

	d.mu.Lock()
	d.data[cfg.Name] = &c
	d.mu.Unlock()

	return nil
}
