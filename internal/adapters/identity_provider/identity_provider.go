package identity_provider

import (
	"firebase.google.com/go/auth"
	"github.com/HOCKNAS/demo-app/internal/adapters/identity_provider/firebase"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
)

type IdentityProvider struct {
	AuthManager ports.AuthManager
}

func NewIdentityProvider(authClient *auth.Client) *IdentityProvider {
	return &IdentityProvider{
		AuthManager: firebase.NewFirebaseIdentityProvider(authClient),
	}
}
