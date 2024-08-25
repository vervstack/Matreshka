package transport

import (
	"context"
	"net"
	"net/http"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
)

type ServersManager struct {
	ctx context.Context

	mux cmux.CMux

	grpcServer
	httpServer
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
		httpServer: newHttpServer(mux.Match(cmux.Any()), serverMux),
	}

	return s, nil
}

func (m *ServersManager) Start(ctx context.Context) error {
	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		// TODO Проверить не возвращается ли из функции ошибка закрытия сервера как в других обработчиках
		err := m.mux.Serve()
		if err != nil {
			return errors.Wrap(err, "error serving main mux")
		}
		return nil
	})

	errGroup.Go(m.grpcServer.start)
	errGroup.Go(m.httpServer.start)

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
	eg, _ := errgroup.WithContext(m.ctx)

	eg.Go(m.grpcServer.stop)
	eg.Go(m.httpServer.stop)

	err := eg.Wait()
	if err != nil {
		return errors.Wrap(err, "error stopping server")
	}

	return nil
}
