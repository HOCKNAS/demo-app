package repository

import (
	"github.com/HOCKNAS/demo-app/internal/adapters/repository/mongodb"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	UsersRepository ports.UsersRepository
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		UsersRepository: mongodb.NewUsersRepository(db),
	}
}
