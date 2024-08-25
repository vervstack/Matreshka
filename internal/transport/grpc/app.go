package grpc

import (
	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

type Impl struct {
	version string
	service service.ConfigService
	storage data.Data

	matreshka_be_api.UnimplementedMatreshkaBeAPIServer
}
