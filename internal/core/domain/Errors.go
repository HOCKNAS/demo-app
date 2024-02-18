package domain

import "errors"

var (
	ErrUserAlreadyExistsDB        = errors.New("El usuario ya existe en la DB")
	ErrEmailAlreadyExistsIdP      = errors.New("El correo ya está registrado en el IdP")
	ErrUserNotFoundForDeletionDB  = errors.New("No se encontró el usuario para eliminar en la DB")
	ErrUserNotFoundForDeletionIdP = errors.New("No se encontró el usuario para eliminar en el IdP")
)
