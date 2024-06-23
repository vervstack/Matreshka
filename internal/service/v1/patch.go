package v1

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka-be/internal/domain"
)

const (
	appInfo            = "APP-INFO"
	environmentSegment = "ENVIRONMENT"
	dataSourceSegment  = "SERVERS"
	serverSegment      = "DATA-SOURCES"
)

func (c *ConfigService) PatchConfig(ctx context.Context, configPatch domain.PatchConfigRequest) error {
	invalidPatchName := make([]string, 0)
	batchesToUpsert := make([]domain.PatchConfig, 0, len(configPatch.Batch))
	batchesToDelete := make([]domain.PatchConfig, 0)

	for patchIdx := range configPatch.Batch {
		configPatch.Batch[patchIdx].FieldName = strings.ToUpper(configPatch.Batch[patchIdx].FieldName)

		var hit bool
		for _, segment := range c.allowedSegments {
			hit = strings.HasPrefix(configPatch.Batch[patchIdx].FieldName, segment)
			if hit {
				break
			}
		}
		if !hit {
			invalidPatchName = append(invalidPatchName, configPatch.Batch[patchIdx].FieldName)
			continue
		}

		patch := domain.PatchConfig{
			FieldName: configPatch.Batch[patchIdx].FieldName,
		}
		var upsert, del bool
		patch.FieldValue, upsert, del = extractUpdateValue(configPatch.Batch[patchIdx].FieldValue)
		if upsert {
			batchesToUpsert = append(batchesToUpsert, patch)
		} else if del {
			batchesToDelete = append(batchesToDelete, patch)
		}
	}

	dataReq := domain.PatchConfigRequest{
		ServiceName: configPatch.ServiceName,
		Batch:       batchesToUpsert,
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

func extractUpdateValue(in any) (value any, toUpsert, toDelete bool) {
	inRef := reflect.ValueOf(in)
	if inRef.Kind() != reflect.Ptr {
		return inRef.Interface(), true, false
	}

	if inRef.IsNil() {
		return nil, false, true
	}

	return inRef.Elem().Interface(), true, false
}
