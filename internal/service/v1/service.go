package v1

import (
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/data"
)

var ErrInvalidPatchName = errors.New("invalid patch name")

type ConfigService struct {
	data data.Data

	allowedSegments []string
}

const (
	environmentSegment = "ENVIRONMENT"
	dataSourceSegment  = "SERVERS"
	serverSegment      = "DATA_SOURCES"
)

func New(data data.Data) *ConfigService {
	return &ConfigService{
		data: data,
		allowedSegments: []string{
			environmentSegment,
			dataSourceSegment,
			serverSegment,
		},
	}
}
