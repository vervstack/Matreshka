package app

import (
	"go.vervstack.ru/matreshka/internal/app"
)

type (
	App       = app.App
	CustomApp = app.Custom
)

// New - application basic constructor for
// e2e tests and external service launch
func New() (App, error) {
	return app.New()
}
