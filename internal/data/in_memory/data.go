package in_memory

import (
	_ "embed"
	"sync"

	"github.com/godverv/matreshka-be/internal/data"
)

type inMemory struct {
	mu   sync.RWMutex
	data map[string]*Config
}

func New() data.Data {
	return &inMemory{
		mu:   sync.RWMutex{},
		data: make(map[string]*Config),
	}
}
