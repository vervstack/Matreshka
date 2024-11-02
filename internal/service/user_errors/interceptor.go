package user_errors

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp any, err error) {
		resp, err = handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		logrus.Error(err)

		return resp, err
	}
}
