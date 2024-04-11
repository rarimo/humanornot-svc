package identityproviders

import (
	"github.com/rarimo/humanornot-svc/internal/data"
	"github.com/rarimo/humanornot-svc/internal/service/core/issuer"
)

type IdentityProvider interface {
	Verify(user *data.User, verifyProviderDataRaw []byte) (*issuer.IdentityProvidersCredentialSubject, []byte, error)
}
