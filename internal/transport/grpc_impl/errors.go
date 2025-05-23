package grpc_impl

import (
	"fmt"

	"go.redsock.ru/rerrors"
	"google.golang.org/grpc/codes"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

var (
	errNoPrefix = rerrors.NewUserError(
		fmt.Sprintf("Expected to have type prefix: %v ", api.ConfigTypePrefix_name),
		codes.InvalidArgument)
)
