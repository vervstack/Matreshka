package main

import (
	"github.com/sirupsen/logrus"
	"go.redsock.ru/toolbox/respect"

	"go.vervstack.ru/matreshka/internal/app"
)

func main() {
	println(respect.Fantasy)

	a, err := app.New()
	if err != nil {
		logrus.Fatal(err)
	}

	err = a.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
