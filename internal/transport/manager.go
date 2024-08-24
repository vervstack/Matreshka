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
	restListener net.Listener
	grpcListener net.Listener

	//	=
	mux       cmux.CMux
	serverMux *http.ServeMux
	// =
	grpc *grpc.Server
	http *http.Server
}

func NewManager(port string) (*ServersManager, error) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, errors.Wrap(err, "error opening listener")
	}

	mux := cmux.New(listener)

	serverMux := http.NewServeMux()

	s := &ServersManager{

		mux:          mux,
		grpcListener: mux.Match(cmux.HTTP2()),
		restListener: mux.Match(cmux.Any()),

		serverMux: serverMux,
		http: &http.Server{
			Handler: setUpCors().Handler(serverMux),
		},

		grpc: grpc.NewServer(),
	}

	return s, nil
}

func (m *ServersManager) AddGrpcServer(newGrpcService func(server *grpc.Server) (gateway http.Handler)) {
	gateway := newGrpcService(m.grpc)
	m.AddRestServer("/api/*", gateway)
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

	errGroup.Go(
		func() error {
			err := m.grpc.Serve(m.grpcListener)
			if err != nil {
				if !errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(err, "error listening grpc server")
				}
			}
			return nil
		})

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
