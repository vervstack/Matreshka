package converters

import (
	"fmt"
)

func ToProtoEnvironment(in map[string]interface{}) map[string]string {
	out := make(map[string]string)

	for k, v := range in {
		out[k] = fmt.Sprintf("%v", v)
	}

	return out
}

func FromProtoEnvironment(in map[string]string) map[string]interface{} {
	out := make(map[string]interface{})

	for k, v := range in {
		out[k] = v
	}

	return out
}
