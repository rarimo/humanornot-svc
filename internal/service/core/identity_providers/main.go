package identityproviders

import (
	"github.com/rarimo/kyc-service/internal/data"
	"github.com/rarimo/kyc-service/internal/service/core/issuer"
)

type IdentityProvider interface {
	Verify(user *data.User, verifyProviderDataRaw []byte) (*issuer.IdentityProvidersCredentialSubject, []byte, error)
}
