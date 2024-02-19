package server

import (
	"net/http"

	"github.com/HOCKNAS/demo-app/internal/adapters/server/http_server"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type Servers struct {
	HTTPServer ports.HTTPServer
}

func NewServer(http *http.Server) *Servers {
	return &Servers{
		HTTPServer: http_server.NewHTTPServer(http),
	}
}
