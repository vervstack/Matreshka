package tests

import (
	"context"
	_ "embed"
	"os"
	"testing"
	"time"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/matreshka-be/internal/app"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type Env struct {
	app        app.App
	grpcClient matreshka_api.MatreshkaBeAPIClient
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
		return errors.Wrap(err, "error during app init")
	}

	_, err = a.DbConn.Exec(`
		DELETE FROM configs;
		DELETE FROM configs_values;`)
	if err != nil {
		return errors.Wrap(err, "error db clean up")
	}

	testEnv.app = a
	go func() {
		err = a.Start()
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	err = initClient()
	if err != nil {
		return errors.Wrap(err, "error during app init")
	}

	return nil
}

func initClient() error {
	cl, err := grpc.NewClient("0.0.0.0:999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "error creating new client")
	}

	testEnv.grpcClient = matreshka_api.NewMatreshkaBeAPIClient(cl)

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

func (e *Env) create(t *testing.T, serviceName string, config []byte) {
	createReq := &matreshka_api.PostConfig_Request{
		Content:     config,
		ServiceName: serviceName,
	}
	ctx := context.Background()

	postResp, err := testEnv.grpcClient.PostConfig(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, postResp)
}

func (e *Env) get(t *testing.T, serviceName string) matreshka.AppConfig {
	ctx := context.Background()
	getReq := &matreshka_api.GetConfig_Request{
		ServiceName: serviceName,
	}
	getResp, err := testEnv.grpcClient.GetConfig(ctx, getReq)
	require.NoError(t, err)

	readConfig := matreshka.NewEmptyConfig()
	err = readConfig.Unmarshal(getResp.Config)
	require.NoError(t, err)

	return readConfig
}
