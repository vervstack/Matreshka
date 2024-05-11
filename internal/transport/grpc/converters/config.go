package converters

import (
	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func ToProtoConfig(cfg matreshka.AppConfig) *matreshka_api.Config {
	return &matreshka_api.Config{
		AppConfig:   ToProtoAppInfo(cfg.AppInfo),
		Resources:   ToProtoResources(cfg.Resources),
		Api:         ToProtoApi(cfg.Servers),
		Environment: ToProtoEnvironment(cfg.Environment),
	}
}

func FromProtoConfig(cfg *matreshka_api.Config) matreshka.AppConfig {
	return matreshka.AppConfig{
		AppInfo:     FromProtoAppInfo(cfg.AppConfig),
		Resources:   FromProtoResources(cfg.Resources),
		Servers:     FromProtoApi(cfg.Api),
		Environment: FromProtoEnvironment(cfg.Environment),
	}
}
