package ports

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
)

type UsersService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	DeactivateUser(ctx context.Context, id string) (*domain.User, error)
}
