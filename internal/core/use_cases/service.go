package use_cases

import (
	identityprovider "github.com/HOCKNAS/demo-app/internal/adapters/identity_provider"
	"github.com/HOCKNAS/demo-app/internal/adapters/repository"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type Services struct {
	Users ports.UsersService
}

type Deps struct {
	Repos            *repository.Repositories
	IdentityProvider *identityprovider.IdentityProvider
	Logs             ports.Logger
}

func NewServices(deps Deps) *Services {
	Users := NewUserService(deps.Repos.UsersRepository, deps.IdentityProvider.AuthManager, deps.Logs)
	return &Services{
		Users: Users,
	}
}
