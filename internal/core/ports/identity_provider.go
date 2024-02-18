package ports

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
)

type AuthManager interface {
	Create(ctx context.Context, user *domain.User) error
}
