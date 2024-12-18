package servicev1

import (
	"strings"

	errors "go.redsock.ru/rerrors"

	"go.verv.tech/matreshka-be/internal/service/user_errors"
)

type validator struct {
	validServiceNameSymbols map[rune]struct{}
}

func newValidator() validator {
	v := validator{
		validServiceNameSymbols: map[rune]struct{}{},
	}

	for _, r := range []rune(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_`) {
		v.validServiceNameSymbols[r] = struct{}{}
	}

	return v
}

func (v *validator) validateServiceName(name string) error {
	if len(name) < 3 {
		return errors.Wrap(user_errors.ErrValidation,
			"Service name must be at least 3 symbols long")
	}

	invalidChars := make(map[rune]struct{})

	for _, r := range []rune(name) {
		if _, ok := v.validServiceNameSymbols[r]; !ok {
			invalidChars[r] = struct{}{}
		}
	}

	if len(invalidChars) > 0 {
		sb := strings.Builder{}
		sb.WriteString("Name contains invalid character")
		if len(invalidChars) > 1 {
			sb.WriteString("s")
		}

		sb.WriteString(": ")

		idx := 0
		for r := range invalidChars {
			sb.WriteRune(r)
			if idx < len(invalidChars)-1 {
				sb.WriteRune(',')
			}

			idx++
		}

		return errors.Wrap(user_errors.ErrValidation, sb.String())
	}

	return nil
}
