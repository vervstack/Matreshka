package v1

import (
	"context"
	"fmt"
	"strings"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/data"
	"github.com/godverv/matreshka-be/internal/domain"
)

var ErrInvalidPatchName = errors.New("invalid patch name")

type ConfigService struct {
	data data.Data

	allowedSegments []string
}

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

const (
	environmentSegment = "ENVIRONMENT"
	dataSourceSegment  = "SERVERS"
	serverSegment      = "DATA_SOURCES"
)

func (c *ConfigService) PatchConfig(ctx context.Context, configPatch domain.PatchConfigRequest) error {
	invalidPatchName := make([]string, 0)
	validBatches := make([]domain.PatchConfig, 0, len(configPatch.Batch))

	for patchIdx := range configPatch.Batch {
		configPatch.Batch[patchIdx].FieldPath = strings.ToUpper(configPatch.Batch[patchIdx].FieldPath)

		var hit bool
		for _, segment := range c.allowedSegments {
			hit = strings.HasPrefix(configPatch.Batch[patchIdx].FieldPath, segment)
			if hit {
				break
			}
		}
		if !hit {
			invalidPatchName = append(invalidPatchName, configPatch.Batch[patchIdx].FieldPath)
		} else {
			validBatches = append(validBatches, configPatch.Batch[patchIdx])
		}
	}

	dataReq := domain.PatchConfigRequest{
		ServiceName: configPatch.ServiceName,
		Batch:       validBatches,
	}

	err := c.data.PatchConfig(ctx, dataReq)
	if err != nil {
		return errors.Wrap(err, "error patching config in data storage")
	}

	if len(invalidPatchName) != 0 {
		return errors.Wrap(ErrInvalidPatchName, fmt.Sprint(invalidPatchName))
	}

	return nil
}
