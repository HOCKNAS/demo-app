package http_server

import (
	"context"
	"net/http"

	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type server struct {
	http *http.Server
}

func NewHTTPServer(http *http.Server) ports.HTTPServer {
	return &server{
		http: http,
	}
}

func (s *server) Start() error {
	return s.http.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
