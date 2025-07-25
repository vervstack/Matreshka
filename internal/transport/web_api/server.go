package web_api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "go.vervstack.ru/matreshka/pkg/matreshka_api"
)

const (
	configNameParam = "config_name"
	versionParam    = "version"
	formatParam     = "format"
)

type Server struct {
	grpcApiServer api.MatreshkaBeAPIServer

	mux http.ServeMux
}

func New(apiServer api.MatreshkaBeAPIServer) http.Handler {
	s := Server{
		grpcApiServer: apiServer,
		mux:           http.ServeMux{},
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/web_api/download/:config_name", s.GetConfig)
	r.POST("/web_api/upload/:config_name", s.UploadConfig)

	s.mux.Handle("/web_api/", r)

	return &s.mux
}
