package http_delivery

import (
	"io"
	"net/http"
	"strings"

	v1 "github.com/HOCKNAS/demo-app/internal/adapters/delivery/http_delivery/v1"
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var docs = strings.Replace(`
You can run it localy via Docker:

^^^sh
# Start the server
$ docker run  -p 8080:8080 demo-app:latest 

# Make a request
$ curl localhost:8080/ping
^^^
`, "^", "`", -1)

type Handler struct {
	Services *use_cases.Services
	Router   *gin.Engine
	api      huma.API
}

func NewHandlerHTTP(services *use_cases.Services) *Handler {

	gin.SetMode(gin.ReleaseMode)

	handler := &Handler{
		Services: services,
		Router:   gin.Default(),
	}

	config := huma.DefaultConfig("DEMO-APP", "1.0.0")
	config.Info.Description = docs
	config.Servers = []*huma.Server{
		{URL: "http://localhost:8080"},
	}

	yamlFormat := huma.Format{
		Marshal: func(writer io.Writer, v any) error {
			return yaml.NewEncoder(writer).Encode(v)
		},
		Unmarshal: func(data []byte, v any) error {
			return yaml.Unmarshal(data, v)
		},
	}
	config.Formats["application/yaml"] = yamlFormat
	config.Formats["yaml"] = yamlFormat

	handler.api = humagin.New(handler.Router, config)
	handler.initMiddleware()
	handler.initAPI()

	return handler
}

func (h *Handler) initMiddleware() {
	h.Router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
}

func (h *Handler) initAPI() {

	handlerV1 := v1.NewHandler(h.Services, &h.api)

	h.Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	rootPath := "/api"
	{
		handlerV1.Init(rootPath)
	}

}

func (h *Handler) GetRouter() *gin.Engine {
	return h.Router
}

func (h *Handler) GetApi() huma.API {
	return h.api
}
