package ports

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
