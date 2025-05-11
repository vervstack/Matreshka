package web_api

import (
	"fmt"
	"net/http"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

type Server struct {
	apiServer api.MatreshkaBeAPIServer

	mux http.ServeMux
}

func New(apiServer api.MatreshkaBeAPIServer) http.Handler {
	s := Server{
		apiServer: apiServer,
		mux:       http.ServeMux{},
	}

	s.mux.HandleFunc("/", s.GetConfig)

	return &s.mux
}

func (s *Server) GetConfig(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	apiReq := &api.GetConfig_Request{
		ConfigName: req.URL.Path[len("/download/"):],
	}

	v := q.Get("version")
	if v != "" {
		apiReq.Version = &v
	}

	format := q.Get("format")
	switch format {
	case api.Format_name[int32(api.Format_env)]:
		apiReq.Format = api.Format_env

	}

	apiResp, err := s.apiServer.GetConfig(req.Context(), apiReq)
	if err != nil {
		return
	}

	const html = `
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

	resp.Header().Set("Content-Type", "text/html")
	_, _ = resp.Write([]byte(fmt.Sprintf(html, apiReq.ConfigName, apiResp.Config)))
}
