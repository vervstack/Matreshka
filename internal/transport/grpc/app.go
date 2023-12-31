package grpc

import (
	"github.com/godverv/matreshka-be/pkg/api/matreshka_api"
)

type App struct {
	matreshka_api.UnimplementedMatreshkaBeAPIServer
}
