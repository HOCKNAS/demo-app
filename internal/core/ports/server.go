package ports

import (
	"context"
)

type HTTPServer interface {
	Start() error
	Stop(ctx context.Context) error
}
