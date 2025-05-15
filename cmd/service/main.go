package main

import (
	"github.com/sirupsen/logrus"

	"go.vervstack.ru/matreshka/internal/app"
)

func main() {
	println(`
▄▖            ▌       
▙▌▛▌▌▌▌█▌▛▘█▌▛▌       
▌ ▙▌▚▚▘▙▖▌ ▙▖▙▌       
                      
  ▄     ▄▖   ▌      ▌ 
  ▙▘▌▌  ▙▘█▌▛▌▛▘▛▌▛▘▙▘
  ▙▘▙▌  ▌▌▙▖▙▌▄▌▙▌▙▖▛▖
    ▄▌                
`)

	a, err := app.New()
	if err != nil {
		logrus.Fatal(err)
	}

	err = a.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
