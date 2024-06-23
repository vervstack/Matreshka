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

	node, err := evon.MarshalEnv(&cfg)
	if err != nil {
		return errors.Wrap(err, "error marshalling config to variables")
	}
	c.nodes = node.InnerNodes
	c.values = evon.NodesToStorage(node.InnerNodes)

	d.mu.Lock()
	d.data[cfg.Name] = &c
	d.mu.Unlock()

	return nil
}
