package app

import (
	"github.com/HOCKNAS/demo-app/internal/adapters/repository"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Users ports.UsersRepository
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users: repository.NewUsersRepository(db),
	}
}
