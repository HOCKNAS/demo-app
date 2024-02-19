package http_delivery

import (
	"net/http"

	v1 "github.com/HOCKNAS/demo-app/internal/adapters/delivery/http_delivery/v1"
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *use_cases.Services
	Router   *gin.Engine
}

func NewHandlerHTTP(services *use_cases.Services) *Handler {

	gin.SetMode(gin.ReleaseMode)

	handler := &Handler{
		Services: services,
		Router:   gin.Default(),
	}

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

	config := huma.DefaultConfig("DEMO-APP", "1.0.0")

	huma := humagin.New(h.Router, config)

	handlerV1 := v1.NewHandlerV1(h.Services, huma)

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
