package in_memory

import (
	"context"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/data"
)

func (d *inMemory) SaveConfig(_ context.Context, cfg matreshka.AppConfig) error {
	c := data.Config{}
	c.Values = map[string]*evon.Node{}
	c.Cfg = &cfg

	node, err := evon.MarshalEnv(&cfg)
	if err != nil {
		return errors.Wrap(err, "error marshalling config to variables")
	}
	c.Nodes = node.InnerNodes
	c.Values = evon.NodesToStorage(node.InnerNodes)

	d.mu.Lock()
	d.data[cfg.Name] = &c
	d.mu.Unlock()

	return nil
}
