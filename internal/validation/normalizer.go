package validation

import (
	"strings"
)

func (v Validator) normalizeAndValidateEnvName(envName string) (string, error) {
	envName = strings.ToUpper(envName)
	return envName, v.IsEnvNameValid(envName)
}
