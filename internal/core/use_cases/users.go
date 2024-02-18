package usecases

import (
	"context"
	"fmt"

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

func (service *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	userDatabase, err := service.repository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	err = service.identity_provider.Create(ctx, userDatabase)
	if err != nil {
		deleteErr := service.repository.Delete(ctx, userDatabase.ID)
		if deleteErr != nil {
			return nil, fmt.Errorf("%w; falló la compensación (eliminación del usuario): %v", domain.ErrCreationFailedIdP, deleteErr)
		}
		return nil, err
	}

	return userDatabase, err
}

func (service *userService) DeleteUser(ctx context.Context, id string) error {

	err := service.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	err = service.identity_provider.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrDeletionFailedIdP, err)
	}

	return nil
}

func (service *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	userDatabase, err := service.repository.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return userDatabase, err
}

func (service *userService) DeactivateUser(ctx context.Context, id string) (*domain.User, error) {
	userDatabase, err := service.repository.Deactivate(ctx, id)

	if err != nil {
		return nil, err
	}

	err = service.identity_provider.Deactivate(ctx, id)

	if err != nil {
		return nil, err
	}

	return userDatabase, nil
}
