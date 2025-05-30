// Code generated by RedSock CLI. DO NOT EDIT.

package config

import (
	"flag"

	"go.redsock.ru/rerrors"
	"go.vervstack.ru/matreshka/pkg/matreshka"
)

var ErrAlreadyLoaded = rerrors.New("config already loaded")

type Config struct {
	AppInfo matreshka.AppInfo

	Servers     ServersConfig
	DataSources DataSourcesConfig
	Environment EnvironmentConfig
	Overrides   matreshka.ServiceDiscovery
}

var defaultConfig Config

const (
	devConfigPath  = "./config/dev.yaml"
	prodConfigPath = "./config/config.yaml"
)

func Load() (Config, error) {
	if defaultConfig.AppInfo.Name != "" {
		return defaultConfig, ErrAlreadyLoaded
	}

	var cfgPath string
	var isDevBuild bool

	flag.StringVar(&cfgPath, "config", "", "Path to configuration file")
	flag.BoolVar(&isDevBuild, "dev", false, "Flag turns on a dev config at ./config/dev.yaml")
	flag.Parse()

	if cfgPath == "" {
		if isDevBuild {
			cfgPath = devConfigPath
		} else {
			cfgPath = prodConfigPath
		}
	}

	rootConfig, err := matreshka.ReadConfigs(cfgPath)
	if err != nil {
		return defaultConfig, rerrors.Wrap(err, "error reading matreshka config")
	}

	defaultConfig.AppInfo = rootConfig.AppInfo
	defaultConfig.Overrides = rootConfig.ServiceDiscovery

	err = rootConfig.Servers.ParseToStruct(&defaultConfig.Servers)
	if err != nil {
		return defaultConfig, rerrors.Wrap(err, "Error parsing servers to config")
	}
	err = rootConfig.DataSources.ParseToStruct(&defaultConfig.DataSources)
	if err != nil {
		return defaultConfig, rerrors.Wrap(err, "error parsing data sources to struct")
	}
	err = rootConfig.Environment.ParseToStruct(&defaultConfig.Environment)
	if err != nil {
		return defaultConfig, rerrors.Wrap(err, "error parsing environment config")
	}

	return defaultConfig, nil
}
