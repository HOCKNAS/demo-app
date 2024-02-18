package ports

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
)

type UsersService interface {
	CreateAccount(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteAccount(ctx context.Context, id string) error
	DeactivateAccount(ctx context.Context, id string) (*domain.User, error)
}
