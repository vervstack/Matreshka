package matreshka_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"go.vervstack.ru/matreshka/internal/web/auth"
)

const AuthHeader = auth.Header

func WithHeader(headerValue string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Add custom header
		md := metadata.Pairs(AuthHeader, headerValue)
		ctx = metadata.NewOutgoingContext(ctx, md)

		// Call the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
