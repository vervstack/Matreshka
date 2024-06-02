package in_memory

import (
	_ "embed"
	"sync"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/data"
)

type inMemory struct {
	m  sync.RWMutex
	mp map[string]*matreshka.AppConfig
}

//go:embed full_config.yaml
var testCfg []byte

func New() data.Data {
	d := &inMemory{
		m:  sync.RWMutex{},
		mp: make(map[string]*matreshka.AppConfig),
	}

	c := matreshka.NewEmptyConfig()
	err := c.Unmarshal(testCfg)
	if err != nil {
		panic(err)
	}

	d.UpsertConfig(nil, c)

	return d
}
