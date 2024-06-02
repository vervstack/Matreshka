package grpc

import (
	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type App struct {
	version    string
	storage    data.Data
	cfgService service.ConfigService

	matreshka_api.UnimplementedMatreshkaBeAPIServer
}
