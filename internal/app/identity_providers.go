package app

import (
	"firebase.google.com/go/auth"
	identityprovider "github.com/HOCKNAS/demo-app/internal/adapters/identity_provider"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type IdentityProvider struct {
	Users ports.AuthManager
}

func NewIdentityProviders(authClient *auth.Client) *IdentityProvider {
	return &IdentityProvider{
		Users: identityprovider.NewIdentityProvider(authClient),
	}
}
