package ports

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
)

type AuthManager interface {
	Create(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}
