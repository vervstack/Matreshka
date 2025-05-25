package matreshka_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"go.vervstack.ru/matreshka/internal/web/auth"
)

const AuthHeader = auth.Header

type AuthType string

const (
	Pass AuthType = "Pass"
)

func WithHeader(authType AuthType, headerValue string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Add custom header
		md := metadata.Pairs(AuthHeader, string(authType)+" "+headerValue)
		ctx = metadata.NewOutgoingContext(ctx, md)

		// Call the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
