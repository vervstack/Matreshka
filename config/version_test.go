package version

import (
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	devVersion = "dev"
)

func Test_VersionFromBuildInfo_Scenarios(t *testing.T) {
	cases := []struct {
		name      string
		buildInfo *debug.BuildInfo
		wanted    string
	}{
		{"matreshka's own binary", newBuildInfo(modulePath), devVersion},
		{"consumer resolved a release", newBuildInfo("go.vervstack.ru/Velez", newModule(modulePath, "v1.0.101")), "v1.0.101"},
		{"consumer replaced with a release", newReplacedBuildInfo("v1.0.102"), "v1.0.102"},
		{"consumer replaced with a local path", newReplacedBuildInfo(""), devVersion},
		{"matreshka not in the build", newBuildInfo("go.vervstack.ru/Velez"), devVersion},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.wanted, versionFromBuildInfo(tc.buildInfo))
		})
	}
}

func newBuildInfo(mainPath string, deps ...*debug.Module) *debug.BuildInfo {
	return &debug.BuildInfo{
		Main: debug.Module{Path: mainPath},
		Deps: deps,
	}
}

func newModule(path, version string) *debug.Module {
	return &debug.Module{
		Path:    path,
		Version: version,
	}
}

func newReplacedBuildInfo(replacementVersion string) *debug.BuildInfo {
	dep := newModule(modulePath, "v1.0.101")
	dep.Replace = newModule("../Matreshka", replacementVersion)

	return newBuildInfo("go.vervstack.ru/Velez", dep)
}
