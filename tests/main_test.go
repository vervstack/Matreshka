package tests

import (
	"context"
	_ "embed"
	"os"
	"testing"

	"github.com/godverv/matreshka"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	errors "go.redsock.ru/rerrors"

	"github.com/godverv/matreshka-be/internal/app"
	"github.com/godverv/matreshka-be/internal/transport/grpc_impl"
	"github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

type Env struct {
	grpcApi *grpc_impl.Impl
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

func initApp() error {
	a, err := app.New()
	if err != nil {
		return errors.Wrap(err, "error initializing config")
	}

	_, err = a.Sqlite.Exec(`
		DELETE 
		FROM configs 	   
	    WHERE true;
		
		DELETE 
		FROM configs_values
		WHERE true;`)
	if err != nil {
		return errors.Wrap(err, "error db clean up")
	}

	testEnv.grpcApi = a.Custom.GrpcImpl

	return nil
}

func getFullConfig(t *testing.T) matreshka.AppConfig {
	fullConfig := matreshka.NewEmptyConfig()
	err := fullConfig.Unmarshal(fullConfigBytes)
	if err != nil {
		t.Fatal(errors.Wrap(err, "error during unmarshalling full config"))
	}

	return fullConfig
}

func (e *Env) create(t *testing.T, serviceName string) {
	createReq := &matreshka_be_api.CreateConfig_Request{
		ServiceName: serviceName,
	}
	ctx := context.Background()

	postResp, err := testEnv.grpcApi.CreateConfig(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, postResp)
}

func (e *Env) get(t *testing.T, serviceName string) matreshka.AppConfig {
	ctx := context.Background()
	getReq := &matreshka_be_api.GetConfig_Request{
		ServiceName: serviceName,
	}
	getResp, err := testEnv.grpcApi.GetConfig(ctx, getReq)
	require.NoError(t, err)

	readConfig := matreshka.NewEmptyConfig()
	err = readConfig.Unmarshal(getResp.Config)
	require.NoError(t, err)

	return readConfig
}
