package v1

import (
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/danielgtaylor/huma/v2"
)

type HandlerV1 struct {
	services *use_cases.Services
	huma     huma.API
}

func NewHandlerV1(services *use_cases.Services, huma huma.API) *HandlerV1 {
	return &HandlerV1{
		services: services,
		huma:     huma,
	}
}

func (h *HandlerV1) Init(rootPath string) {

	s := APIServer{
		RootPath: rootPath + "/v1",
	}

	{
		s.Greeting(h.huma)
	}
}
