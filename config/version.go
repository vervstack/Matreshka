package version

import (
	_ "embed"

	"go.vervstack.ru/matreshka"
)

//go:embed config.yaml
var masterConfig []byte
var version string

func init() {
	m, err := matreshka.ParseConfig(masterConfig)
	if err != nil {
		panic(err)
	}

	version = m.Version
}

func GetVersion() string {
	return version
}
