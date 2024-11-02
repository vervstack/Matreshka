package grpc_impl

import (
	"github.com/godverv/matreshka-be/internal/service"
	"github.com/godverv/matreshka-be/internal/storage"
	"github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

type Impl struct {
	version string
	service service.ConfigService
	storage storage.Data

	matreshka_be_api.UnimplementedMatreshkaBeAPIServer
}
