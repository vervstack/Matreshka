package middleware

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LogInterceptor() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			fields := logrus.Fields{
				"method":  info.FullMethod,
				"request": req,
			}

			defer func() {
				logrus.WithFields(fields).
					Debug("GRPC request:")
			}()

			resp, err = handler(ctx, req)
			if err != nil {
				fields["error"] = err
			} else {
				fields["response"] = resp
			}

			return resp, err
		})
}
