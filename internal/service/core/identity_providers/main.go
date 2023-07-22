package identityproviders

import (
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/issuer"
)

type IdentityProvider interface {
	Verify(user *data.User, verifyProviderDataRaw []byte) (*issuer.IdentityProvidersCredentialSubject, []byte, error)
}
