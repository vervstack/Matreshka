package servicev1

import (
	errors "github.com/Red-Sock/trace-errors"
	"google.golang.org/grpc/codes"

	"github.com/godverv/matreshka-be/internal/data"
)

var ErrInvalidPatchName = errors.New("invalid patched env var name", codes.InvalidArgument)

var allowedSegments = []string{
	appInfo,
	environmentSegment,
	dataSourceSegment,
	serverSegment,
}

type ConfigService struct {
	data data.Data
}

func New(data data.Data) *ConfigService {
	return &ConfigService{
		data: data,
	}
}
