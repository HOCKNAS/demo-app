package usecases

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type userService struct {
	repository        ports.UsersRepository
	identity_provider ports.AuthManager
}

func NewUserService(repository ports.UsersRepository, identity_provider ports.AuthManager) *userService {
	return &userService{
		repository:        repository,
		identity_provider: identity_provider,
	}
}

func (service *userService) CreateAccount(ctx context.Context, input *domain.User) (*domain.User, error) {

	userDatabase, err := service.repository.Create(ctx, input)

	if err != nil {
		return nil, err
	}

	err = service.identity_provider.Create(ctx, userDatabase)

	return userDatabase, err
}

func (service *userService) DeleteAccount(ctx context.Context, id string) error {

	err := service.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	err = service.identity_provider.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
