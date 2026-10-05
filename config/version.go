package version

import (
	"runtime/debug"

	buildversion "go.vervstack.ru/matreshka/internal/version"
)

const (
	modulePath = "go.vervstack.ru/matreshka"
)

// GetVersion returns the matreshka release this code belongs to. Imported as a
// library it is the module version the consumer's go.mod resolved; inside
// matreshka's own binary it is the tag baked in at build time.
func GetVersion() string {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return buildversion.Get()
	}

	return versionFromBuildInfo(buildInfo)
}

func versionFromBuildInfo(buildInfo *debug.BuildInfo) string {
	if buildInfo.Main.Path == modulePath {
		return buildversion.Get()
	}

	for _, dep := range buildInfo.Deps {
		if dep.Path != modulePath {
			continue
		}

		module := dep
		if dep.Replace != nil {
			module = dep.Replace
		}

		if module.Version != "" {
			return module.Version
		}
	}

	return buildversion.Get()
}
