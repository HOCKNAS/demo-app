package v1

import (
	"context"
	"net/http"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/danielgtaylor/huma/v2"
)

type SignUpInputBody struct {
	Name     string `maxLength:"30" example:"Santiago" doc:"Name to user"`
	LastName string `maxLength:"30" example:"Chacon" doc:"Last Name to user"`
	Username string `maxLength:"30" example:"Hocknas" doc:"Username to user"`
	Email    string `maxLength:"30" example:"santiago.chacon99@gmail.com" doc:"Email to user"`
	Password string `maxLength:"30" example:"Ab12345@" doc:"Password to user"`
}

type SignUpInput struct {
	Body    SignUpInputBody
	RawBody []byte
}

type SignUpOutputBody struct {
	ID string
}

type SignUpOutput struct {
	Body SignUpOutputBody
}

func (h *Handler) initUsersRoutes(api APIServer) {
	api.RootPath += "/users"
	api.SignUp(api.Handler, h.Services.Users)
}

func (s *APIServer) SignUp(api huma.API, service ports.UsersService) {

	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Summary:     "Create a new user",
		Method:      http.MethodPost,
		Path:        s.RootPath + "/signup",
	}, func(ctx context.Context, input *SignUpInput) (*SignUpOutput, error) {

		newUser := &domain.User{
			Name:     input.Body.Name,
			LastName: input.Body.LastName,
			Username: input.Body.Username,
			Email:    input.Body.Email,
			Password: input.Body.Password,
			IsAdmin:  false,
			IsActive: true,
		}

		user, err := service.CreateUser(ctx, newUser)
		if err != nil {
			s.Logger.Error("Error creando el usuario", "error", err)
			return nil, err
		}

		return &SignUpOutput{
			Body: SignUpOutputBody{
				ID: user.ID,
			},
		}, nil
	})
}
