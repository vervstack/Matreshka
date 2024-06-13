package in_memory

import (
	_ "embed"
	"sync"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/data"
)

type inMemory struct {
	mu   sync.RWMutex
	data map[string]*Config
}

//go:embed full_config.yaml
var testCfg []byte

func New() data.Data {
	d := &inMemory{
		mu:   sync.RWMutex{},
		data: make(map[string]*Config),
	}

	c := matreshka.NewEmptyConfig()
	err := c.Unmarshal(testCfg)
	if err != nil {
		panic(err)
	}

	d.UpsertConfig(nil, c)

	return d
}
