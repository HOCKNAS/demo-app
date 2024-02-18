package usecases

import (
	"context"

	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type userService struct {
	repository        ports.UsersRepository
	identity_provider ports.AuthManager
	logger            ports.Logger
}

func NewUserService(repository ports.UsersRepository, identity_provider ports.AuthManager, logger ports.Logger) *userService {
	return &userService{
		repository:        repository,
		identity_provider: identity_provider,
		logger:            logger,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userDatabase, err := s.repository.Create(ctx, user)
	if err != nil {
		s.logger.Error("Error al crear el usuario en el Repositorio", "error", err, "userID", user.ID)
		return nil, err
	}

	err = s.identity_provider.Create(ctx, userDatabase)
	if err != nil {
		deleteErr := s.repository.Delete(ctx, userDatabase.ID)
		if deleteErr != nil {
			s.logger.Error("Error al intentar compensar el error de creación del usuario en el Proveedor de Identidad", "error", err, "userID", user.ID)
			return nil, deleteErr
		}
		s.logger.Error("Error al crear el usuario en el Proveedor de Identidad", "error", err, "userID", user.ID)
		return nil, err
	}

	s.logger.Info("Usuario creado exitosamente en el Repositorio y el Proveedor de Identidad", "userID", user.ID)
	return userDatabase, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {

	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Error al eliminar el usuario en el Repositorio", "error", err, "userID", id)
		return err
	}

	err = s.identity_provider.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Error al eliminar el usuario en el Proveedor de Identidad", "error", err, "userID", id)
		return err
	}

	s.logger.Info("Usuario eliminado exitosamente del Repositorio y del Proveedor de Identidad", "userID", id)
	return nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	userDatabase, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Error al obtener el usuario desde el Repositorio", "error", err, "userID", id)
		return nil, err
	}

	s.logger.Info("Se obtuvo exitosamente el usuario desde el Repositorio", "userID", id)
	return userDatabase, nil
}

func (s *userService) DeactivateUser(ctx context.Context, id string) (*domain.User, error) {
	userDatabase, repoErr := s.repository.Deactivate(ctx, id)
	if repoErr != nil {
		s.logger.Error("Error al desactivar el usuario en el Repositorio", "error", repoErr, "userID", id)
		return nil, repoErr
	}

	idpErr := s.identity_provider.Deactivate(ctx, id)
	if idpErr != nil {
		s.logger.Error("Error al desactivar el usuario en el Proveedor de Identidad", "error", idpErr, "userID", id)
		return nil, idpErr
	}

	s.logger.Info("Usuario desactivado exitosamente", "userID", id)
	return userDatabase, nil
}
