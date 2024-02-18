package usecases

import (
	"context"
	"fmt"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/sirupsen/logrus"
)

type userService struct {
	repository        ports.UsersRepository
	identity_provider ports.AuthManager
	logger            *logrus.Logger
}

func NewUserService(repository ports.UsersRepository, identity_provider ports.AuthManager, logger *logrus.Logger) *userService {
	return &userService{
		repository:        repository,
		identity_provider: identity_provider,
		logger:            logger,
	}
}

func (service *userService) logError(contextInfo logrus.Fields, err error) error {
	message := err.Error()
	service.logger.WithFields(contextInfo).Error(message)
	return fmt.Errorf("%s", message)
}

func (service *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userDatabase, err := service.repository.Create(ctx, user)
	if err != nil {
		return nil, service.logError(logrus.Fields{"userID": user.ID}, err)
	}

	err = service.identity_provider.Create(ctx, userDatabase)
	if err != nil {
		deleteErr := service.repository.Delete(ctx, userDatabase.ID)
		if deleteErr != nil {
			return nil, service.logError(logrus.Fields{"userID": user.ID}, deleteErr)
		}
		return nil, service.logError(logrus.Fields{"userID": user.ID}, err)
	}

	service.logger.WithFields(logrus.Fields{"userID": user.ID}).Info("Usuario creado exitosamente en repositorio y proveedor de identidad")
	return userDatabase, nil
}

func (service *userService) DeleteUser(ctx context.Context, id string) error {

	err := service.repository.Delete(ctx, id)
	if err != nil {
		return service.logError(logrus.Fields{"userID": id}, err)
	}

	err = service.identity_provider.Delete(ctx, id)
	if err != nil {
		return service.logError(logrus.Fields{"userID": id}, err)
	}

	service.logger.WithFields(logrus.Fields{"userID": id}).Info("Usuario eliminado exitosamente de repositorio y proveedor de identidad")
	return nil
}

func (service *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	userDatabase, err := service.repository.GetByID(ctx, id)
	if err != nil {
		return nil, service.logError(logrus.Fields{"userID": id}, err)
	}

	return userDatabase, nil
}

func (service *userService) DeactivateUser(ctx context.Context, id string) (*domain.User, error) {
	userDatabase, repoErr := service.repository.Deactivate(ctx, id)
	if repoErr != nil {
		return nil, service.logError(logrus.Fields{"userID": id}, repoErr)
	}

	idpErr := service.identity_provider.Deactivate(ctx, id)
	if idpErr != nil {
		return nil, service.logError(logrus.Fields{"userID": id}, idpErr)
	}

	service.logger.WithFields(logrus.Fields{"userID": id}).Info("Usuario desactivado exitosamente")
	return userDatabase, nil
}
