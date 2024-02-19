package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type APIServer struct {
	RootPath string
}

type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func (s *APIServer) Greeting(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Summary:     "Get a greeting",
		Method:      http.MethodGet,
		Path:        s.RootPath + "/greeting/{name}",
	}, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})
}
