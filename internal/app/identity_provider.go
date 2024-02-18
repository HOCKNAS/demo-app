package app

import (
	"firebase.google.com/go/auth"
	identityprovider "github.com/HOCKNAS/demo-app/internal/adapters/identity_provider"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type IdentityProvider struct {
	AuthManager ports.AuthManager
}

func NewIdentityProvider(authClient *auth.Client) *IdentityProvider {
	return &IdentityProvider{
		AuthManager: identityprovider.NewIdentityProvider(authClient),
	}
}
