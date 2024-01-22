package domain

type User struct {
	ID           string
	Name         string
	LastName     string
	Username     string
	Password     string
	Email        string
	IsAdmin      bool
	IsActive     bool
	CreationDate string
}
