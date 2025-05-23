package tests

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/pkg/matreshka"
	"go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

type Env struct {
	matreshkaApi matreshka_be_api.MatreshkaBeAPIClient
}

//go:embed config/test.config.yaml
var fullConfigBytes []byte

var testEnv Env

func TestMain(m *testing.M) {
	err := initApp()
	if err != nil {
		logrus.Fatal(err)
	}

	var code int
	code = m.Run()
	os.Exit(code)
}

func (e *Env) create(t *testing.T, configName string) {
	createReq := &matreshka_be_api.CreateConfig_Request{
		ConfigName: configName,
	}
	ctx := context.Background()

	postResp, err := testEnv.matreshkaApi.CreateConfig(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, postResp)
}

func (e *Env) updateConfigValues(t *testing.T, cfg matreshka.AppConfig) {
	req := &matreshka_be_api.PatchConfig_Request{
		ConfigName: cfg.ModuleName(),
	}

	nodes, err := evon.MarshalEnv(&cfg)
	require.NoError(t, err)

	storage := evon.NodesToStorage(nodes)

	for k, v := range storage {
		if v.Value != nil {
			req.Patches = append(req.Patches,
				&matreshka_be_api.PatchConfig_Patch{
					FieldName: k,
					Patch: &matreshka_be_api.PatchConfig_Patch_UpdateValue{
						UpdateValue: fmt.Sprint(v.Value),
					},
				})
		}
	}

	ctx := context.Background()

	_, err = e.matreshkaApi.PatchConfig(ctx, req)
	require.NoError(t, err)
}

func (e *Env) get(t *testing.T, configName string) matreshka.AppConfig {
	ctx := context.Background()
	getReq := &matreshka_be_api.GetConfig_Request{
		ConfigName: configName,
	}
	getResp, err := testEnv.matreshkaApi.GetConfig(ctx, getReq)
	require.NoError(t, err)

	readConfig := matreshka.NewEmptyConfig()
	err = readConfig.Unmarshal(getResp.Config)
	require.NoError(t, err)

	return readConfig
}

func getFullConfig(t *testing.T) matreshka.AppConfig {
	fullConfig := matreshka.NewEmptyConfig()
	err := fullConfig.Unmarshal(fullConfigBytes)
	if err != nil {
		t.Fatal(errors.Wrap(err, "error during unmarshalling full config"))
	}

	fullConfig.Name = getServiceNameFromTest(t)

	return fullConfig
}

func getServiceNameFromTest(t *testing.T) string {
	return strings.ReplaceAll(t.Name(), "/", "_")
}
