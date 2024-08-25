package transport

import (
	"context"
	"net"
	"net/http"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type ServersManager struct {
	ctx context.Context

	mux cmux.CMux

	grpcServer
	// Implementations
	// RPC
	grpc         *grpc.Server
	grpcListener net.Listener

	// Http
	http         *http.Server
	serverMux    *http.ServeMux
	restListener net.Listener
}

func NewManager(ctx context.Context, port string) (*ServersManager, error) {
	if port[0] != ':' {
		port = ":" + port
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, errors.Wrap(err, "error opening listener")
	}

	mux := cmux.New(listener)

	serverMux := http.NewServeMux()

	s := &ServersManager{
		ctx: ctx,
		mux: mux,

		grpcServer: newGrpcServer(ctx, mux.Match(cmux.HTTP2()), serverMux),

		restListener: mux.Match(cmux.Any()),

		serverMux: serverMux,
		http: &http.Server{
			Handler: setUpCors().Handler(serverMux),
		},
	}

	return s, nil
}

func (m *ServersManager) AddRestServer(path string, handler http.Handler) {
	m.serverMux.Handle(path, handler)
}

func (m *ServersManager) Start(ctx context.Context) error {
	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		err := m.mux.Serve()
		if err != nil {
			return errors.Wrap(err, "error serving main mux")
		}
		return nil
	})

	errGroup.Go(m.grpcServer.Start)

	errGroup.Go(func() error {
		err := m.http.Serve(m.restListener)
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				return errors.Wrap(err, "error listening http server")
			}
		}
		return nil
	})

	errC := make(chan error, 1)
	go func() {
		errC <- errGroup.Wait()
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-errC:
		return errors.Wrap(err)
	}
}

func (m *ServersManager) Stop() error {
	m.grpc.GracefulStop()

	return nil
}
