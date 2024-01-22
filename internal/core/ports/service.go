package ports

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
)

type UsersService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
}
