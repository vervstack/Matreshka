package in_memory

import (
	"context"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
)

func (d *inMemory) UpsertConfig(_ context.Context, cfg matreshka.AppConfig) error {
	c := Config{}
	c.values = map[string]*evon.Node{}
	c.Cfg = &cfg

	nodes, err := evon.MarshalEnv(&cfg)
	if err != nil {
		return errors.Wrap(err, "error marshalling config to variables")
	}

	for idx := range nodes.InnerNodes {
		c.values[nodes.InnerNodes[idx].Name] = nodes.InnerNodes[idx]
		c.nodes = append(c.nodes, nodes.InnerNodes[idx])
	}

	d.mu.Lock()
	d.data[cfg.Name] = &c
	d.mu.Unlock()

	return nil
}
