package v1

import (
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/autopatch"
)

type Handler struct {
	Services *use_cases.Services
	Huma     *huma.API
	Logger   *ports.Logger
}

type APIServer struct {
	RootPath string
	Handler  huma.API
	Logger   ports.Logger
}

func NewHandler(services *use_cases.Services, huma *huma.API, logger *ports.Logger) *Handler {
	return &Handler{
		Services: services,
		Huma:     huma,
		Logger:   logger,
	}
}

func (h *Handler) Init(rootPath string) {

	api := APIServer{
		RootPath: rootPath + "/v1",
		Handler:  *h.Huma,
		Logger:   *h.Logger,
	}

	huma.AutoRegister(*h.Huma, &api)
	autopatch.AutoPatch(*h.Huma)

	{
		h.initUsersRoutes(api)
	}
}
