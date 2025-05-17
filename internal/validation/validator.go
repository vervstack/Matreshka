package validation

import (
	"fmt"
	"strings"

	errors "go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/service/user_errors"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

type Validator struct {
	validConfigNameSymbols map[rune]struct{}
	validEnvNameSymbols    map[rune]struct{}
}

func New() Validator {
	v := Validator{
		validConfigNameSymbols: make(map[rune]struct{}),
		validEnvNameSymbols:    make(map[rune]struct{}),
	}

	for _, r := range []rune(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_`) {
		v.validConfigNameSymbols[r] = struct{}{}
		v.validEnvNameSymbols[r] = struct{}{}
	}

	for _, r := range []rune(`abcdefghijklmnopqrstuvwxyz`) {
		v.validConfigNameSymbols[r] = struct{}{}
	}

	for _, r := range []rune(`/{}`) {
		v.validEnvNameSymbols[r] = struct{}{}
	}

	return v
}

func (v Validator) IsConfigNameValid(name string) error {
	typePrefix := strings.Split(name, "_")[0]
	_, ok := api.ConfigTypePrefix_value[typePrefix]
	if !ok {
		return errors.Wrap(user_errors.ErrValidation,
			"Service name must start with a typed prefix. eg:"+
				fmt.Sprintf("pg_%[1]s, verv_%[1]s, minio_%[1]s, nginx_%[1]s", name))
	}

	if len(name) < 3 {
		return errors.Wrap(user_errors.ErrValidation,
			"Service name must be at least 3 symbols long")
	}

	invalidChars := make(map[rune]struct{})

	for _, r := range []rune(name) {
		if _, ok := v.validConfigNameSymbols[r]; !ok {
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
func (v Validator) IsEnvNameValid(envName string) error {
	if len(envName) < 3 {
		return errors.Wrap(user_errors.ErrValidation,
			"Variable name must be at least 3 symbols long")
	}

	invalidChars := make(map[rune]struct{})

	for _, r := range []rune(envName) {
		if _, ok := v.validEnvNameSymbols[r]; !ok {
			invalidChars[r] = struct{}{}
		}
	}

	if len(invalidChars) > 0 {
		sb := strings.Builder{}
		sb.WriteString("Variable name contains invalid character")
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
