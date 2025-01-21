package tests

import (
	"context"
	_ "embed"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	errors "go.redsock.ru/rerrors"
	"go.verv.tech/matreshka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"go.verv.tech/matreshka-be/internal/app"
	"go.verv.tech/matreshka-be/internal/transport/grpc_impl"
	"go.verv.tech/matreshka-be/pkg/matreshka_be_api"
)

type Env struct {
	//deprecated use grpcAPI
	grpcImpl *grpc_impl.Impl
	grpcApi  matreshka_be_api.MatreshkaBeAPIClient
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

	testEnv.grpcImpl = a.Custom.GrpcImpl

	const bufSize = 1024 * 1024
	lis := bufconn.Listen(bufSize)

	serv := grpc.NewServer()
	matreshka_be_api.RegisterMatreshkaBeAPIServer(serv, a.Custom.GrpcImpl)
	go func() {
		if err := serv.Serve(lis); err != nil {
			logrus.Fatalf("error serving grpc server for tests %s", err)
		}
	}()

	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatalf("error connecting to test grpc server: %s ", err)
	}
	testEnv.grpcApi = matreshka_be_api.NewMatreshkaBeAPIClient(conn)

	ping, err := testEnv.grpcApi.ApiVersion(a.Ctx, &matreshka_be_api.ApiVersion_Request{})
	if err != nil {
		logrus.Fatalf("error pingin test server: %s", err)
	}

	if ping == nil {
		logrus.Fatalf("error pingin test server")
	}

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

	postResp, err := testEnv.grpcImpl.CreateConfig(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, postResp)
}

func (e *Env) get(t *testing.T, serviceName string) matreshka.AppConfig {
	ctx := context.Background()
	getReq := &matreshka_be_api.GetConfig_Request{
		ServiceName: serviceName,
	}
	getResp, err := testEnv.grpcImpl.GetConfig(ctx, getReq)
	require.NoError(t, err)

	readConfig := matreshka.NewEmptyConfig()
	err = readConfig.Unmarshal(getResp.Config)
	require.NoError(t, err)

	return readConfig
}

func getServiceNameFromTest(t *testing.T) string {
	return strings.ReplaceAll(t.Name(), "/", "_")
}
