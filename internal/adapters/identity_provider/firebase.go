package identityprovider

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/HOCKNAS/demo-app/pkg/auth/firebase_auth"
)

type firebaseIdentityProvider struct {
	authClient *auth.Client
}

func NewIdentityProvider(authClient *auth.Client) ports.AuthManager {
	return &firebaseIdentityProvider{
		authClient: authClient,
	}
}

func (fip *firebaseIdentityProvider) Create(ctx context.Context, user *domain.User) error {

	params := (&auth.UserToCreate{}).
		UID(user.ID).
		Password(user.Password).
		DisplayName(user.Username).
		Email(user.Email).
		EmailVerified(false)

	_, err := fip.authClient.CreateUser(ctx, params)

	if err != nil {
		return firebase_auth.ShowError(err)
	}

	return nil
}

func (fip *firebaseIdentityProvider) Delete(ctx context.Context, id string) error {

	err := fip.authClient.DeleteUser(ctx, id)

	if err != nil {
		return firebase_auth.ShowError(err)
	}

	return nil
}

func (fip *firebaseIdentityProvider) GetByID(ctx context.Context, id string) error {

	_, err := fip.authClient.GetUser(ctx, id)

	if err != nil {
		return firebase_auth.ShowError(err)
	}

	return nil
}

func (fip *firebaseIdentityProvider) Deactivate(ctx context.Context, id string) error {
	params := (&auth.UserToUpdate{}).
		Disabled(true)

	_, err := fip.authClient.UpdateUser(ctx, id, params)

	if err != nil {
		return firebase_auth.ShowError(err)
	}

	return nil
}
