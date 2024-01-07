// Code generated by RedSock CLI. DO NOT EDIT

package grpc

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/godverv/matreshka/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/matreshka-be/internal/config"
	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type Server struct {
	grpcServer *grpc.Server
	gwServer   *http.Server

	grpcAddress string
	gwAddress   string
	m           sync.Mutex
}

func NewServer(cfg config.Config, server *api.GRPC, storage data.Data) (*Server, error) {
	srv := grpc.NewServer()

	// Register your servers here
	matreshka_api.RegisterMatreshkaBeAPIServer(srv, &App{
		version: cfg.AppInfo().Version,
		storage: storage,
	})

	return &Server{
		grpcServer: srv,

		grpcAddress: ":" + server.GetPortStr(),
		gwAddress:   ":" + server.GetGatewayPortStr(),
	}, nil
}

func (s *Server) Start(_ context.Context) error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.grpcAddress != ":" {
		lis, err := net.Listen("tcp", s.grpcAddress)
		if err != nil {
			return errors.Wrapf(err, "error when tried to listen on %s", s.grpcAddress)
		}

		go s.startGrpcServer(lis)
	} else {
		logrus.Warn("no grpc port specified")
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{}))

	if s.gwAddress != ":" {
		err := matreshka_api.RegisterMatreshkaBeAPIHandlerFromEndpoint(
			context.TODO(),
			mux,
			s.grpcAddress,
			[]grpc.DialOption{
				grpc.WithBlock(),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			})
		if err != nil {
			logrus.Errorf("error registering grpc2http handler: %s", err)
		}
		s.gwServer = &http.Server{
			Addr:    s.gwAddress,
			Handler: mux,
		}

		go s.startGrpcGwServer()
	} else {
		logrus.Warn("no grpc gateway port specified")
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	logrus.Infof("Stopping GRPC-GW server at %s", s.gwAddress)
	err := s.gwServer.Shutdown(ctx)
	if err != nil {
		logrus.Errorf("error shutting down grpc-gw server at %s", s.gwAddress)
	}

	logrus.Infof("Stopping GRPC server at %s", s.grpcAddress)
	s.grpcServer.GracefulStop()
	logrus.Infof("GRPC server at %s is stopped", s.grpcAddress)

	return err
}

func (s *Server) startGrpcServer(lis net.Listener) {
	logrus.Infof("Starting GRPC Server at %s (%s)", s.grpcAddress, "tcp")
	err := s.grpcServer.Serve(lis)
	if err != nil {
		logrus.Errorf("error serving grpc: %s", err)
	} else {
		logrus.Infof("GRPC Server at %s is Stopped", s.grpcAddress)
	}
}

func (s *Server) startGrpcGwServer() {
	logrus.Infof("Starting HTTP Server at %s", s.gwAddress)
	err := s.gwServer.ListenAndServe()
	if err != nil {
		logrus.Errorf("error starting grpc2http handler: %s", err)
	}
}
