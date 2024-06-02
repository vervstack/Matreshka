package v1

import (
	"github.com/godverv/matreshka-be/internal/service"
)

type ConfigService struct {
}

func New() service.ConfigService {
	return &ConfigService{}
}
