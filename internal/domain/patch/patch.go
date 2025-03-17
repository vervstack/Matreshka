package patch

import (
	"reflect"
	"strings"

	"go.redsock.ru/evon"

	"go.vervstack.ru/matreshka-be/internal/domain"
)

type Patch struct {
	Invalid   []domain.PatchConfig
	Upsert    []domain.PatchConfig
	EnvUpsert []domain.PatchConfig
	Delete    []domain.PatchConfig
}

func NewPatch(batch []domain.PatchConfig, oldConfig *evon.Node) Patch {
	p := Patch{}
	for _, ptch := range batch {
		var ok bool
		ptch.FieldName, ok = validateName(ptch)
		if !ok {
			p.Invalid = append(p.Invalid, ptch)
			continue
		}

		val := extractValue(ptch.FieldValue)
		if val == nil {
			p.Delete = append(p.Delete, ptch)
		} else {
			if strings.HasPrefix(ptch.FieldName, environmentSegment) {
				p.EnvUpsert = append(p.EnvUpsert, ptch)
			} else {
				p.Upsert = append(p.Upsert, ptch)
			}
		}
	}

	p.normalizeEnvironmentChanges(oldConfig)
	return p
}

func (p *Patch) NameChanged() (newName string, nameChanged bool) {
	for _, v := range p.Upsert {
		if v.FieldName == applicationName && v.FieldValue != nil {
			return *v.FieldValue, true
		}
	}

	return "", false
}

func (p *Patch) normalizeEnvironmentChanges(cfg *evon.Node) {
	nodeStorage := evon.NodesToStorage(cfg.InnerNodes)

	newEnvValues := make(map[string]domain.PatchConfig)
	typesMap := make(map[string]domain.PatchConfig)
	enumMap := make(map[string]domain.PatchConfig)

	envUpsert := make([]domain.PatchConfig, 0, len(newEnvValues))

	for _, valuePatch := range p.EnvUpsert {
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
			p.Invalid = append(p.Invalid, patchVal)
			continue
		}

		envUpsert = append(envUpsert, patchVal, typeVal)
		enumVal, ok := enumMap[key]
		if ok {
			envUpsert = append(envUpsert, enumVal)
		}
	}

	p.EnvUpsert = envUpsert
}

func extractValue(in any) any {
	inRef := reflect.ValueOf(in)
	if inRef.IsNil() {
		return nil
	}

	if inRef.Kind() != reflect.Ptr {
		return inRef.Interface()
	}

	return inRef.Elem().Interface()
}

func validateName(patch domain.PatchConfig) (newName string, ok bool) {
	for _, segment := range allowedSegments {
		if strings.HasPrefix(patch.FieldName, segment) {
			return strings.ToUpper(patch.FieldName), true
		}
	}

	return patch.FieldName, false
}
