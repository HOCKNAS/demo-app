package domain

import "errors"

var (
	ErrUserAlreadyExistsDB        = errors.New("El usuario ya existe en la Base de Datos")
	ErrEmailAlreadyExistsIdP      = errors.New("El correo ya está registrado en el proveedor de identidad")
	ErrUserNotFoundForDeletionDB  = errors.New("No se encontró el usuario para eliminar en la Base de Datos")
	ErrUserNotFoundForDeletionIdP = errors.New("No se encontró el usuario para eliminar en el proveedor de identidad")
	ErrCreationFailedIdP          = errors.New("Falló la creación en el proveedor de identidad")
	ErrDeletionFailedIdP          = errors.New("Falló la eliminación en el proveedor de identidad, se requiere acción manual")
)
