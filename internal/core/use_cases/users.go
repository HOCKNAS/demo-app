package usecases

import (
	"context"
	"errors"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type userService struct {
	repository ports.UsersRepository
}

func NewUserService(repository ports.UsersRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) Register(ctx context.Context, input *domain.User) (*domain.User, error) {

	userDatabase, err := service.repository.Create(ctx, input)

	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return nil, err
		}

		return nil, err
	}

	return userDatabase, err
}
