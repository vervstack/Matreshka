package tests

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"go.redsock.ru/rerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"go.vervstack.ru/matreshka/internal/app"
	"go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func initApp() error {
	a, err := app.New()
	if err != nil {
		return rerrors.Wrap(err, "error initializing config")
	}

	_, err = a.Sqlite.Exec(`
		DELETE 
		FROM configs 	   
	    WHERE true;
		
		DELETE 
		FROM configs_values
		WHERE true;`)
	if err != nil {
		return rerrors.Wrap(err, "error db clean up")
	}

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

	conn, err := grpc.NewClient("[::]:"+a.Cfg.Servers.MASTER.Port,
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatalf("error connecting to test grpc server: %s ", err)
	}

	testEnv.matreshkaApi = matreshka_be_api.NewMatreshkaBeAPIClient(conn)

	ping, err := testEnv.matreshkaApi.ApiVersion(a.Ctx, &matreshka_be_api.ApiVersion_Request{})
	if err != nil {
		logrus.Fatalf("error pingin test server: %s", err)
	}

	if ping == nil {
		logrus.Fatalf("error pingin test server")
	}

	return nil
}
