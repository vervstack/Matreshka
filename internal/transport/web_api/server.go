package web_api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
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

	r.GET("/download/:config_name", s.GetConfig)
	r.POST("/upload/:config_name", s.UploadConfig)

	s.mux.Handle("/", r)

	return &s.mux
}
