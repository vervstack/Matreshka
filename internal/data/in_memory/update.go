package in_memory

import (
	"context"

	"github.com/Red-Sock/evon"
	"github.com/godverv/matreshka"
)

func (d *inMemory) UpsertConfig(_ context.Context, cfg matreshka.AppConfig) error {

	// TODO now it's total update. Thinking in advance to merge two configs before inserting
	splited := evon.MarshalEnv(&cfg)
	c := config{
		appConfig: cfg,
		values:    map[string]interface{}{},
	}
	for _, s := range splited {
		c.values[s.Name] = s.Value
	}

	d.mu.Lock()
	d.data[cfg.Name] = &c
	d.mu.Unlock()

	return nil
}
