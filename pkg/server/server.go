package server

import (
	"net/http"

	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/HOCKNAS/demo-app/pkg/server/http_server"
)

type Servers struct {
	HTTPServer ports.HTTPServer
}

func NewServer(http *http.Server) *Servers {
	return &Servers{
		HTTPServer: http_server.NewHTTPServer(http),
	}
}
