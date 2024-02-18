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

func userToFirebase(user *domain.User) *auth.UserToCreate {
	return (&auth.UserToCreate{}).
		UID(user.ID).
		Password(user.Password).
		DisplayName(user.Username).
		Email(user.Email).
		EmailVerified(false)
}

func NewIdentityProvider(authClient *auth.Client) ports.AuthManager {
	return &firebaseIdentityProvider{
		authClient: authClient,
	}
}

func (fip *firebaseIdentityProvider) Create(ctx context.Context, user *domain.User) error {

	params := userToFirebase(user)

	_, err := fip.authClient.CreateUser(ctx, params)

	if err != nil {
		return firebase_auth.ShowError(err)
	}

	return nil
}
