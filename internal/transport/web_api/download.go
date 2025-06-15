package web_api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

const downloadBucketHTML = `
<!DOCTYPE html>
<html>
<head>
  <title>%s</title>
  <meta charset="UTF-8">
	<link rel="icon" href="https://s3-api.redsock.ru/verv/matreshka/icon.png" type="image/x-icon">
</head>
<body>
  <pre>%s</pre>
</body>
</html>`

func (s *Server) GetConfig(c *gin.Context) {
	apiReq := &api.GetConfig_Request{
		ConfigName: c.Param(configNameParam),
		Version:    extractVersion(c),
		Format:     extractFormat(c),
	}

	apiResp, err := s.grpcApiServer.GetConfig(c, apiReq)
	if err != nil {
		return
	}

	c.Data(http.StatusOK, "text/html",
		[]byte(fmt.Sprintf(downloadBucketHTML, apiReq.ConfigName, apiResp.Config)))
}
