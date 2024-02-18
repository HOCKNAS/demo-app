package domain

import "errors"

var (
	ErrUserAlreadyExistsDB   = errors.New("El usuario ya existe en la DB")
	ErrEmailAlreadyExistsIdP = errors.New("El correo ya está registrado en el IdP")
)
