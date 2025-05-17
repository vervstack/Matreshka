package Matreshka

import (
	"go.vervstack.ru/matreshka/pkg/matreshka"
)

type AppConfig matreshka.AppConfig
type AppInfo matreshka.AppInfo

type DataSources matreshka.DataSources
type Environment matreshka.Environment

var NewEmptyConfig = matreshka.NewEmptyConfig
var ReadConfigs = matreshka.ReadConfigs
var ParseConfig = matreshka.ParseConfig
var MergeConfigs = matreshka.MergeConfigs

type Servers matreshka.Servers

var ServerName = matreshka.ServerName

var ServiceDiscovery matreshka.ServiceDiscovery
