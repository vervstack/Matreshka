package validation

import (
	"strings"

	"go.redsock.ru/evon"

	"go.vervstack.ru/matreshka/internal/domain"
)

const (
	environmentSegment = "ENVIRONMENT"
)

type VervConfigValidationResult struct {
	EnvUpsert []domain.PatchUpdate
	Upsert    []domain.PatchUpdate

	Invalid []domain.PatchUpdate
}

func (v Validator) AsVerv(original *evon.Node, patch *domain.PatchConfigRequest) VervConfigValidationResult {
	vervV := newVervConfigValidator(*patch)

	vervV.parseEnvironmentChanges(original)

	return vervV
}

func newVervConfigValidator(patch domain.PatchConfigRequest) VervConfigValidationResult {
	vervV := VervConfigValidationResult{}

	for _, p := range patch.Update {
		if strings.HasPrefix(p.FieldName, environmentSegment) {
			vervV.EnvUpsert = append(vervV.EnvUpsert, p)
		} else {
			vervV.Upsert = append(vervV.Upsert, p)
		}
	}

	return vervV
}

func (v *VervConfigValidationResult) parseEnvironmentChanges(original *evon.Node) {
	nodeStorage := evon.NodesToStorage(original)

	newEnvValues := make(map[string]domain.PatchUpdate)
	typesMap := make(map[string]domain.PatchUpdate)
	enumMap := make(map[string]domain.PatchUpdate)

	envUpsert := make([]domain.PatchUpdate, 0, len(newEnvValues))

	for _, valuePatch := range v.EnvUpsert {
		// already exists -> simply update value
		_, ok := nodeStorage[valuePatch.FieldName]
		if ok {
			envUpsert = append(envUpsert, valuePatch)
			continue
		}

		if strings.HasSuffix(valuePatch.FieldName, "_TYPE") {
			typesMap[valuePatch.FieldName[:len(valuePatch.FieldName)-5]] = valuePatch
			continue
		}

		if strings.HasSuffix(valuePatch.FieldName, "_ENUM") {
			enumMap[valuePatch.FieldName[:len(valuePatch.FieldName)-5]] = valuePatch
			continue
		}

		newEnvValues[valuePatch.FieldName] = valuePatch
	}

	for key, patchVal := range newEnvValues {
		typeVal, ok := typesMap[key]
		if !ok {
			v.Invalid = append(v.Invalid, patchVal)
			continue
		}

		envUpsert = append(envUpsert, patchVal, typeVal)
		enumVal, ok := enumMap[key]
		if ok {
			envUpsert = append(envUpsert, enumVal)
		}
	}

	v.EnvUpsert = envUpsert
}
