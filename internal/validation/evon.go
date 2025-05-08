package validation

import (
	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (v Validator) AsEvon(patch domain.PatchConfigRequest) error {
	for _, p := range patch.Update {
		err := v.IsEnvNameValid(p.FieldName)
		if err != nil {
			return rerrors.Wrap(err)
		}
	}

	for _, p := range patch.RenameTo {
		err := v.IsEnvNameValid(p.OldName)
		if err != nil {
			return rerrors.Wrap(err)
		}

		err = v.IsEnvNameValid(p.NewName)
		if err != nil {
			return rerrors.Wrap(err)
		}
	}

	for _, fieldName := range patch.Delete {
		err := v.IsEnvNameValid(fieldName)
		if err != nil {
			return rerrors.Wrap(err)
		}
	}
	return nil
}
