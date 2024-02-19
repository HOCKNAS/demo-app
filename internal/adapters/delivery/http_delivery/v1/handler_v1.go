package v1

import (
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/danielgtaylor/huma/v2"
)

type Handler struct {
	Services *use_cases.Services
	Huma     *huma.API
}

type APIServer struct {
	RootPath string
	Handler  huma.API
}

func NewHandler(services *use_cases.Services, huma *huma.API) *Handler {
	return &Handler{
		Services: services,
		Huma:     huma,
	}
}

func (h *Handler) Init(rootPath string) {

	api := APIServer{
		RootPath: rootPath + "/v1",
		Handler:  *h.Huma,
	}

	{
		h.initUsersRoutes(api)
	}
}
