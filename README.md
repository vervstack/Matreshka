# Mono repository for matreshka configuration system

[![Docker Image Version](https://img.shields.io/docker/v/godverv/velez?style=for-the-badge&logo=docker&label=Velez%20image&labelColor=white&color=blue)](https://hub.docker.com/r/vervstack/matreshka/tags)

## Matreshka configuration module
Are at /pkg/matreshka

Here is how you can use it in your Go project

```go
package main

import (
	"log"

	"go.vervstack.ru/matreshka/pkg/matreshka"
)

func main() {
	cfg, err := matreshka.ReadConfigs("./config.yaml")
	if err != nil {
		log.Fatal("error during config initialization" + err.Error())
	}
	
	log.Print("Running app with version: "+cfg.AppInfo.Version)
}
```


## Matreshka service 
**Public entrypoint** is located at `pkg/app/app.go`
From there you can run your own service in same execution context (for test purposes instead of mocking service)

**Internals** of application is located at internal/app and generated with [rscli](https://github.com/Red-Sock/rscli)

**Binary** is assembled at cmd/service

##### generated with love for coding by [RedSock CLI](https://github.com/Red-Sock/rscli)
