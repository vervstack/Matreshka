package service

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"google.golang.org/grpc/codes"

	"github.com/godverv/matreshka-be/internal/domain"
)

var ErrConfigExists = errors.New("config already exist", codes.AlreadyExists)

type ConfigService interface {
	PatchConfig(ctx context.Context, configPatch domain.PatchConfigRequest) error
	CreateConfig(ctx context.Context, serviceName string, cfg matreshka.AppConfig) (*Response, error)
}

type Response struct {
	UserError string
	HTTPCode  int
}
