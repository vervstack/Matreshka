package in_memory

import (
	"sync"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/internal/data"
)

type inMemory struct {
	m  sync.RWMutex
	mp map[string]*matreshka.AppConfig
}

func New() data.Data {
	return &inMemory{
		m:  sync.RWMutex{},
		mp: make(map[string]*matreshka.AppConfig),
	}
}
