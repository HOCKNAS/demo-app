package app

import (
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	usecases "github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Users ports.UsersService
}

type Deps struct {
	Repos            *Repositories
	IdentityProvider *IdentityProvider
	Logger           *logrus.Logger
}

func NewServices(deps Deps) *Services {
	Users := usecases.NewUserService(deps.Repos.Users, deps.IdentityProvider.AuthManager, deps.Logger)

	return &Services{
		Users: Users,
	}
}
