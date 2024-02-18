package firebase_auth

import (
	"context"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/HOCKNAS/demo-app/internal/core/domain"

	"google.golang.org/api/option"
)

const timeout = 10 * time.Second

func NewApp(credentials string) (*firebase.App, error) {

	opt := option.WithCredentialsFile(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func ShowError(err error) error {
	if auth.IsEmailAlreadyExists(err) {
		return domain.ErrEmailAlreadyExistsIdP
	}

	return err
}
