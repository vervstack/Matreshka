package converters

import (
	"time"

	"github.com/godverv/matreshka"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func ToProtoAppInfo(info matreshka.AppInfo) *matreshka_api.Config_AppConfig {
	return &matreshka_api.Config_AppConfig{
		Name:               info.Name,
		Version:            info.Version,
		StartupDurationSec: uint32(info.StartupDuration.Seconds()),
	}
}

func FromProtoAppInfo(info *matreshka_api.Config_AppConfig) (out matreshka.AppInfo) {
	return matreshka.AppInfo{
		Name:            info.GetName(),
		Version:         info.GetVersion(),
		StartupDuration: time.Second * time.Duration(info.GetStartupDurationSec()),
	}
}
