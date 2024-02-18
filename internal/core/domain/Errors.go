package domain

import "errors"

var (
	ErrUserAlreadyExistsDB        = errors.New("El usuario ya existe en la Base de Datos")
	ErrEmailAlreadyExistsIdP      = errors.New("El correo ya está registrado en el Proveedor de Identidad")
	ErrUserNotFoundForDeletionDB  = errors.New("No se encontró el usuario para eliminar en la Base de Datos")
	ErrUserNotFoundForDeletionIdP = errors.New("No se encontró el usuario para eliminar en el Proveedor de Identidad")
	ErrCreationFailedIdP          = errors.New("Falló la creación en el Proveedor de Identidad")
	ErrDeletionFailedIdP          = errors.New("Falló la eliminación en el Proveedor de Identidad, se requiere acción manual")
	ErrUserNotFoundDB             = errors.New("No se encontró el usuario en la Base de Datos")
	ErrDatabase                   = errors.New("Error en la Base de Datos")
	ErrUserAlreadyDeactivatedIdP  = errors.New("El usuario ya está desactivado en el Proveedor de Identidad")
	ErrUserAlreadyDeactivatedDB   = errors.New("El usuario ya está desactivado en la Base de Datos")
)
